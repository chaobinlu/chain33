// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
package store

//store package store the world - state data
=======
// Package store store the world - state data
package store

>>>>>>> upstream/master
import (
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/system/store"
	"github.com/33cn/chain33/types"
)

<<<<<<< HEAD
=======
// New new store queue module
>>>>>>> upstream/master
func New(cfg *types.Store, sub map[string][]byte) queue.Module {
	s, err := store.Load(cfg.Name)
	if err != nil {
		panic("Unsupported store type:" + cfg.Name + " " + err.Error())
	}
	subcfg, ok := sub[cfg.Name]
	if !ok {
		subcfg = nil
	}
	return s(cfg, subcfg)
}
