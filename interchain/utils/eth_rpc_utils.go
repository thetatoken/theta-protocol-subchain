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

type RPCResult struct {
	Jsonrpc string    `json:"jsonrpc"`
	Id      int64     `json:"id"`
	Result  []LogData `json:"result"`
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
}

func QueryInterChainEventLog(queriedChainID *big.Int, fromBlock *big.Int, toBlock *big.Int, contractAddr common.Address, imceType score.InterChainMessageEventType, url string) []*score.InterChainMessageEvent {
	var events []*score.InterChainMessageEvent
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), EventSelectors[imceType])
	var jsonData = []byte(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logger.Fatalf("response error : %v", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	var rpcres RPCResult

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Printf("error decoding response: %v\n", err)
		if e, ok := err.(*json.SyntaxError); ok {
			fmt.Printf("syntax error at byte offset %d\n", e.Offset)
		}
		fmt.Printf("response: %q\n", body)
	}

	for _, logData := range rpcres.Result {
		logData := logData
		switch imceType {

		// TokenLock events
		case score.IMCEventTypeCrossChainTokenLockTFuel:
			extractTFuelTokenLockedEvent(queriedChainID, logData, &events)
		case score.IMCEventTypeCrossChainTokenLockTNT20:
			extractTNT20TokenLockedEvent(queriedChainID, logData, &events)
		case score.IMCEventTypeCrossChainTokenLockTNT721:
			extractTNT721TokenLockedEvent(queriedChainID, logData, &events)

		// TokenUnlock events
		case score.IMCEventTypeCrossChainVoucherMintTFuel:
			extractTFuelVoucherMintedEvent(queriedChainID, logData, &events)
		case score.IMCEventTypeCrossChainVoucherMintTNT20:
			extractTNT20VoucherMintedEvent(queriedChainID, logData, &events)
		case score.IMCEventTypeCrossChainVoucherMintTNT721:
			extractTNT721VoucherMintedEvent(queriedChainID, logData, &events)

		// VoucherBurn events
		case score.IMCEventTypeCrossChainVoucherBurnTFuel:
			extractTFuelVoucherBurnedEvent(queriedChainID, logData, &events)
		case score.IMCEventTypeCrossChainVoucherBurnTNT20:
			extractTNT20VoucherBurnedEvent(queriedChainID, logData, &events)
		case score.IMCEventTypeCrossChainVoucherBurnTNT721:
			extractTNT721VoucherBurnedEvent(queriedChainID, logData, &events)

		// TokenUnlock events
		case score.IMCEventTypeCrossChainTokenUnlockTFuel:
			extractTFuelTokenUnlockedEvent(queriedChainID, logData, &events)
		case score.IMCEventTypeCrossChainTokenUnlockTNT20:
			extractTNT20TokenUnlockedEvent(queriedChainID, logData, &events)
		case score.IMCEventTypeCrossChainTokenUnlockTNT721:
			extractTNT721TokenUnlockedEvent(queriedChainID, logData, &events)

		default:
		}
	}
	return events
}

func extractTFuelTokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTFuelTokenLockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TFuelTokenLocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got tfuel event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT20TokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT20TokenLockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT20TokenLocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got tnt20 event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT721TokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT721TokenLockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT721TokenLocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got tnt721 event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTFuelVoucherMintedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTFuelVoucherMintedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TFuelVoucherMinted", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got tfuel voucher mint event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT20VoucherMintedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT20VoucherMintedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT20VoucherMinted", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got tfuel voucher mint event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT721VoucherMintedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	// data, _ := hex.DecodeString(logData.Data[2:])
	// var tma score.CrossChainTNT721TokenLockedEvent
	// contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	// contractAbi.UnpackIntoInterface(&tma, "TNT20VoucherMinted", data)
	// blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	// originatedChainID, err := score.ExtractOriginatedChainIDFromDenom(tma.Denom)
	// if err != nil {
	// 	logger.Warnf("Failed to extract originated chain ID from denom: %v", tma.Denom)
	// }
	// event := &score.InterChainMessageEvent{
	// 	Type:          score.IMCEventTypeCrossChainVoucherMintTNT20,
	// 	SourceChainID: originatedChainID,
	// 	TargetChainID: targetChainID,
	// 	Sender:        common.Address{}, // don't care
	// 	Receiver:      tma.TargetChainVoucherReceiver,
	// 	Data:          data,
	// 	Nonce:         tma.VoucherMintNonce,
	// 	BlockHeight:   blockHeight,
	// }
	// logger.Infof("got tnt721 voucher mint event : %v, logdata : %v", tma, logData)
	// *events = append(*events, event)
}

func extractTFuelVoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTFuelVoucherBurnedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TFuelVoucherBurned", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got tfuel voucher burn event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT20VoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTFuelVoucherBurnedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT20VoucherBurned", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got tnt20 voucher burn event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT721VoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {

}

func extractTFuelTokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTFuelTokenUnlockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TFuelTokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TFuelTokenUnlocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got tfuel unlock event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT20TokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT20TokenUnlockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT20TokenUnlocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got tnt20 unlock event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT721TokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {

}
