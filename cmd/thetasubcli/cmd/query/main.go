package query

import (
	"github.com/spf13/cobra"
)

var (
	purposeFlag          uint8
	heightFlag           uint64
	addressFlag          string
	previewFlag          bool
	resourceIDFlag       string
	hashFlag             string
	startFlag            uint64
	endFlag              uint64
	skipEdgeNodeFlag     bool
	includeEthTxHashFlag bool
	tokenTypeFlag        int
)

// QueryCmd represents the query command
var QueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query entities stored in blockchain",
}

func init() {
	QueryCmd.AddCommand(statusCmd)
	QueryCmd.AddCommand(accountCmd)
	QueryCmd.AddCommand(blockCmd)
	QueryCmd.AddCommand(txCmd)
	QueryCmd.AddCommand(vsCmd)
	QueryCmd.AddCommand(peersCmd)
	QueryCmd.AddCommand(versionCmd)
	QueryCmd.AddCommand(tokenBankAddrCmd)
}
