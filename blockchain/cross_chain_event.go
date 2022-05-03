package blockchain

import (
	"math/big"

	score "github.com/thetatoken/thetasubchain/core"
)

// AddCrossChainEventToIndex adds a cross chain event to index.
func (ch *Chain) AddCrossChainEventToIndex(event score.CrossChainTransferEvent) {
	key := score.CrossChainEventIndexKey(event.Nonce)
	err := ch.store.Put(key, event)
	if err != nil {
		logger.Panic(err)
	}
}

// FindCrossChainEventByHash looks up votes by nonce.
func (ch *Chain) FindCrossChainEventByHash(nonce *big.Int) *score.CrossChainTransferEvent {
	event := &score.CrossChainTransferEvent{}
	ch.store.Get(score.CrossChainEventIndexKey(nonce), event)
	return event
}

// RemoveCrossChainEventByHash removes votes for givin nonce.
func (ch *Chain) RemoveCrossChainEventByHash(nonce *big.Int) {
	ch.store.Delete(score.CrossChainEventIndexKey(nonce))
}
