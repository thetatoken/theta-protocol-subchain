package cmd

import (
	"fmt"
	"math/big"

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

var deploySubchainMockTokensCmd = &cobra.Command{
	Use:   "DeployTokens",
	Short: "Deploy mock tokens to both the mainchain and subchain.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.DeployTokens()
	},
}
var StakeSpecialCmd = &cobra.Command{
	Use:   "StakeSpecial",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		subchainIDInt, success := big.NewInt(0).SetString(subchainID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read subchainID: %v", subchainIDInt))
		}
		tools.StakeSpecialCases(validatorAddrStr)
	},
}

var CollateralSpecialCmd = &cobra.Command{
	Use:   "CollateralSpecial",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		subchainIDInt, success := big.NewInt(0).SetString(subchainID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read subchainID: %v", subchainIDInt))
		}
		tools.CollateralSpecialCases(accountID, validatorAddrStr, subchainIDInt)
	},
}

func init() {
	rootCmd.AddCommand(startOneAccountRegisterCmd)
	rootCmd.AddCommand(startOneAccountStakeCmd)
	rootCmd.AddCommand(deploySubchainMockTokensCmd)

	rootCmd.AddCommand(StakeSpecialCmd)
	rootCmd.AddCommand(CollateralSpecialCmd)

	StakeSpecialCmd.PersistentFlags().StringVar(&validatorAddrStr, "validator", "0x33a027c2ac93B66987b3C8eA2Bf5bD9F19E2A004", "validator")

	CollateralSpecialCmd.PersistentFlags().IntVar(&accountID, "accountID", 9, "accountID")
	CollateralSpecialCmd.PersistentFlags().StringVar(&validatorAddrStr, "validator", "0x33a027c2ac93B66987b3C8eA2Bf5bD9F19E2A004", "validator")
	CollateralSpecialCmd.PersistentFlags().StringVar(&subchainID, "subchainID", "360777", "subchainID")

	startOneAccountStakeCmd.PersistentFlags().IntVar(&accountID, "accountID", 1, "accountID")
	startOneAccountStakeCmd.PersistentFlags().StringVar(&validatorAddrStr, "validator", "0x2E833968E5bB786Ae419c4d13189fB081Cc43bab", "validator")
}
