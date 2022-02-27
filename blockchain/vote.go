package blockchain

import (
	"github.com/thetatoken/theta/common"
	score "github.com/thetatoken/thetasubchain/core"
)

// voteIndexKey constructs the DB key for the given block hash.
func voteIndexKey(hash common.Hash) common.Bytes {
	return append(common.Bytes("vt/"), hash[:]...)
}

// AddVoteToIndex adds a vote to index.
func (ch *Chain) AddVoteToIndex(vote score.Vote) {
	if vote.Block.IsEmpty() {
		return
	}
	key := voteIndexKey(vote.Block)
	voteSet := score.NewVoteSet()
	ch.store.Get(key, voteSet)
	voteSet.AddVote(vote)
	err := ch.store.Put(key, voteSet)
	if err != nil {
		logger.Panic(err)
	}
}

// FindVotesByHash looks up votes by hash.
func (ch *Chain) FindVotesByHash(hash common.Hash) *score.VoteSet {
	voteSet := score.NewVoteSet()
	ch.store.Get(voteIndexKey(hash), voteSet)
	return voteSet
}

// RemoveVotesByHash removes votes for givin block.
func (ch *Chain) RemoveVotesByHash(hash common.Hash) {
	ch.store.Delete(voteIndexKey(hash))
}
