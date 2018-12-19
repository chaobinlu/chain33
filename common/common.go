// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package common contains various helper functions.
>>>>>>> upstream/master
package common

import (
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

<<<<<<< HEAD
=======
//MinInt32 min
>>>>>>> upstream/master
func MinInt32(left, right int32) int32 {
	if left > right {
		return right
	}
	return left
}

<<<<<<< HEAD
=======
//MaxInt32 max
>>>>>>> upstream/master
func MaxInt32(left, right int32) int32 {
	if left > right {
		return left
	}
	return right
}

<<<<<<< HEAD
=======
//GetRandBytes 获取随机字节
>>>>>>> upstream/master
func GetRandBytes(min, max int) []byte {
	length := max
	if min < max {
		length = min + random.Intn(max-min)
	}
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = byte(random.Intn(256))
	}
	return result
}

<<<<<<< HEAD
=======
//GetRandString 获取随机字符串
>>>>>>> upstream/master
func GetRandString(length int) string {
	return string(GetRandBytes(length, length))
}

var printString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

<<<<<<< HEAD
=======
//GetRandPrintString 获取随机可打印字符串
>>>>>>> upstream/master
func GetRandPrintString(min, max int) string {
	l := max
	if min < max {
		l = min + random.Intn(max-min)
	}
	result := make([]byte, l)
	for i := 0; i < l; i++ {
		result[i] = printString[random.Intn(len(printString))]
	}
	return string(result)
}
