// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dapp

import (
	"fmt"

	"github.com/33cn/chain33/types"
)

<<<<<<< HEAD
=======
// HeightIndexStr height and index format string
>>>>>>> upstream/master
func HeightIndexStr(height, index int64) string {
	v := height*types.MaxTxsPerBlock + index
	return fmt.Sprintf("%018d", v)
}
