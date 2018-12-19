// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package strategy 实现开发者工具实现不同策略的功能
>>>>>>> upstream/master
package strategy

import (
	"fmt"

<<<<<<< HEAD
	"github.com/pkg/errors"
	"github.com/33cn/chain33/cmd/tools/types"
	"github.com/33cn/chain33/common/log/log15"
=======
	"github.com/33cn/chain33/cmd/tools/types"
	"github.com/33cn/chain33/common/log/log15"
	"github.com/pkg/errors"
>>>>>>> upstream/master
)

var (
	mlog = log15.New("module", "strategy")
)

<<<<<<< HEAD
=======
//Strategy 策略
>>>>>>> upstream/master
type Strategy interface {
	SetParam(key string, value string)
	Run() error
}

<<<<<<< HEAD
=======
//New new
>>>>>>> upstream/master
func New(name string) Strategy {
	switch name {
	case types.KeyImportPackage:
		return &importPackageStrategy{
			strategyBasic: strategyBasic{
				params: make(map[string]string),
			},
		}
	case types.KeyCreateSimpleExecProject:
		return &simpleCreateExecProjStrategy{
			strategyBasic: strategyBasic{
				params: make(map[string]string),
			},
		}
	case types.KeyCreateAdvanceExecProject:
		return &advanceCreateExecProjStrategy{
			strategyBasic: strategyBasic{
				params: make(map[string]string),
			},
		}
	case types.KeyUpdateInit:
		return &updateInitStrategy{
			strategyBasic: strategyBasic{
				params: make(map[string]string),
			},
		}
<<<<<<< HEAD
=======
	case types.KeyCreatePlugin:
		return &createPluginStrategy{
			strategyBasic: strategyBasic{
				params: make(map[string]string),
			},
		}
>>>>>>> upstream/master
	}
	return nil
}

type strategyBasic struct {
	params map[string]string
}

<<<<<<< HEAD
func (this *strategyBasic) SetParam(key string, value string) {
	this.params[key] = value
}

func (this *strategyBasic) getParam(key string) (string, error) {
	if v, ok := this.params[key]; ok {
		return v, nil
	}
	return "", errors.New(fmt.Sprintf("Key:%v not existed.", key))
}

func (this *strategyBasic) Run() error {
=======
//SetParam 设置参数
func (s *strategyBasic) SetParam(key string, value string) {
	s.params[key] = value
}

func (s *strategyBasic) getParam(key string) (string, error) {
	if v, ok := s.params[key]; ok {
		return v, nil
	}
	return "", fmt.Errorf("Key:%v not exist", key)
}

//Run 运行
func (s *strategyBasic) Run() error {
>>>>>>> upstream/master
	return errors.New("NotSupport")
}
