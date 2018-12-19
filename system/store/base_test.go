// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"testing"

	"os"

<<<<<<< HEAD
	"github.com/stretchr/testify/assert"
	"github.com/33cn/chain33/common/log"
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
)

var store_cfg0 = &types.Store{"base_test", "leveldb", "/tmp/base_test0", 100}
var store_cfg1 = &types.Store{"base_test", "leveldb", "/tmp/base_test1", 100}
=======
	"github.com/33cn/chain33/common/log"
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
)

var storecfg0 = &types.Store{Name: "base_test", Driver: "leveldb", DbPath: "/tmp/base_test0", DbCache: 100}
var storecfg1 = &types.Store{Name: "base_test", Driver: "leveldb", DbPath: "/tmp/base_test1", DbCache: 100}
>>>>>>> upstream/master

type storeChild struct {
}

func (s *storeChild) Set(datas *types.StoreSet, sync bool) ([]byte, error) {
	return []byte{}, nil
}

func (s *storeChild) Get(datas *types.StoreGet) [][]byte {
	return [][]byte{}
}

func (s *storeChild) MemSet(datas *types.StoreSet, sync bool) ([]byte, error) {
	return []byte{}, nil
}

func (s *storeChild) Commit(hash *types.ReqHash) ([]byte, error) {
	return []byte{}, nil
}

func (s *storeChild) Rollback(req *types.ReqHash) ([]byte, error) {
	return []byte{}, nil
}

func (s *storeChild) Del(req *types.StoreDel) ([]byte, error) {
	return []byte{}, nil
}

func (s *storeChild) IterateRangeByStateHash(statehash []byte, start []byte, end []byte, ascending bool, fn func(key, value []byte) bool) {

}

func (s *storeChild) ProcEvent(msg queue.Message) {}

func init() {
	log.SetLogLevel("error")
}

func TestBaseStore_NewClose(t *testing.T) {
<<<<<<< HEAD
	os.RemoveAll(store_cfg0.DbPath)
	store := NewBaseStore(store_cfg0)
=======
	os.RemoveAll(storecfg0.DbPath)
	store := NewBaseStore(storecfg0)
>>>>>>> upstream/master
	assert.NotNil(t, store)

	db := store.GetDB()
	assert.NotNil(t, db)

	store.Close()
}

func TestBaseStore_Queue(t *testing.T) {
<<<<<<< HEAD
	os.RemoveAll(store_cfg1.DbPath)
	store := NewBaseStore(store_cfg1)
=======
	os.RemoveAll(storecfg1.DbPath)
	store := NewBaseStore(storecfg1)
>>>>>>> upstream/master
	assert.NotNil(t, store)

	var q = queue.New("channel")
	store.SetQueueClient(q.Client())
	queueClinet := store.GetQueueClient()

	child := &storeChild{}
	store.SetChild(child)

	var kv []*types.KeyValue
<<<<<<< HEAD
	kv = append(kv, &types.KeyValue{[]byte("k1"), []byte("v1")})
	kv = append(kv, &types.KeyValue{[]byte("k2"), []byte("v2")})
	datas := &types.StoreSet{
		EmptyRoot[:],
		kv,
		0}
	set := &types.StoreSetWithSync{datas, true}
=======
	kv = append(kv, &types.KeyValue{Key: []byte("k1"), Value: []byte("v1")})
	kv = append(kv, &types.KeyValue{Key: []byte("k2"), Value: []byte("v2")})
	datas := &types.StoreSet{
		StateHash: EmptyRoot[:],
		KV:        kv,
	}
	set := &types.StoreSetWithSync{Storeset: datas, Sync: true}
>>>>>>> upstream/master
	msg := queueClinet.NewMessage("store", types.EventStoreSet, set)
	err := queueClinet.Send(msg, true)
	assert.Nil(t, err)
	resp, err := queueClinet.Wait(msg)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(types.EventStoreSetReply), resp.Ty)

<<<<<<< HEAD
	get := &types.StoreGet{EmptyRoot[:], [][]byte{}}
=======
	get := &types.StoreGet{StateHash: EmptyRoot[:], Keys: [][]byte{}}
>>>>>>> upstream/master
	msg = queueClinet.NewMessage("store", types.EventStoreGet, get)
	err = queueClinet.Send(msg, true)
	assert.Nil(t, err)
	resp, err = queueClinet.Wait(msg)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(types.EventStoreGetReply), resp.Ty)

	memset := set
	msg = queueClinet.NewMessage("store", types.EventStoreMemSet, memset)
	err = queueClinet.Send(msg, true)
	assert.Nil(t, err)
	resp, err = queueClinet.Wait(msg)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(types.EventStoreSetReply), resp.Ty)

<<<<<<< HEAD
	commit := &types.ReqHash{EmptyRoot[:]}
=======
	commit := &types.ReqHash{Hash: EmptyRoot[:]}
>>>>>>> upstream/master
	msg = queueClinet.NewMessage("store", types.EventStoreCommit, commit)
	err = queueClinet.Send(msg, true)
	assert.Nil(t, err)
	resp, err = queueClinet.Wait(msg)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(types.EventStoreCommit), resp.Ty)

<<<<<<< HEAD
	rollback := &types.ReqHash{EmptyRoot[:]}
=======
	rollback := &types.ReqHash{Hash: EmptyRoot[:]}
>>>>>>> upstream/master
	msg = queueClinet.NewMessage("store", types.EventStoreRollback, rollback)
	err = queueClinet.Send(msg, true)
	assert.Nil(t, err)
	resp, err = queueClinet.Wait(msg)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(types.EventStoreRollback), resp.Ty)

<<<<<<< HEAD
	totalCoins := &types.IterateRangeByStateHash{EmptyRoot[:], []byte(""), []byte(""), 100}
=======
	totalCoins := &types.IterateRangeByStateHash{
		StateHash: EmptyRoot[:],
		Start:     []byte(""),
		End:       []byte(""),
		Count:     100,
	}
>>>>>>> upstream/master
	msg = queueClinet.NewMessage("store", types.EventStoreGetTotalCoins, totalCoins)
	err = queueClinet.Send(msg, true)
	assert.Nil(t, err)
	resp, err = queueClinet.Wait(msg)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(types.EventGetTotalCoinsReply), resp.Ty)

}
