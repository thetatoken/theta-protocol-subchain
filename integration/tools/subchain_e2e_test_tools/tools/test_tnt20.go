package tools

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

func MainchainTNT20Lock(lockAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT20 cross-chain transfer...\n")

	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	sender := mainchainSelectAccount(mainchainClient, 1)
	receiver := accountList[1].fromAddress

	mainchainTNT20ContractAddress := common.HexToAddress("0x4fb87c52Bb6D194f78cd4896E3e574028fedBAB9")
	mainchainTNT20Contract, err := ct.NewMockTNT20(mainchainTNT20ContractAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	mainchainTNT20Name, _ := mainchainTNT20Contract.Name(nil)
	mainchainTNT20Symbol, _ := mainchainTNT20Contract.Symbol(nil)
	mainchainTNT20Decimals, _ := mainchainTNT20Contract.Decimals(nil)
	subchainTNT20VoucherAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	instaceSubchainTNT20VoucherContract, err := ct.NewTNT20VoucherContract(subchainTNT20VoucherAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}
	instanceTNT20TokenBank, err := ct.NewTNT20TokenBank(mainchainTNT20TokenBankAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	mintAmount := big.NewInt(1).Mul(lockAmount, big.NewInt(10))
	fmt.Printf("Mainchain TNT20 contract address: %v, Name: %v, Symbol: %v, Decimals: %v\n", mainchainTNT20ContractAddress, mainchainTNT20Name, mainchainTNT20Symbol, mainchainTNT20Decimals)
	fmt.Printf("Minting %v TNT20 tokens (Wei) on the Mainchain to address %v\n", mintAmount, sender.From)
	_, err = mainchainTNT20Contract.Mint(mainchainSelectAccount(mainchainClient, 0), sender.From, mintAmount)
	if err != nil {
		log.Fatal(err)
	}
	_, err = mainchainTNT20Contract.Approve(sender, mainchainTNT20TokenBankAddress, lockAmount)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(12 * time.Second)

	senderTNT20Balance, _ := mainchainTNT20Contract.BalanceOf(nil, sender.From)
	receiverTNT20VoucherBalance, _ := instaceSubchainTNT20VoucherContract.BalanceOf(nil, receiver)
	fmt.Printf("Mainchain sender : %v, TNT20 balance on Mainchain       : %v\n", sender.From, senderTNT20Balance)
	fmt.Printf("Subchain receiver: %v, TNT20 voucher balance on Subchain: %v\n\n", receiver, receiverTNT20VoucherBalance)
	// printTNT20TokenBankLockedAmount(mainchainTNT20TokenBankAddress, mainchainTNT20ContractAddress, subchainID, mainchainClient)

	sender = mainchainSelectAccount(mainchainClient, 1)
	sender.Value.Set(crossChainFee)
	lockTx, err := instanceTNT20TokenBank.LockTokens(sender, subchainID, mainchainTNT20ContractAddress, receiver, lockAmount)
	sender.Value.Set(common.Big0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT20 Token Lock tx hash (Mainchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering %v TNT20 tokens (Wei) from the Mainchain to Subchain %v...\n\n", lockAmount, subchainID)

	fmt.Printf("Start transfer, timestamp      : %v\n", time.Now())
	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	receipt, err := mainchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("lock error")
	}
	fmt.Printf("Token lock confirmed, timestamp: %v\n", time.Now())

	var subchainVoucherAddress common.Address
	for {
		time.Sleep(1 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getSubchainTNT20VoucherMintLogs(int(fromHeight), int(toHeight), subchainTNT20TokenBankAddress, receiver)
		if result != nil {
			subchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Subchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("End transfer, timestamp        : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// printTNT20TokenBankLockedAmount(mainchainTNT20TokenBankAddress, mainchainTNT20ContractAddress, subchainID, mainchainClient)

	senderTNT20Balance, _ = mainchainTNT20Contract.BalanceOf(nil, sender.From)
	instanceSubchainTNT20VoucherContract, _ := ct.NewTNT20VoucherContract(subchainVoucherAddress, subchainClient)
	receiverSubchainTNT20VoucherBalance, _ := instanceSubchainTNT20VoucherContract.BalanceOf(nil, receiver)

	subchainTNT20Name, _ := instanceSubchainTNT20VoucherContract.Name(nil)
	subchainTNT20Symbol, _ := instanceSubchainTNT20VoucherContract.Symbol(nil)
	subchainTNT20Decimals, _ := instanceSubchainTNT20VoucherContract.Decimals(nil)
	fmt.Printf("Subchain TNT20 voucher contract address: %v, Name: %v, Symbol: %v, Decimals: %v\n", subchainVoucherAddress, subchainTNT20Name, subchainTNT20Symbol, subchainTNT20Decimals)
	fmt.Printf("Mainchain sender : %v, TNT20 balance on Mainchain       : %v\n", sender.From, senderTNT20Balance)
	fmt.Printf("Subchain receiver: %v, TNT20 voucher balance on Subchain: %v\n\n", receiver, receiverSubchainTNT20VoucherBalance)
}

func SubchainTNT20Burn(burnAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT20 cross-chain transfer...\n")

	sender := accountList[1].fromAddress
	receiver := accountList[1].fromAddress

	subchainTNT20TokenBank, _ := ct.NewTNT20TokenBank(subchainTNT20TokenBankAddress, subchainClient)

	subchainTNT20VoucherAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	// subchainTNT20VoucherAddress := subchainTNT20TokenAddress
	subchainTNT20VoucherContract, _ := ct.NewTNT20VoucherContract(subchainTNT20VoucherAddress, subchainClient)
	senderSubchainTNT20VoucherBalance, _ := subchainTNT20VoucherContract.BalanceOf(nil, sender)
	subchainTNT20Name, _ := subchainTNT20VoucherContract.Name(nil)
	subchainTNT20Symbol, _ := subchainTNT20VoucherContract.Symbol(nil)
	subchainTNT20Decimals, _ := subchainTNT20VoucherContract.Decimals(nil)

	mainchainTNT20ContractAddress := common.HexToAddress("0x4fb87c52Bb6D194f78cd4896E3e574028fedBAB9")
	mainchainTNT20Contract, _ := ct.NewTNT20VoucherContract(mainchainTNT20ContractAddress, mainchainClient)
	receiverMainchainTNT20TokenBalance, _ := mainchainTNT20Contract.BalanceOf(nil, receiver)

	fmt.Printf("Subchain TNT20 Voucher address: %v, Name: %v, Symbol: %v, Decimals: %v\n", subchainTNT20VoucherAddress, subchainTNT20Name, subchainTNT20Symbol, subchainTNT20Decimals)
	fmt.Printf("Subchain sender   : %v, TNT20 Voucher balance on Subchain: %v\n", sender, senderSubchainTNT20VoucherBalance)
	fmt.Printf("Mainchain receiver: %v, TNT20 Token balance on Mainchin  : %v\n\n", receiver, receiverMainchainTNT20TokenBalance)
	// printTNT20TokenBankLockedAmount(mainchainTNT20TokenBankAddress, mainchainTNT20ContractAddress, subchainID, mainchainClient)

	authUser := subchainSelectAccount(subchainClient, 1)
	subchainTNT20VoucherContract.Approve(authUser, subchainTNT20TokenBankAddress, burnAmount)

	authUser = subchainSelectAccount(subchainClient, 1)
	authUser.Value.Set(crossChainFee)
	burnTx, err := subchainTNT20TokenBank.BurnVouchers(authUser, subchainTNT20VoucherAddress, sender, burnAmount)
	authUser.Value.Set(common.Big0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT20 Voucher Burn tx hash (Subchain): %v\n", burnTx.Hash().Hex())
	fmt.Printf("Burn %v TNT20 Vouchers (Wei) on Subchain %v to recover the authentic tokens on the Mainchain...\n\n", burnAmount, subchainID)

	fmt.Printf("Start transfer, timestamp        : %v\n", time.Now())
	receipt, err := subchainClient.TransactionReceipt(context.Background(), burnTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("lock error")
	}
	fmt.Printf("Voucher burn confirmed, timestamp: %v\n", time.Now())

	for {
		time.Sleep(1 * time.Second)
		updatedBalance, _ := mainchainTNT20Contract.BalanceOf(nil, receiver)
		if receiverMainchainTNT20TokenBalance.Cmp(updatedBalance) != 0 {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion...\n")
	}
	fmt.Printf("End transfer, timestamp          : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// printTNT20TokenBankLockedAmount(mainchainTNT20TokenBankAddress, mainchainTNT20ContractAddress, subchainID, mainchainClient)

	senderSubchainTNT20VoucherBalance, _ = subchainTNT20VoucherContract.BalanceOf(nil, sender)
	receiverMainchainTNT20TokenBalance, _ = mainchainTNT20Contract.BalanceOf(nil, receiver)
	mainchainTNT20Name, _ := mainchainTNT20Contract.Name(nil)
	mainchainTNT20Symbol, _ := mainchainTNT20Contract.Symbol(nil)
	mainchainTNT20Decimals, _ := mainchainTNT20Contract.Decimals(nil)
	fmt.Printf("Mainchain TNT20 token contract address: %v, Name: %v, Symbol: %v, Decimals: %v\n", mainchainTNT20ContractAddress, mainchainTNT20Name, mainchainTNT20Symbol, mainchainTNT20Decimals)
	fmt.Printf("Subchain sender   : %v, TNT20 Voucher balance on Subchain: %v\n", sender, senderSubchainTNT20VoucherBalance)
	fmt.Printf("Mainchain receiver: %v, TNT20 Token balance on Mainchain : %v\n\n", receiver, receiverMainchainTNT20TokenBalance)
}

func SubchainTNT20Lock(lockAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT20 cross-chain transfer...\n")

	subchainTNT20Address := common.HexToAddress("0x9572eCEA04Fe74B642400dBb04952E91049C9B3D")

	sender := accountList[1].fromAddress
	receiver := accountList[6].fromAddress
	subchainTNT20TokenBankInstance, _ := ct.NewTNT20TokenBank(subchainTNT20TokenBankAddress, subchainClient)
	subchainTNT20Instance, _ := ct.NewMockTNT20(subchainTNT20Address, subchainClient)

	mintAmount := big.NewInt(1).Mul(big.NewInt(5), lockAmount)
	fmt.Printf("Minting %v TNT20 tokens\n", mintAmount)

	authUser := subchainSelectAccount(subchainClient, 1)
	subchainTNT20Instance.Mint(authUser, sender, mintAmount)
	time.Sleep(6 * time.Second)

	senderTNT20Balance, _ := subchainTNT20Instance.BalanceOf(nil, sender)
	subchainTNT20Name, _ := subchainTNT20Instance.Name(nil)
	subchainTNT20Symbol, _ := subchainTNT20Instance.Symbol(nil)
	subchainTNT20Decimals, _ := subchainTNT20Instance.Decimals(nil)

	expectedMainchainTNT20VoucherAddress := common.HexToAddress("0x2164892B95EF1eA9C586643a507D0b613215BbD7")
	expectedMainchainTNT20VoucherContractInstance, _ := ct.NewTNT20VoucherContract(expectedMainchainTNT20VoucherAddress, mainchainClient)
	receiverMainchainTNT20VoucherBalance, _ := expectedMainchainTNT20VoucherContractInstance.BalanceOf(nil, receiver)

	fmt.Printf("Subchain TNT20 contract address: %v, Name: %v, Symbol: %v, Decimals: %v\n", subchainTNT20Address, subchainTNT20Name, subchainTNT20Symbol, subchainTNT20Decimals)
	fmt.Printf("Subchain sender   : %v, TNT20 balance on Subchain         : %v\n", sender, senderTNT20Balance)
	fmt.Printf("Mainchain receiver: %v, TNT20 voucher balance on Mainchain: %v\n\n", receiver, receiverMainchainTNT20VoucherBalance)
	// mainchainID, _ := mainchainClient.ChainID(context.Background())
	// printTNT20TokenBankLockedAmount(subchainTNT20TokenBankAddress, subchainTNT20Address, mainchainID, subchainClient)

	authUser = subchainSelectAccount(subchainClient, 1)
	subchainTNT20Instance.Approve(authUser, subchainTNT20TokenBankAddress, lockAmount)

	authUser = subchainSelectAccount(subchainClient, 1)
	authUser.Value.Set(crossChainFee)
	lockTx, err := subchainTNT20TokenBankInstance.LockTokens(authUser, big.NewInt(366), subchainTNT20Address, accountList[6].fromAddress, lockAmount)
	authUser.Value.Set(common.Big0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT20 Token Lock tx hash (Subchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering %v TNT20 tokens (Wei) from to Subchain %v to the Mainchain...\n\n", lockAmount, subchainID)

	fmt.Printf("Start transfer, timestamp      : %v\n", time.Now())
	receipt, err := subchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("lock error")
	}
	fmt.Printf("Token lock confirmed, timestamp: %v\n", time.Now())

	fromHeight, _ := mainchainClient.BlockNumber(context.Background())
	var mainchainVoucherAddress common.Address
	for {
		time.Sleep(1 * time.Second)
		toHeight, _ := mainchainClient.BlockNumber(context.Background())
		result := getMainchainTNT20VoucherMintLogs(int(fromHeight), int(toHeight), mainchainTNT20TokenBankAddress, accountList[6].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Mainchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("End transfer, timestamp        : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// printTNT20TokenBankLockedAmount(subchainTNT20TokenBankAddress, subchainTNT20Address, mainchainID, subchainClient)

	senderTNT20Balance, _ = subchainTNT20Instance.BalanceOf(nil, sender)
	instanceMainchainTNT20VoucherContract, _ := ct.NewTNT20VoucherContract(mainchainVoucherAddress, mainchainClient)
	receiverMainchainTNT20VoucherBalance, _ = instanceMainchainTNT20VoucherContract.BalanceOf(nil, receiver)
	mainchainTNT20Name, _ := instanceMainchainTNT20VoucherContract.Name(nil)
	mainchainTNT20Symbol, _ := instanceMainchainTNT20VoucherContract.Symbol(nil)
	mainchainTNT20Decimals, _ := instanceMainchainTNT20VoucherContract.Decimals(nil)
	fmt.Printf("Mainchain TNT20 voucher contract address: %v, Name: %v, Symbol: %v, Decimals: %v\n", mainchainVoucherAddress, mainchainTNT20Name, mainchainTNT20Symbol, mainchainTNT20Decimals)
	fmt.Printf("Subchain sender   : %v, TNT20 balance on Subchain         : %v\n", sender, senderTNT20Balance)
	fmt.Printf("Mainchain receiver: %v, TNT20 voucher balance on Mainchain: %v\n\n", receiver, receiverMainchainTNT20VoucherBalance)
}

func MainchainTNT20Burn(burnAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT20 cross-chain transfer...\n")

	sender := accountList[6].fromAddress
	receiver := accountList[1].fromAddress

	mainchainTNT20TokenBankInstance, _ := ct.NewTNT20TokenBank(mainchainTNT20TokenBankAddress, mainchainClient)

	mainchainTNT20VoucherAddress := common.HexToAddress("0x2164892B95EF1eA9C586643a507D0b613215BbD7")
	mainchainTNT20VoucherContract, _ := ct.NewTNT20VoucherContract(mainchainTNT20VoucherAddress, mainchainClient)
	senderMainchainTNT20VoucherBalance, _ := mainchainTNT20VoucherContract.BalanceOf(nil, sender)
	mainchainTNT20Name, _ := mainchainTNT20VoucherContract.Name(nil)
	mainchainTNT20Symbol, _ := mainchainTNT20VoucherContract.Symbol(nil)
	mainchainTNT20Decimals, _ := mainchainTNT20VoucherContract.Decimals(nil)

	subchainTNT20TokenAddress := common.HexToAddress("0x9572eCEA04Fe74B642400dBb04952E91049C9B3D")
	// subchainTNT20TokenAddress := deployMockTNT20(subchainClient)
	subchainTNT20TokenContract, _ := ct.NewTNT20VoucherContract(subchainTNT20TokenAddress, subchainClient)
	receiverSubchainTNT20TokenBalance, _ := subchainTNT20TokenContract.BalanceOf(nil, receiver)

	fmt.Printf("Mainchain TNT20 Voucher address: %v, Name: %v, Symbol: %v, Decimals: %v\n", mainchainTNT20VoucherAddress, mainchainTNT20Name, mainchainTNT20Symbol, mainchainTNT20Decimals)
	fmt.Printf("Mainchain sender : %v, TNT20 Voucher balance on Mainchain: %v\n", sender, senderMainchainTNT20VoucherBalance)
	fmt.Printf("Subchain receiver: %v, TNT20 Token balance on Subchain   : %v\n\n", receiver, receiverSubchainTNT20TokenBalance)
	// mainchainID, _ := mainchainClient.ChainID(context.Background())
	// printTNT20TokenBankLockedAmount(subchainTNT20TokenBankAddress, subchainTNT20TokenAddress, mainchainID, subchainClient)

	authUser := mainchainSelectAccount(mainchainClient, 6)
	mainchainTNT20VoucherContract.Approve(authUser, mainchainTNT20TokenBankAddress, burnAmount)

	authUser = mainchainSelectAccount(mainchainClient, 6)
	authUser.Value.Set(crossChainFee)
	burnTx, err := mainchainTNT20TokenBankInstance.BurnVouchers(authUser, mainchainTNT20VoucherAddress, accountList[1].fromAddress, burnAmount)
	authUser.Value.Set(common.Big0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT20 Voucher Burn tx hash (Mainchain): %v\n", burnTx.Hash().Hex())
	fmt.Printf("Burn %v TNT20 Vouchers (Wei) on the Mainchain to recover the authentic tokens on Subchain %v...\n\n", burnAmount, subchainID)

	fmt.Printf("Start transfer, timestamp        : %v\n", time.Now())
	receipt, err := mainchainClient.TransactionReceipt(context.Background(), burnTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("lock error")
	}
	fmt.Printf("Voucher burn confirmed, timestamp: %v\n", time.Now())

	for {
		time.Sleep(1 * time.Second)
		updatedBalance, _ := subchainTNT20TokenContract.BalanceOf(nil, receiver)
		if receiverSubchainTNT20TokenBalance.Cmp(updatedBalance) != 0 {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion...\n")
	}
	fmt.Printf("End transfer, timestamp          : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// printTNT20TokenBankLockedAmount(subchainTNT20TokenBankAddress, subchainTNT20TokenAddress, mainchainID, subchainClient)

	senderMainchainTNT20VoucherBalance, _ = mainchainTNT20VoucherContract.BalanceOf(nil, sender)
	receiverSubchainTNT20TokenBalance, _ = subchainTNT20TokenContract.BalanceOf(nil, receiver)
	subchainTNT20Name, _ := subchainTNT20TokenContract.Name(nil)
	subchainTNT20Symbol, _ := subchainTNT20TokenContract.Symbol(nil)
	subchainTNT20Decimals, _ := subchainTNT20TokenContract.Decimals(nil)

	fmt.Printf("Subchain TNT20 token contract address: %v, Name: %v, Symbol: %v, Decimals: %v\n", subchainTNT20TokenAddress, subchainTNT20Name, subchainTNT20Symbol, subchainTNT20Decimals)
	fmt.Printf("Mainchain sender : %v, TNT20 Voucher balance on Mainchain: %v\n", sender, senderMainchainTNT20VoucherBalance)
	fmt.Printf("Subchain receiver: %v, TNT20 Token balance on Subchain   : %v\n\n", receiver, receiverSubchainTNT20TokenBalance)
}

func QueryTNT20(chainID int64, contractAddress, accountAddress string) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	tnt20ContractAddress := common.HexToAddress(contractAddress) // FIXME: should instantiate a mock TNT20 instead of using the Voucher contract (which causes confusion)
	var tnt20Contract *ct.MockTNT20
	if chainID == 366 {
		fmt.Printf("Preparing for TNT20 mainchain query...\n")
		tnt20Contract, err = ct.NewMockTNT20(tnt20ContractAddress, mainchainClient)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Preparing for TNT20 subchain query...\n")
		tnt20Contract, err = ct.NewMockTNT20(tnt20ContractAddress, subchainClient)
		if err != nil {
			log.Fatal(err)
		}
	}
	balance, _ := tnt20Contract.BalanceOf(nil, common.HexToAddress(accountAddress))
	fmt.Println("Account ", accountAddress, " TNT20 balance in ", contractAddress, " is ", balance)

}

// func printTNT20TokenBankLockedAmount(tokenBankAddr common.Address, tokenAddr common.Address, chainID *big.Int, ethclient *ethclient.Client) {
// 	tnt20TokenBank, _ := ct.NewTNT20TokenBank(tokenBankAddr, ethclient)
// 	tokenBankLockedAmount, _ := tnt20TokenBank.TotalLockedAmounts(nil, chainID, tokenAddr)
// 	fmt.Printf("TNT20TokenBank locked amount: %v\n\n", tokenBankLockedAmount)
// }
