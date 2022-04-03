package consensus

import (
	"math/big"
	"math/rand"

	log "github.com/sirupsen/logrus"
	"github.com/thetatoken/theta/common"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/witness"
)

const MaxValidatorCount int = 31

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
	valSet := selectTopStakeHoldersAsValidatorsForBlock(m.consensus, blockHash, false)
	return valSet
}

// GetNextValidatorSet returns the validator set for given block hash's next block.
func (m *FixedValidatorManager) GetNextValidatorSet(blockHash common.Hash) *score.ValidatorSet {
	valSet := selectTopStakeHoldersAsValidatorsForBlock(m.consensus, blockHash, true)
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
	valSet := selectTopStakeHoldersAsValidatorsForBlock(m.consensus, blockHash, false)
	return valSet
}

// GetNextValidatorSet returns the validator set for given block's next block.
func (m *RotatingValidatorManager) GetNextValidatorSet(blockHash common.Hash) *score.ValidatorSet {
	valSet := selectTopStakeHoldersAsValidatorsForBlock(m.consensus, blockHash, true)
	return valSet
}

//
// -------------------------------- SubchainRotatingValidatorManager ----------------------------------
//
var _ score.ValidatorManager = &SubchainRotatingValidatorManager{}

// SubchainRotatingValidatorManager is an implementation of ValidatorManager interface that selects a random validator as
// the proposer using validator's stake as weight.
type SubchainRotatingValidatorManager struct {
	consensus        score.ConsensusEngine
	mainchainWitness *witness.MainchainWitness
}

// NewSubchainRotatingValidatorManager creates an instance of SubchainRotatingValidatorManager.
func NewSubchainRotatingValidatorManager() *SubchainRotatingValidatorManager {
	srvm := &SubchainRotatingValidatorManager{}
	return srvm
}

// SetConsensusEngine implements ValidatorManager interface.
func (srvm *SubchainRotatingValidatorManager) SetConsensusEngine(consensus score.ConsensusEngine) {
	srvm.consensus = consensus
}

// SetMainchainWitness set the mainchain monitor for the validator manager
func (srvm *SubchainRotatingValidatorManager) SetMainchainWitness(mainchainWitness *witness.MainchainWitness) {
	srvm.mainchainWitness = mainchainWitness
}

// GetProposer implements ValidatorManager interface.
func (srvm *SubchainRotatingValidatorManager) GetProposer(blockHash common.Hash, epoch uint64) score.Validator {
	return score.Validator{} // TODO: implement
}

// GetNextProposer implements ValidatorManager interface.
func (srvm *SubchainRotatingValidatorManager) GetNextProposer(blockHash common.Hash, epoch uint64) score.Validator {
	return score.Validator{} // TODO: implement
}

func (srvm *SubchainRotatingValidatorManager) getProposerFromValidators(valSet *score.ValidatorSet, epoch uint64) score.Validator {
	return score.Validator{} // TODO: implement
}

// GetValidatorSet returns the validator set for given block.
func (srvm *SubchainRotatingValidatorManager) GetValidatorSet(blockHash common.Hash) *score.ValidatorSet {
	return nil // TODO: implement
}

// GetNextValidatorSet returns the validator set for given block's next block.
func (srvm *SubchainRotatingValidatorManager) GetNextValidatorSet(blockHash common.Hash) *score.ValidatorSet {
	return nil // TODO: implement
}

//
// -------------------------------- Utilities ----------------------------------
//

func SelectTopStakeHoldersAsValidators(vcp *score.ValidatorCandidatePool) *score.ValidatorSet {
	maxNumValidators := MaxValidatorCount
	topStakeHolders := vcp.GetTopStakeHolders(maxNumValidators)

	valSet := score.NewValidatorSet()
	for _, stakeHolder := range topStakeHolders {
		valAddr := stakeHolder.Holder.Hex()
		valStake := stakeHolder.TotalStake()
		if valStake.Cmp(score.Zero) == 0 {
			continue
		}
		validator := score.NewValidator(valAddr, valStake)
		valSet.AddValidator(validator)
	}

	return valSet
}

func selectTopStakeHoldersAsValidatorsForBlock(consensus score.ConsensusEngine, blockHash common.Hash, isNext bool) *score.ValidatorSet {
	vcp, err := consensus.GetLedger().GetFinalizedValidatorCandidatePool(blockHash, isNext)
	if err != nil {
		log.Panicf("Failed to get the validator candidate pool, blockHash: %v, isNext: %v, err: %v", blockHash.Hex(), isNext, err)
	}
	if vcp == nil {
		log.Panic("Failed to retrieve the validator candidate pool")
	}

	return SelectTopStakeHoldersAsValidators(vcp)
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
