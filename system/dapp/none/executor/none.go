// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
package executor

//package none execer for unknow execer
//all none transaction exec ok, execept nofee
//nofee transaction will not pack into block
=======
// Package executor none执行器
package executor

// package none execer for unknow execer
// all none transaction exec ok, execept nofee
// nofee transaction will not pack into block
>>>>>>> upstream/master

import (
	drivers "github.com/33cn/chain33/system/dapp"
)

var driverName = "none"

<<<<<<< HEAD
=======
// Init register newnone
>>>>>>> upstream/master
func Init(name string, sub []byte) {
	if name != driverName {
		panic("system dapp can't be rename")
	}
	driverName = name
	drivers.Register(name, newNone, 0)
}

<<<<<<< HEAD
//执行时候的名称
=======
// GetName return name at execution time
>>>>>>> upstream/master
func GetName() string {
	return newNone().GetName()
}

<<<<<<< HEAD
=======
// None defines a none type
>>>>>>> upstream/master
type None struct {
	drivers.DriverBase
}

func newNone() drivers.Driver {
	n := &None{}
	n.SetChild(n)
	return n
}

<<<<<<< HEAD
//驱动注册时候的名称
=======
// GetDriverName return dcrivername at register
>>>>>>> upstream/master
func (n *None) GetDriverName() string {
	return driverName
}
