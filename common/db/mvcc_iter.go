// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"bytes"

	"github.com/33cn/chain33/types"
)

<<<<<<< HEAD
//mvcc 迭代器版本
=======
//MVCCIter mvcc迭代器版本
>>>>>>> upstream/master
//支持db 原生迭代器接口
//为了支持快速迭代，我这里采用了复制数据的做法
type MVCCIter struct {
	*MVCCHelper
}

<<<<<<< HEAD
=======
//NewMVCCIter new
>>>>>>> upstream/master
func NewMVCCIter(db DB) *MVCCIter {
	return &MVCCIter{MVCCHelper: NewMVCC(db)}
}

<<<<<<< HEAD
=======
//AddMVCC add
>>>>>>> upstream/master
func (m *MVCCIter) AddMVCC(kvs []*types.KeyValue, hash []byte, prevHash []byte, version int64) ([]*types.KeyValue, error) {
	kvlist, err := m.MVCCHelper.AddMVCC(kvs, hash, prevHash, version)
	if err != nil {
		return nil, err
	}
	//添加last
	for _, v := range kvs {
		last := getLastKey(v.Key)
		kv := &types.KeyValue{Key: last, Value: v.Value}
		kvlist = append(kvlist, kv)
	}
	return kvlist, nil
}

<<<<<<< HEAD
=======
//DelMVCC del
>>>>>>> upstream/master
func (m *MVCCIter) DelMVCC(hash []byte, version int64, strict bool) ([]*types.KeyValue, error) {
	kvs, err := m.GetDelKVList(version)
	if err != nil {
		return nil, err
	}
	kvlist, err := m.MVCCHelper.delMVCC(kvs, hash, version, strict)
	if err != nil {
		return nil, err
	}
	//更新last, 读取上次版本的 lastv值，更新last
	for _, v := range kvs {
		if version > 0 {
			lastv, err := m.GetV(v.Key, version-1)
			if err == types.ErrNotFound {
				kvlist = append(kvlist, &types.KeyValue{Key: getLastKey(v.Key)})
				continue
			}
			if err != nil {
				return nil, err
			}
			kvlist = append(kvlist, &types.KeyValue{Key: getLastKey(v.Key), Value: lastv})
		}
	}
	return kvlist, nil
}

<<<<<<< HEAD
=======
//Iterator 迭代
>>>>>>> upstream/master
func (m *MVCCIter) Iterator(start, end []byte, reserver bool) Iterator {
	if start == nil {
		start = mvccLast
	} else {
		start = getLastKey(start)
	}
	if end != nil {
		end = getLastKey(end)
	} else {
		end = bytesPrefix(start)
	}
	return &mvccIt{m.db.Iterator(start, end, reserver)}
}

type mvccIt struct {
	Iterator
}

<<<<<<< HEAD
=======
//Prefix 前缀
>>>>>>> upstream/master
func (dbit *mvccIt) Prefix() []byte {
	return mvccLast
}

<<<<<<< HEAD
=======
//Key key
>>>>>>> upstream/master
func (dbit *mvccIt) Key() []byte {
	key := dbit.Iterator.Key()
	if bytes.HasPrefix(key, dbit.Prefix()) {
		return key[len(dbit.Prefix()):]
	}
	return nil
}

<<<<<<< HEAD
=======
//Valid 检查合法性
>>>>>>> upstream/master
func (dbit *mvccIt) Valid() bool {
	if !dbit.Iterator.Valid() {
		return false
	}
	return dbit.Key() != nil
}
