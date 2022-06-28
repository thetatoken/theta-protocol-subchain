package core

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/rlp"
	scta "github.com/thetatoken/thetasubchain/contracts/accessors"
	"github.com/thetatoken/thetasubchain/eth/abi"
)

// var logger *log.Entry = log.WithFields(log.Fields{"prefix": "core"})

type InterChainMessageEventType uint64

const (
	IMCEventTypeUnknown InterChainMessageEventType = 0
	// 1 - 9999 reserved for future use

	IMCEventTypeCrossChainTokenLock       InterChainMessageEventType = 10000
	IMCEventTypeCrossChainTokenLockTFuel  InterChainMessageEventType = 10001
	IMCEventTypeCrossChainTokenLockTNT20  InterChainMessageEventType = 10002
	IMCEventTypeCrossChainTokenLockTNT721 InterChainMessageEventType = 10003

	IMCEventTypeCrossChainVoucherBurn       InterChainMessageEventType = 20000
	IMCEventTypeCrossChainVoucherBurnTFuel  InterChainMessageEventType = 20001
	IMCEventTypeCrossChainVoucherBurnTNT20  InterChainMessageEventType = 20002
	IMCEventTypeCrossChainVoucherBurnTNT721 InterChainMessageEventType = 20003

	IMCEventTypeCrossChainTokenUnlock       InterChainMessageEventType = 40000
	IMCEventTypeCrossChainTokenUnlockTFuel  InterChainMessageEventType = 40001
	IMCEventTypeCrossChainTokenUnlockTNT20  InterChainMessageEventType = 40002
	IMCEventTypeCrossChainTokenUnlockTNT721 InterChainMessageEventType = 40003
)

// InterChainMessageEvent contains the public information of a crosschain transfer event.
type InterChainMessageEvent struct {
	Type          InterChainMessageEventType
	SourceChainID *big.Int
	TargetChainID *big.Int
	Sender        common.Address
	Receiver      common.Address
	Data          common.Bytes // generic data field that can be used to encode arbitrary data for inter-chain messaging
	Nonce         *big.Int
	BlockNumber   *big.Int
}

// NewInterChainMessageEvent creates a new crosschain transfer event instance.
func NewInterChainMessageEvent(eventType InterChainMessageEventType, sourceChainID *big.Int, targetChainID *big.Int, sender common.Address, receiver common.Address,
	data common.Bytes, nonce *big.Int, blockNumber *big.Int) *InterChainMessageEvent {
	return &InterChainMessageEvent{eventType, sourceChainID, targetChainID, sender, receiver, data, nonce, blockNumber}
}

// ID returns the ID of the crosschain transfer event, which is the string representation of its address.
func (c *InterChainMessageEvent) ID() string {
	return strconv.FormatUint(uint64(c.Type), 10) + "/" + c.Nonce.String()
}

// Equals checks whether the crosschain transfer event is the same as another crosschain transfer event
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
	if c.BlockNumber.Cmp(x.BlockNumber) != 0 {
		return false
	}
	return true
}

// String represents the string representation of the validator
func (c *InterChainMessageEvent) String() string {
	return fmt.Sprintf("{ID: %v, Type: %v, SourceChainID: %v, TargetChainID: %v, Sender: %v, Receiver: %v,  Data: %v, Nonce: %v, BlockNumber: %v}",
		c.ID(), c.Type, c.SourceChainID, c.TargetChainID, c.Sender.Hex(), c.Receiver.Hex(), string(c.Data), c.Nonce.String(), c.BlockNumber.String())
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

	sourceChainID := big.NewInt(0)
	err = stream.Decode(&sourceChainID)
	if err != nil {
		return err
	}
	c.SourceChainID = sourceChainID

	targetChainID := ""
	err = stream.Decode(&targetChainID)
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

// ------------------------------------ Cross-Chain Voucher Burn --------------------------------------------
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

type VoucherBurnData struct {
	SubchainID  *big.Int
	Dynasty     *big.Int
	TxHash      common.Hash
	Denom       string
	Receiver    common.Address
	Amount      *big.Int
	TFuelAmount *big.Int
}

type TFuelVoucherBurnMetaData struct {
	VoucherOwner           common.Address
	MainchainTokenReceiver common.Address
	Amount                 *big.Int
	Nonce                  *big.Int
}

type TNT20VoucherBurnMetaData struct {
	TxHash  common.Hash
	Denom   string
	Amount  *big.Int
	TokenId *big.Int
}

type TFuelVoucherUnlockMetaData struct {
	SourceChainID          *big.Int
	MainchainTokenReceiver common.Address
	SubchainTokenSender    common.Address
	UnlockedAmount         *big.Int
	Nonce                  *big.Int
}

// ------------------------------------ Cross-Chain Asset Transfer --------------------------------------------

// Cross-Chain TFuel Transfer

type TfuelTransferMetaData struct {
	TargetChainID         *big.Int
	MainchainTokenSender  common.Address
	SubchainTokenReceiver common.Address
	LockedAmount          *big.Int
	Nonce                 *big.Int
	Denom                 string
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
	if icme.Type != IMCEventTypeCrossChainTokenLockTFuel {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var tma TfuelTransferMetaData
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.MainchainTFuelTokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&tma, "TFuelTokenLocked", icme.Data)
	if err := ValidateDenom(tma.Denom); err != nil {
		return nil, err
	}
	extractedSourceChainID, err := ExtractSourceChainIDFromDenom(tma.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID != extractedSourceChainID {
		return nil, fmt.Errorf("source chain ID mismatch for TFuel transfer: %v vs %v", icme.SourceChainID, extractedSourceChainID)
	}

	ccatEvent := NewCrossChainTFuelTransferEvent(icme.Sender, icme.Receiver, tma.Denom, tma.LockedAmount, icme.Nonce, icme.BlockNumber)
	return ccatEvent, nil
}

// Cross-Chain TNT20 Transfer

type TNT20TransferMetaData struct {
	TargetChainID         *big.Int
	MainchainTokenSender  common.Address
	SubchainTokenReceiver common.Address
	LockedAmount          *big.Int
	TNT20Contract         common.Address
	Name                  string
	Symbol                string
	Decimal               uint8
	Nonce                 *big.Int
	Denom                 string
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
	if icme.Type != IMCEventTypeCrossChainTokenLockTNT20 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var tma TNT20TransferMetaData
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.MainchainTNT20TokenBankABI)))
	if err != nil {
		return nil, err
	}
	contractAbi.UnpackIntoInterface(&tma, "TNT20TokenLocked", icme.Data)
	tma.Denom = strings.ToLower(tma.Denom)
	if err := ValidateDenom(tma.Denom); err != nil {
		return nil, err
	}
	extractedSourceChainID, err := ExtractSourceChainIDFromDenom(tma.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID != extractedSourceChainID {
		return nil, fmt.Errorf("source chain ID mismatch for TNT20 transfer: %v vs %v", icme.SourceChainID, extractedSourceChainID)
	}
	ccatEvent := NewCrossChainTNT20TransferEvent(icme.Sender, icme.Receiver, tma.Denom, tma.Name, tma.Symbol, tma.Decimal, tma.LockedAmount, icme.Nonce, icme.BlockNumber)
	return ccatEvent, nil
}

// Cross-Chain TNT721 Transfer

type TNT721TransferMetaData struct {
	Denom    string
	Name     string
	Symbol   string
	TokenID  *big.Int
	TokenURI string
}

type CrossChainTNT721TransferEvent struct {
	Sender      common.Address
	Receiver    common.Address
	Denom       string
	Name        string
	Symbol      string
	TokenID     *big.Int
	TokenURI    string
	Nonce       *big.Int
	BlockNumber *big.Int
}

func NewCrossChainTNT721TransferEvent(sender common.Address, receiver common.Address, denom string,
	name string, symbol string, tokenID *big.Int, tokenURI string, nonce *big.Int, blockNumber *big.Int) *CrossChainTNT721TransferEvent {
	return &CrossChainTNT721TransferEvent{sender, receiver, denom, name, symbol, tokenID, tokenURI, nonce, blockNumber}
}

func ParseToCrossChainTNT721TransferEvent(icme *InterChainMessageEvent) (*CrossChainTNT721TransferEvent, error) {
	if icme.Type != IMCEventTypeCrossChainTokenLockTNT721 {
		return nil, fmt.Errorf("invalid inter-chain message event type: %v", icme.Type)
	}

	var tma TNT721TransferMetaData
	if err := rlp.DecodeBytes(icme.Data, &tma); err != nil {
		return nil, err
	}
	if err := ValidateDenom(tma.Denom); err != nil {
		return nil, err
	}
	extractedSourceChainID, err := ExtractSourceChainIDFromDenom(tma.Denom)
	if err != nil {
		return nil, err
	}
	if icme.SourceChainID != extractedSourceChainID {
		return nil, fmt.Errorf("source chain ID mismatch for TNT721 transfer: %v vs %v", icme.SourceChainID, extractedSourceChainID)
	}

	ccatEvent := NewCrossChainTNT721TransferEvent(icme.Sender, icme.Receiver, tma.Denom, tma.Name, tma.Symbol, tma.TokenID, tma.TokenURI, icme.Nonce, icme.BlockNumber)
	return ccatEvent, nil
}

// ------------------------------------ Denom Utils ----------------------------------------------

type CrossChainTokenType int

const (
	CrossChainTokenTypeInvalid CrossChainTokenType = -1
	CrossChainTokenTypeTFuel   CrossChainTokenType = 0
	CrossChainTokenTypeTNT20   CrossChainTokenType = 20
	CrossChainTokenTypeTNT721  CrossChainTokenType = 721
)

const tfuelAddressPlaceholder = "0x0000000000000000000000000000000000000000"

// sourceChainID: the chainID of the chain that the token was originated
func TFuelDenom(sourceChainID string) string {
	return strings.ToLower(fmt.Sprintf("%v/%v/%v", sourceChainID, CrossChainTokenTypeTFuel, tfuelAddressPlaceholder)) // normalize to lower case to prevent duplication
}

func TNT20Denom(sourceChainID string, contractAddress common.Address) string {
	return strings.ToLower(fmt.Sprintf("%v/%v/%v", sourceChainID, CrossChainTokenTypeTNT20, contractAddress.Hex())) // normalize to lower case to prevent duplication
}

func TNT721Denom(sourceChainID string, contractAddress common.Address) string {
	return strings.ToLower(fmt.Sprintf("%v/%v/%v", sourceChainID, CrossChainTokenTypeTNT721, contractAddress.Hex())) // normalize to lower case to prevent duplication
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
	default:
		return fmt.Errorf("invalid denom (unknown token type): %v", denom)
	}

	return nil
}

func ExtractSourceChainIDFromDenom(denom string) (*big.Int, error) {
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
