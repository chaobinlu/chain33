// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mavl

import (
	"bytes"
	"fmt"
	"sync"

<<<<<<< HEAD
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/golang-lru"
	"github.com/33cn/chain33/common"
	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/types"
)

const (
	leafKeyCountPrefix = "..mk.."
	delMapPoolPrefix   = "_..md.._"
	blockHeightStrLen  = 10
	pruningStateStart  = 1
	pruningStateEnd    = 0
=======
	"sort"
	"sync/atomic"
	"time"

	"github.com/33cn/chain33/common"
	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/types"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/golang-lru"
)

const (
	leafKeyCountPrefix     = "..mk.."
	oldLeafKeyCountPrefix  = "..mok.."
	secLvlPruningHeightKey = "_..mslphk.._"
	delMapPoolPrefix       = "_..md.._"
	blockHeightStrLen      = 10
	pruningStateStart      = 1
	pruningStateEnd        = 0
>>>>>>> upstream/master
	//删除节点pool以hash的首字母为key因此有256个
	delNodeCacheSize = 256 + 1
	//每个del Pool下存放默认4096个hash
	perDelNodePoolSize = 4096
<<<<<<< HEAD
=======
	//二级裁剪高度，达到此高度未裁剪则放入该处
	secondLevelPruningHeight = 1000000
	//三级裁剪高度，达到此高度还没有裁剪，则不进行裁剪
	threeLevelPruningHeight = 1500000
	onceScanCount           = 100000
>>>>>>> upstream/master
)

var (
	// 是否开启mavl裁剪
	enablePrune bool
	// 每个10000裁剪一次
<<<<<<< HEAD
	pruneHeight int = 10000
	// 裁剪状态
	pruningState int32
	delPoolCache *lru.Cache
	wg           sync.WaitGroup
	quit         bool
=======
	pruneHeight = 10000
	// 裁剪状态
	pruningState   int32
	delPoolCache   *lru.Cache
	wg             sync.WaitGroup
	quit           bool
	secLvlPruningH int64
>>>>>>> upstream/master
)

func init() {
	cache, err := lru.New(delNodeCacheSize)
	if err != nil {
		panic(fmt.Sprint("new delNodeCache lru fail", err))
	}
	delPoolCache = cache
}

type delNodeValuePool struct {
	delCache *lru.Cache
}

type hashData struct {
	height int64
	hash   []byte
}

<<<<<<< HEAD
func NewDelNodeValuePool(cacSize int) *delNodeValuePool {
=======
// newDelNodeValuePool 创建记录已经删除节点缓存池
func newDelNodeValuePool(cacSize int) *delNodeValuePool {
>>>>>>> upstream/master
	cache, err := lru.New(cacSize)
	if err != nil {
		return nil
	}
	dNodePool := &delNodeValuePool{}
	dNodePool.delCache = cache
	return dNodePool
}

<<<<<<< HEAD
=======
// EnablePrune 使能裁剪
>>>>>>> upstream/master
func EnablePrune(enable bool) {
	enablePrune = enable
	//开启裁剪需要同时开启前缀
	enableMavlPrefix = enable
}

<<<<<<< HEAD
=======
// SetPruneHeight 设置每次裁剪高度
>>>>>>> upstream/master
func SetPruneHeight(height int) {
	pruneHeight = height
}

<<<<<<< HEAD
=======
// ClosePrune 关闭裁剪
>>>>>>> upstream/master
func ClosePrune() {
	quit = true
	wg.Wait()
	//防止BaseStore没有关闭再次进入
	setPruning(pruningStateStart)
}

func genLeafCountKey(key, hash []byte, height int64) (hashkey []byte) {
	hashkey = []byte(fmt.Sprintf("%s%s%010d%s", leafKeyCountPrefix, string(key), height, string(hash)))
	return hashkey
}

func getKeyFromLeafCountKey(hashkey []byte, hashlen int) ([]byte, error) {
	if len(hashkey) <= len(leafKeyCountPrefix)+hashlen+blockHeightStrLen {
		return nil, types.ErrSize
	}
	if !bytes.Contains(hashkey, []byte(leafKeyCountPrefix)) {
		return nil, types.ErrSize
	}
	k := bytes.TrimPrefix(hashkey, []byte(leafKeyCountPrefix))
	k = k[:len(k)-hashlen-blockHeightStrLen]
	return k, nil
}

<<<<<<< HEAD
=======
func genOldLeafCountKey(key, hash []byte, height int64) (hashkey []byte) {
	hashkey = []byte(fmt.Sprintf("%s%s%010d%s", oldLeafKeyCountPrefix, string(key), height, string(hash)))
	return hashkey
}

func getKeyFromOldLeafCountKey(hashkey []byte, hashlen int) ([]byte, error) {
	if len(hashkey) <= len(oldLeafKeyCountPrefix)+hashlen+blockHeightStrLen {
		return nil, types.ErrSize
	}
	if !bytes.Contains(hashkey, []byte(oldLeafKeyCountPrefix)) {
		return nil, types.ErrSize
	}
	k := bytes.TrimPrefix(hashkey, []byte(oldLeafKeyCountPrefix))
	k = k[:len(k)-hashlen-blockHeightStrLen]
	return k, nil
}

func genOldLeafCountKeyFromKey(hashk []byte) (oldhashk []byte) {
	if len(hashk) < len(leafKeyCountPrefix) {
		return hashk
	}
	oldhashk = []byte(fmt.Sprintf("%s%s", oldLeafKeyCountPrefix, string(hashk[len(leafKeyCountPrefix):])))
	return oldhashk
}

>>>>>>> upstream/master
func genDeletePoolKey(hash []byte) (key, value []byte) {
	if len(hash) < 32 {
		panic("genDeletePoolKey error hash len illegal")
	}
	hashLen := len(common.Hash{})
	if len(hash) > hashLen {
		value = hash[len(hash)-hashLen:]
	} else {
		value = hash
	}
	key = value[:1]
	key = []byte(fmt.Sprintf("%s%s", delMapPoolPrefix, string(key)))
	return key, value
}

func isPruning() bool {
	return atomic.LoadInt32(&pruningState) == 1
}

func setPruning(state int32) {
	atomic.StoreInt32(&pruningState, state)
}

<<<<<<< HEAD
=======
func getSecLvlPruningHeight(db dbm.DB) int64 {
	value, err := db.Get([]byte(secLvlPruningHeightKey))
	if len(value) == 0 || err != nil {
		return 0
	}
	h := &types.Int64{}
	err = proto.Unmarshal(value, h)
	if err != nil {
		return 0
	}
	return h.Data
}

func setSecLvlPruningHeight(db dbm.DB, height int64) error {
	h := &types.Int64{}
	h.Data = height
	value, err := proto.Marshal(h)
	if err != nil {
		return err
	}
	return db.Set([]byte(secLvlPruningHeightKey), value)
}

>>>>>>> upstream/master
func pruningTree(db dbm.DB, curHeight int64) {
	wg.Add(1)
	defer wg.Add(-1)
	setPruning(pruningStateStart)
<<<<<<< HEAD
	treelog.Info("pruningTree", "start curHeight:", curHeight)
	start := time.Now()
	pruningTreeLeafNode(db, curHeight)
	end := time.Now()
	treelog.Info("pruningTree", "curHeight:", curHeight, "pruning leafNode cost time:", end.Sub(start))
	setPruning(pruningStateEnd)
}

func pruningTreeLeafNode(db dbm.DB, curHeight int64) {
=======
	// 一级遍历
	pruningFirstLevel(db, curHeight)
	// 二级遍历
	pruningSecondLevel(db, curHeight)
	setPruning(pruningStateEnd)
}

func pruningFirstLevel(db dbm.DB, curHeight int64) {
	treelog.Info("pruningTree pruningFirstLevel", "start curHeight:", curHeight)
	start := time.Now()
	pruningFirstLevelNode(db, curHeight)
	end := time.Now()
	treelog.Info("pruningTree pruningFirstLevel", "curHeight:", curHeight, "pruning leafNode cost time:", end.Sub(start))
}

func pruningFirstLevelNode(db dbm.DB, curHeight int64) {
>>>>>>> upstream/master
	prefix := []byte(leafKeyCountPrefix)
	it := db.Iterator(prefix, nil, true)
	defer it.Close()

<<<<<<< HEAD
	const onceScanCount = 100000
	mp := make(map[string][]hashData)
=======
	mp := make(map[string][]hashData)
	var kvs []*types.KeyValue
>>>>>>> upstream/master
	count := 0
	for it.Rewind(); it.Valid(); it.Next() {
		if quit {
			//该处退出
			return
		}
		//copy key
		hashK := make([]byte, len(it.Key()))
		copy(hashK, it.Key())

		value := it.Value()
		var pData types.PruneData
		err := proto.Unmarshal(value, &pData)
		if err != nil {
			panic("Unmarshal mavl leafCountKey fail")
		}
<<<<<<< HEAD
		hashLen := int(pData.Lenth)
		key, err := getKeyFromLeafCountKey(hashK, hashLen)
		if err == nil {
			data := hashData{
				height: pData.Height,
				hash:   hashK[len(hashK)-hashLen:],
			}
			mp[string(key)] = append(mp[string(key)], data)
			count++
			if count >= onceScanCount {
				deleteNode(db, mp, curHeight, key)
				count = 0
			}
=======
		var key []byte
		if curHeight < pData.Height+secondLevelPruningHeight {
			hashLen := int(pData.Lenth)
			key, err = getKeyFromLeafCountKey(hashK, hashLen)
			if err == nil {
				data := hashData{
					height: pData.Height,
					hash:   hashK[len(hashK)-hashLen:],
				}
				mp[string(key)] = append(mp[string(key)], data)
				count++
			}
		} else {
			value := make([]byte, len(it.Value()))
			copy(value, it.Value())
			kvs = append(kvs, &types.KeyValue{Key: hashK, Value: value})
		}
		if count >= onceScanCount {
			deleteNode(db, mp, curHeight, key)
			count = 0
		}
		if len(kvs) >= onceScanCount/2 {
			addLeafCountKeyToSecondLevel(db, kvs)
			kvs = kvs[:0]
>>>>>>> upstream/master
		}
	}
	if count > 0 {
		deleteNode(db, mp, curHeight, nil)
	}
<<<<<<< HEAD
=======
	if len(kvs) > 0 {
		addLeafCountKeyToSecondLevel(db, kvs)
	}
}

func addLeafCountKeyToSecondLevel(db dbm.DB, kvs []*types.KeyValue) {
	batch := db.NewBatch(true)
	for _, kv := range kvs {
		batch.Delete(kv.Key)
		batch.Set(genOldLeafCountKeyFromKey(kv.Key), kv.Value)
	}
	batch.Write()
>>>>>>> upstream/master
}

func deleteNode(db dbm.DB, mp map[string][]hashData, curHeight int64, lastKey []byte) {
	if len(mp) == 0 {
		return
	}
	var tmp []hashData
	//del
	if lastKey != nil {
		if _, ok := mp[string(lastKey)]; ok {
			tmp = mp[string(lastKey)]
			delete(mp, string(lastKey))
		}
	}
	delMp := make(map[string]bool)
	batch := db.NewBatch(true)
	for key, vals := range mp {
		if len(vals) > 1 {
			if vals[1].height != vals[0].height { //防止相同高度时候出现的误删除
				for _, val := range vals[1:] { //从第二个开始判断
					if curHeight >= val.height+int64(pruneHeight) {
						//batch.Delete(val.hash) //叶子节点hash值的删除放入pruningHashNode中
						batch.Delete(genLeafCountKey([]byte(key), val.hash, val.height))
						delMp[string(val.hash)] = true
					}
				}
			}
		}
		delete(mp, key)
	}
	batch.Write()
	//add
	if lastKey != nil {
		if _, ok := mp[string(lastKey)]; ok {
			mp[string(lastKey)] = tmp
		}
	}
	//裁剪hashNode
	pruningHashNode(db, delMp)
}

func pruningHashNode(db dbm.DB, mp map[string]bool) {
	if len(mp) == 0 {
		return
	}
<<<<<<< HEAD
	ndb := newMarkNodeDB(db, 1024*10)
	var delNodeStrs []string
	for key := range mp {
=======
	//对mp排序
	sortKeys := make([]string, len(mp)+1)
	for key := range mp {
		sortKeys = append(sortKeys, key)
	}
	sort.Strings(sortKeys)
	ndb := newMarkNodeDB(db, 1024*10)
	var delNodeStrs []string
	for _, key := range sortKeys {
>>>>>>> upstream/master
		mNode, err := ndb.LoadLeaf([]byte(key))
		if err == nil {
			delNodeStrs = append(delNodeStrs, mNode.getHashNode(ndb)...)
		}
	}
	//根据keyMap进行归类
	mpAddDel := make(map[string][][]byte)
	for _, s := range delNodeStrs {
		key, hash := genDeletePoolKey([]byte(s))
		mpAddDel[string(key)] = append(mpAddDel[string(key)], hash)
	}

	//更新pool数据
	batch := db.NewBatch(true)
	for mpk, mpV := range mpAddDel {
		dep := ndb.getPool(mpk)
		if dep != nil {
			for _, aHsh := range mpV {
				dep.delCache.Add(string(aHsh), true)
			}
			ndb.updateDelHash(batch, mpk, dep)
		}
	}
	batch.Write()

	//加入要删除的hash节点
	count1 := 0
	for _, str := range delNodeStrs {
		mp[str] = true
		count1++
	}
	count := 0
	batch = db.NewBatch(true)
	for key := range mp {
		batch.Delete([]byte(key))
		count++
	}
	batch.Write()
	//fmt.Printf("pruningHashNode ndb.count %d delete %d \n", ndb.count, count1)
	treelog.Info("pruningHashNode ", "delNodeStrs", count1, "delete node mp count", count)
}

//获取要删除的hash节点
func (node *MarkNode) getHashNode(ndb *markNodeDB) (delNodeStrs []string) {
	for {
		parN := node.fetchParentNode(ndb)
		if parN != nil {
			delNodeStrs = append(delNodeStrs, string(node.hash))
			node = parN
		} else {
			delNodeStrs = append(delNodeStrs, string(node.hash))
			break
		}
	}
	return delNodeStrs
}

func (ndb *markNodeDB) updateDelHash(batch dbm.Batch, key string, dep *delNodeValuePool) {
	if dep == nil {
		return
	}
	//这里指针暂时不需要赋值
	//if ndb.delPoolCache != nil {
	//	ndb.delPoolCache.Add(key, dep)
	//}
	stp := &types.StoreValuePool{}
	for _, k := range dep.delCache.Keys() {
		stp.Values = append(stp.Values, []byte(k.(string)))
	}
	v, err := proto.Marshal(stp)
	if err != nil {
		panic(fmt.Sprint("types.DeleteNodeMap fail", err))
	}
	batch.Set([]byte(key), v)
}

func (ndb *markNodeDB) getPool(str string) (dep *delNodeValuePool) {
	if ndb.delPoolCache != nil {
		elem, ok := ndb.delPoolCache.Get(str)
		if ok {
			dep = elem.(*delNodeValuePool)
			return dep
		}
	}
	v, err := ndb.db.Get([]byte(str))
	if err != nil || len(v) == 0 {
		//如果不存在说明是新的则
		//创建一个空的集
		ndb.judgeDelNodeCache()
<<<<<<< HEAD
		dep = NewDelNodeValuePool(perDelNodePoolSize)
=======
		dep = newDelNodeValuePool(perDelNodePoolSize)
>>>>>>> upstream/master
		if dep != nil {
			ndb.delPoolCache.Add(str, dep)
		}
	} else {
		stp := &types.StoreValuePool{}
		err = proto.Unmarshal(v, stp)
		if err != nil {
			panic(fmt.Sprint("types.StoreValuePool fail", err))
		}
		if ndb.delPoolCache != nil {
			ndb.judgeDelNodeCache()
<<<<<<< HEAD
			dep = NewDelNodeValuePool(perDelNodePoolSize)
=======
			dep = newDelNodeValuePool(perDelNodePoolSize)
>>>>>>> upstream/master
			if dep != nil {
				for _, k := range stp.Values {
					dep.delCache.Add(string(k), true)
				}
				ndb.delPoolCache.Add(str, dep)
			}
		}
	}
	return dep
}

func (ndb *markNodeDB) judgeDelNodeCache() {
	if ndb.delPoolCache.Len() >= delNodeCacheSize {
		strs := ndb.delPoolCache.Keys()
		elem, ok := ndb.delPoolCache.Get(strs[0])
		if ok {
			mp := elem.(*delNodeValuePool)
			stp := &types.StoreValuePool{}
			for _, k := range mp.delCache.Keys() {
				stp.Values = append(stp.Values, []byte(k.(string)))
			}
			v, err := proto.Marshal(stp)
			if err != nil {
				panic(fmt.Sprint("types.StoreValuePool fail", err))
			}
			ndb.db.Set([]byte(strs[0].(string)), v)
		}
	}
}

<<<<<<< HEAD
=======
// MarkNode 用于裁剪的节点结构体
>>>>>>> upstream/master
type MarkNode struct {
	height     int32
	hash       []byte
	leftHash   []byte
	rightHash  []byte
	parentHash []byte
	parentNode *MarkNode
}

type markNodeDB struct {
<<<<<<< HEAD
	mtx          sync.Mutex
=======
	//mtx          sync.Mutex
>>>>>>> upstream/master
	cache        *lru.Cache // 缓存当前批次已经删除的节点,
	delPoolCache *lru.Cache // 缓存全部的已经删除的节点
	db           dbm.DB
}

func newMarkNodeDB(db dbm.DB, cache int) *markNodeDB {
	cach, _ := lru.New(cache)
	ndb := &markNodeDB{
		cache:        cach,
		delPoolCache: delPoolCache,
		db:           db,
	}
	return ndb
}

<<<<<<< HEAD
=======
// LoadLeaf 载入叶子节点
>>>>>>> upstream/master
func (ndb *markNodeDB) LoadLeaf(hash []byte) (node *MarkNode, err error) {
	if !bytes.Equal(hash, emptyRoot[:]) {
		leaf, err := ndb.fetchNode(hash)
		return leaf, err
	}
	return nil, types.ErrNotFound
}

func (node *MarkNode) fetchParentNode(ndb *markNodeDB) *MarkNode {
	if node.parentNode != nil {
		return node.parentNode
<<<<<<< HEAD
	} else {
		pNode, err := ndb.fetchNode(node.parentHash)
		if err != nil {
			return nil
		}
		return pNode
	}
=======
	}
	pNode, err := ndb.fetchNode(node.parentHash)
	if err != nil {
		return nil
	}
	return pNode
>>>>>>> upstream/master
}

func (ndb *markNodeDB) fetchNode(hash []byte) (*MarkNode, error) {
	if len(hash) == 0 {
		return nil, ErrNodeNotExist
	}
<<<<<<< HEAD
	ndb.mtx.Lock()
	defer ndb.mtx.Unlock()
=======
	//ndb.mtx.Lock()
	//defer ndb.mtx.Unlock()
>>>>>>> upstream/master

	var mNode *MarkNode
	// cache
	if ndb.cache != nil {
		_, ok := ndb.cache.Get(string(hash))
		if ok {
			//缓存已经删除的节点,
			return nil, ErrNodeNotExist
		}
	}
	if mNode == nil {
		// 先判断是否已经删除掉,如果删除掉查找比较耗时
		key, hsh := genDeletePoolKey(hash)
		mp := ndb.getPool(string(key))
		if mp != nil {
			if _, ok := mp.delCache.Get(string(hsh)); ok {
				return nil, ErrNodeNotExist
			}
		}
		var buf []byte
		buf, err := ndb.db.Get(hash)
		if len(buf) == 0 || err != nil {
			return nil, err
		}
		node, err := MakeNode(buf, nil)
		if err != nil {
			panic(fmt.Sprintf("Error reading IAVLNode. bytes: %X  error: %v", buf, err))
		}
		node.hash = hash
		mNode = &MarkNode{
			height:     node.height,
			hash:       node.hash,
			leftHash:   node.leftHash,
			rightHash:  node.rightHash,
			parentHash: node.parentHash,
		}
		if ndb.cache != nil {
			ndb.cache.Add(string(hash), mNode)
		}
	}
	return mNode, nil
}

<<<<<<< HEAD
=======
func pruningSecondLevel(db dbm.DB, curHeight int64) {
	if secLvlPruningH == 0 {
		secLvlPruningH = getSecLvlPruningHeight(db)
	}
	if curHeight/secondLevelPruningHeight > 1 &&
		curHeight/secondLevelPruningHeight != secLvlPruningH/secondLevelPruningHeight {
		treelog.Info("pruningTree pruningSecondLevel", "start curHeight:", curHeight)
		start := time.Now()
		pruningSecondLevelNode(db, curHeight)
		end := time.Now()
		treelog.Info("pruningTree pruningSecondLevel", "curHeight:", curHeight, "pruning leafNode cost time:", end.Sub(start))
		setSecLvlPruningHeight(db, curHeight)
		secLvlPruningH = curHeight
	}
}

func pruningSecondLevelNode(db dbm.DB, curHeight int64) {
	prefix := []byte(oldLeafKeyCountPrefix)
	it := db.Iterator(prefix, nil, true)
	defer it.Close()

	mp := make(map[string][]hashData)
	count := 0
	for it.Rewind(); it.Valid(); it.Next() {
		if quit {
			//该处退出
			return
		}
		//copy key
		hashK := make([]byte, len(it.Key()))
		copy(hashK, it.Key())

		value := it.Value()
		var pData types.PruneData
		err := proto.Unmarshal(value, &pData)
		if err != nil {
			panic("Unmarshal mavl leafCountKey fail")
		}
		hashLen := int(pData.Lenth)
		key, err := getKeyFromOldLeafCountKey(hashK, hashLen)
		if err == nil {
			data := hashData{
				height: pData.Height,
				hash:   hashK[len(hashK)-hashLen:],
			}
			mp[string(key)] = append(mp[string(key)], data)
			count++
			if count >= onceScanCount {
				deleteOldNode(db, mp, curHeight, key)
				count = 0
			}
		}
	}
	if count > 0 {
		deleteOldNode(db, mp, curHeight, nil)
	}
}

func deleteOldNode(db dbm.DB, mp map[string][]hashData, curHeight int64, lastKey []byte) {
	if len(mp) == 0 {
		return
	}
	var tmp []hashData
	//del
	if lastKey != nil {
		if _, ok := mp[string(lastKey)]; ok {
			tmp = mp[string(lastKey)]
			delete(mp, string(lastKey))
		}
	}
	delMp := make(map[string]bool)
	batch := db.NewBatch(true)
	for key, vals := range mp {
		if len(vals) > 1 {
			if vals[1].height != vals[0].height { //防止相同高度时候出现的误删除
				for _, val := range vals[1:] { //从第二个开始判断
					if curHeight >= val.height+int64(pruneHeight) {
						batch.Delete(genOldLeafCountKey([]byte(key), val.hash, val.height))
						delMp[string(val.hash)] = true
					}
				}
			} else {
				// 删除第三层存储索引key
				for _, val := range vals {
					if curHeight >= val.height+threeLevelPruningHeight {
						batch.Delete(genOldLeafCountKey([]byte(key), val.hash, val.height))
					}
				}
			}
		} else if len(vals) == 1 && curHeight >= vals[0].height+threeLevelPruningHeight { // 删除第三层存储索引key
			batch.Delete(genLeafCountKey([]byte(key), vals[0].hash, vals[0].height))
		}
		delete(mp, key)
	}
	batch.Write()
	//add
	if lastKey != nil {
		if _, ok := mp[string(lastKey)]; ok {
			mp[string(lastKey)] = tmp
		}
	}
	//裁剪hashNode
	pruningHashNode(db, delMp)
}

// PruningTreePrintDB pruning tree print db
>>>>>>> upstream/master
func PruningTreePrintDB(db dbm.DB, prefix []byte) {
	it := db.Iterator(prefix, nil, true)
	defer it.Close()
	count := 0
	for it.Rewind(); it.Valid(); it.Next() {
		if bytes.Equal(prefix, []byte(leafKeyCountPrefix)) {
			hashK := it.Key()
			value := it.Value()
			var pData types.PruneData
			err := proto.Unmarshal(value, &pData)
			if err == nil {
				hashLen := int(pData.Lenth)
				key, err := getKeyFromLeafCountKey(hashK, hashLen)
				if err == nil {
					treelog.Debug("pruningTree:", "key:", string(key), "height", pData.Height)
				}
			}
		} else if bytes.Equal(prefix, []byte(hashNodePrefix)) {
			treelog.Debug("pruningTree:", "key:", string(it.Key()))
		} else if bytes.Equal(prefix, []byte(leafNodePrefix)) {
			treelog.Debug("pruningTree:", "key:", string(it.Key()))
		} else if bytes.Equal(prefix, []byte(delMapPoolPrefix)) {
			value := it.Value()
			var pData types.StoreValuePool
			err := proto.Unmarshal(value, &pData)
			if err == nil {
				for _, k := range pData.Values {
					treelog.Debug("delMapPool value ", "hash:", common.Bytes2Hex(k[:2]))
				}
			}
		}
		count++
	}
	fmt.Printf("prefix %s All count:%d \n", string(prefix), count)
	treelog.Info("pruningTree:", "prefix:", string(prefix), "All count", count)
}

<<<<<<< HEAD
=======
// PruningTree 裁剪树
>>>>>>> upstream/master
func PruningTree(db dbm.DB, curHeight int64) {
	pruningTree(db, curHeight)
}
