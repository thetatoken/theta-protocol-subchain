package consensus

import (
	"math/big"
	"math/rand"

	log "github.com/sirupsen/logrus"
	"github.com/thetatoken/theta/common"
	score "github.com/thetatoken/thetasubchain/core"
)

const MaxValidatorCount int = 121

//
// -------------------------------- FixedValidatorManager ----------------------------------
//
var _ score.ValidatorManager = &FixedValidatorManager{}

// FixedValidatorManager is an implementation of ValidatorManager interface that selects a fixed validator as the proposer.
type FixedValidatorManager struct {
	consensus score.ConsensusEngine
}

// NewFixedValidatorManager creates an instance of FixedValidatorManager.
func NewFixedValidatorManager() *FixedValidatorManager {
	m := &FixedValidatorManager{
		consensus: nil,
	}
	return m
}

// SetConsensusEngine mplements ValidatorManager interface.
func (m *FixedValidatorManager) SetConsensusEngine(consensus score.ConsensusEngine) {
	m.consensus = consensus
}

// GetProposer implements ValidatorManager interface.
func (m *FixedValidatorManager) GetProposer(blockHash common.Hash, _ uint64) score.Validator {
	return m.getProposerFromValidators(m.GetValidatorSet(blockHash))
}

// GetNextProposer implements ValidatorManager interface.
func (m *FixedValidatorManager) GetNextProposer(blockHash common.Hash, _ uint64) score.Validator {
	return m.getProposerFromValidators(m.GetNextValidatorSet(blockHash))
}

func (m *FixedValidatorManager) getProposerFromValidators(valSet *score.ValidatorSet) score.Validator {
	if valSet.Size() == 0 {
		log.Panic("No validators have been added")
	}

	return valSet.Validators()[0]
}

// GetValidatorSet returns the validator set for given block hash.
func (m *FixedValidatorManager) GetValidatorSet(blockHash common.Hash) *score.ValidatorSet {
	valSet := selectValidatorsForBlock(m.consensus, blockHash, false)
	return valSet
}

// GetNextValidatorSet returns the validator set for given block hash's next block.
func (m *FixedValidatorManager) GetNextValidatorSet(blockHash common.Hash) *score.ValidatorSet {
	valSet := selectValidatorsForBlock(m.consensus, blockHash, true)
	return valSet
}

//
// -------------------------------- RotatingValidatorManager ----------------------------------
//
var _ score.ValidatorManager = &RotatingValidatorManager{}

// RotatingValidatorManager is an implementation of ValidatorManager interface that selects a random validator as
// the proposer using validator's stake as weight.
type RotatingValidatorManager struct {
	consensus score.ConsensusEngine
}

// NewRotatingValidatorManager creates an instance of RotatingValidatorManager.
func NewRotatingValidatorManager() *RotatingValidatorManager {
	m := &RotatingValidatorManager{}
	return m
}

// SetConsensusEngine mplements ValidatorManager interface.
func (m *RotatingValidatorManager) SetConsensusEngine(consensus score.ConsensusEngine) {
	m.consensus = consensus
}

// GetProposer implements ValidatorManager interface.
func (m *RotatingValidatorManager) GetProposer(blockHash common.Hash, epoch uint64) score.Validator {
	return m.getProposerFromValidators(m.GetValidatorSet(blockHash), epoch)
}

// GetNextProposer implements ValidatorManager interface.
func (m *RotatingValidatorManager) GetNextProposer(blockHash common.Hash, epoch uint64) score.Validator {
	return m.getProposerFromValidators(m.GetNextValidatorSet(blockHash), epoch)
}

func (m *RotatingValidatorManager) getProposerFromValidators(valSet *score.ValidatorSet, epoch uint64) score.Validator {
	if valSet.Size() == 0 {
		log.Panic("No validators have been added")
	}

	totalStake := valSet.TotalStake()
	scalingFactor := new(big.Int).Div(totalStake, common.BigMaxUint32)
	scalingFactor = new(big.Int).Add(scalingFactor, common.Big1)
	scaledTotalStake := scaleDown(totalStake, scalingFactor)

	// TODO: replace with more secure randomness.
	rnd := rand.New(rand.NewSource(int64(epoch)))
	r := randUint64(rnd, scaledTotalStake)
	curr := uint64(0)
	validators := valSet.Validators()
	for _, v := range validators {
		curr += scaleDown(v.Stake, scalingFactor)
		if r < curr {
			return v
		}
	}

	// Should not reach here.
	log.Panic("Failed to randomly select a validator")
	panic("Should not reach here")
}

// GetValidatorSet returns the validator set for given block.
func (m *RotatingValidatorManager) GetValidatorSet(blockHash common.Hash) *score.ValidatorSet {
	valSet := selectValidatorsForBlock(m.consensus, blockHash, false)
	return valSet
}

// GetNextValidatorSet returns the validator set for given block's next block.
func (m *RotatingValidatorManager) GetNextValidatorSet(blockHash common.Hash) *score.ValidatorSet {
	valSet := selectValidatorsForBlock(m.consensus, blockHash, true)
	return valSet
}

//
// -------------------------------- Utilities ----------------------------------
//

func FilterValidators(vs *score.ValidatorSet) *score.ValidatorSet {
	valSet := score.NewValidatorSet(vs.Dynasty())
	for _, validatorCandidate := range vs.Validators() {
		valAddr := validatorCandidate.Address.Hex()
		valStake := validatorCandidate.Stake
		if valStake.Cmp(score.Zero) == 0 {
			continue
		}
		validator := score.NewValidator(valAddr, valStake)
		valSet.AddValidator(validator)
	}

	return valSet
}

func selectValidatorsForBlock(consensus score.ConsensusEngine, blockHash common.Hash, isNext bool) *score.ValidatorSet {
	vs, err := consensus.GetLedger().GetFinalizedValidatorSet(blockHash, isNext)
	if err != nil {
		log.Panicf("Failed to get the validator set, blockHash: %v, isNext: %v, err: %v", blockHash.Hex(), isNext, err)
	}
	if vs == nil {
		log.Panic("Failed to retrieve the validator set")
	}

	return FilterValidators(vs)
}

// Generate a random uint64 in [0, max)
func randUint64(rnd *rand.Rand, max uint64) uint64 {
	const maxInt64 uint64 = 1<<63 - 1
	if max <= maxInt64 {
		return uint64(rnd.Int63n(int64(max)))
	}
	for {
		r := rnd.Uint64()
		if r < max {
			return r
		}
	}
}

func scaleDown(x *big.Int, scalingFactor *big.Int) uint64 {
	if scalingFactor.Cmp(common.Big0) == 0 {
		log.Panic("scalingFactor is zero")
	}
	scaledX := new(big.Int).Div(x, scalingFactor)
	scaledXUint64 := scaledX.Uint64()
	return scaledXUint64
}
