package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var amount string
var rootCmd = &cobra.Command{
	Use:   "subchain_e2e_test_tools",
	Short: "subchain end-to-end test tools",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
