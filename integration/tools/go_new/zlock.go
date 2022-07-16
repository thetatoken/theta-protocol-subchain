package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	ct "github.com/thetatoken/thetasubchain/integration/tools/go_new/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// rg "chainRegistrarOnMainchain" // for demo
)

type accounts struct {
	priKey      string
	privateKey  *ecdsa.PrivateKey
	fromAddress common.Address
}

var accountList []accounts

func accountsInit() {
	var map1 []string

	map1 = append(map1, "1111111111111111111111111111111111111111111111111111111111111111")
	map1 = append(map1, "2222222222222222222222222222222222222222222222222222222222222222")
	map1 = append(map1, "3333333333333333333333333333333333333333333333333333333333333333")
	map1 = append(map1, "4444444444444444444444444444444444444444444444444444444444444444")
	map1 = append(map1, "5555555555555555555555555555555555555555555555555555555555555555")
	map1 = append(map1, "6666666666666666666666666666666666666666666666666666666666666666")
	map1 = append(map1, "7777777777777777777777777777777777777777777777777777777777777777")
	map1 = append(map1, "8888888888888888888888888888888888888888888888888888888888888888")
	map1 = append(map1, "9999999999999999999999999999999999999999999999999999999999999999")
	map1 = append(map1, "1000000000000000000000000000000000000000000000000000000000000000")
	// privateKey, err := crypto.HexToECDSA("1111111111111111111111111111111111111111111111111111111111111111")
	for _, value := range map1 {

		privateKey, err := crypto.HexToECDSA(value)

		// privateKey, err := crypto.HexToECDSA("2dad160420b1e9b6fc152cd691a686a7080a0cee41b98754597a2ce57cc5dab1")
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		fmt.Println(value, "-----", fromAddress)
		accountList = append(accountList, accounts{priKey: value, privateKey: privateKey, fromAddress: fromAddress})
	}

}
func chooseAccount(client *ethclient.Client, id int) (*ethclient.Client, *bind.TransactOpts) {

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fromAddress := accountList[id].fromAddress
	privateKey := accountList[id].privateKey
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	// auth.Value = big.NewInt(20000000000000000000) // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	return client, auth
}
func main() {
	accountsInit()
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainID := big.NewInt(118)
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	//validatorStakingAmount := new(big.Int).Mul(dec18, big.NewInt(100000))
	//validator1 := accountList[1].fromAddress
	//WthetaAddress := common.HexToAddress("0x17deB845AAcCA09873D1984784EbdC59DBB38846")
	//WthetaAddress := common.HexToAddress("0x5395019D1d4794eb5Ac0AC51976ee48995bdA694")
	RegisterOnMainchainAddress := common.HexToAddress("0x1EaE66323eAeb16dAD8e162520251D49D4E2f6a3")
	GovernanceTokenAddress := common.HexToAddress("0x6c1d3aBA75Fea74000a10BbfbE86be30f20a2F3D")
	//TNT20TokenBankAddress := common.HexToAddress("0x1f629139b3b4A03799c6e6655b7F59a1F01598E7")
	// clientNew, _ := chooseAccount(client, 3) //chainGuarantor
	// instance, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, clientNew)
	// //height, _ := client.BlockNumber(context.Background())
	// //height1 := big.NewInt(int64(height))
	// tx, err := instance.GetStakeSnapshotHeights(nil, subchainID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(tx)
	clientValidator3, _ := chooseAccount(client, 3) //Validator1
	instanceChainRegistrarValidator3, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, clientValidator3)
	validatorStakeManagerAddr, _ := instanceChainRegistrarValidator3.Vsm(nil)
	fmt.Println("validator3 allowance")
	instanceGovernanceTokengovTokenInitDistrWallet, err := ct.NewSubchainGovernanceToken(GovernanceTokenAddress, client)
	fmt.Println(instanceGovernanceTokengovTokenInitDistrWallet.Allowance(nil, accountList[3].fromAddress, validatorStakeManagerAddr))
	fmt.Println("validator3 balance")
	fmt.Println(instanceGovernanceTokengovTokenInitDistrWallet.BalanceOf(nil, accountList[3].fromAddress))

	// clientValidator3, authValidator3 = chooseAccount(client, 3) //Validator1
	// _, err = instanceChainRegistrarValidator3.DepositStake(authValidator3, subchainID, accountList[3].fromAddress, validatorStakingAmount)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	time.Sleep(2 * time.Second)
	tx, err := instanceChainRegistrarValidator3.GetStakeSnapshotHeights(nil, subchainID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx)
}
