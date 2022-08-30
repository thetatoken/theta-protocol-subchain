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

func MainchainTNT20Lock(mintLockAmount *big.Int) {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	user := accountList[1].fromAddress

	instanceTNT20VoucherContract, err := ct.NewTNT20VoucherContract(tnt20VoucherContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	instanceTNT20TokenBank, err := ct.NewTNT20TokenBank(tnt20TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := mainchainSelectAccount(client, 0)
	_, err = instanceTNT20VoucherContract.Mint(authAccount0, user, mintLockAmount)
	if err != nil {
		log.Fatal(err)
	}
	authUser := mainchainSelectAccount(client, 1)
	_, err = instanceTNT20VoucherContract.Approve(authUser, tnt20TokenBankAddress, mintLockAmount)
	if err != nil {
		log.Fatal(err)
	}
	//
	initNum, _ := instanceTNT20VoucherContract.BalanceOf(nil, accountList[1].fromAddress)
	fmt.Println("account ", accountList[1].fromAddress, " balance on mainchain  is ", initNum)
	//
	authUser = mainchainSelectAccount(client, 1)
	LockTx, err := instanceTNT20TokenBank.LockTokens(authUser, subchainID, tnt20VoucherContractAddress, user, mintLockAmount)
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, _ := ethclient.Dial("http://localhost:19888/rpc")
	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	receipt, err := client.TransactionReceipt(context.Background(), LockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("lock error")
	}
	fmt.Println("lock tx hash is", LockTx.Hash().Hex())
	//
	lastNum, _ := instanceTNT20VoucherContract.BalanceOf(nil, accountList[1].fromAddress)
	fmt.Println("account ", accountList[1].fromAddress, " balance on mainchain  is ", lastNum)
	fmt.Println("---------Detecting---------")
	//
	var subchainVoucherAddress common.Address
	for {
		time.Sleep(2 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getSubchainTNT20VoucherMintLogs(int(fromHeight), int(toHeight), subchainTNT20TokenBankAddress, user)
		if result != nil {
			subchainVoucherAddress = *result
			break
		}
	}
	fmt.Println("Subchain TNT20 voucher contract address is", subchainVoucherAddress)
	instanceSubchainTNT20VoucherContractAddress, _ := ct.NewTNT20VoucherContract(subchainVoucherAddress, subchainClient)
	//
	balanceOfSubchain, _ := instanceSubchainTNT20VoucherContractAddress.BalanceOf(nil, accountList[1].fromAddress)
	fmt.Println("account ", accountList[1].fromAddress, " balance on subchain is ", balanceOfSubchain)
	//
}
func SubchainTNT20Lock(mintLockAmount *big.Int) {
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
	tx2, err := subchainTNT20VoucherAddressInstance.Approve(authUser, subchainTNT20TokenBankAddress, mintLockAmount)

	authUser = subchainSelectAccount(client, 1)
	tx2, err = testSubchainTNT20TokenBankInstance.LockTokens(authUser, big.NewInt(366), subchainTNT20VoucherAddress, accountList[6].fromAddress, mintLockAmount)
	//tx, err := testSubchainTNT20TokenBankInstance.GetDenom(nil, subchainTNT20VoucherAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("lock tx hash is", tx2.Hash().Hex())

	receipt, err := client.TransactionReceipt(context.Background(), tx2.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("lock error")
	}
	tx, err = subchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
	fmt.Println(accountList[1].fromAddress, "subchain_account_balance is", tx)
	fmt.Println("---------Detecting---------")
	mainchainClient, _ := ethclient.Dial("http://localhost:18888/rpc")
	fromHeight, _ := mainchainClient.BlockNumber(context.Background())
	var mainchainVoucherAddress common.Address
	for {
		time.Sleep(2 * time.Second)
		toHeight, _ := mainchainClient.BlockNumber(context.Background())
		result := getMainchainTNT20VoucherMintLogs(int(fromHeight), int(toHeight), tnt20TokenBankAddress, accountList[6].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
	}
	fmt.Println("Mainchain TNT20 voucher contract address is", mainchainVoucherAddress)
	instanceMainchainTNT20VoucherContractAddress, _ := ct.NewTNT20VoucherContract(mainchainVoucherAddress, mainchainClient)
	//
	balanceOfSubchain, _ := instanceMainchainTNT20VoucherContractAddress.BalanceOf(nil, accountList[6].fromAddress)
	fmt.Println("account ", accountList[6].fromAddress, " balance on mainchain is ", balanceOfSubchain)
	//
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
func TNT20Query(chainID int64, contractAddress, accountAddress string) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	tnt20ContractAddress := common.HexToAddress(contractAddress) // FIXME: should instantiate a mock TNT20 instead of using the Voucher contract (which causes confusion)
	var instaceTNT20Contract *ct.TNT20VoucherContract
	if chainID == 366 {
		fmt.Printf("Preparing for TNT20 mainchain query...\n")
		instaceTNT20Contract, err = ct.NewTNT20VoucherContract(tnt20ContractAddress, mainchainClient)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Preparing for TNT20 subchain query...\n")
		instaceTNT20Contract, err = ct.NewTNT20VoucherContract(tnt20ContractAddress, subchainClient)
		if err != nil {
			log.Fatal(err)
		}
	}
	balance ,_:=instaceTNT20Contract.BalanceOf(nil,common.HexToAddress(accountAddress))
	fmt.Println("Account ",accountAddress," TNT20 balance in ",contractAddress," is ",balance)

}
