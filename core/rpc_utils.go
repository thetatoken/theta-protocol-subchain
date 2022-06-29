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

func QueryInterChainEventLog(sourceChainID *big.Int, fromBlock *big.Int, toBlock *big.Int, contractAddr common.Address, imceType InterChainMessageEventType, url string) []*InterChainMessageEvent {
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
			extractTFuelTokenLockEvent(sourceChainID, logData, &events)
		case IMCEventTypeCrossChainTokenLockTNT20:
			extractTNT20TokenLockEvent(sourceChainID, logData, &events)
		case IMCEventTypeCrossChainTokenLockTNT721:
			extractTNT721TokenLockEvent(sourceChainID, logData, &events)

		// TokenUnlock events
		case IMCEventTypeCrossChainTokenUnlockTFuel:
			extractTFuelTokenUnlockEvent(sourceChainID, logData, &events)
		case IMCEventTypeCrossChainTokenUnlockTNT20:
			extractTNT20TokenUnlockEvent(sourceChainID, logData, &events)
		case IMCEventTypeCrossChainTokenUnlockTNT721:
			extractTNT721TokenUnlockEvent(sourceChainID, logData, &events)

		// VoucherBurn events
		case IMCEventTypeCrossChainVoucherBurnTFuel:
			extractTFuelVoucherBurnEvent(sourceChainID, logData, &events)
		case IMCEventTypeCrossChainVoucherBurnTNT20:
			extractTNT20VoucherBurnEvent(sourceChainID, logData, &events)
		case IMCEventTypeCrossChainVoucherBurnTNT721:
			extractTNT721VoucherBurnEvent(sourceChainID, logData, &events)
		default:
		}
	}
	return events
}

func extractTFuelTokenLockEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
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
		Nonce:         tma.Nonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got tfuel event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT20TokenLockEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
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
		Nonce:         tma.Nonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got tnt20 event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT721TokenLockEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
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

func extractTFuelTokenUnlockEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma TFuelTokenUnlockEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTFuelTokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TFuelTokenUnlocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	event := &InterChainMessageEvent{
		Type:          IMCEventTypeCrossChainTokenLockTFuel,
		SourceChainID: sourceChainID,
		TargetChainID: tma.TargetChainID,
		Sender:        tma.SubchainTokenSender,
		Receiver:      tma.MainchainTokenReceiver,
		Data:          data,
		Nonce:         tma.Nonce,
		BlockHeight:   blockHeight,
	}
	logger.Infof("got tfuel unlock event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT20TokenUnlockEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {

}

func extractTNT721TokenUnlockEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {

}

func extractTFuelVoucherBurnEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {
	var vma CrossChainTFuelVoucherBurnEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.SubchainTFuelTokenBankABI)))
	contractAbi.UnpackIntoInterface(&vma, "TFuelVoucherBurned", data)
	blockNumber, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
	event := &InterChainMessageEvent{
		Type:          IMCEventTypeCrossChainVoucherBurnTFuel,
		SourceChainID: sourceChainID,
		TargetChainID: targetChainID,
		Sender:        vma.VoucherOwner,
		Receiver:      vma.MainchainTokenReceiver,
		Data:          data,
		Nonce:         vma.Nonce,
		BlockNumber:   blockNumber,
	}
	logger.Infof("got tfuel voucher burn event : %v, logdata : %v", vma, logData)
	events = append(events, event)
}

func extractTNT20VoucherBurnEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {

}

func extractTNT721VoucherBurnEvent(sourceChainID *big.Int, logData LogData, events *[]*InterChainMessageEvent) {

}

// func QueryVoucherBurnEventLog(fromBlock *big.Int, toBlock *big.Int, contractAddr common.Address, imceType InterChainMessageEventType, url string, sourceChainID string, targetChainID string) []*InterChainMessageEvent {
// 	var events []*InterChainMessageEvent
// 	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), EventSelectors[imceType])
// 	var jsonData = []byte(queryStr)

// 	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	request.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		logger.Fatalf("response error : %v", err)
// 	}
// 	defer response.Body.Close()

// 	body, _ := ioutil.ReadAll(response.Body)
// 	var rpcres RPCResult

// 	err = json.Unmarshal(body, &rpcres)
// 	if err != nil {
// 		fmt.Printf("error decoding response: %v\n", err)
// 		if e, ok := err.(*json.SyntaxError); ok {
// 			fmt.Printf("syntax error at byte offset %d\n", e.Offset)
// 		}
// 		fmt.Printf("response: %q\n", body)
// 	}

// 	for _, logData := range rpcres.Result {
// 		logData := logData
// 		data, _ := hex.DecodeString(logData.Data[2:])
// 		switch imceType {
// 		case IMCEventTypeCrossChainVoucherBurnTFuel:
// 			var vma CrossChainTFuelVoucherBurnEvent
// 			contractAbi, _ := abi.JSON(strings.NewReader(string(scta.SubchainTFuelTokenBankABI)))
// 			contractAbi.UnpackIntoInterface(&vma, "TFuelVoucherBurned", data)
// 			blockNumber, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
// 			event := &InterChainMessageEvent{
// 				Type:          IMCEventTypeCrossChainVoucherBurnTFuel,
// 				SourceChainID: sourceChainID,
// 				TargetChainID: targetChainID,
// 				Sender:        vma.VoucherOwner,
// 				Receiver:      vma.MainchainTokenReceiver,
// 				Data:          data,
// 				Nonce:         vma.Nonce,
// 				BlockNumber:   blockNumber,
// 			}
// 			logger.Infof("got tfuel voucher burn event : %v, logdata : %v", vma, logData)
// 			events = append(events, event)
// 		case IMCEventTypeCrossChainVoucherBurnTNT20:
// 		default:
// 		}
// 	}

// 	return events
// }
