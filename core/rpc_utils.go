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

var TransferTypes = []InterChainMessageEventType{
	IMCEventTypeCrossChainLockTFuel,
	IMCEventTypeCrossChainLockTNT20,
	IMCEventTypeCrossChainLockTNT721,
}

var VoucherBurnTypes = []InterChainMessageEventType{
	IMCEventTypeCrossChainVoucherBurnTFuel,
	IMCEventTypeCrossChainVoucherBurnTNT20,
	IMCEventTypeCrossChainVoucherBurnTNT721,
}

var UnlockTypes = []InterChainMessageEventType{
	IMCEventTypeCrossChainUnlockTFuel,
	IMCEventTypeCrossChainUnlockTNT20,
	IMCEventTypeCrossChainUnlockTNT721,
}

var EventSelectors = map[InterChainMessageEventType]string{
	IMCEventTypeCrossChainLockTFuel: crypto.Keccak256Hash([]byte("TFuelTokenLocked(uint256,address,address,uint256,uint256,string)")).Hex(),
	IMCEventTypeCrossChainLockTNT20: crypto.Keccak256Hash([]byte("TNT20TokenLocked(uint256,address,address,uint256,address,string,string,uint8,uint256,string)")).Hex(),
	// IMCEventTypeCrossChainLockTNT721 : crypto.Keccak256Hash([]byte("")).Hex(),
	IMCEventTypeCrossChainVoucherBurnTFuel:  crypto.Keccak256Hash([]byte("BurnTFuelVouchers(address,address,uint256,uint256)")).Hex(),
	IMCEventTypeCrossChainVoucherBurnTNT20:  crypto.Keccak256Hash([]byte("BurnTNT20Vouchers(string,address,address,uint256,uint256)")).Hex(),
	IMCEventTypeCrossChainVoucherBurnTNT721: crypto.Keccak256Hash([]byte("BurnTNT721Vouchers(string,address,address,uint256,uint256)")).Hex(),
	IMCEventTypeCrossChainUnlockTFuel:       crypto.Keccak256Hash([]byte("TFuelTokenUnlocked(uint256,address,uint256,uint256)")).Hex(),
	IMCEventTypeCrossChainUnlockTNT20:       crypto.Keccak256Hash([]byte("")).Hex(),
	IMCEventTypeCrossChainUnlockTNT721:      crypto.Keccak256Hash([]byte("")).Hex(),
}

func QueryInterChainEventLog(fromBlock *big.Int, toBlock *big.Int, contractAddr common.Address, imceType InterChainMessageEventType, url string) []*InterChainMessageEvent {
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
		data, _ := hex.DecodeString(logData.Data[2:])
		switch imceType {
		case IMCEventTypeCrossChainLockTFuel:
			var tma TfuelTransferMetaData
			contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTFuelTokenBankABI)))
			contractAbi.UnpackIntoInterface(&tma, "TFuelTokenLocked", data)
			sourceChainID, _ := ExtractSourceChainIDFromDenom(tma.Denom)
			blockNumber, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
			event := &InterChainMessageEvent{
				Type:          IMCEventTypeCrossChainLockTFuel,
				SourceChainID: sourceChainID,
				TargetChainID: tma.TargetChainID.String(),
				Sender:        tma.MainchainTokenSender,
				Receiver:      tma.SubchainTokenReceiver,
				Data:          data,
				Nonce:         tma.Nonce,
				BlockNumber:   blockNumber,
			}
			logger.Infof("got tfuel event : %v, logdata : %v", tma, logData)
			events = append(events, event)
		case IMCEventTypeCrossChainLockTNT20:
			var tma TNT20TransferMetaData
			contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTNT20TokenBankABI)))
			contractAbi.UnpackIntoInterface(&tma, "TNT20TokenLocked", data)
			sourceChainID, _ := ExtractSourceChainIDFromDenom(tma.Denom)
			blockNumber, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
			event := &InterChainMessageEvent{
				Type:          IMCEventTypeCrossChainLockTNT20,
				SourceChainID: sourceChainID,
				TargetChainID: tma.TargetChainID.String(),
				Sender:        tma.MainchainTokenSender,
				Receiver:      tma.SubchainTokenReceiver,
				Data:          data,
				Nonce:         tma.Nonce,
				BlockNumber:   blockNumber,
			}
			logger.Infof("got tnt20 event : %v, logdata : %v", tma, logData)
			events = append(events, event)
		case IMCEventTypeCrossChainLockTNT721:
		case IMCEventTypeCrossChainUnlockTFuel:
			var tma TFuelVoucherUnlockMetaData
			contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTFuelTokenBankABI)))
			contractAbi.UnpackIntoInterface(&tma, "TFuelTokenUnlocked", data)
			blockNumber, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
			event := &InterChainMessageEvent{
				Type:          IMCEventTypeCrossChainLockTFuel,
				SourceChainID: tma.SourceChainID.String(),
				TargetChainID: MainnetChainID,
				Sender:        tma.SubchainTokenSender,
				Receiver:      tma.MainchainTokenReceiver,
				Data:          data,
				Nonce:         tma.Nonce,
				BlockNumber:   blockNumber,
			}
			logger.Infof("got tfuel unlock event : %v, logdata : %v", tma, logData)
			events = append(events, event)
		default:
		}
	}
	return events
}

func QueryVoucherBurnEventLog(fromBlock *big.Int, toBlock *big.Int, contractAddr common.Address, imceType InterChainMessageEventType, url string, sourceChainID string, targetChainID string) []*InterChainMessageEvent {
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
		data, _ := hex.DecodeString(logData.Data[2:])
		switch imceType {
		case IMCEventTypeCrossChainVoucherBurnTFuel:
			var vma TFuelVoucherBurnMetaData
			contractAbi, _ := abi.JSON(strings.NewReader(string(scta.SubchainTFuelTokenBankABI)))
			contractAbi.UnpackIntoInterface(&vma, "BurnTFuelVouchers", data)
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
		case IMCEventTypeCrossChainVoucherBurnTNT20:
		default:
		}
	}

	return events
}
