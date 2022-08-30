package cmd

import (
	"math/big"

	"github.com/spf13/cobra"
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/tools"
)

var tokenID int64
var startMainchainTNT721LockCmd = &cobra.Command{
	Use: "MainchainTNT721Lock",
	Run: func(cmd *cobra.Command, args []string) {
		tools.MainchainTNT721Lock(big.NewInt(tokenID))
	},
}

var startSubchainTNT721LockCmd = &cobra.Command{
	Use: "SubchainTNT721Lock",
	Run: func(cmd *cobra.Command, args []string) {
		tools.SubchainTNT721Lock(big.NewInt(tokenID))
	},
}

var startSubchainTNT721BurnCmd = &cobra.Command{
	Use: "SubchainTNT721Burn",
	Run: func(cmd *cobra.Command, args []string) {
		tools.SubchainTNT721Burn(big.NewInt(tokenID))
	},
}

var startMainchainTNT721BurnCmd = &cobra.Command{
	Use: "MainchainTNT721Burn",
	Run: func(cmd *cobra.Command, args []string) {
		tools.MainchainTNT721Burn(big.NewInt(tokenID))
	},
}

var starTNT721QueryCmd = &cobra.Command{
	Use: "MainchainTNT721Query",
	Run: func(cmd *cobra.Command, args []string) {
		tools.QueryTNT721(big.NewInt(tokenID))
	},
}

func init() {
	rootCmd.AddCommand(startMainchainTNT721LockCmd)
	rootCmd.AddCommand(startSubchainTNT721LockCmd)
	rootCmd.AddCommand(startSubchainTNT721BurnCmd)
	rootCmd.AddCommand(startMainchainTNT721BurnCmd)
	rootCmd.AddCommand(starTNT721QueryCmd)

	startMainchainTNT721LockCmd.PersistentFlags().Int64Var(&tokenID, "tokenID", 10, "tokenID")
	startSubchainTNT721LockCmd.PersistentFlags().Int64Var(&tokenID, "tokenID", 10, "tokenID")
	startSubchainTNT721BurnCmd.PersistentFlags().Int64Var(&tokenID, "tokenID", 10, "tokenID")
	startMainchainTNT721BurnCmd.PersistentFlags().Int64Var(&tokenID, "tokenID", 10, "tokenID")
	starTNT721QueryCmd.PersistentFlags().Int64Var(&tokenID, "tokenID", 10, "tokenID")
	starTNT721QueryCmd.PersistentFlags().StringVar(&accountAddress, "accountAddress", "", "accountAddress")
	starTNT721QueryCmd.PersistentFlags().Int64Var(&chainID, "chainID", 366, "chainID")
}
