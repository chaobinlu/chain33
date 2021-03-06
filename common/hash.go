// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"

	"github.com/33cn/chain33/common/crypto/sha3"
	"golang.org/x/crypto/ripemd160"
)

const (
	hashLength = 32
)

<<<<<<< HEAD
=======
//Hash hash
>>>>>>> upstream/master
type Hash [hashLength]byte

//BytesToHash []byte -> hash
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

//StringToHash string -> hash
func StringToHash(s string) Hash { return BytesToHash([]byte(s)) }
<<<<<<< HEAD
func BigToHash(b *big.Int) Hash  { return BytesToHash(b.Bytes()) }
func HexToHash(s string) Hash {
	b, _ := FromHex(s)
	return BytesToHash(b)
}
=======
>>>>>>> upstream/master

//BigToHash *big.Int -> hash
func BigToHash(b *big.Int) Hash { return BytesToHash(b.Bytes()) }

//HexToHash hex -> hash
func HexToHash(s string) Hash {
	b, _ := FromHex(s)
	return BytesToHash(b)
}

//Str Get the string representation of the underlying hash
func (h Hash) Str() string { return string(h[:]) }

//Bytes Get the []byte representation of the underlying hash
func (h Hash) Bytes() []byte { return h[:] }

//Hex Get the hex representation of the underlying hash
func (h Hash) Hex() string { return hexEncode(h[:]) }

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (h Hash) TerminalString() string {
	return fmt.Sprintf("%x…%x", h[:3], h[29:])
}

// String implements the stringer interface and is used also by the logger when
// doing full logging into a file.
func (h Hash) String() string {
	return h.Hex()
}

// Format implements fmt.Formatter, forcing the byte slice to be formatted as is,
// without going through the stringer interface used for logging.
func (h Hash) Format(s fmt.State, c rune) {
	fmt.Fprintf(s, "%"+string(c), h[:])
}

//SetBytes Sets the hash to the value of b. If b is larger than len(h), 'b' will be cropped (from the left).
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-hashLength:]
	}

	copy(h[hashLength-len(b):], b)
}

//SetString Set string `s` to h. If s is larger than len(h) s will be cropped (from left) to fit.
func (h *Hash) SetString(s string) { h.SetBytes([]byte(s)) }

//Set Sets h to other
func (h *Hash) Set(other Hash) {
	for i, v := range other {
		h[i] = v
	}
}

//EmptyHash hash是否为空
func EmptyHash(h Hash) bool {
	return h == Hash{}
}

func hexEncode(b []byte) string {
	enc := make([]byte, len(b)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], b)
	return string(enc)
}

<<<<<<< HEAD
=======
//ToHex []byte -> hex
>>>>>>> upstream/master
func ToHex(b []byte) string {
	hex := Bytes2Hex(b)
	// Prefer output of "0x0" instead of "0x"
	if len(hex) == 0 {
		return ""
	}
	return "0x" + hex
}

<<<<<<< HEAD
=======
//HashHex []byte -> hex
>>>>>>> upstream/master
func HashHex(d []byte) string {
	var buf [64]byte
	hex.Encode(buf[:], d)
	return string(buf[:])
}

<<<<<<< HEAD
=======
//FromHex hex -> []byte
>>>>>>> upstream/master
func FromHex(s string) ([]byte, error) {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
		if len(s)%2 == 1 {
			s = "0" + s
		}
		return Hex2Bytes(s)
	}
	return []byte{}, nil
}

<<<<<<< HEAD
// Copy bytes
//
// Returns an exact copy of the provided bytes
=======
// CopyBytes Returns an exact copy of the provided bytes
>>>>>>> upstream/master
func CopyBytes(b []byte) (copiedBytes []byte) {
	if b == nil {
		return nil
	}
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)

	return
}

<<<<<<< HEAD
=======
//HasHexPrefix 是否包含0x前缀
>>>>>>> upstream/master
func HasHexPrefix(str string) bool {
	l := len(str)
	return l >= 2 && str[0:2] == "0x"
}

<<<<<<< HEAD
=======
//IsHex 是否是hex字符串
>>>>>>> upstream/master
func IsHex(str string) bool {
	l := len(str)
	return l >= 4 && l%2 == 0 && str[0:2] == "0x"
}

<<<<<<< HEAD
=======
//Bytes2Hex []byte -> hex
>>>>>>> upstream/master
func Bytes2Hex(d []byte) string {
	return hex.EncodeToString(d)
}

<<<<<<< HEAD
=======
//Sha256 加密
>>>>>>> upstream/master
func Sha256(b []byte) []byte {
	data := sha256.Sum256(b)
	return data[:]
}

<<<<<<< HEAD
=======
//ShaKeccak256 加密
>>>>>>> upstream/master
func ShaKeccak256(b []byte) []byte {
	data := sha3.KeccakSum256(b)
	return data[:]
}

<<<<<<< HEAD
=======
//Hex2Bytes hex -> []byte
>>>>>>> upstream/master
func Hex2Bytes(str string) ([]byte, error) {
	return hex.DecodeString(str)
}

func sha2Hash(b []byte, out []byte) {
	s := sha256.New()
	s.Write(b[:])
	tmp := s.Sum(nil)
	s.Reset()
	s.Write(tmp)
	copy(out[:], s.Sum(nil))
}

<<<<<<< HEAD
// Returns hash: SHA256( SHA256( data ) )
=======
// Sha2Sum Returns hash: SHA256( SHA256( data ) )
>>>>>>> upstream/master
// Where possible, using ShaHash() should be a bit faster
func Sha2Sum(b []byte) (out [32]byte) {
	sha2Hash(b, out[:])
	return
}

func rimpHash(in []byte, out []byte) {
	sha := sha256.New()
	sha.Write(in)
	rim := ripemd160.New()
	rim.Write(sha.Sum(nil)[:])
	copy(out, rim.Sum(nil))
}

<<<<<<< HEAD
// Returns hash: RIMP160( SHA256( data ) )
=======
// Rimp160AfterSha256 Returns hash: RIMP160( SHA256( data ) )
>>>>>>> upstream/master
// Where possible, using RimpHash() should be a bit faster
func Rimp160AfterSha256(b []byte) (out [20]byte) {
	rimpHash(b, out[:])
	return
}

<<<<<<< HEAD
=======
//RandKey 随机key
>>>>>>> upstream/master
func RandKey() (ret [32]byte) {
	_, err := io.ReadFull(rand.Reader, ret[:])
	if err != nil {
		panic(err)
	}
	return
}
