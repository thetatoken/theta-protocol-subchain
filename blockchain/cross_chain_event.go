package blockchain

import (
	"github.com/thetatoken/theta/common"
	score "github.com/thetatoken/thetasubchain/core"
	"math/big"
)

// crossChainEventIndexKey constructs the DB key for the given block hash.
func crossChainEventIndexKey(nonce *big.Int) common.Bytes {
	return common.Bytes("cce/" + nonce.String())
}

// AddCrossChainEventToIndex adds a cross chain event to index.
func (ch *Chain) AddCrossChainEventToIndex(event score.CrossChainTransferEvent) {
	key := crossChainEventIndexKey(event.Nonce)
	err := ch.store.Put(key, event)
	if err != nil {
		logger.Panic(err)
	}
}

// FindCrossChainEventByHash looks up votes by nonce.
func (ch *Chain) FindCrossChainEventByHash(nonce *big.Int) *score.CrossChainTransferEvent {
	event := &score.CrossChainTransferEvent{}
	ch.store.Get(crossChainEventIndexKey(nonce), event)
	return event
}

// RemoveCrossChainEventByHash removes votes for givin nonce.
func (ch *Chain) RemoveCrossChainEventByHash(nonce *big.Int) {
	ch.store.Delete(crossChainEventIndexKey(nonce))
}

