package state

import (
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

// ValidatorSetKey returns the state key for the validator stake holder set
func ValidatorSetKey() common.Bytes {
	return common.Bytes("ls/vs")
}

// EventNonceKey returns the state key for the last processed event nonce
func EventNonceKey() common.Bytes {
	return common.Bytes("ls/evn")
}

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
