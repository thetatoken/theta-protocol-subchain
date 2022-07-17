package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getlog() {

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

	queryStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"1b68","toBlock":"1b70", "address":"0xc7253857256391E518c4aF60aDa5Eb5972Dd6Dbc","topics":["0xe5d8852bc02bf44f2a49b2d7722fa497ff83b689a28de1253304d2bc43d7b1cb"]}],"id":74}`
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
	fmt.Println(rpcres)

}
