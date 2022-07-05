package main

import (
	"bytes"
	"log"

	// "encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"

	// "strings"

	scomm "github.com/thetatoken/thetasubchain/common"

	"github.com/spf13/viper"
	"github.com/thetatoken/theta/common"

	// scta "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
	score "github.com/thetatoken/thetasubchain/core"
	siu "github.com/thetatoken/thetasubchain/interchain/utils"
	// "github.com/thetatoken/thetasubchain/eth/abi"
	// "github.com/thetatoken/theta/crypto"
)

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

func queryEventLog(fromBlock *big.Int, toBlock *big.Int, contractAddr common.Address, imceType score.InterChainMessageEventType, interChainEveneCache *siu.InterChainEventCache) {
	url := viper.GetString(scomm.CfgSubchainEthRpcURL)
	queryStr := fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"eth_getLogs",
		"params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],
		"id":74
	}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), siu.EventSelectors[imceType])
	fmt.Println(queryStr)
	var jsonData = []byte(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	var rpcres RPCResult
	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rpcres.Result[0].Data)

	// for _, logData := range rpcres.Result {
	// 	logData := logData
	// 	h, _ := hex.DecodeString(logData.Data[2:])
	// 	// if ok, _ := interChainEveneCache.Exists(imceType, event.Nonce); ok {
	// 	// 	continue
	// 	// }
	// 	switch imceType {
	// 	case score.IMCEventTypeCrossChainTokenLockTFuel:
	// 		var tma score.TfuelTransferMetaData
	// 		contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTFuelTokenBankABI)))
	// 		contractAbi.UnpackIntoInterface(&tma, "TFuelTokenLocked", h)
	// 		fmt.Println(tma)
	// 	case score.IMCEventTypeCrossChainTokenLockTNT20:
	// 		var tma score.CrossChainTNT20TokenLockedEvent
	// 		contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTNT20TokenBankABI)))
	// 		contractAbi.UnpackIntoInterface(&tma, "TFuelTokenLocked", h)
	// 		fmt.Println(tma)
	// 	case score.IMCEventTypeCrossChainTokenLockTNT721:
	// 	default:
	// 	}
	// }
}

// func hex(i int) string {
// 	i64 := int64(i)
// 	return strconv.FormatInt(i64, 16)
// }

func main() {
	// jsonContent := `{"logIndex":1,"blockNumber":9,"blockHash":"0x5c6e3d707969617b11cd229554f029ab717a8412664df716d1fe7f2e824853e7","transactionHash":"0x7a5784667c10ae49797c6fd758acafde1450abc0bfb3d5aaa31abe031b752664","transactionIndex":0,"address":"0xd2a5bC10698FD955D1Fe6cb468a17809A08fd005","data":"0x00000000000000000000000000000000000000000000000000000000003715910000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc4000000000000000000000000d8b934580fce35a11b58c6d73adee468a2833fa8000000000000000000000000000000000000000000000000000000000003eb09000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000002e312f302f307830303030303030303030303030303030303030303030303030303030303030303030303030303030000000000000000000000000000000000000","topics":["0xb0d72a5b25cc833270260cccd76510dddd5e3a0aa8e49018804c93ed3b256838"],"id":"log_51e1089e"}`
	// var logData LogData
	// body := []byte(jsonContent)
	// err := json.Unmarshal(body, &logData)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// h, _ := hex.DecodeString(logData.Data[2:])
	// var tma score.TfuelTransferMetaData
	// fmt.Println(logData.Data[2:])
	// contractAbi, _ := abi.JSON(strings.NewReader(string(scta.MainchainTFuelTokenBankABI)))
	// contractAbi.UnpackIntoInterface(&tma, "TFuelTokenLocked", h)
	// fmt.Println(tma)

	fbk := flag.Int("fbk", 0, "from block")
	tbk := flag.Int("tbk", 0, "to block")
	// stateHashPtr := flag.String("state_hash", "", "hash of state root")
	flag.Parse()
	s := siu.NewInterChainEventCache(nil)
	// fmt.Println(s.EventSelectors)
	queryEventLog(big.NewInt(int64(*fbk)), big.NewInt(int64(*tbk)), common.HexToAddress("0xde9f866B980C3c1197b316d482D50F70b2854C95"), score.IMCEventTypeCrossChainTokenLockTFuel, s)
	// fromBlock := oct(*fbk)
	// toBlock := oct(*tbk)
	// contractAddr :=

	os.Exit(0)
}
