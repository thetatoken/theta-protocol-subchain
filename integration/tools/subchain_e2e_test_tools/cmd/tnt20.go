package cmd

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/tools"
)

var startMainchainTNT20LockCmd = &cobra.Command{
	Use: "MainchainTNT20Lock",
	Run: func(cmd *cobra.Command, args []string) {
		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}
		tools.MainchainTNT20Lock(amountInt)
	},
}

var startSubchainTNT20LockCmd = &cobra.Command{
	Use: "SubchainTNT20Lock",
	Run: func(cmd *cobra.Command, args []string) {
		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}
		tools.SubchainTNT20Lock(amountInt)
	},
}

var startSubchainTNT20BurnCmd = &cobra.Command{
	Use: "SubchainTNT20Burn",
	Run: func(cmd *cobra.Command, args []string) {
		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}
		tools.SubchainTNT20Burn(amountInt)
	},
}

var startMainchainTNT20BurnCmd = &cobra.Command{
	Use: "MainchainTNT20Burn",
	Run: func(cmd *cobra.Command, args []string) {
		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}
		tools.MainchainTNT20Burn(amountInt)
	},
}
var contractAddress, accountAddress string
var chainID int64
var starTNT20QueryCmd = &cobra.Command{
	Use: "MainchainTNT20Query",
	Run: func(cmd *cobra.Command, args []string) {
		tools.QueryTNT20(chainID, contractAddress, accountAddress)
	},
}

func init() {
	rootCmd.AddCommand(startMainchainTNT20LockCmd)
	rootCmd.AddCommand(startSubchainTNT20LockCmd)
	rootCmd.AddCommand(startSubchainTNT20BurnCmd)
	rootCmd.AddCommand(startMainchainTNT20BurnCmd)
	rootCmd.AddCommand(starTNT20QueryCmd)

	startMainchainTNT20LockCmd.PersistentFlags().StringVar(&amount, "amount", "10", "amount")
	startSubchainTNT20LockCmd.PersistentFlags().StringVar(&amount, "amount", "10", "amount")
	startSubchainTNT20BurnCmd.PersistentFlags().StringVar(&amount, "amount", "10", "amount")
	startMainchainTNT20BurnCmd.PersistentFlags().StringVar(&amount, "amount", "10", "amount")
	starTNT20QueryCmd.PersistentFlags().StringVar(&contractAddress, "contractAddress", "", "contractAddress")
	starTNT20QueryCmd.PersistentFlags().StringVar(&accountAddress, "accountAddress", "", "accountAddress")
	starTNT20QueryCmd.PersistentFlags().Int64Var(&chainID, "chainID", 366, "chainID")
}
