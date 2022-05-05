package query

import (
	"encoding/json"
	"fmt"

	"github.com/thetatoken/thetasubchain/cmd/thetasubcli/cmd/utils"
	"github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/rpc"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	rpcc "github.com/ybbus/jsonrpc"
)

// accountCmd represents the account command.
// Example:
//		thetacli query token_bank_address --token_type=0
var tokenBankAddrCmd = &cobra.Command{
	Use:     "token_bank_address",
	Short:   "Get token bank address",
	Long:    `Get token bank address.`,
	Example: `thetacli query token_bank_address`,
	Run:     doTokenBankAddrCmd,
}

func doTokenBankAddrCmd(cmd *cobra.Command, args []string) {
	client := rpcc.NewRPCClient(viper.GetString(utils.CfgRemoteRPCEndpoint))

	res, err := client.Call("theta.GetTokenBankContractAddress", rpc.GetTokenBankContractAddressArgs{
		TokenType: core.CrossChainTokenType(tokenTypeFlag),
	})
	if err != nil {
		utils.Error("Failed to get token bank address: %v\n", err)
	}
	if res.Error != nil {
		utils.Error("Failed to get token bank address: %v\n", res.Error)
	}
	json, err := json.MarshalIndent(res.Result, "", "    ")
	if err != nil {
		utils.Error("Failed to parse server response: %v\n%v\n", err, string(json))
	}
	fmt.Println(string(json))
}

func init() {
	tokenBankAddrCmd.Flags().IntVar(&tokenTypeFlag, "token_type", int(0), "token type")
	tokenBankAddrCmd.MarkFlagRequired("token_type")
}
