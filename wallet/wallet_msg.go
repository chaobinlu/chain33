// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wallet

import (
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
	wcom "github.com/33cn/chain33/wallet/common"
)

<<<<<<< HEAD
func (wallet *Wallet) ProcRecvMsg() {
	defer wallet.wg.Done()
	for msg := range wallet.client.Recv() {
		walletlog.Debug("wallet recv", "msg", types.GetEventName(int(msg.Ty)), "Id", msg.Id)
=======
// ProcRecvMsg 处理消息循环
func (wallet *Wallet) ProcRecvMsg() {
	defer wallet.wg.Done()
	for msg := range wallet.client.Recv() {
		walletlog.Debug("wallet recv", "msg", types.GetEventName(int(msg.Ty)), "Id", msg.ID)
>>>>>>> upstream/master
		beg := types.Now()
		reply, err := wallet.ExecWallet(&msg)
		if err != nil {
			//only for test ,del when test end
			msg.Reply(wallet.api.NewMessage("", 0, err))
		} else {
			msg.Reply(wallet.api.NewMessage("", 0, reply))
		}
<<<<<<< HEAD
		walletlog.Debug("end process", "msg.id", msg.Id, "cost", types.Since(beg))
	}
}

=======
		walletlog.Debug("end process", "msg.id", msg.ID, "cost", types.Since(beg))
	}
}

// On_WalletGetAccountList 响应获取账户列表
>>>>>>> upstream/master
func (wallet *Wallet) On_WalletGetAccountList(req *types.ReqAccountList) (types.Message, error) {
	reply, err := wallet.ProcGetAccountList(req)
	if err != nil {
		walletlog.Error("onWalletGetAccountList", "err", err.Error())
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_NewAccount 响应新建账号
>>>>>>> upstream/master
func (wallet *Wallet) On_NewAccount(req *types.ReqNewAccount) (types.Message, error) {
	reply, err := wallet.ProcCreateNewAccount(req)
	if err != nil {
		walletlog.Error("onNewAccount", "err", err.Error())
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_WalletTransactionList 响应获取钱包交易列表
>>>>>>> upstream/master
func (wallet *Wallet) On_WalletTransactionList(req *types.ReqWalletTransactionList) (types.Message, error) {
	reply, err := wallet.ProcWalletTxList(req)
	if err != nil {
		walletlog.Error("ProcWalletTxList", "err", err.Error())
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_WalletImportPrivkey 响应导入私钥
>>>>>>> upstream/master
func (wallet *Wallet) On_WalletImportPrivkey(req *types.ReqWalletImportPrivkey) (types.Message, error) {
	reply, err := wallet.ProcImportPrivKey(req)
	if err != nil {
		walletlog.Error("ProcImportPrivKey", "err", err.Error())
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_WalletSendToAddress 响应钱包想地址转账
>>>>>>> upstream/master
func (wallet *Wallet) On_WalletSendToAddress(req *types.ReqWalletSendToAddress) (types.Message, error) {
	reply, err := wallet.ProcSendToAddress(req)
	if err != nil {
		walletlog.Error("ProcSendToAddress", "err", err.Error())
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_WalletSetFee 响应设置钱包手续费
>>>>>>> upstream/master
func (wallet *Wallet) On_WalletSetFee(req *types.ReqWalletSetFee) (types.Message, error) {
	reply := &types.Reply{
		IsOk: true,
	}
	err := wallet.ProcWalletSetFee(req)
	if err != nil {
		walletlog.Error("ProcWalletSetFee", "err", err.Error())
		reply.IsOk = false
		reply.Msg = []byte(err.Error())
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_WalletSetLabel 处理钱包设置标签
>>>>>>> upstream/master
func (wallet *Wallet) On_WalletSetLabel(req *types.ReqWalletSetLabel) (types.Message, error) {
	reply, err := wallet.ProcWalletSetLabel(req)
	if err != nil {
		walletlog.Error("ProcWalletSetLabel", "err", err.Error())
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_WalletMergeBalance 响应钱包合并金额
>>>>>>> upstream/master
func (wallet *Wallet) On_WalletMergeBalance(req *types.ReqWalletMergeBalance) (types.Message, error) {
	reply, err := wallet.ProcMergeBalance(req)
	if err != nil {
		walletlog.Error("ProcMergeBalance", "err", err.Error())
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_WalletSetPasswd 处理钱包设置密码
>>>>>>> upstream/master
func (wallet *Wallet) On_WalletSetPasswd(req *types.ReqWalletSetPasswd) (types.Message, error) {
	reply := &types.Reply{
		IsOk: true,
	}
	err := wallet.ProcWalletSetPasswd(req)
	if err != nil {
		walletlog.Error("ProcWalletSetPasswd", "err", err.Error())
		reply.IsOk = false
		reply.Msg = []byte(err.Error())
	}
	return reply, nil
}

<<<<<<< HEAD
=======
// On_WalletLock 处理钱包加锁
>>>>>>> upstream/master
func (wallet *Wallet) On_WalletLock(req *types.ReqNil) (types.Message, error) {
	reply := &types.Reply{
		IsOk: true,
	}
	err := wallet.ProcWalletLock()
	if err != nil {
		walletlog.Error("ProcWalletLock", "err", err.Error())
		reply.IsOk = false
		reply.Msg = []byte(err.Error())
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_WalletUnLock 处理钱包解锁
>>>>>>> upstream/master
func (wallet *Wallet) On_WalletUnLock(req *types.WalletUnLock) (types.Message, error) {
	reply := &types.Reply{
		IsOk: true,
	}
	err := wallet.ProcWalletUnLock(req)
	if err != nil {
		walletlog.Error("ProcWalletLock", "err", err.Error())
		reply.IsOk = false
		reply.Msg = []byte(err.Error())
	}
	return reply, nil
}

<<<<<<< HEAD
=======
// On_AddBlock 处理新增区块
>>>>>>> upstream/master
func (wallet *Wallet) On_AddBlock(block *types.BlockDetail) (types.Message, error) {
	wallet.updateLastHeader(block, 1)
	wallet.ProcWalletAddBlock(block)
	return nil, nil
}

<<<<<<< HEAD
=======
// On_DelBlock 处理删除区块
>>>>>>> upstream/master
func (wallet *Wallet) On_DelBlock(block *types.BlockDetail) (types.Message, error) {
	wallet.updateLastHeader(block, -1)
	wallet.ProcWalletDelBlock(block)
	return nil, nil
}

<<<<<<< HEAD
=======
// On_GenSeed 处理创建SEED
>>>>>>> upstream/master
func (wallet *Wallet) On_GenSeed(req *types.GenSeedLang) (types.Message, error) {
	reply, err := wallet.genSeed(req.Lang)
	if err != nil {
		walletlog.Error("genSeed", "err", err.Error())
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_GetSeed 处理获取Seed
>>>>>>> upstream/master
func (wallet *Wallet) On_GetSeed(req *types.GetSeedByPw) (types.Message, error) {
	reply := &types.ReplySeed{}
	seed, err := wallet.getSeed(req.Passwd)
	if err != nil {
		walletlog.Error("getSeed", "err", err.Error())
	} else {
		reply.Seed = seed
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_SaveSeed 处理保存SEED
>>>>>>> upstream/master
func (wallet *Wallet) On_SaveSeed(req *types.SaveSeedByPw) (types.Message, error) {
	reply := &types.Reply{
		IsOk: true,
	}
	ok, err := wallet.saveSeed(req.Passwd, req.Seed)
	if !ok {
		walletlog.Error("[saveSeed]", "err", err.Error())
		reply.IsOk = false
		reply.Msg = []byte(err.Error())
	}
	return reply, nil
}

<<<<<<< HEAD
=======
// On_GetWalletStatus 处理获取钱包状态
>>>>>>> upstream/master
func (wallet *Wallet) On_GetWalletStatus(req *types.ReqNil) (types.Message, error) {
	reply := wallet.GetWalletStatus()
	return reply, nil
}

<<<<<<< HEAD
=======
// On_DumpPrivkey 处理到处私钥
>>>>>>> upstream/master
func (wallet *Wallet) On_DumpPrivkey(req *types.ReqString) (types.Message, error) {
	reply := &types.ReplyString{}
	privkey, err := wallet.ProcDumpPrivkey(req.Data)
	if err != nil {
		walletlog.Error("ProcDumpPrivkey", "err", err.Error())
	} else {
		reply.Data = privkey
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_SignRawTx 处理交易签名
>>>>>>> upstream/master
func (wallet *Wallet) On_SignRawTx(req *types.ReqSignRawTx) (types.Message, error) {
	reply := &types.ReplySignRawTx{}
	txhex, err := wallet.ProcSignRawTx(req)
	if err != nil {
		walletlog.Error("ProcSignRawTx", "err", err.Error())
	} else {
		reply.TxHex = txhex
	}
	return reply, err
}

<<<<<<< HEAD
=======
// On_ErrToFront 错误通知
>>>>>>> upstream/master
func (wallet *Wallet) On_ErrToFront(req *types.ReportErrEvent) (types.Message, error) {
	wallet.setFatalFailure(req)
	return nil, nil
}

<<<<<<< HEAD
// onFatalFailure 定时查询是否有致命性故障产生
=======
// On_FatalFailure 定时查询是否有致命性故障产生
>>>>>>> upstream/master
func (wallet *Wallet) On_FatalFailure(req *types.ReqNil) (types.Message, error) {
	reply := &types.Int32{
		Data: wallet.getFatalFailure(),
	}
	return reply, nil
}

<<<<<<< HEAD
=======
// ExecWallet 执行钱包的功能
>>>>>>> upstream/master
func (wallet *Wallet) ExecWallet(msg *queue.Message) (types.Message, error) {
	if param, ok := msg.Data.(*types.ChainExecutor); ok {
		return wallet.execWallet(param, 0)
	}
	var data []byte
	if msg.Data != nil {
		if d, ok := msg.Data.(types.Message); ok {
			data = types.Encode(d)
		} else {
			return nil, types.ErrInvalidParam
		}
	}
	param := &types.ChainExecutor{
		Driver: "wallet",
		Param:  data,
	}
	return wallet.execWallet(param, msg.Ty)
}

<<<<<<< HEAD
func (wallet *Wallet) execWallet(param *types.ChainExecutor, eventId int64) (reply types.Message, err error) {
	if param.FuncName == "" && eventId > 0 {
		param.FuncName = types.GetEventName(int(eventId))
=======
func (wallet *Wallet) execWallet(param *types.ChainExecutor, eventID int64) (reply types.Message, err error) {
	if param.FuncName == "" && eventID > 0 {
		param.FuncName = types.GetEventName(int(eventID))
>>>>>>> upstream/master
		if len(param.FuncName) <= 5 {
			return nil, types.ErrActionNotSupport
		}
		param.FuncName = param.FuncName[5:]
	}
	var paramIn types.Message
	if param.Param == nil {
		paramIn = &types.ReqNil{}
	} else {
		paramIn, err = wcom.QueryData.Decode(param.Driver, param.FuncName, param.Param)
		if err != nil {
			return nil, err
		}
	}
	//这里不判断类型是否可以调用，直接按照名字调用，如果发生panic，用recover 恢复
	return wcom.QueryData.Call(param.Driver, param.FuncName, paramIn)
}
