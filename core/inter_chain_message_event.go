package core

import (
	// "bytes"
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
	"sync"

	// "sort"

	// log "github.com/sirupsen/logrus"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/rlp"
	ts "github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/theta/store/kvstore"
)

// var logger *log.Entry = log.WithFields(log.Fields{"prefix": "core"})

type InterChainMessageEventType uint64

const (
	IMCEventTypeUnknown                  InterChainMessageEventType = 0
	IMCEventTypeCrossChainTFuelTransfer  InterChainMessageEventType = 10001
	IMCEventTypeCrossChainTNT20Transfer  InterChainMessageEventType = 10002
	IMCEventTypeCrossChainTNT721Transfer InterChainMessageEventType = 10003
)

// InterChainMessageEvent contains the public information of a crosschain transfer event.
type InterChainMessageEvent struct {
	Type        InterChainMessageEventType
	Sender      common.Address
	Receiver    common.Address
	Data        common.Bytes // generic data field that can be used to encode arbitrary data for inter-chain messaging
	Nonce       *big.Int
	BlockNumber *big.Int
}

// NewInterChainMessageEvent creates a new crosschain transfer event instance.
func NewInterChainMessageEvent(eventType InterChainMessageEventType, sender common.Address, receiver common.Address,
	data common.Bytes, amount *big.Int, nonce *big.Int, blockNumber *big.Int) InterChainMessageEvent {
	return InterChainMessageEvent{eventType, sender, receiver, data, nonce, blockNumber}
}

// ID returns the ID of the crosschain transfer event, which is the string representation of its address.
func (c *InterChainMessageEvent) ID() *big.Int {
	return c.Nonce
}

// Equals checks whether the crosschain transfer event is the same as another crosschain transfer event
func (c *InterChainMessageEvent) Equals(x *InterChainMessageEvent) bool {
	if c.Nonce.Cmp(x.Nonce) != 0 {
		return false
	}
	if c.Sender.Hex() != x.Sender.Hex() {
		return false
	}
	if c.Receiver.Hex() != x.Receiver.Hex() {
		return false
	}
	if !bytes.Equal(c.Data, x.Data) {
		return false
	}
	if c.BlockNumber.Cmp(x.BlockNumber) != 0 {
		return false
	}
	return true
}

// String represents the string representation of the validator
func (c *InterChainMessageEvent) String() string {
	return fmt.Sprintf("{ID: %v, Type: %v, Sender: %v, Receiver: %v,  Data: %v, Nonce: %v, BlockNumber: %v}",
		c.ID(), c.Type, c.Sender.Hex(), c.Receiver.Hex(), string(c.Data), c.Nonce.String(), c.BlockNumber.String())
}

// ByID implements sort.Interface for InterChainMessageEvent based on ID (Nonce).
type ICMEByID []InterChainMessageEvent

func (b ICMEByID) Len() int           { return len(b) }
func (b ICMEByID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ICMEByID) Less(i, j int) bool { return b[i].Nonce.Cmp(b[j].Nonce) < 0 }

var _ rlp.Encoder = (*InterChainMessageEvent)(nil)

// EncodeRLP implements RLP Encoder interface.
func (c *InterChainMessageEvent) EncodeRLP(w io.Writer) error {
	if c == nil {
		return rlp.Encode(w, &InterChainMessageEvent{})
	}
	return rlp.Encode(w, []interface{}{
		c.Type,
		c.Sender,
		c.Receiver,
		c.Data,
		c.Nonce,
		c.BlockNumber,
	})
}

// DecodeRLP implements RLP Decoder interface.
func (c *InterChainMessageEvent) DecodeRLP(stream *rlp.Stream) error {
	_, err := stream.List()
	if err != nil {
		return err
	}

	eventType := InterChainMessageEventType(0)
	err = stream.Decode(&eventType)
	if err != nil {
		return err
	}
	c.Type = eventType

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

	data := common.Bytes{}
	err = stream.Decode(&data)
	if err != nil {
		return err
	}
	c.Data = data

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

// ------------------------------------ Inter-Chain Event Cache ----------------------------------------------

var (
	// ErrInterChainMessageEventrNotFound for ID is not found in crosschain transfer event set.
	ErrInterChainMessageEventNotFound      = errors.New("InterChainMessageEventNotFound")
	ErrInterChainMessageEventExisted       = errors.New("InterChainMessageEventrExisted")
	ErrInterChainMessageEventPersistFailed = errors.New("InterChainMessageEventPersistFailed")
)

// InterChainEventIndexKey constructs the DB key for the given block hash.
func InterChainEventIndexKey(nonce *big.Int) common.Bytes {
	return common.Bytes("ice/" + nonce.String())
}

type InterChainEventCache struct {
	mutex *sync.Mutex // mutex to for concurrency protection, e.g., the witness thread and consensus thread may access it concurrently
	db    database.Database
}

// NewInterChainEventCache creates a new crosschain transfer event cache instance.
func NewInterChainEventCache(db database.Database) *InterChainEventCache {
	cache := &InterChainEventCache{
		mutex: &sync.Mutex{},
		db:    db,
	}
	return cache
}

func (c *InterChainEventCache) Insert(event *InterChainMessageEvent) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(InterChainEventIndexKey(event.Nonce), event)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) Delete(nonce *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Delete(InterChainEventIndexKey(nonce))
	return err // the caller should handle the error
}

func (c *InterChainEventCache) Get(nonce *big.Int) (*InterChainMessageEvent, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	event := InterChainMessageEvent{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(InterChainEventIndexKey(nonce), &event)
	return &event, err // the caller should handle the error
}

func (c *InterChainEventCache) Exists(nonce *big.Int) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	event := InterChainMessageEvent{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(InterChainEventIndexKey(nonce), &event)
	if err == nil {
		return true, nil
	}

	if err == ts.ErrKeyNotFound {
		return false, nil
	}

	return false, err // the caller should handle the error
}

// ------------------------------------ Cross-Chain Asset Transfer --------------------------------------------

// Cross-Chain TFuel Transfer

type TfuelTransferMetaData struct {
	Denom  string
	Amount *big.Int
}

type CrossChainTFuelTransferEvent struct {
	Sender      common.Address
	Receiver    common.Address
	Denom       string
	Amount      *big.Int
	Nonce       *big.Int
	BlockNumber *big.Int
}

func NewCrossChainTFuelTransferEvent(sender common.Address, receiver common.Address, denom string,
	amount *big.Int, nonce *big.Int, blockNumber *big.Int) *CrossChainTFuelTransferEvent {
	return &CrossChainTFuelTransferEvent{sender, receiver, denom, amount, nonce, blockNumber}
}

func ParseToCrossChainTFuelTransferEvent(icme *InterChainMessageEvent) (*CrossChainTFuelTransferEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTFuelTransfer {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var tma TfuelTransferMetaData
	if err := rlp.DecodeBytes(icme.Data, &tma); err != nil {
		return nil, err
	}
	if err := ValidateDenom(tma.Denom); err != nil {
		return nil, err
	}

	ccatEvent := NewCrossChainTFuelTransferEvent(icme.Sender, icme.Receiver, tma.Denom, tma.Amount, icme.Nonce, icme.BlockNumber)
	return ccatEvent, nil
}

// Cross-Chain TNT20 Transfer

type TNT20TransferMetaData struct {
	Denom    string
	Name     string
	Symbol   string
	Decimals uint8
	Amount   *big.Int
}

type CrossChainTNT20TransferEvent struct {
	Sender      common.Address
	Receiver    common.Address
	Denom       string
	Name        string
	Symbol      string
	Decimals    uint8
	Amount      *big.Int
	Nonce       *big.Int
	BlockNumber *big.Int
}

func NewCrossChainTNT20TransferEvent(sender common.Address, receiver common.Address, denom string,
	name string, symbol string, decimals uint8, amount *big.Int, nonce *big.Int, blockNumber *big.Int) *CrossChainTNT20TransferEvent {
	return &CrossChainTNT20TransferEvent{sender, receiver, denom, name, symbol, decimals, amount, nonce, blockNumber}
}

func ParseToCrossChainTNT20TransferEvent(icme *InterChainMessageEvent) (*CrossChainTNT20TransferEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTNT20Transfer {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var tma TNT20TransferMetaData
	if err := rlp.DecodeBytes(icme.Data, &tma); err != nil {
		return nil, err
	}
	if err := ValidateDenom(tma.Denom); err != nil {
		return nil, err
	}

	ccatEvent := NewCrossChainTNT20TransferEvent(icme.Sender, icme.Receiver, tma.Denom, tma.Name, tma.Symbol, tma.Decimals, tma.Amount, icme.Nonce, icme.BlockNumber)
	return ccatEvent, nil
}

// ------------------------------------ Denom Utils ----------------------------------------------

type CrossChainTokenType int

const (
	CrossChainTokenTypeInvalid CrossChainTokenType = -1
	CrossChainTokenTypeTFuel   CrossChainTokenType = 0
	CrossChainTokenTypeTNT20   CrossChainTokenType = 1
	CrossChainTokenTypeTNT721  CrossChainTokenType = 2
)

const tfuelAddressPlaceholder = "0x0000000000000000000000000000000000000000"

// sourceChainID: the chainID of the chain that the token was originated
func TFuelDenom(sourceChainID string) string {
	return fmt.Sprintf("%v/%v/%v", sourceChainID, CrossChainTokenTypeTFuel, tfuelAddressPlaceholder)
}

func TNT20Denom(sourceChainID string, contractAddress common.Address) string {
	return fmt.Sprintf("%v/%v/%v", sourceChainID, CrossChainTokenTypeTNT20, strings.ToLower(contractAddress.Hex())) // normalize the address to lower case to prevent duplication
}

func TNT721Denom(sourceChainID string, contractAddress common.Address) string {
	return fmt.Sprintf("%v/%v/%v", sourceChainID, CrossChainTokenTypeTNT721, strings.ToLower(contractAddress.Hex())) // normalize the address to lower case to prevent duplication
}

func ValidateDenom(denom string) error {
	parts := strings.Split(denom, "/")
	if len(parts) != 3 {
		return fmt.Errorf("invalid denom (incorrect format): %v", denom)
	}

	tokenType, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid denom (failed to parse token type): %v, %v", denom, err)
	}

	switch CrossChainTokenType(tokenType) {
	case CrossChainTokenTypeTFuel:
		if parts[2] != tfuelAddressPlaceholder {
			return fmt.Errorf("invalid TFuel denom: %v", denom)
		}
	case CrossChainTokenTypeTNT20:
		if !common.IsHexAddress(parts[2]) {
			return fmt.Errorf("invalid TNT20 denom: %v", denom)
		}
	case CrossChainTokenTypeTNT721:
		if !common.IsHexAddress(parts[2]) {
			return fmt.Errorf("invalid TNT20 denom: %v", denom)
		}
	default:
		return fmt.Errorf("invalid denom (unknown token type): %v", denom)
	}

	return nil
}

func ExtractSourceChainIDFromDenom(denom string) (string, error) {
	parts := strings.Split(denom, "/")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid denom: %v", denom)
	}

	return parts[0], nil
}

func ExtractCrossChainTokenTypeFromDenom(denom string) (CrossChainTokenType, error) {
	parts := strings.Split(denom, "/")
	if len(parts) != 3 {
		return CrossChainTokenTypeInvalid, fmt.Errorf("invalid denom: %v", denom)
	}
	tokenType, err := strconv.Atoi(parts[1])
	if err != nil {
		return CrossChainTokenTypeInvalid, fmt.Errorf("invalid denom: %v", denom)
	}

	if (tokenType != int(CrossChainTokenTypeTFuel)) && (tokenType != int(CrossChainTokenTypeTNT20)) && (tokenType != int(CrossChainTokenTypeTNT721)) {
		return CrossChainTokenTypeInvalid, fmt.Errorf("invalid denom: %v", denom)
	}

	return CrossChainTokenType(tokenType), nil
}
