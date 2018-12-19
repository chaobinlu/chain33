// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"testing"

<<<<<<< HEAD
	"github.com/stretchr/testify/assert"
	"github.com/33cn/chain33/types"
=======
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
>>>>>>> upstream/master
)

func TestTypeReflact(t *testing.T) {
	ty := NewType()
	assert.NotNil(t, ty)
	//创建一个json字符串
<<<<<<< HEAD
	data, err := types.PBToJson(&types.AssetsTransfer{Amount: 10})
=======
	data, err := types.PBToJSON(&types.AssetsTransfer{Amount: 10})
>>>>>>> upstream/master
	assert.Nil(t, err)
	raw := json.RawMessage(data)
	tx, err := ty.CreateTx("Transfer", raw)
	assert.Nil(t, err)
	name, val, err := ty.DecodePayloadValue(tx)
	assert.Nil(t, err)
	assert.Equal(t, "Transfer", name)
	assert.Equal(t, !types.IsNil(val) && val.CanInterface(), true)
	if !types.IsNil(val) && val.CanInterface() {
		assert.Equal(t, int64(10), val.Interface().(*types.AssetsTransfer).GetAmount())
	}
}
