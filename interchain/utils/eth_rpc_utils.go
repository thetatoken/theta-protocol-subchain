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

	//TNT1155
	score.IMCEventTypeCrossChainTokenLockTNT1155:   crypto.Keccak256Hash([]byte("TNT1155TokenLocked(string,address,uint256,address,uint256,uint256,string,uint256)")).Hex(),
	score.IMCEventTypeCrossChainVoucherMintTNT1155: crypto.Keccak256Hash([]byte("TNT1155VoucherMinted(string,address,address,uint256,uint256,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainVoucherBurnTNT1155: crypto.Keccak256Hash([]byte("TNT1155VoucherBurned(string,address,address,uint256,uint256,uint256)")).Hex(),
	score.IMCEventTypeCrossChainTokenUnlockTNT1155: crypto.Keccak256Hash([]byte("TNT1155TokenUnlocked(string,address,uint256,uint256,uint256,uint256)")).Hex(),
}

func QueryInterChainEventLog(queriedChainID *big.Int, fromBlock *big.Int, toBlock *big.Int, tfuelTokenbankAddress common.Address, tnt20TokenBankAddress common.Address, tnt721TokenBankAddress common.Address, tnt1155TokenBankAddress common.Address, queryTopics string, url string) []*score.InterChainMessageEvent {

	var events []*score.InterChainMessageEvent

	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":[%v],"topics":[[%v]]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), fmt.Sprintf("\"%v\",\"%v\",\"%v\",\"%v\"", tfuelTokenbankAddress, tnt20TokenBankAddress, tnt721TokenBankAddress, tnt1155TokenBankAddress), queryTopics)
	var jsonData = []byte(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		// logger.Fatal(err)
		logger.Warnf("Failed to post to %v, err: %v", url, err)
		return events // ignore, the query is repeated periodically anyway
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		// logger.Fatalf("response error : %v", err)
		logger.Warnf("RPC response error %v, err: %v", url, err)
		return events // ignore, the query is repeated periodically anyway
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
		switch logData.Topics[0] {

		// TokenLock events
		case EventSelectors[score.IMCEventTypeCrossChainTokenLockTFuel]:
			extractTFuelTokenLockedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainTokenLockTNT20]:
			extractTNT20TokenLockedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainTokenLockTNT721]:
			extractTNT721TokenLockedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainTokenLockTNT1155]:
			extractTNT1155TokenLockedEvent(queriedChainID, logData, &events)

		// VoucherMint events
		case EventSelectors[score.IMCEventTypeCrossChainVoucherMintTFuel]:
			extractTFuelVoucherMintedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainVoucherMintTNT20]:
			extractTNT20VoucherMintedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainVoucherMintTNT721]:
			extractTNT721VoucherMintedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainVoucherMintTNT1155]:
			extractTNT1155VoucherMintedEvent(queriedChainID, logData, &events)

		// VoucherBurn events
		case EventSelectors[score.IMCEventTypeCrossChainVoucherBurnTFuel]:
			extractTFuelVoucherBurnedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainVoucherBurnTNT20]:
			extractTNT20VoucherBurnedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainVoucherBurnTNT721]:
			extractTNT721VoucherBurnedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainVoucherBurnTNT1155]:
			extractTNT1155VoucherBurnedEvent(queriedChainID, logData, &events)

		// TokenUnlock events
		case EventSelectors[score.IMCEventTypeCrossChainTokenUnlockTFuel]:
			extractTFuelTokenUnlockedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainTokenUnlockTNT20]:
			extractTNT20TokenUnlockedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainTokenUnlockTNT721]:
			extractTNT721TokenUnlockedEvent(queriedChainID, logData, &events)
		case EventSelectors[score.IMCEventTypeCrossChainTokenUnlockTNT1155]:
			extractTNT1155TokenUnlockedEvent(queriedChainID, logData, &events)

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
	logger.Infof("got TFuel locked event : %v, logdata : %v", tma, logData)
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
	logger.Infof("got TNT20 locked event : %v, logdata : %v", tma, logData)
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
	logger.Infof("got TNT721 locked event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT1155TokenLockedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT1155TokenLockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT1155TokenLocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got TFuel voucher mint event : %v, logdata : %v", tma, logData)
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
	logger.Infof("got TNT20 voucher mint event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT721VoucherMintedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT721VoucherMintedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT721VoucherMinted", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
}

func extractTNT1155VoucherMintedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT1155VoucherMintedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT1155VoucherMinted", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got TFuel voucher burn event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT20VoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT20VoucherBurnedEvent
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
	logger.Infof("got TNT20 voucher burn event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT721VoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT721VoucherBurnedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT721VoucherBurned", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
}

func extractTNT1155VoucherBurnedEvent(sourceChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT1155VoucherBurnedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT1155VoucherBurned", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
	logger.Infof("got TFuel unlock event : %v, logdata : %v", tma, logData)
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
	logger.Infof("got TNT20 unlock event : %v, logdata : %v", tma, logData)
	*events = append(*events, event)
}

func extractTNT721TokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT721TokenUnlockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT721TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT721TokenUnlocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
}

func extractTNT1155TokenUnlockedEvent(targetChainID *big.Int, logData LogData, events *[]*score.InterChainMessageEvent) {
	data, _ := hex.DecodeString(logData.Data[2:])
	var tma score.CrossChainTNT1155TokenUnlockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT1155TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT1155TokenUnlocked", data)
	blockHeight, _ := new(big.Int).SetString(logData.BlockNumber[2:], 16)
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
}
