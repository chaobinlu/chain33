// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"errors"
	"sync"
	"time"

	"github.com/33cn/chain33/types"
)

<<<<<<< HEAD
=======
//Task 任务
>>>>>>> upstream/master
type Task struct {
	sync.Mutex
	cond     *sync.Cond
	start    int64
	end      int64
	isruning bool
	ticker   *time.Timer
	timeout  time.Duration
	cb       func()
	donelist map[int64]struct{}
}

func newTask(timeout time.Duration) *Task {
	t := &Task{}
	t.timeout = timeout
	t.ticker = time.NewTimer(t.timeout)
	t.cond = sync.NewCond(t)
	go t.tick()
	return t
}

func (t *Task) tick() {
	for {
		t.cond.L.Lock()
		for !t.isruning {
			t.cond.Wait()
		}
		t.cond.L.Unlock()
		_, ok := <-t.ticker.C
		if !ok {
			chainlog.Error("task is done", "timer is stop", t.start)
			continue
		}
		t.Lock()
		if err := t.stop(false); err == nil {
			chainlog.Debug("task is done", "timer is stop", t.start)
		}
		t.Unlock()
	}
}

<<<<<<< HEAD
=======
//InProgress 是否在执行
>>>>>>> upstream/master
func (t *Task) InProgress() bool {
	t.Lock()
	defer t.Unlock()
	return t.isruning
}

<<<<<<< HEAD
=======
//TimerReset 计时器重置
>>>>>>> upstream/master
func (t *Task) TimerReset(timeout time.Duration) {
	t.TimerStop()
	t.ticker.Reset(timeout)
}

<<<<<<< HEAD
=======
//TimerStop 计时器停止
>>>>>>> upstream/master
func (t *Task) TimerStop() {
	if !t.ticker.Stop() {
		select {
		case <-t.ticker.C:
		default:
		}
	}
}

<<<<<<< HEAD
=======
//Start 计时器启动
>>>>>>> upstream/master
func (t *Task) Start(start, end int64, cb func()) error {
	t.Lock()
	defer t.Unlock()
	if t.isruning {
<<<<<<< HEAD
		return errors.New("task is runing")
=======
		return errors.New("task is running")
>>>>>>> upstream/master
	}
	if start > end {
		return types.ErrStartBigThanEnd
	}
	chainlog.Debug("task start:", "start", start, "end", end)
	t.isruning = true
	t.TimerReset(t.timeout)
	t.start = start
	t.end = end
	t.cb = cb
	t.donelist = make(map[int64]struct{})
	t.cond.Signal()
	return nil
}

<<<<<<< HEAD
=======
//Done 任务完成
>>>>>>> upstream/master
func (t *Task) Done(height int64) {
	t.Lock()
	defer t.Unlock()
	if !t.isruning {
		return
	}
	if height >= t.start && height <= t.end {
		chainlog.Debug("done", "height", height)
		t.done(height)
		t.TimerReset(t.timeout)
	}
}

func (t *Task) stop(runcb bool) error {
	if !t.isruning {
<<<<<<< HEAD
		return errors.New("not runing")
=======
		return errors.New("not running")
>>>>>>> upstream/master
	}
	t.isruning = false
	if t.cb != nil && runcb {
		go t.cb()
	}
	t.TimerStop()
	return nil
}

<<<<<<< HEAD
=======
//Cancel 任务取消
>>>>>>> upstream/master
func (t *Task) Cancel() error {
	t.Lock()
	defer t.Unlock()
	chainlog.Warn("----task is cancel----")
	return t.stop(false)
}

func (t *Task) done(height int64) {
	if height == t.start {
		t.start = t.start + 1
		for i := t.start; i <= t.end; i++ {
			_, ok := t.donelist[i]
			if !ok {
				break
			}
			delete(t.donelist, i)
			t.start = i + 1
			//任务完成
		}
		if t.start > t.end {
			chainlog.Debug("----task is done----")
			t.stop(true)
		}
	}
	t.donelist[height] = struct{}{}
}
