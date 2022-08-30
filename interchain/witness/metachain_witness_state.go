package witness

import (
	"math/big"
	"sync"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/theta/store/kvstore"
)

func lastWitnessQueryedHeightKey(sourceChainID *big.Int) common.Bytes {
	return common.Bytes("mw/lwqh/" + sourceChainID.String())
}

type metachainWitnessState struct {
	mutex *sync.Mutex // mutex to for concurrency protection, e.g., the witness thread and consensus thread may access it concurrently
	db    database.Database
}

func newMetachainWitnessState(db database.Database) *metachainWitnessState {
	cache := &metachainWitnessState{
		mutex: &sync.Mutex{},
		db:    db,
	}
	return cache
}

func (mws *metachainWitnessState) getLastQueryedHeightForType(sourceChainID *big.Int) (*big.Int, error) {
	mws.mutex.Lock()
	defer mws.mutex.Unlock()

	height := big.NewInt(0)
	store := kvstore.NewKVStore(mws.db)
	err := store.Get(lastWitnessQueryedHeightKey(sourceChainID), &height)
	if err == nil {
		return height, nil
	}
	return big.NewInt(0), err
}

func (mws *metachainWitnessState) setLastQueryedHeightForType(sourceChainID *big.Int, height *big.Int) error {
	mws.mutex.Lock()
	defer mws.mutex.Unlock()

	store := kvstore.NewKVStore(mws.db)
	err := store.Put(lastWitnessQueryedHeightKey(sourceChainID), height)
	return err
}
