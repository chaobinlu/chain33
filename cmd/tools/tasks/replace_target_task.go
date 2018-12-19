// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tasks

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/33cn/chain33/cmd/tools/types"
	"github.com/33cn/chain33/util"
)

// ReplaceTargetTask 替换指定目录下所有文件的标志性文字
// 可替换的名字列表如下：
// ${PROJECTNAME}
// ${CLASSNAME}
// ${ACTIONNAME}
// ${EXECNAME}
type ReplaceTargetTask struct {
	TaskBase
	OutputPath  string
	ProjectName string
	ClassName   string
	ActionName  string
	ExecName    string
}

<<<<<<< HEAD
func (this *ReplaceTargetTask) GetName() string {
=======
//GetName 获取name
func (r *ReplaceTargetTask) GetName() string {
>>>>>>> upstream/master
	return "ReplaceTargetTask"
}

// Execute 执行具体的替换动作
// 1. 扫描指定的output路径
// 2. 打开每一个文件，根据替换规则替换内部的所有标签
// 3. 保存时查看文件名是否要替换，如果要则替换后保存，否则直接保存
// 4. 一直到所有的文件都替换完毕
<<<<<<< HEAD
func (this *ReplaceTargetTask) Execute() error {
	mlog.Info("Execute replace target task.")
	err := filepath.Walk(this.OutputPath, func(path string, info os.FileInfo, err error) error {
=======
func (r *ReplaceTargetTask) Execute() error {
	mlog.Info("Execute replace target task.")
	err := filepath.Walk(r.OutputPath, func(path string, info os.FileInfo, err error) error {
>>>>>>> upstream/master
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
<<<<<<< HEAD
		if err := this.replaceTarget(path); err != nil {
=======
		if err := r.replaceTarget(path); err != nil {
>>>>>>> upstream/master
			mlog.Error("replaceTarget error", "error", err, "path", path)
			return err
		}
		return nil
	})
	return err
}

<<<<<<< HEAD
func (this *ReplaceTargetTask) replaceTarget(file string) error {
=======
func (r *ReplaceTargetTask) replaceTarget(file string) error {
>>>>>>> upstream/master
	replacePairs := []struct {
		src string
		dst string
	}{
<<<<<<< HEAD
		{src: types.TagProjectName, dst: this.ProjectName},
		{src: types.TagClassName, dst: this.ClassName},
		{src: types.TagActionName, dst: this.ActionName},
		{src: types.TagExecName, dst: this.ExecName},
=======
		{src: types.TagProjectName, dst: r.ProjectName},
		{src: types.TagClassName, dst: r.ClassName},
		{src: types.TagActionName, dst: r.ActionName},
		{src: types.TagExecName, dst: r.ExecName},
>>>>>>> upstream/master
	}
	bcontent, err := util.ReadFile(file)
	if err != nil {
		return err
	}
	content := string(bcontent)
	for _, pair := range replacePairs {
		content = strings.Replace(content, pair.src, pair.dst, -1)
	}
	util.DeleteFile(file)
	_, err = util.WriteStringToFile(file, content)
	return err
}
