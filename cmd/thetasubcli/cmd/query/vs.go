package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/cmd/thetasubcli/cmd/utils"
	"github.com/thetatoken/thetasubchain/rpc"

	rpcc "github.com/ybbus/jsonrpc"
)

// vsCmd represents the validator set command.
// Example:
//		thetacli query validators --height=10
var vsCmd = &cobra.Command{
	Use:     "vs",
	Short:   "Get validator set at a given height",
	Example: `thetacli query vs --height=10`,
	Run:     doVsCmd,
}

func doVsCmd(cmd *cobra.Command, args []string) {
	client := rpcc.NewRPCClient(viper.GetString(utils.CfgRemoteRPCEndpoint))

	height := heightFlag
	res, err := client.Call("theta.GetValidatorSetByHeight", rpc.GetValidatorSetByHeightArgs{Height: common.JSONUint64(height)})
	if err != nil {
		utils.Error("Failed to get validator candidate pool: %v\n", err)
	}
	if res.Error != nil {
		utils.Error("Failed to get validator candidate pool: %v\n", res.Error)
	}
	json, err := json.MarshalIndent(res.Result, "", "    ")
	if err != nil {
		utils.Error("Failed to parse server response: %v\n%s\n", err, string(json))
	}
	fmt.Println(string(json))
}

func init() {
	vsCmd.Flags().Uint64Var(&heightFlag, "height", uint64(0), "height of the block")
	vsCmd.MarkFlagRequired("height")
}
