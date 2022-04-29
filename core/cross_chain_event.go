package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"sort"

	log "github.com/sirupsen/logrus"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/rlp"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "core"})

// CrossChainTransferEvent contains the public information of a crosschain transfer event.
type CrossChainTransferEvent struct {
	Sender      common.Address
	Receiver    common.Address
	Denom       string
	Amount      *big.Int
	Nonce       *big.Int
	BlockNumber *big.Int
}

// NewCrossChainTransferEvent creates a new crosschain transfer event instance.
func NewCrossChainTransferEvent(senderAddr string, receiverAddr string, denom string, amount *big.Int, eventNonce *big.Int, blockNumber *big.Int) CrossChainTransferEvent {
	sender := common.HexToAddress(senderAddr)
	receiver := common.HexToAddress(receiverAddr)
	return CrossChainTransferEvent{sender, receiver, denom, amount, eventNonce, blockNumber}
}

// ID returns the ID of the crosschain transfer event, which is the string representation of its address.
func (c CrossChainTransferEvent) ID() *big.Int {
	return c.EventNonce
}

// Equals checks whether the crosschain transfer event is the same as another crosschain transfer event
func (c CrossChainTransferEvent) Equals(x CrossChainTransferEvent) bool {
	if v.EventNonce.Cmp(x.EventNonce) != 0 {
		return false
	}
	if v.Sender.Hex() != x.Sender.Hex() {
		return false
	}
	if v.Receiver.Hex() != x.Receiver.Hex() {
		return false
	}
	if v.BlockNumber.Cmp(x.BlockNumber) != 0 {
		return false
	}
	if v.Amount.Cmp(x.Amount) != 0 {
		return false
	}
	return true
}

// String represents the string representation of the validator
func (c CrossChainTransferEvent) String() string {
	return fmt.Sprintf("{ID: %v, Denom: %v, from: %v, to: %v}", c.ID(), c.Denom, c.Sender.Hex(), c.Receiver.Hex())
}

// ByID implements sort.Interface for CrossChainTransferEvent based on ID (Nonce).
type ByID []CrossChainTransferEvent

func (b ByID) Len() int           { return len(b) }
func (b ByID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByID) Less(i, j int) bool { return b[i].EventNonce.Cmp(b[j].EventNonce) < 0 }

var _ rlp.Encoder = (*ValidatorSet)(nil)

// EncodeRLP implements RLP Encoder interface.
func (c *CrossChainTransferEvent) EncodeRLP(w io.Writer) error {
	if c == nil {
		return rlp.Encode(w, &CrossChainTransferEvent{})
	}
	return rlp.Encode(w, CrossChainTransferEvent{
		Sender:      c.Sender,
		Receiver:    c.Receiver,
		Denom:       c.Denom,
		Amount:      c.Amount,
		Nonce:       c.Nonce,
		BlockNumber: c.BlockNumber,
	})
}

// DecodeRLP implements RLP Decoder interface.
func (c *CrossChainTransferEvent) DecodeRLP(stream *rlp.Stream) error {
	_, err := stream.List()
	if err != nil {
		return err
	}

	var event CrossChainTransferEvent
	err = stream.Decode(&event)
	if err != nil {
		return err
	}
	c.Sender, c.Receiver, c.Denom, c.Amount, c.Nonce, c.BlockNumber = event.Sender, event.Receiver, event.Denom, event.Amount, event.Nonce, event.BlockNumber

	return stream.ListEnd()
}

// ------------------------------------Event Cache----------------------------------------------

var (
	// ErrCrossChainTransferEventrNotFound for ID is not found in crosschain transfer event set.
	ErrCrossChainTransferEventNotFound      = errors.New("CrossChainTransferEventNotFound")
	ErrCrossChainTransferEventExisted       = errors.New("CrossChainTransferEventrExisted")
	ErrCrossChainTransferEventPersistFailed = errors.New("CrossChainTransferEventPersistFailed")
)

type CrossChainEventCache struct {
	eventCache map[*big.Int]*CrossChainTransferEvent
	db         database.Database
}

// NewCrossChainEventCache creates a new crosschain transfer event cache instance.
func NewCrossChainEventCache(db database.Database) *CrossChainEventCache {
	res := &CrossChainEventCache{
		eventCache: map[*big.Int]*CrossChainTransferEvent{},
		db: db,
	}
	return res
}

func (c *CrossChainEventCache) Size() int {
	return len(c.EventCache)
}

func (c *CrossChainEventCache) IsEmpty() int {
	return len(c.EventCache) == 0
}

func (c *CrossChainEventCache) Insert(event *CrossChainTransferEvent) error {
	if _, ok := c.EventCache[event.Nonce]; ok {
		return ErrCrossChainTransferEventrExisted
	}
	c.EventCache[event.Nonce] = event
	persistent(event)

}

func (c *CrossChainEventCache) persistent(event *CrossChainTransferEvent) error {
	store := kvstore.NewKVStore(db)
	err := store.Put(blockchain.CrossChainEventIndexKey(event.Nonce), event)
	if err != nil {
		c.Delete(event.Nonce)
		logger.Panic(err)
	}
}

func (c *CrossChainEventCache) Delete(nonce *big.Int) {
	delete(c.EventCache, nonce)
}

func (c *CrossChainEventCache) Get(nonce *big.Int) (event *CrossChainTransferEvent, error) {
	if res, ok := c.EventCache[event.Nonce]; ok {
		return res, nil
	} else {
		return nil, ErrCrossChainTransferEventNotFound
	}
}
