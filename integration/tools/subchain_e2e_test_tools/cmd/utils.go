package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/tools"
)

var startOneAccountRegisterCmd = &cobra.Command{
	Use:   "RegisterSubchain",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.OneAccountRegister()
	},
}
var accountID int
var startOneAcoountStakeCmd = &cobra.Command{
	Use:   "AcoountStake",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.OneAcoountStake(accountID)
	},
}

func init() {
	rootCmd.AddCommand(startOneAccountRegisterCmd)
	rootCmd.AddCommand(startOneAcoountStakeCmd)
	startOneAcoountStakeCmd.PersistentFlags().IntVar(&accountID, "id", 1, "id")
}
