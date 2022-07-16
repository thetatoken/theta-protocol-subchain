package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	// rg "chainRegistrarOnMainchain" // for demo
	ct "github.com/thetatoken/thetasubchain/integration/tools/go_new/contract"
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
		//fmt.Println(value, "-----", fromAddress)
		accountList = append(accountList, accounts{priKey: value, privateKey: privateKey, fromAddress: fromAddress})
	}

}
func selectAccount(client *ethclient.Client, id int) *bind.TransactOpts {
	time.Sleep(1 * time.Second)
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

	return auth
}
func registerAndStake() {
	accountsInit()
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	subchainID := big.NewInt(120)

	WthetaAddress := common.HexToAddress("0xD46439c8D157Df0fe214DaD9e73Bd7920654ecA0")
	RegisterOnMainchainAddress := common.HexToAddress("0x1EaE66323eAeb16dAD8e162520251D49D4E2f6a3")
	GovernanceTokenAddress := common.HexToAddress("0x6c1d3aBA75Fea74000a10BbfbE86be30f20a2F3D")
	//TNT20TokenBankAddress := common.HexToAddress("0x1f629139b3b4A03799c6e6655b7F59a1F01598E7")
	validator1 := accountList[1].fromAddress
	validator2 := accountList[2].fromAddress
	validator3 := accountList[3].fromAddress

	auth := selectAccount(client, 7) //chainGuarantor
	instanceWrappedTheta, err := ct.NewMockWrappedTheta(WthetaAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	amount := new(big.Int).Mul(dec18, big.NewInt(200000))
	chainGuarantor := accountList[7].fromAddress

	tx, err := instanceWrappedTheta.Mint(auth, chainGuarantor, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	num := new(big.Int).Mul(dec18, big.NewInt(50000))
	authchainGuarantor := selectAccount(client, 7)
	fmt.Println(instanceWrappedTheta.BalanceOf(nil, chainGuarantor))
	tx, err = instanceWrappedTheta.Approve(authchainGuarantor, RegisterOnMainchainAddress, num)
	if err != nil {
		log.Fatal(err)
	}

	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	num1 := new(big.Int).Mul(dec18, big.NewInt(40000))
	authchainGuarantor = selectAccount(client, 7)
	tx, err = instanceChainRegistrar.RegisterSubchain(authchainGuarantor, subchainID, GovernanceTokenAddress, num1, "111111")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Second)
	x, _ := instanceChainRegistrar.GetAllSubchainIDs(nil)
	fmt.Println(x)

	validatorCollateralManagerAddr, _ := instanceChainRegistrar.Vcm(nil)
	validatorStakeManagerAddr, _ := instanceChainRegistrar.Vsm(nil)

	//Deposit wTHETA collateral to validators
	validatorCollateral := new(big.Int).Mul(dec18, big.NewInt(2000))
	//Validator1
	authValidator1 := selectAccount(client, 1) //Validator1
	instanceWrappedThetaValidator1, err := ct.NewMockWrappedTheta(WthetaAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}

	tx, err = instanceWrappedThetaValidator1.Mint(authValidator1, validator1, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	authValidator1 = selectAccount(client, 1) //Validator1
	tx, err = instanceWrappedThetaValidator1.Approve(authValidator1, validatorCollateralManagerAddr, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("validator1 collateral")
	fmt.Println(instanceWrappedThetaValidator1.BalanceOf(nil, accountList[1].fromAddress))
	authValidator1 = selectAccount(client, 1) //Validator1
	instanceChainRegistrarValidator1, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, client)
	tx, err = instanceChainRegistrarValidator1.DepositCollateral(authValidator1, subchainID, validator1, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}

	//Validator2
	authValidator2 := selectAccount(client, 2) //Validator2
	// fmt.Println("validator1 collateral")
	// fmt.Println(instanceWrappedThetaValidator1.BalanceOf(nil, accountList[1].fromAddress))
	// fmt.Println("validator2 collateral")
	// fmt.Println(instanceWrappedThetaValidator1.BalanceOf(nil, accountList[2].fromAddress))
	instanceWrappedThetaValidator2, err := ct.NewMockWrappedTheta(WthetaAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}

	tx, err = instanceWrappedThetaValidator2.Mint(authValidator2, validator2, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	authValidator2 = selectAccount(client, 2) //Validator2
	tx, err = instanceWrappedThetaValidator2.Approve(authValidator2, validatorCollateralManagerAddr, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("validator2 collateral")
	fmt.Println(instanceWrappedThetaValidator1.BalanceOf(nil, accountList[2].fromAddress))
	authValidator2 = selectAccount(client, 2) //Validator2
	instanceChainRegistrarValidator2, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, client)
	tx, err = instanceChainRegistrarValidator2.DepositCollateral(authValidator2, subchainID, validator2, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}

	//Validator3
	authValidator3 := selectAccount(client, 3) //Validator3
	instanceWrappedThetaValidator3, err := ct.NewMockWrappedTheta(WthetaAddress, client)
	if err != nil {
		log.Fatal("hhh", err)
	}

	tx, err = instanceWrappedThetaValidator3.Mint(authValidator3, validator3, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	authValidator3 = selectAccount(client, 3) //Validator3
	tx, err = instanceWrappedThetaValidator3.Approve(authValidator3, validatorCollateralManagerAddr, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	authValidator3 = selectAccount(client, 3) //Validator3
	instanceChainRegistrarValidator3, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, client)
	tx, err = instanceChainRegistrarValidator3.DepositCollateral(authValidator3, subchainID, validator3, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}

	//Stake to the validators
	validatorStakingAmount := new(big.Int).Mul(dec18, big.NewInt(100000))
	validatorStakingAmountMint := new(big.Int)
	validatorStakingAmountMint.Mul(validatorStakingAmount, big.NewInt(10))
	//govTokenInitDistrWalletAddress:=accountList[6].fromAddress
	authGovTokenInitDistrWallet := selectAccount(client, 6)
	instanceGovernanceTokengovTokenInitDistrWallet, err := ct.NewSubchainGovernanceToken(GovernanceTokenAddress, client)
	fmt.Println("validator1 deposit")
	fmt.Println(instanceGovernanceTokengovTokenInitDistrWallet.BalanceOf(nil, validator1))
	tx, err = instanceGovernanceTokengovTokenInitDistrWallet.Transfer(authGovTokenInitDistrWallet, validator1, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}
	authGovTokenInitDistrWallet = selectAccount(client, 6)
	tx, err = instanceGovernanceTokengovTokenInitDistrWallet.Transfer(authGovTokenInitDistrWallet, validator2, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}
	authGovTokenInitDistrWallet = selectAccount(client, 6)
	tx, err = instanceGovernanceTokengovTokenInitDistrWallet.Transfer(authGovTokenInitDistrWallet, validator3, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("validator1 deposit")
	fmt.Println(instanceGovernanceTokengovTokenInitDistrWallet.BalanceOf(nil, validator1))
	authValidator1 = selectAccount(client, 1) //Validator1
	instanceGovernanceTokengovValidator1, err := ct.NewSubchainGovernanceToken(GovernanceTokenAddress, client)
	tx, err = instanceGovernanceTokengovValidator1.Approve(authValidator1, validatorStakeManagerAddr, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}
	authValidator2 = selectAccount(client, 2) //Validator2
	instanceGovernanceTokengovValidator2, err := ct.NewSubchainGovernanceToken(GovernanceTokenAddress, client)
	tx, err = instanceGovernanceTokengovValidator2.Approve(authValidator2, validatorStakeManagerAddr, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}
	authValidator3 = selectAccount(client, 3) //Validator3
	instanceGovernanceTokengovValidator3, err := ct.NewSubchainGovernanceToken(GovernanceTokenAddress, client)
	tx, err = instanceGovernanceTokengovValidator3.Approve(authValidator3, validatorStakeManagerAddr, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("validator1 allowance")
	fmt.Println(instanceGovernanceTokengovTokenInitDistrWallet.Allowance(nil, accountList[1].fromAddress, validatorStakeManagerAddr))
	fmt.Println("wallet deposited")
	fmt.Println(instanceGovernanceTokengovTokenInitDistrWallet.BalanceOf(nil, validatorStakeManagerAddr))

	authValidator1 = selectAccount(client, 1) //Validator1
	tx, err = instanceChainRegistrarValidator1.DepositStake(authValidator1, subchainID, validator1, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("validator2 deposited")
	fmt.Println(instanceGovernanceTokengovTokenInitDistrWallet.BalanceOf(nil, accountList[2].fromAddress))

	//height1 := big.NewInt(int64(height))
	authValidator2 = selectAccount(client, 2) //Validator2
	tx, err = instanceChainRegistrarValidator2.DepositStake(authValidator2, subchainID, validator2, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}

	authValidator3 = selectAccount(client, 3) //Validator3
	tx, err = instanceChainRegistrarValidator3.DepositStake(authValidator3, subchainID, validator3, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx)

	time.Sleep(5 * time.Second)
	height, _ := client.BlockNumber(context.Background())
	fmt.Println(big.NewInt(int64(height)))
	dynasty := big.NewInt(int64(height/100 + 1))
	tx1, tx2 := instanceChainRegistrarValidator2.GetValidatorSet(nil, subchainID, dynasty)
	fmt.Println(tx1)
	fmt.Println(tx2)
	tx3, err := instanceChainRegistrarValidator3.GetStakeSnapshotHeights(nil, subchainID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx3)
}
