// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
package executor

/*
manage 负责管理配置
 1. 添加管理
 1. 添加运营人员
 1. （未来）修改某些配置项
*/

=======
// Package executor 管理插件执行器
package executor

>>>>>>> upstream/master
import (
	log "github.com/33cn/chain33/common/log/log15"
	drivers "github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
)

var (
	clog       = log.New("module", "execs.manage")
	driverName = "manage"
	conf       = types.ConfSub(driverName)
)

func init() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&Manage{}))
}

<<<<<<< HEAD
=======
// Init resister a dirver
>>>>>>> upstream/master
func Init(name string, sub []byte) {
	drivers.Register(GetName(), newManage, types.GetDappFork(driverName, "Enable"))
}

<<<<<<< HEAD
=======
// GetName return manage name
>>>>>>> upstream/master
func GetName() string {
	return newManage().GetName()
}

<<<<<<< HEAD
=======
// Manage defines Manage object
>>>>>>> upstream/master
type Manage struct {
	drivers.DriverBase
}

func newManage() drivers.Driver {
	c := &Manage{}
	c.SetChild(c)
	c.SetExecutorType(types.LoadExecutorType(driverName))
	return c
}

<<<<<<< HEAD
=======
// GetDriverName return a drivername
>>>>>>> upstream/master
func (c *Manage) GetDriverName() string {
	return driverName
}

<<<<<<< HEAD
=======
// CheckTx checkout transaction
>>>>>>> upstream/master
func (c *Manage) CheckTx(tx *types.Transaction, index int) error {
	return nil
}

<<<<<<< HEAD
=======
// IsSuperManager is supper manager or not
>>>>>>> upstream/master
func IsSuperManager(addr string) bool {
	for _, m := range conf.GStrList("superManager") {
		if addr == m {
			return true
		}
	}
	return false
}
<<<<<<< HEAD
=======

// CheckReceiptExecOk return true to check if receipt ty is ok
func (c *Manage) CheckReceiptExecOk() bool {
	return true
}
>>>>>>> upstream/master
