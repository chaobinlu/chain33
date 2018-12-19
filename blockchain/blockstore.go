// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"

<<<<<<< HEAD
	"github.com/golang/protobuf/proto"
=======
>>>>>>> upstream/master
	"github.com/33cn/chain33/common"
	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/common/difficulty"
	"github.com/33cn/chain33/common/version"
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
<<<<<<< HEAD
)

=======
	"github.com/golang/protobuf/proto"
)

//var
>>>>>>> upstream/master
var (
	blockLastHeight       = []byte("blockLastHeight")
	bodyPerfix            = []byte("Body:")
	LastSequence          = []byte("LastSequence")
	headerPerfix          = []byte("Header:")
	heightToHeaderPerfix  = []byte("HH:")
	hashPerfix            = []byte("Hash:")
	tdPerfix              = []byte("TD:")
	heightToHashKeyPerfix = []byte("Height:")
	seqToHashKey          = []byte("Seq:")
	HashToSeqPerfix       = []byte("HashToSeq:")
<<<<<<< HEAD

	storeLog       = chainlog.New("submodule", "store")
	lastheaderlock sync.Mutex
	AddBlock       int64 = 1
	DelBlock       int64 = 2
)

=======
	seqCBPrefix           = []byte("SCB:")
	seqCBLastNumPrefix    = []byte("SCBL:")
	storeLog              = chainlog.New("submodule", "store")
	lastheaderlock        sync.Mutex
	AddBlock              int64 = 1
	DelBlock              int64 = 2
)

//GetLocalDBKeyList 获取本地键值列表
>>>>>>> upstream/master
func GetLocalDBKeyList() [][]byte {
	return [][]byte{
		blockLastHeight, bodyPerfix, LastSequence, headerPerfix, heightToHeaderPerfix,
		hashPerfix, tdPerfix, heightToHashKeyPerfix, seqToHashKey, HashToSeqPerfix,
	}
}

//存储block hash对应的blockbody信息
func calcHashToBlockBodyKey(hash []byte) []byte {
	return append(bodyPerfix, hash...)
}

<<<<<<< HEAD
=======
//并发访问的可能性(每次开辟新内存)
func calcSeqCBKey(name []byte) []byte {
	return append(append([]byte{}, seqCBPrefix...), name...)
}

//并发访问的可能性(每次开辟新内存)
func caclSeqCBLastNumKey(name []byte) []byte {
	return append(append([]byte{}, seqCBLastNumPrefix...), name...)
}

>>>>>>> upstream/master
//存储block hash对应的header信息
func calcHashToBlockHeaderKey(hash []byte) []byte {
	return append(headerPerfix, hash...)
}

func calcHeightToBlockHeaderKey(height int64) []byte {
	return append(heightToHeaderPerfix, []byte(fmt.Sprintf("%012d", height))...)
}

//存储block hash对应的block height
func calcHashToHeightKey(hash []byte) []byte {
	return append(hashPerfix, hash...)
}

//存储block hash对应的block总难度TD
func calcHashToTdKey(hash []byte) []byte {
	return append(tdPerfix, hash...)
}

//存储block height 对应的block  hash
func calcHeightToHashKey(height int64) []byte {
	return append(heightToHashKeyPerfix, []byte(fmt.Sprintf("%v", height))...)
}

//存储block操作序列号对应的block hash,KEY=Seq:sequence
func calcSequenceToHashKey(sequence int64) []byte {
	return append(seqToHashKey, []byte(fmt.Sprintf("%v", sequence))...)
}

//存储block hash对应的seq序列号，KEY=Seq:sequence，只用于平行链addblock操作，方便delblock回退是查找对应seq的hash
func calcHashToSequenceKey(hash []byte) []byte {
	return append(HashToSeqPerfix, hash...)
}

<<<<<<< HEAD
=======
//BlockStore 区块存储
>>>>>>> upstream/master
type BlockStore struct {
	db        dbm.DB
	client    queue.Client
	height    int64
	lastBlock *types.Block
}

<<<<<<< HEAD
=======
//NewBlockStore new
>>>>>>> upstream/master
func NewBlockStore(db dbm.DB, client queue.Client) *BlockStore {
	height, err := LoadBlockStoreHeight(db)
	if err != nil {
		chainlog.Info("init::LoadBlockStoreHeight::database may be crash", "err", err.Error())
		if err != types.ErrHeightNotExist {
			panic(err)
		}
	}
	blockStore := &BlockStore{
		height: height,
		db:     db,
		client: client,
	}
	if height == -1 {
		chainlog.Info("load block height error, may be init database", "height", height)
		if types.IsEnable("quickIndex") {
			blockStore.saveQuickIndexFlag()
		}
	} else {
		blockdetail, err := blockStore.LoadBlockByHeight(height)
		if err != nil {
			chainlog.Error("init::LoadBlockByHeight::database may be crash")
			panic(err)
		}
		blockStore.lastBlock = blockdetail.GetBlock()
		flag, err := blockStore.loadFlag(types.FlagTxQuickIndex)
		if err != nil {
			panic(err)
		}
		if types.IsEnable("quickIndex") {
			if flag == 0 {
				blockStore.initQuickIndex(height)
			}
		} else {
			if flag != 0 {
				panic("toml config disable tx quick index, but database enable quick index")
			}
		}
	}
	return blockStore
}

//步骤:
//检查数据库是否已经进行quickIndex改造
//如果没有，那么进行下面的步骤
//1. 先把hash 都给改成 TX:hash
//2. 把所有的 Tx:hash 都加一个 8字节的index
//3. 10000个区块 处理一次，并且打印进度
//4. 全部处理完成了,添加quickIndex 的标记
func (bs *BlockStore) initQuickIndex(height int64) {
	batch := bs.db.NewBatch(true)
	var count int
	for i := int64(0); i <= height; i++ {
		blockdetail, err := bs.LoadBlockByHeight(i)
		if err != nil {
			panic(err)
		}
		for _, tx := range blockdetail.Block.Txs {
			hash := tx.Hash()
			txresult, err := bs.db.Get(hash)
			if err != nil {
				panic(err)
			}
			count++
			batch.Set(types.CalcTxKey(hash), txresult)
			batch.Set(types.CalcTxShortKey(hash), []byte("1"))
		}
		if count > 100000 {
			storeLog.Info("initQuickIndex", "height", i)
			err := batch.Write()
			if err != nil {
				panic(err)
			}
			batch = bs.db.NewBatch(true)
			count = 0
		}
	}
	if count > 0 {
		err := batch.Write()
		if err != nil {
			panic(err)
		}
		storeLog.Info("initQuickIndex", "height", height)
	}
	bs.saveQuickIndexFlag()
}

<<<<<<< HEAD
=======
func (bs *BlockStore) isSeqCBExist(name string) bool {
	value, err := bs.db.Get(calcSeqCBKey([]byte(name)))
	if err == nil {
		var cb types.BlockSeqCB
		err = types.Decode(value, &cb)
		return err == nil
	}
	return false
}

func (bs *BlockStore) seqCBNum() int64 {
	counts := dbm.NewListHelper(bs.db).PrefixCount(seqCBPrefix)
	return counts
}

func (bs *BlockStore) addBlockSeqCB(cb *types.BlockSeqCB) error {
	if len(cb.Name) > 128 || len(cb.URL) > 1024 {
		return types.ErrInvalidParam
	}
	storeLog.Info("addBlockSeqCB", "key", string(calcSeqCBKey([]byte(cb.Name))), "value", cb)

	return bs.db.SetSync(calcSeqCBKey([]byte(cb.Name)), types.Encode(cb))
}

func (bs *BlockStore) listSeqCB() (cbs []*types.BlockSeqCB, err error) {
	values := dbm.NewListHelper(bs.db).PrefixScan(seqCBPrefix)
	if values == nil {
		return nil, types.ErrNotFound
	}
	for _, value := range values {
		var cb types.BlockSeqCB
		err := types.Decode(value, &cb)
		if err != nil {
			return nil, err
		}
		cbs = append(cbs, &cb)
	}
	return cbs, nil
}

>>>>>>> upstream/master
func (bs *BlockStore) delAllKeys() {
	var allkeys [][]byte
	allkeys = append(allkeys, GetLocalDBKeyList()...)
	allkeys = append(allkeys, version.GetLocalDBKeyList()...)
	allkeys = append(allkeys, types.GetLocalDBKeyList()...)
	var lastkey []byte
	isvalid := true
	for isvalid {
		lastkey, isvalid = bs.delKeys(lastkey, allkeys)
	}
}

func (bs *BlockStore) delKeys(seek []byte, allkeys [][]byte) ([]byte, bool) {
	it := bs.db.Iterator(seek, types.EmptyValue, false)
	defer it.Close()
	i := 0
	count := 0
	var lastkey []byte
	for it.Rewind(); it.Valid(); it.Next() {
		key := it.Key()
		lastkey = key
		if it.Error() != nil {
			panic(it.Error())
		}
		has := false
		for _, prefix := range allkeys {
			if bytes.HasPrefix(key, prefix) {
				has = true
				break
			}
		}
		if !has {
			i++
			if i > 0 && i%10000 == 0 {
				chainlog.Info("del key count", "count", i)
			}
			err := bs.db.Delete(key)
			if err != nil {
				panic(err)
			}
		}
		count++
		if count == 1000000 {
			break
		}
	}
	return lastkey, it.Valid()
}

func (bs *BlockStore) saveQuickIndexFlag() {
	kv := types.FlagKV(types.FlagTxQuickIndex, 1)
	err := bs.db.Set(kv.Key, kv.Value)
	if err != nil {
		panic(err)
	}
}

func (bs *BlockStore) loadFlag(key []byte) (int64, error) {
	flag := &types.Int64{}
	flagBytes, err := bs.db.Get(key)
	if err == nil {
		err = types.Decode(flagBytes, flag)
		if err != nil {
			return 0, err
		}
		return flag.GetData(), nil
	} else if err == types.ErrNotFound || err == dbm.ErrNotFoundInDb {
		return 0, nil
	}
	return 0, err
}

<<<<<<< HEAD
=======
//HasTx 是否包含该交易
>>>>>>> upstream/master
func (bs *BlockStore) HasTx(key []byte) (bool, error) {
	if types.IsEnable("quickIndex") {
		if _, err := bs.db.Get(types.CalcTxShortKey(key)); err != nil {
			if err == dbm.ErrNotFoundInDb {
				return false, nil
			}
			return false, err
		}
		return true, nil
	}
	if _, err := bs.db.Get(types.CalcTxKey(key)); err != nil {
		if err == dbm.ErrNotFoundInDb {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

<<<<<<< HEAD
// 返回BlockStore保存的当前block高度
=======
//Height 返回BlockStore保存的当前block高度
>>>>>>> upstream/master
func (bs *BlockStore) Height() int64 {
	return atomic.LoadInt64(&bs.height)
}

<<<<<<< HEAD
// 更新db中的block高度到BlockStore.Height
=======
//UpdateHeight 更新db中的block高度到BlockStore.Height
>>>>>>> upstream/master
func (bs *BlockStore) UpdateHeight() {
	height, _ := LoadBlockStoreHeight(bs.db)
	atomic.StoreInt64(&bs.height, height)
	storeLog.Debug("UpdateHeight", "curblockheight", height)
}

<<<<<<< HEAD
=======
//UpdateHeight2 更新指定的block高度到BlockStore.Height
>>>>>>> upstream/master
func (bs *BlockStore) UpdateHeight2(height int64) {
	atomic.StoreInt64(&bs.height, height)
	storeLog.Debug("UpdateHeight2", "curblockheight", height)
}

<<<<<<< HEAD
// 返回BlockStore保存的当前blockheader
=======
//LastHeader 返回BlockStore保存的当前blockheader
>>>>>>> upstream/master
func (bs *BlockStore) LastHeader() *types.Header {
	lastheaderlock.Lock()
	defer lastheaderlock.Unlock()

	// 通过lastBlock获取lastheader
	var blockheader = types.Header{}
	if bs.lastBlock != nil {
		blockheader.Version = bs.lastBlock.Version
		blockheader.ParentHash = bs.lastBlock.ParentHash
		blockheader.TxHash = bs.lastBlock.TxHash
		blockheader.StateHash = bs.lastBlock.StateHash
		blockheader.Height = bs.lastBlock.Height
		blockheader.BlockTime = bs.lastBlock.BlockTime
		blockheader.Signature = bs.lastBlock.Signature
		blockheader.Difficulty = bs.lastBlock.Difficulty

		blockheader.Hash = bs.lastBlock.Hash()
		blockheader.TxCount = int64(len(bs.lastBlock.Txs))
	}
	return &blockheader
}

<<<<<<< HEAD
// 更新LastBlock到缓存中
=======
//UpdateLastBlock 更新LastBlock到缓存中
>>>>>>> upstream/master
func (bs *BlockStore) UpdateLastBlock(hash []byte) {
	blockdetail, err := bs.LoadBlockByHash(hash)
	if err != nil {
		storeLog.Error("UpdateLastBlock", "hash", common.ToHex(hash), "error", err)
		return
	}
	lastheaderlock.Lock()
	defer lastheaderlock.Unlock()
	if blockdetail != nil {
		bs.lastBlock = blockdetail.Block
	}
	storeLog.Debug("UpdateLastBlock", "UpdateLastBlock", blockdetail.Block.Height, "LastHederhash", common.ToHex(blockdetail.Block.Hash()))
}

<<<<<<< HEAD
=======
//UpdateLastBlock2 更新LastBlock到缓存中
>>>>>>> upstream/master
func (bs *BlockStore) UpdateLastBlock2(block *types.Block) {
	lastheaderlock.Lock()
	defer lastheaderlock.Unlock()
	bs.lastBlock = block
	storeLog.Debug("UpdateLastBlock", "UpdateLastBlock", block.Height, "LastHederhash", common.ToHex(block.Hash()))
}

<<<<<<< HEAD
//获取最新的block信息
=======
//LastBlock 获取最新的block信息
>>>>>>> upstream/master
func (bs *BlockStore) LastBlock() *types.Block {
	lastheaderlock.Lock()
	defer lastheaderlock.Unlock()
	if bs.lastBlock != nil {
		return bs.lastBlock
	}
	return nil
}

<<<<<<< HEAD
=======
//Get get
>>>>>>> upstream/master
func (bs *BlockStore) Get(keys *types.LocalDBGet) *types.LocalReplyValue {
	var reply types.LocalReplyValue
	for i := 0; i < len(keys.Keys); i++ {
		key := keys.Keys[i]
		value, _ := bs.db.Get(key)
		reply.Values = append(reply.Values, value)
	}
	return &reply
}

<<<<<<< HEAD
//通过height高度获取BlockDetail信息
=======
//LoadBlockByHeight 通过height高度获取BlockDetail信息
>>>>>>> upstream/master
func (bs *BlockStore) LoadBlockByHeight(height int64) (*types.BlockDetail, error) {
	//首先通过height获取block hash从db中
	hash, err := bs.GetBlockHashByHeight(height)
	if err != nil {
		return nil, err
	}
	return bs.LoadBlockByHash(hash)
}

<<<<<<< HEAD
//通过hash获取BlockDetail信息
=======
//LoadBlockByHash 通过hash获取BlockDetail信息
>>>>>>> upstream/master
func (bs *BlockStore) LoadBlockByHash(hash []byte) (*types.BlockDetail, error) {
	var blockdetail types.BlockDetail
	var blockheader types.Header
	var blockbody types.BlockBody
	var block types.Block

	//通过hash获取blockheader
	header, err := bs.db.Get(calcHashToBlockHeaderKey(hash))
	if header == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("LoadBlockByHash calcHashToBlockHeaderKey", "hash", common.ToHex(hash), "err", err)
		}
		return nil, types.ErrHashNotExist
	}
	err = proto.Unmarshal(header, &blockheader)
	if err != nil {
		storeLog.Error("LoadBlockByHash", "err", err)
		return nil, err
	}
	//通过hash获取blockbody
	body, err := bs.db.Get(calcHashToBlockBodyKey(hash))
	if body == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("LoadBlockByHash calcHashToBlockBodyKey ", "err", err)
		}
		return nil, types.ErrHashNotExist
	}
	err = proto.Unmarshal(body, &blockbody)
	if err != nil {
		storeLog.Error("LoadBlockByHash", "err", err)
		return nil, err
	}
	block.Version = blockheader.Version
	block.ParentHash = blockheader.ParentHash
	block.TxHash = blockheader.TxHash
	block.StateHash = blockheader.StateHash
	block.Height = blockheader.Height
	block.BlockTime = blockheader.BlockTime
	block.Signature = blockheader.Signature
	block.Difficulty = blockheader.Difficulty
	block.Txs = blockbody.Txs

	blockdetail.Receipts = blockbody.Receipts
	blockdetail.Block = &block

	//storeLog.Info("LoadBlockByHash", "Height", block.Height, "Difficulty", blockdetail.Block.Difficulty)

	return &blockdetail, nil
}

<<<<<<< HEAD
//  批量保存blocks信息到db数据库中
func (bs *BlockStore) SaveBlock(storeBatch dbm.Batch, blockdetail *types.BlockDetail, sequence int64) error {

=======
//SaveBlock 批量保存blocks信息到db数据库中,并返回最新的sequence值
func (bs *BlockStore) SaveBlock(storeBatch dbm.Batch, blockdetail *types.BlockDetail, sequence int64) (int64, error) {
	var lastSequence int64 = -1
>>>>>>> upstream/master
	height := blockdetail.Block.Height
	if len(blockdetail.Receipts) == 0 && len(blockdetail.Block.Txs) != 0 {
		storeLog.Error("SaveBlock Receipts is nil ", "height", height)
	}
	hash := blockdetail.Block.Hash()

	// Save blockbody通过block hash
	var blockbody types.BlockBody
	blockbody.Txs = blockdetail.Block.Txs
	blockbody.Receipts = blockdetail.Receipts

	body, err := proto.Marshal(&blockbody)
	if err != nil {
		storeLog.Error("SaveBlock Marshal blockbody", "height", height, "hash", common.ToHex(hash), "error", err)
<<<<<<< HEAD
		return err
=======
		return lastSequence, err
>>>>>>> upstream/master
	}
	storeBatch.Set(calcHashToBlockBodyKey(hash), body)

	// Save blockheader通过block hash
	var blockheader types.Header
	blockheader.Version = blockdetail.Block.Version
	blockheader.ParentHash = blockdetail.Block.ParentHash
	blockheader.TxHash = blockdetail.Block.TxHash
	blockheader.StateHash = blockdetail.Block.StateHash
	blockheader.Height = blockdetail.Block.Height
	blockheader.BlockTime = blockdetail.Block.BlockTime
	blockheader.Signature = blockdetail.Block.Signature
	blockheader.Difficulty = blockdetail.Block.Difficulty

	blockheader.Hash = hash
	blockheader.TxCount = int64(len(blockdetail.Block.Txs))

	header, err := proto.Marshal(&blockheader)
	if err != nil {
		storeLog.Error("SaveBlock Marshal blockheader", "height", height, "hash", common.ToHex(hash), "error", err)
<<<<<<< HEAD
		return err
=======
		return lastSequence, err
>>>>>>> upstream/master
	}

	storeBatch.Set(calcHashToBlockHeaderKey(hash), header)
	storeBatch.Set(calcHeightToBlockHeaderKey(height), header)

	//更新最新的block 高度
<<<<<<< HEAD
	heightbytes := types.Encode(&types.Int64{height})
=======
	heightbytes := types.Encode(&types.Int64{Data: height})
>>>>>>> upstream/master
	storeBatch.Set(blockLastHeight, heightbytes)

	//存储block hash和height的对应关系，便于通过hash查询block
	storeBatch.Set(calcHashToHeightKey(hash), heightbytes)

	//存储block height和block hash的对应关系，便于通过height查询block
	storeBatch.Set(calcHeightToHashKey(height), hash)

	if isRecordBlockSequence || isParaChain {
		//存储记录block序列执行的type add
<<<<<<< HEAD
		err = bs.SaveBlockSequence(storeBatch, hash, height, AddBlock, sequence)
		if err != nil {
			storeLog.Error("SaveBlock SaveBlockSequence", "height", height, "hash", common.ToHex(hash), "error", err)
			return err
		}
	}
	storeLog.Debug("SaveBlock success", "blockheight", height, "hash", common.ToHex(hash))
	return nil
}

// 删除block信息从db数据库中
func (bs *BlockStore) DelBlock(storeBatch dbm.Batch, blockdetail *types.BlockDetail, sequence int64) error {

=======
		lastSequence, err = bs.saveBlockSequence(storeBatch, hash, height, AddBlock, sequence)
		if err != nil {
			storeLog.Error("SaveBlock SaveBlockSequence", "height", height, "hash", common.ToHex(hash), "error", err)
			return lastSequence, err
		}
	}
	storeLog.Debug("SaveBlock success", "blockheight", height, "hash", common.ToHex(hash))
	return lastSequence, nil
}

//DelBlock 删除block信息从db数据库中
func (bs *BlockStore) DelBlock(storeBatch dbm.Batch, blockdetail *types.BlockDetail, sequence int64) (int64, error) {
	var lastSequence int64 = -1
>>>>>>> upstream/master
	height := blockdetail.Block.Height
	hash := blockdetail.Block.Hash()

	//更新最新的block高度为前一个高度
<<<<<<< HEAD
	bytes := types.Encode(&types.Int64{height - 1})
=======
	bytes := types.Encode(&types.Int64{Data: height - 1})
>>>>>>> upstream/master
	storeBatch.Set(blockLastHeight, bytes)

	//删除block hash和height的对应关系
	storeBatch.Delete(calcHashToHeightKey(hash))

	//删除block height和block hash的对应关系，便于通过height查询block
	storeBatch.Delete(calcHeightToHashKey(height))
	storeBatch.Delete(calcHeightToBlockHeaderKey(height))

	if isRecordBlockSequence || isParaChain {
		//存储记录block序列执行的type del
<<<<<<< HEAD
		err := bs.SaveBlockSequence(storeBatch, hash, height, DelBlock, sequence)
		if err != nil {
			storeLog.Error("DelBlock SaveBlockSequence", "height", height, "hash", common.ToHex(hash), "error", err)
			return err
=======
		lastSequence, err := bs.saveBlockSequence(storeBatch, hash, height, DelBlock, sequence)
		if err != nil {
			storeLog.Error("DelBlock SaveBlockSequence", "height", height, "hash", common.ToHex(hash), "error", err)
			return lastSequence, err
>>>>>>> upstream/master
		}
	}

	storeLog.Debug("DelBlock success", "blockheight", height, "hash", common.ToHex(hash))
<<<<<<< HEAD
	return nil
}

// 通过tx hash 从db数据库中获取tx交易信息
=======
	return lastSequence, nil
}

//GetTx 通过tx hash 从db数据库中获取tx交易信息
>>>>>>> upstream/master
func (bs *BlockStore) GetTx(hash []byte) (*types.TxResult, error) {
	if len(hash) == 0 {
		err := errors.New("input hash is null")
		return nil, err
	}
	rawBytes, err := bs.db.Get(types.CalcTxKey(hash))
	if rawBytes == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("GetTx", "hash", common.ToHex(hash), "err", err)
		}
		err = errors.New("tx not exist")
		return nil, err
	}

	var txResult types.TxResult
	err = proto.Unmarshal(rawBytes, &txResult)
	if err != nil {
		return nil, err
	}
	return &txResult, nil
}

<<<<<<< HEAD
// 通过批量存储tx信息到db中
=======
//AddTxs 通过批量存储tx信息到db中
>>>>>>> upstream/master
func (bs *BlockStore) AddTxs(storeBatch dbm.Batch, blockDetail *types.BlockDetail) error {
	kv, err := bs.getLocalKV(blockDetail)
	if err != nil {
		storeLog.Error("indexTxs getLocalKV err", "Height", blockDetail.Block.Height, "err", err)
		return err
	}
	//storelog.Info("add txs kv num", "n", len(kv.KV))
	for i := 0; i < len(kv.KV); i++ {
		if kv.KV[i].Value == nil {
			storeBatch.Delete(kv.KV[i].Key)
		} else {
			storeBatch.Set(kv.KV[i].Key, kv.KV[i].Value)
		}
	}
	return nil
}

<<<<<<< HEAD
//通过批量删除tx信息从db中
=======
//DelTxs 通过批量删除tx信息从db中
>>>>>>> upstream/master
func (bs *BlockStore) DelTxs(storeBatch dbm.Batch, blockDetail *types.BlockDetail) error {
	//存储key:addr:flag:height ,value:txhash
	//flag :0-->from,1--> to
	//height=height*10000+index 存储账户地址相关的交易
	kv, err := bs.getDelLocalKV(blockDetail)
	if err != nil {
		storeLog.Error("indexTxs getLocalKV err", "Height", blockDetail.Block.Height, "err", err)
		return err
	}
	for i := 0; i < len(kv.KV); i++ {
		if kv.KV[i].Value == nil {
			storeBatch.Delete(kv.KV[i].Key)
		} else {
			storeBatch.Set(kv.KV[i].Key, kv.KV[i].Value)
		}
	}

	return nil
}

<<<<<<< HEAD
//从db数据库中获取指定hash对应的block高度
=======
//GetHeightByBlockHash 从db数据库中获取指定hash对应的block高度
>>>>>>> upstream/master
func (bs *BlockStore) GetHeightByBlockHash(hash []byte) (int64, error) {

	heightbytes, err := bs.db.Get(calcHashToHeightKey(hash))
	if heightbytes == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("GetHeightByBlockHash", "error", err)
		}
		return -1, types.ErrHashNotExist
	}
	return decodeHeight(heightbytes)
}

func decodeHeight(heightbytes []byte) (int64, error) {
	var height types.Int64
	err := types.Decode(heightbytes, &height)
	if err != nil {
		//may be old database format json...
		err = json.Unmarshal(heightbytes, &height.Data)
		if err != nil {
			storeLog.Error("GetHeightByBlockHash Could not unmarshal height bytes", "error", err)
			return -1, types.ErrUnmarshal
		}
	}
	return height.Data, nil
}

<<<<<<< HEAD
//从db数据库中获取指定height对应的blockhash
=======
//GetBlockHashByHeight 从db数据库中获取指定height对应的blockhash
>>>>>>> upstream/master
func (bs *BlockStore) GetBlockHashByHeight(height int64) ([]byte, error) {

	hash, err := bs.db.Get(calcHeightToHashKey(height))
	if hash == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("GetBlockHashByHeight", "error", err)
		}
		return nil, types.ErrHeightNotExist
	}
	return hash, nil
}

<<<<<<< HEAD
//通过blockheight获取blockheader
=======
//GetBlockHeaderByHeight 通过blockheight获取blockheader
>>>>>>> upstream/master
func (bs *BlockStore) GetBlockHeaderByHeight(height int64) (*types.Header, error) {
	//从最新版本的key里面获取header，找不到找老版本的数据库
	blockheader, err := bs.db.Get(calcHeightToBlockHeaderKey(height))
	var flagFoundInOldDB bool
	if err != nil {
		//首先通过height获取block hash从db中
		var hash []byte
		hash, err = bs.GetBlockHashByHeight(height)
		if err != nil {
			return nil, err
		}
		blockheader, err = bs.db.Get(calcHashToBlockHeaderKey(hash))
		flagFoundInOldDB = true
	}
	if blockheader == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("GetBlockHeaderByHeight calcHashToBlockHeaderKey", "error", err)
		}
		return nil, types.ErrHashNotExist
	}
	var header types.Header
	err = proto.Unmarshal(blockheader, &header)
	if err != nil {
		storeLog.Error("GetBlockHerderByHeight", "Could not unmarshal blockheader:", blockheader)
		return nil, err
	}
	if flagFoundInOldDB {
		bs.db.Set(calcHeightToBlockHeaderKey(height), blockheader)
	}
	return &header, nil
}

<<<<<<< HEAD
//通过blockhash获取blockheader
=======
//GetBlockHeaderByHash 通过blockhash获取blockheader
>>>>>>> upstream/master
func (bs *BlockStore) GetBlockHeaderByHash(hash []byte) (*types.Header, error) {

	var header types.Header
	blockheader, err := bs.db.Get(calcHashToBlockHeaderKey(hash))
	if blockheader == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("GetBlockHerderByHash calcHashToBlockHeaderKey ", "err", err)
		}
		return nil, types.ErrHashNotExist
	}
	err = proto.Unmarshal(blockheader, &header)
	if err != nil {
		storeLog.Error("GetBlockHerderByHash", "err", err)
		return nil, err
	}
	return &header, nil
}

func (bs *BlockStore) getLocalKV(detail *types.BlockDetail) (*types.LocalDBSet, error) {
	if bs.client == nil {
		panic("client not bind message queue.")
	}
	msg := bs.client.NewMessage("execs", types.EventAddBlock, detail)
	bs.client.Send(msg, true)
	resp, err := bs.client.Wait(msg)
	if err != nil {
		return nil, err
	}
	kv := resp.GetData().(*types.LocalDBSet)
	return kv, nil
}

func (bs *BlockStore) getDelLocalKV(detail *types.BlockDetail) (*types.LocalDBSet, error) {
	if bs.client == nil {
		panic("client not bind message queue.")
	}
	msg := bs.client.NewMessage("execs", types.EventDelBlock, detail)
	bs.client.Send(msg, true)
	resp, err := bs.client.Wait(msg)
	if err != nil {
		return nil, err
	}
	localDBSet := resp.GetData().(*types.LocalDBSet)
	return localDBSet, nil
}

<<<<<<< HEAD
//从db数据库中获取指定blockhash对应的block总难度td
=======
//GetTdByBlockHash 从db数据库中获取指定blockhash对应的block总难度td
>>>>>>> upstream/master
func (bs *BlockStore) GetTdByBlockHash(hash []byte) (*big.Int, error) {

	blocktd, err := bs.db.Get(calcHashToTdKey(hash))
	if blocktd == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("GetTdByBlockHash ", "error", err)
		}
		return nil, types.ErrHashNotExist
	}
	td := new(big.Int)
	return td.SetBytes(blocktd), nil
}

<<<<<<< HEAD
//保存block hash对应的总难度到db中
=======
//SaveTdByBlockHash 保存block hash对应的总难度到db中
>>>>>>> upstream/master
func (bs *BlockStore) SaveTdByBlockHash(storeBatch dbm.Batch, hash []byte, td *big.Int) error {
	if td == nil {
		return types.ErrInvalidParam
	}

	storeBatch.Set(calcHashToTdKey(hash), td.Bytes())
	return nil
}

<<<<<<< HEAD
=======
//NewBatch new
>>>>>>> upstream/master
func (bs *BlockStore) NewBatch(sync bool) dbm.Batch {
	storeBatch := bs.db.NewBatch(sync)
	return storeBatch
}

<<<<<<< HEAD
=======
//LoadBlockStoreHeight 加载区块高度
>>>>>>> upstream/master
func LoadBlockStoreHeight(db dbm.DB) (int64, error) {
	bytes, err := db.Get(blockLastHeight)
	if bytes == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("LoadBlockStoreHeight", "error", err)
		}
		return -1, types.ErrHeightNotExist
	}
	return decodeHeight(bytes)
}

// 将收到的block都暂时存储到db中，加入主链之后会重新覆盖。主要是用于chain重组时获取侧链的block使用
func (bs *BlockStore) dbMaybeStoreBlock(blockdetail *types.BlockDetail, sync bool) error {
	if blockdetail == nil {
		return types.ErrInvalidParam
	}
	height := blockdetail.Block.GetHeight()
	hash := blockdetail.Block.Hash()
	storeBatch := bs.NewBatch(sync)

	// Save blockbody通过block hash
	var blockbody types.BlockBody
	blockbody.Txs = blockdetail.Block.Txs
	blockbody.Receipts = blockdetail.Receipts

	body, err := proto.Marshal(&blockbody)
	if err != nil {
		storeLog.Error("dbMaybeStoreBlock Marshal blockbody", "height", height, "hash", common.ToHex(hash), "error", err)
		return types.ErrMarshal
	}
	storeBatch.Set(calcHashToBlockBodyKey(hash), body)

	// Save blockheader通过block hash
	var blockheader types.Header
	blockheader.Version = blockdetail.Block.Version
	blockheader.ParentHash = blockdetail.Block.ParentHash
	blockheader.TxHash = blockdetail.Block.TxHash
	blockheader.StateHash = blockdetail.Block.StateHash
	blockheader.Height = blockdetail.Block.Height
	blockheader.BlockTime = blockdetail.Block.BlockTime
	blockheader.Signature = blockdetail.Block.Signature
	blockheader.Difficulty = blockdetail.Block.Difficulty
	blockheader.Hash = hash
	blockheader.TxCount = int64(len(blockdetail.Block.Txs))

	header, err := proto.Marshal(&blockheader)
	if err != nil {
		storeLog.Error("dbMaybeStoreBlock Marshal blockheader", "height", height, "hash", common.ToHex(hash), "error", err)
		return types.ErrMarshal
	}
	storeBatch.Set(calcHashToBlockHeaderKey(hash), header)

	//保存block的总难度到db中
	parentHash := blockdetail.Block.ParentHash

	//转换自己的难度成big.int
	difficulty := difficulty.CalcWork(blockdetail.Block.Difficulty)
	//chainlog.Error("dbMaybeStoreBlock Difficulty", "height", height, "Block.Difficulty", blockdetail.Block.Difficulty)
	//chainlog.Error("dbMaybeStoreBlock Difficulty bigint", "height", height, "self.Difficulty", difficulty.BigToCompact(difficulty))

	var blocktd *big.Int
	if height == 0 {
		blocktd = difficulty
	} else {
		parenttd, err := bs.GetTdByBlockHash(parentHash)
		if err != nil {
			chainlog.Error("dbMaybeStoreBlock GetTdByBlockHash", "height", height, "parentHash", common.ToHex(parentHash))
			return err
		}
		blocktd = new(big.Int).Add(difficulty, parenttd)
		//chainlog.Error("dbMaybeStoreBlock Difficulty", "height", height, "parenttd.td", difficulty.BigToCompact(parenttd))
		//chainlog.Error("dbMaybeStoreBlock Difficulty", "height", height, "self.td", difficulty.BigToCompact(blocktd))
	}

	err = bs.SaveTdByBlockHash(storeBatch, blockdetail.Block.Hash(), blocktd)
	if err != nil {
		chainlog.Error("dbMaybeStoreBlock SaveTdByBlockHash:", "height", height, "hash", common.ToHex(hash), "err", err)
		return err
	}

	err = storeBatch.Write()
	if err != nil {
		chainlog.Error("dbMaybeStoreBlock storeBatch.Write:", "err", err)
		return types.ErrDataBaseDamage
	}
	return nil
}

<<<<<<< HEAD
//获取当前最新的block操作序列号
=======
//LoadBlockLastSequence 获取当前最新的block操作序列号
>>>>>>> upstream/master
func (bs *BlockStore) LoadBlockLastSequence() (int64, error) {
	bytes, err := bs.db.Get(LastSequence)
	if bytes == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("LoadBlockLastSequence", "error", err)
		}
		return -1, types.ErrHeightNotExist
	}
	return decodeHeight(bytes)
}

<<<<<<< HEAD
//存储block 序列执行的类型用于blockchain的恢复
//获取当前的序列号，将此序列号加1存储本block的hash ，当主链使能isRecordBlockSequence
// 平行链使能isParaChain时，sequence序列号是传入的
func (bs *BlockStore) SaveBlockSequence(storeBatch dbm.Batch, hash []byte, height int64, Type int64, sequence int64) error {
=======
func (bs *BlockStore) setSeqCBLastNum(name []byte, num int64) error {
	return bs.db.SetSync(caclSeqCBLastNumKey(name), types.Encode(&types.Int64{Data: num}))
}

//Seq的合法值从0开始的，所以没有获取到或者获取失败都应该返回-1
func (bs *BlockStore) getSeqCBLastNum(name []byte) int64 {
	bytes, err := bs.db.Get(caclSeqCBLastNumKey(name))
	if bytes == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("getSeqCBLastNum", "error", err)
		}
		return -1
	}
	n, err := decodeHeight(bytes)
	if err != nil {
		return -1
	}
	storeLog.Error("getSeqCBLastNum", "name", string(name), "num", n)

	return n
}

//SaveBlockSequence 存储block 序列执行的类型用于blockchain的恢复
//获取当前的序列号，将此序列号加1存储本block的hash ，当主链使能isRecordBlockSequence
// 平行链使能isParaChain时，sequence序列号是传入的
func (bs *BlockStore) saveBlockSequence(storeBatch dbm.Batch, hash []byte, height int64, Type int64, sequence int64) (int64, error) {
>>>>>>> upstream/master

	var blockSequence types.BlockSequence
	var newSequence int64

	if isRecordBlockSequence {
		Sequence, err := bs.LoadBlockLastSequence()
		if err != nil {
			storeLog.Error("SaveBlockSequence", "LoadBlockLastSequence err", err)
			if err != types.ErrHeightNotExist {
				panic(err)
			}
		}

		newSequence = Sequence + 1
		//开启isRecordBlockSequence功能必须从0开始同步数据，不允许从非0高度开启此功能
		if newSequence == 0 && height != 0 {
			storeLog.Error("isRecordBlockSequence is true must Synchronizing data from zero block", "height", height)
			panic(errors.New("isRecordBlockSequence is true must Synchronizing data from zero block"))
		}
	} else if isParaChain {
		newSequence = sequence
	}
	blockSequence.Hash = hash
	blockSequence.Type = Type

	BlockSequenceByte, err := proto.Marshal(&blockSequence)
	if err != nil {
		storeLog.Error("SaveBlockSequence Marshal BlockSequence", "hash", common.ToHex(hash), "error", err)
<<<<<<< HEAD
		return err
=======
		return newSequence, err
>>>>>>> upstream/master
	}

	// seq->hash
	storeBatch.Set(calcSequenceToHashKey(newSequence), BlockSequenceByte)

	//parachain  hash->seq 只记录add block时的hash和seq对应关系
<<<<<<< HEAD
	if Type == AddBlock && isParaChain {
		Sequencebytes := types.Encode(&types.Int64{newSequence})
		storeBatch.Set(calcHashToSequenceKey(hash), Sequencebytes)
	}
	Sequencebytes := types.Encode(&types.Int64{newSequence})
	storeBatch.Set(LastSequence, Sequencebytes)

	return nil
}

//通过seq高度获取BlockDetail信息
=======
	if Type == AddBlock {
		Sequencebytes := types.Encode(&types.Int64{Data: newSequence})
		storeBatch.Set(calcHashToSequenceKey(hash), Sequencebytes)
	}
	Sequencebytes := types.Encode(&types.Int64{Data: newSequence})
	storeBatch.Set(LastSequence, Sequencebytes)

	return newSequence, nil
}

//LoadBlockBySequence 通过seq高度获取BlockDetail信息
>>>>>>> upstream/master
func (bs *BlockStore) LoadBlockBySequence(Sequence int64) (*types.BlockDetail, error) {
	//首先通过Sequence序列号获取对应的blockhash和操作类型从db中
	BlockSequence, err := bs.GetBlockSequence(Sequence)
	if err != nil {
		return nil, err
	}
	return bs.LoadBlockByHash(BlockSequence.Hash)
}

<<<<<<< HEAD
//从db数据库中获取指定Sequence对应的block序列操作信息
=======
//GetBlockSequence 从db数据库中获取指定Sequence对应的block序列操作信息
>>>>>>> upstream/master
func (bs *BlockStore) GetBlockSequence(Sequence int64) (*types.BlockSequence, error) {

	var blockSeq types.BlockSequence
	blockSeqByte, err := bs.db.Get(calcSequenceToHashKey(Sequence))
	if blockSeqByte == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("GetBlockSequence", "error", err)
		}
		return nil, types.ErrHeightNotExist
	}

	err = proto.Unmarshal(blockSeqByte, &blockSeq)
	if err != nil {
		storeLog.Error("GetBlockSequence", "err", err)
		return nil, err
	}
	return &blockSeq, nil
}

<<<<<<< HEAD
//通过block还是获取对应的seq，只提供给parachain使用
=======
//GetSequenceByHash 通过block还是获取对应的seq，只提供给parachain使用
>>>>>>> upstream/master
func (bs *BlockStore) GetSequenceByHash(hash []byte) (int64, error) {
	var seq types.Int64
	seqbytes, err := bs.db.Get(calcHashToSequenceKey(hash))
	if seqbytes == nil || err != nil {
		if err != dbm.ErrNotFoundInDb {
			storeLog.Error("GetSequenceByHash", "error", err)
		}
<<<<<<< HEAD
		return -1, types.ErrHeightNotExist
=======
		return -1, types.ErrHashNotExist
>>>>>>> upstream/master
	}

	err = types.Decode(seqbytes, &seq)
	if err != nil {
		storeLog.Error("GetSequenceByHash  types.Decode", "error", err)
		return -1, types.ErrUnmarshal
	}
	return seq.Data, nil
}

<<<<<<< HEAD
//获取blockchain的数据库版本号
=======
//GetDbVersion 获取blockchain的数据库版本号
>>>>>>> upstream/master
func (bs *BlockStore) GetDbVersion() int64 {
	ver := types.Int64{}
	version, err := bs.db.Get(version.BlockChainVerKey)
	if err != nil && err != types.ErrNotFound {
		storeLog.Info("GetDbVersion", "err", err)
		return 0
	}
	if len(version) == 0 {
		storeLog.Info("GetDbVersion len(version)==0")
		return 0
	}
	err = types.Decode(version, &ver)
	if err != nil {
		storeLog.Info("GetDbVersion", "types.Decode err", err)
		return 0
	}
	storeLog.Info("GetDbVersion", "blockchain db version", ver.Data)
	return ver.Data
}

<<<<<<< HEAD
//获取blockchain的数据库版本号
=======
//SetDbVersion 获取blockchain的数据库版本号
>>>>>>> upstream/master
func (bs *BlockStore) SetDbVersion(versionNo int64) error {
	ver := types.Int64{Data: versionNo}
	verByte := types.Encode(&ver)

	storeLog.Info("SetDbVersion", "blcokchain db version", versionNo)

	return bs.db.SetSync(version.BlockChainVerKey, verByte)
}

<<<<<<< HEAD
=======
//GetUpgradeMeta 获取blockchain的数据库版本号
>>>>>>> upstream/master
func (bs *BlockStore) GetUpgradeMeta() (*types.UpgradeMeta, error) {
	ver := types.UpgradeMeta{}
	version, err := bs.db.Get(version.LocalDBMeta)
	if err != nil && err != dbm.ErrNotFoundInDb {
		return nil, err
	}
	if len(version) == 0 {
		return &types.UpgradeMeta{Version: "0.0.0"}, nil
	}
	err = types.Decode(version, &ver)
	if err != nil {
		return nil, err
	}
	storeLog.Info("GetUpgradeMeta", "blockchain db version", ver)
	return &ver, nil
}

<<<<<<< HEAD
//获取blockchain的数据库版本号
=======
//SetUpgradeMeta 设置blockchain的数据库版本号
>>>>>>> upstream/master
func (bs *BlockStore) SetUpgradeMeta(meta *types.UpgradeMeta) error {
	verByte := types.Encode(meta)
	storeLog.Info("SetUpgradeMeta", "meta", meta)
	return bs.db.SetSync(version.LocalDBMeta, verByte)
}
