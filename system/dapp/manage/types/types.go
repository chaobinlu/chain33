// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
package types

import (
	"encoding/json"
=======
// Package types 管理插件相关的定义
package types

import (
>>>>>>> upstream/master
	"reflect"

	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/types"
)

var (
<<<<<<< HEAD
=======
	// ManageX defines a global string
>>>>>>> upstream/master
	ManageX    = "manage"
	actionName = map[string]int32{
		"Modify": ManageActionModifyConfig,
	}
	logmap = map[int64]*types.LogInfo{
		// 这里reflect.TypeOf类型必须是proto.Message类型，且是交易的回持结构
		TyLogModifyConfig: {reflect.TypeOf(types.ReceiptConfig{}), "LogModifyConfig"},
	}
)

func init() {
	types.AllowUserExec = append(types.AllowUserExec, []byte(ManageX))
	types.RegistorExecutor(ManageX, NewType())

	types.RegisterDappFork(ManageX, "Enable", 120000)
	types.RegisterDappFork(ManageX, "ForkManageExec", 400000)
}

<<<<<<< HEAD
=======
// ManageType defines managetype
>>>>>>> upstream/master
type ManageType struct {
	types.ExecTypeBase
}

<<<<<<< HEAD
=======
// NewType new a managetype object
>>>>>>> upstream/master
func NewType() *ManageType {
	c := &ManageType{}
	c.SetChild(c)
	return c
}

<<<<<<< HEAD
func (at *ManageType) GetPayload() types.Message {
	return &ManageAction{}
}

=======
// GetPayload return manageaction
func (m *ManageType) GetPayload() types.Message {
	return &ManageAction{}
}

// ActionName return action a string name
>>>>>>> upstream/master
func (m ManageType) ActionName(tx *types.Transaction) string {
	return "config"
}

<<<<<<< HEAD
=======
// Amount amount
>>>>>>> upstream/master
func (m ManageType) Amount(tx *types.Transaction) (int64, error) {
	return 0, nil
}

<<<<<<< HEAD
// TODO 暂时不修改实现， 先完成结构的重构
func (m ManageType) CreateTx(action string, message json.RawMessage) (*types.Transaction, error) {
	var tx *types.Transaction
	return tx, nil
}

=======
// GetLogMap  get log for map
>>>>>>> upstream/master
func (m *ManageType) GetLogMap() map[int64]*types.LogInfo {
	return logmap
}

<<<<<<< HEAD
// GetRealToAddr 重载该函数主要原因是manage的协议在实现过程中，不同高度的To地址规范不一样
=======
// GetRealToAddr main reason for overloading this function is because of the manage protocol
// during implementation, to address specification varies fron height to height
>>>>>>> upstream/master
func (m ManageType) GetRealToAddr(tx *types.Transaction) string {
	if len(tx.To) == 0 {
		// 如果To地址为空，则认为是早期低于types.ForkV11ManageExec高度的交易，直接返回合约地址
		return address.ExecAddress(string(tx.Execer))
	}
	return tx.To
}

<<<<<<< HEAD
func (m ManageType) GetTypeMap() map[string]int32 {
	return actionName
}
=======
// GetTypeMap return typename of actionname
func (m ManageType) GetTypeMap() map[string]int32 {
	return actionName
}

// GetName reset name
func (m *ManageType) GetName() string {
	return ManageX
}
>>>>>>> upstream/master
