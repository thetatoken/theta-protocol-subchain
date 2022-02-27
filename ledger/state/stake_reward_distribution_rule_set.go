package state

import (
	"log"

	"github.com/thetatoken/theta/common"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/theta/ledger/types"
)

type StakeRewardDistributionRuleSet struct {
	sv *StoreView
}

// NewStakeRewardDistributionRuleSet creates a new instance of StakeRewardDistributionRuleSet.
func NewStakeRewardDistributionRuleSet(sv *StoreView) *StakeRewardDistributionRuleSet {
	return &StakeRewardDistributionRuleSet{
		sv: sv,
	}
}

// Get returns the reward distribution of a stake holder address in the pool. Returns nil if not found.
func (srdr *StakeRewardDistributionRuleSet) Get(stakeHolder common.Address) *score.RewardDistribution {
	rewardDistrKey := StakeRewardDistributionRuleSetKey(stakeHolder)
	data := srdr.sv.Get(rewardDistrKey)
	if data == nil || len(data) == 0 {
		return nil
	}

	rewardDistr := &score.RewardDistribution{}
	err := types.FromBytes(data, rewardDistr)
	if err != nil {
		log.Panicf("StakeRewardDistributionRuleSet.Get: Error reading reward distribution rule %X, error: %v",
			data, err.Error())
	}

	return rewardDistr
}

// Upsert update or inserts a stake reward distribution to the rule set
func (srdr *StakeRewardDistributionRuleSet) Upsert(rd *score.RewardDistribution) {
	rewardDistrKey := StakeRewardDistributionRuleSetKey(rd.StakeHolder)
	data, err := types.ToBytes(rd)
	if err != nil {
		log.Panicf("StakeRewardDistributionRuleSet.Upsert: Error serializing the reward distribution rule %X, error: %v",
			data, err.Error())
	}
	srdr.sv.Set(rewardDistrKey, data)
}

// Remove removes an elite edge node from the pool; returns false if guardian is not found.
func (srdr *StakeRewardDistributionRuleSet) Remove(stakeHolder common.Address) {
	rewardDistrKey := StakeRewardDistributionRuleSetKey(stakeHolder)
	srdr.sv.Delete(rewardDistrKey)
}

// GetAll returns all the reward distribution rules
func (srdr *StakeRewardDistributionRuleSet) GetAll() []*score.RewardDistribution {
	prefix := StakeRewardDistributionRuleSetKeyPrefix()
	rewardDistrList := []*score.RewardDistribution{}
	cb := func(k, v common.Bytes) bool {
		rewardDistr := &score.RewardDistribution{}
		err := types.FromBytes(v, rewardDistr)
		if err != nil {
			log.Panicf("StakeRewardDistributionRuleSet.GetAll: Error reading reward distribution rule %X, error: %v",
				v, err.Error())
		}
		rewardDistrList = append(rewardDistrList, rewardDistr)
		return true
	}

	srdr.sv.Traverse(prefix, cb)

	return rewardDistrList
}
