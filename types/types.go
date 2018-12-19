// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package types 实现了chain33基础结构体、接口、常量等的定义
>>>>>>> upstream/master
package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

<<<<<<< HEAD
	"github.com/golang/protobuf/proto"
=======
>>>>>>> upstream/master
	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/crypto"
	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types/jsonpb"
<<<<<<< HEAD

=======
	"github.com/golang/protobuf/proto"

	// 注册system的crypto 加密算法
>>>>>>> upstream/master
	_ "github.com/33cn/chain33/system/crypto/init"
)

var tlog = log.New("module", "types")

<<<<<<< HEAD
const Size_1K_shiftlen uint = 10

type Message proto.Message

type Query4Cli struct {
	Execer   string      `json:"execer"`
	FuncName string      `json:"funcName"`
	Payload  interface{} `json:"payload"`
}

//交易组的接口，Transactions 和 Transaction 都符合这个接口
=======
// Size1Kshiftlen tx消息大小1k
const Size1Kshiftlen uint = 10

// Message 声明proto.Message
type Message proto.Message

//TxGroup 交易组的接口，Transactions 和 Transaction 都符合这个接口
>>>>>>> upstream/master
type TxGroup interface {
	Tx() *Transaction
	GetTxGroup() (*Transactions, error)
	CheckSign() bool
}

<<<<<<< HEAD
func ExecName(name string) string {
=======
//ExecName  执行器name
func ExecName(name string) string {
	if len(name) > 1 && name[0] == '#' {
		return name[1:]
	}
>>>>>>> upstream/master
	if IsParaExecName(name) {
		return name
	}
	if IsPara() {
		return GetTitle() + name
	}
	return name
}

<<<<<<< HEAD
//默认的allow 规则->根据 GetRealExecName 来判断
=======
//IsAllowExecName 默认的allow 规则->根据 GetRealExecName 来判断
>>>>>>> upstream/master
//name 必须大于3 小于 100
func IsAllowExecName(name []byte, execer []byte) bool {
	// name长度不能超过系统限制
	if len(name) > address.MaxExecNameLength || len(execer) > address.MaxExecNameLength {
		return false
	}
	if len(name) < 3 || len(execer) < 3 {
		return false
	}
	// name中不允许有 "-"
<<<<<<< HEAD
	if bytes.Contains(name, slash) {
=======
	if bytes.Contains(name, slash) || bytes.Contains(name, sharp) {
>>>>>>> upstream/master
		return false
	}
	if !bytes.Equal(name, execer) && !bytes.Equal(name, GetRealExecName(execer)) {
		return false
	}
	if bytes.HasPrefix(name, UserKey) {
		return true
	}
	for i := range AllowUserExec {
		if bytes.Equal(AllowUserExec[i], name) {
			return true
		}
	}
	return false
}

var bytesExec = []byte("exec-")
var commonPrefix = []byte("mavl-")

<<<<<<< HEAD
=======
//GetExecKey  获取执行器key
>>>>>>> upstream/master
func GetExecKey(key []byte) (string, bool) {
	n := 0
	start := 0
	end := 0
	for i := len(commonPrefix); i < len(key); i++ {
		if key[i] == '-' {
			n = n + 1
			if n == 2 {
				start = i + 1
			}
			if n == 3 {
				end = i
				break
			}
		}
	}
	if start > 0 && end > 0 {
		if bytes.Equal(key[start:end+1], bytesExec) {
			//find addr
			start = end + 1
			for k := end; k < len(key); k++ {
				if key[k] == ':' { //end+1
					end = k
					return string(key[start:end]), true
				}
			}
		}
	}
	return "", false
}

<<<<<<< HEAD
=======
//FindExecer  查找执行器
>>>>>>> upstream/master
func FindExecer(key []byte) (execer []byte, err error) {
	if !bytes.HasPrefix(key, commonPrefix) {
		return nil, ErrMavlKeyNotStartWithMavl
	}
	for i := len(commonPrefix); i < len(key); i++ {
		if key[i] == '-' {
			return key[len(commonPrefix):i], nil
		}
	}
	return nil, ErrNoExecerInMavlKey
}

<<<<<<< HEAD
=======
//GetParaExec  获取平行链执行
>>>>>>> upstream/master
func GetParaExec(execer []byte) []byte {
	//必须是平行链
	if !IsPara() {
		return execer
	}
	//必须是相同的平行链
	if !strings.HasPrefix(string(execer), GetTitle()) {
		return execer
	}
	return execer[len(GetTitle()):]
}

<<<<<<< HEAD
func getParaExecName(execer []byte) []byte {
=======
//GetParaExecName 获取平行链上的执行器
func GetParaExecName(execer []byte) []byte {
>>>>>>> upstream/master
	if !bytes.HasPrefix(execer, ParaKey) {
		return execer
	}
	count := 0
	for i := 0; i < len(execer); i++ {
		if execer[i] == '.' {
			count++
		}
		if count == 3 && i < (len(execer)-1) {
			newexec := execer[i+1:]
			return newexec
		}
	}
	return execer
}

<<<<<<< HEAD
func GetRealExecName(execer []byte) []byte {
	//平行链执行器，获取真实执行器的规则
	execer = getParaExecName(execer)
=======
//GetRealExecName  获取真实的执行器name
func GetRealExecName(execer []byte) []byte {
	//平行链执行器，获取真实执行器的规则
	execer = GetParaExecName(execer)
>>>>>>> upstream/master
	//平行链嵌套平行链是不被允许的
	if bytes.HasPrefix(execer, ParaKey) {
		return execer
	}
	if bytes.HasPrefix(execer, UserKey) {
		//不是user.p. 的情况, 而是user. 的情况
		count := 0
		index := 0
		for i := 0; i < len(execer); i++ {
			if execer[i] == '.' {
				count++
			}
			index = i
			if count == 2 {
				index--
				break
			}
		}
		e := execer[len(UserKey) : index+1]
		if len(e) > 0 {
			return e
		}
	}
	return execer
}

<<<<<<< HEAD
=======
//Encode  编码
>>>>>>> upstream/master
func Encode(data proto.Message) []byte {
	b, err := proto.Marshal(data)
	if err != nil {
		panic(err)
	}
	return b
}

<<<<<<< HEAD
=======
//Size  消息大小
>>>>>>> upstream/master
func Size(data proto.Message) int {
	return proto.Size(data)
}

<<<<<<< HEAD
=======
//Decode  解码
>>>>>>> upstream/master
func Decode(data []byte, msg proto.Message) error {
	return proto.Unmarshal(data, msg)
}

<<<<<<< HEAD
func JsonToPB(data []byte, msg proto.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(data), msg)
}

=======
//JSONToPB  JSON格式转换成protobuffer格式
func JSONToPB(data []byte, msg proto.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(data), msg)
}

//Hash  计算叶子节点的hash
>>>>>>> upstream/master
func (leafnode *LeafNode) Hash() []byte {
	data, err := proto.Marshal(leafnode)
	if err != nil {
		panic(err)
	}
	return common.Sha256(data)
}

<<<<<<< HEAD
=======
//Hash  计算中间节点的hash
>>>>>>> upstream/master
func (innernode *InnerNode) Hash() []byte {
	rightHash := innernode.RightHash
	leftHash := innernode.LeftHash
	hashLen := len(common.Hash{})
	if len(innernode.RightHash) > hashLen {
		innernode.RightHash = innernode.RightHash[len(innernode.RightHash)-hashLen:]
	}
	if len(innernode.LeftHash) > hashLen {
		innernode.LeftHash = innernode.LeftHash[len(innernode.LeftHash)-hashLen:]
	}
	data, err := proto.Marshal(innernode)
	if err != nil {
		panic(err)
	}
	innernode.RightHash = rightHash
	innernode.LeftHash = leftHash
	return common.Sha256(data)
}

<<<<<<< HEAD
func NewErrReceipt(err error) *Receipt {
	berr := err.Error()
	errlog := &ReceiptLog{TyLogErr, []byte(berr)}
	return &Receipt{ExecErr, nil, []*ReceiptLog{errlog}}
}

=======
//NewErrReceipt  new一个新的Receipt
func NewErrReceipt(err error) *Receipt {
	berr := err.Error()
	errlog := &ReceiptLog{Ty: TyLogErr, Log: []byte(berr)}
	return &Receipt{Ty: ExecErr, KV: nil, Logs: []*ReceiptLog{errlog}}
}

//CheckAmount  检测转账金额
>>>>>>> upstream/master
func CheckAmount(amount int64) bool {
	if amount <= 0 || amount >= MaxCoin {
		return false
	}
	return true
}

<<<<<<< HEAD
=======
//GetEventName  获取时间name通过事件id
>>>>>>> upstream/master
func GetEventName(event int) string {
	name, ok := eventName[event]
	if ok {
		return name
	}
	return "unknow-event"
}

<<<<<<< HEAD
=======
//GetSignName  获取签名类型
>>>>>>> upstream/master
func GetSignName(execer string, signType int) string {
	//优先加载执行器的签名类型
	if execer != "" {
		exec := LoadExecutorType(execer)
		if exec != nil {
			name, err := exec.GetCryptoDriver(signType)
			if err == nil {
				return name
			}
		}
	}
	//加载系统执行器的签名类型
	return crypto.GetName(signType)
}

<<<<<<< HEAD
=======
//GetSignType  获取签名类型
>>>>>>> upstream/master
func GetSignType(execer string, name string) int {
	//优先加载执行器的签名类型
	if execer != "" {
		exec := LoadExecutorType(execer)
		if exec != nil {
			ty, err := exec.GetCryptoType(name)
			if err == nil {
				return ty
			}
		}
	}
	//加载系统执行器的签名类型
	return crypto.GetType(name)
}

<<<<<<< HEAD
var ConfigPrefix = "mavl-config-"

// 原来实现有bug， 但生成的key在状态树里， 不可修改
// mavl-config–{key}  key 前面两个-
=======
// ConfigPrefix 配置前缀key
var ConfigPrefix = "mavl-config-"

// ConfigKey 原来实现有bug， 但生成的key在状态树里， 不可修改
// mavl-config–{key}  key 前面两个-
>>>>>>> upstream/master
func ConfigKey(key string) string {
	return fmt.Sprintf("%s-%s", ConfigPrefix, key)
}

<<<<<<< HEAD
var ManagePrefix = "mavl-"

=======
// ManagePrefix 超级管理员账户配置前缀key
var ManagePrefix = "mavl-"

//ManageKey 超级管理员账户key
>>>>>>> upstream/master
func ManageKey(key string) string {
	return fmt.Sprintf("%s-%s", ManagePrefix+"manage", key)
}

<<<<<<< HEAD
func ManaeKeyWithHeigh(key string, height int64) string {
	if IsFork(height, "ForkExecKey") {
		return ManageKey(key)
	} else {
		return ConfigKey(key)
	}
}

=======
//ManaeKeyWithHeigh 超级管理员账户key
func ManaeKeyWithHeigh(key string, height int64) string {
	if IsFork(height, "ForkExecKey") {
		return ManageKey(key)
	}
	return ConfigKey(key)
}

//ReceiptDataResult 回执数据
>>>>>>> upstream/master
type ReceiptDataResult struct {
	Ty     int32               `json:"ty"`
	TyName string              `json:"tyname"`
	Logs   []*ReceiptLogResult `json:"logs"`
}

<<<<<<< HEAD
=======
//ReceiptLogResult 回执log数据
>>>>>>> upstream/master
type ReceiptLogResult struct {
	Ty     int32       `json:"ty"`
	TyName string      `json:"tyname"`
	Log    interface{} `json:"log"`
	RawLog string      `json:"rawlog"`
}

<<<<<<< HEAD
=======
//DecodeReceiptLog 编码回执数据
>>>>>>> upstream/master
func (r *ReceiptData) DecodeReceiptLog(execer []byte) (*ReceiptDataResult, error) {
	result := &ReceiptDataResult{Ty: r.GetTy()}
	switch r.Ty {
	case 0:
		result.TyName = "ExecErr"
	case 1:
		result.TyName = "ExecPack"
	case 2:
		result.TyName = "ExecOk"
	default:
		return nil, ErrLogType
	}

	logs := r.GetLogs()
	for _, l := range logs {
		var lTy string
		var logIns interface{}
		lLog, err := hex.DecodeString(common.ToHex(l.GetLog())[2:])
		if err != nil {
			return nil, err
		}

		logType := LoadLog(execer, int64(l.Ty))
		if logType == nil {
			//tlog.Error("DecodeReceiptLog:", "Faile to decodeLog with type value logtype", l.Ty)
			return nil, ErrLogType
		}

<<<<<<< HEAD
		logIns, err = logType.Decode(lLog)
=======
		logIns, _ = logType.Decode(lLog)
>>>>>>> upstream/master
		lTy = logType.Name()

		result.Logs = append(result.Logs, &ReceiptLogResult{Ty: l.Ty, TyName: lTy, Log: logIns, RawLog: common.ToHex(l.GetLog())})
	}
	return result, nil
}

<<<<<<< HEAD
=======
//OutputReceiptDetails 输出回执数据详情
>>>>>>> upstream/master
func (r *ReceiptData) OutputReceiptDetails(execer []byte, logger log.Logger) {
	rds, err := r.DecodeReceiptLog(execer)
	if err == nil {
		logger.Debug("receipt decode", "receipt data", rds)
		for _, rdl := range rds.Logs {
			logger.Debug("receipt log", "log", rdl)
		}
	} else {
		logger.Error("decodelogerr", "err", err)
	}
}

<<<<<<< HEAD
=======
//IterateRangeByStateHash 迭代查找
>>>>>>> upstream/master
func (t *ReplyGetTotalCoins) IterateRangeByStateHash(key, value []byte) bool {
	fmt.Println("ReplyGetTotalCoins.IterateRangeByStateHash", "key", string(key))
	var acc Account
	err := Decode(value, &acc)
	if err != nil {
		tlog.Error("ReplyGetTotalCoins.IterateRangeByStateHash", "err", err)
		return true
	}
	//tlog.Info("acc:", "value", acc)
	if t.Num >= t.Count {
		t.NextKey = key
		return true
	}
	t.Num++
	t.Amount += acc.Balance
	return false
}

// GetTxTimeInterval 获取交易有效期
func GetTxTimeInterval() time.Duration {
	return time.Second * 120
}

<<<<<<< HEAD
=======
// ParaCrossTx 平行跨链交易
>>>>>>> upstream/master
type ParaCrossTx interface {
	IsParaCrossTx() bool
}

<<<<<<< HEAD
func PBToJson(r Message) ([]byte, error) {
=======
// PBToJSON 消息类型转换
func PBToJSON(r Message) ([]byte, error) {
>>>>>>> upstream/master
	encode := &jsonpb.Marshaler{EmitDefaults: true}
	var buf bytes.Buffer
	if err := encode.Marshal(&buf, r); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

<<<<<<< HEAD
func MustDecode(data []byte, v interface{}) {
=======
//MustPBToJSON panic when error
func MustPBToJSON(req Message) []byte {
	data, err := PBToJSON(req)
	if err != nil {
		panic(err)
	}
	return data
}

// MustDecode 数据是否已经编码
func MustDecode(data []byte, v interface{}) {
	if data == nil {
		return
	}
>>>>>>> upstream/master
	err := json.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}
}
<<<<<<< HEAD
=======

// AddItem 添加item
func (t *ReplyGetExecBalance) AddItem(execAddr, value []byte) {
	var acc Account
	err := Decode(value, &acc)
	if err != nil {
		tlog.Error("ReplyGetExecBalance.AddItem", "err", err)
		return
	}
	tlog.Info("acc:", "value", acc)
	t.Amount += acc.Balance
	t.Amount += acc.Frozen

	t.AmountActive += acc.Balance
	t.AmountFrozen += acc.Frozen

	item := &ExecBalanceItem{ExecAddr: execAddr, Frozen: acc.Frozen, Active: acc.Balance}
	t.Items = append(t.Items, item)
}
>>>>>>> upstream/master
