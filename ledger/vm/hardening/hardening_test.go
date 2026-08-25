package hardening

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/store/database/backend"
	"github.com/thetatoken/thetasubchain/ledger/state"
	"github.com/thetatoken/thetasubchain/ledger/vm"
)

// buildCall returns the bytecode for CALL(target) with empty args and no value,
// followed by POP to discard the success flag.
func buildCall(target common.Address) []byte {
	code := []byte{
		0x60, 0x00, // retSize
		0x60, 0x00, // retOffset
		0x60, 0x00, // argsSize
		0x60, 0x00, // argsOffset
		0x60, 0x00, // value
		0x73, // PUSH20 target
	}
	code = append(code, target.Bytes()...)
	code = append(code, 0x61, 0xff, 0xff, // PUSH2 gas
		0xf1, // CALL
		0x50) // POP
	return code
}

// Logs emitted inside a reverted sub-call must not survive into the receipt:
// the frame's state changes are unwound, so the events it emitted have to be
// unwound with them.
func TestLogsDiscardedOnNestedRevert(t *testing.T) {
	assert := assert.New(t)

	storeView := state.NewStoreView(0, common.Hash{}, backend.NewMemDatabase())

	caller := types.MakeAccWithInitBalance("caller_a", types.NewCoins(90000000, 50000000000))
	storeView.SetAccount(caller.Address, &caller.Account)
	storeView.IncrementHeight()
	storeView.Save()

	innerAddr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	outerAddr := common.HexToAddress("0x2000000000000000000000000000000000000002")

	// inner: LOG0; REVERT
	innerCode := []byte{0x60, 0x00, 0x60, 0x00, 0xa0, 0x60, 0x00, 0x60, 0x00, 0xfd}

	// outer: CALL inner; POP; STOP
	outerCode := append(buildCall(innerAddr), 0x00)

	storeView.SetCode(innerAddr, innerCode)
	storeView.SetCode(outerAddr, outerCode)

	tx := &types.SmartContractTx{
		From:     types.TxInput{Address: caller.Address},
		To:       types.TxOutput{Address: outerAddr},
		GasLimit: 200000,
		GasPrice: big.NewInt(50),
		Data:     nil,
	}

	blockInfo := vm.NewBlockInfo(1, big.NewInt(1), "test_chain")
	_, _, _, evmErr := vm.Execute(blockInfo, tx, storeView)

	assert.Nil(evmErr)
	assert.Equal(0, len(storeView.PopLogs()))
}

// Logs from a successful sibling sub-call must survive when a later sibling
// reverts, i.e. the fix must not over-truncate.
func TestLogsKeptFromSuccessfulSiblingCall(t *testing.T) {
	assert := assert.New(t)

	storeView := state.NewStoreView(0, common.Hash{}, backend.NewMemDatabase())

	caller := types.MakeAccWithInitBalance("caller_b", types.NewCoins(90000000, 50000000000))
	storeView.SetAccount(caller.Address, &caller.Account)
	storeView.IncrementHeight()
	storeView.Save()

	successAddr := common.HexToAddress("0x3000000000000000000000000000000000000003")
	revertAddr := common.HexToAddress("0x4000000000000000000000000000000000000004")
	outerAddr := common.HexToAddress("0x5000000000000000000000000000000000000005")

	successCode := []byte{0x60, 0x00, 0x60, 0x00, 0xa0, 0x00}                        // LOG0; STOP
	revertCode := []byte{0x60, 0x00, 0x60, 0x00, 0xa0, 0x60, 0x00, 0x60, 0x00, 0xfd} // LOG0; REVERT

	outerCode := append(buildCall(successAddr), buildCall(revertAddr)...)
	outerCode = append(outerCode, 0x00)

	storeView.SetCode(successAddr, successCode)
	storeView.SetCode(revertAddr, revertCode)
	storeView.SetCode(outerAddr, outerCode)

	tx := &types.SmartContractTx{
		From:     types.TxInput{Address: caller.Address},
		To:       types.TxOutput{Address: outerAddr},
		GasLimit: 200000,
		GasPrice: big.NewInt(50),
		Data:     nil,
	}

	blockInfo := vm.NewBlockInfo(1, big.NewInt(1), "test_chain")
	_, _, _, evmErr := vm.Execute(blockInfo, tx, storeView)

	assert.Nil(evmErr)
	assert.Equal(1, len(storeView.PopLogs()))
}

// A reverted frame and an earlier successful sub-call can share the same state
// root, since neither LOG nor a no-op CALL mutates the trie. Checkpoints must
// therefore be identified by ID, not by root hash: keying on the root would
// make RevertToSnapshot() pick the inner frame's checkpoint and keep the log.
func TestLogsDroppedWhenRevertedFrameSharesHashWithSubcall(t *testing.T) {
	assert := assert.New(t)

	storeView := state.NewStoreView(0, common.Hash{}, backend.NewMemDatabase())

	caller := types.MakeAccWithInitBalance("caller_l", types.NewCoins(0, 50000000000))
	storeView.SetAccount(caller.Address, &caller.Account)
	storeView.IncrementHeight()
	storeView.Save()

	outerAddr := common.HexToAddress("0x0e01000000000000000000000000000000000001")
	middleAddr := common.HexToAddress("0x0e02000000000000000000000000000000000002")
	noopAddr := common.HexToAddress("0x0e03000000000000000000000000000000000003")

	// noop: STOP. Returns success without touching state.
	noopCode := []byte{0x00}

	// middle: LOG0; CALL noop; POP; REVERT
	middleCode := []byte{0x60, 0x00, 0x60, 0x00, 0xa0} // LOG0
	middleCode = append(middleCode, buildCall(noopAddr)...)
	middleCode = append(middleCode, 0x60, 0x00, 0x60, 0x00, 0xfd) // REVERT

	// outer: CALL middle; POP; STOP. Outer succeeds.
	outerCode := append(buildCall(middleAddr), 0x00)

	storeView.SetCode(noopAddr, noopCode)
	storeView.SetCode(middleAddr, middleCode)
	storeView.SetCode(outerAddr, outerCode)

	tx := &types.SmartContractTx{
		From:     types.TxInput{Address: caller.Address},
		To:       types.TxOutput{Address: outerAddr},
		GasLimit: 300000,
		GasPrice: big.NewInt(50),
		Data:     nil,
	}

	blockInfo := vm.NewBlockInfo(1, big.NewInt(1), "test_chain")
	_, _, _, evmErr := vm.Execute(blockInfo, tx, storeView)

	assert.Nil(evmErr)
	assert.Equal(0, len(storeView.PopLogs()))
}
