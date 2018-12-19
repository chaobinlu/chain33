// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package crypto 加解密、签名接口定义
>>>>>>> upstream/master
package crypto

import (
	"fmt"
	"sync"
)

<<<<<<< HEAD
=======
//PrivKey 私钥
>>>>>>> upstream/master
type PrivKey interface {
	Bytes() []byte
	Sign(msg []byte) Signature
	PubKey() PubKey
	Equals(PrivKey) bool
}

<<<<<<< HEAD
=======
//Signature 签名
>>>>>>> upstream/master
type Signature interface {
	Bytes() []byte
	IsZero() bool
	String() string
	Equals(Signature) bool
}

<<<<<<< HEAD
=======
//PubKey 公钥
>>>>>>> upstream/master
type PubKey interface {
	Bytes() []byte
	KeyString() string
	VerifyBytes(msg []byte, sig Signature) bool
	Equals(PubKey) bool
}

<<<<<<< HEAD
=======
//Crypto 加密
>>>>>>> upstream/master
type Crypto interface {
	GenKey() (PrivKey, error)
	SignatureFromBytes([]byte) (Signature, error)
	PrivKeyFromBytes([]byte) (PrivKey, error)
	PubKeyFromBytes([]byte) (PubKey, error)
}

var (
	drivers     = make(map[string]Crypto)
	driversType = make(map[string]int)
)

var driverMutex sync.Mutex

<<<<<<< HEAD
=======
//const
>>>>>>> upstream/master
const (
	SignNameED25519 = "ed25519"
)

<<<<<<< HEAD
=======
//Register 注册
>>>>>>> upstream/master
func Register(name string, driver Crypto) {
	driverMutex.Lock()
	defer driverMutex.Unlock()
	if driver == nil {
		panic("crypto: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("crypto: Register called twice for driver " + name)
	}
	drivers[name] = driver
}

<<<<<<< HEAD
=======
//RegisterType 注册类型
>>>>>>> upstream/master
func RegisterType(name string, ty int) {
	driverMutex.Lock()
	defer driverMutex.Unlock()
	if _, dup := driversType[name]; dup {
		panic("crypto: Register(ty) called twice for driver " + name)
	}
	driversType[name] = ty
}

<<<<<<< HEAD
=======
//GetName 获取name
>>>>>>> upstream/master
func GetName(ty int) string {
	for name, t := range driversType {
		if t == ty {
			return name
		}
	}
	return "unknown"
}

<<<<<<< HEAD
=======
//GetType 获取type
>>>>>>> upstream/master
func GetType(name string) int {
	if ty, ok := driversType[name]; ok {
		return ty
	}
	return 0
}

<<<<<<< HEAD
=======
//New new
>>>>>>> upstream/master
func New(name string) (c Crypto, err error) {
	driverMutex.Lock()
	defer driverMutex.Unlock()
	c, ok := drivers[name]
	if !ok {
		err = fmt.Errorf("unknown driver %q", name)
		return
	}

	return c, nil
}

<<<<<<< HEAD
=======
//CertSignature 签名
>>>>>>> upstream/master
type CertSignature struct {
	Signature []byte
	Cert      []byte
}
