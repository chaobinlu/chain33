// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tasks

import "github.com/33cn/chain33/common/log/log15"

var (
	mlog = log15.New("module", "task")
)

<<<<<<< HEAD
=======
//TaskBase task base
>>>>>>> upstream/master
type TaskBase struct {
	NextTask Task
}

<<<<<<< HEAD
func (this *TaskBase) Execute() error {
	return nil
}

func (this *TaskBase) SetNext(t Task) {
	this.NextTask = t
}

func (this *TaskBase) Next() Task {
	return this.NextTask
=======
//Execute 执行
func (t *TaskBase) Execute() error {
	return nil
}

//SetNext set next
func (t *TaskBase) SetNext(task Task) {
	t.NextTask = task
}

//Next 获取next
func (t *TaskBase) Next() Task {
	return t.NextTask
>>>>>>> upstream/master
}
