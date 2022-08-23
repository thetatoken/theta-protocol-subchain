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
var startOneAccountStakeCmd = &cobra.Command{
	Use:   "AccountStake",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.OneAccountStake(accountID)
	},
}

func init() {
	rootCmd.AddCommand(startOneAccountRegisterCmd)
	rootCmd.AddCommand(startOneAccountStakeCmd)
	startOneAccountStakeCmd.PersistentFlags().IntVar(&accountID, "id", 1, "id")
}
