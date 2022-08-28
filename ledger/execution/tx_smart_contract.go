package execution

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/result"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/ledger/types"

	sbc "github.com/thetatoken/thetasubchain/blockchain"
	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
	"github.com/thetatoken/thetasubchain/ledger/vm"
	svm "github.com/thetatoken/thetasubchain/ledger/vm"
)

var _ TxExecutor = (*SmartContractTxExecutor)(nil)

const contractAddrInfoKey string = "contract_address"

// ------------------------------- SmartContractTx Transaction -----------------------------------

// SmartContractTxExecutor implements the TxExecutor interface
type SmartContractTxExecutor struct {
	state  *slst.LedgerState
	chain  *sbc.Chain
	ledger score.Ledger
	valMgr score.ValidatorManager
}

// NewSmartContractTxExecutor creates a new instance of SmartContractTxExecutor
func NewSmartContractTxExecutor(chain *sbc.Chain, state *slst.LedgerState, ledger score.Ledger, valMgr score.ValidatorManager) *SmartContractTxExecutor {
	return &SmartContractTxExecutor{
		state:  state,
		chain:  chain,
		ledger: ledger,
		valMgr: valMgr,
	}
}

func (exec *SmartContractTxExecutor) sanityCheck(chainID string, view *slst.StoreView, viewSel score.ViewSelector, transaction types.Tx) result.Result {
	blockHeight := getBlockHeight(exec.state)
	tx := transaction.(*types.SmartContractTx)

	// Validate from, basic
	res := tx.From.ValidateBasic()
	if res.IsError() {
		return res
	}

	// Check signatures
	signBytes := tx.SignBytes(chainID)
	nativeSignatureValid := tx.From.Signature.Verify(signBytes, tx.From.Address)
	if blockHeight >= common.HeightTxWrapperExtension {
		signBytesV2 := types.ChangeEthereumTxWrapper(signBytes, 2)
		nativeSignatureValid = nativeSignatureValid || tx.From.Signature.Verify(signBytesV2, tx.From.Address)
	}

	if !nativeSignatureValid {
		if blockHeight < common.HeightRPCCompatibility {
			return result.Error("Signature verification failed, SignBytes: %v",
				hex.EncodeToString(signBytes)).WithErrorCode(result.CodeInvalidSignature)
		}

		// interpret the signature as ETH tx signature
		if tx.From.Coins.ThetaWei.Cmp(big.NewInt(0)) != 0 {
			return result.Error("Sending Theta with ETH transaction is not allowed") // extra check, since ETH transaction only signs the TFuel part (i.e., value, gasPrice, gasLimit, etc)
		}

		ethSigningHash := tx.EthSigningHash(chainID, blockHeight)
		err := crypto.ValidateEthSignature(tx.From.Address, ethSigningHash, tx.From.Signature)
		if err != nil {
			return result.Error("ETH Signature verification failed, SignBytes: %v, error: %v",
				hex.EncodeToString(signBytes), err.Error()).WithErrorCode(result.CodeInvalidSignature)
		}
	}

	// Get input account
	fromAccount, success := getInput(view, tx.From)
	if success.IsError() {
		return result.Error("Failed to get the account (the address has no Theta nor TFuel)")
	}

	// Validate input, advanced

	// Check sequence/coins
	seq, balance := fromAccount.Sequence, fromAccount.Balance
	if seq+1 != tx.From.Sequence {
		return result.Error("ValidateInputAdvanced: Got %v, expected %v. (acc.seq=%v)",
			tx.From.Sequence, seq+1, fromAccount.Sequence).WithErrorCode(result.CodeInvalidSequence)
	}

	// Check amount
	if !balance.IsGTE(tx.From.Coins) {
		return result.Error("Insufficient fund: balance is %v, tried to send %v",
			balance, tx.From.Coins).WithErrorCode(result.CodeInsufficientFund)
	}

	coins := tx.From.Coins.NoNil()
	if !coins.IsNonnegative() {
		return result.Error("Invalid value to transfer").
			WithErrorCode(result.CodeInvalidValueToTransfer)
	}

	if !sanityCheckForGasPrice(view, tx.From.Address, tx.To.Address, tx.Data, exec.ledger, exec.valMgr, tx.GasPrice, blockHeight) {
		minimumGasPrice := scom.GetMinimumGasPrice()
		return result.Error("Insufficient gas price. Gas price needs to be at least %v TFuelWei", minimumGasPrice).
			WithErrorCode(result.CodeInvalidGasPrice)
	}

	maxGasLimit := types.GetMaxGasLimit(blockHeight)
	if new(big.Int).SetUint64(tx.GasLimit).Cmp(maxGasLimit) > 0 {
		return result.Error("Invalid gas limit. Gas limit needs to be at most %v", maxGasLimit).
			WithErrorCode(result.CodeInvalidGasLimit)
	}

	err := exec.checkIntrinsicGas(tx)
	if err != nil {
		return result.Error("Intrinsic gas check failed: %v", err).
			WithErrorCode(result.CodeInvalidGasLimit)
	}

	zero := big.NewInt(0)
	feeLimit := new(big.Int).Mul(tx.GasPrice, new(big.Int).SetUint64(tx.GasLimit))
	if feeLimit.BitLen() > 255 || feeLimit.Cmp(zero) < 0 {
		// There is no explicit upper limit for big.Int. Just be conservative
		// here to prevent potential overflow attack
		return result.Error("Fee limit too high").
			WithErrorCode(result.CodeFeeLimitTooHigh)
	}

	if coins.ThetaWei.Cmp(types.Zero) != 0 { // subchains do not support native THETA
		return result.Error("Subchain does not support native THETA").
			WithErrorCode(result.CodeDoNotSupportNativeThetaInSubchain)
	}

	var minimalBalance types.Coins
	value := coins.TFuelWei // NoNil() already guarantees value is NOT nil
	minimalBalance = types.Coins{
		ThetaWei: zero,
		TFuelWei: feeLimit.Add(feeLimit, value),
	}

	if !fromAccount.Balance.IsGTE(minimalBalance) {
		logger.Infof(fmt.Sprintf("Source did not have enough balance %v", tx.From.Address.Hex()))
		return result.Error("Source balance is %v, but required minimal balance is %v",
			fromAccount.Balance, minimalBalance).WithErrorCode(result.CodeInsufficientFund)
	}

	return result.OK
}

func (exec *SmartContractTxExecutor) process(chainID string, view *slst.StoreView, viewSel score.ViewSelector, transaction types.Tx) (common.Hash, result.Result) {
	tx := transaction.(*types.SmartContractTx)

	view.ResetLogs()
	view.ResetBalanceChanges()

	// Note: for contract deployment, vm.Execute() might transfer coins from the fromAccount to the
	//       deployed smart contract. Thus, we should call vm.Execute() before calling getInput().
	//       Otherwise, the fromAccount returned by getInput() will have incorrect balance.
	pb := exec.state.ParentBlock()
	parentBlockInfo := svm.NewBlockInfo(pb.Height, pb.Timestamp, pb.ChainID)
	evmRet, contractAddr, gasUsed, evmErr := svm.Execute(parentBlockInfo, tx, view)

	fromAddress := tx.From.Address
	fromAccount, success := getInput(view, tx.From)
	if success.IsError() {
		return common.Hash{}, result.Error("Failed to get the from account")
	}

	feeAmount := new(big.Int).Mul(tx.GasPrice, new(big.Int).SetUint64(gasUsed))
	fee := types.Coins{
		ThetaWei: big.NewInt(int64(0)),
		TFuelWei: feeAmount,
	}
	if !chargeFee(fromAccount, fee) {
		return common.Hash{}, result.Error("failed to charge transaction fee")
	}

	createContract := (tx.To.Address == common.Address{})
	if !createContract { // svm.create() increments the sequence of the from account
		fromAccount.Sequence++
	}
	view.SetAccount(fromAddress, fromAccount)

	txHash := types.TxID(chainID, tx)

	// TODO: Add tx receipt: status and events
	logs := view.PopLogs()
	balanceChanges := view.PopBalanceChanges()

	if evmErr != nil {
		// Do not record events if transaction is reverted
		logs = nil
		balanceChanges = nil
	}

	if viewSel == score.DeliveredView { // only record the receipt for the delivered views
		exec.chain.AddTxReceipt(exec.ledger.GetCurrentBlock(), tx, logs, balanceChanges, evmRet, contractAddr, gasUsed, evmErr)
	}

	contractInfo := result.Info{}
	contractInfo[contractAddrInfoKey] = contractAddr

	return txHash, result.OKWith(contractInfo)
}

func (exec *SmartContractTxExecutor) checkIntrinsicGas(tx *types.SmartContractTx) error {
	contractAddr := tx.To.Address
	createContract := (contractAddr == common.Address{})
	intrinsicGas, err := vm.CalculateIntrinsicGas(tx.Data, createContract)
	if err != nil {
		return err
	}

	gasLimit := tx.GasLimit
	if intrinsicGas > gasLimit {
		return vm.ErrOutOfGas
	}

	return nil
}

func (exec *SmartContractTxExecutor) getTxInfo(transaction types.Tx) *score.TxInfo {
	tx := transaction.(*types.SmartContractTx)
	return &score.TxInfo{
		Address:           tx.From.Address,
		Sequence:          tx.From.Sequence,
		EffectiveGasPrice: exec.calculateEffectiveGasPrice(transaction),
	}
}

func (exec *SmartContractTxExecutor) calculateEffectiveGasPrice(transaction types.Tx) *big.Int {
	tx := transaction.(*types.SmartContractTx)
	return tx.GasPrice
}
