package tx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/thetasubchain/rpc"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thetatoken/thetasubchain/cmd/thetasubcli/cmd/utils"
	scom "github.com/thetatoken/thetasubchain/common"
	stypes "github.com/thetatoken/thetasubchain/ledger/types"

	rpcc "github.com/ybbus/jsonrpc"
)

// smartContractCmd represents the smart_contract command. It will submit a smart contract transaction
// to the blockchain, which will modify the global consensus state when it is included in the blockchain
// Examples:
//   * Deploy a smart contract
//      thetasubcli tx smart_contract --chain="tsub360777" --from=2E833968E5bB786Ae419c4d13189fB081Cc43bab --value=0 --gas_price=100000000wei --gas_limit=200000 --data=608060405234801561001057600080fd5b50610148806100206000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633fa4f2451461005c578063b5a0241a14610087578063ed8b0706146100b2575b600080fd5b34801561006857600080fd5b506100716100df565b6040518082815260200191505060405180910390f35b34801561009357600080fd5b5061009c6100e5565b6040518082815260200191505060405180910390f35b3480156100be57600080fd5b506100dd60048036038101908080359060200190929190505050610112565b005b60005481565b6000806000546000540290506000546000548281151561010157fe5b0414151561010b57fe5b8091505090565b80600081905550505600a165627a7a72305820459c07c1668e919ca760d663b8df04e80634c53ebd49393dca83e81c58ae2a660029 --seq=1
//   * Call an API of a smart contract
//		thetasubcli tx smart_contract --chain="tsub360777" --from=2E833968E5bB786Ae419c4d13189fB081Cc43bab --to=0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA --data=7ff75b46 --gas_price=100000000wei --gas_limit=100000 --seq=9
var smartContractCmd = &cobra.Command{
	Use:   "smart_contract",
	Short: "Call or deploy a smart contract",
	Example: `
	[Deploy a smart contract] 
	thetasubcli tx smart_contract --chain="privatenet" --from=2E833968E5bB786Ae419c4d13189fB081Cc43bab --value=1680 --gas_price=3 --gas_limit=50000 --data=600a600c600039600a6000f3600360135360016013f3 --seq=1	
	
	[Call an API of a smart contract]
	thetasubcli tx smart_contract --chain="privatenet" --from=2E833968E5bB786Ae419c4d13189fB081Cc43bab --to=0x7ad6cea2bc3162e30a3c98d84f821b3233c22647 --gas_price=3 --gas_limit=50000 --seq=2`,
	Long: "smartContractCmd represents the smart_contract command. It will submit a smart contract transaction to the blockchain, which will modify the global consensus state when it is included in the blockchain",
	Run:  doSmartContractCmd,
}

func doSmartContractCmd(cmd *cobra.Command, args []string) {
	wallet, fromAddress, err := walletUnlock(cmd, fromFlag, passwordFlag)
	if err != nil {
		return
	}
	defer wallet.Lock(fromAddress)

	value, ok := types.ParseCoinAmount(valueFlag)
	if !ok {
		utils.Error("Failed to parse value")
	}

	from := types.TxInput{
		Address: common.HexToAddress(fromFlag),
		Coins: types.Coins{
			ThetaWei: new(big.Int).SetUint64(0),
			TFuelWei: value,
		},
		Sequence: seqFlag,
	}

	to := types.TxOutput{
		Address: common.HexToAddress(toFlag),
	}

	gasPrice, ok := types.ParseCoinAmount(gasPriceFlag)
	if !ok {
		utils.Error("Failed to parse gas price")
	}

	data, err := hex.DecodeString(dataFlag)
	if err != nil {
		return
	}

	smartContractTx := &types.SmartContractTx{
		From:     from,
		To:       to,
		GasLimit: gasLimitFlag,
		GasPrice: gasPrice,
		Data:     data,
	}

	sig, err := wallet.Sign(fromAddress, smartContractTx.SignBytes(chainIDFlag))
	if err != nil {
		utils.Error("Failed to sign transaction: %v\n", err)
	}
	smartContractTx.SetSignature(fromAddress, sig)

	raw, err := stypes.TxToBytes(smartContractTx)
	if err != nil {
		utils.Error("Failed to encode transaction: %v\n", err)
	}
	signedTx := hex.EncodeToString(raw)

	client := rpcc.NewRPCClient(viper.GetString(utils.CfgRemoteRPCEndpoint))

	var res *rpcc.RPCResponse
	if asyncFlag {
		res, err = client.Call("theta.BroadcastRawTransactionAsync", rpc.BroadcastRawTransactionArgs{TxBytes: signedTx})
	} else {
		res, err = client.Call("theta.BroadcastRawTransaction", rpc.BroadcastRawTransactionArgs{TxBytes: signedTx})
	}
	if err != nil {
		utils.Error("Failed to broadcast transaction: %v\n", err)
	}
	if res.Error != nil {
		utils.Error("Server returned error: %v\n", res.Error)
	}
	result := &rpc.BroadcastRawTransactionResult{}
	err = res.GetObject(result)
	if err != nil {
		utils.Error("Failed to parse server response: %v\n", err)
	}
	formatted, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		utils.Error("Failed to parse server response: %v\n", err)
	}
	fmt.Printf("Successfully broadcasted transaction:\n%s\n", formatted)
}

func init() {
	smartContractCmd.Flags().StringVar(&chainIDFlag, "chain", "", "Chain ID")
	smartContractCmd.Flags().StringVar(&fromFlag, "from", "", "The caller address")
	smartContractCmd.Flags().StringVar(&toFlag, "to", "", "The smart contract address")
	smartContractCmd.Flags().StringVar(&valueFlag, "value", "0", "Value to be transferred")
	smartContractCmd.Flags().StringVar(&gasPriceFlag, "gas_price", fmt.Sprintf("%dwei", scom.MinimumGasPrice), "The gas price")
	smartContractCmd.Flags().Uint64Var(&gasLimitFlag, "gas_limit", 0, "The gas limit")
	smartContractCmd.Flags().StringVar(&dataFlag, "data", "", "The data for the smart contract")
	smartContractCmd.Flags().Uint64Var(&seqFlag, "seq", 0, "Sequence number of the transaction")
	smartContractCmd.Flags().StringVar(&walletFlag, "wallet", "soft", "Wallet type (soft|nano)")
	smartContractCmd.Flags().BoolVar(&asyncFlag, "async", false, "block until tx has been included in the blockchain")
	smartContractCmd.Flags().StringVar(&passwordFlag, "password", "", "password to unlock the wallet")

	smartContractCmd.MarkFlagRequired("chain")
	smartContractCmd.MarkFlagRequired("from")
	smartContractCmd.MarkFlagRequired("gas_price")
	smartContractCmd.MarkFlagRequired("gas_limit")
	smartContractCmd.MarkFlagRequired("seq")
}
