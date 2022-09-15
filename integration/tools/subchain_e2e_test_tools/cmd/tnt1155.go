package cmd

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/tools"
)

var startMainchainTNT1155LockCmd = &cobra.Command{
	Use: "MainchainTNT1155Lock",
	Run: func(cmd *cobra.Command, args []string) {
		tokenIDInt, success := big.NewInt(0).SetString(tokenID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read tokenID: %v", tokenID))
		}

		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}

		tools.MainchainTNT1155Lock(tokenIDInt, amountInt)
	},
}

var startSubchainTNT1155LockCmd = &cobra.Command{
	Use: "SubchainTNT1155Lock",
	Run: func(cmd *cobra.Command, args []string) {
		tokenIDInt, success := big.NewInt(0).SetString(tokenID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read tokenID: %v", tokenID))
		}

		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}

		tools.SubchainTNT1155Lock(tokenIDInt, amountInt)
	},
}

var startSubchainTNT1155BurnCmd = &cobra.Command{
	Use: "SubchainTNT1155Burn",
	Run: func(cmd *cobra.Command, args []string) {
		tokenIDInt, success := big.NewInt(0).SetString(tokenID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read tokenID: %v", tokenID))
		}

		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}

		tools.SubchainTNT1155Burn(tokenIDInt, amountInt)
	},
}

var startMainchainTNT1155BurnCmd = &cobra.Command{
	Use: "MainchainTNT1155Burn",
	Run: func(cmd *cobra.Command, args []string) {
		tokenIDInt, success := big.NewInt(0).SetString(tokenID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read tokenID: %v", tokenID))
		}

		amountInt, success := big.NewInt(0).SetString(amount, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read amount: %v", amount))
		}

		tools.MainchainTNT1155Burn(tokenIDInt, amountInt)
	},
}

// var startTNT1155QueryCmd = &cobra.Command{
// 	Use: "MainchainTNT1155Query",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		tokenIDInt, success := big.NewInt(0).SetString(tokenID, 10)
// 		if !success {
// 			panic(fmt.Sprintf("Failed to read tokenID: %v", tokenID))
// 		}
// 		tools.QueryTNT1155(chainID, contractAddress, tokenIDInt)
// 	},
// }

func init() {
	rootCmd.AddCommand(startMainchainTNT1155LockCmd)
	rootCmd.AddCommand(startSubchainTNT1155LockCmd)
	rootCmd.AddCommand(startSubchainTNT1155BurnCmd)
	rootCmd.AddCommand(startMainchainTNT1155BurnCmd)
	// rootCmd.AddCommand(starTNT1155QueryCmd)

	startMainchainTNT1155LockCmd.PersistentFlags().StringVar(&tokenID, "tokenID", "8", "tokenID")
	startMainchainTNT1155LockCmd.PersistentFlags().StringVar(&amount, "amount", "100", "amount")

	startSubchainTNT1155BurnCmd.PersistentFlags().StringVar(&tokenID, "tokenID", "8", "tokenID")
	startSubchainTNT1155BurnCmd.PersistentFlags().StringVar(&amount, "amount", "100", "amount")

	startSubchainTNT1155LockCmd.PersistentFlags().StringVar(&tokenID, "tokenID", "8", "tokenID")
	startSubchainTNT1155LockCmd.PersistentFlags().StringVar(&amount, "amount", "100", "amount")

	startMainchainTNT1155BurnCmd.PersistentFlags().StringVar(&tokenID, "tokenID", "8", "tokenID")
	startMainchainTNT1155BurnCmd.PersistentFlags().StringVar(&amount, "amount", "100", "amount")

	// 	startTNT1155QueryCmd.PersistentFlags().StringVar(&tokenID, "tokenID", "10", "tokenID")
	// 	startTNT1155QueryCmd.PersistentFlags().StringVar(&contractAddress, "contractAddress", "", "contractAddress")
	// 	startTNT1155QueryCmd.PersistentFlags().Int64Var(&chainID, "chainID", 366, "chainID")
}
