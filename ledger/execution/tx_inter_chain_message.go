package execution

import (
	"fmt"
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/result"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/store/database"

	sbc "github.com/thetatoken/thetasubchain/blockchain"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/ledger/execution/interchain"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
	stypes "github.com/thetatoken/thetasubchain/ledger/types"
	svm "github.com/thetatoken/thetasubchain/ledger/vm"
	"github.com/thetatoken/thetasubchain/witness"
)

var _ TxExecutor = (*InterChainMessageTxExecutor)(nil)

// ------------------------------- InterChainMessage Transaction -----------------------------------

// InterChainMessageTxExecutor implements the TxExecutor interface
type InterChainMessageTxExecutor struct {
	db               database.Database
	chain            *sbc.Chain
	state            *slst.LedgerState
	consensus        score.ConsensusEngine
	valMgr           score.ValidatorManager
	mainchainWitness witness.ChainWitness
}

// NewInterChainMessageTxExecutor creates a new instance of InterChainMessageTxExecutor
func NewInterChainMessageTxExecutor(db database.Database, chain *sbc.Chain, state *slst.LedgerState, consensus score.ConsensusEngine,
	valMgr score.ValidatorManager, mainchainWitness witness.ChainWitness) *InterChainMessageTxExecutor {
	return &InterChainMessageTxExecutor{
		db:               db,
		chain:            chain,
		state:            state,
		consensus:        consensus,
		valMgr:           valMgr,
		mainchainWitness: mainchainWitness,
	}
}

func (exec *InterChainMessageTxExecutor) sanityCheck(chainID string, view *slst.StoreView, transaction types.Tx) result.Result {
	// basic checks, and verify the tx proposer is a validator
	tx := transaction.(*stypes.InterChainMessageTx)
	validatorSet := getValidatorSet(exec.consensus.GetLedger(), exec.valMgr)
	validatorAddresses := getValidatorAddresses(validatorSet)

	// Validate proposer, basic
	res := tx.Proposer.ValidateBasic()
	if res.IsError() {
		return res
	}

	// verify the proposer is one of the validators
	res = isAValidator(tx.Proposer.Address, validatorAddresses)
	if res.IsError() {
		return res
	}

	icmNonce := tx.Event.Nonce
	if view.InterChainMessageTransactionProcessed(icmNonce) {
		return result.Error("inter-chain message %v has already been processed", icmNonce)
	}

	// verify the proposer's signature
	proposerAccount, res := getOrMakeInput(view, tx.Proposer)
	if res.IsError() {
		return res
	}

	signBytes := tx.SignBytes(chainID)
	if !tx.Proposer.Signature.Verify(signBytes, proposerAccount.Address) {
		return result.Error("SignBytes: %X", signBytes)
	}

	return result.OK
}

func (exec *InterChainMessageTxExecutor) process(chainID string, view *slst.StoreView, transaction types.Tx) (common.Hash, result.Result) {
	tx := transaction.(*stypes.InterChainMessageTx)
	icmNonce := tx.Event.Nonce

	if !view.ShouldProcessThisInterChainMessageEvent(icmNonce) {
		return common.Hash{}, result.Error("inter-chain message nonce %v mismatches with the expected nonce %v", icmNonce, view.GetLastProcessedEventNonce())
	}

	eventCache := exec.mainchainWitness.GetInterChainEventCache()
	witnessedEvent, err := eventCache.Get(tx.Event.Type, icmNonce)
	if err != nil || witnessedEvent == nil {
		return common.Hash{}, result.UndecidedWith(result.Info{"event not seen yet": tx.Event, "err": err}) // not seen on mainchain yet
	}

	if view.InterChainMessageTransactionProcessed(icmNonce) {
		return common.Hash{}, result.Error("inter-chain message with nonce %v has already been processed", icmNonce)
	}

	// the leader could be malicious, so we need to verify if the inter-chain message event proposed by the leader is consistent with the query results
	if !witnessedEvent.Equals(&tx.Event) {
		return common.Hash{}, result.Error("event mismatch, inter-chain message verification failed for ID %v", icmNonce)
	}

	pb := exec.state.ParentBlock()
	parentBlockInfo := svm.NewBlockInfo(pb.Height, pb.Timestamp, pb.ChainID)
	proxySctx, err := exec.constructProxySctx(view, tx)
	if err != nil {
		return common.Hash{}, result.Error("error constructing proxy mint voucher sctx: %v", err)
	}

	evmRet, contractAddr, gasUsed, evmErr := svm.Execute(parentBlockInfo, proxySctx, view, svm.PreviledgedAccess)

	// TODO: Do we need to increase the sequence number of the proposer?
	proposerAddress := tx.Proposer.Address
	proposerAccount, success := getInput(view, tx.Proposer)
	if success.IsError() {
		logger.Fatalf("Failed to get proposer account: %v", success) // should never happen, since the proposer account must exist
	}

	createContract := (proxySctx.To.Address == common.Address{})
	if !createContract { // svm.create() increments the sequence of the from account
		proposerAccount.Sequence++
	}
	view.SetAccount(proposerAddress, proposerAccount)

	// TODO: Add tx receipt: status and events
	logs := view.PopLogs()
	if evmErr != nil {
		// Do not record events if transaction is reverted
		logs = nil
	}
	exec.chain.AddTxReceipt(tx, logs, evmRet, contractAddr, gasUsed, evmErr)

	view.SetInterChainMessageTransactionProcessed(icmNonce)
	txHash := types.TxID(chainID, tx)
	return txHash, result.OK
}

func (exec *InterChainMessageTxExecutor) constructProxySctx(view *slst.StoreView, tx *stypes.InterChainMessageTx) (*types.SmartContractTx, error) {
	var proxySctx *types.SmartContractTx
	var err error

	eventType := tx.Event.Type
	switch eventType {
	case score.IMCEventTypeCrossChainTFuelTransfer:
		proxySctx, err = interchain.ConstructMintTFuelVoucherProxySctx(tx.Proposer.Address, view, &tx.Event)
	case score.IMCEventTypeCrossChainTNT20Transfer:
		proxySctx, err = interchain.ConstructMintTNT20VoucherProxySctx(tx.Proposer.Address, view, &tx.Event)
	case score.IMCEventTypeCrossChainTNT721Transfer:
		proxySctx, err = interchain.ConstructMintTNT721VoucherProxySctx(tx.Proposer.Address, view, &tx.Event)
	default:
		return nil, fmt.Errorf("unsupported event type: %v", eventType)
	}

	if err != nil {
		return nil, err
	}

	return proxySctx, nil
}

func (exec *InterChainMessageTxExecutor) getTxInfo(transaction types.Tx) *score.TxInfo {
	return &score.TxInfo{
		EffectiveGasPrice: exec.calculateEffectiveGasPrice(transaction),
	}
}

func (exec *InterChainMessageTxExecutor) calculateEffectiveGasPrice(transaction types.Tx) *big.Int {
	return new(big.Int).SetUint64(0)
}
