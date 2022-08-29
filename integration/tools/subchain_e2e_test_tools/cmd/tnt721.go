package cmd

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/tools"
)

var startMainchainTNT721LockCmd = &cobra.Command{
	Use: "MainchainTNT721Lock",
	Run: func(cmd *cobra.Command, args []string) {
		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}
		tools.MainchainTNT721Lock(amountInt)
	},
}

var startSubchainTNT721LockCmd = &cobra.Command{
	Use: "SubchainTNT721Lock",
	Run: func(cmd *cobra.Command, args []string) {
		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}
		tools.SubchainTNT721Lock(amountInt)
	},
}

var startSubchainTNT721BurnCmd = &cobra.Command{
	Use: "SubchainTNT721Burn",
	Run: func(cmd *cobra.Command, args []string) {
		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}
		tools.SubchainTNT721Burn(amountInt)
	},
}

var startMainchainTNT721BurnCmd = &cobra.Command{
	Use: "MainchainTNT721Burn",
	Run: func(cmd *cobra.Command, args []string) {
		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}
		tools.MainchainTNT721Burn(amountInt)
	},
}

func init() {
	rootCmd.AddCommand(startMainchainTNT721LockCmd)
	rootCmd.AddCommand(startSubchainTNT721LockCmd)
	rootCmd.AddCommand(startSubchainTNT721BurnCmd)
	rootCmd.AddCommand(startMainchainTNT721BurnCmd)
}
