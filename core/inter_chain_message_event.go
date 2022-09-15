package core

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/rlp"
	"github.com/thetatoken/thetasubchain/eth/abi"
	scta "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

// var logger *log.Entry = log.WithFields(log.Fields{"prefix": "core"})

type InterChainMessageEventType uint64

const (
	IMCEventTypeUnknown InterChainMessageEventType = 0
	// 1 - 9999 reserved for future use

	IMCEventTypeCrossChainTokenLockTFuel   InterChainMessageEventType = 10001
	IMCEventTypeCrossChainTokenLockTNT20   InterChainMessageEventType = 10002
	IMCEventTypeCrossChainTokenLockTNT721  InterChainMessageEventType = 10003
	IMCEventTypeCrossChainTokenLockTNT1155 InterChainMessageEventType = 10004

	IMCEventTypeCrossChainVoucherMintTFuel   InterChainMessageEventType = 20001
	IMCEventTypeCrossChainVoucherMintTNT20   InterChainMessageEventType = 20002
	IMCEventTypeCrossChainVoucherMintTNT721  InterChainMessageEventType = 20003
	IMCEventTypeCrossChainVoucherMintTNT1155 InterChainMessageEventType = 20004

	IMCEventTypeCrossChainTokenUnlockTFuel   InterChainMessageEventType = 30001
	IMCEventTypeCrossChainTokenUnlockTNT20   InterChainMessageEventType = 30002
	IMCEventTypeCrossChainTokenUnlockTNT721  InterChainMessageEventType = 30003
	IMCEventTypeCrossChainTokenUnlockTNT1155 InterChainMessageEventType = 30004

	IMCEventTypeCrossChainVoucherBurnTFuel   InterChainMessageEventType = 40001
	IMCEventTypeCrossChainVoucherBurnTNT20   InterChainMessageEventType = 40002
	IMCEventTypeCrossChainVoucherBurnTNT721  InterChainMessageEventType = 40003
	IMCEventTypeCrossChainVoucherBurnTNT1155 InterChainMessageEventType = 40004
)

// InterChainMessageEvent represents an inter-chain messaging event.
type InterChainMessageEvent struct {
	Type          InterChainMessageEventType
	SourceChainID *big.Int
	TargetChainID *big.Int
	Sender        common.Address // sender of the message on the source chain
	Receiver      common.Address // receiver of the msssage on the target chain
	Data          common.Bytes   // generic data field that can be used to encode arbitrary data for inter-chain messaging
	Nonce         *big.Int
	BlockHeight   *big.Int
}

// NewInterChainMessageEvent creates a new inter-chain messaging event instance.
func NewInterChainMessageEvent(eventType InterChainMessageEventType, sourceChainID *big.Int, targetChainID *big.Int, sender common.Address, receiver common.Address,
	data common.Bytes, nonce *big.Int, blockHeight *big.Int) *InterChainMessageEvent {
	return &InterChainMessageEvent{eventType, sourceChainID, targetChainID, sender, receiver, data, nonce, blockHeight}
}

// ID returns the ID of the inter-chain messaging event.
func (c *InterChainMessageEvent) ID() string {
	eventStr := strconv.FormatUint(uint64(c.Type), 10) + "/" + c.SourceChainID.String() + "/" + c.Nonce.String()
	id := hex.EncodeToString(crypto.Keccak256([]byte(eventStr)))
	return id
}

// Equals checks whether an inter-chain messaging event is identical to the other
func (c *InterChainMessageEvent) Equals(x *InterChainMessageEvent) bool {
	if c.Type != x.Type {
		return false
	}
	if c.SourceChainID != x.SourceChainID {
		return false
	}
	if c.TargetChainID != x.TargetChainID {
		return false
	}
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
	if c.BlockHeight.Cmp(x.BlockHeight) != 0 {
		return false
	}
	return true
}

// String represents the string representation of the event
func (c *InterChainMessageEvent) String() string {
	return fmt.Sprintf("{ID: %v, Type: %v, SourceChainID: %v, TargetChainID: %v, Sender: %v, Receiver: %v,  Data: %v, Nonce: %v, BlockHeight: %v}",
		c.ID(), c.Type, c.SourceChainID, c.TargetChainID, c.Sender.Hex(), c.Receiver.Hex(), string(c.Data), c.Nonce.String(), c.BlockHeight.String())
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
		c.SourceChainID,
		c.TargetChainID,
		c.Sender,
		c.Receiver,
		c.Data,
		c.Nonce,
		c.BlockHeight,
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

	sourceChainID := big.NewInt(0)
	err = stream.Decode(&sourceChainID)
	if err != nil {
		return err
	}
	c.SourceChainID = sourceChainID

	targetChainID := big.NewInt(0)
	err = stream.Decode(&targetChainID)
	if err != nil {
		return err
	}
	c.TargetChainID = targetChainID

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

	blockHeight := big.NewInt(0)
	err = stream.Decode(blockHeight)
	if err != nil {
		return err
	}
	c.BlockHeight = blockHeight

	return stream.ListEnd()
}

// ------------------------------------ Cross-Chain: Token Lock --------------------------------------------

// Cross-Chain TFuel Lock

type CrossChainTFuelTokenLockedEvent struct { // corresponding to the "TFuelTokenLocked" event
	Denom                      string
	SourceChainTokenSender     common.Address
	TargetChainID              *big.Int // targetChain: the chain to send the token to (i.e. on which vouchers will be minted)
	TargetChainVoucherReceiver common.Address
	LockedAmount               *big.Int
	TokenLockNonce             *big.Int
}

func ParseToCrossChainTFuelTokenLockedEvent(icme *InterChainMessageEvent) (*CrossChainTFuelTokenLockedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTokenLockTFuel {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTFuelTokenLockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TFuelTokenLocked", icme.Data)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID.Cmp(originatedChainID) != 0 {
		// Token Lock events can only happen on the chain where the authenic token was deployed. Thus, the "source chain", i.e. where the token is sending from
		// needs to be the same as the "originated chain".
		return nil, fmt.Errorf("source chain ID mismatch for TFuel lock: %v vs %v", icme.SourceChainID, originatedChainID)
	}

	return &event, nil
}

// Cross-Chain TNT20 Lock

type CrossChainTNT20TokenLockedEvent struct { // corresponding to the "TNT20TokenLocked" event
	Denom                      string
	SourceChainTokenSender     common.Address
	TargetChainID              *big.Int // targetChain: the chain to send the token to (i.e. on which vouchers will be minted)
	TargetChainVoucherReceiver common.Address
	LockedAmount               *big.Int
	Name                       string
	Symbol                     string
	Decimals                   uint8
	TokenLockNonce             *big.Int
}

func ParseToCrossChainTNT20TokenLockedEvent(icme *InterChainMessageEvent) (*CrossChainTNT20TokenLockedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTokenLockTNT20 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT20TokenLockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT20TokenLocked", icme.Data)
	event.Denom = strings.ToLower(event.Denom)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID.Cmp(originatedChainID) != 0 {
		// Token Lock events can only happen on the chain where the authenic token was deployed. Thus, the "source chain", i.e. where the token is sending from
		// needs to be the same as the "originated chain".
		return nil, fmt.Errorf("source chain ID mismatch for TNT20 lock: %v vs %v", icme.SourceChainID, originatedChainID)
	}
	return &event, nil
}

// Cross-Chain TNT721 Lock

type CrossChainTNT721TokenLockedEvent struct { // corresponding to the "TNT721TokenLocked" event
	Denom                      string
	SourceChainTokenSender     common.Address
	TargetChainID              *big.Int // targetChain: the chain to send the token to (i.e. on which vouchers will be minted)
	TargetChainVoucherReceiver common.Address
	TokenID                    *big.Int
	Name                       string
	Symbol                     string
	TokenURI                   string
	TokenLockNonce             *big.Int
}

func ParseToCrossChainTNT721TokenLockedEvent(icme *InterChainMessageEvent) (*CrossChainTNT721TokenLockedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTokenLockTNT721 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT721TokenLockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT721TokenLocked", icme.Data)
	event.Denom = strings.ToLower(event.Denom)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID.Cmp(originatedChainID) != 0 {
		// Token Lock events can only happen on the chain where the authenic token was deployed. Thus, the "source chain", i.e. where the token is sending from
		// needs to be the same as the "originated chain".
		return nil, fmt.Errorf("source chain ID mismatch for TNT721 lock: %v vs %v", icme.SourceChainID, originatedChainID)
	}

	return &event, nil
	//return nil, nil // TODO: implementation
}

// Cross-Chain TNT1155 Lock

type CrossChainTNT1155TokenLockedEvent struct { // corresponding to the "TNT1155TokenLocked" event
	Denom                      string
	SourceChainTokenSender     common.Address
	TargetChainID              *big.Int // targetChain: the chain to send the token to (i.e. on which vouchers will be minted)
	TargetChainVoucherReceiver common.Address
	TokenID                    *big.Int
	LockedAmount               *big.Int
	TokenURI                   string
	TokenLockNonce             *big.Int
}

func ParseToCrossChainTNT1155TokenLockedEvent(icme *InterChainMessageEvent) (*CrossChainTNT1155TokenLockedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTokenLockTNT1155 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT1155TokenLockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT1155TokenLocked", icme.Data)
	event.Denom = strings.ToLower(event.Denom)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID.Cmp(originatedChainID) != 0 {
		// Token Lock events can only happen on the chain where the authenic token was deployed. Thus, the "source chain", i.e. where the token is sending from
		// needs to be the same as the "originated chain".
		return nil, fmt.Errorf("source chain ID mismatch for TNT1155 lock: %v vs %v", icme.SourceChainID, originatedChainID)
	}

	return &event, nil
	//return nil, nil // TODO: implementation
}

// ------------------------------------ Cross-Chain: Voucher Mint --------------------------------------------

type CrossChainTFuelVoucherMintedEvent struct { // corresponding to the "TFuelVoucherMinted" event
	Denom                      string
	TargetChainVoucherReceiver common.Address
	MintedAmount               *big.Int
	SourceChainTokenLockNonce  *big.Int
	VoucherMintNonce           *big.Int
}

func ParseToCrossChainTFuelVoucherMintedEvent(icme *InterChainMessageEvent) (*CrossChainTFuelVoucherMintedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainVoucherMintTFuel {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTFuelVoucherMintedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TFuelVoucherMinted", icme.Data)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("source chain ID mismatch for TNT20 voucher mint: %v vs %v", icme.SourceChainID, originatedChainID)
	}

	return &event, nil
}

type CrossChainTNT20VoucherMintedEvent struct { // corresponding to the "TNT20VoucherMinted" event
	Denom                      string
	TargetChainVoucherReceiver common.Address
	VoucherContract            common.Address
	MintedAmount               *big.Int
	SourceChainTokenLockNonce  *big.Int
	VoucherMintNonce           *big.Int
}

func ParseToCrossChainTNT20VoucherMintedEvent(icme *InterChainMessageEvent) (*CrossChainTNT20VoucherMintedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainVoucherMintTNT20 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT20VoucherMintedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT20VoucherMinted", icme.Data)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("source chain ID mismatch for TNT20 voucher mint: %v vs %v", icme.SourceChainID, originatedChainID)
	}

	return &event, nil
}

type CrossChainTNT721VoucherMintedEvent struct { // corresponding to the "TNT721TokenLocked" event
	Denom                      string
	TargetChainVoucherReceiver common.Address
	VoucherContract            common.Address
	TokenID                    *big.Int
	SourceChainTokenLockNonce  *big.Int
	VoucherMintNonce           *big.Int
}

func ParseToCrossChainTNT721VoucherMintedEvent(icme *InterChainMessageEvent) (*CrossChainTNT721VoucherMintedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainVoucherMintTNT721 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT721VoucherMintedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT721VoucherMinted", icme.Data)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("source chain ID mismatch for TNT721 voucher mint: %v vs %v", icme.SourceChainID, originatedChainID)
	}

	return &event, nil
}

type CrossChainTNT1155VoucherMintedEvent struct { // corresponding to the "TNT1155VoucherMinted" event
	Denom                      string
	TargetChainVoucherReceiver common.Address
	VoucherContract            common.Address
	TokenID                    *big.Int
	MintedAmount               *big.Int
	SourceChainTokenLockNonce  *big.Int
	VoucherMintNonce           *big.Int
}

func ParseToCrossChainTNT1155VoucherMintedEvent(icme *InterChainMessageEvent) (*CrossChainTNT1155VoucherMintedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainVoucherMintTNT1155 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT1155VoucherMintedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT1155VoucherMinted", icme.Data)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("source chain ID mismatch for TNT1155 voucher mint: %v vs %v", icme.SourceChainID, originatedChainID)
	}

	return &event, nil
}

// ------------------------------------ Cross-Chain: Voucher Burn --------------------------------------------

type VoucherBurnEventStatus byte

const (
	VoucherBurnEventStatusPending VoucherBurnEventStatus = VoucherBurnEventStatus(iota)
	VoucherBurnEventStatusProcessed
	VoucherBurnEventStatusFinalized
	VoucherBurnEventStatusFailed
)

type VoucherBurnEventStatusInfo struct {
	Type                     InterChainMessageEventType
	Nonce                    *big.Int
	Status                   VoucherBurnEventStatus
	LastProcessedBlockHeight *big.Int
	RetriedTime              uint
}

type CrossChainTFuelVoucherBurnedEvent struct { // corresponding to the "TFuelVoucherBurned" event
	Denom                    string
	SourceChainVoucherOwner  common.Address
	TargetChainTokenReceiver common.Address
	BurnedAmount             *big.Int
	VoucherBurnNonce         *big.Int
}

func ParseToCrossChainTFuelVoucherBurnedEvent(icme *InterChainMessageEvent) (*CrossChainTFuelVoucherBurnedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainVoucherBurnTFuel {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTFuelVoucherBurnedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TFuelVoucherBurned", icme.Data)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.TargetChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("target chain ID mismatch for TFuel voucher burn: %v vs %v", icme.TargetChainID, originatedChainID)
	}

	return &event, nil
}

type CrossChainTNT20VoucherBurnedEvent struct { // corresponding to the "TNT20VoucherBurned" event
	Denom                    string
	SourceChainVoucherOwner  common.Address
	TargetChainTokenReceiver common.Address
	BurnedAmount             *big.Int
	VoucherBurnNonce         *big.Int
}

func ParseToCrossChainTNT20VoucherBurnedEvent(icme *InterChainMessageEvent) (*CrossChainTNT20VoucherBurnedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainVoucherBurnTNT20 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT20VoucherBurnedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT20VoucherBurned", icme.Data)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.TargetChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("target chain ID mismatch for TNT20 voucher burn: %v vs %v", icme.TargetChainID, originatedChainID)
	}

	return &event, nil
}

type CrossChainTNT721VoucherBurnedEvent struct { // corresponding to the "TNT721VoucherBurned" event
	Denom                    string
	SourceChainVoucherOwner  common.Address
	TargetChainTokenReceiver common.Address
	TokenID                  *big.Int
	VoucherBurnNonce         *big.Int
}

func ParseToCrossChainTNT721VoucherBurnedEvent(icme *InterChainMessageEvent) (*CrossChainTNT721VoucherBurnedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainVoucherBurnTNT721 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT721VoucherBurnedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT721VoucherBurned", icme.Data)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.TargetChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("target chain ID mismatch for TNT20 voucher burn: %v vs %v", icme.TargetChainID, originatedChainID)
	}

	return &event, nil
}

type CrossChainTNT1155VoucherBurnedEvent struct { // corresponding to the "TNT1155VoucherBurned" event
	Denom                    string
	SourceChainVoucherOwner  common.Address
	TargetChainTokenReceiver common.Address
	TokenID                  *big.Int
	BurnedAmount             *big.Int
	VoucherBurnNonce         *big.Int
}

func ParseToCrossChainTNT1155VoucherBurnedEvent(icme *InterChainMessageEvent) (*CrossChainTNT1155VoucherBurnedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainVoucherBurnTNT1155 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT1155VoucherBurnedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT1155VoucherBurned", icme.Data)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.TargetChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("target chain ID mismatch for TNT1155 voucher burn: %v vs %v", icme.TargetChainID, originatedChainID)
	}

	return &event, nil
}

// ------------------------------------ Cross-Chain: Token Unlock --------------------------------------------

// Cross-Chain TFuel Unlock

type CrossChainTFuelTokenUnlockedEvent struct { // corresponding to the "TFuelTokenUnlocked" event
	Denom                       string
	TargetChainTokenReceiver    common.Address
	UnlockedAmount              *big.Int
	SourceChainVoucherBurnNonce *big.Int
	TokenUnlockNonce            *big.Int
}

func ParseToCrossChainTFuelTokenUnlockedEvent(icme *InterChainMessageEvent) (*CrossChainTFuelTokenUnlockedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTokenUnlockTFuel {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTFuelTokenUnlockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TFuelTokenUnlocked", icme.Data)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.TargetChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("target chain ID mismatch for TFuel unlock: %v vs %v", icme.TargetChainID, originatedChainID)
	}

	return &event, nil
}

// Cross-Chain TNT20 Unlock

type CrossChainTNT20TokenUnlockedEvent struct { // corresponding to the "TNT20TokenUnlocked" event
	Denom                       string
	TargetChainTokenReceiver    common.Address
	UnlockedAmount              *big.Int
	SourceChainVoucherBurnNonce *big.Int
	TokenUnlockNonce            *big.Int
}

func ParseToCrossChainTNT20TokenUnlockedEvent(icme *InterChainMessageEvent) (*CrossChainTNT20TokenUnlockedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTokenUnlockTNT20 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT20TokenUnlockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT20TokenUnlocked", icme.Data)
	event.Denom = strings.ToLower(event.Denom)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.TargetChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("target chain ID mismatch for TNT20 unlock: %v vs %v", icme.TargetChainID, originatedChainID)
	}
	return &event, nil
}

// Cross-Chain TNT721 Unlock

type CrossChainTNT721TokenUnlockedEvent struct { // corresponding to the "TNT721TokenUnlocked" event
	Denom                       string
	TargetChainTokenReceiver    common.Address
	TokenID                     *big.Int
	SourceChainVoucherBurnNonce *big.Int
	TokenUnlockNonce            *big.Int
}

func ParseToCrossChainTNT721TokenUnlockedEvent(icme *InterChainMessageEvent) (*CrossChainTNT721TokenUnlockedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTokenUnlockTNT721 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT721TokenUnlockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT721TokenUnlocked", icme.Data)
	event.Denom = strings.ToLower(event.Denom)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.TargetChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("target chain ID mismatch for TNT721 unlock: %v vs %v", icme.TargetChainID, originatedChainID)
	}
	return &event, nil
}

// Cross-Chain TNT1155 Unlock

type CrossChainTNT1155TokenUnlockedEvent struct { // corresponding to the "TNT1155TokenUnlocked" event
	Denom                       string
	TargetChainTokenReceiver    common.Address
	TokenID                     *big.Int
	UnlockedAmount              *big.Int
	SourceChainVoucherBurnNonce *big.Int
	TokenUnlockNonce            *big.Int
}

func ParseToCrossChainTNT1155TokenUnlockedEvent(icme *InterChainMessageEvent) (*CrossChainTNT1155TokenUnlockedEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTokenUnlockTNT1155 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var event CrossChainTNT1155TokenUnlockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&event, "TNT1155TokenUnlocked", icme.Data)
	event.Denom = strings.ToLower(event.Denom)
	if err := ValidateDenom(event.Denom); err != nil {
		return nil, err
	}
	originatedChainID, err := ExtractOriginatedChainIDFromDenom(event.Denom)
	if err != nil {
		return nil, err
	}
	if icme.TargetChainID.Cmp(originatedChainID) != 0 {
		return nil, fmt.Errorf("target chain ID mismatch for TNT1155 unlock: %v vs %v", icme.TargetChainID, originatedChainID)
	}
	return &event, nil
}

// ------------------------------------ Denom Utils ----------------------------------------------

type CrossChainTokenType int

const (
	CrossChainTokenTypeInvalid CrossChainTokenType = -1
	CrossChainTokenTypeTFuel   CrossChainTokenType = 0
	CrossChainTokenTypeTNT20   CrossChainTokenType = 20
	CrossChainTokenTypeTNT721  CrossChainTokenType = 721
	CrossChainTokenTypeTNT1155 CrossChainTokenType = 1155
)

const tfuelAddressPlaceholder = "0x0000000000000000000000000000000000000000"

// originatedChainID: the chainID of the chain that the token was originated
func TFuelDenom(originatedChainID *big.Int) string {
	return strings.ToLower(fmt.Sprintf("%v/%v/%v", originatedChainID.String(), CrossChainTokenTypeTFuel, tfuelAddressPlaceholder)) // normalize to lower case to prevent duplication
}

func TNT20Denom(originatedChainID *big.Int, contractAddress common.Address) string {
	return strings.ToLower(fmt.Sprintf("%v/%v/%v", originatedChainID.String(), CrossChainTokenTypeTNT20, contractAddress.Hex())) // normalize to lower case to prevent duplication
}

func TNT721Denom(originatedChainID *big.Int, contractAddress common.Address) string {
	return strings.ToLower(fmt.Sprintf("%v/%v/%v", originatedChainID.String(), CrossChainTokenTypeTNT721, contractAddress.Hex())) // normalize to lower case to prevent duplication
}

func TNT1155Denom(originatedChainID *big.Int, contractAddress common.Address) string {
	return strings.ToLower(fmt.Sprintf("%v/%v/%v", originatedChainID.String(), CrossChainTokenTypeTNT1155, contractAddress.Hex())) // normalize to lower case to prevent duplication
}

func ValidateDenom(denom string) error {
	if !isLowerCase(denom) {
		return fmt.Errorf("invalid denom (must be lower case): %v", denom)
	}

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
	case CrossChainTokenTypeTNT1155:
		if !common.IsHexAddress(parts[2]) {
			return fmt.Errorf("invalid TNT1155 denom: %v", denom)
		}
	default:
		return fmt.Errorf("invalid denom (unknown token type): %v", denom)
	}

	return nil
}

func ExtractOriginatedChainIDFromDenom(denom string) (*big.Int, error) {
	parts := strings.Split(denom, "/")
	if len(parts) != 3 {
		return big.NewInt(0), fmt.Errorf("invalid denom: %v", denom)
	}

	chainID, ok := big.NewInt(0).SetString(parts[0], 10)
	if !ok {
		return big.NewInt(0), fmt.Errorf("invalid denom: %v", denom)
	}

	return chainID, nil
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

func isLowerCase(str string) bool {
	return str == strings.ToLower(str)
}
