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
	ct "github.com/thetatoken/thetasubchain/integration/tools/go_new/accessors"
)

type accounts struct {
	priKey      string
	privateKey  *ecdsa.PrivateKey
	fromAddress common.Address
}

var WthetaAddress common.Address
var RegisterOnMainchainAddress common.Address
var GovernanceTokenAddress common.Address
var TNT20VoucherContractAddress common.Address
var TNT20TokenBankAddress common.Address
var SubchainTNT20TokenBankAddress common.Address
var accountList []accounts
var TNT721TokenBankAddress common.Address
var TNT721VoucherContractAddress common.Address

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
func AccountsInit() {
	var map1 []string

	map1 = append(map1, "1111111111111111111111111111111111111111111111111111111111111111")
	map1 = append(map1, "93a90ea508331dfdf27fb79757d4250b4e84954927ba0073cd67454ac432c737")
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

		fromAddress := pubkeyToAddress(*publicKeyECDSA)
		//fmt.Println(value, "-----", fromAddress)
		accountList = append(accountList, accounts{priKey: value, privateKey: privateKey, fromAddress: fromAddress})
	}

	WthetaAddress = common.HexToAddress("0x7d73424a8256C0b2BA245e5d5a3De8820E45F390")
	RegisterOnMainchainAddress = common.HexToAddress("0x08425D9Df219f93d5763c3e85204cb5B4cE33aAa")
	GovernanceTokenAddress = common.HexToAddress("0x6E05f58eEddA592f34DD9105b1827f252c509De0")
	TNT20VoucherContractAddress = common.HexToAddress("0x4fb87c52Bb6D194f78cd4896E3e574028fedBAB9")
	TNT20TokenBankAddress = common.HexToAddress("0x2Ce636d6240f8955d085a896e12429f8B3c7db26")
	SubchainTNT20TokenBankAddress = common.HexToAddress("0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D")

	TNT721TokenBankAddress = common.HexToAddress("0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA")
	TNT721VoucherContractAddress = common.HexToAddress("0x560A0c0CA6B0A67895024dae77442C5fd3DC473e")
}
func SelectAccount(client *ethclient.Client, id int) *bind.TransactOpts {
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
func SubchainSelectAccount(client *ethclient.Client, id int) *bind.TransactOpts {
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

	subchainID := big.NewInt(360777)

	//TNT20TokenBankAddress := common.HexToAddress("0x1f629139b3b4A03799c6e6655b7F59a1F01598E7")

	chainGuarantor := accountList[7].fromAddress

	instanceWrappedTheta, err := ct.NewMockWrappedTheta(WthetaAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	if err != nil {
		log.Fatal("hhh", err)
	}
	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	amount := new(big.Int).Mul(dec18, big.NewInt(200000))

	auth := SelectAccount(client, 7) //chainGuarantor
	tx, err := instanceWrappedTheta.Mint(auth, chainGuarantor, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx)
	approveAmount := new(big.Int).Mul(dec18, big.NewInt(50000))
	authchainGuarantor := SelectAccount(client, 7)
	fmt.Println(instanceWrappedTheta.BalanceOf(nil, chainGuarantor))
	tx, err = instanceWrappedTheta.Approve(authchainGuarantor, RegisterOnMainchainAddress, approveAmount)
	if err != nil {
		log.Fatal(err)
	}

	registerAmount := new(big.Int).Mul(dec18, big.NewInt(40000))
	authchainGuarantor = SelectAccount(client, 7)
	tx, err = instanceChainRegistrar.RegisterSubchain(authchainGuarantor, subchainID, GovernanceTokenAddress, registerAmount, "111111")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	x, _ := instanceChainRegistrar.GetAllSubchainIDs(nil)
	fmt.Println(x)

}
func oneAcoountStake() {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	subchainID := big.NewInt(360777)
	instanceWrappedTheta, err := ct.NewMockWrappedTheta(WthetaAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	instanceGovernanceToken, err := ct.NewSubchainGovernanceToken(GovernanceTokenAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	validator1 := accountList[1].fromAddress
	validatorCollateralManagerAddr, _ := instanceChainRegistrar.Vcm(nil)
	validatorStakeManagerAddr, _ := instanceChainRegistrar.Vsm(nil)

	//Deposit wTHETA collateral to validators
	validatorCollateral := new(big.Int).Mul(dec18, big.NewInt(2000))
	authValidator1 := SelectAccount(client, 1) //Validator1

	tx, err := instanceWrappedTheta.Mint(authValidator1, validator1, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	authValidator1 = SelectAccount(client, 1) //Validator1
	tx, err = instanceWrappedTheta.Approve(authValidator1, validatorCollateralManagerAddr, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("validator1 collateral")
	// fmt.Println(instanceWrappedTheta.BalanceOf(nil, accountList[1].fromAddress))
	authValidator1 = SelectAccount(client, 1) //Validator1
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
	authGovTokenInitDistrWallet := SelectAccount(client, 6)
	tx, err = instanceGovernanceToken.Transfer(authGovTokenInitDistrWallet, validator1, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("validator1 deposit")
	// fmt.Println(instanceGovernanceToken.BalanceOf(nil, validator1))
	authValidator1 = SelectAccount(client, 1) //Validator1
	tx, err = instanceGovernanceToken.Approve(authValidator1, validatorStakeManagerAddr, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("validator1 allowance")
	// fmt.Println(instanceGovernanceToken.Allowance(nil, accountList[1].fromAddress, validatorStakeManagerAddr))
	// fmt.Println("wallet deposited")
	// fmt.Println(instanceGovernanceToken.BalanceOf(nil, validatorStakeManagerAddr))

	authValidator1 = SelectAccount(client, 1) //Validator1
	tx, err = instanceChainRegistrar.DepositStake(authValidator1, subchainID, validator1, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("validator2 deposited")
	// fmt.Println(instanceGovernanceToken.BalanceOf(nil, accountList[2].fromAddress))

	//height1 := big.NewInt(int64(height))
	fmt.Println(tx)

	time.Sleep(5 * time.Second)
	// height, _ := client.BlockNumber(context.Background())
	// fmt.Println(big.NewInt(int64(height)))
	// dynasty := big.NewInt(int64(height/100 + 1))
	// tx1, tx2 := instanceChainRegistrar.GetValidatorSet(nil, subchainID, dynasty)
	// fmt.Println(tx1)
	// fmt.Println(tx2)
	tx3, err := instanceChainRegistrar.GetStakeSnapshotHeights(nil, subchainID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx3)
}
