package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/crypto/sha3"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/thetasubchain/eth/ethclient"

	// rg "chainRegistrarOnMainchain" // for demo
	ct "github.com/thetatoken/thetasubchain/interchain/contracts/contract"
)

type accounts struct {
	priKey      string
	privateKey  *ecdsa.PrivateKey
	fromAddress common.Address
}

var wthetaAddress common.Address
var registerOnMainchainAddress common.Address
var governanceTokenAddress common.Address
var tnt20VoucherContractAddress common.Address
var tnt20TokenBankAddress common.Address
var subchaintnt20TokenBankAddress common.Address
var accountList []accounts
var tnt721TokenBankAddress common.Address
var tnt721VoucherContractAddress common.Address
var tfuelTokenbankAddress common.Address
var subchainTfuelTokenBank common.Address
var Subchaintnt721TokenBankAddress common.Address
var subchainID *big.Int

func keccak256(data ...[]byte) []byte {
	d := sha3.NewKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}
func pubkeyToAddress(p ecdsa.PublicKey) common.Address {
	pubBytes := crypto.FromECDSAPub(&p)
	return common.BytesToAddress(keccak256(pubBytes[1:])[12:])
}
func init() {
	map1 := config()

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

		fromAddress := pubkeyToAddress(*publicKeyECDSA)
		fmt.Println(value, "-----", fromAddress)
		accountList = append(accountList, accounts{priKey: value, privateKey: privateKey, fromAddress: fromAddress})
	}

	// var dec18 = new(big.Int)
	// dec18.SetString("1000000000000000000", 10)
	// //amount := new(big.Int).Mul(dec18, big.NewInt(200000))
	// client, err := ethclient.Dial("http://localhost:18888/rpc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// nonce, err := client.PendingNonceAt(context.Background(), accountList[9].fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// queryStr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"thetacli.Send","params":[{"chain_id":"privatenet", "from":"%v", "to":"%v", "thetawei":"9900000", "tfuelwei":"88000000", "fee":"100000000", "sequence":"%v", "async":true}],"id":1}`, accountList[9].fromAddress, accountList[10].fromAddress, fmt.Sprintf("%d", nonce))
	// var jsonData = []byte(queryStr)

	// request, err := http.NewRequest("POST", "http://localhost:18888/rpc", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// request.Header.Set("Content-Type", "application/json")

	// client1 := &http.Client{}
	// response, err := client1.Do(request)
	// if err != nil {
	// 	log.Fatalf("response error : %v", err)
	// }
	// defer response.Body.Close()
	// fmt.Println(response)
}
func mainchainSelectAccount(client *ethclient.Client, id int) *bind.TransactOpts {
	time.Sleep(1 * time.Second)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//chainID := big.NewInt(360777)
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

	auth, err := bind.NewKeyedTransactorWithChainID(crypto.ECDSAToPrivKey(privateKey), chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	// auth.Value = big.NewInt(20000000000000000000) // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth
}
func subchainSelectAccount(client *ethclient.Client, id int) *bind.TransactOpts {
	time.Sleep(1 * time.Second)
	// chainID, err := client.ChainID(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	chainID := big.NewInt(360777)
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

	auth, err := bind.NewKeyedTransactorWithChainID(crypto.ECDSAToPrivKey(privateKey), chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	// auth.Value = big.NewInt(20000000000000000000) // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth
}
func oneAccountRegister() {
	//func main() {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	//tnt20TokenBankAddress := common.HexToAddress("0x1f629139b3b4A03799c6e6655b7F59a1F01598E7")

	chainGuarantor := accountList[7].fromAddress

	instanceWrappedTheta, err := ct.NewMockWrappedTheta(wthetaAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	if err != nil {
		log.Fatal("hhh", err)
	}
	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(registerOnMainchainAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	amount := new(big.Int).Mul(dec18, big.NewInt(200000))

	auth := mainchainSelectAccount(client, 7) //chainGuarantor
	tx, err := instanceWrappedTheta.Mint(auth, chainGuarantor, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx)
	approveAmount := new(big.Int).Mul(dec18, big.NewInt(50000))
	authchainGuarantor := mainchainSelectAccount(client, 7)
	fmt.Println(instanceWrappedTheta.BalanceOf(nil, chainGuarantor))
	tx, err = instanceWrappedTheta.Approve(authchainGuarantor, registerOnMainchainAddress, approveAmount)
	if err != nil {
		log.Fatal(err)
	}

	registerAmount := new(big.Int).Mul(dec18, big.NewInt(40000))
	authchainGuarantor = mainchainSelectAccount(client, 7)
	tx, err = instanceChainRegistrar.RegisterSubchain(authchainGuarantor, subchainID, governanceTokenAddress, registerAmount, "111111")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	x, _ := instanceChainRegistrar.GetAllSubchainIDs(nil)
	fmt.Println(x)

}
func oneAcoountStake(id int) {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	instanceWrappedTheta, err := ct.NewMockWrappedTheta(wthetaAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	instanceGovernanceToken, err := ct.NewSubchainGovernanceToken(governanceTokenAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(registerOnMainchainAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	validator1 := accountList[id].fromAddress
	validatorCollateralManagerAddr, _ := instanceChainRegistrar.Vcm(nil)
	validatorStakeManagerAddr, _ := instanceChainRegistrar.Vsm(nil)

	//Deposit wTHETA collateral to validators
	validatorCollateral := new(big.Int).Mul(dec18, big.NewInt(2000))
	authValidator1 := mainchainSelectAccount(client, id) //Validator1

	tx, err := instanceWrappedTheta.Mint(authValidator1, validator1, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	authValidator1 = mainchainSelectAccount(client, id) //Validator1
	tx, err = instanceWrappedTheta.Approve(authValidator1, validatorCollateralManagerAddr, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("validator1 collateral")
	// fmt.Println(instanceWrappedTheta.BalanceOf(nil, accountList[1].fromAddress))
	authValidator1 = mainchainSelectAccount(client, id) //Validator1
	tx, err = instanceChainRegistrar.DepositCollateral(authValidator1, subchainID, validator1, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	//-------------------------------------
	//Stake to the validators
	validatorStakingAmount := new(big.Int).Mul(dec18, big.NewInt(100000))
	validatorStakingAmountMint := new(big.Int)
	validatorStakingAmountMint.Mul(validatorStakingAmount, big.NewInt(10))
	//govTokenInitDistrWalletAddress:=accountList[6].fromAddress

	// fmt.Println("validator1 deposit")
	// fmt.Println(instanceGovernanceToken.BalanceOf(nil, validator1))
	authGovTokenInitDistrWallet := mainchainSelectAccount(client, 6)
	tx, err = instanceGovernanceToken.Transfer(authGovTokenInitDistrWallet, validator1, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("validator1 deposit")
	// fmt.Println(instanceGovernanceToken.BalanceOf(nil, validator1))
	authValidator1 = mainchainSelectAccount(client, id) //Validator1
	tx, err = instanceGovernanceToken.Approve(authValidator1, validatorStakeManagerAddr, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("validator1 allowance")
	// fmt.Println(instanceGovernanceToken.Allowance(nil, accountList[1].fromAddress, validatorStakeManagerAddr))
	// fmt.Println("wallet deposited")
	// fmt.Println(instanceGovernanceToken.BalanceOf(nil, validatorStakeManagerAddr))

	authValidator1 = mainchainSelectAccount(client, id) //Validator1
	tx, err = instanceChainRegistrar.DepositStake(authValidator1, subchainID, validator1, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("validator2 deposited")
	// fmt.Println(instanceGovernanceToken.BalanceOf(nil, accountList[2].fromAddress))

	//height1 := big.NewInt(int64(height))
	fmt.Println(tx)

	time.Sleep(5 * time.Second)
	height, _ := client.BlockNumber(context.Background())
	fmt.Println(big.NewInt(int64(height)))
	dynasty := big.NewInt(int64(height/100 + 1))
	tx1, tx2 := instanceChainRegistrar.GetValidatorSet(nil, subchainID, dynasty)
	fmt.Println("set", tx1)
	fmt.Println("amout", tx2)
	// tx3, err := instanceChainRegistrar.GetStakeSnapshotHeights(nil, subchainID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(tx3)
}
func claimStake() {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(registerOnMainchainAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	var validatorStakingAmount = new(big.Int)

	//validatorStakingAmount := new(big.Int).Mul(dec18, big.NewInt(100000))
	// authGovTokenInitDistrWallet := mainchainSelectAccount(client, 6)
	// validator1 := accountList[12].fromAddress
	// instanceGovernanceToken, err := ct.NewSubchainGovernanceToken(governanceTokenAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	// tx, err := instanceGovernanceToken.Transfer(authGovTokenInitDistrWallet, validator1, validatorStakingAmount)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//validatorStakeManagerAddr, _ := instanceChainRegistrar.Vsm(nil)
	// authValidator1 := mainchainSelectAccount(client, 12) //Validator1
	// tx, err = instanceGovernanceToken.Approve(authValidator1, validatorStakeManagerAddr, validatorStakingAmount)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// authValidator1 = mainchainSelectAccount(client, 12) //Validator1
	// tx, err = instanceChainRegistrar.DepositStake(authValidator1, subchainID, validator1, validatorStakingAmount)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	height, _ := client.BlockNumber(context.Background())
	fmt.Println(big.NewInt(int64(height)))
	dynasty := big.NewInt(int64(height/100 + 1))
	re, _ := instanceChainRegistrar.GetValidatorSet(nil, subchainID, dynasty)
	fmt.Println(re)
	//validatorStakingAmount.SetString("1621710387562809105166", 10)
	validatorStakingAmount = re.ShareAmounts[0]
	authValidator1 := mainchainSelectAccount(client, 10) //Validator1
	tx, err := instanceChainRegistrar.WithdrawStake(authValidator1, subchainID, accountList[10].fromAddress, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("error")
	}
	// validatorStakingAmount = re.ShareAmounts[1]
	// authValidator1 = mainchainSelectAccount(client, 10) //Validator1
	// tx, err = instanceChainRegistrar.WithdrawStake(authValidator1, subchainID, accountList[10].fromAddress, validatorStakingAmount)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(tx.Hash().Hex())
	// receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if receipt.Status != 1 {
	// 	fmt.Println("error")
	// }
}
