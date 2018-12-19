// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"bytes"
	"sync"

	"sort"
	"strings"

	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
)

var mlog = log.New("module", "db.memdb")

// memdb 应该无需区分同步与异步操作

func init() {
	dbCreator := func(name string, dir string, cache int) (DB, error) {
		return NewGoMemDB(name, dir, cache)
	}
	registerDBCreator(memDBBackendStr, dbCreator, false)
}

<<<<<<< HEAD
=======
//GoMemDB db
>>>>>>> upstream/master
type GoMemDB struct {
	TransactionDB
	db   map[string][]byte
	lock sync.RWMutex
}

<<<<<<< HEAD
=======
//NewGoMemDB new
>>>>>>> upstream/master
func NewGoMemDB(name string, dir string, cache int) (*GoMemDB, error) {
	// memdb 不需要创建文件，后续考虑增加缓存数目
	return &GoMemDB{
		db: make(map[string][]byte),
	}, nil
}

<<<<<<< HEAD
=======
//CopyBytes 复制字节
>>>>>>> upstream/master
func CopyBytes(b []byte) (copiedBytes []byte) {
	/* 兼容leveldb
	if b == nil {
		return nil
	}
	*/
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)

	return copiedBytes
}

<<<<<<< HEAD
=======
//Get get
>>>>>>> upstream/master
func (db *GoMemDB) Get(key []byte) ([]byte, error) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	if entry, ok := db.db[string(key)]; ok {
		return CopyBytes(entry), nil
	}
	return nil, ErrNotFoundInDb
}

<<<<<<< HEAD
=======
//Set set
>>>>>>> upstream/master
func (db *GoMemDB) Set(key []byte, value []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	//debug.PrintStack()
	//println("--", string(key)[0:4], common.ToHex(key))
	db.db[string(key)] = CopyBytes(value)
	if db.db[string(key)] == nil {
		mlog.Error("Set", "error have no mem")
	}
	return nil
}

<<<<<<< HEAD
=======
//SetSync 设置同步
>>>>>>> upstream/master
func (db *GoMemDB) SetSync(key []byte, value []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	//debug.PrintStack()
	//println("--", string(key)[0:4], common.ToHex(key))
	db.db[string(key)] = CopyBytes(value)
	if db.db[string(key)] == nil {
		mlog.Error("Set", "error have no mem")
	}
	return nil
}

<<<<<<< HEAD
=======
//Delete 删除
>>>>>>> upstream/master
func (db *GoMemDB) Delete(key []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	delete(db.db, string(key))
	return nil
}

<<<<<<< HEAD
=======
//DeleteSync 删除同步
>>>>>>> upstream/master
func (db *GoMemDB) DeleteSync(key []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	delete(db.db, string(key))
	return nil
}

<<<<<<< HEAD
=======
//DB db
>>>>>>> upstream/master
func (db *GoMemDB) DB() map[string][]byte {
	return db.db
}

<<<<<<< HEAD
=======
//Close 关闭
>>>>>>> upstream/master
func (db *GoMemDB) Close() {

}

<<<<<<< HEAD
=======
//Print 打印
>>>>>>> upstream/master
func (db *GoMemDB) Print() {
	for key, value := range db.db {
		mlog.Info("Print", "key", key, "value", string(value))
	}
}

<<<<<<< HEAD
=======
//Stats ...
>>>>>>> upstream/master
func (db *GoMemDB) Stats() map[string]string {
	//TODO
	return nil
}

<<<<<<< HEAD
=======
//Iterator 迭代器
>>>>>>> upstream/master
func (db *GoMemDB) Iterator(start []byte, end []byte, reverse bool) Iterator {
	db.lock.RLock()
	defer db.lock.RUnlock()
	if end == nil {
		end = bytesPrefix(start)
	}
	if bytes.Equal(end, types.EmptyValue) {
		end = nil
	}
	base := itBase{start, end, reverse}

	var keys []string
	for k := range db.db {
		if base.checkKey([]byte(k)) {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	var index int
	return &goMemDBIt{base, index, keys, db}
}

<<<<<<< HEAD
=======
//BatchGet 批量获取
>>>>>>> upstream/master
func (db *GoMemDB) BatchGet(keys [][]byte) (value [][]byte, err error) {
	mlog.Error("BatchGet", "Need to implement")
	return nil, nil
}

type goMemDBIt struct {
	itBase
	index   int      // 记录当前索引
	keys    []string // 记录所有keys值
	goMemDb *GoMemDB
}

<<<<<<< HEAD
=======
//Seek 查找
>>>>>>> upstream/master
func (dbit *goMemDBIt) Seek(key []byte) bool { //指向当前的index值
	for i, k := range dbit.keys {
		if 0 == strings.Compare(k, string(key)) {
			dbit.index = i
			return true
		}
	}
	return false
}

<<<<<<< HEAD
=======
//Close 关闭
>>>>>>> upstream/master
func (dbit *goMemDBIt) Close() {
	dbit.goMemDb.Close()
}

<<<<<<< HEAD
=======
//Next next
>>>>>>> upstream/master
func (dbit *goMemDBIt) Next() bool {
	if dbit.reverse { // 反向
		dbit.index-- //将当前key值指向前一个
		return dbit.Valid()
<<<<<<< HEAD
	} else { // 正向
		dbit.index++ //将当前key值指向后一个
		return dbit.Valid()
	}
}

=======
	}
	// 正向
	dbit.index++ //将当前key值指向后一个
	return dbit.Valid()

}

//Rewind ...
>>>>>>> upstream/master
func (dbit *goMemDBIt) Rewind() bool {
	if dbit.reverse { // 反向
		if (len(dbit.keys) > 0) && dbit.Valid() {
			dbit.index = len(dbit.keys) - 1 // 将当前key值指向最后一个
			return true
<<<<<<< HEAD
		} else {
			return false
		}
	} else { // 正向
		if dbit.Valid() {
			dbit.index = 0 // 将当前key值指向第一个
			return true
		} else {
			return false
		}
	}
}

=======
		}
		return false
	}
	// 正向
	if dbit.Valid() {
		dbit.index = 0 // 将当前key值指向第一个
		return true
	}
	return false
}

//Key key
>>>>>>> upstream/master
func (dbit *goMemDBIt) Key() []byte {
	return []byte(dbit.keys[dbit.index])
}

<<<<<<< HEAD
=======
//Value value
>>>>>>> upstream/master
func (dbit *goMemDBIt) Value() []byte {
	value, _ := dbit.goMemDb.Get([]byte(dbit.keys[dbit.index]))
	return value
}

func (dbit *goMemDBIt) ValueCopy() []byte {
	v, _ := dbit.goMemDb.Get([]byte(dbit.keys[dbit.index]))
	value := make([]byte, len(v))
	copy(value, v)
	return value
}

func (dbit *goMemDBIt) Valid() bool {

	if (dbit.goMemDb == nil) && (len(dbit.keys) == 0) {
		return false
	}

	if len(dbit.keys) > dbit.index && dbit.index >= 0 {
		return true
<<<<<<< HEAD
	} else {
		return false
	}
=======
	}
	return false

>>>>>>> upstream/master
}

func (dbit *goMemDBIt) Error() error {
	return nil
}

type kv struct{ k, v []byte }
type memBatch struct {
	db     *GoMemDB
	writes []kv
	size   int
}

<<<<<<< HEAD
=======
//NewBatch new
>>>>>>> upstream/master
func (db *GoMemDB) NewBatch(sync bool) Batch {
	return &memBatch{db: db}
}

func (b *memBatch) Set(key, value []byte) {
	//println("-b-", string(key)[0:4], common.ToHex(key))
	b.writes = append(b.writes, kv{CopyBytes(key), CopyBytes(value)})
	b.size += len(value)
}

func (b *memBatch) Delete(key []byte) {
	b.writes = append(b.writes, kv{CopyBytes(key), CopyBytes(nil)})
<<<<<<< HEAD
	b.size += 1
=======
	b.size++
>>>>>>> upstream/master
}

func (b *memBatch) Write() error {
	b.db.lock.Lock()
	defer b.db.lock.Unlock()

	for _, kv := range b.writes {
		if kv.v == nil {
			//println("[d]", string(kv.k))
			delete(b.db.db, string(kv.k))
		} else {
			//println("[i]", string(kv.k))
			b.db.db[string(kv.k)] = kv.v
		}
	}
	return nil
}

func (b *memBatch) ValueSize() int {
	return b.size
}

func (b *memBatch) Reset() {
	b.writes = b.writes[:0]
	b.size = 0
}
