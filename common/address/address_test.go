// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package address

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
	"time"

<<<<<<< HEAD
	"github.com/stretchr/testify/require"
	"github.com/33cn/chain33/common/crypto"
=======
	"github.com/33cn/chain33/common/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
>>>>>>> upstream/master

	_ "github.com/33cn/chain33/system/crypto/init"
)

<<<<<<< HEAD
func TestAddress(t *testing.T) {
	c, err := crypto.New("secp256k1")
	if err != nil {
		t.Error(err)
		return
	}
	key, err := c.GenKey()
	if err != nil {
		t.Error(err)
		return
	}
=======
func genkey() crypto.PrivKey {
	c, err := crypto.New("secp256k1")
	if err != nil {
		panic(err)
	}
	key, err := c.GenKey()
	if err != nil {
		panic(err)
	}
	return key
}
func TestAddress(t *testing.T) {
	key := genkey()
>>>>>>> upstream/master
	t.Logf("%X", key.Bytes())
	addr := PubKeyToAddress(key.PubKey().Bytes())
	t.Log(addr)
}

<<<<<<< HEAD
=======
func TestMultiSignAddress(t *testing.T) {
	key := genkey()
	addr1 := MultiSignAddress(key.PubKey().Bytes())
	addr := MultiSignAddress(key.PubKey().Bytes())
	assert.Equal(t, addr1, addr)
	err := CheckAddress(addr)
	assert.Equal(t, ErrCheckVersion, err)
	err = CheckMultiSignAddress(addr)
	assert.Nil(t, err)
	t.Log(addr)
}

>>>>>>> upstream/master
func TestPubkeyToAddress(t *testing.T) {
	pubkey := "024a17b0c6eb3143839482faa7e917c9b90a8cfe5008dff748789b8cea1a3d08d5"
	b, err := hex.DecodeString(pubkey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%X", b)
	addr := PubKeyToAddress(b)
	t.Log(addr)
}

func TestCheckAddress(t *testing.T) {
	c, err := crypto.New("secp256k1")
	if err != nil {
		t.Error(err)
		return
	}
	key, err := c.GenKey()
	if err != nil {
		t.Error(err)
		return
	}
	addr := PubKeyToAddress(key.PubKey().Bytes())
	err = CheckAddress(addr.String())
	require.NoError(t, err)
}

<<<<<<< HEAD
=======
func TestExecAddress(t *testing.T) {
	assert.Equal(t, "16htvcBNSEA7fZhAdLJphDwQRQJaHpyHTp", ExecAddress("ticket"))
	assert.Equal(t, "16htvcBNSEA7fZhAdLJphDwQRQJaHpyHTp", ExecAddress("ticket"))
	addr, err := NewAddrFromString(ExecAddress("ticket"))
	assert.Nil(t, err)
	assert.Equal(t, addr.Version, NormalVer)
}

>>>>>>> upstream/master
func BenchmarkExecAddress(b *testing.B) {
	start := time.Now().UnixNano() / 1000000
	fmt.Println(start)
	for i := 0; i < b.N; i++ {
		ExecAddress("ticket")
	}
	end := time.Now().UnixNano() / 1000000
	fmt.Println(end)
	duration := end - start
	fmt.Println("duration with cache:", strconv.FormatInt(duration, 10))

	start = time.Now().UnixNano() / 1000000
	fmt.Println(start)
	for i := 0; i < b.N; i++ {
		GetExecAddress("ticket")
	}
	end = time.Now().UnixNano() / 1000000
	fmt.Println(end)
	duration = end - start
	fmt.Println("duration without cache:", strconv.FormatInt(duration, 10))
}
