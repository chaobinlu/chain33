// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"encoding/hex"
	"time"

	pb "github.com/33cn/chain33/types"
	"golang.org/x/net/context"
)

<<<<<<< HEAD
=======
// SendTransaction send transaction by network
>>>>>>> upstream/master
func (g *Grpc) SendTransaction(ctx context.Context, in *pb.Transaction) (*pb.Reply, error) {
	return g.cli.SendTx(in)
}

<<<<<<< HEAD
=======
// CreateNoBalanceTransaction create transaction with no balance
>>>>>>> upstream/master
func (g *Grpc) CreateNoBalanceTransaction(ctx context.Context, in *pb.NoBalanceTx) (*pb.ReplySignRawTx, error) {
	reply, err := g.cli.CreateNoBalanceTransaction(in)
	if err != nil {
		return nil, err
	}
	tx := pb.Encode(reply)
	return &pb.ReplySignRawTx{TxHex: hex.EncodeToString(tx)}, nil
}

<<<<<<< HEAD
=======
// CreateRawTransaction create rawtransaction of grpc
>>>>>>> upstream/master
func (g *Grpc) CreateRawTransaction(ctx context.Context, in *pb.CreateTx) (*pb.UnsignTx, error) {
	reply, err := g.cli.CreateRawTransaction(in)
	if err != nil {
		return nil, err
	}
	return &pb.UnsignTx{Data: reply}, nil
}

<<<<<<< HEAD
func (g *Grpc) CreateTransaction(ctx context.Context, in *pb.CreateTxIn) (*pb.UnsignTx, error) {
	exec := pb.LoadExecutorType(string(in.Execer))
=======
// CreateTransaction create transaction of grpc
func (g *Grpc) CreateTransaction(ctx context.Context, in *pb.CreateTxIn) (*pb.UnsignTx, error) {
	execer := pb.ExecName(string(in.Execer))
	exec := pb.LoadExecutorType(execer)
>>>>>>> upstream/master
	if exec == nil {
		log.Error("callExecNewTx", "Error", "exec not found")
		return nil, pb.ErrNotSupport
	}
	msg, err := exec.GetAction(in.ActionName)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	reply, err := pb.CallCreateTx(string(in.Execer), in.ActionName, msg)
=======
	//decode protocol buffer
	err = pb.Decode(in.Payload, msg)
	if err != nil {
		return nil, err
	}
	reply, err := pb.CallCreateTx(execer, in.ActionName, msg)
>>>>>>> upstream/master
	if err != nil {
		return nil, err
	}
	return &pb.UnsignTx{Data: reply}, nil
}

<<<<<<< HEAD
=======
// CreateRawTxGroup create rawtransaction for group
>>>>>>> upstream/master
func (g *Grpc) CreateRawTxGroup(ctx context.Context, in *pb.CreateTransactionGroup) (*pb.UnsignTx, error) {
	reply, err := g.cli.CreateRawTxGroup(in)
	if err != nil {
		return nil, err
	}
	return &pb.UnsignTx{Data: reply}, nil
}

<<<<<<< HEAD
=======
// SendRawTransaction send rawtransaction
>>>>>>> upstream/master
func (g *Grpc) SendRawTransaction(ctx context.Context, in *pb.SignedTx) (*pb.Reply, error) {
	return g.cli.SendRawTransaction(in)
}

<<<<<<< HEAD
=======
// QueryTransaction query transaction by grpc
>>>>>>> upstream/master
func (g *Grpc) QueryTransaction(ctx context.Context, in *pb.ReqHash) (*pb.TransactionDetail, error) {
	return g.cli.QueryTx(in)
}

<<<<<<< HEAD
func (g *Grpc) GetBlocks(ctx context.Context, in *pb.ReqBlocks) (*pb.Reply, error) {
	reply, err := g.cli.GetBlocks(&pb.ReqBlocks{in.Start, in.End, in.IsDetail, []string{""}})
=======
// GetBlocks get blocks by grpc
func (g *Grpc) GetBlocks(ctx context.Context, in *pb.ReqBlocks) (*pb.Reply, error) {
	reply, err := g.cli.GetBlocks(&pb.ReqBlocks{
		Start:    in.Start,
		End:      in.End,
		IsDetail: in.IsDetail,
	})
>>>>>>> upstream/master
	if err != nil {
		return nil, err
	}
	return &pb.Reply{
			IsOk: true,
			Msg:  pb.Encode(reply)},
		nil
}

<<<<<<< HEAD
=======
// GetLastHeader get lastheader information
>>>>>>> upstream/master
func (g *Grpc) GetLastHeader(ctx context.Context, in *pb.ReqNil) (*pb.Header, error) {
	return g.cli.GetLastHeader()
}

<<<<<<< HEAD
=======
// GetTransactionByAddr get transaction by address
>>>>>>> upstream/master
func (g *Grpc) GetTransactionByAddr(ctx context.Context, in *pb.ReqAddr) (*pb.ReplyTxInfos, error) {
	return g.cli.GetTransactionByAddr(in)
}

<<<<<<< HEAD
=======
// GetHexTxByHash get hex transaction by hash
>>>>>>> upstream/master
func (g *Grpc) GetHexTxByHash(ctx context.Context, in *pb.ReqHash) (*pb.HexTx, error) {
	reply, err := g.cli.QueryTx(in)
	if err != nil {
		return nil, err
	}
	tx := reply.GetTx()
	if tx == nil {
		return &pb.HexTx{}, nil
	}
	return &pb.HexTx{Tx: hex.EncodeToString(pb.Encode(reply.GetTx()))}, nil
}

<<<<<<< HEAD
=======
// GetTransactionByHashes get transaction by hashes
>>>>>>> upstream/master
func (g *Grpc) GetTransactionByHashes(ctx context.Context, in *pb.ReqHashes) (*pb.TransactionDetails, error) {
	return g.cli.GetTransactionByHash(in)
}

<<<<<<< HEAD
=======
// GetMemPool get mempool contents
>>>>>>> upstream/master
func (g *Grpc) GetMemPool(ctx context.Context, in *pb.ReqNil) (*pb.ReplyTxList, error) {
	return g.cli.GetMempool()
}

<<<<<<< HEAD
=======
// GetAccounts get  accounts
>>>>>>> upstream/master
func (g *Grpc) GetAccounts(ctx context.Context, in *pb.ReqNil) (*pb.WalletAccounts, error) {
	req := &pb.ReqAccountList{WithoutBalance: false}
	return g.cli.WalletGetAccountList(req)
}

<<<<<<< HEAD
=======
// NewAccount produce new account
>>>>>>> upstream/master
func (g *Grpc) NewAccount(ctx context.Context, in *pb.ReqNewAccount) (*pb.WalletAccount, error) {
	return g.cli.NewAccount(in)
}

<<<<<<< HEAD
=======
// WalletTransactionList transaction list of wallet
>>>>>>> upstream/master
func (g *Grpc) WalletTransactionList(ctx context.Context, in *pb.ReqWalletTransactionList) (*pb.WalletTxDetails, error) {
	return g.cli.WalletTransactionList(in)
}

<<<<<<< HEAD
=======
// ImportPrivkey import privkey
>>>>>>> upstream/master
func (g *Grpc) ImportPrivkey(ctx context.Context, in *pb.ReqWalletImportPrivkey) (*pb.WalletAccount, error) {
	return g.cli.WalletImportprivkey(in)
}

<<<<<<< HEAD
=======
// SendToAddress send to address of coins
>>>>>>> upstream/master
func (g *Grpc) SendToAddress(ctx context.Context, in *pb.ReqWalletSendToAddress) (*pb.ReplyHash, error) {
	return g.cli.WalletSendToAddress(in)
}

<<<<<<< HEAD
=======
// SetTxFee set tx fee
>>>>>>> upstream/master
func (g *Grpc) SetTxFee(ctx context.Context, in *pb.ReqWalletSetFee) (*pb.Reply, error) {
	return g.cli.WalletSetFee(in)
}

<<<<<<< HEAD
=======
// SetLabl set labl
>>>>>>> upstream/master
func (g *Grpc) SetLabl(ctx context.Context, in *pb.ReqWalletSetLabel) (*pb.WalletAccount, error) {
	return g.cli.WalletSetLabel(in)
}

<<<<<<< HEAD
=======
// MergeBalance merge balance of wallet
>>>>>>> upstream/master
func (g *Grpc) MergeBalance(ctx context.Context, in *pb.ReqWalletMergeBalance) (*pb.ReplyHashes, error) {
	return g.cli.WalletMergeBalance(in)
}

<<<<<<< HEAD
=======
// SetPasswd set password
>>>>>>> upstream/master
func (g *Grpc) SetPasswd(ctx context.Context, in *pb.ReqWalletSetPasswd) (*pb.Reply, error) {
	return g.cli.WalletSetPasswd(in)
}

<<<<<<< HEAD
=======
// Lock wallet lock
>>>>>>> upstream/master
func (g *Grpc) Lock(ctx context.Context, in *pb.ReqNil) (*pb.Reply, error) {
	return g.cli.WalletLock()
}

<<<<<<< HEAD
=======
// UnLock wallet unlock
>>>>>>> upstream/master
func (g *Grpc) UnLock(ctx context.Context, in *pb.WalletUnLock) (*pb.Reply, error) {
	return g.cli.WalletUnLock(in)
}

<<<<<<< HEAD
=======
// GetPeerInfo get peer information
>>>>>>> upstream/master
func (g *Grpc) GetPeerInfo(ctx context.Context, in *pb.ReqNil) (*pb.PeerList, error) {
	return g.cli.PeerInfo()
}

<<<<<<< HEAD
=======
// GetHeaders return headers
>>>>>>> upstream/master
func (g *Grpc) GetHeaders(ctx context.Context, in *pb.ReqBlocks) (*pb.Headers, error) {
	return g.cli.GetHeaders(in)
}

<<<<<<< HEAD
=======
// GetLastMemPool return last mempool contents
>>>>>>> upstream/master
func (g *Grpc) GetLastMemPool(ctx context.Context, in *pb.ReqNil) (*pb.ReplyTxList, error) {
	return g.cli.GetLastMempool()
}

<<<<<<< HEAD
//add by hyb
//GetBlockOverview(parm *types.ReqHash) (*types.BlockOverview, error)
=======
// GetBlockOverview get block overview
// GetBlockOverview(parm *types.ReqHash) (*types.BlockOverview, error)   //add by hyb
>>>>>>> upstream/master
func (g *Grpc) GetBlockOverview(ctx context.Context, in *pb.ReqHash) (*pb.BlockOverview, error) {
	return g.cli.GetBlockOverview(in)
}

<<<<<<< HEAD
=======
// GetAddrOverview get address overview
>>>>>>> upstream/master
func (g *Grpc) GetAddrOverview(ctx context.Context, in *pb.ReqAddr) (*pb.AddrOverview, error) {
	return g.cli.GetAddrOverview(in)
}

<<<<<<< HEAD
=======
// GetBlockHash get block  hash
>>>>>>> upstream/master
func (g *Grpc) GetBlockHash(ctx context.Context, in *pb.ReqInt) (*pb.ReplyHash, error) {
	return g.cli.GetBlockHash(in)
}

<<<<<<< HEAD
//seed
=======
// GenSeed seed
>>>>>>> upstream/master
func (g *Grpc) GenSeed(ctx context.Context, in *pb.GenSeedLang) (*pb.ReplySeed, error) {
	return g.cli.GenSeed(in)
}

<<<<<<< HEAD
=======
// GetSeed get seed
>>>>>>> upstream/master
func (g *Grpc) GetSeed(ctx context.Context, in *pb.GetSeedByPw) (*pb.ReplySeed, error) {
	return g.cli.GetSeed(in)
}

<<<<<<< HEAD
=======
// SaveSeed save seed
>>>>>>> upstream/master
func (g *Grpc) SaveSeed(ctx context.Context, in *pb.SaveSeedByPw) (*pb.Reply, error) {
	return g.cli.SaveSeed(in)
}

<<<<<<< HEAD
=======
// GetWalletStatus get wallet status
>>>>>>> upstream/master
func (g *Grpc) GetWalletStatus(ctx context.Context, in *pb.ReqNil) (*pb.WalletStatus, error) {
	return g.cli.GetWalletStatus()
}

<<<<<<< HEAD
=======
// GetBalance get balance
>>>>>>> upstream/master
func (g *Grpc) GetBalance(ctx context.Context, in *pb.ReqBalance) (*pb.Accounts, error) {
	reply, err := g.cli.GetBalance(in)
	if err != nil {
		return nil, err
	}
	return &pb.Accounts{Acc: reply}, nil
}

<<<<<<< HEAD
=======
// GetAllExecBalance get balance of exec
>>>>>>> upstream/master
func (g *Grpc) GetAllExecBalance(ctx context.Context, in *pb.ReqAddr) (*pb.AllExecBalance, error) {
	return g.cli.GetAllExecBalance(in)
}

<<<<<<< HEAD
=======
// QueryConsensus query consensus
>>>>>>> upstream/master
func (g *Grpc) QueryConsensus(ctx context.Context, in *pb.ChainExecutor) (*pb.Reply, error) {
	msg, err := g.cli.QueryConsensus(in)
	if err != nil {
		return nil, err
	}
	var reply pb.Reply
	reply.IsOk = true
	reply.Msg = pb.Encode(msg)
	return &reply, nil
}

<<<<<<< HEAD
=======
// QueryChain query chain
>>>>>>> upstream/master
func (g *Grpc) QueryChain(ctx context.Context, in *pb.ChainExecutor) (*pb.Reply, error) {
	msg, err := g.cli.QueryChain(in)
	if err != nil {
		return nil, err
	}
	var reply pb.Reply
	reply.IsOk = true
	reply.Msg = pb.Encode(msg)
	return &reply, nil
}

<<<<<<< HEAD
=======
// ExecWallet  exec wallet
>>>>>>> upstream/master
func (g *Grpc) ExecWallet(ctx context.Context, in *pb.ChainExecutor) (*pb.Reply, error) {
	msg, err := g.cli.ExecWallet(in)
	if err != nil {
		return nil, err
	}
	var reply pb.Reply
	reply.IsOk = true
	reply.Msg = pb.Encode(msg)
	return &reply, nil
}

<<<<<<< HEAD
=======
// DumpPrivkey dump Privkey
>>>>>>> upstream/master
func (g *Grpc) DumpPrivkey(ctx context.Context, in *pb.ReqString) (*pb.ReplyString, error) {

	return g.cli.DumpPrivkey(in)
}

<<<<<<< HEAD
func (g *Grpc) Version(ctx context.Context, in *pb.ReqNil) (*pb.Reply, error) {
=======
// Version version
func (g *Grpc) Version(ctx context.Context, in *pb.ReqNil) (*pb.VersionInfo, error) {
>>>>>>> upstream/master

	return g.cli.Version()
}

<<<<<<< HEAD
=======
// IsSync is the sync
>>>>>>> upstream/master
func (g *Grpc) IsSync(ctx context.Context, in *pb.ReqNil) (*pb.Reply, error) {

	return g.cli.IsSync()
}

<<<<<<< HEAD
=======
// IsNtpClockSync is ntp clock sync
>>>>>>> upstream/master
func (g *Grpc) IsNtpClockSync(ctx context.Context, in *pb.ReqNil) (*pb.Reply, error) {

	return g.cli.IsNtpClockSync()
}

<<<<<<< HEAD
=======
// NetInfo net information
>>>>>>> upstream/master
func (g *Grpc) NetInfo(ctx context.Context, in *pb.ReqNil) (*pb.NodeNetInfo, error) {

	return g.cli.GetNetInfo()
}

<<<<<<< HEAD
=======
// GetFatalFailure return  fatal of failure
>>>>>>> upstream/master
func (g *Grpc) GetFatalFailure(ctx context.Context, in *pb.ReqNil) (*pb.Int32, error) {
	return g.cli.GetFatalFailure()
}

<<<<<<< HEAD
=======
// CloseQueue close queue
>>>>>>> upstream/master
func (g *Grpc) CloseQueue(ctx context.Context, in *pb.ReqNil) (*pb.Reply, error) {
	go func() {
		time.Sleep(time.Millisecond * 100)
		g.cli.CloseQueue()
	}()

	return &pb.Reply{IsOk: true}, nil
}

<<<<<<< HEAD
func (g *Grpc) GetLastBlockSequence(ctx context.Context, in *pb.ReqNil) (*pb.Int64, error) {
	return g.cli.GetLastBlockSequence()
}
func (g *Grpc) GetBlockSequences(ctx context.Context, in *pb.ReqBlocks) (*pb.BlockSequences, error) {
	return g.cli.GetBlockSequences(in)
}
=======
// GetLastBlockSequence get last block sequence
func (g *Grpc) GetLastBlockSequence(ctx context.Context, in *pb.ReqNil) (*pb.Int64, error) {
	return g.cli.GetLastBlockSequence()
}

// GetBlockSequences get block sequeces
func (g *Grpc) GetBlockSequences(ctx context.Context, in *pb.ReqBlocks) (*pb.BlockSequences, error) {
	return g.cli.GetBlockSequences(in)
}

// GetBlockByHashes get block by hashes
>>>>>>> upstream/master
func (g *Grpc) GetBlockByHashes(ctx context.Context, in *pb.ReqHashes) (*pb.BlockDetails, error) {
	return g.cli.GetBlockByHashes(in)
}

<<<<<<< HEAD
func (g *Grpc) SignRawTx(ctx context.Context, in *pb.ReqSignRawTx) (*pb.ReplySignRawTx, error) {
	return g.cli.SignRawTx(in)
}
=======
// GetSequenceByHash get block sequece by hash
func (g *Grpc) GetSequenceByHash(ctx context.Context, in *pb.ReqHash) (*pb.Int64, error) {
	return g.cli.GetSequenceByHash(in)
}

// SignRawTx signature rawtransaction
func (g *Grpc) SignRawTx(ctx context.Context, in *pb.ReqSignRawTx) (*pb.ReplySignRawTx, error) {
	return g.cli.SignRawTx(in)
}

// QueryRandNum query randHash from ticket
func (g *Grpc) QueryRandNum(ctx context.Context, in *pb.ReqRandHash) (*pb.ReplyHash, error) {
	reply, err := g.cli.Query(in.ExecName, "RandNumHash", in)
	if err != nil {
		return nil, err
	}
	return reply.(*pb.ReplyHash), nil
}
>>>>>>> upstream/master
