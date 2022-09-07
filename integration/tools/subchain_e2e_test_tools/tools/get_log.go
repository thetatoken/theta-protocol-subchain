package tools

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
	"github.com/thetatoken/thetasubchain/eth/abi"
)

func getSubchainTFuelVoucherMintLogs(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
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
		"name": "TFuelVoucherMinted",
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
		Id      int64     `json:"id"`
		Result  []LogData `json:"result"`
	}

	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TFuelVoucherMinted(string,address,uint256,uint256,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	////fmt.Println(queryStr)

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
	//fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rpcres)
	for _, logData := range rpcres.Result {
		//fmt.Println(idx)
		logData := logData
		type TransferEvt struct {
			Denom                      string
			TargetChainVoucherReceiver common.Address
			MintedAmount               *big.Int
			SourceChainTokenLockNonce  *big.Int
			VoucherMintNonce           *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TFuelVoucherMinted", h)
		if err != nil {
			fmt.Println(err)
		}
		if event.TargetChainVoucherReceiver == receiver {
			return &(event.TargetChainVoucherReceiver)
		}
	}
	return nil
}

func getMainchainTFuelUnlockLogs(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
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
				"name": "targetChainTokenReceiver",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "unlockedAmount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "sourceChainVoucherBurnNonce",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenUnlockNonce",
				"type": "uint256"
			}
		],
		"name": "TFuelTokenUnlocked",
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
		Id      int64     `json:"id"`
		Result  []LogData `json:"result"`
	}

	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TFuelTokenUnlocked(string,address,uint256,uint256,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	////fmt.Println(queryStr)

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
	//fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rpcres)
	for _, logData := range rpcres.Result {
		//fmt.Println(idx)
		logData := logData
		type TransferEvt struct {
			Denom                       string
			TargetChainTokenReceiver    common.Address
			UnlockedAmount              *big.Int
			SourceChainVoucherBurnNonce *big.Int
			TokenUnlockNonce            *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TFuelTokenUnlocked", h)
		if err != nil {
			fmt.Println(err)
		}
		if event.TargetChainTokenReceiver == receiver {
			return &(event.TargetChainTokenReceiver)
		}
	}
	return nil
}

func getSubchainTNT20VoucherMintLogs(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
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
		Id      int64     `json:"id"`
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	////fmt.Println(queryStr)

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
	//fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rpcres)
	for _, logData := range rpcres.Result {
		//fmt.Println(idx)
		logData := logData
		//fmt.Println(logData.Data)
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

func getSubchainTNT721VoucherMintLogs(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
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
		Id      int64     `json:"id"`
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT721VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
	//fmt.Println("query str", queryStr)
	var jsonData = []byte(queryStr)
	//fmt.Println(queryStr)

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
	//fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rpcres)
	for _, logData := range rpcres.Result {
		//fmt.Println(idx)
		logData := logData
		//fmt.Println(logData.Data)
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

// func getMainchainMintlog(fromBlock, toBlock int, contractAddr common.Address) {
// 	const RawABI = `
// 	[
// 		{
// 			"anonymous": false,
// 			"inputs": [
// 				{
// 					"indexed": false,
// 					"internalType": "string",
// 					"name": "denom",
// 					"type": "string"
// 				},
// 				{
// 					"indexed": false,
// 					"internalType": "address",
// 					"name": "targetChainVoucherReceiver",
// 					"type": "address"
// 				},
// 				{
// 					"indexed": false,
// 					"internalType": "address",
// 					"name": "voucherContact",
// 					"type": "address"
// 				},
// 				{
// 					"indexed": false,
// 					"internalType": "uint256",
// 					"name": "mintedAmount",
// 					"type": "uint256"
// 				},
// 				{
// 					"indexed": false,
// 					"internalType": "uint256",
// 					"name": "sourceChainTokenLockNonce",
// 					"type": "uint256"
// 				},
// 				{
// 					"indexed": false,
// 					"internalType": "uint256",
// 					"name": "voucherMintNonce",
// 					"type": "uint256"
// 				}
// 			],
// 			"name": "TNT20VoucherMinted",
// 			"type": "event"
// 		}
// 	]`
// 	url := "http://127.0.0.1:18888/rpc"

// 	type LogData struct {
// 		LogIndex         string   `json:"logIndex"`
// 		TransactionIndex string   `json:"transactionIndex"`
// 		TransactionHash  string   `json:"transactionHash"`
// 		BlockHash        string   `json:"blockHash"`
// 		BlockNumber      string   `json:"blockNumber"`
// 		Address          string   `json:"address"`
// 		Data             string   `json:"data"`
// 		Topics           []string `json:"topics"`
// 		Type             string   `json:"type"`
// 	}

// 	type RPCResult struct {
// 		Jsonrpc string    `json:"jsonrpc"`
// 		Id      int64     `json:"id"`
// 		Result  []LogData `json:"result"`
// 	}

// 	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
// 	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
// 	var jsonData = []byte(queryStr)
// 	//fmt.Println(queryStr)

// 	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		// logger.Fatal(err)
// 	}
// 	request.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		// logger.Fatalf("response error : %v", err)
// 	}
// 	defer response.Body.Close()

// 	body, _ := ioutil.ReadAll(response.Body)

// 	var rpcres RPCResult
// 	//fmt.Println("-----------RPC res----------------")

// 	err = json.Unmarshal(body, &rpcres)
// 	if err != nil {
// 		fmt.Println(err)
// 		// logger.Fatalf("json unmarshal error : %v", err)
// 	}
// 	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//fmt.Println(rpcres)
// 	for _, logData := range rpcres.Result {
// 		//fmt.Println(idx)
// 		logData := logData
// 		//fmt.Println(logData.Data)
// 		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
// 		// if err != nil {
// 		// 	fmt.Println(err)
// 		// }
// 		// fmt.Println(out)
// 		//TNT20VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
// 		type TransferEvt struct {
// 			Denom                      string
// 			TargetChainVoucherReceiver common.Address
// 			VoucherContact             common.Address
// 			MintedAmount               *big.Int
// 			SourceChainTokenLockNonce  *big.Int
// 			VoucherMintNonce           *big.Int
// 		}
// 		var event TransferEvt
// 		txData := logData.Data
// 		h, _ := hex.DecodeString(txData[2:])
// 		err = contractAbi.UnpackIntoInterface(&event, "TNT20VoucherMinted", h)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(event.VoucherContact)
// 		fmt.Println(event.MintedAmount)
// 		fmt.Println(event.TargetChainVoucherReceiver)
// 	}
// }

func getMainchainTNT20VoucherMintLogs(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
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
		Id      int64     `json:"id"`
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	//fmt.Println(queryStr)

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
	//fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rpcres)
	for _, logData := range rpcres.Result {
		//fmt.Println(idx)
		logData := logData
		//fmt.Println(logData.Data)
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

// func getSubchainTNT20Mintlog(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
// 	const RawABI = `
// [
// 	{
// 		"anonymous": false,
// 		"inputs": [
// 			{
// 				"indexed": false,
// 				"internalType": "string",
// 				"name": "denom",
// 				"type": "string"
// 			},
// 			{
// 				"indexed": false,
// 				"internalType": "address",
// 				"name": "targetChainVoucherReceiver",
// 				"type": "address"
// 			},
// 			{
// 				"indexed": false,
// 				"internalType": "address",
// 				"name": "voucherContract",
// 				"type": "address"
// 			},
// 			{
// 				"indexed": false,
// 				"internalType": "uint256",
// 				"name": "mintedAmount",
// 				"type": "uint256"
// 			},
// 			{
// 				"indexed": false,
// 				"internalType": "uint256",
// 				"name": "sourceChainTokenLockNonce",
// 				"type": "uint256"
// 			},
// 			{
// 				"indexed": false,
// 				"internalType": "uint256",
// 				"name": "voucherMintNonce",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "TNT20VoucherMinted",
// 		"type": "event"
// 	}
// ]`
// 	url := "http://127.0.0.1:18888/rpc"

// 	type LogData struct {
// 		LogIndex         string   `json:"logIndex"`
// 		TransactionIndex string   `json:"transactionIndex"`
// 		TransactionHash  string   `json:"transactionHash"`
// 		BlockHash        string   `json:"blockHash"`
// 		BlockNumber      string   `json:"blockNumber"`
// 		Address          string   `json:"address"`
// 		Data             string   `json:"data"`
// 		Topics           []string `json:"topics"`
// 		Type             string   `json:"type"`
// 	}

// 	type RPCResult struct {
// 		Jsonrpc string    `json:"jsonrpc"`
// 		Id      int64     `json:"id"`
// 		Result  []LogData `json:"result"`
// 	}

// 	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
// 	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
// 	var jsonData = []byte(queryStr)
// 	//fmt.Println(queryStr)

// 	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		// logger.Fatal(err)
// 	}
// 	request.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		// logger.Fatalf("response error : %v", err)
// 	}
// 	defer response.Body.Close()

// 	body, _ := ioutil.ReadAll(response.Body)

// 	var rpcres RPCResult
// 	//fmt.Println("-----------RPC res----------------")

// 	err = json.Unmarshal(body, &rpcres)
// 	if err != nil {
// 		fmt.Println(err)
// 		// logger.Fatalf("json unmarshal error : %v", err)
// 	}
// 	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//fmt.Println(rpcres)
// 	for _, logData := range rpcres.Result {
// 		//fmt.Println(idx)
// 		logData := logData
// 		//fmt.Println(logData.Data)
// 		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
// 		// if err != nil {
// 		// 	fmt.Println(err)
// 		// }
// 		// fmt.Println(out)
// 		//TNT20VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
// 		type TransferEvt struct {
// 			Denom                      string
// 			TargetChainVoucherReceiver common.Address
// 			VoucherContract            common.Address
// 			MintedAmount               *big.Int
// 			SourceChainTokenLockNonce  *big.Int
// 			VoucherMintNonce           *big.Int
// 		}
// 		var event TransferEvt
// 		txData := logData.Data
// 		h, _ := hex.DecodeString(txData[2:])
// 		err = contractAbi.UnpackIntoInterface(&event, "TNT20VoucherMinted", h)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		if event.TargetChainVoucherReceiver == receiver {
// 			return &(event.VoucherContract)
// 		}
// 	}
// 	return nil
// }
func getMainchainTNT721VoucherMintLogs(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
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
		Id      int64     `json:"id"`
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT20VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT721VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex())
	//fmt.Println("query str", queryStr)
	var jsonData = []byte(queryStr)
	//fmt.Println(queryStr)

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
	//fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rpcres)
	for _, logData := range rpcres.Result {
		//fmt.Println(idx)
		logData := logData
		//fmt.Println(logData.Data)
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

func getMainchainTNT1155VoucherMintLogs(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
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
		"name": "TNT1155VoucherMinted",
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
		Id      int64     `json:"id"`
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT1155VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT1155VoucherMinted(string,address,address,uint256,uint256,uint256,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	//fmt.Println(queryStr)

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
	//fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rpcres)
	for _, logData := range rpcres.Result {
		//fmt.Println(idx)
		logData := logData
		//fmt.Println(logData.Data)
		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(out)
		//TNT1155VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
		type TransferEvt struct {
			Denom                      string
			TargetChainVoucherReceiver common.Address
			VoucherContract            common.Address
			TokenID                    *big.Int
			MintedAmount               *big.Int
			SourceChainTokenLockNonce  *big.Int
			VoucherMintNonce           *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TNT1155VoucherMinted", h)
		if err != nil {
			fmt.Println(err)
		}
		if event.TargetChainVoucherReceiver == receiver {
			return &(event.VoucherContract)
		}
	}
	return nil
}
func getSubchainTNT1155VoucherMintLogs(fromBlock, toBlock int, contractAddr common.Address, receiver common.Address) *common.Address {
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
		"name": "TNT1155VoucherMinted",
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
		Id      int64     `json:"id"`
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT1155VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), contractAddr.Hex(), crypto.Keccak256Hash([]byte("TNT1155VoucherMinted(string,address,address,uint256,uint256,uint256,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	//fmt.Println(queryStr)

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
	//fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rpcres)
	for _, logData := range rpcres.Result {
		//fmt.Println(idx)
		logData := logData
		//fmt.Println(logData.Data)
		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(out)
		//TNT1155VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
		type TransferEvt struct {
			Denom                      string
			TargetChainVoucherReceiver common.Address
			VoucherContract            common.Address
			TokenID                    *big.Int
			MintedAmount               *big.Int
			SourceChainTokenLockNonce  *big.Int
			VoucherMintNonce           *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "TNT1155VoucherMinted", h)
		if err != nil {
			fmt.Println(err)
		}
		if event.TargetChainVoucherReceiver == receiver {
			return &(event.VoucherContract)
		}
	}
	return nil
}
func getSubchainTNT1155VoucherMintLog(fromBlock, toBlock int) *common.Address {
	const RawABI = `
[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "x",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "y",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "z",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "q",
				"type": "uint256"
			}
		],
		"name": "has",
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
		Id      int64     `json:"id"`
		Result  []LogData `json:"result"`
	}

	//queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0","toBlock":"83c", "address":"0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"}`, crypto.Keccak256Hash([]byte("TNT1155VoucherMinted(string,address,address,uint256,uint256,uint256)")).Hex()
	queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"%v","toBlock":"%v", "address":"%v","topics":["%v"]}],"id":74}`, fmt.Sprintf("%x", fromBlock), fmt.Sprintf("%x", toBlock), "0x47c5e40890bce4a473a49d7501808b9633f29782", crypto.Keccak256Hash([]byte("has(uint256,uint256,address,uint256)")).Hex())
	var jsonData = []byte(queryStr)
	//fmt.Println(queryStr)

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
	//fmt.Println("-----------RPC res----------------")

	err = json.Unmarshal(body, &rpcres)
	if err != nil {
		fmt.Println(err)
		// logger.Fatalf("json unmarshal error : %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rpcres)
	for _, logData := range rpcres.Result {
		//fmt.Println(idx)
		logData := logData
		//fmt.Println(logData.Data)
		//out, err := contractAbi.Unpack("SendToSubchainEvent", logData.Data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(out)
		//TNT1155VoucherMinted(string denom, address targetChainVoucherReceiver, address voucherContact, uint mintedAmount, uint sourceChainTokenLockNonce, uint voucherMintNonce);
		type TransferEvt struct {
			X *big.Int
			Y *big.Int
			Z common.Address
			Q *big.Int
		}
		var event TransferEvt
		txData := logData.Data
		h, _ := hex.DecodeString(txData[2:])
		err = contractAbi.UnpackIntoInterface(&event, "has", h)
		if err != nil {
			fmt.Println(err)
		}

	}
	return nil
}
