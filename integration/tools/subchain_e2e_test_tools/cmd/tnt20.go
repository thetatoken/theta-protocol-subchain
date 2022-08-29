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

func init() {
	rootCmd.AddCommand(startMainchainTNT20LockCmd)
	rootCmd.AddCommand(startSubchainTNT20LockCmd)
	rootCmd.AddCommand(startSubchainTNT20BurnCmd)
	rootCmd.AddCommand(startMainchainTNT20BurnCmd)

}
