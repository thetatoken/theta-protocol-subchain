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
func InterChainEventIndexKey(chainID *big.Int, icmeType InterChainMessageEventType, nonce *big.Int) common.Bytes {
	return common.Bytes("ice/" + chainID.String() + "/" + strconv.FormatUint(uint64(icmeType), 10) + "/" + nonce.String())
}

func InterChainTokenLockEventNextNonceKey(chainID *big.Int, icmeType InterChainMessageEventType) common.Bytes {
	return common.Bytes("ictlenn/" + chainID.String() + "/" + strconv.FormatUint(uint64(icmeType), 10))
}

func LastQueryedHeightKey(icmeType InterChainMessageEventType) common.Bytes {
	switch icmeType {
	case IMCEventTypeCrossChainLock:
		return common.Bytes("tlq")
	case IMCEventTypeCrossChainVoucherBurn:
		return common.Bytes("vblq")
	default:
		return nil
	}
}

func InterChainVoucherBurnEventNextNonceKey(icmeType InterChainMessageEventType) common.Bytes {
	return common.Bytes("vblp" + strconv.FormatUint(uint64(icmeType), 10))
}

func VoucherBurnStatusInfoKey(icmeType InterChainMessageEventType, nonce *big.Int) common.Bytes {
	return common.Bytes("vbsi" + strconv.FormatUint(uint64(icmeType), 10) + "/" + nonce.String())
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
	err := store.Put(InterChainEventIndexKey(event.Type, event.Nonce), event)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) InsertList(events []*InterChainMessageEvent) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	for _, event := range events {
		err := store.Put(InterChainEventIndexKey(event.Type, event.Nonce), event)
		if err != nil {
			return err // the caller should handle the error
		}
	}
	return nil
}

func (c *InterChainEventCache) Delete(imceType InterChainMessageEventType, nonce *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Delete(InterChainEventIndexKey(imceType, nonce))
	return err // the caller should handle the error
}

func (c *InterChainEventCache) Get(imceType InterChainMessageEventType, nonce *big.Int) (*InterChainMessageEvent, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	event := InterChainMessageEvent{}
	store := kvstore.NewKVStore(c.db)
	logger.Debugf("Getting %v event %v", imceType, nonce)
	err := store.Get(InterChainEventIndexKey(imceType, nonce), &event)
	return &event, err // the caller should handle the error
}

func (c *InterChainEventCache) Exists(imceType InterChainMessageEventType, nonce *big.Int) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	event := InterChainMessageEvent{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(InterChainEventIndexKey(imceType, nonce), &event)
	if err == nil {
		return true, nil
	}

	if err == ts.ErrKeyNotFound {
		return false, nil
	}

	return false, err // the caller should handle the error
}

// ------------------------------------ Getters and Setters for utility values --------------------------------------------

func (c *InterChainEventCache) GetLastQueryedHeightForType(icmeType InterChainMessageEventType) (*big.Int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	height := big.NewInt(0)
	store := kvstore.NewKVStore(c.db)
	err := store.Get(LastQueryedHeightKey(icmeType), &height)
	if err == nil {
		return height, nil
	}
	return big.NewInt(0), err
}

func (c *InterChainEventCache) SetLastQueryedHeightForType(icmeType InterChainMessageEventType, height *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(LastQueryedHeightKey(icmeType), height)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) GetNextVoucherBurnNonceForType(icmeType InterChainMessageEventType) (*big.Int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	nonce := big.NewInt(0)
	store := kvstore.NewKVStore(c.db)
	err := store.Get(InterChainVoucherBurnEventNextNonceKey(icmeType), &nonce)
	return nonce, err
}

func (c *InterChainEventCache) SetNextVoucherBurnNonceForType(icmeType InterChainMessageEventType, nonce *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(InterChainVoucherBurnEventNextNonceKey(icmeType), nonce)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) GetVoucherBurnStatus(icmeType InterChainMessageEventType, nonce *big.Int) (*VoucherBurnEventStatusInfo, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	statusInfo := VoucherBurnEventStatusInfo{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(VoucherBurnStatusInfoKey(icmeType, nonce), &statusInfo)
	return &statusInfo, err
}

func (c *InterChainEventCache) SetVoucherBurnStatus(statusInfo *VoucherBurnEventStatusInfo) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(VoucherBurnStatusInfoKey(statusInfo.Type, statusInfo.Nonce), statusInfo)
	return err // the caller should handle the error
}

func (c *InterChainEventCache) VoucherBurnNonceExists(icmeType InterChainMessageEventType, nonce *big.Int) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	statusInfo := VoucherBurnEventStatusInfo{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(VoucherBurnStatusInfoKey(icmeType, nonce), &statusInfo)
	if err == nil {
		return true, nil
	}

	if err == ts.ErrKeyNotFound {
		return false, nil
	}

	return false, err // the caller should handle the error
}

func (c *InterChainEventCache) GetNextTransferNonceForType(icmeType InterChainMessageEventType) (*big.Int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	nonce := big.NewInt(0)
	store := kvstore.NewKVStore(c.db)
	err := store.Get(InterChainTokenLockEventNextNonceKey(icmeType), &nonce)
	return nonce, err
}

func (c *InterChainEventCache) SetNextTransferNonceForType(icmeType InterChainMessageEventType, nonce *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(InterChainTokenLockEventNextNonceKey(icmeType), nonce)
	return err // the caller should handle the error
}
