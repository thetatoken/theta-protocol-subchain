package execution

import (
	"encoding/hex"
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/result"
	"github.com/thetatoken/theta/core"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/ledger/types"

	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
)

// --------------------------------- Execution Utilities -------------------------------------

// TODO: need to implement the following two functions
// // Read genesis file.
// func ReadGenesisFile() (genDoc *ttypes.GenesisDoc, err error) {
// 	cfg, err := tcmd.ParseConfig()
// 	if err != nil {
// 		return
// 	}

// 	return ReadGenesisFileByPath(cfg.GenesisFile())
// }

// func ReadGenesisFileByPath(path string) (genDoc *ttypes.GenesisDoc, err error) {
// 	genDocJSON, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return
// 	}

// 	genDoc, err = ttypes.GenesisDocFromJSON(genDocJSON)
// 	return
// }

// getValidatorAddresses returns the validator set
func getValidatorSet(ledger score.Ledger, valMgr score.ValidatorManager) *score.ValidatorSet {
	currentBlock := ledger.GetCurrentBlock()
	if currentBlock == nil {
		panic("ledger.currentBlock is nil")
	}
	parentBlkHash := currentBlock.Parent
	validatorSet := valMgr.GetNextValidatorSet(parentBlkHash)
	return validatorSet
}

// getValidatorAddresses returns validators' addresses
func getValidatorAddresses(validatorSet *score.ValidatorSet) []common.Address {
	validators := validatorSet.Validators()
	validatorAddresses := make([]common.Address, len(validators))
	for i, v := range validators {
		validatorAddresses[i] = v.Address
	}
	return validatorAddresses
}

func isAValidator(address common.Address, validatorAddresses []common.Address) result.Result {
	proposerIsAValidator := false
	for _, validatorAddr := range validatorAddresses {
		if address == validatorAddr {
			proposerIsAValidator = true
			break
		}
	}
	if !proposerIsAValidator {
		return result.Error("The coinbaseTx proposer is not a validator")
	}

	return result.OK
}

// The accounts from the TxInputs must either already have
// crypto.PubKey.(type) != nil, (it must be known),
// or it must be specified in the TxInput.
func getInputs(view *slst.StoreView, ins []types.TxInput) (map[string]*types.Account, result.Result) {
	accounts := map[string]*types.Account{}
	for _, in := range ins {
		// Account shouldn't be duplicated
		if _, ok := accounts[string(in.Address[:])]; ok {
			return nil, result.Error("getInputs - Duplicated address: %v", in.Address)
		}

		acc, success := getAccount(view, in.Address)
		if success.IsError() {
			return nil, result.Error("getInputs - Unknown address: %v", in.Address)
		}

		accounts[string(in.Address[:])] = acc
	}
	return accounts, result.OK
}

func getInput(view *slst.StoreView, in types.TxInput) (*types.Account, result.Result) {
	return getOrMakeInputImpl(view, in, false)
}

func getOrMakeInput(view *slst.StoreView, in types.TxInput) (*types.Account, result.Result) {
	return getOrMakeInputImpl(view, in, true)
}

// This function guarantees the public key of the retrieved account is not empty
func getOrMakeInputImpl(view *slst.StoreView, in types.TxInput, makeNewAccount bool) (*types.Account, result.Result) {
	acc, success := getOrMakeAccountImpl(view, in.Address, makeNewAccount)
	if success.IsError() {
		return nil, result.Error("getOrMakeInputImpl - Unknown address: %v", in.Address)
	}

	return acc, result.OK
}

func getAccount(view *slst.StoreView, address common.Address) (*types.Account, result.Result) {
	return getOrMakeAccountImpl(view, address, false)
}

func getOrMakeAccount(view *slst.StoreView, address common.Address) *types.Account {
	acc, _ := getOrMakeAccountImpl(view, address, true)
	// Impossible to have error.
	return acc
}

func getOrMakeAccountImpl(view *slst.StoreView, address common.Address, makeNewAccount bool) (*types.Account, result.Result) {
	acc := view.GetAccount(address)
	if acc == nil {
		if !makeNewAccount {
			return nil, result.Error("getOrMakeAccountImpl - Unknown address: %v", address)
		}
		acc = types.NewAccount(address)
		acc.LastUpdatedBlockHeight = view.Height()
	}
	acc.UpdateToHeight(view.Height())

	return acc, result.OK
}

func getOrMakeOutputs(view *slst.StoreView, accounts map[string]*types.Account, outs []types.TxOutput) (map[string]*types.Account, result.Result) {
	if accounts == nil {
		accounts = make(map[string]*types.Account)
	}

	for _, out := range outs {
		// Account shouldn't be duplicated
		if _, ok := accounts[string(out.Address[:])]; ok {
			return nil, result.Error("getOrMakeOutputs - Duplicated address: %v", out.Address)
		}

		acc := getOrMakeAccount(view, out.Address)
		accounts[string(out.Address[:])] = acc
	}
	return accounts, result.OK
}

// Validate inputs basic structure
func validateInputsBasic(ins []types.TxInput) result.Result {
	for _, in := range ins {
		// Check TxInput basic
		if res := in.ValidateBasic(); res.IsError() {
			return res
		}

		// subchain does not support native Thta coins
		if in.Coins.NoNil().ThetaWei.Cmp(core.Zero) != 0 {
			return result.Error("Subchain should not transfer native Theta coin: %v", in)
		}
	}
	return result.OK
}

// Validate inputs and compute total amount of coins
func validateInputsAdvanced(accounts map[string]*types.Account, signBytes []byte, ins []types.TxInput, blockHeight uint64) (total types.Coins, res result.Result) {
	total = types.NewCoins(0, 0)
	for _, in := range ins {
		acc := accounts[string(in.Address[:])]
		if acc == nil {
			panic("validateInputsAdvanced() expects account in accounts")
		}
		res = validateInputAdvanced(acc, signBytes, in, blockHeight)
		if res.IsError() {
			return
		}
		// Good. Add amount to total
		total = total.Plus(in.Coins)
	}
	return total, result.OK
}

func validateInputAdvanced(acc *types.Account, signBytes []byte, in types.TxInput, blockHeight uint64) result.Result {
	// Check sequence/coins
	seq, balance := acc.Sequence, acc.Balance
	if seq+1 != in.Sequence {
		return result.Error("ValidateInputAdvanced: Got %v, expected %v. (acc.seq=%v)",
			in.Sequence, seq+1, acc.Sequence).WithErrorCode(result.CodeInvalidSequence)
	}

	// Check amount
	if !balance.IsGTE(in.Coins) {
		return result.Error("Insufficient fund: balance is %v, tried to send %v",
			balance, in.Coins).WithErrorCode(result.CodeInsufficientFund)
	}

	// Check signatures
	signatureValid := in.Signature.Verify(signBytes, acc.Address)
	if blockHeight >= common.HeightTxWrapperExtension {
		signBytesV2 := types.ChangeEthereumTxWrapper(signBytes, 2)
		signatureValid = signatureValid || in.Signature.Verify(signBytesV2, acc.Address)
	}

	if !signatureValid {
		return result.Error("Signature verification failed, SignBytes: %v",
			hex.EncodeToString(signBytes)).WithErrorCode(result.CodeInvalidSignature)
	}

	return result.OK
}

func validateOutputsBasic(outs []types.TxOutput) result.Result {
	for _, out := range outs {
		// Check TxOutput basic
		if res := out.ValidateBasic(); res.IsError() {
			return res
		}
	}
	return result.OK
}

func sumOutputs(outs []types.TxOutput) types.Coins {
	total := types.NewCoins(0, 0)
	for _, out := range outs {
		total = total.Plus(out.Coins)
	}
	return total
}

// Note: Since totalInput == totalOutput + fee, the transaction fee is charged implicitly
//       by the following adjustByInputs() function. No special handling needed
func adjustByInputs(view *slst.StoreView, accounts map[string]*types.Account, ins []types.TxInput) {
	for _, in := range ins {
		acc := accounts[string(in.Address[:])]
		if acc == nil {
			panic("adjustByInputs() expects account in accounts")
		}
		if !acc.Balance.IsGTE(in.Coins) {
			panic("adjustByInputs() expects sufficient funds")
		}
		acc.Balance = acc.Balance.Minus(in.Coins)
		acc.Sequence++
		view.SetAccount(in.Address, acc)
	}
}

func adjustByOutputs(view *slst.StoreView, accounts map[string]*types.Account, outs []types.TxOutput) {
	for _, out := range outs {
		acc := accounts[string(out.Address[:])]
		if acc == nil {
			panic("adjustByOutputs() expects account in accounts")
		}
		acc.Balance = acc.Balance.Plus(out.Coins)
		view.SetAccount(out.Address, acc)
	}
}

func sanityCheckForGasPrice(view *slst.StoreView, fromAddr common.Address, contractAddr common.Address, calldata common.Bytes,
	ledger score.Ledger, valMgr score.ValidatorManager, gasPrice *big.Int, blockHeight uint64) bool {
	if gasPrice == nil {
		return false
	}

	if isWhitelistedOperation(view, fromAddr, contractAddr, calldata, ledger, valMgr) && gasPrice.Cmp(common.Big0) == 0 {
		return true
	}

	minimumGasPrice := scom.GetMinimumGasPrice()
	if gasPrice.Cmp(minimumGasPrice) < 0 {
		return false
	}

	return true
}

// Allow subchain validators to vote for inter-chain events without gas fees
func isWhitelistedOperation(view *slst.StoreView, fromAddr common.Address, contractAddr common.Address, calldata common.Bytes,
	ledger score.Ledger, valMgr score.ValidatorManager) bool {
	logger.Debugf("Checking whitelisted operation...")

	// check if the from address is a validator
	validatorSet := view.GetValidatorSet()
	validatorAddresses := getValidatorAddresses(validatorSet)
	res := isAValidator(fromAddr, validatorAddresses)

	logger.Debugf("Checking whitelisted operation, Validators: %v, fromAddr: %v, isValidator: %v", validatorAddresses, fromAddr, !res.IsError())
	if res.IsError() {
		return false
	}

	// check if this tx is for inter-chain event voting
	if !isVotingForInterchainEvents(contractAddr, calldata, ledger) {
		logger.Debugf("Checking whitelisted operation, isWhitelistedOperation: %v", false)
		return false
	}

	logger.Debugf("Checking whitelisted operation, isWhitelistedOperation: %v", true)
	return true
}

func isVotingForInterchainEvents(contractAddr common.Address, calldata common.Bytes, ledger score.Ledger) bool {
	if len(calldata) < 4 {
		logger.Debugf("Checking whitelisted operation, calldata.len: %v", len(calldata))
		return false
	}

	logger.Debugf("Checking whitelisted operation, contract address: %v", contractAddr.Hex())
	whitelistedMethodSignatures := []string{}
	if contractAddr == *ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTFuel) {
		whitelistedMethodSignatures = append(whitelistedMethodSignatures, "mintVouchers(string,address,uint256,uint256,uint256)") // TFuelTokenBank.mintVouchers
		// Note: TFuelTokenBank.unlockTokens is NOT whitelisted since it can only be called on the main chain
	} else if contractAddr == *ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT20) {
		whitelistedMethodSignatures = append(whitelistedMethodSignatures, "mintVouchers(string,string,string,uint8,address,uint256,uint256,uint256)") // TNT20TokenBank.mintVouchers
		whitelistedMethodSignatures = append(whitelistedMethodSignatures, "unlockTokens(uint256,string,address,uint256,uint256,uint256)")             // TNT20TokenBank.unlockTokens
	} else if contractAddr == *ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT721) {
		whitelistedMethodSignatures = append(whitelistedMethodSignatures, "mintVouchers(string,string,string,address,uint256,string,uint256,uint256)") // TNT721TokenBank.mintVouchers
		whitelistedMethodSignatures = append(whitelistedMethodSignatures, "unlockTokens(uint256,string,address,uint256,uint256,uint256)")              // TNT721TokenBank.unlockTokens
	} else if contractAddr == *ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT1155) {
		whitelistedMethodSignatures = append(whitelistedMethodSignatures, "mintVouchers(string,address,uint256,uint256,string,uint256,uint256)")  // TNT1155TokenBank.mintVouchers
		whitelistedMethodSignatures = append(whitelistedMethodSignatures, "unlockTokens(uint256,string,address,uint256,uint256,uint256,uint256)") // TN1155TokenBank.unlockTokens
	} else {
		logger.Debugf("Checking whitelisted operation, contract is not a TokenBank")
		return false
	}

	for _, wms := range whitelistedMethodSignatures {
		funcSelector := hex.EncodeToString(crypto.Keccak256(common.Bytes(wms))[:4])
		calldataLeadingBytes := hex.EncodeToString(calldata[:4])
		logger.Debugf("Checking whitelisted operation, calldataLeadingBytes: 0x%v, funcSelector : 0x%v", calldataLeadingBytes, funcSelector)

		if calldataLeadingBytes == funcSelector {
			logger.Debugf("Checking whitelisted operation, the method called IS whitelisted")
			return true
		}
	}
	logger.Debugf("Checking whitelisted operation, the method called is NOT whitelisted")

	return false
}

func sanityCheckForFee(fee types.Coins, blockHeight uint64) (minimumFee *big.Int, success bool) {
	fee = fee.NoNil()
	minimumFee = types.GetMinimumTransactionFeeTFuelWei(blockHeight)
	success = (fee.ThetaWei.Cmp(types.Zero) == 0 && fee.TFuelWei.Cmp(minimumFee) >= 0)

	return minimumFee, success
}

func sanityCheckForSendTxFee(fee types.Coins, numAccountsAffected uint64, blockHeight uint64) (minimumFee *big.Int, success bool) {
	fee = fee.NoNil()
	minimumFee = types.GetSendTxMinimumTransactionFeeTFuelWei(numAccountsAffected, blockHeight)
	success = (fee.ThetaWei.Cmp(types.Zero) == 0 && fee.TFuelWei.Cmp(minimumFee) >= 0)

	return minimumFee, success
}

func chargeFee(account *types.Account, fee types.Coins) bool {
	if !account.Balance.IsGTE(fee) {
		return false
	}

	account.Balance = account.Balance.Minus(fee)
	return true
}

func getBlockHeight(ledgerState *slst.LedgerState) uint64 {
	blockHeight := ledgerState.Height() + 1
	return blockHeight
}

func getRegularTxGas(ledgerState *slst.LedgerState) uint64 {
	blockHeight := getBlockHeight(ledgerState)
	if blockHeight < common.HeightJune2021FeeAdjustment {
		return types.GasRegularTx
	}
	return types.GasRegularTxJune2021
}
