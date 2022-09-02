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
var validatorAddrStr string

var startOneAccountStakeCmd = &cobra.Command{
	Use:   "AccountStake",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.StakeToValidatorFromAccount(accountID, validatorAddrStr)
	},
}

func init() {
	rootCmd.AddCommand(startOneAccountRegisterCmd)
	rootCmd.AddCommand(startOneAccountStakeCmd)
	startOneAccountStakeCmd.PersistentFlags().IntVar(&accountID, "id", 1, "id")
	startOneAccountStakeCmd.PersistentFlags().StringVar(&validatorAddrStr, "validator", "0x2E833968E5bB786Ae419c4d13189fB081Cc43bab", "validator")
}
