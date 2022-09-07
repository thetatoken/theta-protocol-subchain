package execution

import (
	log "github.com/sirupsen/logrus"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/result"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/store/database"
	sbc "github.com/thetatoken/thetasubchain/blockchain"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/interchain/witness"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
	stypes "github.com/thetatoken/thetasubchain/ledger/types"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "ledger"})

//
// TxExecutor defines the interface of the transaction executors
//
type TxExecutor interface {
	sanityCheck(chainID string, view *slst.StoreView, viewSel score.ViewSelector, transaction types.Tx) result.Result
	process(chainID string, view *slst.StoreView, viewSel score.ViewSelector, transaction types.Tx) (common.Hash, result.Result)
	getTxInfo(transaction types.Tx) *score.TxInfo
}

//
// Executor executes the transactions
//
type Executor struct {
	db        database.Database
	chain     *sbc.Chain
	state     *slst.LedgerState
	consensus score.ConsensusEngine
	valMgr    score.ValidatorManager
	ledger    score.Ledger

	coinbaseTxExec                   *CoinbaseTxExecutor
	subchainValidatorSetUpdateTxExec *SubchainValidatorSetUpdateTxExecutor
	sendTxExec                       *SendTxExecutor
	smartContractTxExec              *SmartContractTxExecutor

	skipSanityCheck bool
}

// NewExecutor creates a new instance of Executor
func NewExecutor(db database.Database, chain *sbc.Chain, state *slst.LedgerState, consensus score.ConsensusEngine,
	valMgr score.ValidatorManager, ledger score.Ledger, metachainWitness witness.ChainWitness) *Executor {
	executor := &Executor{
		db:                               db,
		chain:                            chain,
		state:                            state,
		consensus:                        consensus,
		valMgr:                           valMgr,
		coinbaseTxExec:                   NewCoinbaseTxExecutor(db, chain, state, consensus, valMgr),
		subchainValidatorSetUpdateTxExec: NewSubchainValidatorSetUpdateTxExecutor(db, chain, state, consensus, valMgr, metachainWitness),
		sendTxExec:                       NewSendTxExecutor(state),
		smartContractTxExec:              NewSmartContractTxExecutor(chain, state, ledger, valMgr),
		skipSanityCheck:                  false,
	}

	return executor
}

// SetSkipSanityCheck sets the flag for sanity check.
// Skip checks while replaying commmitted blocks.
func (exec *Executor) SetSkipSanityCheck(skip bool) {
	exec.skipSanityCheck = skip
}

// ExecuteTx executes the given transaction
func (exec *Executor) ExecuteTx(tx types.Tx) (common.Hash, result.Result) {
	return exec.processTx(tx, score.DeliveredView)
}

// CheckTx checks the validity of the given transaction
func (exec *Executor) CheckTx(tx types.Tx) (common.Hash, result.Result) {
	return exec.processTx(tx, score.CheckedView)
}

// ScreenTx checks the validity of the given transaction
func (exec *Executor) ScreenTx(tx types.Tx) (common.Hash, result.Result) {
	return exec.processTx(tx, score.ScreenedView)
}

// GetTxInfo extracts tx information used by mempool to sort Txs.
func (exec *Executor) GetTxInfo(tx types.Tx) (*score.TxInfo, result.Result) {
	txExecutor := exec.getTxExecutor(tx)
	if txExecutor == nil {
		return nil, result.Error("Unknown tx type")
	}

	txInfo := txExecutor.getTxInfo(tx)
	return txInfo, result.OK
}

// processTx contains the main logic to process the transaction. If the tx is invalid, a TMSP error will be returned.
func (exec *Executor) processTx(tx types.Tx, viewSel score.ViewSelector) (common.Hash, result.Result) {
	chainID := exec.state.GetChainID()
	var view *slst.StoreView
	switch viewSel {
	case score.DeliveredView:
		view = exec.state.Delivered()
	case score.CheckedView:
		view = exec.state.Checked()
	default:
		view = exec.state.Screened()
	}

	sanityCheckResult := exec.sanityCheck(chainID, view, viewSel, tx)
	if sanityCheckResult.IsError() {
		return common.Hash{}, sanityCheckResult
	}
	if sanityCheckResult.IsUndecided() {
		return common.Hash{}, sanityCheckResult
	}

	txHash, processResult := exec.process(chainID, view, viewSel, tx)
	return txHash, processResult
}

func (exec *Executor) sanityCheck(chainID string, view *slst.StoreView, viewSel score.ViewSelector, tx types.Tx) result.Result {
	if exec.skipSanityCheck { // Skip checks, e.g. while replaying commmitted blocks.
		return result.OK
	}

	if !exec.isTxTypeSupported(view, tx) {
		return result.Error("tx type not supported yet")
	}

	var sanityCheckResult result.Result
	txExecutor := exec.getTxExecutor(tx)
	if txExecutor != nil {
		sanityCheckResult = txExecutor.sanityCheck(chainID, view, viewSel, tx)
	} else {
		sanityCheckResult = result.Error("Unknown tx type")
	}

	return sanityCheckResult
}

func (exec *Executor) process(chainID string, view *slst.StoreView, viewSel score.ViewSelector, tx types.Tx) (common.Hash, result.Result) {
	var processResult result.Result
	var txHash common.Hash

	if !exec.isTxTypeSupported(view, tx) {
		return txHash, result.Error("tx type not supported yet")
	}

	txExecutor := exec.getTxExecutor(tx)
	if txExecutor != nil {
		txHash, processResult = txExecutor.process(chainID, view, viewSel, tx)
		if processResult.IsError() {
			logger.Warnf("Tx processing error: %v", processResult.Message)
		}
	} else {
		processResult = result.Error("Unknown tx type")
	}

	return txHash, processResult
}

func (exec *Executor) isTxTypeSupported(view *slst.StoreView, tx types.Tx) bool {
	blockHeight := view.Height() + 1

	switch tx.(type) {
	case *types.SmartContractTx:
		if blockHeight < common.HeightEnableSmartContract {
			return false
		}
	case *types.StakeRewardDistributionTx:
		if blockHeight < common.HeightEnableTheta3 {
			return false
		}
	default:
		return true
	}

	return true
}

func (exec *Executor) getTxExecutor(tx types.Tx) TxExecutor {
	var txExecutor TxExecutor
	switch tx.(type) {
	case *types.CoinbaseTx:
		txExecutor = exec.coinbaseTxExec
	case *stypes.SubchainValidatorSetUpdateTx:
		txExecutor = exec.subchainValidatorSetUpdateTxExec
	case *types.SendTx:
		txExecutor = exec.sendTxExec
	case *types.SmartContractTx:
		txExecutor = exec.smartContractTxExec
	default:
		txExecutor = nil
	}
	return txExecutor
}
