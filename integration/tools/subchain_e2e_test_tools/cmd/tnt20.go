package cmd

import (
	"math/big"

	"github.com/spf13/cobra"
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/tools"
)

var startMainchainTNT20LockCmd = &cobra.Command{
	Use:   "MainchainTNT20Lock",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.MainchainTNT20Lock(big.NewInt(amount))
	},
}

var startSubchainTNT20LockCmd = &cobra.Command{
	Use:   "SubchainTNT20Lock",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.SubchainTNT20Lock(big.NewInt(amount))
	},
}

var startSubchainTNT20BurnCmd = &cobra.Command{
	Use:   "SubchainTNT20Burn",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.SubchainTNT20Burn(big.NewInt(amount))
	},
}

var startMainchainTNT20BurnCmd = &cobra.Command{
	Use:   "MainchainTNT20Burn",
	Short: "Start Thetasubchain node.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.MainchainTNT20Burn(big.NewInt(amount))
	},
}

func init() {
	rootCmd.AddCommand(startMainchainTNT20LockCmd)
	rootCmd.AddCommand(startSubchainTNT20LockCmd)
	rootCmd.AddCommand(startSubchainTNT20BurnCmd)
	rootCmd.AddCommand(startMainchainTNT20BurnCmd)

}
