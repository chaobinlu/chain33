// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package db 数据库操作底层接口定义以及实现包括：leveldb、
// memdb、mvcc、badgerdb、pegasus、ssdb
>>>>>>> upstream/master
package db

import (
	"bytes"
<<<<<<< HEAD
	"errors"
	"fmt"

	lru "github.com/hashicorp/golang-lru"
	"github.com/33cn/chain33/types"
)

var ErrNotFoundInDb = errors.New("ErrNotFoundInDb")

=======
	"fmt"

	"github.com/33cn/chain33/types"
	lru "github.com/hashicorp/golang-lru"
)

//ErrNotFoundInDb error
var ErrNotFoundInDb = types.ErrNotFound

//Lister 列表接口
>>>>>>> upstream/master
type Lister interface {
	List(prefix, key []byte, count, direction int32) ([][]byte, error)
	PrefixCount(prefix []byte) int64
}

<<<<<<< HEAD
=======
//KV kv
>>>>>>> upstream/master
type KV interface {
	Get(key []byte) ([]byte, error)
	BatchGet(keys [][]byte) (values [][]byte, err error)
	Set(key []byte, value []byte) (err error)
	Begin()
	Rollback()
	Commit()
}

<<<<<<< HEAD
=======
//KVDB kvdb
>>>>>>> upstream/master
type KVDB interface {
	KV
	Lister
}

<<<<<<< HEAD
=======
//DB db
>>>>>>> upstream/master
type DB interface {
	KV
	IteratorDB
	SetSync([]byte, []byte) error
	Delete([]byte) error
	DeleteSync([]byte) error
	Close()
	NewBatch(sync bool) Batch
	// For debugging
	Print()
	Stats() map[string]string
	SetCacheSize(size int)
	GetCache() *lru.ARCCache
}

<<<<<<< HEAD
=======
//KVDBList list
>>>>>>> upstream/master
type KVDBList struct {
	DB
	list *ListHelper
}

<<<<<<< HEAD
=======
//List 列表
>>>>>>> upstream/master
func (l *KVDBList) List(prefix, key []byte, count, direction int32) ([][]byte, error) {
	vals := l.list.List(prefix, key, count, direction)
	if vals == nil {
		return nil, types.ErrNotFound
	}
	return vals, nil
}

<<<<<<< HEAD
=======
//PrefixCount 前缀长度
>>>>>>> upstream/master
func (l *KVDBList) PrefixCount(prefix []byte) int64 {
	return l.list.PrefixCount(prefix)
}

<<<<<<< HEAD
=======
//NewKVDB new
>>>>>>> upstream/master
func NewKVDB(db DB) KVDB {
	return &KVDBList{DB: db, list: NewListHelper(db)}
}

<<<<<<< HEAD
=======
//Batch batch
>>>>>>> upstream/master
type Batch interface {
	Set(key, value []byte)
	Delete(key []byte)
	Write() error
	ValueSize() int // amount of data in the batch
	Reset()         // Reset resets the batch for reuse
}

<<<<<<< HEAD
=======
//IteratorSeeker ...
>>>>>>> upstream/master
type IteratorSeeker interface {
	Rewind() bool
	Seek(key []byte) bool
	Next() bool
}

<<<<<<< HEAD
=======
//Iterator 迭代器
>>>>>>> upstream/master
type Iterator interface {
	IteratorSeeker
	Valid() bool
	Key() []byte
	Value() []byte
	ValueCopy() []byte
	Error() error
	Prefix() []byte
	Close()
}

type itBase struct {
	start   []byte
	end     []byte
	reverse bool
}

func (it *itBase) checkKey(key []byte) bool {
	//key must in start and end
	var startok = true
	var endok = true
	if it.start != nil {
		startok = bytes.Compare(key, it.start) >= 0
	}
	if it.end != nil {
		endok = bytes.Compare(key, it.end) <= 0
	}
	ok := startok && endok
	return ok
}

<<<<<<< HEAD
=======
//Prefix 前缀
>>>>>>> upstream/master
func (it *itBase) Prefix() []byte {
	return nil
}

<<<<<<< HEAD
=======
//IteratorDB 迭代
>>>>>>> upstream/master
type IteratorDB interface {
	Iterator(start []byte, end []byte, reserver bool) Iterator
}

func bytesPrefix(prefix []byte) []byte {
	var limit []byte
	for i := len(prefix) - 1; i >= 0; i-- {
		c := prefix[i]
		if c < 0xff {
			limit = make([]byte, i+1)
			copy(limit, prefix)
			limit[i] = c + 1
			break
		}
	}
	return limit
}

const (
	levelDBBackendStr     = "leveldb" // legacy, defaults to goleveldb.
	goLevelDBBackendStr   = "goleveldb"
	memDBBackendStr       = "memdb"
	goBadgerDBBackendStr  = "gobadgerdb"
	ssDBBackendStr        = "ssdb"
	goPegasusDbBackendStr = "pegasus"
)

type dbCreator func(name string, dir string, cache int) (DB, error)

var backends = map[string]dbCreator{}

func registerDBCreator(backend string, creator dbCreator, force bool) {
	_, ok := backends[backend]
	if !force && ok {
		return
	}
	backends[backend] = creator
}

<<<<<<< HEAD
=======
//NewDB new
>>>>>>> upstream/master
func NewDB(name string, backend string, dir string, cache int32) DB {
	dbCreator, ok := backends[backend]
	if !ok {
		fmt.Printf("Error initializing DB: %v\n", backend)
		panic("initializing DB error")
	}
	db, err := dbCreator(name, dir, int(cache))
	if err != nil {
		fmt.Printf("Error initializing DB: %v\n", err)
		panic("initializing DB error")
	}
	return db
}

<<<<<<< HEAD
=======
//TransactionDB 交易缓存
>>>>>>> upstream/master
type TransactionDB struct {
	cache *lru.ARCCache
}

<<<<<<< HEAD
=======
//Begin 启动
>>>>>>> upstream/master
func (db *TransactionDB) Begin() {

}

<<<<<<< HEAD
=======
//Rollback 回滚
>>>>>>> upstream/master
func (db *TransactionDB) Rollback() {

}

<<<<<<< HEAD
=======
//Commit 提交
>>>>>>> upstream/master
func (db *TransactionDB) Commit() {

}

<<<<<<< HEAD
=======
//GetCache 获取缓存
>>>>>>> upstream/master
func (db *TransactionDB) GetCache() *lru.ARCCache {
	return db.cache
}

<<<<<<< HEAD
=======
//SetCacheSize 设置缓存大小
>>>>>>> upstream/master
func (db *TransactionDB) SetCacheSize(size int) {
	if db.cache != nil {
		return
	}
	db.cache, _ = lru.NewARC(size)
}
