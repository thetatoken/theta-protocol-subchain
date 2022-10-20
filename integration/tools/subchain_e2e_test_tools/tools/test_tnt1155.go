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

func MainchainTNT1155Lock(tokenID *big.Int, lockAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT1155 cross-chain transfer...\n")

	sender := accountList[1].fromAddress
	receiver := accountList[1].fromAddress

	mainchainTNT1155ContractAddress := common.HexToAddress("0x47eb28D8139A188C5686EedE1E9D8EDE3Afdd543")
	mainchainTNT1155Contract, err := ct.NewMockTNT1155(mainchainTNT1155ContractAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	mainchainTNT1155TokenBank, err := ct.NewTNT1155TokenBank(mainchainTNT1155TokenBankAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := mainchainSelectAccount(mainchainClient, 0)
	_, err = mainchainTNT1155Contract.Mint(authAccount0, sender, tokenID, lockAmount, []byte(""))
	if err != nil {
		log.Fatal(err)
	}

	mainchainTokenURI, _ := mainchainTNT1155Contract.Uri(nil, tokenID)
	fmt.Printf("Mainchain TNT1155 contract address: %v, tokenID: %v, amount to be sent: %v, tokenURI: %v\n",
		mainchainTNT1155ContractAddress, tokenID, lockAmount, mainchainTokenURI)
	time.Sleep(6 * time.Second)

	expectedSubchainVoucherContractAddress := common.HexToAddress("0x0ede92cAc9161F6C397A604DE508Dcd1e6f43E61")
	subchainTNT1155VoucherContract, err := ct.NewTNT1155VoucherContract(expectedSubchainVoucherContractAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}

	senderMainchainTNT1155TokenBalance, _ := mainchainTNT1155Contract.BalanceOf(nil, sender, tokenID)
	receiverSubchainTNT1155VoucherBalance, _ := subchainTNT1155VoucherContract.BalanceOf(nil, receiver, tokenID)
	fmt.Printf("Mainchain TNT1155 NFT sender         : %v, token balance on Mainchain : %v\n", sender, senderMainchainTNT1155TokenBalance)
	fmt.Printf("Subchain TNT1155 NFT Voucher receiver: %v, voucher balance on Subchain: %v\n\n", receiver, receiverSubchainTNT1155VoucherBalance)

	authAccount1 := mainchainSelectAccount(mainchainClient, 1)
	_, err = mainchainTNT1155Contract.SetApprovalForAll(authAccount1, mainchainTNT1155TokenBankAddress, true)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(6 * time.Second)

	// tokenBankTNT1155BalanceBefore, _ := mainchainTNT1155Contract.BalanceOf(nil, mainchainTNT1155TokenBankAddress, tokenID)
	// fmt.Printf("Mainchain TNT1155TokenBank balance for NFT %v before tranfer: %v\n", tokenID, tokenBankTNT1155BalanceBefore)
	// printTNT1155TokenBankLockedAmount(mainchainTNT1155TokenBankAddress, mainchainTNT1155ContractAddress, subchainID, tokenID, mainchainClient)

	authAccount1 = mainchainSelectAccount(mainchainClient, 1)
	authAccount1.Value.Set(crossChainFee)
	lockTx, err := mainchainTNT1155TokenBank.LockTokens(authAccount1, subchainID, mainchainTNT1155ContractAddress, receiver, tokenID, lockAmount)
	authAccount1.Value.Set(common.Big0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("TNT1155 Token Lock tx hash (Mainchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering TNT1155 NFT (tokenID: %v, amount: %v) from the Mainchain to Subchain %v...\n\n", tokenID, lockAmount, subchainID)

	fmt.Printf("Start transfer, timestamp      : %v\n", time.Now())
	receipt, err := mainchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("burn error")
		return
	}
	fmt.Printf("Token lock confirmed, timestamp: %v\n", time.Now())

	fromHeight, _ := subchainClient.BlockNumber(context.Background())
	var subchainVoucherAddress common.Address
	for {
		time.Sleep(1 * time.Second)
		toHeight, _ := subchainClient.BlockNumber(context.Background())
		result := getSubchainTNT1155VoucherMintLogs(int(fromHeight), int(toHeight), subchainTNT1155TokenBankAddress, accountList[1].fromAddress)
		if result != nil {
			subchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Subchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("End transfer, timestamp        : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// tokenBankTNT1155BalanceAfter, _ := mainchainTNT1155Contract.BalanceOf(nil, mainchainTNT1155TokenBankAddress, tokenID)
	subchainVoucherContract, _ := ct.NewTNT1155VoucherContract(subchainVoucherAddress, subchainClient)
	subchainTokenURI, _ := subchainVoucherContract.Uri(nil, tokenID)

	// printTNT1155TokenBankLockedAmount(mainchainTNT1155TokenBankAddress, mainchainTNT1155ContractAddress, subchainID, tokenID, mainchainClient)

	// fmt.Printf("Mainchain TNT1155TokenBank balance for NFT %v after tranfer: %v\n", tokenID, tokenBankTNT1155BalanceAfter)
	fmt.Printf("Subchain TNT1155 Voucher contract address: %v, tokenID: %v, TokenURI: %v\n", subchainVoucherAddress, tokenID, subchainTokenURI)

	senderMainchainTNT1155TokenBalance, _ = mainchainTNT1155Contract.BalanceOf(nil, sender, tokenID)
	receiverSubchainTNT1155VoucherBalance, _ = subchainTNT1155VoucherContract.BalanceOf(nil, receiver, tokenID)
	fmt.Printf("Mainchain TNT1155 NFT sender         : %v, token balance on Mainchain: %v\n", sender, senderMainchainTNT1155TokenBalance)
	fmt.Printf("Subchain TNT1155 NFT Voucher receiver: %v, voucher balance on Subchain: %v\n\n", receiver, receiverSubchainTNT1155VoucherBalance)
}

func SubchainTNT1155Burn(tokenID *big.Int, burnAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT1155 cross-chain transfer...\n")

	sender := accountList[1].fromAddress
	receiver := accountList[6].fromAddress

	subchainTNT1155TokenBank, _ := ct.NewTNT1155TokenBank(subchainTNT1155TokenBankAddress, subchainClient)

	subchainTNT1155VoucherAddress := common.HexToAddress("0x0ede92cac9161f6c397a604de508dcd1e6f43e61")
	subchainTNT1155VoucherContract, _ := ct.NewTNT1155VoucherContract(subchainTNT1155VoucherAddress, subchainClient)
	senderSubchainTNT1155VoucherBalance, _ := subchainTNT1155VoucherContract.BalanceOf(nil, sender, tokenID)

	mainchainTNT1155ContractAddress := common.HexToAddress("0x47eb28D8139A188C5686EedE1E9D8EDE3Afdd543")
	mainchainTNT1155Contract, _ := ct.NewTNT1155VoucherContract(mainchainTNT1155ContractAddress, mainchainClient)
	receiverMainchainTNT1155TokenBalance, _ := mainchainTNT1155Contract.BalanceOf(nil, receiver, tokenID)
	subchainTokenURI, _ := subchainTNT1155VoucherContract.Uri(nil, tokenID)

	fmt.Printf("Subchain TNT1155 Voucher address: %v, tokenID: %v, tokenURI: %v\n", subchainTNT1155VoucherAddress, tokenID, subchainTokenURI)
	fmt.Printf("Subchain sender   : %v, TNT1155 Voucher balance on Subchain: %v\n", sender, senderSubchainTNT1155VoucherBalance)
	fmt.Printf("Mainchain receiver: %v, TNT1155 Token balance on Mainchin  : %v\n\n", receiver, receiverMainchainTNT1155TokenBalance)

	// printTNT1155TokenBankLockedAmount(mainchainTNT1155TokenBankAddress, mainchainTNT1155ContractAddress, subchainID, tokenID, mainchainClient)

	authUser := subchainSelectAccount(subchainClient, 1)
	subchainTNT1155VoucherContract.SetApprovalForAll(authUser, subchainTNT1155TokenBankAddress, true)

	authUser = subchainSelectAccount(subchainClient, 1)
	authUser.Value.Set(crossChainFee)
	burnTx, err := subchainTNT1155TokenBank.BurnVouchers(authUser, subchainTNT1155VoucherAddress, receiver, tokenID, burnAmount)
	authUser.Value.Set(common.Big0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT1155 Voucher Burn tx hash (Subchain): %v\n", burnTx.Hash().Hex())
	fmt.Printf("Burn %v TNT1155 Vouchers (tokenID:%v) on Subchain %v to recover the authentic tokens on the Mainchain...\n\n", burnAmount, tokenID, subchainID)

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
		updatedBalance, _ := mainchainTNT1155Contract.BalanceOf(nil, receiver, tokenID)
		// fallbackReceiverBalance, _ := mainchainTNT1155Contract.BalanceOf(nil, common.HexToAddress("0x19E7E376E7C213B7E7e7e46cc70A5dD086DAff2A"), tokenID)

		if receiverMainchainTNT1155TokenBalance.Cmp(updatedBalance) != 0 {
			break
		}
		// fmt.Printf("orig balance: %v, updated balance: %v, fallback receiver balance: %v\n", receiverMainchainTNT1155TokenBalance, updatedBalance, fallbackReceiverBalance)
		fmt.Printf("waiting for cross-chain transfer completion...\n")
	}
	fmt.Printf("End transfer, timestamp          : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// printTNT1155TokenBankLockedAmount(mainchainTNT1155TokenBankAddress, mainchainTNT1155ContractAddress, subchainID, tokenID, mainchainClient)

	senderSubchainTNT1155VoucherBalance, _ = subchainTNT1155VoucherContract.BalanceOf(nil, sender, tokenID)
	receiverMainchainTNT1155TokenBalance, _ = mainchainTNT1155Contract.BalanceOf(nil, receiver, tokenID)
	mainchainTokenURI, _ := mainchainTNT1155Contract.Uri(nil, tokenID)

	fmt.Printf("Mainchain TNT1155 token contract address: %v, tokenID: %v, tokenURI: %v\n", mainchainTNT1155ContractAddress, tokenID, mainchainTokenURI)
	fmt.Printf("Subchain sender   : %v, TNT1155 Voucher balance on Subchain: %v\n", sender, senderSubchainTNT1155VoucherBalance)
	fmt.Printf("Mainchain receiver: %v, TNT1155 Token balance on Mainchain  : %v\n\n", receiver, receiverMainchainTNT1155TokenBalance)
}

func SubchainTNT1155Lock(tokenID *big.Int, lockAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT1155 cross-chain transfer...\n")

	sender := accountList[1].fromAddress
	receiver := accountList[6].fromAddress

	subchainTNT1155ContractAddress := common.HexToAddress("0xFfF2bb6198F181C412bb250825a5b3AfF1bB2d3a")
	subchainTNT1155Contract, _ := ct.NewMockTNT1155(subchainTNT1155ContractAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}
	subchainTNT1155TokenBank, err := ct.NewTNT1155TokenBank(subchainTNT1155TokenBankAddress, subchainClient)
	if err != nil {
		log.Fatal(err)
	}

	auth := subchainSelectAccount(subchainClient, 1)
	_, err = subchainTNT1155Contract.Mint(auth, sender, tokenID, lockAmount, []byte(""))
	if err != nil {
		log.Fatal(err)
	}

	subchainTokenURI, _ := subchainTNT1155Contract.Uri(nil, tokenID)
	fmt.Printf("Subchain TNT1155 contract address: %v, tokenID: %v, amount to be sent: %v, tokenURI: %v\n",
		subchainTNT1155ContractAddress, tokenID, lockAmount, subchainTokenURI)
	time.Sleep(6 * time.Second)

	expectedMainchainVoucherContractAddress := common.HexToAddress("0xb0DBBcba1Be5B71Dcb42aB1935773B3675e645e8")
	mainchainTNT1155VoucherContract, err := ct.NewTNT1155VoucherContract(expectedMainchainVoucherContractAddress, mainchainClient)
	if err != nil {
		log.Fatal(err)
	}

	senderSubchainTNT1155TokenBalance, _ := subchainTNT1155Contract.BalanceOf(nil, sender, tokenID)
	receiverMainchainTNT1155VoucherBalance, _ := mainchainTNT1155VoucherContract.BalanceOf(nil, receiver, tokenID)
	fmt.Printf("Subchain TNT1155 NFT sender           : %v, token balance on Subchain   : %v\n", sender, senderSubchainTNT1155TokenBalance)
	fmt.Printf("Mainchain TNT1155 NFT Voucher receiver: %v, voucher balance on Mainchain: %v\n\n", receiver, receiverMainchainTNT1155VoucherBalance)

	auth = subchainSelectAccount(subchainClient, 1)
	_, err = subchainTNT1155Contract.SetApprovalForAll(auth, subchainTNT1155TokenBankAddress, true)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(6 * time.Second)

	// tokenBankTNT1155BalanceBefore, _ := subchainTNT1155Contract.BalanceOf(nil, subchainTNT1155TokenBankAddress, tokenID)
	// fmt.Printf("Subchain TNT1155TokenBank balance for NFT %v before tranfer (locked amount): %v\n\n", tokenID, tokenBankTNT1155BalanceBefore)
	mainchainID, _ := mainchainClient.ChainID(context.Background())
	// printTNT1155TokenBankLockedAmount(subchainTNT1155TokenBankAddress, subchainTNT1155ContractAddress, mainchainID, tokenID, subchainClient)

	auth = subchainSelectAccount(subchainClient, 1)
	auth.Value.Set(crossChainFee)
	lockTx, err := subchainTNT1155TokenBank.LockTokens(auth, mainchainID, subchainTNT1155ContractAddress, receiver, tokenID, lockAmount)
	// lockTx, err := subchainTNT1155TokenBank.LockTokens(auth, subchainID, subchainTNT1155ContractAddress, receiver, tokenID, lockAmount) // for testing a corner case: sending to the same chain
	auth.Value.Set(common.Big0)
	if err != nil {
		log.Fatal(err)
	}

	// totalLockAmount, _ := subchainTNT1155TokenBank.TotalLockedAmounts(nil, subchainID, subchainTNT1155ContractAddress, tokenID)
	// fmt.Printf("Total amount locked for tokenID %v: %v\n", tokenID, totalLockAmount)
	fmt.Printf("TNT1155 Token Lock tx hash (Subchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering TNT1155 NFT (tokenID: %v, amount: %v) from Subchain %v to the Mainchain...\n\n", tokenID, lockAmount, subchainID)

	fmt.Printf("Start transfer, timestamp      : %v\n", time.Now())
	receipt, err := subchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		fmt.Println("lock error")
		return
	}
	fmt.Printf("Token lock confirmed, timestamp: %v\n", time.Now())

	fromHeight, _ := mainchainClient.BlockNumber(context.Background())
	var mainchainVoucherAddress common.Address
	for {
		time.Sleep(1 * time.Second)
		toHeight, _ := mainchainClient.BlockNumber(context.Background())
		result := getMainchainTNT1155VoucherMintLogs(int(fromHeight), int(toHeight), mainchainTNT1155TokenBankAddress, receiver)
		if result != nil {
			mainchainVoucherAddress = *result
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion (scanning Mainchain from height %v to %v)...\n", fromHeight, toHeight)
	}
	fmt.Printf("End transfer, timestamp        : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// printTNT1155TokenBankLockedAmount(subchainTNT1155TokenBankAddress, subchainTNT1155ContractAddress, mainchainID, tokenID, subchainClient)

	// tokenBankTNT1155BalanceAfter, _ := subchainTNT1155Contract.BalanceOf(nil, subchainTNT1155TokenBankAddress, tokenID)
	mainchainTNT1155VoucherContract, _ = ct.NewTNT1155VoucherContract(mainchainVoucherAddress, mainchainClient)
	mainchainTokenURI, _ := mainchainTNT1155VoucherContract.Uri(nil, tokenID)

	// fmt.Printf("Subchain TNT1155TokenBank balance for NFT %v after tranfer (locked amount): %v\n\n", tokenID, tokenBankTNT1155BalanceAfter)
	fmt.Printf("Mainchain TNT1155 Voucher contract address: %v, tokenID: %v, TokenURI: %v\n", mainchainVoucherAddress, tokenID, mainchainTokenURI)

	senderSubchainTNT1155TokenBalance, _ = subchainTNT1155Contract.BalanceOf(nil, sender, tokenID)
	receiverMainchainTNT1155VoucherBalance, _ = mainchainTNT1155VoucherContract.BalanceOf(nil, receiver, tokenID)
	fmt.Printf("Subchain TNT1155 NFT sender           : %v, token balance on Subchain   : %v\n", sender, senderSubchainTNT1155TokenBalance)
	fmt.Printf("Mainchain TNT1155 NFT Voucher receiver: %v, voucher balance on Mainchain: %v\n\n", receiver, receiverMainchainTNT1155VoucherBalance)
}

func MainchainTNT1155Burn(tokenID *big.Int, burnAmount *big.Int) {
	mainchainClient, err := ethclient.Dial("http://localhost:18888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT1155 cross-chain transfer...\n")

	sender := accountList[6].fromAddress
	receiver := accountList[1].fromAddress

	mainchainTNT1155TokenBankInstance, _ := ct.NewTNT1155TokenBank(mainchainTNT1155TokenBankAddress, mainchainClient)

	mainchainTNT1155VoucherAddress := common.HexToAddress("0xb0DBBcba1Be5B71Dcb42aB1935773B3675e645e8")
	mainchainTNT1155VoucherContract, _ := ct.NewTNT1155VoucherContract(mainchainTNT1155VoucherAddress, mainchainClient)
	senderMainchainTNT1155VoucherBalance, _ := mainchainTNT1155VoucherContract.BalanceOf(nil, sender, tokenID)

	subchainTNT1155TokenAddress := common.HexToAddress("0xFfF2bb6198F181C412bb250825a5b3AfF1bB2d3a")
	subchainTNT1155TokenContract, _ := ct.NewTNT1155VoucherContract(subchainTNT1155TokenAddress, subchainClient)
	receiverSubchainTNT1155TokenBalance, _ := subchainTNT1155TokenContract.BalanceOf(nil, receiver, tokenID)
	mainchainTokenURI, _ := mainchainTNT1155VoucherContract.Uri(nil, tokenID)

	fmt.Printf("Mainchain TNT1155 Voucher address: %v, tokenID: %v, tokenURI: %v\n", mainchainTNT1155VoucherAddress, tokenID, mainchainTokenURI)
	fmt.Printf("Mainchain sender : %v, TNT1155 Voucher balance on Mainchain: %v\n", sender, senderMainchainTNT1155VoucherBalance)
	fmt.Printf("Subchain receiver: %v, TNT1155 Token balance on Subchain   : %v\n\n", receiver, receiverSubchainTNT1155TokenBalance)

	// subchainTNT1155TokenBankInstance, _ := ct.NewTNT1155TokenBank(subchainTNT1155TokenBankAddress, subchainClient)
	// tokenBankLockedAmount, _ := subchainTNT1155TokenBankInstance.TotalLockedAmounts(nil, big.NewInt(366), subchainTNT1155TokenAddress, tokenID)
	// subchainTNT1155TokenBankTNT1155Balance, _ := subchainTNT1155TokenContract.BalanceOf(nil, subchainTNT1155TokenBankAddress, tokenID)
	// fmt.Printf("TokenBank locked %v for tokenID %v\n", tokenBankLockedAmount, tokenID)
	// fmt.Printf("TokenBank balance for tokenID %v: %v\n\n", tokenID, subchainTNT1155TokenBankTNT1155Balance)
	// mainchainID, _ := mainchainClient.ChainID(context.Background())
	// printTNT1155TokenBankLockedAmount(subchainTNT1155TokenBankAddress, subchainTNT1155TokenAddress, mainchainID, tokenID, subchainClient)

	authUser := mainchainSelectAccount(mainchainClient, 6)
	mainchainTNT1155VoucherContract.SetApprovalForAll(authUser, mainchainTNT1155TokenBankAddress, true)

	authUser = mainchainSelectAccount(mainchainClient, 6)
	authUser.Value.Set(crossChainFee)
	burnTx, err := mainchainTNT1155TokenBankInstance.BurnVouchers(authUser, mainchainTNT1155VoucherAddress, accountList[1].fromAddress, tokenID, burnAmount)
	authUser.Value.Set(common.Big0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT1155 Voucher Burn tx hash (Mainchain): %v\n", burnTx.Hash().Hex())
	fmt.Printf("Burn %v TNT1155 Vouchers on the Mainchain to recover the authentic tokens on Subchain %v...\n\n", burnAmount, subchainID)

	fmt.Printf("Start transfer, timestamp        : %v\n", time.Now())
	receipt, err := mainchainClient.TransactionReceipt(context.Background(), burnTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("burn error")
		return
	}
	fmt.Printf("Voucher burn confirmed, timestamp: %v\n", time.Now())

	for {
		time.Sleep(1 * time.Second)
		updatedBalance, _ := subchainTNT1155TokenContract.BalanceOf(nil, receiver, tokenID)
		if receiverSubchainTNT1155TokenBalance.Cmp(updatedBalance) != 0 {
			break
		}
		fmt.Printf("waiting for cross-chain transfer completion...\n")
	}
	fmt.Printf("End transfer, timestamp          : %v\n", time.Now())
	fmt.Printf("Cross-chain transfer completed.\n\n")

	// printTNT1155TokenBankLockedAmount(subchainTNT1155TokenBankAddress, subchainTNT1155TokenAddress, mainchainID, tokenID, subchainClient)

	senderMainchainTNT1155VoucherBalance, _ = mainchainTNT1155VoucherContract.BalanceOf(nil, sender, tokenID)
	receiverSubchainTNT1155TokenBalance, _ = subchainTNT1155TokenContract.BalanceOf(nil, receiver, tokenID)
	subchainTokenURI, _ := subchainTNT1155TokenContract.Uri(nil, tokenID)

	fmt.Printf("Subchain TNT1155 token contract address: %v, tokenID: %v, tokenURI: %v\n", subchainTNT1155TokenAddress, tokenID, subchainTokenURI)
	fmt.Printf("Mainchain sender : %v, TNT1155 Voucher balance on Mainchain: %v\n", sender, senderMainchainTNT1155VoucherBalance)
	fmt.Printf("Subchain receiver: %v, TNT1155 Token balance on Subchain   : %v\n\n", receiver, receiverSubchainTNT1155TokenBalance)
}

// func printTNT1155TokenBankLockedAmount(tokenBankAddr common.Address, tokenAddr common.Address, chainID *big.Int, tokenID *big.Int, ethclient *ethclient.Client) {
// 	tnt1155TokenBank, _ := ct.NewTNT1155TokenBank(tokenBankAddr, ethclient)
// 	tokenBankLockedAmount, _ := tnt1155TokenBank.TotalLockedAmounts(nil, chainID, tokenAddr, tokenID)
// 	fmt.Printf("TNT1155TokenBank locked amount for tokenID %v: %v\n\n", tokenID, tokenBankLockedAmount)
// }
