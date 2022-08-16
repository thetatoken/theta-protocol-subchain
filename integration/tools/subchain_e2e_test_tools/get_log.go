package main

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
)

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

func TestGetMintlog(fromBlock, toBlock uint64, contractAddr common.Address, receiver common.Address) *common.Address {
	const RawABI = `
[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "denom",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "targetChainVoucherReceiver",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "voucherContact",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "mintedAmount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "sourceChainTokenLockNonce",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "voucherMintNonce",
				"type": "uint256"
			}
		],
		"name": "TNT20VoucherMinted",
		"type": "event"
	}
]`
	url := "http://127.0.0.1:18888/rpc"

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
		Id      int64     `json:"id"` // 声明对应的json key
		Result  []LogData `json:"result"`
	}
	var topics string
	for _,value :=range(EventSelectors){
		topics=topics+",\""+value+"\""
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":[%v],"topics":[[%v]]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), fmt.Sprintf("\"%v\",\"%v\",\"%v\"", tfuelTokenbankAddress, tnt20TokenBankAddress, tnt721TokenBankAddress),topics[1:])// fmt.Sprintf("[\"%v\",\"%v\",\"%v\"]", crypto.Keccak256Hash([]byte("TNT20TokenLocked(string,address,uint256,address,uint256,string,string,uint8,uint256)")).Hex(), crypto.Keccak256Hash([]byte("TFuelTokenLocked(string,address,uint256,address,uint256,uint256)")).Hex(), crypto.Keccak256Hash([]byte("TNT721TokenLocked(string,address,uint256,address,uint256,string,string,string,uint256)")).Hex()))
	var jsonData = []byte(queryStr)
	fmt.Println(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		// logger.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		// logger.Fatalf("response error : %v", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var rpcres RPCResult
	fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rpcres)
	for idx, logData := range rpcres.Result {
		fmt.Println(idx)
		logData := logData
		fmt.Println(logData.Data)
		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(out)
		//TNT20VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
		type TransferEvt struct {
			Denom                      string
			TargetChainVoucherReceiver common.Address
			VoucherContact             common.Address
			MintedAmount               *big.Int
			SourceChainTokenLockNonce  *big.Int
			VoucherMintNonce           *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TNT20VoucherMinted", h)
		if err != nil {
			fmt.Println(err)
		}
		if event.TargetChainVoucherReceiver == receiver {
			return &(event.VoucherContact)
		}
	}
	return nil
}

func getMintlog(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
	const RawABI = `
[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "denom",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "targetChainVoucherReceiver",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "voucherContact",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "mintedAmount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "sourceChainTokenLockNonce",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "voucherMintNonce",
				"type": "uint256"
			}
		],
		"name": "TNT20VoucherMinted",
		"type": "event"
	}
]`
	url := "http://127.0.0.1:19888/rpc"

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
		Id      int64     `json:"id"` // 声明对应的json key
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	fmt.Println(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		// logger.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		// logger.Fatalf("response error : %v", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var rpcres RPCResult
	fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rpcres)
	for idx, logData := range rpcres.Result {
		fmt.Println(idx)
		logData := logData
		fmt.Println(logData.Data)
		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(out)
		//TNT20VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
		type TransferEvt struct {
			Denom                      string
			TargetChainVoucherReceiver common.Address
			VoucherContact             common.Address
			MintedAmount               *big.Int
			SourceChainTokenLockNonce  *big.Int
			VoucherMintNonce           *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TNT20VoucherMinted", h)
		if err != nil {
			fmt.Println(err)
		}
		if event.TargetChainVoucherReceiver == receiver {
			return &(event.VoucherContact)
		}
	}
	return nil
}

func get721Mintlog(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
	const RawABI = `
[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "denom",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "targetChainVoucherReceiver",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "voucherContract",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "sourceChainTokenLockNonce",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "voucherMintNonce",
				"type": "uint256"
			}
		],
		"name": "TNT721VoucherMinted",
		"type": "event"
	}
]`
	url := "http://127.0.0.1:19888/rpc"

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
		Id      int64     `json:"id"` // 声明对应的json key
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT721VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
	fmt.Println("query str", queryStr)
	var jsonData = []byte(queryStr)
	fmt.Println(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		// logger.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		// logger.Fatalf("response error : %v", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var rpcres RPCResult
	fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rpcres)
	for idx, logData := range rpcres.Result {
		fmt.Println(idx)
		logData := logData
		fmt.Println(logData.Data)
		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(out)
		//TNT20VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
		type TransferEvt struct {
			Denom                      string
			TargetChainVoucherReceiver common.Address
			VoucherContract            common.Address
			TokenID                    *big.Int
			SourceChainTokenLockNonce  *big.Int
			VoucherMintNonce           *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TNT721VoucherMinted", h)
		if err != nil {
			fmt.Println(err)
		}
		if event.TargetChainVoucherReceiver == receiver {
			return &(event.VoucherContract)
		}
	}
	return nil
}

func getMainchainMintlog(fromBlock, toBlock int, contractAddr common.Address) {
	const RawABI = `
	[
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": false,
					"internalType": "string",
					"name": "denom",
					"type": "string"
				},
				{
					"indexed": false,
					"internalType": "address",
					"name": "targetChainVoucherReceiver",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "address",
					"name": "voucherContact",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "mintedAmount",
					"type": "uint256"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "sourceChainTokenLockNonce",
					"type": "uint256"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "voucherMintNonce",
					"type": "uint256"
				}
			],
			"name": "TNT20VoucherMinted",
			"type": "event"
		}
	]`
	url := "http://127.0.0.1:18888/rpc"

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
		Id      int64     `json:"id"` // 声明对应的json key
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	fmt.Println(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		// logger.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		// logger.Fatalf("response error : %v", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var rpcres RPCResult
	fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rpcres)
	for idx, logData := range rpcres.Result {
		fmt.Println(idx)
		logData := logData
		fmt.Println(logData.Data)
		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(out)
		//TNT20VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
		type TransferEvt struct {
			Denom                      string
			TargetChainVoucherReceiver common.Address
			VoucherContact             common.Address
			MintedAmount               *big.Int
			SourceChainTokenLockNonce  *big.Int
			VoucherMintNonce           *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TNT20VoucherMinted", h)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(event.VoucherContact)
		fmt.Println(event.MintedAmount)
		fmt.Println(event.TargetChainVoucherReceiver)
	}
}
func getMainchainTNT20Mintlog(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
	const RawABI = `
[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "denom",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "targetChainVoucherReceiver",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "voucherContract",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "mintedAmount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "sourceChainTokenLockNonce",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "voucherMintNonce",
				"type": "uint256"
			}
		],
		"name": "TNT20VoucherMinted",
		"type": "event"
	}
]`
	url := "http://127.0.0.1:18888/rpc"

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
		Id      int64     `json:"id"` // 声明对应的json key
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	fmt.Println(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		// logger.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		// logger.Fatalf("response error : %v", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var rpcres RPCResult
	fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rpcres)
	for idx, logData := range rpcres.Result {
		fmt.Println(idx)
		logData := logData
		fmt.Println(logData.Data)
		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(out)
		//TNT20VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
		type TransferEvt struct {
			Denom                      string
			TargetChainVoucherReceiver common.Address
			VoucherContract            common.Address
			MintedAmount               *big.Int
			SourceChainTokenLockNonce  *big.Int
			VoucherMintNonce           *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TNT20VoucherMinted", h)
		if err != nil {
			fmt.Println(err)
		}
		if event.TargetChainVoucherReceiver == receiver {
			return &(event.VoucherContract)
		}
	}
	return nil
}
func getSubchainTNT20Mintlog(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
	const RawABI = `
[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "denom",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "targetChainVoucherReceiver",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "voucherContract",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "mintedAmount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "sourceChainTokenLockNonce",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "voucherMintNonce",
				"type": "uint256"
			}
		],
		"name": "TNT20VoucherMinted",
		"type": "event"
	}
]`
	url := "http://127.0.0.1:18888/rpc"

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
		Id      int64     `json:"id"` // 声明对应的json key
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	fmt.Println(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		// logger.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		// logger.Fatalf("response error : %v", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var rpcres RPCResult
	fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rpcres)
	for idx, logData := range rpcres.Result {
		fmt.Println(idx)
		logData := logData
		fmt.Println(logData.Data)
		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(out)
		//TNT20VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
		type TransferEvt struct {
			Denom                      string
			TargetChainVoucherReceiver common.Address
			VoucherContract            common.Address
			MintedAmount               *big.Int
			SourceChainTokenLockNonce  *big.Int
			VoucherMintNonce           *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TNT20VoucherMinted", h)
		if err != nil {
			fmt.Println(err)
		}
		if event.TargetChainVoucherReceiver == receiver {
			return &(event.VoucherContract)
		}
	}
	return nil
}
func getMainchainTNT721Mintlog(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
	const RawABI = `
[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "denom",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "targetChainVoucherReceiver",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "voucherContract",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "sourceChainTokenLockNonce",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "voucherMintNonce",
				"type": "uint256"
			}
		],
		"name": "TNT721VoucherMinted",
		"type": "event"
	}
]`
	url := "http://127.0.0.1:18888/rpc"

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
		Id      int64     `json:"id"` // 声明对应的json key
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT721VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
	fmt.Println("query str", queryStr)
	var jsonData = []byte(queryStr)
	fmt.Println(queryStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		// logger.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		// logger.Fatalf("response error : %v", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var rpcres RPCResult
	fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rpcres)
	for idx, logData := range rpcres.Result {
		fmt.Println(idx)
		logData := logData
		fmt.Println(logData.Data)
		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(out)
		//TNT20VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
		type TransferEvt struct {
			Denom                      string
			TargetChainVoucherReceiver common.Address
			VoucherContract            common.Address
			TokenID                    *big.Int
			SourceChainTokenLockNonce  *big.Int
			VoucherMintNonce           *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TNT721VoucherMinted", h)
		if err != nil {
			fmt.Println(err)
		}
		if event.TargetChainVoucherReceiver == receiver {
			return &(event.VoucherContract)
		}
	}
	return nil
}
