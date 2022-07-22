package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/thetatoken/thetasubchain/eth/abi"
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
	tx3, err := instanceChainRegistrar.GetStakeSnapshot(nil, big.NewInt(360777), big.NewInt(44816))
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
	user := accountList[6].fromAddress

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
	authUser := SelectAccount(client, 6)
	_, err = instanceTNT20VoucherContract.Approve(authUser, TNT20TokenBankAddress, big.NewInt(30))
	if err != nil {
		log.Fatal(err)
	}
	authUser = SelectAccount(client, 6)
	_, err = instanceTNT20TokenBank.LockTokens(authUser, subchainID, TNT20VoucherContractAddress, user, big.NewInt(20))
	if err != nil {
		log.Fatal(err)
	}
}
func mint() {
	AccountsInit()

	// oneAccountRegister()
	// oneAcoountStake()
	client, _ := ethclient.Dial("http://localhost:19888/rpc")

	auth := SubchainSelectAccount(client, 1)
	instance, _ := ct.NewTestSubchainTNT20TokenBank(SubchainTNT20TokenBankAddress, client)
	tx, err := instance.MintVouchers(auth, "366/20/"+"0x7f1C87Bd3a22159b8a2E5D195B1a3283D10ea895", "TNT20Name", "vTNT20Symbol", uint8(1), accountList[1].fromAddress, big.NewInt(20), big.NewInt(7007), big.NewInt(1))
	fmt.Println("-------------")
	fmt.Println(tx.Hash().Hex(), err)
	height, _ := client.BlockNumber(context.Background())
	fmt.Println(height)
}

const RawABI = `
[
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "address",
						"name": "validator",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "shareAmount",
						"type": "uint256"
					}
				],
				"internalType": "struct TestRegister.ValidatorAddrSharePair[]",
				"name": "a",
				"type": "tuple[]"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	}
]`

type hhh struct {
	Validator   common.Address
	ShareAmount *big.Int
}

func main() {
	parsed, err := abi.JSON(strings.NewReader(RawABI))
	if err != nil {
		panic(err)
	}
	var a []hhh
	c := &hhh{
		Validator:   common.HexToAddress("0x7f1C87Bd3a22159b8a2E5D195B1a3283D10ea895"),
		ShareAmount: big.NewInt(2),
	}
	a = append(a, *c)
	valueInput, err := parsed.Pack("", a)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(valueInput))
	//mint()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("ok!")

	//fmt.Println(common.BytesToAddress([]byte("ZD7\x04\xddKYK8,\"\xa0\x83\xe2\xbd0\x90\xa6\xfe\xf3")))
	//00000000000000000000000000000000000000b7

}
