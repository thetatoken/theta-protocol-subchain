package core

import (
	"errors"
	"math/big"
	"strconv"
	"sync"

	"github.com/thetatoken/theta/common"
	ts "github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/theta/store/kvstore"
)

// ------------------------------------ Inter-Chain Event Cache ----------------------------------------------

var (
	// ErrInterChainMessageEventrNotFound for ID is not found in crosschain transfer event set.
	ErrInterChainMessageEventNotFound      = errors.New("InterChainMessageEventNotFound")
	ErrInterChainMessageEventExisted       = errors.New("InterChainMessageEventrExisted")
	ErrInterChainMessageEventPersistFailed = errors.New("InterChainMessageEventPersistFailed")
)

// InterChainEventIndexKey constructs the DB key for the given block hash.
func InterChainEventIndexKey(sourceChainID *big.Int, icmeType InterChainMessageEventType, nonce *big.Int) common.Bytes {
	return common.Bytes("ice/" + sourceChainID.String() + "/" + strconv.FormatUint(uint64(icmeType), 10) + "/" + nonce.String())
}

func InterChainTokenLockEventNextNonceKey(sourceChainID *big.Int, icmeType InterChainMessageEventType) common.Bytes {
	return common.Bytes("ictlenn/" + sourceChainID.String() + "/" + strconv.FormatUint(uint64(icmeType), 10))
}

func LastQueryedHeightKey(sourceChainID *big.Int, icmeType InterChainMessageEventType) common.Bytes {
	switch icmeType {
	case IMCEventTypeCrossChainTokenLock:
		return common.Bytes("tlq/" + sourceChainID.String())
	case IMCEventTypeCrossChainVoucherBurn:
		return common.Bytes("vblq/" + sourceChainID.String())
	default:
		return nil
	}
}

func InterChainVoucherBurnEventNextNonceKey(sourceChainID *big.Int, icmeType InterChainMessageEventType) common.Bytes {
	return common.Bytes("vblp" + sourceChainID.String() + "/" + strconv.FormatUint(uint64(icmeType), 10))
}

func VoucherBurnStatusInfoKey(sourceChainID *big.Int, icmeType InterChainMessageEventType, nonce *big.Int) common.Bytes {
	return common.Bytes("vbsi" + sourceChainID.String() + "/" + strconv.FormatUint(uint64(icmeType), 10) + "/" + nonce.String())
}

type InterChainEventCache struct {
	mutex *sync.Mutex // mutex to for concurrency protection, e.g., the witness thread and consensus thread may access it concurrently
	db    database.Database
}

// NewInterChainEventCache creates a new crosschain transfer event cache instance.
func NewInterChainEventCache(db database.Database) *InterChainEventCache {
	cache := &InterChainEventCache{
		mutex: &sync.Mutex{},
		db:    db,
	}
	return cache
}

func (c *InterChainEventCache) Insert(event *InterChainMessageEvent) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(InterChainEventIndexKey(event.SourceChainID, event.Type, event.Nonce), event)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) InsertList(events []*InterChainMessageEvent) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	for _, event := range events {
		err := store.Put(InterChainEventIndexKey(event.SourceChainID, event.Type, event.Nonce), event)
		if err != nil {
			return err // the caller should handle the error
		}
	}
	return nil
}

func (c *InterChainEventCache) Delete(sourceChainID *big.Int, imceType InterChainMessageEventType, nonce *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Delete(InterChainEventIndexKey(sourceChainID, imceType, nonce))
	return err // the caller should handle the error
}

func (c *InterChainEventCache) Get(sourceChainID *big.Int, imceType InterChainMessageEventType, nonce *big.Int) (*InterChainMessageEvent, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	event := InterChainMessageEvent{}
	store := kvstore.NewKVStore(c.db)
	logger.Debugf("Source chain %v getting %v event %v", sourceChainID.String(), imceType, nonce)
	err := store.Get(InterChainEventIndexKey(sourceChainID, imceType, nonce), &event)
	return &event, err // the caller should handle the error
}

func (c *InterChainEventCache) Exists(sourceChainID *big.Int, imceType InterChainMessageEventType, nonce *big.Int) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	event := InterChainMessageEvent{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(InterChainEventIndexKey(sourceChainID, imceType, nonce), &event)
	if err == nil {
		return true, nil
	}

	if err == ts.ErrKeyNotFound {
		return false, nil
	}

	return false, err // the caller should handle the error
}

// ------------------------------------ Getters and Setters for utility values --------------------------------------------

func (c *InterChainEventCache) GetLastQueryedHeightForType(sourceChainID *big.Int, icmeType InterChainMessageEventType) (*big.Int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	height := big.NewInt(0)
	store := kvstore.NewKVStore(c.db)
	err := store.Get(LastQueryedHeightKey(sourceChainID, icmeType), &height)
	if err == nil {
		return height, nil
	}
	return big.NewInt(0), err
}

func (c *InterChainEventCache) SetLastQueryedHeightForType(sourceChainID *big.Int, icmeType InterChainMessageEventType, height *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(LastQueryedHeightKey(sourceChainID, icmeType), height)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) GetNextVoucherBurnNonceForType(sourceChainID *big.Int, icmeType InterChainMessageEventType) (*big.Int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	nonce := big.NewInt(0)
	store := kvstore.NewKVStore(c.db)
	err := store.Get(InterChainVoucherBurnEventNextNonceKey(sourceChainID, icmeType), &nonce)
	return nonce, err
}

func (c *InterChainEventCache) SetNextVoucherBurnNonceForType(sourceChainID *big.Int, icmeType InterChainMessageEventType, nonce *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(InterChainVoucherBurnEventNextNonceKey(sourceChainID, icmeType), nonce)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) GetVoucherBurnStatus(sourceChainID *big.Int, icmeType InterChainMessageEventType, nonce *big.Int) (*VoucherBurnEventStatusInfo, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	statusInfo := VoucherBurnEventStatusInfo{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(VoucherBurnStatusInfoKey(sourceChainID, icmeType, nonce), &statusInfo)
	return &statusInfo, err
}

func (c *InterChainEventCache) SetVoucherBurnStatus(sourceChainID *big.Int, statusInfo *VoucherBurnEventStatusInfo) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(VoucherBurnStatusInfoKey(sourceChainID, statusInfo.Type, statusInfo.Nonce), statusInfo)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) VoucherBurnNonceExists(sourceChainID *big.Int, icmeType InterChainMessageEventType, nonce *big.Int) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	statusInfo := VoucherBurnEventStatusInfo{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(VoucherBurnStatusInfoKey(sourceChainID, icmeType, nonce), &statusInfo)
	if err == nil {
		return true, nil
	}

	if err == ts.ErrKeyNotFound {
		return false, nil
	}

	return false, err // the caller should handle the error
}

func (c *InterChainEventCache) GetNextTransferNonceForType(sourceChainID *big.Int, icmeType InterChainMessageEventType) (*big.Int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	nonce := big.NewInt(0)
	store := kvstore.NewKVStore(c.db)
	err := store.Get(InterChainTokenLockEventNextNonceKey(sourceChainID, icmeType), &nonce)
	return nonce, err
}

func (c *InterChainEventCache) SetNextTransferNonceForType(sourceChainID *big.Int, icmeType InterChainMessageEventType, nonce *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(InterChainTokenLockEventNextNonceKey(sourceChainID, icmeType), nonce)
	return err // the caller should handle the error
}
