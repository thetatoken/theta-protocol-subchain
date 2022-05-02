package execution

import (
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/result"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/store/database"

	sbc "github.com/thetatoken/thetasubchain/blockchain"
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

	// crossChainTransferID := exec.state.GetCrossChainTransferID(tx)
	crossChainTransferID := tx.Event.Nonce
	if view.CrossChainTransferTransactionProcessed(crossChainTransferID) {
		return result.Error("cross-chain chain transfer %v has already been processed", crossChainTransferID)
	}

	// smartContracTx, err := constructSmartContractTx(tx) // construct a smart contract tx from the special tx (proposed and siged by the leader)
	// if err != nil {
	// 	return result.Error("error constructing smart contract tx: %v", err)
	// }

	// Other checks, e.g. signature check..

	proposerAccount, res := getOrMakeInput(view, tx.Proposer)
	if res.IsError() {
		return res
	}

	// verify the proposer's signature
	signBytes := tx.SignBytes(chainID)
	if !tx.Proposer.Signature.Verify(signBytes, proposerAccount.Address) {
		return result.Error("SignBytes: %X", signBytes)
	}

	// _, result := exec.smartContractExecutor.sanityCheck(chainID, view, smartContractTx)
	// if result.IsError() {
	// 	return result
	// }

	return result.OK
}

func (exec *CrossChainTransferTxExecutor) process(chainID string, view *slst.StoreView, transaction types.Tx) (common.Hash, result.Result) {
	tx := transaction.(*stypes.CrossChainTransferTx)
	// crossChainTransferID := exec.state.GetCrossChainTransferID(tx)
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

	// smartContracTx, err := constructSmartContractTx(tx) // construct a smart contract tx from the special tx (proposed and siged by the leader)
	// if err != nil {
	// 	return common.Hash{}, result.Error("error constructing smart contract tx: %v", err)
	// }

	// mainchainTokenContractAddress := tx.GetSourceTokenContractAddrOnMainChain()
	// isContractDeployment := (smartContracTx.To.Address == commmon.Address{})
	// tokenVoucherContractAddress := view.GetTokenVoucherContract(mainchainTokenContractAddress)
	// if isContractDeployment && (tokenVoucherContractAddress != commmon.Address{}) {
	// 	// the token voucher contract should not have been deployed, otherwise the tx is malicious
	// 	return result.Error("token voucher contract for %v has already been deployed at %v, the tx is invalid", mainchainTokenContractAddress.Hex(), tokenVoucherContractAddress.Hex())
	// }
	// if !isContractDeployment && (tokenVoucherContractAddress != smartContracTx.To) {
	// 	return result.Error("token voucher contract mismatch, %v vs %v", tokenVoucherContractAddress.Hex(), smartContracTx.To.Hex())
	// }

	// _, result := exec.smartContractExecutor.process(chainID, view, smartContractTx)
	// if result.IsError() {
	// 	return result
	// }

	// if isContractDeployment {
	// 	tokenVoucherContractAddress = (result.Info[contractAddrInfoKey]).(commmon.Address)
	// 	view.SetTokenVoucherContract(mainchainTokenContractAddress, tokenVoucherContractAddress)
	// }

	// sourceTokenContractAddrOnMainChain := xferDetails.SourceTokenContractAddrOnMainChain
	// if sourceTokenContractAddrOnMainChain == nil

	view.SetCrossChainTransferTransactionProcessed(crossChainTransferID)
	txHash := types.TxID(chainID, tx)
	return txHash, result.OK
}

// to be called by Ledger to compose the special tx for cross-chain transfers
// func GenerateCrossChainTransferTx(proposerAddress common.Address, view *slst.StoreView) (*types.SmartContractTx, error) {
// 	xferDetails, err := exec.mainchainWitness.GetCrossChainTransfer(crossChainTransferID)
// 	if err != nil || xferDetails == nil {
// 		return nil, err
// 	}

// 	proposerSeq := getProposerSequence(proposerAddress)

// 	from := types.TxInput{
// 		Address: proposerAddr,
// 		Coins: types.Coins{
// 			ThetaWei: new(big.Int).SetUint64(0),
// 			TFuelWei: new(big.Int).SetUint64(0),
// 		},
// 		Sequence: proposerSeq + 1,
// 	}

// 	sourceTokenContractAddrOnMainChain := xferDetails.SourceTokenContractAddrOnMainChain
// 	tokenVoucherContract := view.GetTokenVoucherContract(sourceTokenContractAddrOnMainChain)
// 	var to types.TxOutput
// 	var data []byte
// 	if tokenVoucherContract == nil { // token voucher contract not yet exists, should deploy and mint
// 		// generate a tx with "deploy and mint" data field
// 		to.Address = commmon.Address{}
// 		data = xxxx // need to handle both TFuel and TNT tokens, will add a precomiled contract for minting TFuel
// 	} else { // token voucher contract already exists, should mint the token voucher
// 		// generate a tx with "mint" data field
// 		to.Address = tokenVoucherContract.Address
// 		data = xxxx // need to handle both TFuel and TNT tokens
// 	}

// 	sctx := &types.SmartContractTx{
// 		From:     from,
// 		To:       to,
// 		GasLimit: gasLimit,
// 		GasPrice: gasPrice,
// 		Data:     data,
// 	}

// 	return sctx, nil
// }

func (exec *CrossChainTransferTxExecutor) getTxInfo(transaction types.Tx) *score.TxInfo {
	return &score.TxInfo{
		EffectiveGasPrice: exec.calculateEffectiveGasPrice(transaction),
	}
}

func (exec *CrossChainTransferTxExecutor) calculateEffectiveGasPrice(transaction types.Tx) *big.Int {
	return new(big.Int).SetUint64(0)
}

