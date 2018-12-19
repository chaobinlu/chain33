// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"bytes"
	"path"

<<<<<<< HEAD
=======
	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
>>>>>>> upstream/master
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
<<<<<<< HEAD
	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
=======
>>>>>>> upstream/master
)

var llog = log.New("module", "db.goleveldb")

func init() {
	dbCreator := func(name string, dir string, cache int) (DB, error) {
		return NewGoLevelDB(name, dir, cache)
	}
	registerDBCreator(levelDBBackendStr, dbCreator, false)
	registerDBCreator(goLevelDBBackendStr, dbCreator, false)
}

<<<<<<< HEAD
=======
//GoLevelDB db
>>>>>>> upstream/master
type GoLevelDB struct {
	TransactionDB
	db *leveldb.DB
}

<<<<<<< HEAD
=======
//NewGoLevelDB new
>>>>>>> upstream/master
func NewGoLevelDB(name string, dir string, cache int) (*GoLevelDB, error) {
	dbPath := path.Join(dir, name+".db")
	if cache == 0 {
		cache = 64
	}
	handles := cache
	if handles < 16 {
		handles = 16
	}
	if cache < 4 {
		cache = 4
	}
	// Open the db and recover any potential corruptions
	db, err := leveldb.OpenFile(dbPath, &opt.Options{
		OpenFilesCacheCapacity: handles,
		BlockCacheCapacity:     cache / 2 * opt.MiB,
		WriteBuffer:            cache / 4 * opt.MiB, // Two of these are used internally
		Filter:                 filter.NewBloomFilter(10),
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(dbPath, nil)
	}
	if err != nil {
		return nil, err
	}
	database := &GoLevelDB{db: db}
	return database, nil
}

<<<<<<< HEAD
=======
//Get get
>>>>>>> upstream/master
func (db *GoLevelDB) Get(key []byte) ([]byte, error) {
	res, err := db.db.Get(key, nil)
	if err != nil {
		if err == errors.ErrNotFound {
			return nil, ErrNotFoundInDb
<<<<<<< HEAD
		} else {
			llog.Error("Get", "error", err)
			return nil, err
		}
=======
		}
		llog.Error("Get", "error", err)
		return nil, err

>>>>>>> upstream/master
	}
	return res, nil
}

<<<<<<< HEAD
=======
//Set set
>>>>>>> upstream/master
func (db *GoLevelDB) Set(key []byte, value []byte) error {
	err := db.db.Put(key, value, nil)
	if err != nil {
		llog.Error("Set", "error", err)
		return err
	}
	return nil
}

<<<<<<< HEAD
=======
//SetSync 同步
>>>>>>> upstream/master
func (db *GoLevelDB) SetSync(key []byte, value []byte) error {
	err := db.db.Put(key, value, &opt.WriteOptions{Sync: true})
	if err != nil {
		llog.Error("SetSync", "error", err)
		return err
	}
	return nil
}

<<<<<<< HEAD
=======
//Delete 删除
>>>>>>> upstream/master
func (db *GoLevelDB) Delete(key []byte) error {
	err := db.db.Delete(key, nil)
	if err != nil {
		llog.Error("Delete", "error", err)
		return err
	}
	return nil
}

<<<<<<< HEAD
=======
//DeleteSync 删除同步
>>>>>>> upstream/master
func (db *GoLevelDB) DeleteSync(key []byte) error {
	err := db.db.Delete(key, &opt.WriteOptions{Sync: true})
	if err != nil {
		llog.Error("DeleteSync", "error", err)
		return err
	}
	return nil
}

<<<<<<< HEAD
=======
//DB db
>>>>>>> upstream/master
func (db *GoLevelDB) DB() *leveldb.DB {
	return db.db
}

<<<<<<< HEAD
=======
//Close 关闭
>>>>>>> upstream/master
func (db *GoLevelDB) Close() {
	db.db.Close()
}

<<<<<<< HEAD
=======
//Print 打印
>>>>>>> upstream/master
func (db *GoLevelDB) Print() {
	str, _ := db.db.GetProperty("leveldb.stats")
	llog.Info("Print", "stats", str)

	iter := db.db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		//fmt.Printf("[%X]:\t[%X]\n", key, value)
		llog.Info("Print", "key", string(key), "value", string(value))
	}
}

<<<<<<< HEAD
=======
//Stats ...
>>>>>>> upstream/master
func (db *GoLevelDB) Stats() map[string]string {
	keys := []string{
		"leveldb.num-files-at-level{n}",
		"leveldb.stats",
		"leveldb.sstables",
		"leveldb.blockpool",
		"leveldb.cachedblock",
		"leveldb.openedtables",
		"leveldb.alivesnaps",
		"leveldb.aliveiters",
	}

	stats := make(map[string]string)
	for _, key := range keys {
		str, err := db.db.GetProperty(key)
		if err == nil {
			stats[key] = str
		}
	}
	return stats
}

<<<<<<< HEAD
=======
//Iterator 迭代器
>>>>>>> upstream/master
func (db *GoLevelDB) Iterator(start []byte, end []byte, reverse bool) Iterator {
	if end == nil {
		end = bytesPrefix(start)
	}
	if bytes.Equal(end, types.EmptyValue) {
		end = nil
	}
<<<<<<< HEAD
	r := &util.Range{start, end}
=======
	r := &util.Range{Start: start, Limit: end}
>>>>>>> upstream/master
	it := db.db.NewIterator(r, nil)
	return &goLevelDBIt{it, itBase{start, end, reverse}}
}

<<<<<<< HEAD
=======
//BatchGet 批量获取
>>>>>>> upstream/master
func (db *GoLevelDB) BatchGet(keys [][]byte) (value [][]byte, err error) {
	llog.Error("BatchGet", "Need to implement")
	return nil, nil
}

type goLevelDBIt struct {
	iterator.Iterator
	itBase
}

<<<<<<< HEAD
=======
//Close 关闭
>>>>>>> upstream/master
func (dbit *goLevelDBIt) Close() {
	dbit.Iterator.Release()
}

<<<<<<< HEAD
=======
//Next next
>>>>>>> upstream/master
func (dbit *goLevelDBIt) Next() bool {
	if dbit.reverse {
		return dbit.Iterator.Prev() && dbit.Valid()
	}
	return dbit.Iterator.Next() && dbit.Valid()
}

<<<<<<< HEAD
=======
//Rewind ...
>>>>>>> upstream/master
func (dbit *goLevelDBIt) Rewind() bool {
	if dbit.reverse {
		return dbit.Iterator.Last() && dbit.Valid()
	}
	return dbit.Iterator.First() && dbit.Valid()
}

func (dbit *goLevelDBIt) Value() []byte {
	return dbit.Iterator.Value()
}

func cloneByte(v []byte) []byte {
	value := make([]byte, len(v))
	copy(value, v)
	return value
}

func (dbit *goLevelDBIt) ValueCopy() []byte {
	v := dbit.Iterator.Value()
	return cloneByte(v)
}

func (dbit *goLevelDBIt) Valid() bool {
	return dbit.Iterator.Valid() && dbit.checkKey(dbit.Key())
}

type goLevelDBBatch struct {
	db    *GoLevelDB
	batch *leveldb.Batch
	wop   *opt.WriteOptions
	size  int
}

<<<<<<< HEAD
=======
//NewBatch new
>>>>>>> upstream/master
func (db *GoLevelDB) NewBatch(sync bool) Batch {
	batch := new(leveldb.Batch)
	wop := &opt.WriteOptions{Sync: sync}
	return &goLevelDBBatch{db, batch, wop, 0}
}

func (mBatch *goLevelDBBatch) Set(key, value []byte) {
	mBatch.batch.Put(key, value)
	mBatch.size += len(value)
}

func (mBatch *goLevelDBBatch) Delete(key []byte) {
	mBatch.batch.Delete(key)
<<<<<<< HEAD
	mBatch.size += 1
=======
	mBatch.size++
>>>>>>> upstream/master
}

func (mBatch *goLevelDBBatch) Write() error {
	err := mBatch.db.db.Write(mBatch.batch, mBatch.wop)
	if err != nil {
		llog.Error("Write", "error", err)
		return err
	}
	return nil
}

func (mBatch *goLevelDBBatch) ValueSize() int {
	return mBatch.size
}

func (mBatch *goLevelDBBatch) Reset() {
	mBatch.batch.Reset()
	mBatch.size = 0
}
