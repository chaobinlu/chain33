// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"sync"

	"github.com/33cn/chain33/client"
	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
)

<<<<<<< HEAD
=======
//Query 检索
>>>>>>> upstream/master
type Query struct {
	db        dbm.DB
	stateHash []byte
	client    queue.Client
	mu        sync.Mutex
	api       client.QueueProtocolAPI
}

<<<<<<< HEAD
=======
//NewQuery new
>>>>>>> upstream/master
func NewQuery(db dbm.DB, qclient queue.Client, stateHash []byte) *Query {
	query := &Query{db: db, client: qclient, stateHash: stateHash}
	query.api, _ = client.New(qclient, nil)
	return query
}

<<<<<<< HEAD
=======
//Query 检索
>>>>>>> upstream/master
func (q *Query) Query(driver string, funcname string, param types.Message) (types.Message, error) {
	query := &types.ChainExecutor{
		Driver:    driver,
		FuncName:  funcname,
		Param:     types.Encode(param),
		StateHash: q.getStateHash(),
	}
	return q.api.QueryChain(query)
}

func (q *Query) updateStateHash(stateHash []byte) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.stateHash = stateHash
}

func (q *Query) getStateHash() (stateHash []byte) {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.stateHash
}
