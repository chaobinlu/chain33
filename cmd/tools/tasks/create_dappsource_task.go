// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tasks

import (
	"fmt"
	"regexp"
	"strings"

<<<<<<< HEAD
	"github.com/pkg/errors"
=======
>>>>>>> upstream/master
	"github.com/33cn/chain33/cmd/tools/types"
	"github.com/33cn/chain33/util"
)

type actionInfoItem struct {
	memberName string
	memberType string
}

// CreateDappSourceTask 通过生成好的pb.go和预先设计的模板，生成反射程序源码
type CreateDappSourceTask struct {
	TaskBase
	TemplatePath       string // 生成最终源码时的模板路径
	OutputPath         string
	ProjectName        string
	ClsName            string // 生成源码的类名
	ActionName         string // 生成源码的Action类名
	TypeName           string
	ExecuteName        string
	ProtoFile          string // 推导的原始proto文件
	ExecHeaderTempFile string
	TypeTempFile       string
	TypeOutputFile     string

	actionInfos           []*actionInfoItem // Action中的成员变量名称PB格式
	execHeaderTempContent string
}

<<<<<<< HEAD
func (this *CreateDappSourceTask) GetName() string {
	return "CreateDappSourceTask"
}

func (this *CreateDappSourceTask) Execute() error {
	mlog.Info("Execute create build app source task.")
	if err := this.init(); err != nil {
		return err
	}
	if err := this.readActionMemberNames(); err != nil {
		return err
	}
	if err := this.createExecFile(); err != nil {
		return err
	}
	if err := this.createExecLocalFile(); err != nil {
		return err
	}
	if err := this.createExecDelLocalFile(); err != nil {
		return err
	}
	if err := this.createTypeExecuteFile(); err != nil {
=======
//GetName 获取name
func (c *CreateDappSourceTask) GetName() string {
	return "CreateDappSourceTask"
}

//Execute 执行
func (c *CreateDappSourceTask) Execute() error {
	mlog.Info("Execute create build app source task.")
	if err := c.init(); err != nil {
		return err
	}
	if err := c.readActionMemberNames(); err != nil {
		return err
	}
	if err := c.createExecFile(); err != nil {
		return err
	}
	if err := c.createExecLocalFile(); err != nil {
		return err
	}
	if err := c.createExecDelLocalFile(); err != nil {
		return err
	}
	if err := c.createTypeExecuteFile(); err != nil {
>>>>>>> upstream/master
		return err
	}
	return nil
}

<<<<<<< HEAD
func (this *CreateDappSourceTask) init() error {
	if !util.CheckFileIsExist(this.ExecHeaderTempFile) {
		return errors.New(fmt.Sprintf("File %s not existed.", this.ExecHeaderTempFile))
	}
	contentbt, err := util.ReadFile(this.ExecHeaderTempFile)
	if err != nil {
		return errors.New(fmt.Sprintf("Read file %s failed. error %q", this.ExecHeaderTempFile, err))
	}
	content := strings.Replace(string(contentbt), types.TagClassName, this.ClsName, -1)
	content = strings.Replace(content, types.TagExecName, this.ExecuteName, -1)
	this.execHeaderTempContent = content
=======
func (c *CreateDappSourceTask) init() error {
	if !util.CheckFileIsExist(c.ExecHeaderTempFile) {
		return fmt.Errorf("file %s not exist", c.ExecHeaderTempFile)
	}
	contentbt, err := util.ReadFile(c.ExecHeaderTempFile)
	if err != nil {
		return fmt.Errorf("read file %s failed. error %q", c.ExecHeaderTempFile, err)
	}
	content := strings.Replace(string(contentbt), types.TagClassName, c.ClsName, -1)
	content = strings.Replace(content, types.TagExecName, c.ExecuteName, -1)
	c.execHeaderTempContent = content
>>>>>>> upstream/master
	return nil
}

/**
通过正则获取Action的成员变量名和类型，其具体操作步骤如下：
1. 读取需要解析的proto文件
2. 通过搜索，定位到指定Action的起始为止
3. 使用正则获取该Action中的oneof Value的内容
4. 使用正则解析oneof Value中的内容，获取变量名和类型名
5. 将获取到的变量名去除空格，并将首字母大写
*/
<<<<<<< HEAD
func (this *CreateDappSourceTask) readActionMemberNames() error {
	pbContext, err := util.ReadFile(this.ProtoFile)
=======
func (c *CreateDappSourceTask) readActionMemberNames() error {
	pbContext, err := util.ReadFile(c.ProtoFile)
>>>>>>> upstream/master
	if err != nil {
		return err
	}
	context := string(pbContext)
	// 如果文件中含有与ActionName部分匹配的文字，则会造成搜索到多个
<<<<<<< HEAD
	index := strings.Index(context, this.ActionName)
	if index < 0 {
		return errors.New(fmt.Sprintf("Action %s Not Existed", this.ActionName))
=======
	index := strings.Index(context, c.ActionName)
	if index < 0 {
		return fmt.Errorf("Action %s Not Existed", c.ActionName)
>>>>>>> upstream/master
	}
	expr := fmt.Sprintf(`\s*oneof\s+value\s*{\s+([\w\s=;]*)\}`)
	reg := regexp.MustCompile(expr)
	oneOfValueStrs := reg.FindAllStringSubmatch(string(pbContext), index)

	expr = fmt.Sprintf(`\s+(\w+)([\s\w]+)=\s+(\d+);`)
	reg = regexp.MustCompile(expr)
	members := reg.FindAllStringSubmatch(oneOfValueStrs[0][0], -1)

<<<<<<< HEAD
	this.actionInfos = make([]*actionInfoItem, 0)
=======
	c.actionInfos = make([]*actionInfoItem, 0)
>>>>>>> upstream/master
	for _, member := range members {
		memberType := strings.Replace(member[1], " ", "", -1)
		memberName := strings.Replace(member[2], " ", "", -1)
		// 根据proto生成pb.go的规则，成员变量首字母必须大写
		memberName, _ = util.MakeStringToUpper(memberName, 0, 1)
<<<<<<< HEAD
		this.actionInfos = append(this.actionInfos, &actionInfoItem{
=======
		c.actionInfos = append(c.actionInfos, &actionInfoItem{
>>>>>>> upstream/master
			memberName: memberName,
			memberType: memberType,
		})
	}
<<<<<<< HEAD
	if len(this.actionInfos) == 0 {
		return errors.New(fmt.Sprintf("Can Not Find %s Member Info", this.ActionName))
=======
	if len(c.actionInfos) == 0 {
		return fmt.Errorf("Can Not Find %s Member Info", c.ActionName)
>>>>>>> upstream/master
	}
	return nil
}

<<<<<<< HEAD
func (this *CreateDappSourceTask) createExecFile() error {
	fnFmtStr := `func (this *%s) Exec_%s(payload *ptypes.%s, tx *types.Transaction, index int) (*types.Receipt, error) {
=======
func (c *CreateDappSourceTask) createExecFile() error {
	fnFmtStr := `func (c *%s) Exec_%s(payload *ptypes.%s, tx *types.Transaction, index int) (*types.Receipt, error) {
>>>>>>> upstream/master
	return &types.Receipt{}, nil
}

`
<<<<<<< HEAD
	content := this.execHeaderTempContent
	for _, info := range this.actionInfos {
		content += fmt.Sprintf(fnFmtStr, this.ClsName, info.memberName, info.memberType)
	}
	fileName := fmt.Sprintf("%s/executor/exec.go", this.OutputPath)
=======
	content := c.execHeaderTempContent
	for _, info := range c.actionInfos {
		content += fmt.Sprintf(fnFmtStr, c.ClsName, info.memberName, info.memberType)
	}
	fileName := fmt.Sprintf("%s/executor/exec.go", c.OutputPath)
>>>>>>> upstream/master
	_, err := util.WriteStringToFile(fileName, content)
	if err != nil {
		mlog.Error(fmt.Sprintf("Write to file %s failed. error %q", fileName, err))
		return err
	}
	return nil
}

<<<<<<< HEAD
func (this *CreateDappSourceTask) createExecLocalFile() error {
	fnFmtStr := `func (this *%s) ExecLocal_%s(payload *ptypes.%s, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
=======
func (c *CreateDappSourceTask) createExecLocalFile() error {
	fnFmtStr := `func (c *%s) ExecLocal_%s(payload *ptypes.%s, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
>>>>>>> upstream/master
	return &types.LocalDBSet{}, nil
}

`
<<<<<<< HEAD
	content := this.execHeaderTempContent
	for _, info := range this.actionInfos {
		content += fmt.Sprintf(fnFmtStr, this.ClsName, info.memberName, info.memberType)
	}
	fileName := fmt.Sprintf("%s/executor/exec_local.go", this.OutputPath)
=======
	content := c.execHeaderTempContent
	for _, info := range c.actionInfos {
		content += fmt.Sprintf(fnFmtStr, c.ClsName, info.memberName, info.memberType)
	}
	fileName := fmt.Sprintf("%s/executor/exec_local.go", c.OutputPath)
>>>>>>> upstream/master
	_, err := util.WriteStringToFile(fileName, content)
	if err != nil {
		mlog.Error(fmt.Sprintf("Write to file %s failed. error %q", fileName, err))
		return err
	}
	return nil
}

<<<<<<< HEAD
func (this *CreateDappSourceTask) createExecDelLocalFile() error {
	fnFmtStr := `func (this *%s) ExecDelLocal_%s(payload *ptypes.%s, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
=======
func (c *CreateDappSourceTask) createExecDelLocalFile() error {
	fnFmtStr := `func (c *%s) ExecDelLocal_%s(payload *ptypes.%s, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
>>>>>>> upstream/master
	return &types.LocalDBSet{}, nil
}

`
<<<<<<< HEAD
	content := this.execHeaderTempContent
	for _, info := range this.actionInfos {
		content += fmt.Sprintf(fnFmtStr, this.ClsName, info.memberName, info.memberType)
	}
	fileName := fmt.Sprintf("%s/executor/exec_del_local.go", this.OutputPath)
=======
	content := c.execHeaderTempContent
	for _, info := range c.actionInfos {
		content += fmt.Sprintf(fnFmtStr, c.ClsName, info.memberName, info.memberType)
	}
	fileName := fmt.Sprintf("%s/executor/exec_del_local.go", c.OutputPath)
>>>>>>> upstream/master
	_, err := util.WriteStringToFile(fileName, content)
	if err != nil {
		mlog.Error(fmt.Sprintf("Write to file %s failed. error %q", fileName, err))
		return err
	}
	return nil
}

// 组成规则是 TyLog+ActionName + ActionMemberName
<<<<<<< HEAD
func (this *CreateDappSourceTask) buildActionLogTypeText() (text string, err error) {
	items := fmt.Sprintf("TyLog%sUnknown = iota\r\n", this.ExecuteName)
	for _, info := range this.actionInfos {
		items += fmt.Sprintf("TyLog%s%s\r\n", this.ExecuteName, info.memberName)
=======
func (c *CreateDappSourceTask) buildActionLogTypeText() (text string, err error) {
	items := fmt.Sprintf("TyLog%sUnknown = iota\r\n", c.ExecuteName)
	for _, info := range c.actionInfos {
		items += fmt.Sprintf("TyLog%s%s\r\n", c.ExecuteName, info.memberName)
>>>>>>> upstream/master
	}
	text = fmt.Sprintf("const (\r\n%s)\r\n", items)
	return
}

// 组成规则是 ActionName + ActionMemberName
<<<<<<< HEAD
func (this *CreateDappSourceTask) buildActionIDText() (text string, err error) {
	var items string
	for index, info := range this.actionInfos {
		items += fmt.Sprintf("%sAction%s = %d\r\n", this.ExecuteName, info.memberName, index)
=======
func (c *CreateDappSourceTask) buildActionIDText() (text string, err error) {
	var items string
	for index, info := range c.actionInfos {
		items += fmt.Sprintf("%sAction%s = %d\r\n", c.ExecuteName, info.memberName, index)
>>>>>>> upstream/master
	}
	text = fmt.Sprintf("const (\r\n%s)\r\n", items)
	return
}

// 返回 map[int64]*types.LogInfo
<<<<<<< HEAD
func (this *CreateDappSourceTask) buildLogMapText() (text string, err error) {
	var items string
	for _, info := range this.actionInfos {
		items += fmt.Sprintf("\"%s\": %sAction%s,\r\n", info.memberName, this.ExecuteName, info.memberName)
=======
func (c *CreateDappSourceTask) buildLogMapText() (text string, err error) {
	var items string
	for _, info := range c.actionInfos {
		items += fmt.Sprintf("\"%s\": %sAction%s,\r\n", info.memberName, c.ExecuteName, info.memberName)
>>>>>>> upstream/master
	}
	text = fmt.Sprintf("map[string]int32{\r\n%s}", items)
	return
}

// 返回 map[string]*types.LogInfo
<<<<<<< HEAD
func (this *CreateDappSourceTask) buidTypeMapText() (text string, err error) {
=======
func (c *CreateDappSourceTask) buidTypeMapText() (text string, err error) {
>>>>>>> upstream/master
	text = fmt.Sprintf("map[int64]*types.LogInfo{\r\n}")
	return
}

/**
createTypeExecuteFile 根据自己的需求，创建一个types中与执行器同名的Type对照关系
需要处理的内容：
1. 定义TyLogXXXX的常量，规则是 TyLog + 变量名称
2. 定义类型常量，规则是 ActionName + 变量名称
3. 实现GetLogMap()
4. 实现GetTypeMap()
*/
<<<<<<< HEAD
func (this *CreateDappSourceTask) createTypeExecuteFile() error {
	logText, err := this.buildActionLogTypeText() // ${TYLOGACTIONTYPE}
	if err != nil {
		return err
	}
	actionIDText, err := this.buildActionIDText() // ${ACTIONIDTEXT}
	if err != nil {
		return err
	}
	logMapText, err := this.buildLogMapText() // ${LOGMAPTEXT}
	if err != nil {
		return err
	}
	typeMapText, err := this.buidTypeMapText() // ${TYPEMAPTEXT}
=======
func (c *CreateDappSourceTask) createTypeExecuteFile() error {
	logText, err := c.buildActionLogTypeText() // ${TYLOGACTIONTYPE}
	if err != nil {
		return err
	}
	actionIDText, err := c.buildActionIDText() // ${ACTIONIDTEXT}
	if err != nil {
		return err
	}
	logMapText, err := c.buildLogMapText() // ${LOGMAPTEXT}
	if err != nil {
		return err
	}
	typeMapText, err := c.buidTypeMapText() // ${TYPEMAPTEXT}
>>>>>>> upstream/master
	if err != nil {
		return err
	}

	replacePairs := []struct {
		src string
		dst string
	}{
		{src: types.TagTyLogActionType, dst: logText},
		{src: types.TagActionIDText, dst: actionIDText},
		{src: types.TagLogMapText, dst: logMapText},
		{src: types.TagTypeMapText, dst: typeMapText},
<<<<<<< HEAD
		{src: types.TagTypeName, dst: this.TypeName},
		{src: types.TagExecName, dst: this.ExecuteName},
		{src: types.TagActionName, dst: this.ActionName},
	}
	bcontent, err := util.ReadFile(this.TypeTempFile)
=======
		{src: types.TagTypeName, dst: c.TypeName},
		{src: types.TagExecName, dst: c.ExecuteName},
		{src: types.TagActionName, dst: c.ActionName},
	}
	bcontent, err := util.ReadFile(c.TypeTempFile)
>>>>>>> upstream/master
	if err != nil {
		return err
	}
	content := string(bcontent)
	for _, pair := range replacePairs {
		content = strings.Replace(content, pair.src, pair.dst, -1)
	}
<<<<<<< HEAD
	fileName := fmt.Sprintf("%s%s.go", this.TypeOutputFile, this.ClsName)
	util.DeleteFile(fileName)
	_, err = util.WriteStringToFile(fileName, content)
	return err

	return nil
=======
	fileName := fmt.Sprintf("%s%s.go", c.TypeOutputFile, c.ClsName)
	util.DeleteFile(fileName)
	_, err = util.WriteStringToFile(fileName, content)
	return err
>>>>>>> upstream/master
}
