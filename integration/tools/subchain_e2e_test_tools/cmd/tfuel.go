package cmd

import (
	"math/big"

	"github.com/spf13/cobra"
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/tools"
)

var startMainchainTfuelLockCmd = &cobra.Command{
	Use:   "MainchainTfuelLock",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.MainchainTfuelLock(big.NewInt(amount))
	},
}

var startSubchainTfuelBurnCmd = &cobra.Command{
	Use:   "MainchainTfuelBurn",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.SubchainTfuelBurn(big.NewInt(amount))
	},
}

func init() {
	rootCmd.AddCommand(startMainchainTfuelLockCmd)
	rootCmd.AddCommand(startSubchainTfuelBurnCmd)
}
