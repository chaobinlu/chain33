// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package executor coins执行器
>>>>>>> upstream/master
package executor

/*
coins 是一个货币的exec。内置货币的执行器。

主要提供两种操作：
EventTransfer -> 转移资产
*/

<<<<<<< HEAD
//package none execer for unknow execer
//all none transaction exec ok, execept nofee
//nofee transaction will not pack into block
=======
// package none execer for unknow execer
// all none transaction exec ok, execept nofee
// nofee transaction will not pack into block
>>>>>>> upstream/master

import (
	drivers "github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
)

<<<<<<< HEAD
//var clog = log.New("module", "execs.coins")
var driverName = "coins"

=======
// var clog = log.New("module", "execs.coins")
var driverName = "coins"

// Init defines a register function
>>>>>>> upstream/master
func Init(name string, sub []byte) {
	if name != driverName {
		panic("system dapp can't be rename")
	}
	drivers.Register(driverName, newCoins, types.GetDappFork(driverName, "Enable"))
}

<<<<<<< HEAD
//初始化过程比较重量级，有很多reflact, 所以弄成全局的
=======
// the initialization process is relatively heavyweight, lots of reflact, so it's global
>>>>>>> upstream/master
func init() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&Coins{}))
}

<<<<<<< HEAD
=======
// GetName return name string
>>>>>>> upstream/master
func GetName() string {
	return newCoins().GetName()
}

<<<<<<< HEAD
=======
// Coins defines coins
>>>>>>> upstream/master
type Coins struct {
	drivers.DriverBase
}

func newCoins() drivers.Driver {
	c := &Coins{}
	c.SetChild(c)
	c.SetExecutorType(types.LoadExecutorType(driverName))
	return c
}

<<<<<<< HEAD
=======
// GetDriverName get drive name
>>>>>>> upstream/master
func (c *Coins) GetDriverName() string {
	return driverName
}

<<<<<<< HEAD
func (c *Coins) CheckTx(tx *types.Transaction, index int) error {
	return nil
}

//coins 合约 运行 ticket 合约的挖矿交易
=======
// CheckTx check transaction amount 必须不能为负数
func (c *Coins) CheckTx(tx *types.Transaction, index int) error {
	ety := c.GetExecutorType()
	amount, err := ety.Amount(tx)
	if err != nil {
		return err
	}
	if amount < 0 {
		return types.ErrAmount
	}
	return nil
}

// IsFriend coins contract  the mining transaction that runs the ticket contract
>>>>>>> upstream/master
func (c *Coins) IsFriend(myexec, writekey []byte, othertx *types.Transaction) bool {
	//step1 先判定自己合约的权限
	if !c.AllowIsSame(myexec) {
		return false
	}
	//step2 判定 othertx 的 执行器名称(只允许主链，并且是挖矿的行为)
	if string(othertx.Execer) == "ticket" && othertx.ActionName() == "miner" {
		return true
	}
	return false
}
<<<<<<< HEAD
=======

// CheckReceiptExecOk return true to check if receipt ty is ok
func (c *Coins) CheckReceiptExecOk() bool {
	return true
}
>>>>>>> upstream/master
