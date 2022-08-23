package tools

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

func MainchainTFuelLock(lockAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	user := accountList[1].fromAddress

	tfuelTokenBankInstance, _ := ct.NewTFuelTokenBank(tfuelTokenbankAddress, mainchainClient)
	auth := mainchainSelectAccount(mainchainClient, 3)
	auth.Value = lockAmount // big.NewInt(2000000) //new(big.Int).Mul(dec18, big.NewInt(200000))

	mainchainHeight, _ := mainchainClient.BlockNumber(context.Background())
	senderMainchainTFuelBalance, _ := mainchainClient.BalanceAt(context.Background(), auth.From, big.NewInt(int64(mainchainHeight)))

	subchainHeight, _ := subchainClient.BlockNumber(context.Background())
	receiverSubchainTFuelBalance, _ := subchainClient.BalanceAt(context.Background(), user, big.NewInt(int64(subchainHeight)))

	fmt.Printf("Mainchain sender : %v, TFuel balance on Mainchain       : %v\n", auth.From, senderMainchainTFuelBalance)
	fmt.Printf("Subchain receiver: %v, TFuel voucher balance on Subchain: %v \n", user, receiverSubchainTFuelBalance)
	tx, err := tfuelTokenBankInstance.LockTokens(auth, subchainID, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TFuel Token Lock tx hash (Mainchain): %v\n", tx.Hash().Hex())
	fmt.Printf("Transfering %v TFuelWei from the Mainchain to Subchain %v...\n\n", lockAmount, subchainID)

	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	for {
		time.Sleep(6 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getSubchainTFuelVoucherMintlog(int(fromHeight), int(toHeight), subchainTFuelTokenBankAddress, user)
		if result != nil {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Subchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("Cross-chain transfer completed.\n\n")

	mainchainHeight, _ = mainchainClient.BlockNumber(context.Background())
	senderMainchainTFuelBalance, _ = mainchainClient.BalanceAt(context.Background(), auth.From, big.NewInt(int64(mainchainHeight)))

	subchainHeight, _ = subchainClient.BlockNumber(context.Background())
	receiverSubchainTFuelBalance, _ = subchainClient.BalanceAt(context.Background(), user, big.NewInt(int64(subchainHeight)))

	fmt.Printf("Mainchain sender : %v, TFuel balance on Mainchain       : %v\n", auth.From, senderMainchainTFuelBalance)
	fmt.Printf("Subchain receiver: %v, TFuel voucher balance on Subchain: %v \n", user, receiverSubchainTFuelBalance)
}

func SubchainTFuelBurn(burnAmount *big.Int) {
	// mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	authUser := subchainSelectAccount(subchainClient, 1)
	//subchainTNT721VoucherAddress := common.HexToAddress("0xd5125d7bB9c4Fb222C522c4b1922cabC631E52D7")
	TFuelTokenBankInstance, err := ct.NewTFuelTokenBank(subchainTFuelTokenBankAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}

	authUser = subchainSelectAccount(subchainClient, 1)
	authUser.Value = burnAmount
	tx, err := TFuelTokenBankInstance.BurnVouchers(authUser, accountList[1].fromAddress)
	//tx, err := testSubchainTNT721TokenBankInstance.GetDenom(nil, subchainTNT721VoucherAddress)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tx)
	fmt.Println("burn tx hash is", tx.Hash().Hex())
}
