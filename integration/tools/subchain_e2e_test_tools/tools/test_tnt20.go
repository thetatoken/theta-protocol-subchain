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

func MainchainTNT20Lock(lockAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT20 cross-chain transfer...\n")

	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	sender := mainchainSelectAccount(mainchainClient, 1)
	receiver := accountList[1].fromAddress

	tnt20ContractAddress := tnt20VoucherContractAddress // FIXME: should instantiate a mock TNT20 instead of using the Voucher contract
	instaceTNT20Contract, err := ct.NewTNT20VoucherContract(tnt20ContractAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT20VoucherAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	instaceSubchainTNT20VoucherContract, err := ct.NewTNT20VoucherContract(subchainTNT20VoucherAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}
	instanceTNT20TokenBank, err := ct.NewTNT20TokenBank(tnt20TokenBankAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	mintAmount := big.NewInt(1).Mul(lockAmount, big.NewInt(10))
	fmt.Printf("Mainchain TNT20 contract address: %v\n", tnt20ContractAddress)
	fmt.Printf("Minting %v TNT20 tokens (Wei) on the Mainchain to address %v\n", mintAmount, sender.From)
	_, err = instaceTNT20Contract.Mint(mainchainSelectAccount(mainchainClient, 0), sender.From, mintAmount)
	if err != nil {
		log.Fatal(err)
	}
	_, err = instaceTNT20Contract.Approve(sender, tnt20TokenBankAddress, lockAmount)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(12 * time.Second)

	senderTNT20Balance, _ := instaceTNT20Contract.BalanceOf(nil, sender.From)
	receiverTNT20VoucherBalance, _ := instaceSubchainTNT20VoucherContract.BalanceOf(nil, receiver)
	fmt.Printf("Mainchain sender : %v, TNT20 balance on Mainchain       : %v\n", sender.From, senderTNT20Balance)
	fmt.Printf("Subchain receiver: %v, TNT20 voucher balance on Subchain: %v\n\n", receiver, receiverTNT20VoucherBalance)

	sender = mainchainSelectAccount(mainchainClient, 1)
	lockTx, err := instanceTNT20TokenBank.LockTokens(sender, subchainID, tnt20VoucherContractAddress, receiver, lockAmount)
	if err != nil {
		log.Fatal(err)
	}

	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	receipt, err := mainchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("lock error")
	}

	fmt.Printf("TFuel Token Lock tx hash (Mainchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering %v TNT20 tokens (Wei) from the Mainchain to Subchain %v...\n\n", lockAmount, subchainID)

	var subchainVoucherAddress common.Address
	for {
		time.Sleep(6 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getSubchainTNT20VoucherMintLogs(int(fromHeight), int(toHeight), subchainTNT20TokenBankAddress, receiver)
		if result != nil {
			subchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Subchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("Cross-chain transfer completed.\n\n")

	senderTNT20Balance, _ = instaceTNT20Contract.BalanceOf(nil, sender.From)
	instanceSubchainTNT20VoucherContract, _ := ct.NewTNT20VoucherContract(subchainVoucherAddress, subchainClient)
	receiverSubchainTNT20VoucherBalance, _ := instanceSubchainTNT20VoucherContract.BalanceOf(nil, receiver)

	fmt.Printf("Subchain TNT20 voucher contract address: %v\n", subchainVoucherAddress)
	fmt.Printf("Mainchain sender : %v, TNT20 balance on Mainchain       : %v\n", sender.From, senderTNT20Balance)
	fmt.Printf("Subchain receiver: %v, TNT20 voucher balance on Subchain: %v\n\n", receiver, receiverSubchainTNT20VoucherBalance)
}

func SubchainTNT20Lock(lockAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT20 cross-chain transfer...\n")

	subchainTNT20Address := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	if err != nil {
		log.Fatal(err)
	}
	sender := accountList[1].fromAddress
	receiver := accountList[6].fromAddress
	subchainTNT20TokenBankInstance, _ := ct.NewTNT20TokenBank(subchainTNT20TokenBankAddress, subchainClient)
	subchainTNT20Instance, _ := ct.NewTNT20VoucherContract(subchainTNT20Address, subchainClient)
	senderTNT20Balance, _ := subchainTNT20Instance.BalanceOf(nil, sender)

	expectedMainchainTNT20VoucherAddress := common.HexToAddress("0xb0DBBcba1Be5B71Dcb42aB1935773B3675e645e8")
	expectedMainchainTNT20VoucherContractInstance, _ := ct.NewTNT20VoucherContract(expectedMainchainTNT20VoucherAddress, mainchainClient)
	receiverMainchainTNT20VoucherBalance, _ := expectedMainchainTNT20VoucherContractInstance.BalanceOf(nil, receiver)

	fmt.Printf("Subchain TNT20 contract address: %v\n", subchainTNT20Address)
	fmt.Printf("Subchain sender   : %v, TNT20 balance on Subchain         : %v\n", sender, senderTNT20Balance)
	fmt.Printf("Mainchain receiver: %v, TNT20 voucher balance on Mainchain: %v\n\n", receiver, receiverMainchainTNT20VoucherBalance)

	authUser := subchainSelectAccount(subchainClient, 1)
	subchainTNT20Instance.Approve(authUser, subchainTNT20TokenBankAddress, lockAmount)

	authUser = subchainSelectAccount(subchainClient, 1)
	lockTx, err := subchainTNT20TokenBankInstance.LockTokens(authUser, big.NewInt(366), subchainTNT20Address, accountList[6].fromAddress, lockAmount)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := subchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("lock error")
	}

	fmt.Printf("TFuel Token Lock tx hash (Mainchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering %v TNT20 tokens (Wei) from to Subchain %v to the Mainchain...\n\n", lockAmount, subchainID)

	fromHeight, _ := mainchainClient.BlockNumber(context.Background())
	var mainchainVoucherAddress common.Address
	for {
		time.Sleep(6 * time.Second)
		toHeight, _ := mainchainClient.BlockNumber(context.Background())
		result := getMainchainTNT20VoucherMintLogs(int(fromHeight), int(toHeight), tnt20TokenBankAddress, accountList[6].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Mainchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("Cross-chain transfer completed.\n\n")

	senderTNT20Balance, _ = subchainTNT20Instance.BalanceOf(nil, sender)
	instanceMainchainTNT20VoucherContract, _ := ct.NewTNT20VoucherContract(mainchainVoucherAddress, mainchainClient)
	receiverMainchainTNT20VoucherBalance, _ = instanceMainchainTNT20VoucherContract.BalanceOf(nil, receiver)

	fmt.Printf("Mainchain TNT20 voucher contract address: %v\n", mainchainVoucherAddress)
	fmt.Printf("Subchain sender   : %v, TNT20 balance on Subchain         : %v\n", sender, senderTNT20Balance)
	fmt.Printf("Mainchain receiver: %v, TNT20 voucher balance on Mainchain: %v\n\n", receiver, receiverMainchainTNT20VoucherBalance)
}

func MainchainTNT20Burn(burnAmount *big.Int) {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	mainchainTNT20VoucherAddress := common.HexToAddress("0xb0DBBcba1Be5B71Dcb42aB1935773B3675e645e8")
	mainchainTNT20TokenBankInstance, err := ct.NewTNT20TokenBank(tnt20TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	mainchainTNT20VoucherAddressInstance, _ := ct.NewTNT20VoucherContract(mainchainTNT20VoucherAddress, client)
	tx, err := mainchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[6].fromAddress)
	fmt.Println("mainchain account 6 ", accountList[6].fromAddress, "tnt20 balance is", tx)
	//await TNT20Token.approve(TNT20TokenBank1.address, 20, { from: valGuarantor1 })
	authUser := mainchainSelectAccount(client, 6)
	tx1, err := mainchainTNT20VoucherAddressInstance.Approve(authUser, tnt20TokenBankAddress, burnAmount)

	authUser = mainchainSelectAccount(client, 6)
	tx1, err = mainchainTNT20TokenBankInstance.BurnVouchers(authUser, mainchainTNT20VoucherAddress, accountList[1].fromAddress, burnAmount)
	//tx, err := testSubchainTNT20TokenBankInstance.GetDenom(nil, subchainTNT20VoucherAddress)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tx)
	fmt.Println("burn tx is", tx1.Hash().Hex())
	receipt, err := client.TransactionReceipt(context.Background(), tx1.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("lock error")
	}
	time.Sleep(2 * time.Second)
	tx, err = mainchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[6].fromAddress)
	fmt.Println("mainchain account 6 ", accountList[6].fromAddress, "tnt20 balance is", tx)
	fmt.Println("---------Detecting---------")
	subchainClient, _ := ethclient.Dial("http://localhost:19888/rpc")
	subchainTNT20VoucherAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	subchainTNT20VoucherAddressInstance, _ := ct.NewTNT20VoucherContract(subchainTNT20VoucherAddress, subchainClient)
	tx, err = subchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
	fmt.Println(accountList[1].fromAddress, "subchain_account_balance is", tx)
	for {
		time.Sleep(2 * time.Second)
		tx_subchain, _ := subchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
		if tx.Cmp(tx_subchain) != 0 {
			tx = tx_subchain
			break
		}
	}
	fmt.Println(accountList[1].fromAddress, "subchain_account_balance is", tx)
}

func SubchainTNT20Burn(burnAmount *big.Int) {
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	subchainTNT20VoucherAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	if err != nil {
		log.Fatal(err)
	}
	testSubchainTNT20TokenBankInstance, err := ct.NewTNT20TokenBank(subchainTNT20TokenBankAddress, client)
	subchainTNT20VoucherAddressInstance, _ := ct.NewTNT20VoucherContract(subchainTNT20VoucherAddress, client)
	tx, err := subchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
	fmt.Println(accountList[1].fromAddress, "subchain_account_balance is", tx)
	//await TNT20Token.approve(TNT20TokenBank1.address, 20, { from: valGuarantor1 })
	authUser := subchainSelectAccount(client, 1)
	tx1, err := subchainTNT20VoucherAddressInstance.Approve(authUser, subchainTNT20TokenBankAddress, burnAmount)

	authUser = subchainSelectAccount(client, 1)
	tx1, err = testSubchainTNT20TokenBankInstance.BurnVouchers(authUser, subchainTNT20VoucherAddress, accountList[1].fromAddress, burnAmount)
	//tx, err := testSubchainTNT20TokenBankInstance.GetDenom(nil, subchainTNT20VoucherAddress)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tx)
	fmt.Println("burn tx is", tx1.Hash().Hex())
	receipt, err := client.TransactionReceipt(context.Background(), tx1.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("lock error")
		return
	}
	time.Sleep(2 * time.Second)
	tx, err = subchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
	fmt.Println(accountList[1].fromAddress, "subchain_account_balance is", tx)
	fmt.Println("---------Detecting---------")
	mainchainClient, _ := ethclient.Dial("http://localhost:18888/rpc")
	mainchainTNT20VoucherAddressInstance, _ := ct.NewTNT20VoucherContract(tnt20VoucherContractAddress, mainchainClient)
	tx, err = mainchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
	fmt.Println(accountList[1].fromAddress, "mainchain_account_balance is", tx)
	for {
		time.Sleep(2 * time.Second)
		tx_subchain, _ := mainchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
		//fmt.Println(accountList[1].fromAddress, "mainchain_account_balance is", tx_subchain)
		if tx.Cmp(tx_subchain) != 0 {
			tx = tx_subchain
			break
		}
	}
	fmt.Println(accountList[1].fromAddress, "mainchain_account_balance is", tx)
}
