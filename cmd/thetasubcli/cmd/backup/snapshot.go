package backup

import (
	"encoding/json"
	"fmt"

	"github.com/thetatoken/thetasubchain/cmd/thetasubcli/cmd/utils"
	"github.com/thetatoken/thetasubchain/rpc"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	rpcc "github.com/ybbus/jsonrpc"
)

// snapshotCmd represents the snapshot backup command.
// Example:
//		thetasubcli backup snapshot
var snapshotCmd = &cobra.Command{
	Use:     "snapshot",
	Short:   "backup snapshot",
	Long:    `Backup snapshot.`,
	Example: `thetasubcli backup snapshot --config=../data/subchain/validator --height=888`,
	Run:     doSnapshotCmd,
}

func doSnapshotCmd(cmd *cobra.Command, args []string) {
	client := rpcc.NewRPCClient(viper.GetString(utils.CfgRemoteRPCEndpoint))

	res, err := client.Call("theta.BackupSnapshot", rpc.BackupSnapshotArgs{Config: configFlag, Height: heightFlag, Version: versionFlag})
	if err != nil {
		utils.Error("Failed to get backup snapshot call details: %v\n", err)
	}
	if res.Error != nil {
		utils.Error("Failed to get backup snapshot res details: %v\n", res.Error)
	}
	json, err := json.MarshalIndent(res.Result, "", "    ")
	if err != nil {
		utils.Error("Failed to parse server response: %v\n%v\n", err, string(json))
	}
	fmt.Println(string(json))
}

func init() {
	snapshotCmd.Flags().StringVar(&configFlag, "config", "", "Config dir")
	snapshotCmd.MarkFlagRequired("config")
	snapshotCmd.Flags().Uint64Var(&heightFlag, "height", 0, "Snapshot height")
	snapshotCmd.Flags().Uint64Var(&versionFlag, "version", 0, "Snapshot version.(2 or 3. Default is 2)")
}
