package snapshot

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/rlp"
	"github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/theta/store/database/backend"
	"github.com/thetatoken/theta/store/kvstore"
	"github.com/thetatoken/theta/store/trie"
	sbc "github.com/thetatoken/thetasubchain/blockchain"
	sconsensus "github.com/thetatoken/thetasubchain/consensus"
	score "github.com/thetatoken/thetasubchain/core"
	sld "github.com/thetatoken/thetasubchain/ledger"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "snapshot"})

type SVStack []*slst.StoreView

func (s SVStack) push(sv *slst.StoreView) SVStack {
	return append(s, sv)
}

func (s SVStack) pop() (SVStack, *slst.StoreView) {
	l := len(s)
	if l == 0 {
		return s, nil
	}
	return s[:l-1], s[l-1]
}

func (s SVStack) peek() *slst.StoreView {
	l := len(s)
	if l == 0 {
		return nil
	}
	return s[l-1]
}

// ImportSnapshot loads the snapshot into the given database
func ImportSnapshot(snapshotFilePath, chainImportDirPath, chainCorrectionPath string, chain *sbc.Chain, db database.Database, ledger *sld.Ledger) (snapshotBlockHeader *score.BlockHeader, lastCC *score.ExtendedBlock, err error) {
	logger.Infof("Loading snapshot from: %v", snapshotFilePath)
	snapshotBlockHeader, metadata, err := loadSnapshot(snapshotFilePath, db, "Importing Snapshot")
	if err != nil {
		return nil, nil, err
	}
	logger.Infof("Snapshot loaded successfully.")

	// load previous chain, if any
	err = loadPrevChain(chainImportDirPath, snapshotBlockHeader, metadata, chain, db)
	if err != nil {
		return nil, nil, err
	}

	// load chain correction, if any
	if len(chainCorrectionPath) != 0 {
		headBlock, tailBlock, err := LoadChainCorrection(chainCorrectionPath, snapshotBlockHeader, metadata, chain, db, ledger)
		if err != nil {
			return nil, nil, err
		}

		snapshotBlock := score.ExtendedBlock{}
		kvstore := kvstore.NewKVStore(db)
		kvstore.Get(snapshotBlockHeader.Hash().Bytes(), &snapshotBlock)
		snapshotBlock.Children = []common.Hash{headBlock.Hash()}
		err = kvstore.Put(snapshotBlockHeader.Hash().Bytes(), snapshotBlock)
		if err != nil {
			return nil, nil, err
		}

		lastCC = tailBlock
	}

	return snapshotBlockHeader, lastCC, nil
}

// ValidateSnapshot validates the snapshot using a temporary database
func ValidateSnapshot(snapshotFilePath, chainImportDirPath, chainCorrectionPath string) (*score.BlockHeader, error) {
	logger.Infof("Verifying snapshot: %v", snapshotFilePath)

	tmpdbRoot, err := ioutil.TempDir("", "tmpdb")
	if err != nil {
		log.Panicf("Failed to create temporary db for snapshot verification: %v", err)
	}
	mainTmpDBPath := path.Join(tmpdbRoot, "main")
	refTmpDBPath := path.Join(tmpdbRoot, "ref")
	defer func() {
		os.RemoveAll(mainTmpDBPath)
		os.RemoveAll(refTmpDBPath)
	}()

	tmpdb, err := backend.NewLDBDatabase(mainTmpDBPath, refTmpDBPath, 256, 0)

	snapshotBlockHeader, metadata, err := loadSnapshot(snapshotFilePath, tmpdb, "Validating Snapshot")
	if err != nil {
		return nil, err
	}
	logger.Infof("Snapshot verified.")

	// load previous chain, if any
	err = loadPrevChain(chainImportDirPath, snapshotBlockHeader, metadata, nil, tmpdb)
	if err != nil {
		return nil, err
	}

	// load chain correction, if any
	if len(chainCorrectionPath) != 0 {
		headBlock, _, err := LoadChainCorrection(chainCorrectionPath, snapshotBlockHeader, metadata, nil, tmpdb, nil)
		if err != nil {
			return nil, err
		}

		snapshotBlock := score.ExtendedBlock{}
		kvstore := kvstore.NewKVStore(tmpdb)
		kvstore.Get(snapshotBlockHeader.Hash().Bytes(), &snapshotBlock)
		snapshotBlock.Children = []common.Hash{headBlock.Hash()}
		err = kvstore.Put(snapshotBlockHeader.Hash().Bytes(), snapshotBlock)
		if err != nil {
			return nil, err
		}
	}

	return snapshotBlockHeader, nil
}

func LoadSnapshotCheckpointHeader(snapshotFilePath string) *score.BlockHeader {
	var err error

	snapshotFile, err := os.Open(snapshotFilePath)
	if err != nil {
		return nil
	}
	defer snapshotFile.Close()

	snapshotHeader := &score.SnapshotHeader{}
	_, err = score.ReadRecord(snapshotFile, snapshotHeader)
	if err != nil {
		return nil
	}

	lastCheckpoint := score.LastCheckpoint{}
	_, err = score.ReadRecord(snapshotFile, &lastCheckpoint)
	if err != nil {
		return nil
	}

	metadata := score.SnapshotMetadata{}
	_, err = score.ReadRecord(snapshotFile, &metadata)
	if err != nil {
		return nil
	}

	return metadata.TailTrio.Second.Header
}

func loadSnapshot(snapshotFilePath string, db database.Database, logStr string) (*score.BlockHeader, *score.SnapshotMetadata, error) {
	var err error

	snapshotFile, err := os.Open(snapshotFilePath)
	if err != nil {
		return nil, nil, err
	}
	defer snapshotFile.Close()

	kvstore := kvstore.NewKVStore(db)

	// ------------------------------ Load State ------------------------------ //

	snapshotVersion := uint(1)
	snapshotHeader := &score.SnapshotHeader{}
	_, err = score.ReadRecord(snapshotFile, snapshotHeader)
	if err != nil || snapshotHeader.Magic != score.SnapshotHeaderMagic { // older version, reset snapshotFile
		snapshotFile.Seek(0, 0)
	} else {
		snapshotVersion = snapshotHeader.Version
	}

	logger.Infof("Reading snapshot header, version: %v, magic: %v", snapshotVersion, snapshotHeader.Magic)

	lastCheckpoint := score.LastCheckpoint{}
	if snapshotVersion >= 2 {
		_, err = score.ReadRecord(snapshotFile, &lastCheckpoint)
		if err != nil {
			return nil, nil, fmt.Errorf("Failed to load snapshot last checkpoint, %v", err)
		}

		ckb := score.Block{
			BlockHeader: lastCheckpoint.CheckpointHeader,
		}
		eckb := score.ExtendedBlock{
			Block:  &ckb,
			Status: score.BlockStatusTrusted, // HCC links between all three blocks
		}
		ckbHash := ckb.BlockHeader.Hash()

		existingCkbExt := score.ExtendedBlock{}
		if kvstore.Get(ckbHash[:], &existingCkbExt) != nil {
			logger.Infof("Saving the last checkpoint block: %v", ckbHash.Hex())
			err = kvstore.Put(ckbHash[:], &eckb)
			if err != nil {
				logger.Panicf("Failed to save the last checkpoint: %v, err: %v", ckbHash.Hex(), err)
			}
		}

		for _, intermediateHeader := range lastCheckpoint.IntermediateHeaders {
			ibHash := intermediateHeader.Hash()
			eib := score.ExtendedBlock{
				Block: &score.Block{BlockHeader: intermediateHeader},
			}
			existingEib := score.ExtendedBlock{}
			if kvstore.Get(ibHash[:], &existingEib) != nil {
				logger.Debugf("Saving intermediate blocks: %v", ibHash.Hex())
				err = kvstore.Put(ibHash[:], &eib)
				if err != nil {
					logger.Panicf("Failed to save ntermediate block: %v, err: %v", ibHash.Hex(), err)
				}
			}
		}
	}

	metadata := score.SnapshotMetadata{}
	_, err = score.ReadRecord(snapshotFile, &metadata)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to load snapshot metadata, %v", err)
	}

	fileInfo, err := os.Stat(snapshotFilePath)
	var fileSize uint64
	if err == nil {
		fileSize = uint64(fileInfo.Size()) / 100
	}

	var sv *slst.StoreView
	if snapshotHeader.Version >= 3 {
		err = loadStateV3(snapshotFile, db, fileSize, logStr)
		if err != nil {
			return nil, nil, err
		}
		lfb := metadata.TailTrio.Second
		sv = slst.NewStoreView(lfb.Header.Height, lfb.Header.StateHash, db)
	} else {
		sv, _, err = loadStateV2(snapshotFile, db, fileSize, logStr)
		if err != nil {
			return nil, nil, err
		}
	}

	// ----------------------------- Validity Checks -------------------------- //

	if snapshotVersion >= 4 {
		if err = checkSnapshotV4(sv, &metadata, db); err != nil {
			return nil, nil, fmt.Errorf("Snapshot state validation failed: %v", err)
		}
	} else {
		if err = checkSnapshot(sv, &metadata, db); err != nil {
			return nil, nil, fmt.Errorf("Snapshot state validation failed: %v", err)
		}
	}

	// --------------------- Save Proofs and Tail Blocks  --------------------- //

	for _, blockTrio := range metadata.ProofTrios {
		blockTrioKey := []byte(score.BlockTrioStoreKeyPrefix + strconv.FormatUint(blockTrio.First.Header.Height, 10))
		err = kvstore.Put(blockTrioKey, blockTrio)
		if err != nil {
			logger.Panicf("Failed to save ProofTrios: err: %v", err)
		}
	}

	secondBlockHeader := saveTailBlocks(&metadata, sv, kvstore)

	// ----------------------------- More Validity Checks -------------------------- //

	if snapshotVersion >= 2 {
		if err = checkLastCheckpoint(sv, secondBlockHeader, &lastCheckpoint, db); err != nil {
			return nil, nil, fmt.Errorf("Snapshot last checkpoint validation failed: %v", err)
		}
	}

	// Testing
	valSet := sv.GetValidatorSet()
	numValidators := valSet.Size()
	logger.Infof("Number of validators loaded: %v", numValidators)
	for _, v := range valSet.Validators() {
		logger.Infof("Validator: %v, stake: %v", v.Address.Hex(), v.Stake)
	}

	return secondBlockHeader, &metadata, nil
}

func LoadChainCorrection(chainImportDirPath string, snapshotBlockHeader *score.BlockHeader, metadata *score.SnapshotMetadata, chain *sbc.Chain, db database.Database, ledger *sld.Ledger) (headBlock, tailBlock *score.ExtendedBlock, err error) {
	chainFile, err := os.Open(chainImportDirPath)
	if err != nil {
		return
	}
	defer chainFile.Close()

	kvstore := kvstore.NewKVStore(db)

	var count uint64
	var prevBlock *score.ExtendedBlock
	blockStack := make([]*score.ExtendedBlock, 0)
	for {
		backupBlock := &score.BackupBlock{}
		_, err := score.ReadRecord(chainFile, backupBlock)
		if err != nil {
			if err == io.EOF {
				if prevBlock.Height != snapshotBlockHeader.Height+1 {
					return nil, nil, fmt.Errorf("Chain's head block height: %v, snapshot's height: %v", prevBlock.Height, snapshotBlockHeader.Height)
				}
				break
			}
			return nil, nil, fmt.Errorf("Failed to read backup record, %v", err)
		}
		block := backupBlock.Block

		if count == 0 {
			tailBlock = block
		}

		if block.Height <= snapshotBlockHeader.Height {
			return nil, nil, fmt.Errorf("Block height is %v, must be > than snapshot height %v", block.Height, snapshotBlockHeader.Height)
		}

		blockHash := block.Hash()

		if prevBlock != nil {
			// check chaining
			if block.Height != prevBlock.Height-1 {
				return nil, nil, fmt.Errorf("Block height: %v, prev block height: %v", block.Height, prevBlock.Height)
			}
			if prevBlock.Parent != blockHash {
				return nil, nil, fmt.Errorf("Block at height %v has invalid parent %v vs %v", prevBlock.Height, prevBlock.Parent, blockHash)
			}
		}

		if chain != nil {
			if block.ChainID != chain.ChainID {
				return nil, nil, errors.Errorf("ChainID mismatch: block.ChainID(%s) != %s", block.ChainID, chain.ChainID)
			}
			existingBlock := score.ExtendedBlock{}
			if kvstore.Get(blockHash[:], &existingBlock) != nil {
				block.Status = score.BlockStatusTrusted
				err = kvstore.Put(blockHash[:], block)
				if err != nil {
					return nil, nil, err
				}
				chain.AddBlockByHeightIndex(block.Height, blockHash)
				chain.AddTxsToIndex(block, true)
			} else {
				existingBlock.Txs = block.Txs
				existingBlock.Status = score.BlockStatusTrusted
				err = kvstore.Put(blockHash[:], existingBlock)
				if err != nil {
					return nil, nil, err
				}
				chain.AddBlockByHeightIndex(block.Height, blockHash)
				chain.AddTxsToIndex(block, true)
			}
		}

		count++
		prevBlock = block
		blockStack = append(blockStack, block)
	}

	if ledger != nil {
		for {
			num := len(blockStack)
			if num == 0 {
				break
			}
			block := blockStack[num-1]
			blockStack = blockStack[:num-1]
			parent, _ := chain.FindBlock(block.Parent)

			//result := ledger.ResetState(parent.Height, parent.StateHash)
			result := ledger.ResetState(parent.Block)
			if result.IsError() {
				return nil, nil, fmt.Errorf("%v", result.String())
			}

			_, result = ledger.ApplyBlockTxsForChainCorrection(block.Block)
			if result.IsError() {
				return nil, nil, fmt.Errorf("%v", result.String())
			}
		}
	}

	headBlock = prevBlock
	return
}

func loadPrevChain(chainImportDirPath string, snapshotBlockHeader *score.BlockHeader, metadata *score.SnapshotMetadata, chain *sbc.Chain, db database.Database) error {
	if len(chainImportDirPath) != 0 {
		if _, err := os.Stat(chainImportDirPath); !os.IsNotExist(err) {
			fileInfos, err := ioutil.ReadDir(chainImportDirPath)
			if err != nil {
				return fmt.Errorf("Failed to read chain import directory %v: %v", chainImportDirPath, err)
			}

			var chainFileNames []string
			for _, fileInfo := range fileInfos {
				chainFileNames = append(chainFileNames, fileInfo.Name())
			}

			sort.Slice(chainFileNames, func(i, j int) bool {
				start1, _ := getChainBoundary(chainFileNames[i])
				start2, _ := getChainBoundary(chainFileNames[j])
				return start1 > start2
			})

			var blockEnd uint64
			var prevBlock *score.ExtendedBlock
			for _, fileName := range chainFileNames {
				start, end := getChainBoundary(fileName)
				if start > snapshotBlockHeader.Height {
					continue
				}
				if end < snapshotBlockHeader.Height {
					if end != blockEnd {
						return fmt.Errorf("Missing chain file ending at height %v", blockEnd)
					}
				}

				chainFilePath := path.Join(chainImportDirPath, fileName)
				prevBlock, err = loadChainSegment(chainFilePath, start, end, prevBlock, snapshotBlockHeader, metadata, chain, db)
				if err != nil {
					return err
				}
				blockEnd = start - 1
			}

			start, _ := getChainBoundary(chainFileNames[len(chainFileNames)-1])
			if prevBlock.Height != start {
				return fmt.Errorf("Chain loading started at height %v, but should start at height %v", prevBlock.Height, start)
			}

			logger.Infof("Chain loaded successfully.")
		}
	}
	return nil
}

func loadChainSegment(filePath string, start, end uint64, prevBlock *score.ExtendedBlock, snapshotBlockHeader *score.BlockHeader, metadata *score.SnapshotMetadata, chain *sbc.Chain, db database.Database) (*score.ExtendedBlock, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	kvstore := kvstore.NewKVStore(db)

	var count uint64
	var proofTrio, prevTrio score.SnapshotBlockTrio
	for {
		backupBlock := &score.BackupBlock{}
		_, err := score.ReadRecord(file, backupBlock)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("Failed to read backup record, %v", err)
		}
		block := backupBlock.Block

		if count == 0 {
			if block.Height != end {
				return nil, fmt.Errorf("Last block height doesn't match, %v : %v", block.Height, end)
			}
		}

		if block.Height > snapshotBlockHeader.Height {
			count++
			continue
		}

		// check block itself
		var provenValSet *score.ValidatorSet
		if block.Height == score.GenesisBlockHeight {
			provenValSet, err = checkGenesisBlock(block.BlockHeader, db)
			if err != nil {
				return nil, err
			}
		} else {
			if chain != nil {
				if res := block.Validate(chain.ChainID); res.IsError() {
					return nil, fmt.Errorf("Block %v's header is invalid, %v", block.Height, res)
				}
			}

			for {
				proofTrio = metadata.ProofTrios[len(metadata.ProofTrios)-1]
				if proofTrio.First.Header.Height+2 <= block.Height || len(metadata.ProofTrios) == 1 {
					break
				}
				provenValSet = nil
				metadata.ProofTrios = metadata.ProofTrios[:len(metadata.ProofTrios)-1]
				prevTrio = proofTrio
			}

			if provenValSet == nil {
				if proofTrio.First.Header.Height == score.GenesisBlockHeight {
					provenValSet, err = checkGenesisBlock(proofTrio.Second.Header, db)
				} else {
					provenValSet, err = getValidatorSetFromVSProof(proofTrio.First.Header.StateHash, &proofTrio.First.Proof)
				}
				if err != nil {
					return nil, fmt.Errorf("Failed to retrieve validator set from VS proof: %v", err)
				}
			}
		}

		// check votes
		if err := validateVotes(provenValSet, block.BlockHeader, backupBlock.Votes); err != nil {
			return nil, fmt.Errorf("Failed to validate voteSet, %v", err)
		}

		blockHash := block.Hash()

		// check against snapshot block
		if block.Height == snapshotBlockHeader.Height {
			if blockHash != snapshotBlockHeader.Hash() {
				return nil, fmt.Errorf("Chain loading reached snapshot block, but block hash doesn't match, %v : %v", blockHash, snapshotBlockHeader.Hash())
			}
		} else {
			// check chaining
			if block.Height != prevBlock.Height-1 {
				return nil, fmt.Errorf("Block height missing at %v", prevBlock.Height-1)
			}
			if prevBlock.Parent != blockHash {
				return nil, fmt.Errorf("Block at height %v has invalid parent %v vs %v", prevBlock.Height, prevBlock.Parent, blockHash)
			}

			// check against genesis block
			if block.Height == score.GenesisBlockHeight {
				_, err := checkGenesisBlock(block.BlockHeader, db)
				if err != nil {
					return nil, err
				}
			}
		}

		if chain != nil {
			if block.ChainID != chain.ChainID {
				return nil, errors.Errorf("ChainID mismatch: block.ChainID(%s) != %s", block.ChainID, chain.ChainID)
			}
			existingBlock := score.ExtendedBlock{}
			if kvstore.Get(blockHash[:], &existingBlock) != nil {
				err = kvstore.Put(blockHash[:], block)
				if err != nil {
					return nil, err
				}
				chain.AddBlockByHeightIndex(block.Height, blockHash)
				chain.AddTxsToIndex(block, true)
			} else {
				if block.Height == score.GenesisBlockHeight+1 || block.Height == snapshotBlockHeader.Height || block.Height == snapshotBlockHeader.Height-1 || block.Height == prevTrio.First.Header.Height {
					existingBlock.Txs = block.Txs
					existingBlock.HasValidatorUpdate = true
					err = kvstore.Put(blockHash[:], existingBlock)
					if err != nil {
						return nil, err
					}
					chain.AddBlockByHeightIndex(block.Height, blockHash)
					chain.AddTxsToIndex(block, true)
				}
			}
		}

		count++
		prevBlock = block
	}

	if prevBlock == nil {
		return nil, fmt.Errorf("Empty chain file for height %v to %v", start, end)
	}
	if prevBlock.Height != start {
		return nil, fmt.Errorf("Chain file for height %v to %v has first block height %v", start, end, prevBlock.Height)
	}
	if count != end-start+1 {
		return nil, fmt.Errorf("Expect %v blocks, but actually got %v", end-start+1, count)
	}

	return prevBlock, nil
}

func getChainBoundary(filename string) (start, end uint64) {
	filename = filename[len("theta_chain-"):]
	idx := strings.Index(filename, "-")
	startStr := filename[:idx]
	start, _ = strconv.ParseUint(startStr, 10, 64)
	filename = filename[idx+1:]
	filename = filename[:strings.Index(filename, "-")]
	end, _ = strconv.ParseUint(filename, 10, 64)
	return
}

func loadStateV2(file *os.File, db database.Database, fileSize uint64, logStr string) (*slst.StoreView, common.Hash, error) {
	var hash common.Hash
	var sv *slst.StoreView
	var account *types.Account
	svStack := make(SVStack, 0)
	var progress, curSize uint64
	for {
		record := score.SnapshotTrieRecord{}
		recordSize, err := score.ReadRecord(file, &record)
		if err != nil {
			if err == io.EOF {
				if svStack.peek() != nil {
					return nil, common.Hash{}, fmt.Errorf("Still some storeview unhandled")
				}
				break
			}
			return nil, common.Hash{}, fmt.Errorf("Failed to read snapshot record, %v", err)
		}

		if fileSize > 0 {
			curSize += recordSize
			percentage := curSize / fileSize
			if percentage > progress && percentage <= 100 && percentage%5 == 0 {
				logger.Infof("%s, %v%% done.", logStr, percentage)
				progress = percentage
			}
		}

		if bytes.Equal(record.K, []byte{score.SVStart}) {
			height := score.Bytestoi(record.V)
			sv := slst.NewStoreView(height, common.Hash{}, db)
			svStack = svStack.push(sv)
		} else if bytes.Equal(record.K, []byte{score.SVEnd}) {
			svStack, sv = svStack.pop()
			if sv == nil {
				return nil, common.Hash{}, fmt.Errorf("Missing storeview to handle")
			}
			height := score.Bytestoi(record.V)
			if height != sv.Height() {
				return nil, common.Hash{}, fmt.Errorf("Storeview start and end heights don't match")
			}
			hash = sv.Save()

			if svStack.peek() != nil && height == svStack.peek().Height() {
				// it's a storeview for account storage, verify account
				if account.Root != hash {
					return nil, common.Hash{}, fmt.Errorf("Account storage root doesn't match")
				}
			}
			account = nil
		} else {
			sv := svStack.peek()
			if sv == nil {
				return nil, common.Hash{}, fmt.Errorf("Missing storeview to handle")
			}
			sv.Set(record.K, record.V)

			if account == nil {
				if bytes.HasPrefix(record.K, []byte("ls/a")) {
					acct := &types.Account{}
					err = types.FromBytes([]byte(record.V), acct)
					if err != nil {
						return nil, common.Hash{}, fmt.Errorf("Failed to parse account, %v", err)
					}
					if acct.Root != (common.Hash{}) {
						account = acct
					}
				}
			}
		}
	}
	logger.Infof("%s, 100%% done.", logStr)

	return sv, hash, nil
}

func loadStateV3(file *os.File, db database.Database, fileSize uint64, logStr string) error {
	var progress, curSize uint64
	batch := db.NewBatch()
	record := score.SnapshotTrieRecord{}
	for {
		recordSize, err := score.ReadRecord(file, &record)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("Failed to read snapshot record, %v", err)
		}

		if fileSize > 0 {
			curSize += recordSize
			percentage := curSize / fileSize
			if percentage > progress && percentage <= 100 && percentage%5 == 0 {
				logger.Infof("%s, %v%% done.", logStr, percentage)
				progress = percentage
			}
		}

		err = batch.Put(record.K, record.V)
		if err != nil {
			return fmt.Errorf("Failed to write snapshot record, %v", err)
		}

		// Set the ref count to 3 to be conservative as we have 3 state tries in the snapshot
		for i := 0; i < 3; i++ {
			err = batch.Reference(record.K)
			if err != nil {
				return fmt.Errorf("Failed to create reference of snapshot record, %v", err)
			}
		}

		if batch.ValueSize() > database.IdealBatchSize {
			if err := batch.Write(); err != nil {
				return err
			}
			batch.Reset()
		}
	}
	if err := batch.Write(); err != nil {
		return err
	}

	logger.Infof("%s, 100%% done.", logStr)

	return nil
}

func checkLastCheckpoint(sv *slst.StoreView, snapshotBlockHeader *score.BlockHeader, lastCheckpoint *score.LastCheckpoint, db database.Database) error {
	if snapshotBlockHeader == nil {
		return fmt.Errorf("The snapshot block header is nil")
	}

	lastCheckpointHeader := lastCheckpoint.CheckpointHeader
	if lastCheckpointHeader == nil {
		return fmt.Errorf("The last checkpoint header is nil")
	}

	// Verify that the snapshot block is a descendent of the last checkpoint
	store := kvstore.NewKVStore(db)
	lastCheckpointHash := lastCheckpointHeader.Hash()
	snapshotBlockHash := snapshotBlockHeader.Hash()

	isDescendent := false
	hash := snapshotBlockHash
	for i := int64(0); i < common.CheckpointInterval; i++ {
		if hash == lastCheckpointHash {
			isDescendent = true
			break
		}

		var block score.ExtendedBlock
		err := store.Get(hash[:], &block)
		if err != nil {
			return fmt.Errorf("Failed to find block when verifying the last checkpoint, block hash: %v", hash.Hex())
		}
		hash = block.Parent
	}

	if !isDescendent {
		return fmt.Errorf("Snapshot block %v is not a descendant of the last checkpoint %v",
			snapshotBlockHash.Hex(), lastCheckpointHash.Hex())
	}

	return nil
}

func checkSnapshot(sv *slst.StoreView, metadata *score.SnapshotMetadata, db database.Database) error {
	tailTrio := &metadata.TailTrio
	secondBlock := tailTrio.Second.Header
	expectedStateHash := sv.Hash()
	if bytes.Compare(expectedStateHash.Bytes(), secondBlock.StateHash.Bytes()) != 0 {
		return fmt.Errorf("StateHash not matching: %v vs %s",
			expectedStateHash.Hex(), secondBlock.StateHash.Hex())
	}

	var provenValSet *score.ValidatorSet
	var err error
	if secondBlock.Height != score.GenesisBlockHeight {
		provenValSet, err = checkProofTrios(metadata.ProofTrios, db)
		if err != nil {
			return err
		}
	}

	err = checkTailTrio(sv, provenValSet, tailTrio)
	if err != nil {
		return err
	}

	return nil
}

func checkSnapshotV4(sv *slst.StoreView, metadata *score.SnapshotMetadata, db database.Database) error {
	tailTrio := &metadata.TailTrio
	secondBlock := tailTrio.Second.Header
	expectedStateHash := sv.Hash()
	if bytes.Compare(expectedStateHash.Bytes(), secondBlock.StateHash.Bytes()) != 0 {
		return fmt.Errorf("StateHash not matching: %v vs %s",
			expectedStateHash.Hex(), secondBlock.StateHash.Hex())
	}

	var valSet *score.ValidatorSet
	var err error

	first := tailTrio.First
	valSet, err = getValidatorSetFromVSProof(first.Header.StateHash, &first.Proof)
	if err != nil {
		return fmt.Errorf("Failed to retrieve validator set from VS proof: %v", err)
	}

	logger.Infof("Validators of snapshost: %v", valSet)

	err = checkTailTrio(sv, valSet, tailTrio)
	if err != nil {
		return err
	}

	return nil
}

func checkProofTrios(proofTrios []score.SnapshotBlockTrio, db database.Database) (*score.ValidatorSet, error) {
	logger.Debugf("Check validator set change proofs...")

	var provenValSet *score.ValidatorSet // the proven validator set so far
	var err error
	for idx, blockTrio := range proofTrios {
		first := blockTrio.First
		second := blockTrio.Second
		third := blockTrio.Third
		if idx == 0 {
			// special handling for the genesis block
			provenValSet, err = checkGenesisBlock(second.Header, db)
			if err != nil {
				return nil, fmt.Errorf("Invalid genesis block: %v", err)
			}
		} else {
			if second.Header.Parent != first.Header.Hash() || third.Header.Parent != second.Header.Hash() {
				return nil, fmt.Errorf("block trio has invalid Parent link")
			}

			if second.Header.HCC.BlockHash != first.Header.Hash() || third.Header.HCC.BlockHash != second.Header.Hash() {
				return nil, fmt.Errorf("block trio has invalid HCC link: %v, %v; %v, %v", first.Header.Hash(), second.Header.HCC.BlockHash,
					second.Header.Hash(), third.Header.HCC.BlockHash)
			}

			// third.Header.HCC.Votes contains the votes for the second block in the trio
			if err := validateVotes(provenValSet, second.Header, third.Header.HCC.Votes); err != nil {
				return nil, fmt.Errorf("Failed to validate voteSet, %v", err)
			}
			provenValSet, err = getValidatorSetFromVSProof(first.Header.StateHash, &first.Proof)
			if err != nil {
				return nil, fmt.Errorf("Failed to retrieve validator set from VS proof: %v", err)
			}
		}

		logger.Debugf("Block height: %v, Currently proven validator set: %v", first.Header.Height, provenValSet)
	}

	return provenValSet, nil
}

func checkTailTrio(sv *slst.StoreView, provenValSet *score.ValidatorSet, tailTrio *score.SnapshotBlockTrio) error {
	second := &tailTrio.Second
	third := &tailTrio.Third

	if second.Header.Height == score.GenesisBlockHeight {
		_, err := checkGenesisBlock(second.Header, sv.GetDB())
		if err != nil {
			return err
		}
	} else {
		validateVotes(provenValSet, third.Header, third.VoteSet)
		retrievedValSet := getValidatorSetFromSV(sv)
		if !provenValSet.Equals(retrievedValSet) {
			return fmt.Errorf("The latest proven and retrieved validator set does not match")
		}
	}

	return nil
}

func checkGenesisBlock(block *score.BlockHeader, db database.Database) (*score.ValidatorSet, error) {
	if block.Height != score.GenesisBlockHeight {
		return nil, fmt.Errorf("Invalid genesis block height: %v", block.Height)
	}

	var expectedGenesisHash string
	if block.ChainID == score.MainnetChainID {
		expectedGenesisHash = score.MainnetGenesisBlockHash
	} else {
		expectedGenesisHash = viper.GetString(common.CfgGenesisHash)
	}

	// logger.Infof("Expected genesis hash: %v", expectedGenesisHash)
	// logger.Infof("Acutal   genesis hash: %v", block.Hash().Hex())

	if block.Hash() != common.HexToHash(expectedGenesisHash) {
		return nil, fmt.Errorf("Genesis block hash mismatch, expected: %v, calculated: %v",
			expectedGenesisHash, block.Hash().Hex())
	}

	// now that the block hash matches with the expected genesis block hash,
	// the block and its state trie is considerred valid. We can retrieve the
	// genesis validator set from its state trie
	gsv := slst.NewStoreView(block.Height, block.StateHash, db)

	genesisValidatorSet := getValidatorSetFromSV(gsv)

	return genesisValidatorSet, nil
}

func getValidatorSetFromVSProof(stateHash common.Hash, recoverredVSP *score.ValidatorSetProof) (*score.ValidatorSet, error) {
	serializedVS, _, err := trie.VerifyProof(stateHash, slst.CurrentValidatorSetKey(), recoverredVSP)
	if err != nil {
		return nil, err
	}

	vs := &score.ValidatorSet{}
	err = rlp.DecodeBytes(serializedVS, vs)
	if err != nil {
		return nil, err
	}
	return sconsensus.FilterValidators(vs), nil
}

func getValidatorSetFromSV(sv *slst.StoreView) *score.ValidatorSet {
	vs := sv.GetValidatorSet()
	return sconsensus.FilterValidators(vs)
}

func validateVotes(validatorSet *score.ValidatorSet, block *score.BlockHeader, voteSet *score.VoteSet) error {
	if !validatorSet.HasMajority(voteSet) {
		return fmt.Errorf("block doesn't have majority votes")
	}
	for _, vote := range voteSet.Votes() {
		res := vote.Validate()
		if !res.IsOK() {
			return fmt.Errorf("vote is not valid, %v", res)
		}
		if vote.Block != block.Hash() {
			return fmt.Errorf("vote is not for corresponding block")
		}
		_, err := validatorSet.GetValidator(vote.ID)
		if err != nil {
			return fmt.Errorf("can't find validator for vote")
		}
	}
	return nil
}

func saveTailBlocks(metadata *score.SnapshotMetadata, sv *slst.StoreView, kvstore store.Store) *score.BlockHeader {
	tailBlockTrio := &metadata.TailTrio
	firstBlock := score.Block{BlockHeader: tailBlockTrio.First.Header}
	secondBlock := score.Block{BlockHeader: tailBlockTrio.Second.Header}
	hl := sv.GetValidatorSetUpdateTxHeightList()

	if secondBlock.Height != score.GenesisBlockHeight {
		firstExt := &score.ExtendedBlock{
			Block:              &firstBlock,
			Status:             score.BlockStatusTrusted, // HCC links between all three blocks
			Children:           []common.Hash{secondBlock.Hash()},
			HasValidatorUpdate: hl.Contains(firstBlock.Height),
		}
		firstBlockHash := firstBlock.BlockHeader.Hash()

		existingFirstExt := score.ExtendedBlock{}
		if kvstore.Get(firstBlockHash[:], &existingFirstExt) != nil {
			logger.Infof("Saving the the first snapshot tail block: %v", firstBlockHash.Hex())
			err := kvstore.Put(firstBlockHash[:], firstExt)
			if err != nil {
				logger.Panicf("Failed to save the first snapshot tail block: %v, err: %v", firstBlockHash.Hex(), err)
			}
		}
	}

	secondExt := &score.ExtendedBlock{
		Block:              &secondBlock,
		Status:             score.BlockStatusTrusted,
		Children:           []common.Hash{},
		HasValidatorUpdate: hl.Contains(secondBlock.Height),
	}
	secondBlockHash := secondBlock.BlockHeader.Hash()

	existingSecondExt := score.ExtendedBlock{}
	if kvstore.Get(secondBlockHash[:], &existingSecondExt) != nil {
		logger.Infof("Saving the the second snapshot tail block: %v", secondBlockHash.Hex())
		err := kvstore.Put(secondBlockHash[:], secondExt)
		if err != nil {
			logger.Panicf("Failed to save the second snapshot tail block: %v, err: %v", secondBlockHash.Hex(), err)
		}
	}

	if secondExt.Height != score.GenesisBlockHeight && secondExt.HasValidatorUpdate {
		// TODO: this would lead to mismatch between the proven and retrieved validator set,
		//       need to handle this case properly
		logger.Warnf("The second block in the tail trio contains validator update, may cause valSet mismatch, height: %v", secondBlock.Height)
	}

	return secondBlock.BlockHeader
}
