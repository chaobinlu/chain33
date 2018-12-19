// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package sm2 系统签名包
>>>>>>> upstream/master
package sm2

import (
	"bytes"
	"errors"
	"fmt"

	"crypto/elliptic"
	"encoding/asn1"
	"math/big"

<<<<<<< HEAD
	"github.com/tjfoc/gmsm/sm2"
	"github.com/33cn/chain33/common/crypto"
)

const (
	SM2_RPIVATEKEY_LENGTH = 32
	SM2_PUBLICKEY_LENGTH  = 65
)

type Driver struct{}

func (d Driver) GenKey() (crypto.PrivKey, error) {
	privKeyBytes := [SM2_RPIVATEKEY_LENGTH]byte{}
	copy(privKeyBytes[:], crypto.CRandBytes(SM2_RPIVATEKEY_LENGTH))
=======
	"github.com/33cn/chain33/common/crypto"
	"github.com/tjfoc/gmsm/sm2"
)

//const
const (
	SM2PrivateKeyLength = 32
	SM2PublicKeyLength  = 65
)

//Driver 驱动
type Driver struct{}

//GenKey 生成私钥
func (d Driver) GenKey() (crypto.PrivKey, error) {
	privKeyBytes := [SM2PrivateKeyLength]byte{}
	copy(privKeyBytes[:], crypto.CRandBytes(SM2PrivateKeyLength))
>>>>>>> upstream/master
	priv, _ := privKeyFromBytes(sm2.P256Sm2(), privKeyBytes[:])
	copy(privKeyBytes[:], SerializePrivateKey(priv))
	return PrivKeySM2(privKeyBytes), nil
}

<<<<<<< HEAD
func (d Driver) PrivKeyFromBytes(b []byte) (privKey crypto.PrivKey, err error) {
	if len(b) != SM2_RPIVATEKEY_LENGTH {
		return nil, errors.New("invalid priv key byte")
	}
	privKeyBytes := new([SM2_RPIVATEKEY_LENGTH]byte)
	copy(privKeyBytes[:], b[:SM2_RPIVATEKEY_LENGTH])
=======
//PrivKeyFromBytes 字节转为私钥
func (d Driver) PrivKeyFromBytes(b []byte) (privKey crypto.PrivKey, err error) {
	if len(b) != SM2PrivateKeyLength {
		return nil, errors.New("invalid priv key byte")
	}
	privKeyBytes := new([SM2PrivateKeyLength]byte)
	copy(privKeyBytes[:], b[:SM2PrivateKeyLength])
>>>>>>> upstream/master

	priv, _ := privKeyFromBytes(sm2.P256Sm2(), privKeyBytes[:])

	copy(privKeyBytes[:], SerializePrivateKey(priv))
	return PrivKeySM2(*privKeyBytes), nil
}

<<<<<<< HEAD
func (d Driver) PubKeyFromBytes(b []byte) (pubKey crypto.PubKey, err error) {
	if len(b) != SM2_PUBLICKEY_LENGTH {
		return nil, errors.New("invalid pub key byte")
	}
	pubKeyBytes := new([SM2_PUBLICKEY_LENGTH]byte)
=======
//PubKeyFromBytes 字节转为公钥
func (d Driver) PubKeyFromBytes(b []byte) (pubKey crypto.PubKey, err error) {
	if len(b) != SM2PublicKeyLength {
		return nil, errors.New("invalid pub key byte")
	}
	pubKeyBytes := new([SM2PublicKeyLength]byte)
>>>>>>> upstream/master
	copy(pubKeyBytes[:], b[:])
	return PubKeySM2(*pubKeyBytes), nil
}

<<<<<<< HEAD
=======
//SignatureFromBytes 字节转为签名
>>>>>>> upstream/master
func (d Driver) SignatureFromBytes(b []byte) (sig crypto.Signature, err error) {
	var certSignature crypto.CertSignature
	_, err = asn1.Unmarshal(b, &certSignature)
	if err != nil {
		return SignatureSM2(b), nil
	}

	if len(certSignature.Cert) == 0 {
		return SignatureSM2(b), nil
	}

	return SignatureSM2(certSignature.Signature), nil
}

<<<<<<< HEAD
type PrivKeySM2 [SM2_RPIVATEKEY_LENGTH]byte

func (privKey PrivKeySM2) Bytes() []byte {
	s := make([]byte, SM2_RPIVATEKEY_LENGTH)
=======
//PrivKeySM2 私钥
type PrivKeySM2 [SM2PrivateKeyLength]byte

//Bytes 字节格式
func (privKey PrivKeySM2) Bytes() []byte {
	s := make([]byte, SM2PrivateKeyLength)
>>>>>>> upstream/master
	copy(s, privKey[:])
	return s
}

<<<<<<< HEAD
=======
//Sign 签名
>>>>>>> upstream/master
func (privKey PrivKeySM2) Sign(msg []byte) crypto.Signature {
	priv, _ := privKeyFromBytes(sm2.P256Sm2(), privKey[:])
	r, s, err := sm2.Sign(priv, crypto.Sm3Hash(msg))
	if err != nil {
		return nil
	}

	//sm2不需要LowS转换
	//s = ToLowS(pub, s)
	return SignatureSM2(Serialize(r, s))
}

<<<<<<< HEAD
=======
//PubKey 私钥生成公钥
>>>>>>> upstream/master
func (privKey PrivKeySM2) PubKey() crypto.PubKey {
	_, pub := privKeyFromBytes(sm2.P256Sm2(), privKey[:])
	var pubSM2 PubKeySM2
	copy(pubSM2[:], SerializePublicKey(pub))
	return pubSM2
}

<<<<<<< HEAD
=======
//Equals 公钥
>>>>>>> upstream/master
func (privKey PrivKeySM2) Equals(other crypto.PrivKey) bool {
	if otherSecp, ok := other.(PrivKeySM2); ok {
		return bytes.Equal(privKey[:], otherSecp[:])
	}

	return false
}

func (privKey PrivKeySM2) String() string {
	return fmt.Sprintf("PrivKeySM2{*****}")
}

<<<<<<< HEAD
type PubKeySM2 [SM2_PUBLICKEY_LENGTH]byte

func (pubKey PubKeySM2) Bytes() []byte {
	s := make([]byte, SM2_PUBLICKEY_LENGTH)
=======
//PubKeySM2 公钥
type PubKeySM2 [SM2PublicKeyLength]byte

//Bytes 字节格式
func (pubKey PubKeySM2) Bytes() []byte {
	s := make([]byte, SM2PublicKeyLength)
>>>>>>> upstream/master
	copy(s, pubKey[:])
	return s
}

<<<<<<< HEAD
=======
//VerifyBytes 验证字节
>>>>>>> upstream/master
func (pubKey PubKeySM2) VerifyBytes(msg []byte, sig crypto.Signature) bool {
	if wrap, ok := sig.(SignatureS); ok {
		sig = wrap.Signature
	}

	sigSM2, ok := sig.(SignatureSM2)
	if !ok {
		fmt.Printf("convert failed\n")
		return false
	}

	pub, err := parsePubKey(pubKey[:], sm2.P256Sm2())
	if err != nil {
		fmt.Printf("parse pubkey failed\n")
		return false
	}

	r, s, err := Deserialize(sigSM2)
	if err != nil {
		fmt.Printf("unmarshal sign failed")
		return false
	}

	//国密签名算法和ecdsa不一样，-s验签不通过，所以不需要LowS检查
	//fmt.Printf("verify:%x, r:%d, s:%d\n", crypto.Sm3Hash(msg), r, s)
	//lowS := IsLowS(s)
	//if !lowS {
	//	fmt.Printf("lowS check failed")
	//	return false
	//}

	return sm2.Verify(pub, crypto.Sm3Hash(msg), r, s)
}

func (pubKey PubKeySM2) String() string {
	return fmt.Sprintf("PubKeySM2{%X}", pubKey[:])
}

<<<<<<< HEAD
// Must return the full bytes in hex.
=======
//KeyString Must return the full bytes in hex.
>>>>>>> upstream/master
// Used for map keying, etc.
func (pubKey PubKeySM2) KeyString() string {
	return fmt.Sprintf("%X", pubKey[:])
}

<<<<<<< HEAD
func (pubKey PubKeySM2) Equals(other crypto.PubKey) bool {
	if otherSecp, ok := other.(PubKeySM2); ok {
		return bytes.Equal(pubKey[:], otherSecp[:])
	} else {
		return false
	}
}

type SignatureSM2 []byte

=======
//Equals 相等
func (pubKey PubKeySM2) Equals(other crypto.PubKey) bool {
	if otherSecp, ok := other.(PubKeySM2); ok {
		return bytes.Equal(pubKey[:], otherSecp[:])
	}
	return false

}

//SignatureSM2 签名
type SignatureSM2 []byte

//SignatureS 签名
>>>>>>> upstream/master
type SignatureS struct {
	crypto.Signature
}

<<<<<<< HEAD
=======
//Bytes 字节格式
>>>>>>> upstream/master
func (sig SignatureSM2) Bytes() []byte {
	s := make([]byte, len(sig))
	copy(s, sig[:])
	return s
}

<<<<<<< HEAD
=======
//IsZero 是否为0
>>>>>>> upstream/master
func (sig SignatureSM2) IsZero() bool { return len(sig) == 0 }

func (sig SignatureSM2) String() string {
	fingerprint := make([]byte, len(sig[:]))
	copy(fingerprint, sig[:])
	return fmt.Sprintf("/%X.../", fingerprint)

}

<<<<<<< HEAD
=======
//Equals 相等
>>>>>>> upstream/master
func (sig SignatureSM2) Equals(other crypto.Signature) bool {
	if otherEd, ok := other.(SignatureSM2); ok {
		return bytes.Equal(sig[:], otherEd[:])
	}
	return false
}

<<<<<<< HEAD
const Name = "sm2"
const ID = 3
=======
//const
const (
	Name = "sm2"
	ID   = 3
)
>>>>>>> upstream/master

func init() {
	crypto.Register(Name, &Driver{})
	crypto.RegisterType(Name, ID)
}

func privKeyFromBytes(curve elliptic.Curve, pk []byte) (*sm2.PrivateKey, *sm2.PublicKey) {
	x, y := curve.ScalarBaseMult(pk)

	priv := &sm2.PrivateKey{
		PublicKey: sm2.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(pk),
	}

	return priv, &priv.PublicKey
}
