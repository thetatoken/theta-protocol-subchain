package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/thetatoken/theta/crypto"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/thetasubchain/eth/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.HexToECDSA("1111111111111111111111111111111111111111111111111111111111111111")
	privateKey, err := crypto.HexToECDSA("93a90ea508331dfdf27fb79757d4250b4e84954927ba0073cd67454ac432c737")

	// privateKey, err := crypto.HexToECDSA("2dad160420b1e9b6fc152cd691a686a7080a0cee41b98754597a2ce57cc5dab1")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fr := privateKey.PublicKey

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x79EaFd0B5eC8D3f945E6BB2817ed90b046c0d0Af")
	instance, err := ct.NewSubchainERC(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// 1 mint
	amount, suceess := new(big.Int).SetString("2000000000000000000000", 10)
	if !suceess {
		log.Fatal("Failed string")
	}

	tx, err := instance.Mint(auth, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	// 2 GetBalance
	// toaddress := common.HexToAddress("0x2E833968E5bB786Ae419c4d13189fB081Cc43bab")
	// result, err := instance.GetBalance(nil, toaddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	// 3 DepositCollateral
	// toaddress := common.HexToAddress("0x2E833968E5bB786Ae419c4d13189fB081Cc43bab")
	// amount := new(big.Int).SetUint64(uint64(23580001))
	// tx, err := instance.DepositCollateral(auth, toaddress, amount)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	// 4 GetLegalValidators

	// result, err := instance.GetLegalValidators(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	// 5 Burn
	// amount := new(big.Int).SetUint64(uint64(20001))
	// tx, err := instance.Burn(auth, amount)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	// 6 Approve
	// amount := new(big.Int).SetUint64(uint64(5000000))
	// tnt20Addr := common.HexToAddress("0x031375c16fc20ECEd29fA21642701EE48A5b9090")
	// tx, err := instance.Approve(auth, tnt20Addr, amount)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

}
