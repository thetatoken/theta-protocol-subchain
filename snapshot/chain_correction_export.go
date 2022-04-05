package snapshot

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/crypto"
	"github.com/thetatoken/theta/ledger/types"

	sbc "github.com/thetatoken/thetasubchain/blockchain"
	score "github.com/thetatoken/thetasubchain/core"
	stypes "github.com/thetatoken/thetasubchain/ledger/types"
)

func ExcludeTxs(txs []common.Bytes, exclusionTxMap map[string]bool, chain *sbc.Chain) (results []common.Bytes) {
	for _, tx := range txs {
		t, err := stypes.TxFromBytes(tx)
		if err != nil {
			continue
		}

		// exclude coinbase tx as well
		if _, ok := t.(*types.CoinbaseTx); ok {
			continue
		}
		// exclude stake updating txs as well
		if _, ok := t.(*types.DepositStakeTx); ok {
			continue
		}
		if _, ok := t.(*types.WithdrawStakeTx); ok {
			continue
		}

		hash := crypto.Keccak256Hash(tx).Hex()
		if _, ok := exclusionTxMap[hash]; !ok {
			results = append(results, tx)
		}
	}
	return
}

func ExportChainCorrection(chain *sbc.Chain, ledger score.Ledger, snapshotHeight uint64, endBlockHash common.Hash, backupDir string, exclusionTxs []string) (backupFile string, blockHashMap map[uint64]string, err error) {
	block, err := chain.FindBlock(endBlockHash)
	if err != nil {
		return "", nil, fmt.Errorf("Can't find block for hash %v", endBlockHash)
	}
	if !block.Status.IsFinalized() {
		return "", nil, fmt.Errorf("End Block %v is not finalized yet", endBlockHash)
	}

	if snapshotHeight >= block.Height {
		return "", nil, errors.New("Start height must be < end height")
	}

	backupFile = "theta_chain_correction-" + strconv.FormatUint(snapshotHeight, 10) + "-" + strconv.FormatUint(block.Height, 10)
	backupPath := path.Join(backupDir, backupFile)
	file, err := os.Create(backupPath)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	var stack []*score.ExtendedBlock

	exclusionTxMap := make(map[string]bool)
	for _, exclusion := range exclusionTxs {
		exclusionTxMap[exclusion] = true
	}

	for {
		block.Txs = ExcludeTxs(block.Txs, exclusionTxMap, chain)
		block.TxHash = score.CalculateRootHash(block.Txs)
		block.UpdateHash()
		stack = append(stack, block)

		if block.Height <= snapshotHeight+1 {
			break
		}
		parentBlock, err := chain.FindBlock(block.Parent)
		if err != nil {
			return "", nil, fmt.Errorf("Can't find block for %v", block.Hash())
		}
		block = parentBlock
	}

	// var bh common.Hash
	var snapshot, parent *score.ExtendedBlock
	blocks := chain.FindBlocksByHeight(snapshotHeight)
	for _, block := range blocks {
		if block.Status.IsDirectlyFinalized() {
			snapshot = block
			parent = block
			break
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		block = stack[i]
		block.Parent = parent.Hash()
		block.HCC.BlockHash = snapshot.Hash()
		block.Children = []common.Hash{}

		//result := ledger.ResetState(parent.Height, parent.StateHash)
		result := ledger.ResetState(parent.Block)
		if result.IsError() {
			return "", nil, fmt.Errorf("%v", result.String())
		}

		hash, result := ledger.ApplyBlockTxsForChainCorrection(block.Block)
		if result.IsError() {
			return "", nil, fmt.Errorf("%v", result.String())
		}
		block.StateHash = hash
		block.UpdateHash()

		parent.Children = []common.Hash{block.Hash()}

		parent = block
	}

	blockHashMap = make(map[uint64]string)
	for i := 0; i < len(stack); i++ {
		block = stack[i]

		backupBlock := &score.BackupBlock{Block: block}
		writeBlock(writer, backupBlock)

		blockHashMap[block.Height] = block.Hash().Hex()
	}

	return
}
