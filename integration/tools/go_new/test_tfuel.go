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
	auth.Value = big.NewInt(2000000) //new(big.Int).Mul(dec18, big.NewInt(200000))
	tx, err := tfuelTokenBankInstance.LockTokens(auth, subchainID, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())

}

func subchainTfuelBurn() {
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	authUser := SubchainSelectAccount(client, 1)
	//subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	TfuelTokenBankInstance, err := ct.NewTFuelTokenBank(subchainTfuelTokenBank, client)
	if err != nil {
		log.Fatal(err)
	}

	authUser = SubchainSelectAccount(client, 1)
	authUser.Value = big.NewInt(11111)
	tx, err := TfuelTokenBankInstance.BurnVouchers(authUser, accountList[1].fromAddress)
	//tx, err := testSubchainTNT721TokenBankInstance.GetDenom(nil, subchainTNT721VoucherAddress)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tx)
	fmt.Println(tx.Hash().Hex())
}
