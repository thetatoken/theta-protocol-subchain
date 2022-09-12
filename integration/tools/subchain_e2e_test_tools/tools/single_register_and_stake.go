package tools

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
	scom "github.com/thetatoken/thetasubchain/common"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/thetasubchain/eth/ethclient"

	// rg "chainRegistrarOnMainchain" // for demo
	ct "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

type accounts struct {
	priKey      string
	privateKey  *ecdsa.PrivateKey
	fromAddress common.Address
}

var accountList []accounts

var wthetaAddress common.Address
var registrarOnMainchainAddress common.Address
var governanceTokenAddress common.Address

var mainchainTFuelTokenBankAddress common.Address
var mainchainTNT20TokenBankAddress common.Address
var mainchainTNT721TokenBankAddress common.Address
var mainchainTNT1155TokenBankAddress common.Address

var subchainTFuelTokenBankAddress common.Address
var subchainTNT20TokenBankAddress common.Address
var subchainTNT721TokenBankAddress common.Address
var subchainTNT1155TokenBankAddress common.Address

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

func DeployTokens() {
	// fmt.Printf("Deploying Tokens to the mainchain...\n\n")
	// mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// deployMockTNT20(mainchainClient)
	// deployMockTNT721(mainchainClient)
	// deployMockTNT1155(mainchainClient)
	// fmt.Println("")

	fmt.Printf("Deploying Tokens to the subchain...\n\n")
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	deployMockTNT20(subchainClient)
	deployMockTNT721(subchainClient)
	deployMockTNT1155(subchainClient)
}

func deployMockTNT20(ethClient *ethclient.Client) common.Address {
	fmt.Printf("Deploying mock TNT20 token...\n")
	auth := subchainSelectAccount(ethClient, 1)
	address, _, _, err := ct.DeployMockTNT20(auth, ethClient)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mock TNT20 token deployed, Address:", address)
	// fmt.Println("Deployment tx:", tx.Hash().Hex())
	fmt.Println("")
	time.Sleep(6 * time.Second)
	return address
}

func deployMockTNT721(ethClient *ethclient.Client) common.Address {
	fmt.Printf("Deploying mock TNT721 token...\n")
	auth := subchainSelectAccount(ethClient, 1)
	address, _, _, err := ct.DeployMockTNT721(auth, ethClient)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mock TNT721 token deployed, Address:", address)
	// fmt.Println("Deployment tx:", tx.Hash().Hex())
	fmt.Println("")
	time.Sleep(6 * time.Second)
	return address
}

func deployMockTNT1155(ethClient *ethclient.Client) common.Address {
	fmt.Printf("Deploying mock TNT1155 token...\n")
	auth := subchainSelectAccount(ethClient, 1)
	address, _, _, err := ct.DeployMockTNT1155(auth, ethClient)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mock TNT1155 token deployed, Address:", address)
	// fmt.Println("Deployment tx:", tx.Hash().Hex())
	fmt.Println("")
	time.Sleep(6 * time.Second)
	return address
}

// func deploy_contracts() {
// 	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	auth := subchainSelectAccount(subchainClient, 1)
// 	address, tx, _, err := ct.DeployMockTNT20(auth, subchainClient)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("tnt20", address)
// 	fmt.Println(tx.Hash().Hex())

// 	subchainTNT20TokenAddress = address
// 	auth = subchainSelectAccount(subchainClient, 1)
// 	address, tx, _, err = ct.DeployMockTNT721(auth, subchainClient)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("tnt721", address)
// 	fmt.Println(tx.Hash().Hex())

// 	subchainTNT721TokenAddress = address
// 	auth = subchainSelectAccount(subchainClient, 1)
// 	address, tx, _, err = ct.DeployMockTNT1155(auth, subchainClient)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("tnt1155", address)
// 	fmt.Println(tx.Hash().Hex())

// 	subchainTNT1155TokenAddress = address
// }

func init() {
	subchainID = big.NewInt(360777)

	wthetaAddress = common.HexToAddress("0x7d73424a8256C0b2BA245e5d5a3De8820E45F390")
	registrarOnMainchainAddress = common.HexToAddress("0x08425D9Df219f93d5763c3e85204cb5B4cE33aAa")
	governanceTokenAddress = common.HexToAddress("0x59AF421cB35fc23aB6C8ee42743e6176040031f4")

	mainchainTFuelTokenBankAddress = common.HexToAddress("0xA10A3B175F0f2641Cf41912b887F77D8ef34FAe8")
	subchainTFuelTokenBankAddress = common.HexToAddress("0x5a443704dd4B594B382c22a083e2BD3090A6feF3")

	mainchainTNT20TokenBankAddress = common.HexToAddress("0x6E05f58eEddA592f34DD9105b1827f252c509De0")
	subchainTNT20TokenBankAddress = common.HexToAddress("0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D")

	mainchainTNT721TokenBankAddress = common.HexToAddress("0x79EaFd0B5eC8D3f945E6BB2817ed90b046c0d0Af")
	subchainTNT721TokenBankAddress = common.HexToAddress("0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA")

	mainchainTNT1155TokenBankAddress = common.HexToAddress("0x2Ce636d6240f8955d085a896e12429f8B3c7db26")
	subchainTNT1155TokenBankAddress = common.HexToAddress("0x47c5e40890bcE4a473A49D7501808b9633F29782")

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
	map1 = append(map1, "a249a82c42a282e87b2ddef63404d9dfcf6ea501dcaf5d447761765bd74f666d")
	map1 = append(map1, "d0d53ac0b4cd47d0ce0060dddc179d04145fea2ee2e0b66c3ee1699c6b492013")
	map1 = append(map1, "83f0bb8655139cef4657f90db64a7bb57847038a9bd0ccd87c9b0828e9cbf76d")

	// fmt.Println("-------------------------------------------------------- List of Accounts -------------------------------------------------------")
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
		// fmt.Println("Private key:", value, "address:", fromAddress)
		accountList = append(accountList, accounts{priKey: value, privateKey: privateKey, fromAddress: fromAddress})
	}
	// fmt.Println("---------------------------------------------------------------------------------------------------------------------------------")
	// fmt.Println("")
	// deploy_contracts()
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

func OneAccountRegister() {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}

	chainGuarantor := accountList[7].fromAddress

	instanceWrappedTheta, err := ct.NewMockWrappedTheta(wthetaAddress, client)
	if err != nil {
		log.Fatal("Failed to instantiate the wTHETA contract", err)
	}
	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(registrarOnMainchainAddress, client)
	if err != nil {
		log.Fatal("Failed to instantiate the ChainRegistrar contract", err)
	}
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	amount := new(big.Int).Mul(dec18, big.NewInt(200000))

	auth := mainchainSelectAccount(client, 7) //chainGuarantor
	tx, err := instanceWrappedTheta.Mint(auth, chainGuarantor, amount)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Mint wTHETA tx hash (Mainchain):", tx.Hash().Hex())
	approveAmount := new(big.Int).Mul(dec18, big.NewInt(50000))
	authchainGuarantor := mainchainSelectAccount(client, 7)
	wThetaBalanceOfGuarantor, err := instanceWrappedTheta.BalanceOf(nil, chainGuarantor)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wTheta balance of %v: %v\n", chainGuarantor.Hex(), wThetaBalanceOfGuarantor)

	tx, err = instanceWrappedTheta.Approve(authchainGuarantor, registrarOnMainchainAddress, approveAmount)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Approve wTHETA tx hash (Mainchain):", tx.Hash().Hex())

	allChainIDs, _ := instanceChainRegistrar.GetAllSubchainIDs(nil)
	fmt.Printf("All subchain IDs before subchain registration: %v\n", allChainIDs)
	fmt.Printf("Registering subchain %v\n", subchainID)
	collateralAmount := new(big.Int).Mul(dec18, big.NewInt(40000))
	authchainGuarantor = mainchainSelectAccount(client, 7)
	dummyGenesisHash := "0x012345679abcdef"
	tx, err = instanceChainRegistrar.RegisterSubchain(authchainGuarantor, subchainID, governanceTokenAddress, collateralAmount, dummyGenesisHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Register Subchain tx hash (Mainchain):", tx.Hash().Hex())

	time.Sleep(12 * time.Second)
	allChainIDs, _ = instanceChainRegistrar.GetAllSubchainIDs(nil)
	fmt.Printf("Subchain registered, all subchain IDs: %v\n", allChainIDs)
}

func StakeToValidatorFromAccount(id int, validatorAddrStr string) {
	validator := common.HexToAddress(validatorAddrStr)
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	instanceWrappedTheta, err := ct.NewMockWrappedTheta(wthetaAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	instanceGovernanceToken, err := ct.NewSubchainGovernanceToken(governanceTokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(registrarOnMainchainAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	validatorCollateralManagerAddr, _ := instanceChainRegistrar.Vcm(nil)
	validatorStakeManagerAddr, _ := instanceChainRegistrar.Vsm(nil)

	//
	// The guarantor deposits wTHETA collateral for the validator
	//

	fmt.Println("Prepare for validator collateral deposit...")

	validatorCollateral := new(big.Int).Mul(dec18, big.NewInt(2000))
	guarantor := mainchainSelectAccount(client, id)
	tx, err := instanceWrappedTheta.Mint(guarantor, guarantor.From, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	guarantor = mainchainSelectAccount(client, id)
	tx, err = instanceWrappedTheta.Approve(guarantor, validatorCollateralManagerAddr, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}
	guarantor = mainchainSelectAccount(client, id)
	tx, err = instanceChainRegistrar.DepositCollateral(guarantor, subchainID, validator, validatorCollateral)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Prepare for validator stake deposit...")

	//
	// The staker deposits Gov token stakes for the validator
	//

	staker := mainchainSelectAccount(client, id)
	validatorStakingAmount := new(big.Int).Mul(dec18, big.NewInt(100000))
	validatorStakingAmountMint := new(big.Int)
	validatorStakingAmountMint.Mul(validatorStakingAmount, big.NewInt(10))

	authGovTokenInitDistrWallet := mainchainSelectAccount(client, 6)
	tx, err = instanceGovernanceToken.Transfer(authGovTokenInitDistrWallet, staker.From, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(6 * time.Second)

	staker = mainchainSelectAccount(client, id)
	tx, err = instanceGovernanceToken.Approve(staker, validatorStakeManagerAddr, validatorStakingAmountMint)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(12 * time.Second)

	staker = mainchainSelectAccount(client, id)
	minInitFeeRequired := new(big.Int).Mul(dec18, big.NewInt(100000)) // 100,000 TFuel
	staker.Value.Set(minInitFeeRequired)
	tx, err = instanceChainRegistrar.DepositStake(staker, subchainID, validator, validatorStakingAmount)
	staker.Value.Set(common.Big0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deposit %v Wei Gov Tokens as stake to subchain %v validator %v\n", validatorStakingAmount, subchainID, validator)
	fmt.Printf("Stake deposit tx hash (Mainchain): %v\n", tx.Hash().Hex())

	time.Sleep(12 * time.Second)
	mainchainHeight, _ := client.BlockNumber(context.Background())
	dynasty := scom.CalculateDynasty(big.NewInt(int64(mainchainHeight)))
	fmt.Printf("Maichain block height: %v, dynasty: %v\n", mainchainHeight, dynasty)

	valset, _ := instanceChainRegistrar.GetValidatorSet(nil, subchainID, dynasty)
	fmt.Printf("Validator Set for subchain %v during the current dynasty %v: %v\n", subchainID, dynasty, valset)

	nextDynasty := big.NewInt(0).Add(dynasty, big.NewInt(1))
	valsetNext, _ := instanceChainRegistrar.GetValidatorSet(nil, subchainID, nextDynasty)
	fmt.Printf("Validator Set for subchain %v during the next dynasty    %v: %v\n", subchainID, nextDynasty, valsetNext)
}

func claimStake() {
	client, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	instanceChainRegistrar, err := ct.NewChainRegistrarOnMainchain(registrarOnMainchainAddress, client)
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
