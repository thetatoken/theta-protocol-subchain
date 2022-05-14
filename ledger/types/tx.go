package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/rlp"
	score "github.com/thetatoken/thetasubchain/core"
)

const maxTxSize = 8 * 1024 * 1024

const (
	TxSubchainValidatorSetUpdate types.TxType = 201
	TxInterChainMessage          types.TxType = 202
)

//---------------------------------InterChainMessageTx--------------------------------------------

type InterChainMessageTx struct {
	Proposer    types.TxInput
	BlockNumber *big.Int
	Event       score.InterChainMessageEvent
}

type InterChainMessageTxJSON struct {
	Proposer    types.TxInput                `json:"proposer"`
	BlockNumber *big.Int                     `json:"block_number"`
	Event       score.InterChainMessageEvent `json:"event"`
}

func NewInterChainMessageTxJSON(a InterChainMessageTx) InterChainMessageTxJSON {
	return InterChainMessageTxJSON{
		Proposer:    a.Proposer,
		BlockNumber: a.BlockNumber,
		Event:       a.Event,
	}
}

func (a InterChainMessageTxJSON) CrossChainTransTx() InterChainMessageTx {
	return InterChainMessageTx{
		Proposer:    a.Proposer,
		BlockNumber: a.BlockNumber,
		Event:       a.Event,
	}
}

func (a InterChainMessageTxJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(InterChainMessageTxJSON(a))
}

func (a *InterChainMessageTx) UnmarshalJSON(data []byte) error {
	var b InterChainMessageTxJSON
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	*a = b.CrossChainTransTx()
	return nil
}

func (_ *InterChainMessageTx) AssertIsTx() {}

func (tx *InterChainMessageTx) SignBytes(chainID string) []byte {
	signBytes := encodeToBytes(chainID)
	sig := tx.Proposer.Signature
	tx.Proposer.Signature = nil
	txBytes, _ := TxToBytes(tx)
	signBytes = append(signBytes, txBytes...)
	signBytes = addPrefixForSignBytes(signBytes)

	tx.Proposer.Signature = sig
	return signBytes
}

func (tx *InterChainMessageTx) SetSignature(addr common.Address, sig *crypto.Signature) bool {
	if tx.Proposer.Address == addr {
		tx.Proposer.Signature = sig
		return true
	}
	return false
}

func (tx *InterChainMessageTx) String() string {
	return fmt.Sprintf("InterChainMessageTx with Nonce {%v}", tx.Event.Nonce)
}

//---------------------------------SubchainValidatorSetUpdateTx--------------------------------------------

type SubchainValidatorSetUpdateTx struct {
	Proposer   types.TxInput
	Dynasty    *big.Int
	Validators []score.Validator
}

type SubchainValidatorSetUpdateTxJSON struct {
	Proposer   types.TxInput     `json:"proposer"`
	Dynasty    *big.Int          `json:"dynasty"`
	Validators []score.Validator `json:"validators"`
}

func NewValidatorSetUpdateTxJSON(a SubchainValidatorSetUpdateTx) SubchainValidatorSetUpdateTxJSON {
	return SubchainValidatorSetUpdateTxJSON{
		Proposer:   a.Proposer,
		Dynasty:    a.Dynasty,
		Validators: a.Validators,
	}
}

func (a SubchainValidatorSetUpdateTxJSON) ValidatorSetUpdateTx() SubchainValidatorSetUpdateTx {
	return SubchainValidatorSetUpdateTx{
		Proposer:   a.Proposer,
		Dynasty:    a.Dynasty,
		Validators: a.Validators,
	}
}

func (a SubchainValidatorSetUpdateTxJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(SubchainValidatorSetUpdateTxJSON(a))
}

func (a *SubchainValidatorSetUpdateTx) UnmarshalJSON(data []byte) error {
	var b SubchainValidatorSetUpdateTxJSON
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	*a = b.ValidatorSetUpdateTx()
	return nil
}

func (_ *SubchainValidatorSetUpdateTx) AssertIsTx() {}

func (tx *SubchainValidatorSetUpdateTx) SignBytes(chainID string) []byte {
	signBytes := encodeToBytes(chainID)
	sig := tx.Proposer.Signature
	tx.Proposer.Signature = nil
	txBytes, _ := TxToBytes(tx)
	signBytes = append(signBytes, txBytes...)
	signBytes = addPrefixForSignBytes(signBytes)

	tx.Proposer.Signature = sig
	return signBytes
}

func (tx *SubchainValidatorSetUpdateTx) SetSignature(addr common.Address, sig *crypto.Signature) bool {
	if tx.Proposer.Address == addr {
		tx.Proposer.Signature = sig
		return true
	}
	return false
}

func (tx *SubchainValidatorSetUpdateTx) String() string {
	return fmt.Sprintf("SubchainValidatorSetUpdateTx{%v}", tx.Validators)
}

// --------------- Utils --------------- //

func encodeToBytes(str string) []byte {
	encodedBytes, err := rlp.EncodeToBytes(str)
	if err != nil {
		log.Panicf("Failed to encode %v: %v", str, err)
	}
	return encodedBytes
}

type EthereumTxWrapper struct {
	AccountNonce uint64          `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64          `json:"gas"      gencodec:"required"`
	Recipient    *common.Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"    gencodec:"required"`
	Payload      []byte          `json:"input"    gencodec:"required"`
}

// Need to add the following prefix to the tx signbytes to be compatible with
// the Ethereum tx format
func addPrefixForSignBytes(signBytes common.Bytes) common.Bytes {
	ethTx := EthereumTxWrapper{
		AccountNonce: uint64(0),
		Price:        new(big.Int).SetUint64(0),
		GasLimit:     uint64(0),
		Recipient:    &common.Address{},
		Amount:       new(big.Int).SetUint64(0),
		Payload:      signBytes,
	}
	signBytes, err := rlp.EncodeToBytes(ethTx)
	if err != nil {
		log.Panic(err)
	}
	return signBytes
}

func TxToBytes(t types.Tx) ([]byte, error) {
	var buf bytes.Buffer
	var txType types.TxType
	switch t.(type) {
	case *types.CoinbaseTx:
		txType = types.TxCoinbase
	case *types.SendTx:
		txType = types.TxSend
	case *types.SmartContractTx:
		txType = types.TxSmartContract
	case *SubchainValidatorSetUpdateTx:
		txType = TxSubchainValidatorSetUpdate
	case *InterChainMessageTx:
		txType = TxInterChainMessage
	default:
		return nil, errors.New("unsupported message type")
	}
	err := rlp.Encode(&buf, txType)
	if err != nil {
		return nil, err
	}
	err = rlp.Encode(&buf, t)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func TxFromBytes(raw []byte) (types.Tx, error) {
	var txType types.TxType
	buff := bytes.NewBuffer(raw)
	s := rlp.NewStream(buff, maxTxSize)
	err := s.Decode(&txType)
	if err != nil {
		return nil, err
	}
	if txType == types.TxCoinbase {
		data := &types.CoinbaseTx{}
		err = s.Decode(data)
		return data, err
	} else if txType == types.TxSend {
		data := &types.SendTx{}
		err = s.Decode(data)
		return data, err
	} else if txType == types.TxSmartContract {
		data := &types.SmartContractTx{}
		err = s.Decode(data)
		return data, err
	} else if txType == TxSubchainValidatorSetUpdate {
		data := &SubchainValidatorSetUpdateTx{}
		err = s.Decode(data)
		return data, err
	} else if txType == TxInterChainMessage {
		data := &InterChainMessageTx{}
		err = s.Decode(data)
		return data, err
	} else {
		return nil, fmt.Errorf("Unknown TX type: %v", txType)
	}
}
