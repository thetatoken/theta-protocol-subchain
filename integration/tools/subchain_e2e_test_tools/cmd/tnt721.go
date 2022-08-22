package cmd

import (
	"math/big"

	"github.com/spf13/cobra"
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/tools"
)

var startMainchainTNT721LockCmd = &cobra.Command{
	Use:   "MainchainTNT721Lock",
	Run: func(cmd *cobra.Command, args []string) {
		tools.MainchainTNT721Lock(big.NewInt(amount))
	},
}

var startSubchainTNT721LockCmd = &cobra.Command{
	Use:   "SubchainTNT721Lock",
	Run: func(cmd *cobra.Command, args []string) {
		tools.SubchainTNT721Lock(big.NewInt(amount))
	},
}

var startSubchainTNT721BurnCmd = &cobra.Command{
	Use:   "SubchainTNT721Burn",
	Run: func(cmd *cobra.Command, args []string) {
		tools.SubchainTNT721Burn(big.NewInt(amount))
	},
}

var startMainchainTNT721BurnCmd = &cobra.Command{
	Use:   "MainchainTNT721Burn",
	Run: func(cmd *cobra.Command, args []string) {
		tools.MainchainTNT721Burn(big.NewInt(amount))
	},
}

func init() {
	rootCmd.AddCommand(startMainchainTNT721LockCmd)
	rootCmd.AddCommand(startSubchainTNT721LockCmd)
	rootCmd.AddCommand(startSubchainTNT721BurnCmd)
	rootCmd.AddCommand(startMainchainTNT721BurnCmd)
}
