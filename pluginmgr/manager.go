// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pluginmgr

import (
	"sync"

<<<<<<< HEAD
	"github.com/spf13/cobra"
	"github.com/33cn/chain33/rpc/types"
	wcom "github.com/33cn/chain33/wallet/common"
=======
	"github.com/33cn/chain33/rpc/types"
	wcom "github.com/33cn/chain33/wallet/common"
	"github.com/spf13/cobra"
>>>>>>> upstream/master
)

var pluginItems = make(map[string]Plugin)

var once = &sync.Once{}

<<<<<<< HEAD
=======
// InitExec init exec
>>>>>>> upstream/master
func InitExec(sub map[string][]byte) {
	once.Do(func() {
		for _, item := range pluginItems {
			item.InitExec(sub)
		}
	})
}

<<<<<<< HEAD
=======
// InitWallet init wallet plugin
>>>>>>> upstream/master
func InitWallet(wallet wcom.WalletOperate, sub map[string][]byte) {
	once.Do(func() {
		for _, item := range pluginItems {
			item.InitWallet(wallet, sub)
		}
	})
}

<<<<<<< HEAD
=======
// HasExec check is have the name exec
>>>>>>> upstream/master
func HasExec(name string) bool {
	for _, item := range pluginItems {
		if item.GetExecutorName() == name {
			return true
		}
	}
	return false
}

<<<<<<< HEAD
=======
// Register Register plugin
>>>>>>> upstream/master
func Register(p Plugin) {
	if p == nil {
		panic("plugin param is nil" + p.GetName())
	}
	packageName := p.GetName()
	if len(packageName) == 0 {
		panic("plugin package name is empty")
	}
	if _, ok := pluginItems[packageName]; ok {
		panic("execute plugin item is existed. name = " + packageName)
	}
	pluginItems[packageName] = p
}

<<<<<<< HEAD
=======
// AddCmd add Command for plugin
>>>>>>> upstream/master
func AddCmd(rootCmd *cobra.Command) {
	for _, item := range pluginItems {
		item.AddCmd(rootCmd)
	}
}

<<<<<<< HEAD
=======
// AddRPC add Rpc
>>>>>>> upstream/master
func AddRPC(s types.RPCServer) {
	for _, item := range pluginItems {
		item.AddRPC(s)
	}
}
