package core

import (
	// "bytes"
	"errors"
	"fmt"
	"io"
	"math/big"

	// "sort"

	// log "github.com/sirupsen/logrus"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/rlp"
	ts "github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/theta/store/kvstore"
)

// var logger *log.Entry = log.WithFields(log.Fields{"prefix": "core"})

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
func NewCrossChainTransferEvent(senderAddr string, receiverAddr string, denom string, amount *big.Int, nonce *big.Int, blockNumber *big.Int) CrossChainTransferEvent {
	sender := common.HexToAddress(senderAddr)
	receiver := common.HexToAddress(receiverAddr)
	return CrossChainTransferEvent{sender, receiver, denom, amount, nonce, blockNumber}
}

// ID returns the ID of the crosschain transfer event, which is the string representation of its address.
func (c CrossChainTransferEvent) ID() *big.Int {
	return c.Nonce
}

// Equals checks whether the crosschain transfer event is the same as another crosschain transfer event
func (c CrossChainTransferEvent) Equals(x CrossChainTransferEvent) bool {
	if c.Nonce.Cmp(x.Nonce) != 0 {
		return false
	}
	if c.Sender.Hex() != x.Sender.Hex() {
		return false
	}
	if c.Receiver.Hex() != x.Receiver.Hex() {
		return false
	}
	if c.BlockNumber.Cmp(x.BlockNumber) != 0 {
		return false
	}
	if c.Amount.Cmp(x.Amount) != 0 {
		return false
	}
	return true
}

// String represents the string representation of the validator
func (c CrossChainTransferEvent) String() string {
	return fmt.Sprintf("{ID: %v, Denom: %v, from: %v, to: %v}", c.ID(), c.Denom, c.Sender.Hex(), c.Receiver.Hex())
}

// ByID implements sort.Interface for CrossChainTransferEvent based on ID (Nonce).
type CCTEByID []CrossChainTransferEvent

func (b CCTEByID) Len() int           { return len(b) }
func (b CCTEByID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b CCTEByID) Less(i, j int) bool { return b[i].Nonce.Cmp(b[j].Nonce) < 0 }

var _ rlp.Encoder = (*CrossChainTransferEvent)(nil)

// EncodeRLP implements RLP Encoder interface.
func (c *CrossChainTransferEvent) EncodeRLP(w io.Writer) error {
	if c == nil {
		return rlp.Encode(w, &CrossChainTransferEvent{})
	}
	return rlp.Encode(w, []interface{}{
		c.Sender,
		c.Receiver,
		c.Denom,
		c.Amount,
		c.Nonce,
		c.BlockNumber,
	})
}

// DecodeRLP implements RLP Decoder interface.
func (c *CrossChainTransferEvent) DecodeRLP(stream *rlp.Stream) error {
	_, err := stream.List()
	if err != nil {
		return err
	}

	sender := &common.Address{}
	err = stream.Decode(sender)
	if err != nil {
		return err
	}
	c.Sender = *sender

	receiver := &common.Address{}
	err = stream.Decode(receiver)
	if err != nil {
		return err
	}
	c.Receiver = *receiver

	denom := ""
	err = stream.Decode(&denom)
	if err != nil {
		return err
	}
	c.Denom = denom

	amount := big.NewInt(0)
	err = stream.Decode(amount)
	if err != nil {
		return err
	}
	c.Amount = amount

	nonce := big.NewInt(0)
	err = stream.Decode(nonce)
	if err != nil {
		return err
	}
	c.Nonce = nonce

	blockNumber := big.NewInt(0)
	err = stream.Decode(blockNumber)
	if err != nil {
		return err
	}
	c.BlockNumber = blockNumber

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
	db database.Database
}

// NewCrossChainEventCache creates a new crosschain transfer event cache instance.
func NewCrossChainEventCache(db database.Database) *CrossChainEventCache {
	cache := &CrossChainEventCache{
		db: db,
	}
	return cache
}

// CrossChainEventIndexKey constructs the DB key for the given block hash.
func CrossChainEventIndexKey(nonce *big.Int) common.Bytes {
	return common.Bytes("cce/" + nonce.String())
}

func (c *CrossChainEventCache) Insert(event *CrossChainTransferEvent) error {
	store := kvstore.NewKVStore(c.db)
	err := store.Put(CrossChainEventIndexKey(event.Nonce), event)
	return err // the caller should handle the error
}

func (c *CrossChainEventCache) Delete(nonce *big.Int) error {
	store := kvstore.NewKVStore(c.db)
	err := store.Delete(CrossChainEventIndexKey(nonce))
	return err // the caller should handle the error
}

func (c *CrossChainEventCache) Get(nonce *big.Int) (*CrossChainTransferEvent, error) {
	event := CrossChainTransferEvent{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(CrossChainEventIndexKey(nonce), &event)
	return &event, err // the caller should handle the error
}

func (c *CrossChainEventCache) Exists(nonce *big.Int) (bool, error) {
	event := CrossChainTransferEvent{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(CrossChainEventIndexKey(nonce), &event)
	if err == nil {
		return true, nil
	}

	if err == ts.ErrKeyNotFound {
		return false, nil
	}

	return false, err // the caller should handle the error
}
