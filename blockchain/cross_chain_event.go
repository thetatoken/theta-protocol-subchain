package blockchain

import (
	"github.com/thetatoken/theta/common"
	score "github.com/thetatoken/thetasubchain/core"
)

// crossChainEventIndexKey constructs the DB key for the given block hash.
func CrossChainEventIndexKey(nonce *big.Int) common.Bytes {
	return append(common.Bytes("cce/"), nonce.Bytes())
}

// AddCrossChainEventToIndex adds a cross chain event to index.
func (ch *Chain) AddCrossChainEventToIndex(event score.CrossChainTransferEvent) {
	key := CrossChainEventIndexKey(event.Nonce)
	err := ch.store.Put(key, event)
	if err != nil {
		logger.Panic(err)
	}
}

// FindCrossChainEventByHash looks up votes by nonce.
func (ch *Chain) FindCrossChainEventByHash(nonce *big.Int) *score.VoteSet {
	event := &score.CrossChainTransferEvent{}
	ch.store.Get(CrossChainEventIndexKey(nonce), event)
	return event
}

// RemoveCrossChainEventByHash removes votes for givin nonce.
func (ch *Chain) RemoveCrossChainEventByHash(nonce *big.Int) {
	ch.store.Delete(CrossChainEventIndexKey(nonce))
}
