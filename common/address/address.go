// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package address 计算地址相关的函数
>>>>>>> upstream/master
package address

import (
	"bytes"
	"encoding/hex"
	"errors"

<<<<<<< HEAD
	"github.com/decred/base58"
	lru "github.com/hashicorp/golang-lru"
	. "github.com/33cn/chain33/common"
=======
	"github.com/33cn/chain33/common"
	"github.com/decred/base58"
	lru "github.com/hashicorp/golang-lru"
>>>>>>> upstream/master
)

var addrSeed = []byte("address seed bytes for public key")
var addressCache *lru.Cache
var checkAddressCache *lru.Cache
<<<<<<< HEAD

const MaxExecNameLength = 100

func init() {
	addressCache, _ = lru.New(10240)
	checkAddressCache, _ = lru.New(10240)
}

=======
var multisignCache *lru.Cache
var multiCheckAddressCache *lru.Cache

// ErrCheckVersion :
var ErrCheckVersion = errors.New("check version error")

//ErrCheckChecksum :
var ErrCheckChecksum = errors.New("Address Checksum error")

//MaxExecNameLength 执行器名最大长度
const MaxExecNameLength = 100

//NormalVer 普通地址的版本号
const NormalVer byte = 0

//MultiSignVer 多重签名地址的版本号
const MultiSignVer byte = 5

func init() {
	multisignCache, _ = lru.New(10240)
	addressCache, _ = lru.New(10240)
	checkAddressCache, _ = lru.New(10240)
	multiCheckAddressCache, _ = lru.New(10240)
}

//ExecPubKey 计算公钥
>>>>>>> upstream/master
func ExecPubKey(name string) []byte {
	if len(name) > MaxExecNameLength {
		panic("name too long")
	}
	var bname [200]byte
	buf := append(bname[:0], addrSeed...)
	buf = append(buf, []byte(name)...)
<<<<<<< HEAD
	hash := Sha2Sum(buf)
	return hash[:]
}

//计算量有点大，做一次cache
=======
	hash := common.Sha2Sum(buf)
	return hash[:]
}

//ExecAddress 计算量有点大，做一次cache
>>>>>>> upstream/master
func ExecAddress(name string) string {
	if value, ok := addressCache.Get(name); ok {
		return value.(string)
	}
<<<<<<< HEAD
	addr := PubKeyToAddress(ExecPubkey(name))
=======
	addr := GetExecAddress(name)
>>>>>>> upstream/master
	addrstr := addr.String()
	addressCache.Add(name, addrstr)
	return addrstr
}

<<<<<<< HEAD
=======
//MultiSignAddress create a multi sign address
func MultiSignAddress(pubkey []byte) string {
	if value, ok := multisignCache.Get(string(pubkey)); ok {
		return value.(string)
	}
	addr := HashToAddress(MultiSignVer, pubkey)
	addrstr := addr.String()
	multisignCache.Add(string(pubkey), addrstr)
	return addrstr
}

//ExecPubkey 计算公钥
>>>>>>> upstream/master
func ExecPubkey(name string) []byte {
	if len(name) > MaxExecNameLength {
		panic("name too long")
	}
	var bname [200]byte
	buf := append(bname[:0], addrSeed...)
	buf = append(buf, []byte(name)...)
<<<<<<< HEAD
	hash := Sha2Sum(buf)
	return hash[:]
}

func GetExecAddress(name string) *Address {
	if len(name) > MaxExecNameLength {
		panic("name too long")
	}
	var bname [200]byte
	buf := append(bname[:0], addrSeed...)
	buf = append(buf, []byte(name)...)
	hash := Sha2Sum(buf)
=======
	hash := common.Sha2Sum(buf)
	return hash[:]
}

//GetExecAddress 获取地址
func GetExecAddress(name string) *Address {
	hash := ExecPubkey(name)
>>>>>>> upstream/master
	addr := PubKeyToAddress(hash[:])
	return addr
}

<<<<<<< HEAD
func PubKeyToAddress(in []byte) *Address {
	a := new(Address)
	a.Pubkey = make([]byte, len(in))
	copy(a.Pubkey[:], in[:])
	a.Version = 0
	a.Hash160 = Rimp160AfterSha256(in)
	return a
}

func CheckAddress(addr string) (e error) {
	if value, ok := checkAddressCache.Get(addr); ok {
		if value == nil {
			return nil
		}
		return value.(error)
	}
=======
//PubKeyToAddress 公钥转为地址
func PubKeyToAddress(in []byte) *Address {
	return HashToAddress(NormalVer, in)
}

//HashToAddress hash32 to address
func HashToAddress(version byte, in []byte) *Address {
	a := new(Address)
	a.Pubkey = make([]byte, len(in))
	copy(a.Pubkey[:], in[:])
	a.Version = version
	a.Hash160 = common.Rimp160AfterSha256(in)
	return a
}

func checkAddress(ver byte, addr string) (e error) {
>>>>>>> upstream/master
	dec := base58.Decode(addr)
	if dec == nil {
		e = errors.New("Cannot decode b58 string '" + addr + "'")
		checkAddressCache.Add(addr, e)
		return
	}
	if len(dec) < 25 {
		e = errors.New("Address too short " + hex.EncodeToString(dec))
		checkAddressCache.Add(addr, e)
		return
	}
	if len(dec) == 25 {
<<<<<<< HEAD
		sh := Sha2Sum(dec[0:21])
		if !bytes.Equal(sh[:4], dec[21:25]) {
			e = errors.New("Address Checksum error")
		}
	}
=======
		sh := common.Sha2Sum(dec[0:21])
		if !bytes.Equal(sh[:4], dec[21:25]) {
			e = ErrCheckChecksum
		}
	}
	if dec[0] != ver {
		e = ErrCheckVersion
	}
	return e
}

//CheckMultiSignAddress 检查多重签名地址的有效性
func CheckMultiSignAddress(addr string) (e error) {
	if value, ok := multiCheckAddressCache.Get(addr); ok {
		if value == nil {
			return nil
		}
		return value.(error)
	}
	e = checkAddress(MultiSignVer, addr)
	multiCheckAddressCache.Add(addr, e)
	return
}

//CheckAddress 检查地址
func CheckAddress(addr string) (e error) {
	if value, ok := checkAddressCache.Get(addr); ok {
		if value == nil {
			return nil
		}
		return value.(error)
	}
	e = checkAddress(NormalVer, addr)
>>>>>>> upstream/master
	checkAddressCache.Add(addr, e)
	return
}

<<<<<<< HEAD
=======
//NewAddrFromString new 地址
>>>>>>> upstream/master
func NewAddrFromString(hs string) (a *Address, e error) {
	dec := base58.Decode(hs)
	if dec == nil {
		e = errors.New("Cannot decode b58 string '" + hs + "'")
		return
	}
	if len(dec) < 25 {
		e = errors.New("Address too short " + hex.EncodeToString(dec))
		return
	}
	if len(dec) == 25 {
<<<<<<< HEAD
		sh := Sha2Sum(dec[0:21])
		if !bytes.Equal(sh[:4], dec[21:25]) {
			e = errors.New("Address Checksum error")
=======
		sh := common.Sha2Sum(dec[0:21])
		if !bytes.Equal(sh[:4], dec[21:25]) {
			e = ErrCheckChecksum
>>>>>>> upstream/master
		} else {
			a = new(Address)
			a.Version = dec[0]
			copy(a.Hash160[:], dec[1:21])
			a.Checksum = make([]byte, 4)
			copy(a.Checksum, dec[21:25])
			a.Enc58str = hs
		}
	}
	return
}

<<<<<<< HEAD
=======
//Address 地址
>>>>>>> upstream/master
type Address struct {
	Version  byte
	Hash160  [20]byte // For a stealth address: it's HASH160
	Checksum []byte   // Unused for a stealth address
	Pubkey   []byte   // Unused for a stealth address
	Enc58str string
}

func (a *Address) String() string {
	if a.Enc58str == "" {
		var ad [25]byte
		ad[0] = a.Version
		copy(ad[1:21], a.Hash160[:])
		if a.Checksum == nil {
<<<<<<< HEAD
			sh := Sha2Sum(ad[0:21])
=======
			sh := common.Sha2Sum(ad[0:21])
>>>>>>> upstream/master
			a.Checksum = make([]byte, 4)
			copy(a.Checksum, sh[:4])
		}
		copy(ad[21:25], a.Checksum[:])
		a.Enc58str = base58.Encode(ad[:])
	}
	return a.Enc58str
}
