// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/33cn/chain33/types"
)

const (
<<<<<<< HEAD
	CoinsActionTransfer       = 1
	CoinsActionGenesis        = 2
	CoinsActionWithdraw       = 3
=======
	// CoinsActionTransfer defines const number
	CoinsActionTransfer = 1
	// CoinsActionGenesis  defines const coinsactiongenesis number
	CoinsActionGenesis = 2
	// CoinsActionWithdraw defines const number coinsactionwithdraw
	CoinsActionWithdraw = 3
	// CoinsActionTransferToExec defines const number coinsactiontransfertoExec
>>>>>>> upstream/master
	CoinsActionTransferToExec = 10
)

var (
<<<<<<< HEAD
	CoinsX      = "coins"
=======
	// CoinsX defines a global string
	CoinsX = "coins"
	// ExecerCoins execer coins
>>>>>>> upstream/master
	ExecerCoins = []byte(CoinsX)
	actionName  = map[string]int32{
		"Transfer":       CoinsActionTransfer,
		"TransferToExec": CoinsActionTransferToExec,
		"Withdraw":       CoinsActionWithdraw,
		"Genesis":        CoinsActionGenesis,
	}
	logmap = make(map[int64]*types.LogInfo)
)

func init() {
	types.AllowUserExec = append(types.AllowUserExec, ExecerCoins)
	types.RegistorExecutor("coins", NewType())

	types.RegisterDappFork(CoinsX, "Enable", 0)
}

<<<<<<< HEAD
=======
// CoinsType defines exec type
>>>>>>> upstream/master
type CoinsType struct {
	types.ExecTypeBase
}

<<<<<<< HEAD
=======
// NewType new coinstype
>>>>>>> upstream/master
func NewType() *CoinsType {
	c := &CoinsType{}
	c.SetChild(c)
	return c
}

<<<<<<< HEAD
func (coins *CoinsType) GetPayload() types.Message {
	return &CoinsAction{}
}

func (coins *CoinsType) GetName() string {
	return CoinsX
}

func (coins *CoinsType) GetLogMap() map[int64]*types.LogInfo {
	return logmap
}

=======
// GetPayload  return payload
func (c *CoinsType) GetPayload() types.Message {
	return &CoinsAction{}
}

// GetName  return coins string
func (c *CoinsType) GetName() string {
	return CoinsX
}

// GetLogMap return log for map
func (c *CoinsType) GetLogMap() map[int64]*types.LogInfo {
	return logmap
}

// GetTypeMap return actionname for map
>>>>>>> upstream/master
func (c *CoinsType) GetTypeMap() map[string]int32 {
	return actionName
}

<<<<<<< HEAD
=======
// RPC_Default_Process default process fo rpc
>>>>>>> upstream/master
func (c *CoinsType) RPC_Default_Process(action string, msg interface{}) (*types.Transaction, error) {
	var create *types.CreateTx
	if _, ok := msg.(*types.CreateTx); !ok {
		return nil, types.ErrInvalidParam
	}
	create = msg.(*types.CreateTx)
	if create.IsToken {
		return nil, types.ErrNotSupport
	}
	tx, err := c.AssertCreate(create)
	if err != nil {
		return nil, err
	}
	//to地址的问题,如果是主链交易，to地址就是直接是设置to
	if !types.IsPara() {
		tx.To = create.To
	}
	return tx, err
}

<<<<<<< HEAD
=======
// GetAssets return asset list
>>>>>>> upstream/master
func (c *CoinsType) GetAssets(tx *types.Transaction) ([]*types.Asset, error) {
	assetlist, err := c.ExecTypeBase.GetAssets(tx)
	if err != nil || len(assetlist) == 0 {
		return nil, err
	}
	if assetlist[0].Symbol == "" {
		assetlist[0].Symbol = types.BTY
	}
	return assetlist, nil
}
