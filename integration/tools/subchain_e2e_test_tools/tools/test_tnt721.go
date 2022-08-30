package tools

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

func MainchainTNT721Lock(tokenID *big.Int) {
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
	fmt.Println("Lock tx hash is", tx.Hash().Hex())
	authUser := mainchainSelectAccount(client, 1)
	tx, err = instanceTNT721VoucherContract.Approve(authUser, tnt721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tx.Hash().Hex())
	authUser = mainchainSelectAccount(client, 1)
	//
	firstOnwer, _ := instanceTNT721VoucherContract.OwnerOf(nil, tokenID)
	fmt.Println("The mainchain owner of ", tokenID, "is", firstOnwer)
	//
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
	//
	SecondOnwer, _ := instanceTNT721VoucherContract.OwnerOf(nil, tokenID)
	fmt.Println("The mainchain owner of ", tokenID, "is", SecondOnwer)
	fmt.Println("---------Detecting---------")
	//
	// subchainClient, _ := ethclient.Dial("http://localhost:19888/rpc")
	// fromHeight, _ := subchainClient.BlockNumber(context.Background())
	// receipt, err := client.TransactionReceipt(context.Background(), LockTx.Hash())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if receipt.Status != 1 {
	// 	fmt.Println("lock error")
	// }
	fmt.Println("lock tx hash is ", LockTx.Hash().Hex())
	subchainClient, _ := ethclient.Dial("http://localhost:19888/rpc")
	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	var subchainVoucherAddress common.Address
	for {
		time.Sleep(2 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getSubchainTNT721VoucherMintLogs(int(fromHeight), int(toHeight), subchainTNT721TokenBankAddress, user)
		if result != nil {
			subchainVoucherAddress = *result
			break
		}
	}
	fmt.Println("Subchain TNT721 Voucher contract address is ", subchainVoucherAddress)
	instanceSubchainVoucherAddress, _ := ct.NewTNT721VoucherContract(subchainVoucherAddress, subchainClient)
	//
	thirdOnwer, _ := instanceSubchainVoucherAddress.OwnerOf(nil, tokenID)
	fmt.Println("The subchain owner of ", tokenID, "is", thirdOnwer)
	//
}
func SubchainTNT721Burn(tokenID *big.Int) {
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	subchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(subchainTNT721VoucherAddress, client)
	subchainTNT721TokenBankinstance, err := ct.NewTNT721TokenBank(subchainTNT721TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	authUser := subchainSelectAccount(client, 1)
	tx, err := subchainTNT721VoucherInstance.Approve(authUser, subchainTNT721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	//
	firstOnwer, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("The subchain owner of ", tokenID, "is", firstOnwer)
	//
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
	fmt.Println("burn tx hash is", tx.Hash().Hex())
	//
	secondOnwer, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("The subchain owner of ", tokenID, "is", secondOnwer)
	fmt.Println("---------Detecting---------")
	//
	mainchainClient, _ := ethclient.Dial("http://localhost:18888/rpc")
	mainchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(tnt721VoucherContractAddress, mainchainClient)
	result, _ := mainchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("the mainchain owner of ", tokenID, "is", result)
	for {
		time.Sleep(2 * time.Second)
		new_result, _ := mainchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
		if new_result != result {
			result = new_result
			break
		}
	}
	fmt.Println("the mainchain owner of ", tokenID, "is", result)
}
func SubchainTNT721Lock(tokenID *big.Int) {
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	subchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(subchainTNT721VoucherAddress, client)
	subchainTNT721TokenBankinstance, err := ct.NewTNT721TokenBank(subchainTNT721TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	auth := subchainSelectAccount(client, 1)
	_, err = subchainTNT721VoucherInstance.Approve(auth, subchainTNT721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	//
	firstOnwer, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("The subchain owner of ", tokenID, "is", firstOnwer)
	//
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
	//
	secondOnwer, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("The subchain owner of ", tokenID, "is", secondOnwer)
	fmt.Println("---------Detecting---------")
	//
	mainchainClient, _ := ethclient.Dial("http://localhost:18888/rpc")

	fromHeight, _ := mainchainClient.BlockNumber(context.Background())
	var mainchainVoucherAddress common.Address
	for {
		time.Sleep(2 * time.Second)
		toHeight, _ := mainchainClient.BlockNumber(context.Background())
		result := getMainchainTNT721VoucherMintLogs(int(fromHeight), int(toHeight), tnt721TokenBankAddress, accountList[6].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
	}
	fmt.Println("Mainchain TNT721 Voucher Contract Address is", mainchainVoucherAddress)
	mainchainTNT721VoucherInstance, err := ct.NewTNT721VoucherContract(mainchainVoucherAddress, mainchainClient)
	result, _ := mainchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("The mainchain token owner of ", tokenID, "is", result)
}
func MainchainTNT721Burn(tokenID *big.Int) {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	mainchainTNT721VoucherAddr := common.HexToAddress("0x20B2897c4f95df71a5a8A62Ae812f5843AA92E25")
	mainchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(mainchainTNT721VoucherAddr, client)
	auth := mainchainSelectAccount(client, 6)
	_, err = mainchainTNT721VoucherInstance.Approve(auth, tnt721TokenBankAddress, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	mainchainTNT721TokenBankInstance, _ := ct.NewTNT721TokenBank(tnt721TokenBankAddress, client)
	//
	firstOnwer, _ := mainchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("The mainchain owner of ", tokenID, "is", firstOnwer)
	//
	auth = mainchainSelectAccount(client, 6)
	tx, err := mainchainTNT721TokenBankInstance.BurnVouchers(auth, mainchainTNT721VoucherAddr, accountList[1].fromAddress, tokenID)
	fmt.Println("burn tx hash is", tx.Hash().Hex())
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}
	//
	secondOnwer, _ := mainchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("The mainchain owner of ", tokenID, "is", secondOnwer)
	fmt.Println("---------Detecting---------")
	//
	subchainClient, _ := ethclient.Dial("http://localhost:19888/rpc")
	subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	subchainTNT721VoucherInstance, _ := ct.NewTNT721VoucherContract(subchainTNT721VoucherAddress, subchainClient)
	result, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
	fmt.Println("The subchain owner of ", tokenID, " is", result)
	for {
		time.Sleep(2 * time.Second)
		new_result, _ := subchainTNT721VoucherInstance.OwnerOf(nil, tokenID)
		if new_result != result {
			result = new_result
			break
		}
	}
	fmt.Println("The subchain owner of ", tokenID, " is", result)
}

func TNT721Query(chainID int64, contractAddress string,tokenID *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	
	tnt721ContractAddress := common.HexToAddress(contractAddress) // FIXME: should instantiate a mock TNT721 instead of using the Voucher contract (which causes confusion)
	var instaceTNT721Contract *ct.TNT721VoucherContract
	if chainID == 366 {
		fmt.Printf("Preparing for TNT721 mainchain query...\n")
		instaceTNT721Contract, err = ct.NewTNT721VoucherContract(tnt721ContractAddress, mainchainClient)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Preparing for TNT721 subchain query...\n")
		instaceTNT721Contract, err = ct.NewTNT721VoucherContract(tnt721ContractAddress, subchainClient)
		if err != nil {
			log.Fatal(err)
		}
	}
	owner ,_:=instaceTNT721Contract.OwnerOf(nil,tokenID)
	fmt.Println("TokenID",tokenID," TNT721 owner in ",contractAddress," is ",owner)

}