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
	scta "github.com/thetatoken/thetasubchain/contracts/accessors"
	"github.com/thetatoken/thetasubchain/eth/abi"
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

var LockTypes = []InterChainMessageEventType{
	IMCEventTypeCrossChainTokenLockTFuel,
	IMCEventTypeCrossChainTokenLockTNT20,
	IMCEventTypeCrossChainTokenLockTNT721,
}

var VoucherBurnTypes = []InterChainMessageEventType{
	IMCEventTypeCrossChainVoucherBurnTFuel,
	IMCEventTypeCrossChainVoucherBurnTNT20,
	IMCEventTypeCrossChainVoucherBurnTNT721,
}

var UnlockTypes = []InterChainMessageEventType{
	IMCEventTypeCrossChainTokenUnlockTFuel,
	IMCEventTypeCrossChainTokenUnlockTNT20,
	IMCEventTypeCrossChainTokenUnlockTNT721,
}

var EventSelectors = map[InterChainMessageEventType]string{
	// TokenLock events
	IMCEventTypeCrossChainTokenLockTFuel:  crypto.Keccak256Hash([]byte("TFuelTokenLocked(uint256,address,address,uint256,uint256)")).Hex(),
	IMCEventTypeCrossChainTokenLockTNT20:  crypto.Keccak256Hash([]byte("TNT20TokenLocked(uint256,string,address,address,uint256,string,string,uint8,uint256)")).Hex(),
	IMCEventTypeCrossChainTokenLockTNT721: crypto.Keccak256Hash([]byte("TNT721TokenLocked(uint256,string,address,address,uint256,string,string,string,uint256)")).Hex(),

	// TokenUnlock events
	IMCEventTypeCrossChainTokenUnlockTFuel:  crypto.Keccak256Hash([]byte("TFuelTokenUnlocked(uint256,address,uint256,uint256,uint256)")).Hex(),
	IMCEventTypeCrossChainTokenUnlockTNT20:  crypto.Keccak256Hash([]byte("TNT20TokenUnlocked(uint256,string,address,uint256,uint256,uint256)")).Hex(),
	IMCEventTypeCrossChainTokenUnlockTNT721: crypto.Keccak256Hash([]byte("TNT721TokenUnlocked(uint256,string,address,uint256,uint256,uint256)")).Hex(),

	// VoucherBurn events
	IMCEventTypeCrossChainVoucherBurnTFuel:  crypto.Keccak256Hash([]byte("TFuelVoucherBurned(address,address,uint256,uint256)")).Hex(),
	IMCEventTypeCrossChainVoucherBurnTNT20:  crypto.Keccak256Hash([]byte("TNT20VoucherBurned(uint256,string,address,address,uint256,uint256)")).Hex(),
	IMCEventTypeCrossChainVoucherBurnTNT721: crypto.Keccak256Hash([]byte("TNT721VoucherBurned(uint256,string,address,address,uint256,uint256)")).Hex(),
}

func QueryInterChainEventLog(queriedChainID *big.Int, fromBlock *big.Int, toBlock *big.Int, contractAddr common.Address, imceType InterChainMessageEventType, url string) []*InterChainMessageEvent {
	var events []*InterChainMessageEvent
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
		case IMCEventTypeCrossChainTokenLockTFuel:
			extractTFuelTokenLockedEvent(queriedChainID, logData, &events)
		case IMCEventTypeCrossChainTokenLockTNT20:
			extractTNT20TokenLockedEvent(queriedChainID, logData, &events)
		case IMCEventTypeCrossChainTokenLockTNT721:
			extractTNT721TokenLockedEvent(queriedChainID, logData, &events)

		// TokenUnlock events
		case IMCEventTypeCrossChainTokenUnlockTFuel:
			extractTFuelTokenUnlockedEvent(queriedChainID, logData, &events)
		case IMCEventTypeCrossChainTokenUnlockTNT20:
			extractTNT20TokenUnlockedEvent(queriedChainID, logData, &events)
		case IMCEventTypeCrossChainTokenUnlockTNT721:
			extractTNT721TokenUnlockedEvent(queriedChainID, logData, &events)

		// VoucherBurn events
		case IMCEventTypeCrossChainVoucherBurnTFuel:
			extractTFuelVoucherBurnedEvent(queriedChainID, logData, &events)
		case IMCEventTypeCrossChainVoucherBurnTNT20:
			extractTNT20VoucherBurnedEvent(queriedChainID, logData, &events)
		case IMCEventTypeCrossChainVoucherBurnTNT721:
			extractTNT721VoucherBurnedEvent(queriedChainID, logData, &events)
		default:
		}
	}
	return events
}

func extractTFuelTokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma CrossChainTFuelTokenLockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTFuelTokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TFuelTokenLocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	event := &InterChainMessageEvent{
		Type:          IMCEventTypeCrossChainTokenLockTFuel,
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

func extractTNT20TokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma CrossChainTNT20TokenLockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTNT20TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT20TokenLocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	event := &InterChainMessageEvent{
		Type:          IMCEventTypeCrossChainTokenLockTNT20,
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

func extractTNT721TokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
	// data, _ := hex.DecodeString(logData.Data[2:])
	// var tma CrossChainTNT721TokenLockedEvent
	// contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTNT721TokenBankABI)))
	// contractAbi.UnpackIntoInterface(&tma, "TNT721TokenLocked", data)
	// blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	// event := &InterChainMessageEvent{
	// 	Type:          IMCEventTypeCrossChainTokenLockTNT20,
	// 	SourceChainID: sourceChainID,
	// 	TargetChainID: tma.TargetChainID,
	// 	Sender:        tma.SourceChainTokenSender,
	// 	Receiver:      tma.TargetChainVoucherReceiver,
	// 	Data:          data,
	// 	Nonce:         tma.Nonce,
	// 	BlockHeight:   blockHeight,
	// }
	// logger.Infof("got tnt20 event : %v, logdata : %v", tma, logData)
	// *events = append(*events, event)
}

func extractTFuelTokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma CrossChainTFuelTokenUnlockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTFuelTokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TFuelTokenUnlocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	event := &InterChainMessageEvent{
		Type:          IMCEventTypeCrossChainTokenUnlockTFuel,
		SourceChainID: tma.SourceChainID,
		TargetChainID: targetChainID,
		Sender:        common.Address{}, // not needed
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.TokenUnlockNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got tfuel unlock event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT20TokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma CrossChainTNT20TokenUnlockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTNT20TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT20TokenUnlocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	event := &InterChainMessageEvent{
		Type:          IMCEventTypeCrossChainTokenUnlockTNT20,
		SourceChainID: tma.SourceChainID,
		TargetChainID: targetChainID,
		Sender:        common.Address{}, // not needed
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.TokenUnlockNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got tnt20 unlock event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT721TokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {

}

func extractTFuelVoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma CrossChainTFuelVoucherBurnedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.SubchainTFuelTokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TFuelVoucherBurned", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	event := &InterChainMessageEvent{
		Type:          IMCEventTypeCrossChainVoucherBurnTFuel,
		SourceChainID: sourceChainID,
		TargetChainID: tma.TargetChainID,
		Sender:        tma.SourceChainVoucherOwner,
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.VoucherBurnNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got tfuel voucher burn event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT20VoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma CrossChainTFuelVoucherBurnedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTNT20TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT20VoucherBurned", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	event := &InterChainMessageEvent{
		Type:          IMCEventTypeCrossChainVoucherBurnTNT20,
		SourceChainID: sourceChainID,
		TargetChainID: tma.TargetChainID,
		Sender:        tma.SourceChainVoucherOwner,
		Receiver:      tma.TargetChainTokenReceiver,
		Data:          data,
		Nonce:         tma.VoucherBurnNonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got tnt20 voucher burn event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT721VoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {

}
