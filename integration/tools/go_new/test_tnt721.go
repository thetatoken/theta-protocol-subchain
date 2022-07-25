package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/integration/tools/go_new/accessors"
)

func tnt721Lock() {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainID := big.NewInt(360777)
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	user := accountList[1].fromAddress

	instanceTNT721VoucherContract, err := ct.NewTNT721VoucherContract(TNT721VoucherContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	instanceTNT721TokenBank, err := ct.NewTNT721TokenBank(TNT721TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := SelectAccount(client, 0)
	tx, err := instanceTNT721VoucherContract.Mint(authAccount0, user, big.NewInt(666), "721")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	authUser := SelectAccount(client, 1)
	tx, err = instanceTNT721VoucherContract.Approve(authUser, TNT721TokenBankAddress, big.NewInt(666))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	authUser = SelectAccount(client, 1)
	LockTx, err := instanceTNT721TokenBank.LockTokens(authUser, subchainID, TNT721VoucherContractAddress, user, big.NewInt(666))
	if err != nil {
		log.Fatal(err)
	}
	receipt, err := client.TransactionReceipt(context.Background(), LockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("lock error")
		return
	}
	// subchainClient, _ := ethclient.Dial("http://localhost:19888/rpc")
	// fromHeight, _ := subchainClient.BlockNumber(context.Background())
	// receipt, err := client.TransactionReceipt(context.Background(), LockTx.Hash())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if receipt.Status != 1 {
	// 	fmt.Println("lock error")
	// }
	fmt.Println(LockTx.Hash().Hex())
	subchainClient, _ := ethclient.Dial("http://localhost:19888/rpc")
	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	fmt.Println(LockTx.Hash())
	var subchainVoucherAddress common.Address
	for {
		time.Sleep(2 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := get721Mintlog(int(fromHeight), int(toHeight), SubchainTNT721TokenBankAddress, user)
		if result != nil {
			subchainVoucherAddress = *result
			break
		}
	}
	fmt.Println(subchainVoucherAddress)
}
func subchainTNT721Burn() {
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	subchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(subchainTNT721VoucherAddress, client)
	subchainTNT721TokenBankinstance, err := ct.NewTNT721TokenBank(SubchainTNT721TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	authUser := SubchainSelectAccount(client, 1)
	tx, err := subchainTNT721VoucherInstance.Approve(authUser, SubchainTNT721TokenBankAddress, big.NewInt(666))
	if err != nil {
		log.Fatal(err)
	}
	authUser = SubchainSelectAccount(client, 1)
	tx, err = subchainTNT721TokenBankinstance.BurnVouchers(authUser, subchainTNT721VoucherAddress, accountList[6].fromAddress, big.NewInt(666))
	if err != nil {
		log.Fatal(err)
	}
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}
	fmt.Println(tx.Hash().Hex())
}
