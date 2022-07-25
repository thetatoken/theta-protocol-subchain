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

// rg "chainRegistrarOnMainchain" // for demo
func main1() {
	AccountsInit()
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	tx3, err := instanceChainRegistrar.GetStakeSnapshot(nil, big.NewInt(360777), big.NewInt(463))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx3)
}
func tokenLocked() {
	AccountsInit()
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainID := big.NewInt(360777)
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	user := accountList[1].fromAddress

	instanceTNT20VoucherContract, err := ct.NewTNT20VoucherContract(TNT20VoucherContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	instanceTNT20TokenBank, err := ct.NewTNT20TokenBank(TNT20TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := SelectAccount(client, 0)
	_, err = instanceTNT20VoucherContract.Mint(authAccount0, user, big.NewInt(30))
	if err != nil {
		log.Fatal(err)
	}
	authUser := SelectAccount(client, 1)
	_, err = instanceTNT20VoucherContract.Approve(authUser, TNT20TokenBankAddress, big.NewInt(30))
	if err != nil {
		log.Fatal(err)
	}

	authUser = SelectAccount(client, 1)
	LockTx, err := instanceTNT20TokenBank.LockTokens(authUser, subchainID, TNT20VoucherContractAddress, user, big.NewInt(20))
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
	fmt.Println(LockTx.Hash())
	var subchainVoucherAddress common.Address
	for {
		time.Sleep(2 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getMintlog(int(fromHeight), int(toHeight), SubchainTNT20TokenBankAddress, user)
		if result != nil {
			subchainVoucherAddress = *result
			break
		}
	}
	fmt.Println(subchainVoucherAddress)
}
func mint() {
	AccountsInit()

	// oneAccountRegister()
	// oneAcoountStake()
	client, _ := ethclient.Dial("http://localhost:19888/rpc")

	auth := SubchainSelectAccount(client, 1)
	instance, _ := ct.NewTestSubchainTNT20TokenBank(SubchainTNT20TokenBankAddress, client)
	tx, err := instance.MintVouchers(auth, "366/20/"+"0x7f1C87Bd3a22159b8a2E5D195B1a3283D10ea895", "TNT20Name", "vTNT20Symbol", uint8(1), accountList[1].fromAddress, big.NewInt(20), big.NewInt(139), big.NewInt(1))
	fmt.Println("-------------")
	fmt.Println(tx.Hash().Hex(), err)
	height, _ := client.BlockNumber(context.Background())
	fmt.Println(height)
}

// const RawABI = `
// [
// 	{
// 		"inputs": [
// 			{
// 				"components": [
// 					{
// 						"internalType": "address",
// 						"name": "validator",
// 						"type": "address"
// 					},
// 					{
// 						"internalType": "uint256",
// 						"name": "shareAmount",
// 						"type": "uint256"
// 					}
// 				],
// 				"internalType": "struct TestRegister.ValidatorAddrSharePair[]",
// 				"name": "a",
// 				"type": "tuple[]"
// 			}
// 		],
// 		"stateMutability": "nonpayable",
// 		"type": "constructor"
// 	}
// ]`

// type hhh struct {
// 	Validator   common.Address
// 	ShareAmount *big.Int
// }
func subchainBurnTNT20() {
	AccountsInit()
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	authUser := SubchainSelectAccount(client, 1)
	subchainTNT20VoucherAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	testSubchainTNT20TokenBankInstance, err := ct.NewTestSubchainTNT20TokenBank(SubchainTNT20TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT20VoucherAddressInstance, _ := ct.NewTNT20VoucherContract(subchainTNT20VoucherAddress, client)
	// tx, err := subchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
	//await TNT20Token.approve(TNT20TokenBank1.address, 20, { from: valGuarantor1 })
	tx, err := subchainTNT20VoucherAddressInstance.Approve(authUser, SubchainTNT20TokenBankAddress, big.NewInt(10))

	authUser = SubchainSelectAccount(client, 1)
	tx, err = testSubchainTNT20TokenBankInstance.BurnVouchers(authUser, subchainTNT20VoucherAddress, accountList[6].fromAddress, big.NewInt(10))
	//tx, err := testSubchainTNT20TokenBankInstance.GetDenom(nil, subchainTNT20VoucherAddress)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tx)
	fmt.Println(tx.Hash().Hex())
}
func subchainLockTNT20() {
	AccountsInit()
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	//authUser := SubchainSelectAccount(client, 1)
	subchainTNT20VoucherAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	//testSubchainTNT20TokenBankInstance, err := ct.NewTestSubchainTNT20TokenBank(SubchainTNT20TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT20VoucherAddressInstance, _ := ct.NewTNT20VoucherContract(subchainTNT20VoucherAddress, client)
	tx, err := subchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[1].fromAddress)
	//await TNT20Token.approve(TNT20TokenBank1.address, 20, { from: valGuarantor1 })
	//tx, err := subchainTNT20VoucherAddressInstance.Approve(authUser, SubchainTNT20TokenBankAddress, big.NewInt(20))

	//authUser = SubchainSelectAccount(client, 1)
	//tx, err = testSubchainTNT20TokenBankInstance.LockTokens(authUser, big.NewInt(366), subchainTNT20VoucherAddress, accountList[6].fromAddress, big.NewInt(20))
	//tx, err := testSubchainTNT20TokenBankInstance.GetDenom(nil, subchainTNT20VoucherAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx)
	//fmt.Println(tx.Hash().Hex())
}
func getSubchainBlockHeight() {
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	height, _ := client.BlockNumber(context.Background())
	fmt.Println(big.NewInt(int64(height)))
}
func mainchainBurnTNT20() {
	AccountsInit()
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	mainchainTNT20VoucherAddress := common.HexToAddress("0x20B2897c4f95df71a5a8A62Ae812f5843AA92E25")
	mainchainTNT20TokenBankInstance, err := ct.NewTNT20TokenBank(TNT20TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	mainchainTNT20VoucherAddressInstance, _ := ct.NewTNT20VoucherContract(mainchainTNT20VoucherAddress, client)
	tx, err := mainchainTNT20VoucherAddressInstance.BalanceOf(nil, accountList[6].fromAddress)
	fmt.Println(tx)
	//await TNT20Token.approve(TNT20TokenBank1.address, 20, { from: valGuarantor1 })
	authUser := SelectAccount(client, 6)
	tx1, err := mainchainTNT20VoucherAddressInstance.Approve(authUser, TNT20TokenBankAddress, big.NewInt(11))

	authUser = SelectAccount(client, 6)
	tx1, err = mainchainTNT20TokenBankInstance.BurnVouchers(authUser, mainchainTNT20VoucherAddress, accountList[1].fromAddress, big.NewInt(10))
	//tx, err := testSubchainTNT20TokenBankInstance.GetDenom(nil, subchainTNT20VoucherAddress)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tx)
	fmt.Println(tx1.Hash().Hex())
}
func main() {
	AccountsInit()
	mainchainTfuelLock()
	//tnt721Lock()
	// oneAccountRegister()
	// oneAcoountStake()
	//subchainBurnTNT20()
	//getMintlog(30, 50, SubchainTNT20TokenBankAddress)
	//mint()
	//subchainLockTNT20()
	// tokenLocked()
	// tnt721Lock()
	//mainchainTfuelLock()
	//getMainchainMintlog(18770, 18780, TNT20TokenBankAddress)
	//mainchainBurnTNT20()
	//tokenLocked()
	//subchainBurnTNT20()
	// parsed, err := abi.JSON(strings.NewReader(RawABI))
	// if err != nil {
	// 	panic(err)
	// }
	// var a []hhh
	// c := &hhh{
	// 	Validator:   common.HexToAddress("0x7f1C87Bd3a22159b8a2E5D195B1a3283D10ea895"),
	// 	ShareAmount: big.NewInt(2),
	// }
	// a = append(a, *c)
	// valueInput, err := parsed.Pack("", a)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(hex.EncodeToString(valueInput))
	//mint()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("ok!")

	//fmt.Println(common.BytesToAddress([]byte("ZD7\x04\xddKYK8,\"\xa0\x83\xe2\xbd0\x90\xa6\xfe\xf3")))
	//00000000000000000000000000000000000000b7

}
