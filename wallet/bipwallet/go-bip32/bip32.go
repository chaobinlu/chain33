// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package bip32 A fully compliant implementation of the BIP0032 spec for Hierarchical Deterministic Bitcoin addresses
>>>>>>> upstream/master
package bip32

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

const (
<<<<<<< HEAD
	FirstHardenedChild        = uint32(0x80000000)
=======
	// FirstHardenedChild FirstHardenedChild
	FirstHardenedChild = uint32(0x80000000)
	// PublicKeyCompressedLength 公钥压缩长度
>>>>>>> upstream/master
	PublicKeyCompressedLength = 33
)

var (
<<<<<<< HEAD
	PrivateWalletVersion, _ = hex.DecodeString("0488ADE4")
	PublicWalletVersion, _  = hex.DecodeString("0488B21E")
)

// Represents a bip32 extended key containing key data, chain code, parent information, and other meta data
=======
	// PrivateWalletVersion 私钥钱包版本
	PrivateWalletVersion, _ = hex.DecodeString("0488ADE4")
	// PublicWalletVersion 公钥钱包版本
	PublicWalletVersion, _ = hex.DecodeString("0488B21E")
)

// Key Represents a bip32 extended key containing key data, chain code, parent information, and other meta data
>>>>>>> upstream/master
type Key struct {
	Version     []byte // 4 bytes
	Depth       byte   // 1 bytes
	ChildNumber []byte // 4 bytes
	FingerPrint []byte // 4 bytes
	ChainCode   []byte // 32 bytes
	Key         []byte // 33 bytes
	IsPrivate   bool   // unserialized
}

<<<<<<< HEAD
// Creates a new master extended key from a seed
=======
// NewMasterKey Creates a new master extended key from a seed
>>>>>>> upstream/master
func NewMasterKey(seed []byte) (*Key, error) {
	// Generate key and chaincode
	hmac := hmac.New(sha512.New, []byte("Bitcoin seed"))
	hmac.Write(seed)
	intermediary := hmac.Sum(nil)

	// Split it into our key and chain code
	keyBytes := intermediary[:32]
	chainCode := intermediary[32:]

	// Validate key
	err := validatePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}

	// Create the key struct
	key := &Key{
		Version:     PrivateWalletVersion,
		ChainCode:   chainCode,
		Key:         keyBytes,
		Depth:       0x0,
		ChildNumber: []byte{0x00, 0x00, 0x00, 0x00},
		FingerPrint: []byte{0x00, 0x00, 0x00, 0x00},
		IsPrivate:   true,
	}

	return key, nil
}

<<<<<<< HEAD
// Derives a child key from a given parent as outlined by bip32
=======
// NewChildKey Derives a child key from a given parent as outlined by bip32
>>>>>>> upstream/master
func (key *Key) NewChildKey(childIdx uint32) (*Key, error) {
	hardenedChild := childIdx >= FirstHardenedChild
	childIndexBytes := uint32Bytes(childIdx)

	// Fail early if trying to create hardned child from public key
	if !key.IsPrivate && hardenedChild {
		return nil, errors.New("Can't create hardened child for public key")
	}

	// Get intermediary to create key and chaincode from
	// Hardened children are based on the private key
	// NonHardened children are based on the public key
	var data []byte
	if hardenedChild {
		data = append([]byte{0x0}, key.Key...)
	} else if key.IsPrivate {
		data = publicKeyForPrivateKey(key.Key)
	} else {
		data = key.Key
	}
	data = append(data, childIndexBytes...)

	hmac := hmac.New(sha512.New, key.ChainCode)
	hmac.Write(data)
	intermediary := hmac.Sum(nil)

	// Create child Key with data common to all both scenarios
	childKey := &Key{
		ChildNumber: childIndexBytes,
		ChainCode:   intermediary[32:],
		Depth:       key.Depth + 1,
		IsPrivate:   key.IsPrivate,
	}

	// Bip32 CKDpriv
	if key.IsPrivate {
		childKey.Version = PrivateWalletVersion
		childKey.FingerPrint = hash160(publicKeyForPrivateKey(key.Key))[:4]
		childKey.Key = addPrivateKeys(intermediary[:32], key.Key)

		// Validate key
		err := validatePrivateKey(childKey.Key)
		if err != nil {
			return nil, err
		}
		// Bip32 CKDpub
	} else {
		keyBytes := publicKeyForPrivateKey(intermediary[:32])
		// Validate key
		err := validateChildPublicKey(keyBytes)
		if err != nil {
			return nil, err
		}

		childKey.Version = PublicWalletVersion
		childKey.FingerPrint = hash160(key.Key)[:4]
		childKey.Key = addPublicKeys(keyBytes, key.Key)
	}

	return childKey, nil
}

<<<<<<< HEAD
// Create public version of key or return a copy; 'Neuter' function from the bip32 spec
=======
// PublicKey Create public version of key or return a copy; 'Neuter' function from the bip32 spec
>>>>>>> upstream/master
func (key *Key) PublicKey() *Key {
	keyBytes := key.Key

	if key.IsPrivate {
		keyBytes = publicKeyForPrivateKey(keyBytes)
	}

	return &Key{
		Version:     PublicWalletVersion,
		Key:         keyBytes,
		Depth:       key.Depth,
		ChildNumber: key.ChildNumber,
		FingerPrint: key.FingerPrint,
		ChainCode:   key.ChainCode,
		IsPrivate:   false,
	}
}

<<<<<<< HEAD
// Serialized an Key to a 78 byte byte slice
=======
// Serialize Serialized an Key to a 78 byte byte slice
>>>>>>> upstream/master
func (key *Key) Serialize() []byte {
	// Private keys should be prepended with a single null byte
	keyBytes := key.Key
	if key.IsPrivate {
		keyBytes = append([]byte{0x0}, keyBytes...)
	}

	// Write fields to buffer in order
	buffer := new(bytes.Buffer)
	buffer.Write(key.Version)
	buffer.WriteByte(key.Depth)
	buffer.Write(key.FingerPrint)
	buffer.Write(key.ChildNumber)
	buffer.Write(key.ChainCode)
	buffer.Write(keyBytes)

	// Append the standard doublesha256 checksum
	serializedKey := addChecksumToBytes(buffer.Bytes())

	return serializedKey
}

<<<<<<< HEAD
// Encode the Key in the standard Bitcoin base58 encoding
=======
// String Encode the Key in the standard Bitcoin base58 encoding
>>>>>>> upstream/master
func (key *Key) String() string {
	return string(base58Encode(key.Serialize()))
}

<<<<<<< HEAD
// Cryptographically secure seed
=======
// NewSeed Cryptographically secure seed
>>>>>>> upstream/master
func NewSeed() ([]byte, error) {
	// Well that easy, just make go read 256 random bytes into a slice
	s := make([]byte, 256)
	_, err := rand.Read(s)
	return s, err
}
