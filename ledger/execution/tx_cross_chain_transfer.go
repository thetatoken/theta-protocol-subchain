package execution

import (
	"fmt"
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/result"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/ledger/vm"
	"github.com/thetatoken/theta/store/database"

	sbc "github.com/thetatoken/thetasubchain/blockchain"
	"github.com/thetatoken/thetasubchain/contracts/predeployed"
	"github.com/thetatoken/thetasubchain/core"
	score "github.com/thetatoken/thetasubchain/core"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
	stypes "github.com/thetatoken/thetasubchain/ledger/types"
	"github.com/thetatoken/thetasubchain/witness"
)

var _ TxExecutor = (*CrossChainTransferTxExecutor)(nil)

// ------------------------------- CrossChainTransfer Transaction -----------------------------------

// CrossChainTransferTxExecutor implements the TxExecutor interface
type CrossChainTransferTxExecutor struct {
	db               database.Database
	chain            *sbc.Chain
	state            *slst.LedgerState
	consensus        score.ConsensusEngine
	valMgr           score.ValidatorManager
	mainchainWitness witness.ChainWitness
}

// NewCrossChainTransferTxExecutor creates a new instance of CrossChainTransferTxExecutor
func NewCrossChainTransferTxExecutor(db database.Database, chain *sbc.Chain, state *slst.LedgerState, consensus score.ConsensusEngine,
	valMgr score.ValidatorManager, mainchainWitness witness.ChainWitness) *CrossChainTransferTxExecutor {
	return &CrossChainTransferTxExecutor{
		db:               db,
		chain:            chain,
		state:            state,
		consensus:        consensus,
		valMgr:           valMgr,
		mainchainWitness: mainchainWitness,
	}
}

func (exec *CrossChainTransferTxExecutor) sanityCheck(chainID string, view *slst.StoreView, transaction types.Tx) result.Result {
	// basic checks, and verify the tx proposer is a validator
	tx := transaction.(*stypes.CrossChainTransferTx)
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

	crossChainTransferID := tx.Event.Nonce
	if view.CrossChainTransferTransactionProcessed(crossChainTransferID) {
		return result.Error("cross-chain chain transfer %v has already been processed", crossChainTransferID)
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

func (exec *CrossChainTransferTxExecutor) process(chainID string, view *slst.StoreView, transaction types.Tx) (common.Hash, result.Result) {
	tx := transaction.(*stypes.CrossChainTransferTx)
	crossChainTransferID := tx.Event.Nonce

	if !view.ShouldProcessThisCrossChainTransferEvent(crossChainTransferID) {
		return common.Hash{}, result.Error("cross-chain chain transfer %v is not the next event to process, expected nonce is %v", crossChainTransferID, view.GetLastProcessedEventNonce())
	}

	eventCache := exec.mainchainWitness.GetCrossChainEventCache()
	xferDetails, err := eventCache.Get(crossChainTransferID)
	if err != nil || xferDetails == nil {
		return common.Hash{}, result.UndecidedWith(result.Info{"new event": tx.Event, "err": err}) // not seen on mainchain yet
	}

	if view.CrossChainTransferTransactionProcessed(crossChainTransferID) {
		return common.Hash{}, result.Error("cross-chain chain transfer %v has already been processed", crossChainTransferID)
	}

	// the leader could be malicious, so we need to verify if the cross-chain xfer proposed by the leader is consistent with the query results
	if !xferDetails.Equals(tx.Event) {
		return common.Hash{}, result.Error("cross-chain transfer verification failed for ID %v", crossChainTransferID)
	}

	pb := exec.state.ParentBlock()
	parentBlockInfo := vm.NewBlockInfo(pb.Height, pb.Timestamp, pb.ChainID)

	proxySctx, err := exec.constructProxyMintVoucherSctx(view, tx)
	if err != nil {
		return common.Hash{}, result.Error("error constructing proxy mint voucher sctx: %v", err)
	}

	evmRet, contractAddr, gasUsed, evmErr := vm.Execute(parentBlockInfo, proxySctx, view)

	// TODO: Do we need to increase the sequence number of the proposer?

	// TODO: Add tx receipt: status and events
	logs := view.PopLogs()
	if evmErr != nil {
		// Do not record events if transaction is reverted
		logs = nil
	}
	exec.chain.AddTxReceipt(tx, logs, evmRet, contractAddr, gasUsed, evmErr)

	view.SetCrossChainTransferTransactionProcessed(crossChainTransferID)
	txHash := types.TxID(chainID, tx)
	return txHash, result.OK
}

func (exec *CrossChainTransferTxExecutor) constructProxyMintVoucherSctx(view *slst.StoreView, tx *stypes.CrossChainTransferTx) (*types.SmartContractTx, error) {
	tokenType, err := core.ExtractCrossChainTokenTypeFromDenom(tx.Event.Denom)
	if err != nil || tokenType == core.CrossChainTokenTypeInvalid {
		return nil, fmt.Errorf("failed to extract token type from denom %v: %v", tx.Event.Denom, err)
	}

	sourceChainID, err := core.ExtractSourceChainIDFromDenom(tx.Event.Denom)
	if err != nil {
		return nil, fmt.Errorf("failed to extract source chain ID from denom %v: %v", tx.Event.Denom, err)
	}

	var tokenBank predeployed.TokenBank
	switch tokenType {
	case core.CrossChainTokenTypeTFuel:
		tokenBank = predeployed.NewTFuelTokenBank()
	// case core.CrossChainTokenTypeTNT20:
	// 	tokenBank = predeployed.NewTNT20TokenBank()
	// case core.CrossChainTokenTypeTNT721:
	// 	tokenBank = predeployed.NewTNT721TokenBank()
	default:
		return nil, fmt.Errorf("unsupported token type: %v", tokenType)
	}

	proxySctx, err := tokenBank.GetMintVouchersProxySctx(sourceChainID, view, tx)
	if err != nil {
		return nil, err
	}

	return proxySctx, nil
}

func (exec *CrossChainTransferTxExecutor) getTxInfo(transaction types.Tx) *score.TxInfo {
	return &score.TxInfo{
		EffectiveGasPrice: exec.calculateEffectiveGasPrice(transaction),
	}
}

func (exec *CrossChainTransferTxExecutor) calculateEffectiveGasPrice(transaction types.Tx) *big.Int {
	return new(big.Int).SetUint64(0)
}
