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

var dec18, _ = new(big.Int).SetString("1000000000000000000", 10)
var crossChainFee = new(big.Int).Mul(big.NewInt(10), dec18)

func MainchainTFuelLock(lockAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	receiver := accountList[1].fromAddress

	tfuelTokenBankInstance, _ := ct.NewTFuelTokenBank(mainchainTFuelTokenBankAddress, mainchainClient)
	sender := mainchainSelectAccount(mainchainClient, 3)

	mainchainHeight, _ := mainchainClient.BlockNumber(context.Background())
	senderMainchainTFuelBalance, _ := mainchainClient.BalanceAt(context.Background(), sender.From, big.NewInt(int64(mainchainHeight)))

	subchainHeight, _ := subchainClient.BlockNumber(context.Background())
	receiverSubchainTFuelBalance, _ := subchainClient.BalanceAt(context.Background(), receiver, big.NewInt(int64(subchainHeight)))

	fmt.Printf("Mainchain sender : %v, TFuel balance on Mainchain       : %v\n", sender.From, senderMainchainTFuelBalance)
	fmt.Printf("Subchain receiver: %v, TFuel voucher balance on Subchain: %v\n\n", receiver, receiverSubchainTFuelBalance)
	printTFuelTokenBankLockedAmount(mainchainTFuelTokenBankAddress, subchainID, mainchainClient)

	sender.Value = big.NewInt(0).Add(lockAmount, crossChainFee)
	tx, err := tfuelTokenBankInstance.LockTokens(sender, subchainID, receiver)
	sender.Value = big.NewInt(0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TFuel Token Lock tx hash (Mainchain): %v\n", tx.Hash().Hex())
	fmt.Printf("Transfering %v TFuelWei from the Mainchain to Subchain %v...\n\n", lockAmount, subchainID)

	fmt.Printf("Start transfer, timestamp: %v\n", time.Now())
	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	for {
		time.Sleep(1 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getSubchainTFuelVoucherMintLogs(int(fromHeight), int(toHeight), subchainTFuelTokenBankAddress, receiver)
		if result != nil {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Subchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("End transfer, timestamp  : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	printTFuelTokenBankLockedAmount(mainchainTFuelTokenBankAddress, subchainID, mainchainClient)

	mainchainHeight, _ = mainchainClient.BlockNumber(context.Background())
	senderMainchainTFuelBalance, _ = mainchainClient.BalanceAt(context.Background(), sender.From, big.NewInt(int64(mainchainHeight)))

	subchainHeight, _ = subchainClient.BlockNumber(context.Background())
	receiverSubchainTFuelBalance, _ = subchainClient.BalanceAt(context.Background(), receiver, big.NewInt(int64(subchainHeight)))

	fmt.Printf("Mainchain sender : %v, TFuel balance on Mainchain       : %v\n", sender.From, senderMainchainTFuelBalance)
	fmt.Printf("Subchain receiver: %v, TFuel voucher balance on Subchain: %v\n\n", receiver, receiverSubchainTFuelBalance)
}

func SubchainTFuelBurn(burnAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	TFuelTokenBankInstance, err := ct.NewTFuelTokenBank(subchainTFuelTokenBankAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}

	sender := subchainSelectAccount(subchainClient, 1)
	receiver := accountList[2].fromAddress

	subchainHeight, _ := subchainClient.BlockNumber(context.Background())
	senderSubchainTFuelBalance, _ := subchainClient.BalanceAt(context.Background(), sender.From, big.NewInt(int64(subchainHeight)))

	mainchainHeight, _ := mainchainClient.BlockNumber(context.Background())
	receiverMainchainTFuelBalance, _ := mainchainClient.BalanceAt(context.Background(), receiver, big.NewInt(int64(mainchainHeight)))

	fmt.Printf("Subchain sender   : %v, TFuel voucher balance on Subchain : %v\n", sender.From, senderSubchainTFuelBalance)
	fmt.Printf("Mainchain receiver: %v, TFuel balance on Mainchain        : %v\n\n", receiver, receiverMainchainTFuelBalance)
	printTFuelTokenBankLockedAmount(mainchainTFuelTokenBankAddress, subchainID, mainchainClient)

	sender.Value = big.NewInt(0).Add(burnAmount, crossChainFee)
	tx, err := TFuelTokenBankInstance.BurnVouchers(sender, receiver)
	sender.Value = big.NewInt(0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TFuel Voucher Burn tx hash (Subchain): %v\n", tx.Hash().Hex())
	fmt.Printf("Burn %v TFuelWei on Subchain %v to recover authentic TFuel on the Mainchain...\n\n", burnAmount, subchainID)

	fmt.Printf("Start transfer, timestamp: %v\n", time.Now())
	fromHeight, _ := mainchainClient.BlockNumber(context.Background())
	for {
		time.Sleep(1 * time.Second)
		toHeight, _ := mainchainClient.BlockNumber(context.Background())
		result := getMainchainTFuelUnlockLogs(int(fromHeight), int(toHeight), mainchainTFuelTokenBankAddress, receiver)
		if result != nil {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Mainchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("End transfer, timestamp  : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	printTFuelTokenBankLockedAmount(mainchainTFuelTokenBankAddress, subchainID, mainchainClient)

	subchainHeight, _ = subchainClient.BlockNumber(context.Background())
	senderSubchainTFuelBalance, _ = subchainClient.BalanceAt(context.Background(), sender.From, big.NewInt(int64(subchainHeight)))

	mainchainHeight, _ = mainchainClient.BlockNumber(context.Background())
	receiverMainchainTFuelBalance, _ = mainchainClient.BalanceAt(context.Background(), receiver, big.NewInt(int64(mainchainHeight)))

	fmt.Printf("Subchain sender   : %v, TFuel voucher balance on Subchain : %v\n", sender.From, senderSubchainTFuelBalance)
	fmt.Printf("Mainchain receiver: %v, TFuel balance on Mainchain        : %v\n\n", receiver, receiverMainchainTFuelBalance)
}

func printTFuelTokenBankLockedAmount(tokenBankAddr common.Address, chainID *big.Int, ethclient *ethclient.Client) {
	tfuelTokenBank, _ := ct.NewTFuelTokenBank(tokenBankAddr, ethclient)
	tokenBankLockedAmount, _ := tfuelTokenBank.TotalLockedAmounts(nil, chainID)
	fmt.Printf("TFuelTokenBank locked amount: %v\n\n", tokenBankLockedAmount)
}
