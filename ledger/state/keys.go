package state

import (
	"math/big"

	"github.com/thetatoken/theta/common"
)

//
// ------------------------- Ledger State Keys -------------------------
//

// ChainIDKey returns the key for chainID
func ChainIDKey() common.Bytes {
	return common.Bytes("chainid")
}

// AccountKey constructs the state key for the given address
func AccountKey(addr common.Address) common.Bytes {
	return append(common.Bytes("ls/a/"), addr[:]...)
}

// SplitRuleKeyPrefix returns the prefix for the split rule key
func SplitRuleKeyPrefix() common.Bytes {
	return common.Bytes("ls/ssc/split/") // special smart contract / split rule
}

// SplitRuleKey constructs the state key for the given resourceID
func SplitRuleKey(resourceID string) common.Bytes {
	resourceIDBytes := common.Bytes(resourceID)
	return append(SplitRuleKeyPrefix(), resourceIDBytes[:]...)
}

// CodeKey constructs the state key for the given code hash
func CodeKey(codeHash common.Bytes) common.Bytes {
	return append(common.Bytes("ls/ch/"), codeHash...)
}

// CurrentValidatorSetKey returns the state key for the current validator stake holder set
func CurrentValidatorSetKey() common.Bytes {
	return common.Bytes("ls/vs")
}

// ValidatorSetForChainDuringDynastyKey returns the key for the validator set for a chain during the given dynasty
func ValidatorSetForChainDuringDynastyKey(chainID *big.Int, dynasty *big.Int) common.Bytes {
	key := common.Bytes("ls/vscd/")
	key = append(key, common.Bytes(chainID.String())...)
	key = append(key, common.Bytes("/")...)
	key = append(key, common.Bytes(dynasty.String())...)
	return key
}

// // EventNonceKey returns the state key for the last processed event nonce
// func EventNonceKey(eventType score.InterChainMessageEventType) common.Bytes {
// 	return common.Bytes("ls/evn/" + strconv.FormatUint(uint64(eventType), 10))
// }

// ValidatorSetUpdateTxHeightListKey returns the state key the heights of blocks
// that contain stake related transactions (i.e. StakeDeposit, StakeWithdraw, etc)
func ValidatorSetUpdateTxHeightListKey() common.Bytes {
	return common.Bytes("ls/sthl")
}

// StatePruningProgressKey returns the key for the state pruning progress
func StatePruningProgressKey() common.Bytes {
	return common.Bytes("ls/spp")
}

// TFuelTokenBankContractAddressKey returns the key for looking up the address of the
// TFuel token bank contract deployed in the genesis block
func TFuelTokenBankContractAddressKey() common.Bytes {
	return common.Bytes("ls/tbca/tfuel")
}

// TNT20TokenBankContractAddressKey returns the key for looking up the address of the
// TNT20 token bank contract deployed in the genesis block
func TNT20TokenBankContractAddressKey() common.Bytes {
	return common.Bytes("ls/tbca/tnt20")
}

// TNT721TokenBankContractAddressKey returns the key for looking up the address of the
// TNT721 token bank contract deployed in the genesis block
func TNT721TokenBankContractAddressKey() common.Bytes {
	return common.Bytes("ls/tbca/tnt721")
}
