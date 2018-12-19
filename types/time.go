// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"sync/atomic"
	"time"
)

var deltaTime int64
<<<<<<< HEAD
=======

//NtpHosts ntp hosts
>>>>>>> upstream/master
var NtpHosts = []string{
	"time.windows.com:123",
	"ntp.ubuntu.com:123",
	"pool.ntp.org:123",
	"cn.pool.ntp.org:123",
	"time.apple.com:123",
}

<<<<<<< HEAD
//realtime - localtime
=======
//SetTimeDelta realtime - localtime
>>>>>>> upstream/master
//超过60s 不做修正
//为了系统的安全，我们只做小范围时间错误的修复
func SetTimeDelta(dt int64) {
	if dt > 300*int64(time.Second) || dt < -300*int64(time.Second) {
		dt = 0
	}
	atomic.StoreInt64(&deltaTime, dt)
}

<<<<<<< HEAD
=======
//Now 获取当前时间戳
>>>>>>> upstream/master
func Now() time.Time {
	dt := time.Duration(atomic.LoadInt64(&deltaTime))
	return time.Now().Add(dt)
}

<<<<<<< HEAD
=======
//Since Since时间
>>>>>>> upstream/master
func Since(t time.Time) time.Duration {
	return Now().Sub(t)
}
