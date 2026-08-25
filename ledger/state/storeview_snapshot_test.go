package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/store/database/backend"
)

// Nested checkpoints must unwind the log list to the exact length recorded when
// each checkpoint was taken, and a checkpoint must be consumed once reverted to.
func TestSnapshotUnwindsLogs(t *testing.T) {
	assert := assert.New(t)

	sv := NewStoreView(0, common.Hash{}, backend.NewMemDatabase())
	addr := common.HexToAddress("0x0000000000000000000000000000000000000abc")

	sv.AddLog(&types.Log{Address: addr})
	outer := sv.Snapshot()

	sv.AddLog(&types.Log{Address: addr})
	inner := sv.Snapshot()

	sv.AddLog(&types.Log{Address: addr})
	assert.Equal(3, len(sv.logs))

	sv.RevertToSnapshot(inner)
	assert.Equal(2, len(sv.logs))

	sv.RevertToSnapshot(outer)
	assert.Equal(1, len(sv.logs))

	// The outer checkpoint has been consumed by the revert above; reverting to
	// it a second time must be a no-op rather than a panic or an over-truncation.
	sv.RevertToSnapshot(outer)
	assert.Equal(1, len(sv.logs))

	assert.Equal(1, len(sv.PopLogs()))
	assert.Equal(0, len(sv.logs))
}

// Balance changes are journaled alongside the logs and must unwind with them.
func TestSnapshotUnwindsBalanceChanges(t *testing.T) {
	assert := assert.New(t)

	sv := NewStoreView(0, common.Hash{}, backend.NewMemDatabase())
	addr := common.HexToAddress("0x0000000000000000000000000000000000000def")

	sv.addBalanceChange(&types.BalanceChange{Address: addr})
	id := sv.Snapshot()

	sv.addBalanceChange(&types.BalanceChange{Address: addr})
	sv.addBalanceChange(&types.BalanceChange{Address: addr})
	assert.Equal(3, len(sv.balanceChanges))

	sv.RevertToSnapshot(id)
	assert.Equal(1, len(sv.balanceChanges))
}

// PopLogs()/PopBalanceChanges() end the transaction, so the checkpoint stack
// must not leak across transactions.
func TestSnapshotStackResetBetweenTransactions(t *testing.T) {
	assert := assert.New(t)

	sv := NewStoreView(0, common.Hash{}, backend.NewMemDatabase())
	addr := common.HexToAddress("0x0000000000000000000000000000000000000fed")

	sv.AddLog(&types.Log{Address: addr})
	stale := sv.Snapshot()
	sv.PopLogs()
	assert.Equal(0, len(sv.snapshots))

	// A checkpoint from the previous transaction must not match anything in the
	// new one, and must therefore not discard the new transaction's logs.
	sv.AddLog(&types.Log{Address: addr})
	sv.RevertToSnapshot(stale)
	assert.Equal(1, len(sv.logs))
}
