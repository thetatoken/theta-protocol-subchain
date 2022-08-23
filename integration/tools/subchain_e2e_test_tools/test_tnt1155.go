package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

func MainchainTNT1155Lock() {
	user := accountList[1].fromAddress
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	instanceTNT1155VoucherContract, err := ct.NewTNT1155VoucherContract(mainchainTNT1155VoucherContractAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	instanceTNT1155TokenBank, err := ct.NewTNT1155TokenBank(mainchainTNT1155TokenBankAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := mainchainSelectAccount(mainchainClient, 0)
	tx, err := instanceTNT1155VoucherContract.Mint(authAccount0, user, big.NewInt(1), big.NewInt(1), []byte(""))
	if err != nil {
		log.Fatal(err)
	}
	authAccount1 := mainchainSelectAccount(mainchainClient, 1)
	tx, err = instanceTNT1155VoucherContract.SetApprovalForAll(authAccount1, mainchainTNT1155TokenBankAddress, true)
	if err != nil {
		log.Fatal(err)
	}
	authAccount1 = mainchainSelectAccount(mainchainClient, 1)
	tx, err = instanceTNT1155TokenBank.LockTokens(authAccount1, subchainID, mainchainTNT1155VoucherContractAddress, user, big.NewInt(1), big.NewInt(1), []byte(""))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
}
