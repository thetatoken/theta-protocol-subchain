package execution

import (
	"fmt"
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/result"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/store/database"

	sbc "github.com/thetatoken/thetasubchain/blockchain"
	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/interchain/witness"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
	stypes "github.com/thetatoken/thetasubchain/ledger/types"
)

var _ TxExecutor = (*SubchainValidatorSetUpdateTxExecutor)(nil)

// ------------------------------- SubchainValidatorSet Transaction -----------------------------------

// SubchainValidatorSetUpdateTxExecutor implements the TxExecutor interface
type SubchainValidatorSetUpdateTxExecutor struct {
	db               database.Database
	chain            *sbc.Chain
	state            *slst.LedgerState
	consensus        score.ConsensusEngine
	valMgr           score.ValidatorManager
	metachainWitness witness.ChainWitness
}

// NewSubchainValidatorSetUpdateTxExecutor creates a new instance of SubchainValidatorSetUpdateTxExecutor
func NewSubchainValidatorSetUpdateTxExecutor(db database.Database, chain *sbc.Chain, state *slst.LedgerState, consensus score.ConsensusEngine,
	valMgr score.ValidatorManager, metachainWitness witness.ChainWitness) *SubchainValidatorSetUpdateTxExecutor {
	return &SubchainValidatorSetUpdateTxExecutor{
		db:               db,
		chain:            chain,
		state:            state,
		consensus:        consensus,
		valMgr:           valMgr,
		metachainWitness: metachainWitness,
	}
}

func (exec *SubchainValidatorSetUpdateTxExecutor) sanityCheck(chainID string, view *slst.StoreView, viewSel score.ViewSelector, transaction types.Tx) result.Result {
	tx := transaction.(*stypes.SubchainValidatorSetUpdateTx)
	validatorSet := getValidatorSet(exec.consensus.GetLedger(), exec.valMgr)
	validatorAddresses := getValidatorAddresses(validatorSet)

	// Validate proposer, basic
	res := tx.Proposer.ValidateBasic()
	if res.IsError() {
		return res
	}

	// verify that at most one subchain validator set update transaction is processed for each block
	if view.SubchainValidatorSetTransactionProcessed() {
		return result.Error("Another subchain validator set update transaction has been processed for the current block")
	}

	// verify the proposer is one of the validators
	res = isAValidator(tx.Proposer.Address, validatorAddresses)
	if res.IsError() {
		return res
	}

	proposerAccount, res := getOrMakeInput(view, tx.Proposer)
	if res.IsError() {
		return res
	}

	// verify the proposer's signature
	signBytes := tx.SignBytes(chainID)
	if !tx.Proposer.Signature.Verify(signBytes, proposerAccount.Address) {
		return result.Error("SignBytes: %X", signBytes)
	}

	return result.OK
}

func (exec *SubchainValidatorSetUpdateTxExecutor) process(chainID string, view *slst.StoreView, viewSel score.ViewSelector, transaction types.Tx) (common.Hash, result.Result) {
	tx := transaction.(*stypes.SubchainValidatorSetUpdateTx)

	if view.SubchainValidatorSetTransactionProcessed() {
		return common.Hash{}, result.Error("Another subchain validator set update transaction has been processed for the current block")
	}

	// new validator set and dynasty from the transaction
	newDynasty := tx.Dynasty
	newValidatorSet := score.NewValidatorSet(newDynasty)

	for _, v := range tx.Validators {
		// AddValidator() sorts the validator set, which ensures the order of the validators is determinstic
		newValidatorSet.AddValidator(v)
		validatorAccount := view.GetAccount(v.Address)
		if validatorAccount == nil {
			valAcc := types.NewAccount(v.Address)
			valAcc.LastUpdatedBlockHeight = view.Height()
			view.SetAccount(v.Address, valAcc)
		}
	}

	currentDynasty := view.GetDynasty()
	if newDynasty.Cmp(currentDynasty) <= 0 {
		// This must be an invalid block, since dynasty needs to strictly increase. Without this check,
		// a malicious proposer may attempt to install a validator set of a previous dynasty.
		return common.Hash{}, result.Error(fmt.Sprintf("new dynasty needs to be strictly larger than the current dynasty (new: %v, current: %v)", newDynasty, currentDynasty))
	}

	registrationMainchainHeight, err := exec.metachainWitness.GetSubchainRegistrationHeight()
	if err != nil {
		logger.Warnf("Failed to get subchain registration height: %v", err)
		return common.Hash{}, result.UndecidedWith(result.Info{"newDynasty": newDynasty, "err": err})
	}
	registrationDynasty := scom.CalculateDynasty(registrationMainchainHeight)

	if currentDynasty.Cmp(common.Big0) == 0 && newDynasty.Cmp(registrationDynasty) == 0 {
		// special handling: currentDynasty == 0 and newDynasty == registrationDynasty means the node just loaded the genesis snapshot,
		// and the only purpose of the validatorSetUpdateTx is to update the dynasty in the StoreView. In this case,
		// We should trust the initial ValidatorSet specified by the snapshot
		validatorSetSpecifiedInTheSnapshot := view.GetValidatorSet()

		logger.Debugf("currentDynasty: %v", currentDynasty)
		logger.Debugf("newValidatorSet      : %v", newValidatorSet.String())
		logger.Debugf("validatorSetSpecifiedInTheSnapshot: %v", validatorSetSpecifiedInTheSnapshot.String())

		if !newValidatorSet.EqualsDisregardingDynasty(validatorSetSpecifiedInTheSnapshot) { // the first ever ValidatorSet update essentially only updates the dynasty
			return common.Hash{}, result.Error("validator set mismatch: %v vs %v", *newValidatorSet, *validatorSetSpecifiedInTheSnapshot)
		}
	} else {
		witnessedValidatorSet, err := exec.metachainWitness.GetValidatorSetByDynasty(newDynasty)
		if err != nil {
			return common.Hash{}, result.UndecidedWith(result.Info{"newDynasty": newDynasty, "err": err})
		}

		logger.Debugf("newDynasty: %v", newDynasty)
		logger.Debugf("newValidatorSet      : %v", newValidatorSet.String())
		logger.Debugf("witnessedValidatorSet: %v", witnessedValidatorSet.String())

		if !newValidatorSet.Equals(witnessedValidatorSet) {
			return common.Hash{}, result.Error("validator set mismatch: %v vs %v", *newValidatorSet, *witnessedValidatorSet)
		}
	}

	// update the dynasty and the subchain validator set
	selfChainIDInt := scom.MapChainID(chainID)
	view.UpdateValidatorSet(selfChainIDInt, newValidatorSet)
	view.SetSubchainValidatorSetTransactionProcessed(true)
	txHash := types.TxID(chainID, tx)

	logger.Debugf("Update validator set tx processed, chainID: %v, viewSel: %v, blockHeight: %v", selfChainIDInt, viewSel, view.GetBlockHeight())

	return txHash, result.OK
}

func (exec *SubchainValidatorSetUpdateTxExecutor) getTxInfo(transaction types.Tx) *score.TxInfo {
	return &score.TxInfo{
		EffectiveGasPrice: exec.calculateEffectiveGasPrice(transaction),
	}
}

func (exec *SubchainValidatorSetUpdateTxExecutor) calculateEffectiveGasPrice(transaction types.Tx) *big.Int {
	return new(big.Int).SetUint64(0)
}
