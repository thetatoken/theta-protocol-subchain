package state

import (
	"strconv"

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

// ValidatorCandidatePoolKey returns the state key for the validator stake holder set
func ValidatorCandidatePoolKey() common.Bytes {
	return common.Bytes("ls/vcp")
}

// StakeTransactionHeightListKey returns the state key the heights of blocks
// that contain stake related transactions (i.e. StakeDeposit, StakeWithdraw, etc)
func StakeTransactionHeightListKey() common.Bytes {
	return common.Bytes("ls/sthl")
}

// StatePruningProgressKey returns the key for the state pruning progress
func StatePruningProgressKey() common.Bytes {
	return common.Bytes("ls/spp")
}

// StakeRewardDistributionRuleSetKeyPrefix returns the prefix of the stake reward distribution rule
func StakeRewardDistributionRuleSetKeyPrefix() common.Bytes {
	return common.Bytes("ls/srdrs/")
}

// StakeRewardDistributionRuleSetKey returns the prefix of the stake reward distribution rule
func StakeRewardDistributionRuleSetKey(addr common.Address) common.Bytes {
	prefix := StakeRewardDistributionRuleSetKeyPrefix()
	return append(prefix, addr[:]...)
}

//EliteEdgeNodeStakeReturnsKeyPrefix returns the prefix of the elite edge node stake return key
func EliteEdgeNodeStakeReturnsKeyPrefix() common.Bytes {
	return common.Bytes("ls/eensrk/")
}

//EliteEdgeNodeStakeReturnsKey returns the EEN stake return key for the given height
func EliteEdgeNodeStakeReturnsKey(height uint64) common.Bytes {
	heightStr := strconv.FormatUint(height, 10)
	return common.Bytes(string(EliteEdgeNodeStakeReturnsKeyPrefix()) + heightStr)
}

func EliteEdgeNodesTotalActiveStakeKey() common.Bytes {
	return common.Bytes("ls/eentas")
}
