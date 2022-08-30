package cmd

import (
	"math/big"

	"github.com/spf13/cobra"
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/tools"
)

var startMainchainTFuelLockCmd = &cobra.Command{
	Use: "MainchainTFuelLock",
	Run: func(cmd *cobra.Command, args []string) {
		tools.MainchainTFuelLock(big.NewInt(amount))
	},
}

var startSubchainTFuelBurnCmd = &cobra.Command{
	Use: "SubchainTFuelBurn",
	Run: func(cmd *cobra.Command, args []string) {
		tools.SubchainTFuelBurn(big.NewInt(amount))
	},
}

func init() {
	rootCmd.AddCommand(startMainchainTFuelLockCmd)
	rootCmd.AddCommand(startSubchainTFuelBurnCmd)
	startMainchainTFuelLockCmd.PersistentFlags().Int64Var(&amount, "amount", 10, "amount")
	startSubchainTFuelBurnCmd.PersistentFlags().Int64Var(&amount, "amount", 10, "amount")
}
