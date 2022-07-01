package orchestrator

import (
	"math/big"
	"strconv"
	"sync"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/theta/store/kvstore"

	score "github.com/thetatoken/thetasubchain/core"
)

func lastOrchestratorQueryedHeightKey(sourceChainID *big.Int, icmeType score.InterChainMessageEventType) common.Bytes {
	switch icmeType {
	case score.IMCEventTypeCrossChainTokenLock:
		return common.Bytes("os/tlq/" + sourceChainID.String())
	case score.IMCEventTypeCrossChainVoucherBurn:
		return common.Bytes("os/vblq/" + sourceChainID.String())
	default:
		return nil
	}
}

func InterChainTokenLockEventNextNonceKey(sourceChainID *big.Int, icmeType score.InterChainMessageEventType) common.Bytes {
	return common.Bytes("ictlenn/" + sourceChainID.String() + "/" + strconv.FormatUint(uint64(icmeType), 10))
}

func InterChainVoucherBurnEventNextNonceKey(sourceChainID *big.Int, icmeType score.InterChainMessageEventType) common.Bytes {
	return common.Bytes("vblp" + sourceChainID.String() + "/" + strconv.FormatUint(uint64(icmeType), 10))
}

func VoucherBurnStatusInfoKey(sourceChainID *big.Int, icmeType score.InterChainMessageEventType, nonce *big.Int) common.Bytes {
	return common.Bytes("vbsi" + sourceChainID.String() + "/" + strconv.FormatUint(uint64(icmeType), 10) + "/" + nonce.String())
}

type orchestratorState struct {
	mutex *sync.Mutex // mutex to for concurrency protection, e.g., the witness thread and consensus thread may access it concurrently
	db    database.Database
}

func newOrchestratorState(db database.Database) *orchestratorState {
	cache := &orchestratorState{
		mutex: &sync.Mutex{},
		db:    db,
	}
	return cache
}

func (os *orchestratorState) getLastQueryedHeightForType(sourceChainID *big.Int, icmeType score.InterChainMessageEventType) (*big.Int, error) {
	os.mutex.Lock()
	defer os.mutex.Unlock()

	height := big.NewInt(0)
	store := kvstore.NewKVStore(os.db)
	err := store.Get(lastOrchestratorQueryedHeightKey(sourceChainID, icmeType), &height)
	if err == nil {
		return height, nil
	}
	return big.NewInt(0), err
}

func (os *orchestratorState) setLastQueryedHeightForType(sourceChainID *big.Int, icmeType score.InterChainMessageEventType, height *big.Int) error {
	os.mutex.Lock()
	defer os.mutex.Unlock()

	store := kvstore.NewKVStore(os.db)
	err := store.Put(lastOrchestratorQueryedHeightKey(sourceChainID, icmeType), height)
	return err
}
