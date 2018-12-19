// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package ed25519 ed25519系统加密包
>>>>>>> upstream/master
package ed25519

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/33cn/chain33/common/crypto"
	"github.com/33cn/chain33/common/ed25519"
)

<<<<<<< HEAD
type Driver struct{}

// Crypto
=======
//Driver 驱动
type Driver struct{}

//GenKey 生成私钥
>>>>>>> upstream/master
func (d Driver) GenKey() (crypto.PrivKey, error) {
	privKeyBytes := new([64]byte)
	copy(privKeyBytes[:32], crypto.CRandBytes(32))
	ed25519.MakePublicKey(privKeyBytes)
	return PrivKeyEd25519(*privKeyBytes), nil
}

<<<<<<< HEAD
=======
//PrivKeyFromBytes 字节转为私钥
>>>>>>> upstream/master
func (d Driver) PrivKeyFromBytes(b []byte) (privKey crypto.PrivKey, err error) {
	if len(b) != 64 {
		return nil, errors.New("invalid priv key byte")
	}
	privKeyBytes := new([64]byte)
	copy(privKeyBytes[:32], b[:32])
	ed25519.MakePublicKey(privKeyBytes)
	return PrivKeyEd25519(*privKeyBytes), nil
}

<<<<<<< HEAD
=======
//PubKeyFromBytes 字节转为公钥
>>>>>>> upstream/master
func (d Driver) PubKeyFromBytes(b []byte) (pubKey crypto.PubKey, err error) {
	if len(b) != 32 {
		return nil, errors.New("invalid pub key byte")
	}
	pubKeyBytes := new([32]byte)
	copy(pubKeyBytes[:], b[:])
	return PubKeyEd25519(*pubKeyBytes), nil
}

<<<<<<< HEAD
=======
//SignatureFromBytes 字节转为签名
>>>>>>> upstream/master
func (d Driver) SignatureFromBytes(b []byte) (sig crypto.Signature, err error) {
	sigBytes := new([64]byte)
	copy(sigBytes[:], b[:])
	return SignatureEd25519(*sigBytes), nil
}

<<<<<<< HEAD
// PrivKey
type PrivKeyEd25519 [64]byte

=======
//PrivKeyEd25519 PrivKey
type PrivKeyEd25519 [64]byte

//Bytes 字节格式
>>>>>>> upstream/master
func (privKey PrivKeyEd25519) Bytes() []byte {
	s := make([]byte, 64)
	copy(s, privKey[:])
	return s
}

<<<<<<< HEAD
=======
//Sign 签名
>>>>>>> upstream/master
func (privKey PrivKeyEd25519) Sign(msg []byte) crypto.Signature {
	privKeyBytes := [64]byte(privKey)
	signatureBytes := ed25519.Sign(&privKeyBytes, msg)
	return SignatureEd25519(*signatureBytes)
}

<<<<<<< HEAD
=======
//PubKey 公钥
>>>>>>> upstream/master
func (privKey PrivKeyEd25519) PubKey() crypto.PubKey {
	privKeyBytes := [64]byte(privKey)
	return PubKeyEd25519(*ed25519.MakePublicKey(&privKeyBytes))
}

<<<<<<< HEAD
func (privKey PrivKeyEd25519) Equals(other crypto.PrivKey) bool {
	if otherEd, ok := other.(PrivKeyEd25519); ok {
		return bytes.Equal(privKey[:], otherEd[:])
	} else {
		return false
	}
}

// PubKey
type PubKeyEd25519 [32]byte

=======
//Equals 相等
func (privKey PrivKeyEd25519) Equals(other crypto.PrivKey) bool {
	if otherEd, ok := other.(PrivKeyEd25519); ok {
		return bytes.Equal(privKey[:], otherEd[:])
	}
	return false

}

//PubKeyEd25519 PubKey
type PubKeyEd25519 [32]byte

//Bytes 字节格式
>>>>>>> upstream/master
func (pubKey PubKeyEd25519) Bytes() []byte {
	s := make([]byte, 32)
	copy(s, pubKey[:])
	return s
}

<<<<<<< HEAD
=======
//VerifyBytes 验证字节
>>>>>>> upstream/master
func (pubKey PubKeyEd25519) VerifyBytes(msg []byte, sig crypto.Signature) bool {
	// unwrap if needed
	if wrap, ok := sig.(SignatureS); ok {
		sig = wrap.Signature
	}
	// make sure we use the same algorithm to sign
	sigEd25519, ok := sig.(SignatureEd25519)
	if !ok {
		return false
	}
	pubKeyBytes := [32]byte(pubKey)
	sigBytes := [64]byte(sigEd25519)
	return ed25519.Verify(&pubKeyBytes, msg, &sigBytes)
}

<<<<<<< HEAD
=======
//KeyString 公钥字符串格式
>>>>>>> upstream/master
func (pubKey PubKeyEd25519) KeyString() string {
	return fmt.Sprintf("%X", pubKey[:])
}

<<<<<<< HEAD
func (pubKey PubKeyEd25519) Equals(other crypto.PubKey) bool {
	if otherEd, ok := other.(PubKeyEd25519); ok {
		return bytes.Equal(pubKey[:], otherEd[:])
	} else {
		return false
	}
}

// Signature
type SignatureEd25519 [64]byte

=======
//Equals 相等
func (pubKey PubKeyEd25519) Equals(other crypto.PubKey) bool {
	if otherEd, ok := other.(PubKeyEd25519); ok {
		return bytes.Equal(pubKey[:], otherEd[:])
	}
	return false

}

//SignatureEd25519 Signature
type SignatureEd25519 [64]byte

//SignatureS 签名
>>>>>>> upstream/master
type SignatureS struct {
	crypto.Signature
}

<<<<<<< HEAD
=======
//Bytes 字节格式
>>>>>>> upstream/master
func (sig SignatureEd25519) Bytes() []byte {
	s := make([]byte, 64)
	copy(s, sig[:])
	return s
}

<<<<<<< HEAD
=======
//IsZero 是否是0
>>>>>>> upstream/master
func (sig SignatureEd25519) IsZero() bool { return len(sig) == 0 }

func (sig SignatureEd25519) String() string {
	fingerprint := make([]byte, len(sig[:]))
	copy(fingerprint, sig[:])
	return fmt.Sprintf("/%X.../", fingerprint)
}

<<<<<<< HEAD
func (sig SignatureEd25519) Equals(other crypto.Signature) bool {
	if otherEd, ok := other.(SignatureEd25519); ok {
		return bytes.Equal(sig[:], otherEd[:])
	} else {
		return false
	}
}

const Name = "ed25519"
const ID = 2
=======
//Equals 相等
func (sig SignatureEd25519) Equals(other crypto.Signature) bool {
	if otherEd, ok := other.(SignatureEd25519); ok {
		return bytes.Equal(sig[:], otherEd[:])
	}
	return false

}

//const
const (
	Name = "ed25519"
	ID   = 2
)
>>>>>>> upstream/master

func init() {
	crypto.Register(Name, &Driver{})
	crypto.RegisterType(Name, ID)
}
