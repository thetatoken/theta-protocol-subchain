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
	subchainID = big.NewInt(360777)

	wthetaAddress = common.HexToAddress("0x7d73424a8256C0b2BA245e5d5a3De8820E45F390")
	registerOnMainchainAddress = common.HexToAddress("0x08425D9Df219f93d5763c3e85204cb5B4cE33aAa")
	governanceTokenAddress = common.HexToAddress("0x6E05f58eEddA592f34DD9105b1827f252c509De0")
	tnt20VoucherContractAddress = common.HexToAddress("0x59AF421cB35fc23aB6C8ee42743e6176040031f4")
	tnt20TokenBankAddress = common.HexToAddress("0x2Ce636d6240f8955d085a896e12429f8B3c7db26")
	subchaintnt20TokenBankAddress = common.HexToAddress("0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D")

	tnt721TokenBankAddress = common.HexToAddress("0xEd8d61f42dC1E56aE992D333A4992C3796b22A74")
	tnt721VoucherContractAddress = common.HexToAddress("0x47eb28D8139A188C5686EedE1E9D8EDE3Afdd543")
	Subchaintnt721TokenBankAddress = common.HexToAddress("0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA")

	tfuelTokenbankAddress = common.HexToAddress("0x7f1C87Bd3a22159b8a2E5D195B1a3283D10ea895")
	subchainTfuelTokenBank = common.HexToAddress("0x5a443704dd4B594B382c22a083e2BD3090A6feF3")

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

	map1 = append(map1, "a249a82c42a282e87b2ddef63404d9dfcf6ea501dcaf5d447761765bd74f666d") //10
	map1 = append(map1, "d0d53ac0b4cd47d0ce0060dddc179d04145fea2ee2e0b66c3ee1699c6b492013") //11
	map1 = append(map1, "83f0bb8655139cef4657f90db64a7bb57847038a9bd0ccd87c9b0828e9cbf76d")

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
func OneAccountRegister() {
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
func OneAcoountStake(id int) {
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
	tx1, _ := instanceChainRegistrar.GetValidatorSet(nil, subchainID, dynasty)
	fmt.Println("set info", tx1)
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
