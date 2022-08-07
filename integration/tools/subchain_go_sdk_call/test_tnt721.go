package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

func mainchainTNT721Lock(tokenID *big.Int) {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	user := accountList[1].fromAddress

	instanceTNT721VoucherContract, err := ct.NewTNT721VoucherContract(tnt721VoucherContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	instanceTNT721TokenBank, err := ct.NewTNT721TokenBank(tnt721TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := mainchainSelectAccount(client, 0)
	tx, err := instanceTNT721VoucherContract.Mint(authAccount0, user, tokenID, tokenID.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	authUser := mainchainSelectAccount(client, 1)
	tx, err = instanceTNT721VoucherContract.Approve(authUser, tnt721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	authUser = mainchainSelectAccount(client, 1)
	LockTx, err := instanceTNT721TokenBank.LockTokens(authUser, subchainID, tnt721VoucherContractAddress, user, tokenID)
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
		result := get721Mintlog(int(fromHeight), int(toHeight), Subchaintnt721TokenBankAddress, user)
		if result != nil {
			subchainVoucherAddress = *result
			break
		}
	}
	fmt.Println(subchainVoucherAddress)
}
func subchainTNT721Burn(tokenID *big.Int) {
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	subchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(subchainTNT721VoucherAddress, client)
	subchainTNT721TokenBankinstance, err := ct.NewTNT721TokenBank(Subchaintnt721TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	authUser := subchainSelectAccount(client, 1)
	tx, err := subchainTNT721VoucherInstance.Approve(authUser, Subchaintnt721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	authUser = subchainSelectAccount(client, 1)
	tx, err = subchainTNT721TokenBankinstance.BurnVouchers(authUser, subchainTNT721VoucherAddress, accountList[6].fromAddress, tokenID)
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
	mainchainClient, _ := ethclient.Dial("http://localhost:18888/rpc")
	mainchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(tnt721VoucherContractAddress, mainchainClient)
	result, _ := mainchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("owner is", result)
	for {
		time.Sleep(2 * time.Second)
		new_result, _ := mainchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
		if new_result != result {
			result = new_result
			break
		}
	}
	fmt.Println("owner is", result)
}
func subchainTNT721Lock(tokenID *big.Int) {
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	subchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(subchainTNT721VoucherAddress, client)
	subchainTNT721TokenBankinstance, err := ct.NewTNT721TokenBank(Subchaintnt721TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	auth := subchainSelectAccount(client, 1)
	_, err = subchainTNT721VoucherInstance.Approve(auth, Subchaintnt721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	auth = subchainSelectAccount(client, 1)
	tx, err := subchainTNT721TokenBankinstance.LockTokens(auth, big.NewInt(366), subchainTNT721VoucherAddress, accountList[6].fromAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("lock tx hash is", tx.Hash().Hex())
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}
	mainchainClient, _ := ethclient.Dial("http://localhost:18888/rpc")
	mainchainTNT721VoucherInstance, err := ct.NewTNT721VoucherContract(tnt721VoucherContractAddress, mainchainClient)
	result, _ := mainchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("token owner is", result)
	fromHeight, _ := mainchainClient.BlockNumber(context.Background())
	var mainchainVoucherAddress common.Address
	for {
		time.Sleep(2 * time.Second)
		toHeight, _ := mainchainClient.BlockNumber(context.Background())
		result := getMainchainTNT721Mintlog(int(fromHeight), int(toHeight), tnt721TokenBankAddress, accountList[6].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
	}
	fmt.Println(mainchainVoucherAddress)
}
func mainchainTNT721Burn(tokenID *big.Int) {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	mainchainTNT721VoucherAddr := common.HexToAddress("0xaE65cd74C7aF2C5c1009E2Caa6Fa30e4a832a687")
	mainchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(mainchainTNT721VoucherAddr, client)
	auth := mainchainSelectAccount(client, 6)
	_, err = mainchainTNT721VoucherInstance.Approve(auth, tnt721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	mainchainTNT721TokenBankInstance, _ := ct.NewTNT721TokenBank(tnt721TokenBankAddress, client)
	auth = mainchainSelectAccount(client, 6)
	tx, err := mainchainTNT721TokenBankInstance.BurnVouchers(auth, mainchainTNT721VoucherAddr, accountList[7].fromAddress, tokenID)
	fmt.Println("burn tx hash is", tx.Hash().Hex())
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}
	subchainClient, _ := ethclient.Dial("http://localhost:19888/rpc")
	subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	subchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(subchainTNT721VoucherAddress, subchainClient)
	result, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("owner is", result)
	for {
		time.Sleep(2 * time.Second)
		new_result, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
		if new_result != result {
			result = new_result
			break
		}
	}
	fmt.Println("owner is", result)
}
