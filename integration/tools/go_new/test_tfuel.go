package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/integration/tools/go_new/accessors"
)

func mainchainTfuelLock() {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainID := big.NewInt(360777)
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	user := accountList[1].fromAddress

	tfuelTokenBankInstance, _ := ct.NewTFuelTokenBank(TfuelTokenbankAddress, client)
	auth := SelectAccount(client, 3)
	auth.Value = new(big.Int).Mul(dec18, big.NewInt(200000))
	tx, err := tfuelTokenBankInstance.LockTokens(auth, subchainID, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())

}
