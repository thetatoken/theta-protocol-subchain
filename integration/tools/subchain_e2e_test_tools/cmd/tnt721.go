package cmd

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/tools"
)

var tokenID string
var startMainchainTNT721LockCmd = &cobra.Command{
	Use: "MainchainTNT721Lock",
	Run: func(cmd *cobra.Command, args []string) {
		tokenIDInt, success := big.NewInt(0).SetString(tokenID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read tokenID: %v", tokenID))
		}
		tools.MainchainTNT721Lock(tokenIDInt)
	},
}

var startSubchainTNT721LockCmd = &cobra.Command{
	Use: "SubchainTNT721Lock",
	Run: func(cmd *cobra.Command, args []string) {
		tokenIDInt, success := big.NewInt(0).SetString(tokenID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read tokenID: %v", tokenID))
		}
		tools.SubchainTNT721Lock(tokenIDInt)
	},
}

var startSubchainTNT721BurnCmd = &cobra.Command{
	Use: "SubchainTNT721Burn",
	Run: func(cmd *cobra.Command, args []string) {
		tokenIDInt, success := big.NewInt(0).SetString(tokenID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read tokenID: %v", tokenID))
		}
		tools.SubchainTNT721Burn(tokenIDInt)
	},
}

var startMainchainTNT721BurnCmd = &cobra.Command{
	Use: "MainchainTNT721Burn",
	Run: func(cmd *cobra.Command, args []string) {
		tokenIDInt, success := big.NewInt(0).SetString(tokenID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read tokenID: %v", tokenID))
		}
		tools.MainchainTNT721Burn(tokenIDInt)
	},
}
var starTNT721QueryCmd = &cobra.Command{
	Use: "MainchainTNT721Query",
	Run: func(cmd *cobra.Command, args []string) {
		tokenIDInt, success := big.NewInt(0).SetString(tokenID, 10)
		if !success {
			panic(fmt.Sprintf("Failed to read tokenID: %v", tokenID))
		}
		tools.QueryTNT721(chainID, contractAddress, tokenIDInt)
	},
}

func init() {
	rootCmd.AddCommand(startMainchainTNT721LockCmd)
	rootCmd.AddCommand(startSubchainTNT721LockCmd)
	rootCmd.AddCommand(startSubchainTNT721BurnCmd)
	rootCmd.AddCommand(startMainchainTNT721BurnCmd)
	rootCmd.AddCommand(starTNT721QueryCmd)

	startMainchainTNT721LockCmd.PersistentFlags().StringVar(&tokenID, "tokenID", "10", "tokenID")
	startSubchainTNT721LockCmd.PersistentFlags().StringVar(&tokenID, "tokenID", "10", "tokenID")
	startSubchainTNT721BurnCmd.PersistentFlags().StringVar(&tokenID, "tokenID", "10", "tokenID")
	startMainchainTNT721BurnCmd.PersistentFlags().StringVar(&tokenID, "tokenID", "10", "tokenID")
	starTNT721QueryCmd.PersistentFlags().StringVar(&tokenID, "tokenID", "10", "tokenID")
	starTNT721QueryCmd.PersistentFlags().StringVar(&contractAddress, "contractAddress", "", "contractAddress")
	starTNT721QueryCmd.PersistentFlags().Int64Var(&chainID, "chainID", 366, "chainID")
}
