// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strategy

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

<<<<<<< HEAD
	"github.com/BurntSushi/toml"
	"github.com/33cn/chain33/util"
=======
	"github.com/33cn/chain33/util"
	"github.com/BurntSushi/toml"
>>>>>>> upstream/master
)

const (
	dappFolderName      = "dapp"
	consensusFolderName = "consensus"
	storeFolderName     = "store"
	cryptoFolderName    = "crypto"
)

type pluginConfigItem struct {
	Type    string
	Gitrepo string
	Version string
}

type pluginItem struct {
	name    string
	gitRepo string
	version string
}

type importPackageStrategy struct {
	strategyBasic
	cfgFileName    string
	cfgItems       map[string]*pluginConfigItem
	projRootPath   string
	projPluginPath string
	items          map[string][]*pluginItem
}

<<<<<<< HEAD
func (this *importPackageStrategy) Run() error {
	mlog.Info("Begin run chain33 import packages.")
	defer mlog.Info("Run chain33 import packages finish.")
	return this.runImpl()
}

func (this *importPackageStrategy) runImpl() error {
	type STEP func() error
	steps := []STEP{
		this.readConfig,
		this.initData,
		this.generateImportFile,
		this.fetchPluginPackage,
=======
func (im *importPackageStrategy) Run() error {
	mlog.Info("Begin run chain33 import packages.")
	defer mlog.Info("Run chain33 import packages finish.")
	return im.runImpl()
}

func (im *importPackageStrategy) runImpl() error {
	type STEP func() error
	steps := []STEP{
		im.readConfig,
		im.initData,
		im.generateImportFile,
		im.fetchPluginPackage,
>>>>>>> upstream/master
	}

	for s, step := range steps {
		err := step()
		if err != nil {
			fmt.Println("call", s+1, "step error", err)
			return err
		}
	}
	return nil
}

<<<<<<< HEAD
func (this *importPackageStrategy) readConfig() error {
	mlog.Info("读取配置文件")
	conf, _ := this.getParam("conf")
=======
func (im *importPackageStrategy) readConfig() error {
	mlog.Info("读取配置文件")
	conf, _ := im.getParam("conf")
>>>>>>> upstream/master
	if conf == "" {
		return nil
	}
	if conf != "" {
<<<<<<< HEAD
		this.cfgFileName = conf
	}
	_, err := toml.DecodeFile(this.cfgFileName, &this.cfgItems)
	return err
}

func (this *importPackageStrategy) initData() error {
	mlog.Info("初始化数据")
	this.items = make(map[string][]*pluginItem)
=======
		im.cfgFileName = conf
	}
	_, err := toml.DecodeFile(im.cfgFileName, &im.cfgItems)
	return err
}

func (im *importPackageStrategy) initData() error {
	mlog.Info("初始化数据")
	im.items = make(map[string][]*pluginItem)
>>>>>>> upstream/master
	dappItems := make([]*pluginItem, 0)
	consensusItems := make([]*pluginItem, 0)
	storeItems := make([]*pluginItem, 0)
	cryptoItems := make([]*pluginItem, 0)

	//read current plugin dir
	//(分成两级，并且去掉了 init 目录)
<<<<<<< HEAD
	path, _ := this.getParam("path")
	dirlist, err := this.readPluginDir(path)
	if err != nil {
		return err
	}
	if this.cfgItems == nil {
		this.cfgItems = make(map[string]*pluginConfigItem)
	}
	for name, value := range dirlist {
		this.cfgItems[name] = value
	}
	out, _ := this.getParam("out")
	//输出新的配置文件
	if out != "" {
		buf := new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(this.cfgItems)
=======
	path, _ := im.getParam("path")
	dirlist, err := im.readPluginDir(path)
	if err != nil {
		return err
	}
	if im.cfgItems == nil {
		im.cfgItems = make(map[string]*pluginConfigItem)
	}
	for name, value := range dirlist {
		im.cfgItems[name] = value
	}
	out, _ := im.getParam("out")
	//输出新的配置文件
	if out != "" {
		buf := new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(im.cfgItems)
>>>>>>> upstream/master
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(out, buf.Bytes(), 0666)
		if err != nil {
			return err
		}
	}
<<<<<<< HEAD
	if len(this.cfgItems) == 0 {
		return errors.New("Config is empty.")
	}
	for name, cfgItem := range this.cfgItems {
=======
	if len(im.cfgItems) == 0 {
		return errors.New("empty config")
	}
	for name, cfgItem := range im.cfgItems {
>>>>>>> upstream/master
		splitdata := strings.Split(name, "-")
		if len(splitdata) == 2 {
			cfgItem.Type = splitdata[0]
			name = splitdata[1]
		}
		item := &pluginItem{
			name:    name,
			gitRepo: cfgItem.Gitrepo,
			version: cfgItem.Version,
		}
		switch cfgItem.Type {
		case dappFolderName:
			dappItems = append(dappItems, item)
		case consensusFolderName:
			consensusItems = append(consensusItems, item)
		case storeFolderName:
			storeItems = append(storeItems, item)
		case cryptoFolderName:
			cryptoItems = append(cryptoItems, item)
		default:
			fmt.Printf("type %s is not supported.\n", cfgItem.Type)
<<<<<<< HEAD
			return errors.New("Config error.")
		}
	}
	this.items[dappFolderName] = dappItems
	this.items[consensusFolderName] = consensusItems
	this.items[storeFolderName] = storeItems
	this.items[cryptoFolderName] = cryptoItems
	this.projRootPath = ""
	this.projPluginPath, _ = this.getParam("path")
=======
			return errors.New("config error")
		}
	}
	im.items[dappFolderName] = dappItems
	im.items[consensusFolderName] = consensusItems
	im.items[storeFolderName] = storeItems
	im.items[cryptoFolderName] = cryptoItems
	im.projRootPath = ""
	im.projPluginPath, _ = im.getParam("path")
>>>>>>> upstream/master
	return nil
}

func getDirList(path string) ([]string, error) {
	dirlist, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	dirs := make([]string, 0)
	for _, f := range dirlist {
		if f.IsDir() {
			if f.Name() == "." || f.Name() == ".." || f.Name() == "init" || f.Name() == ".git" {
				continue
			}
			dirs = append(dirs, f.Name())
		}
	}
	return dirs, nil
}

<<<<<<< HEAD
func (this *importPackageStrategy) readPluginDir(path string) (map[string]*pluginConfigItem, error) {
=======
func (im *importPackageStrategy) readPluginDir(path string) (map[string]*pluginConfigItem, error) {
>>>>>>> upstream/master
	dirlist, err := getDirList(path)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	packname, _ := this.getParam("packname")
=======
	packname, _ := im.getParam("packname")
>>>>>>> upstream/master
	conf := make(map[string]*pluginConfigItem)
	for _, ty := range dirlist {
		names, err := getDirList(path + "/" + ty)
		if err != nil {
			return nil, err
		}
		for _, name := range names {
			key := ty + "-" + name
			item := &pluginConfigItem{
				Type:    ty,
				Gitrepo: packname + "/" + ty + "/" + name,
			}
			conf[key] = item
		}
	}
	return conf, nil
}

<<<<<<< HEAD
func (this *importPackageStrategy) generateImportFile() error {
	mlog.Info("生成引用文件")
	importStrs := map[string]string{}
	for name, plugins := range this.items {
		for _, item := range plugins {
			importStrs[name] += fmt.Sprintf("\r\n_ \"%s\"", item.gitRepo)
=======
func (im *importPackageStrategy) generateImportFile() error {
	mlog.Info("生成引用文件")
	importStrs := map[string]string{}
	for name, plugins := range im.items {
		for _, item := range plugins {
			importStrs[name] += fmt.Sprintf("\r\n_ \"%s\" //auto gen", item.gitRepo)
>>>>>>> upstream/master
		}
	}
	for key, value := range importStrs {
		content := fmt.Sprintf("package init\r\n\r\nimport(%s\r\n)", value)
<<<<<<< HEAD
		initFile := fmt.Sprintf("%s/%s/init/init.go", this.projPluginPath, key)
=======
		initFile := fmt.Sprintf("%s/%s/init/init.go", im.projPluginPath, key)
>>>>>>> upstream/master
		util.MakeDir(initFile)

		{ // 写入到文件中
			util.DeleteFile(initFile)
			file, err := util.OpenFile(initFile)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.WriteString(file, content)
			if err != nil {
				return err
			}
		}
		// 格式化生成的文件
		cmd := exec.Command("gofmt", "-l", "-s", "-w", initFile)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

<<<<<<< HEAD
func (this *importPackageStrategy) fetchPlugin(gitrepo, version string) error {
=======
func (im *importPackageStrategy) fetchPlugin(gitrepo, version string) error {
>>>>>>> upstream/master
	var param string
	if len(version) > 0 {
		param = fmt.Sprintf("%s@%s", gitrepo, version)
	} else {
		param = gitrepo
	}
	cmd := exec.Command("govendor", "fetch", param)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// fetchPluginPackage 使用govendor来下载依赖包
<<<<<<< HEAD
func (this *importPackageStrategy) fetchPluginPackage() error {
	mlog.Info("下载插件源码包")
	pwd := util.Pwd()
	os.Chdir(this.projRootPath)
	defer os.Chdir(pwd)
	for _, plugins := range this.items {
=======
func (im *importPackageStrategy) fetchPluginPackage() error {
	mlog.Info("下载插件源码包")
	pwd := util.Pwd()
	os.Chdir(im.projRootPath)
	defer os.Chdir(pwd)
	for _, plugins := range im.items {
>>>>>>> upstream/master
		for _, plugin := range plugins {
			mlog.Info("同步插件", "repo", plugin.gitRepo, "version", plugin.version)
			if plugin.version == "" {
				//留给后面的 fetch +m
				continue
			}
<<<<<<< HEAD
			err := this.fetchPlugin(plugin.gitRepo, plugin.version)
=======
			err := im.fetchPlugin(plugin.gitRepo, plugin.version)
>>>>>>> upstream/master
			if err != nil {
				mlog.Info("同步插件包出错", "repo", plugin.gitRepo, "error", err.Error())
				return err
			}
		}
	}
	return nil
}
