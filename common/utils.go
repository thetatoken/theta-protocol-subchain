package common

import (
	scom "github.com/thetatoken/thetasubchain/common"
	"github.com/thetatoken/thetasubchain/witness"
	"math/big"
)

const NumMainchainBlocksPerDynasty int64 = 100

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

func CalculateDynasty(mainchainHeight *big.Int) *big.Int {
	return new(big.Int).Div(mainchainHeight, big.NewInt(NumMainchainBlocksPerDynasty))
}

func RpcEventLogQuery(fromBlock int, toBlock int, contractAddr common.Address, witnessXTransferCache *map[*big.Int]*witness.CrossChainTransferEvent) []witness.CrossChainTransferEvent {
	url := "http://127.0.0.1:18888/rpc"
	queryStr := fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"eth_getLogs",
		"params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["0xfcaf9544852e1f0902dbcf8a118d2ae8235ed282692bbb5dc2aff212e4888a41"]}],
		"id":74
	}`, fromBlock, toBlock, mw.RegisterContractAddr.Hex())
	var jsonData = []byte(queryStr)

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	var rpcres RPCResult
	err := json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rpcres.Result[0].Data)

	contractAbi, err := abi.JSON(strings.NewReader(string(ct.SubchainRegisterABI)))
	if err != nil {
		fmt.Println(err)
	}

	crossChainTransferEventArr := make([]witness.CrossChainTransferEvent, len(rpcres.Result))

	for _, logData := range rpcres.Result {
		logData := logData
		var event TransferEvent
		h, _ := hex.DecodeString(logData.Data[2:])
		err := contractAbi.UnpackIntoInterface(&event, "SendToSubchainEvent", h)
		if _, ok := witnessXTransferCache[event.Nonce], ok{
			continue
		}
		resEventArr = append(resEventArr, event)
		if err != nil {
			fmt.Println(err)
		}
		BlockNumberDec, _ := strconv.ParseUint(logData.BlockNumber[2:], 16, 32)
		crossChainTransferEvent := &witness.CrossChainTransferEvent{
			Sender:      common.HexToAddress(logData.Topics[1]),
			Denom:       event.Denom,
			Amount:      event.Amount,
			EventNonce:  event.Nonce,
			BlockNumber: big.NewInt(BlockNumberDec),
		}
		crossChainTransferEventArr = appen(crossChainTransferEventArr, crossChainTransferEvent)
	}

	return crossChainTransferEventArr
}
