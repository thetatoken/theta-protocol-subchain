package core

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/eth/abi"
	scta "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

// RPC related

type LogData struct {
	LogIndex         string   `json:"logIndex"`
	TransactionIndex string   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
	Type             string   `json:"type"`
}

// eventLogQueryTimeout bounds the eth_getLogs call. The default http.Client has no
// timeout at all, so an unresponsive node would hold the witness loop open forever.
const eventLogQueryTimeout = 30 * time.Second

type RPCResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int64  `json:"id"`
	// Result is a pointer so that an absent or null "result" is distinguishable from
	// an empty array. Both decode to a zero-length slice otherwise, which the caller
	// would treat as "this range contained no events" and advance past it.
	Result *[]LogData `json:"result"`
	Error  *RPCError  `json:"error"`
}

// RPCError carries a JSON-RPC error object. Without it, an error response
// unmarshals into a RPCResult with a nil Result and no indication that anything
// went wrong, which is indistinguishable from "this block range contained no
// events".
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// decodeLogData decodes a log's hex payload. logData.Data[2:] was previously sliced
// unconditionally, which panics on a value shorter than the "0x" prefix.
func decodeLogData(logData LogData) ([]byte, error) {
	if !strings.HasPrefix(logData.Data, "0x") {
		return nil, fmt.Errorf("log data is not 0x-prefixed: %.64q", logData.Data)
	}
	data, err := hex.DecodeString(logData.Data[2:])
	if err != nil {
		return nil, fmt.Errorf("log data is not valid hex: %v", err)
	}
	return data, nil
}

// parseLogBlockNumber parses a log's block number. The previous code discarded the
// failure, yielding an event with a nil BlockHeight that no longer corresponds to
// anything on chain.
func parseLogBlockNumber(logData LogData) (*big.Int, error) {
	if !strings.HasPrefix(logData.BlockNumber, "0x") {
		return nil, fmt.Errorf("log block number is not 0x-prefixed: %.64q", logData.BlockNumber)
	}
	blockHeight, ok := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	if !ok {
		return nil, fmt.Errorf("log block number is not valid hex: %.64q", logData.BlockNumber)
	}
	return blockHeight, nil
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("JSON-RPC error %v: %v", e.Code, e.Message)
}

type TransferEvent struct {
	Denom  string
	Amount *big.Int
	Nonce  *big.Int
}

var LockTypes = []score.InterChainMessageEventType{
	score.IMCEventTypeCrossChainTokenLockTFuel,
	score.IMCEventTypeCrossChainTokenLockTNT20,
	score.IMCEventTypeCrossChainTokenLockTNT721,
}

var VoucherBurnTypes = []score.InterChainMessageEventType{
	score.IMCEventTypeCrossChainVoucherBurnTFuel,
	score.IMCEventTypeCrossChainVoucherBurnTNT20,
	score.IMCEventTypeCrossChainVoucherBurnTNT721,
}

var UnlockTypes = []score.InterChainMessageEventType{
	score.IMCEventTypeCrossChainTokenUnlockTFuel,
	score.IMCEventTypeCrossChainTokenUnlockTNT20,
	score.IMCEventTypeCrossChainTokenUnlockTNT721,
}

var EventSelectors = map[score.InterChainMessageEventType]string{
	// TokenLock events
	score.IMCEventTypeCrossChainTokenLockTFuel:  crypto.Keccak256Hash([]byte("TFuelTokenLocked(string,address,uint256,address,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainTokenLockTNT20:  crypto.Keccak256Hash([]byte("TNT20TokenLocked(string,address,uint256,address,uint256,string,string,uint8,uint256)")).Hex(),
	score.IMCEventTypeCrossChainTokenLockTNT721: crypto.Keccak256Hash([]byte("TNT721TokenLocked(string,address,uint256,address,uint256,string,string,string,uint256)")).Hex(),

	// VoucherMint events
	score.IMCEventTypeCrossChainVoucherMintTFuel:  crypto.Keccak256Hash([]byte("TFuelVoucherMinted(string,address,uint256,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainVoucherMintTNT20:  crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainVoucherMintTNT721: crypto.Keccak256Hash([]byte("TNT721VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex(),

	// VoucherBurn events
	score.IMCEventTypeCrossChainVoucherBurnTFuel:  crypto.Keccak256Hash([]byte("TFuelVoucherBurned(string,address,address,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainVoucherBurnTNT20:  crypto.Keccak256Hash([]byte("TNT20VoucherBurned(string,address,address,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainVoucherBurnTNT721: crypto.Keccak256Hash([]byte("TNT721VoucherBurned(string,address,address,uint256,uint256)")).Hex(),

	// TokenUnlock events
	score.IMCEventTypeCrossChainTokenUnlockTFuel:  crypto.Keccak256Hash([]byte("TFuelTokenUnlocked(string,address,uint256,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainTokenUnlockTNT20:  crypto.Keccak256Hash([]byte("TNT20TokenUnlocked(string,address,uint256,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainTokenUnlockTNT721: crypto.Keccak256Hash([]byte("TNT721TokenUnlocked(string,address,uint256,uint256,uint256)")).Hex(),

	//TNT1155
	score.IMCEventTypeCrossChainTokenLockTNT1155:   crypto.Keccak256Hash([]byte("TNT1155TokenLocked(string,address,uint256,address,uint256,uint256,string,uint256)")).Hex(),
	score.IMCEventTypeCrossChainVoucherMintTNT1155: crypto.Keccak256Hash([]byte("TNT1155VoucherMinted(string,address,address,uint256,uint256,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainVoucherBurnTNT1155: crypto.Keccak256Hash([]byte("TNT1155VoucherBurned(string,address,address,uint256,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainTokenUnlockTNT1155: crypto.Keccak256Hash([]byte("TNT1155TokenUnlocked(string,address,uint256,uint256,uint256,uint256)")).Hex(),
}

// QueryInterChainEventLog returns the inter-chain events in [fromBlock, toBlock].
//
// It returns an error for every failure mode rather than an empty slice. The caller
// advances its scan checkpoint on the strength of this result, so "the query failed"
// and "the range contained no events" must not be conflated: doing so permanently
// skips the range.
func QueryInterChainEventLog(queriedChainID *big.Int, fromBlock *big.Int, toBlock *big.Int, tfuelTokenbankAddress common.Address, tnt20TokenBankAddress common.Address, tnt721TokenBankAddress common.Address, tnt1155TokenBankAddress common.Address, queryTopics string, url string) ([]*score.InterChainMessageEvent, error) {

	var events []*score.InterChainMessageEvent

	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":[%v],"topics":[[%v]]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), fmt.Sprintf("\"%v\",\"%v\",\"%v\",\"%v\"", tfuelTokenbankAddress, tnt20TokenBankAddress, tnt721TokenBankAddress, tnt1155TokenBankAddress), queryTopics)
	var jsonData = []byte(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to build the eth_getLogs request for %v: %v", url, err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: eventLogQueryTimeout}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("eth_getLogs request to %v failed: %v", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("eth_getLogs to %v returned HTTP %v", url, response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the eth_getLogs response from %v: %v", url, err)
	}

	var rpcres RPCResult
	if err := json.Unmarshal(body, &rpcres); err != nil {
		return nil, fmt.Errorf("failed to decode the eth_getLogs response from %v: %v (body: %.256q)", url, err, body)
	}
	if rpcres.Error != nil {
		return nil, fmt.Errorf("eth_getLogs to %v: %v", url, rpcres.Error)
	}
	if rpcres.Result == nil {
		return nil, fmt.Errorf("eth_getLogs to %v returned no result array (body: %.256q)", url, body)
	}

	for _, logData := range *rpcres.Result {
		logData := logData
		// A log with no topics cannot be one of ours, but indexing Topics[0] blindly
		// would panic on it. The query filters by topic, so this should not occur --
		// which is exactly why it must be reported rather than assumed away.
		if len(logData.Topics) == 0 {
			return nil, fmt.Errorf("eth_getLogs to %v returned a log with no topics at block %v", url, logData.BlockNumber)
		}
		switch logData.Topics[0] {

		// TokenLock events
		case EventSelectors[score.IMCEventTypeCrossChainTokenLockTFuel]:
			if err := extractTFuelTokenLockedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TFuelTokenLocked log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainTokenLockTNT20]:
			if err := extractTNT20TokenLockedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT20TokenLocked log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainTokenLockTNT721]:
			if err := extractTNT721TokenLockedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT721TokenLocked log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainTokenLockTNT1155]:
			if err := extractTNT1155TokenLockedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT1155TokenLocked log at block %v: %v", logData.BlockNumber, err)
			}

		// VoucherMint events
		case EventSelectors[score.IMCEventTypeCrossChainVoucherMintTFuel]:
			if err := extractTFuelVoucherMintedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TFuelVoucherMinted log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainVoucherMintTNT20]:
			if err := extractTNT20VoucherMintedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT20VoucherMinted log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainVoucherMintTNT721]:
			if err := extractTNT721VoucherMintedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT721VoucherMinted log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainVoucherMintTNT1155]:
			if err := extractTNT1155VoucherMintedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT1155VoucherMinted log at block %v: %v", logData.BlockNumber, err)
			}

		// VoucherBurn events
		case EventSelectors[score.IMCEventTypeCrossChainVoucherBurnTFuel]:
			if err := extractTFuelVoucherBurnedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TFuelVoucherBurned log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainVoucherBurnTNT20]:
			if err := extractTNT20VoucherBurnedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT20VoucherBurned log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainVoucherBurnTNT721]:
			if err := extractTNT721VoucherBurnedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT721VoucherBurned log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainVoucherBurnTNT1155]:
			if err := extractTNT1155VoucherBurnedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT1155VoucherBurned log at block %v: %v", logData.BlockNumber, err)
			}

		// TokenUnlock events
		case EventSelectors[score.IMCEventTypeCrossChainTokenUnlockTFuel]:
			if err := extractTFuelTokenUnlockedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TFuelTokenUnlocked log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainTokenUnlockTNT20]:
			if err := extractTNT20TokenUnlockedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT20TokenUnlocked log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainTokenUnlockTNT721]:
			if err := extractTNT721TokenUnlockedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT721TokenUnlocked log at block %v: %v", logData.BlockNumber, err)
			}
		case EventSelectors[score.IMCEventTypeCrossChainTokenUnlockTNT1155]:
			if err := extractTNT1155TokenUnlockedEvent(queriedChainID, logData, &events); err != nil {
				return nil, fmt.Errorf("malformed TNT1155TokenUnlocked log at block %v: %v", logData.BlockNumber, err)
			}

		default:
		}
	}
	return events, nil
}

func extractTFuelTokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTFuelTokenLockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TFuelTokenLocked", data); err != nil {
		return fmt.Errorf("failed to decode a TFuelTokenLocked log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainTokenLockTFuel,
		SourceChainID: sourceChainID,
		TargetChainID: tma.TargetChainID,
		Sender:        tma.SourceChainTokenSender,
		Receiver:      tma.TargetChainVoucherReceiver,
		Data:          data,
		Nonce:         tma.TokenLockNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TFuel locked event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT20TokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT20TokenLockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT20TokenLocked", data); err != nil {
		return fmt.Errorf("failed to decode a TNT20TokenLocked log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainTokenLockTNT20,
		SourceChainID: sourceChainID,
		TargetChainID: tma.TargetChainID,
		Sender:        tma.SourceChainTokenSender,
		Receiver:      tma.TargetChainVoucherReceiver,
		Data:          data,
		Nonce:         tma.TokenLockNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT20 locked event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT721TokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT721TokenLockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT721TokenLocked", data); err != nil {
		return fmt.Errorf("failed to decode a TNT721TokenLocked log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainTokenLockTNT721,
		SourceChainID: sourceChainID,
		TargetChainID: tma.TargetChainID,
		Sender:        tma.SourceChainTokenSender,
		Receiver:      tma.TargetChainVoucherReceiver,
		Data:          data,
		Nonce:         tma.TokenLockNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT721 locked event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT1155TokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT1155TokenLockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT1155TokenLocked", data); err != nil {
		return fmt.Errorf("failed to decode a TNT1155TokenLocked log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainTokenLockTNT1155,
		SourceChainID: sourceChainID,
		TargetChainID: tma.TargetChainID,
		Sender:        tma.SourceChainTokenSender,
		Receiver:      tma.TargetChainVoucherReceiver,
		Data:          data,
		Nonce:         tma.TokenLockNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT1155 locked event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTFuelVoucherMintedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTFuelVoucherMintedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TFuelVoucherMinted", data); err != nil {
		return fmt.Errorf("failed to decode a TFuelVoucherMinted log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	originatedChainID, err := score.ExtractOriginatedChainIDFromDenom(tma.Denom)
	if err != nil {
		logger.Warnf("Failed to extract originated chain ID from denom: %v", tma.Denom)
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherMintTFuel,
		SourceChainID: originatedChainID,
		TargetChainID: targetChainID,
		Sender:        common.Address{}, // don't care
		Receiver:      tma.TargetChainVoucherReceiver,
		Data:          data,
		Nonce:         tma.VoucherMintNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TFuel voucher mint event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT20VoucherMintedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT20VoucherMintedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT20VoucherMinted", data); err != nil {
		return fmt.Errorf("failed to decode a TNT20VoucherMinted log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	originatedChainID, err := score.ExtractOriginatedChainIDFromDenom(tma.Denom)
	if err != nil {
		logger.Warnf("Failed to extract originated chain ID from denom: %v", tma.Denom)
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherMintTNT20,
		SourceChainID: originatedChainID,
		TargetChainID: targetChainID,
		Sender:        common.Address{}, // don't care
		Receiver:      tma.TargetChainVoucherReceiver,
		Data:          data,
		Nonce:         tma.VoucherMintNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT20 voucher mint event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT721VoucherMintedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT721VoucherMintedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT721VoucherMinted", data); err != nil {
		return fmt.Errorf("failed to decode a TNT721VoucherMinted log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	originatedChainID, err := score.ExtractOriginatedChainIDFromDenom(tma.Denom)
	if err != nil {
		logger.Warnf("Failed to extract originated chain ID from denom: %v", tma.Denom)
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherMintTNT721,
		SourceChainID: originatedChainID,
		TargetChainID: targetChainID,
		Sender:        common.Address{}, // don't care
		Receiver:      tma.TargetChainVoucherReceiver,
		Data:          data,
		Nonce:         tma.VoucherMintNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT721 voucher mint event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT1155VoucherMintedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT1155VoucherMintedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT1155VoucherMinted", data); err != nil {
		return fmt.Errorf("failed to decode a TNT1155VoucherMinted log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	originatedChainID, err := score.ExtractOriginatedChainIDFromDenom(tma.Denom)
	if err != nil {
		logger.Warnf("Failed to extract originated chain ID from denom: %v", tma.Denom)
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherMintTNT1155,
		SourceChainID: originatedChainID,
		TargetChainID: targetChainID,
		Sender:        common.Address{}, // don't care
		Receiver:      tma.TargetChainVoucherReceiver,
		Data:          data,
		Nonce:         tma.VoucherMintNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT1155 voucher mint event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTFuelVoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTFuelVoucherBurnedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TFuelVoucherBurned", data); err != nil {
		return fmt.Errorf("failed to decode a TFuelVoucherBurned log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	originatedChainID, err := score.ExtractOriginatedChainIDFromDenom(tma.Denom)
	if err != nil {
		logger.Warnf("Failed to extract originated chain ID from denom: %v", tma.Denom)
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherBurnTFuel,
		SourceChainID: sourceChainID,
		TargetChainID: originatedChainID,
		Sender:        tma.SourceChainVoucherOwner,
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.VoucherBurnNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TFuel voucher burn event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT20VoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT20VoucherBurnedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT20VoucherBurned", data); err != nil {
		return fmt.Errorf("failed to decode a TNT20VoucherBurned log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	originatedChainID, err := score.ExtractOriginatedChainIDFromDenom(tma.Denom)
	if err != nil {
		logger.Warnf("Failed to extract originated chain ID from denom: %v", tma.Denom)
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherBurnTNT20,
		SourceChainID: sourceChainID,
		TargetChainID: originatedChainID,
		Sender:        tma.SourceChainVoucherOwner,
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.VoucherBurnNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT20 voucher burn event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT721VoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT721VoucherBurnedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT721VoucherBurned", data); err != nil {
		return fmt.Errorf("failed to decode a TNT721VoucherBurned log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	originatedChainID, err := score.ExtractOriginatedChainIDFromDenom(tma.Denom)
	if err != nil {
		logger.Warnf("Failed to extract originated chain ID from denom: %v", tma.Denom)
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherBurnTNT721,
		SourceChainID: sourceChainID,
		TargetChainID: originatedChainID,
		Sender:        tma.SourceChainVoucherOwner,
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.VoucherBurnNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT721 voucher burn event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT1155VoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT1155VoucherBurnedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT1155VoucherBurned", data); err != nil {
		return fmt.Errorf("failed to decode a TNT1155VoucherBurned log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	originatedChainID, err := score.ExtractOriginatedChainIDFromDenom(tma.Denom)
	if err != nil {
		logger.Warnf("Failed to extract originated chain ID from denom: %v", tma.Denom)
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainVoucherBurnTNT1155,
		SourceChainID: sourceChainID,
		TargetChainID: originatedChainID,
		Sender:        tma.SourceChainVoucherOwner,
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.VoucherBurnNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT1155 voucher burn event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTFuelTokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTFuelTokenUnlockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TFuelTokenUnlocked", data); err != nil {
		return fmt.Errorf("failed to decode a TFuelTokenUnlocked log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainTokenUnlockTFuel,
		SourceChainID: nil, // don't care
		TargetChainID: targetChainID,
		Sender:        common.Address{}, // don't care
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.TokenUnlockNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TFuel unlock event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT20TokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT20TokenUnlockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT20TokenUnlocked", data); err != nil {
		return fmt.Errorf("failed to decode a TNT20TokenUnlocked log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainTokenUnlockTNT20,
		SourceChainID: nil, // don't care
		TargetChainID: targetChainID,
		Sender:        common.Address{}, // don't care
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.TokenUnlockNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT20 unlock event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT721TokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT721TokenUnlockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT721TokenUnlocked", data); err != nil {
		return fmt.Errorf("failed to decode a TNT721TokenUnlocked log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainTokenUnlockTNT721,
		SourceChainID: nil, // don't care
		TargetChainID: targetChainID,
		Sender:        common.Address{}, // don't care
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.TokenUnlockNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT721 unlock event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}

func extractTNT1155TokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) error {
	data, err := decodeLogData(logData)
	if err != nil {
		return err
	}
	var tma score.CrossChainTNT1155TokenUnlockedEvent
	contractAbi, err := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	if err != nil {
		return err
	}
	if err := contractAbi.UnpackIntoInterface(&tma, "TNT1155TokenUnlocked", data); err != nil {
		return fmt.Errorf("failed to decode a TNT1155TokenUnlocked log: %v", err)
	}
	blockHeight, err := parseLogBlockNumber(logData)
	if err != nil {
		return err
	}
	event := &score.InterChainMessageEvent{
		Type:          score.IMCEventTypeCrossChainTokenUnlockTNT1155,
		SourceChainID: nil, // don't care
		TargetChainID: targetChainID,
		Sender:        common.Address{}, // don't care
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.TokenUnlockNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got TNT1155 unlock event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
	return nil
}
