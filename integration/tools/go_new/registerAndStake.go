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
func main1() {
	accountsInit()
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	WthetaAddress := common.HexToAddress("0x17deB845AAcCA09873D1984784EbdC59DBB38846")
	RegisterOnMainchainAddress := common.HexToAddress("0xDE06587B53D752ac280c7Dae1D11C3ffc341A509")
	GovernanceTokenAddress := common.HexToAddress("0xD323FbDE9D287A288d22F9F4B30E61d2500aa962")
	//TNT20TokenBankAddress := common.HexToAddress("0x1f629139b3b4A03799c6e6655b7F59a1F01598E7")

	clientNew, auth := chooseAccount(client, 7) //chainGuarantor
	instanceWrappedTheta, err := ct.NewMockWrappedTheta(WthetaAddress, clientNew)
	if err != nil {
		log.Fatal("hhh", err)
	}
	chainGuarantor := accountList[7].fromAddress
	num := big.NewInt(200000)
	tx, err := instanceWrappedTheta.Mint(auth, chainGuarantor, num)
	if err != nil {
		log.Fatal(err)
	}
	num = big.NewInt(50000)
	clientNew, auth = chooseAccount(client, 7)
	tx, err = instanceWrappedTheta.Approve(auth, RegisterOnMainchainAddress, num)
	if err != nil {
		log.Fatal(err)
	}
	subchainID := big.NewInt(9988)
	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, clientNew)
	if err != nil {
		log.Fatal(err)
	}
	clientNew, auth = chooseAccount(client, 7)
	tx, err = instanceChainRegistrar.Registersubchain(auth, subchainID, GovernanceTokenAddress, num, "111")
	if err != nil {
		log.Fatal(err)
	}
	validatorCollateralManagerAddr, _ := instanceChainRegistrar.Vcm(nil)
	validatorStakeManagerAddr, _ := instanceChainRegistrar.Vsm(nil)

	//Deposit wTHETA collateral to validators
	validatorCollateral := big.NewInt(2000)
	//Validator1
	clientValidator1, authValidator1 := chooseAccount(client, 1) //Validator1
	instanceWrappedThetaValidator1, err := ct.NewMockWrappedTheta(WthetaAddress, clientValidator1)
	if err != nil {
		log.Fatal("hhh", err)
	}
	validator1 := accountList[1].fromAddress
	num = big.NewInt(200000)
	tx, err = instanceWrappedThetaValidator1.Mint(authValidator1, validator1, num)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator1, authValidator1 = chooseAccount(client, 1) //Validator1
	tx, err = instanceWrappedThetaValidator1.Approve(authValidator1, validatorCollateralManagerAddr, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator1, authValidator1 = chooseAccount(client, 1) //Validator1
	instanceChainRegistrarValidator1, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, clientValidator1)
	tx, err = instanceChainRegistrarValidator1.Depositcollateral(authValidator1, subchainID, validator1, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}

	//Validator2
	clientValidator2, authValidator2 := chooseAccount(client, 2) //Validator2
	instanceWrappedThetaValidator2, err := ct.NewMockWrappedTheta(WthetaAddress, clientValidator2)
	if err != nil {
		log.Fatal("hhh", err)
	}
	validator2 := accountList[2].fromAddress
	num = big.NewInt(200000)
	tx, err = instanceWrappedThetaValidator2.Mint(authValidator2, validator2, num)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator2, authValidator2 = chooseAccount(client, 2) //Validator2
	tx, err = instanceWrappedThetaValidator2.Approve(authValidator2, validatorCollateralManagerAddr, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator2, authValidator2 = chooseAccount(client, 2) //Validator2
	instanceChainRegistrarValidator2, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, clientValidator2)
	tx, err = instanceChainRegistrarValidator2.Depositcollateral(authValidator2, subchainID, validator2, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}

	//Validator3
	clientValidator3, authValidator3 := chooseAccount(client, 3) //Validator3
	instanceWrappedThetaValidator3, err := ct.NewMockWrappedTheta(WthetaAddress, clientValidator3)
	if err != nil {
		log.Fatal("hhh", err)
	}
	validator3 := accountList[3].fromAddress
	num = big.NewInt(200000)
	tx, err = instanceWrappedThetaValidator3.Mint(authValidator3, validator3, num)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator3, authValidator3 = chooseAccount(client, 3) //Validator3
	tx, err = instanceWrappedThetaValidator3.Approve(authValidator3, validatorCollateralManagerAddr, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator3, authValidator3 = chooseAccount(client, 3) //Validator3
	instanceChainRegistrarValidator3, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, clientValidator3)
	tx, err = instanceChainRegistrarValidator3.Depositcollateral(authValidator3, subchainID, validator3, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}

	//Stake to the validators
	validatorStakingAmount := big.NewInt(100000)
	//govTokenInitDistrWalletAddress:=accountList[6].fromAddress
	ClientGovTokenInitDistrWallet, authGovTokenInitDistrWallet := chooseAccount(client, 6)
	instanceGovernanceTokengovTokenInitDistrWallet, err := ct.NewSubchainGovernanceToken(GovernanceTokenAddress, ClientGovTokenInitDistrWallet)
	tx, err = instanceGovernanceTokengovTokenInitDistrWallet.Transfer(authGovTokenInitDistrWallet, validator1, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}
	ClientGovTokenInitDistrWallet, authGovTokenInitDistrWallet = chooseAccount(client, 6)
	tx, err = instanceGovernanceTokengovTokenInitDistrWallet.Transfer(authGovTokenInitDistrWallet, validator2, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}
	ClientGovTokenInitDistrWallet, authGovTokenInitDistrWallet = chooseAccount(client, 6)
	tx, err = instanceGovernanceTokengovTokenInitDistrWallet.Transfer(authGovTokenInitDistrWallet, validator3, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}

	clientValidator1, authValidator1 = chooseAccount(client, 1) //Validator1
	instanceGovernanceTokengovValidator1, err := ct.NewSubchainGovernanceToken(GovernanceTokenAddress, clientValidator1)
	tx, err = instanceGovernanceTokengovValidator1.Approve(authValidator1, validatorStakeManagerAddr, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator2, authValidator2 = chooseAccount(client, 2) //Validator2
	instanceGovernanceTokengovValidator2, err := ct.NewSubchainGovernanceToken(GovernanceTokenAddress, clientValidator2)
	tx, err = instanceGovernanceTokengovValidator2.Approve(authValidator2, validatorStakeManagerAddr, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator3, authValidator3 = chooseAccount(client, 3) //Validator3
	instanceGovernanceTokengovValidator3, err := ct.NewSubchainGovernanceToken(GovernanceTokenAddress, clientValidator3)
	tx, err = instanceGovernanceTokengovValidator3.Approve(authValidator3, validatorStakeManagerAddr, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator1, authValidator1 = chooseAccount(client, 1) //Validator1
	tx, err = instanceChainRegistrarValidator1.Depositstake(authValidator1, subchainID, validator1, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator2, authValidator2 = chooseAccount(client, 2) //Validator2
	tx, err = instanceChainRegistrarValidator2.Depositstake(authValidator2, subchainID, validator2, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}
	clientValidator3, authValidator3 = chooseAccount(client, 3) //Validator3
	tx, err = instanceChainRegistrarValidator2.Depositstake(authValidator3, subchainID, validator3, validatorStakingAmount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx)
}
