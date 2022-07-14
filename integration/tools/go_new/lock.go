package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	// rg "chainRegistrarOnMainchain" // for demo
	ct "github.com/thetatoken/thetasubchain/integration/tools/go_new/accessors"
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
	var success bool
	auth.Value, success = new(big.Int).SetString("3000000000000000000000", 10)
	if !success {
		log.Fatal("Failed string")
	}
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

	//WthetaAddress := common.HexToAddress("0x17deB845AAcCA09873D1984784EbdC59DBB38846")
	RegisterOnMainchainAddress := common.HexToAddress("0xDE06587B53D752ac280c7Dae1D11C3ffc341A509")
	// GovernanceTokenAddress := common.HexToAddress("0xD323FbDE9D287A288d22F9F4B30E61d2500aa962")
	clientNew, _ := chooseAccount(client, 3) //chainGuarantor
	instance, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, clientNew)
	//subchainID := big.NewInt(9988)
	tx, err := instance.Getallsubchainids(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx)
}
