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

	tnt20ContractAddress := tnt20VoucherContractAddress // FIXME: should instantiate a mock TNT20 instead of using the Voucher contract (which causes confusion)
	instaceTNT20Contract, err := ct.NewTNT20VoucherContract(tnt20ContractAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT20VoucherAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	instaceSubchainTNT20VoucherContract, err := ct.NewTNT20VoucherContract(subchainTNT20VoucherAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}
	instanceTNT20TokenBank, err := ct.NewTNT20TokenBank(tnt20TokenBankAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	mintAmount := big.NewInt(1).Mul(lockAmount, big.NewInt(10))
	fmt.Printf("Mainchain TNT20 contract address: %v\n", tnt20ContractAddress)
	fmt.Printf("Minting %v TNT20 tokens (Wei) on the Mainchain to address %v\n", mintAmount, sender.From)
	_, err = instaceTNT20Contract.Mint(mainchainSelectAccount(mainchainClient, 0), sender.From, mintAmount)
	if err != nil {
		log.Fatal(err)
	}
	_, err = instaceTNT20Contract.Approve(sender, tnt20TokenBankAddress, lockAmount)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(12 * time.Second)

	senderTNT20Balance, _ := instaceTNT20Contract.BalanceOf(nil, sender.From)
	receiverTNT20VoucherBalance, _ := instaceSubchainTNT20VoucherContract.BalanceOf(nil, receiver)
	fmt.Printf("Mainchain sender : %v, TNT20 balance on Mainchain       : %v\n", sender.From, senderTNT20Balance)
	fmt.Printf("Subchain receiver: %v, TNT20 voucher balance on Subchain: %v\n\n", receiver, receiverTNT20VoucherBalance)

	sender = mainchainSelectAccount(mainchainClient, 1)
	sender.Value.Set(crossChainFee)
	lockTx, err := instanceTNT20TokenBank.LockTokens(sender, subchainID, tnt20VoucherContractAddress, receiver, lockAmount)
	if err != nil {
		log.Fatal(err)
	}
	sender.Value.Set(common.Big0)

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

	senderTNT20Balance, _ = instaceTNT20Contract.BalanceOf(nil, sender.From)
	instanceSubchainTNT20VoucherContract, _ := ct.NewTNT20VoucherContract(subchainVoucherAddress, subchainClient)
	receiverSubchainTNT20VoucherBalance, _ := instanceSubchainTNT20VoucherContract.BalanceOf(nil, receiver)

	fmt.Printf("Subchain TNT20 voucher contract address: %v\n", subchainVoucherAddress)
	fmt.Printf("Mainchain sender : %v, TNT20 balance on Mainchain       : %v\n", sender.From, senderTNT20Balance)
	fmt.Printf("Subchain receiver: %v, TNT20 voucher balance on Subchain: %v\n\n", receiver, receiverSubchainTNT20VoucherBalance)
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

	subchainTNT20Address := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f")
	if err != nil {
		log.Fatal(err)
	}
	sender := accountList[1].fromAddress
	receiver := accountList[6].fromAddress
	subchainTNT20TokenBankInstance, _ := ct.NewTNT20TokenBank(subchainTNT20TokenBankAddress, subchainClient)
	subchainTNT20Instance, _ := ct.NewTNT20VoucherContract(subchainTNT20Address, subchainClient)
	senderTNT20Balance, _ := subchainTNT20Instance.BalanceOf(nil, sender)

	expectedMainchainTNT20VoucherAddress := common.HexToAddress("0xb0DBBcba1Be5B71Dcb42aB1935773B3675e645e8")
	expectedMainchainTNT20VoucherContractInstance, _ := ct.NewTNT20VoucherContract(expectedMainchainTNT20VoucherAddress, mainchainClient)
	receiverMainchainTNT20VoucherBalance, _ := expectedMainchainTNT20VoucherContractInstance.BalanceOf(nil, receiver)

	fmt.Printf("Subchain TNT20 contract address: %v\n", subchainTNT20Address)
	fmt.Printf("Subchain sender   : %v, TNT20 balance on Subchain         : %v\n", sender, senderTNT20Balance)
	fmt.Printf("Mainchain receiver: %v, TNT20 voucher balance on Mainchain: %v\n\n", receiver, receiverMainchainTNT20VoucherBalance)

	authUser := subchainSelectAccount(subchainClient, 1)
	subchainTNT20Instance.Approve(authUser, subchainTNT20TokenBankAddress, lockAmount)

	authUser = subchainSelectAccount(subchainClient, 1)
	lockTx, err := subchainTNT20TokenBankInstance.LockTokens(authUser, big.NewInt(366), subchainTNT20Address, accountList[6].fromAddress, lockAmount)
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
		result := getMainchainTNT20VoucherMintLogs(int(fromHeight), int(toHeight), tnt20TokenBankAddress, accountList[6].fromAddress)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Mainchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("End transfer, timestamp        : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	senderTNT20Balance, _ = subchainTNT20Instance.BalanceOf(nil, sender)
	instanceMainchainTNT20VoucherContract, _ := ct.NewTNT20VoucherContract(mainchainVoucherAddress, mainchainClient)
	receiverMainchainTNT20VoucherBalance, _ = instanceMainchainTNT20VoucherContract.BalanceOf(nil, receiver)

	fmt.Printf("Mainchain TNT20 voucher contract address: %v\n", mainchainVoucherAddress)
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

	mainchainTNT20TokenBankInstance, _ := ct.NewTNT20TokenBank(tnt20TokenBankAddress, mainchainClient)

	mainchainTNT20VoucherAddress := common.HexToAddress("0xb0DBBcba1Be5B71Dcb42aB1935773B3675e645e8")
	mainchainTNT20VoucherContract, _ := ct.NewTNT20VoucherContract(mainchainTNT20VoucherAddress, mainchainClient)
	senderMainchainTNT20VoucherBalance, _ := mainchainTNT20VoucherContract.BalanceOf(nil, sender)

	subchainTNT20TokenAddress := common.HexToAddress("0x7D7e270b7E279C94b265A535CdbC00Eb62E6e68f") // FIXME: should instantiate a mock TNT20 instead of using the Voucher contract (which causes confusion)
	subchainTNT20TokenContract, _ := ct.NewTNT20VoucherContract(subchainTNT20TokenAddress, subchainClient)
	receiverSubchainTNT20TokenBalance, _ := subchainTNT20TokenContract.BalanceOf(nil, receiver)

	fmt.Printf("Mainchain TNT20 Voucher address: %v\n", mainchainTNT20VoucherAddress)
	fmt.Printf("Mainchain sender : %v, TNT20 Voucher balance on Mainchain: %v\n", sender, senderMainchainTNT20VoucherBalance)
	fmt.Printf("Subchain receiver: %v, TNT20 Token balance on Subchain   : %v\n\n", receiver, receiverSubchainTNT20TokenBalance)

	authUser := mainchainSelectAccount(mainchainClient, 6)
	mainchainTNT20VoucherContract.Approve(authUser, tnt20TokenBankAddress, burnAmount)

	authUser = mainchainSelectAccount(mainchainClient, 6)
	authUser.Value.Set(crossChainFee)
	burnTx, err := mainchainTNT20TokenBankInstance.BurnVouchers(authUser, mainchainTNT20VoucherAddress, accountList[1].fromAddress, burnAmount)
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

	senderMainchainTNT20VoucherBalance, _ = mainchainTNT20VoucherContract.BalanceOf(nil, sender)
	receiverSubchainTNT20TokenBalance, _ = subchainTNT20TokenContract.BalanceOf(nil, receiver)

	fmt.Printf("Subchain TNT20 token contract address: %v\n", subchainTNT20TokenAddress)
	fmt.Printf("Mainchain sender : %v, TNT20 Voucher balance on Mainchain: %v\n", sender, senderMainchainTNT20VoucherBalance)
	fmt.Printf("Subchain receiver: %v, TNT20 Token balance on Subchain   : %v\n\n", receiver, receiverSubchainTNT20TokenBalance)
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
	subchainTNT20VoucherContract, _ := ct.NewTNT20VoucherContract(subchainTNT20VoucherAddress, subchainClient)
	senderSubchainTNT20VoucherBalance, _ := subchainTNT20VoucherContract.BalanceOf(nil, sender)

	mainchainTNT20ContractAddress := tnt20VoucherContractAddress // FIXME: should instantiate a mock TNT20 instead of using the Voucher contract (which causes confusion)
	mainchainTNT20Contract, _ := ct.NewTNT20VoucherContract(mainchainTNT20ContractAddress, mainchainClient)
	receiverMainchainTNT20TokenBalance, _ := mainchainTNT20Contract.BalanceOf(nil, receiver)

	fmt.Printf("Subchain TNT20 Voucher address: %v\n", subchainTNT20VoucherAddress)
	fmt.Printf("Subchain sender   : %v, TNT20 Voucher balance on Subchain: %v\n", sender, senderSubchainTNT20VoucherBalance)
	fmt.Printf("Mainchain receiver: %v, TNT20 Token balance on Mainchin  : %v\n\n", receiver, receiverMainchainTNT20TokenBalance)

	authUser := subchainSelectAccount(subchainClient, 1)
	subchainTNT20VoucherContract.Approve(authUser, subchainTNT20TokenBankAddress, burnAmount)

	authUser = subchainSelectAccount(subchainClient, 1)
	burnTx, err := subchainTNT20TokenBank.BurnVouchers(authUser, subchainTNT20VoucherAddress, sender, burnAmount)
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

	senderSubchainTNT20VoucherBalance, _ = subchainTNT20VoucherContract.BalanceOf(nil, sender)
	receiverMainchainTNT20TokenBalance, _ = mainchainTNT20Contract.BalanceOf(nil, receiver)

	fmt.Printf("Mainchain TNT20 token contract address: %v\n", mainchainTNT20ContractAddress)
	fmt.Printf("Subchain sender   : %v, TNT20 Voucher balance on Subchain: %v\n", sender, senderSubchainTNT20VoucherBalance)
	fmt.Printf("Mainchain receiver: %v, TNT20 Token balance on Mainchain : %v\n\n", receiver, receiverMainchainTNT20TokenBalance)
}
