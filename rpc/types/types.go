// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package types rpc相关的一些结构体定义以及转化函数
>>>>>>> upstream/master
package types

import (
	"encoding/json"

	"github.com/33cn/chain33/types"
)

<<<<<<< HEAD
=======
// TransParm transport parameter
>>>>>>> upstream/master
type TransParm struct {
	Execer    string     `json:"execer"`
	Payload   string     `json:"payload"`
	Signature *Signature `json:"signature"`
	Fee       int64      `json:"fee"`
}

<<<<<<< HEAD
=======
// SignedTx signature tx
>>>>>>> upstream/master
type SignedTx struct {
	Unsign string `json:"unsignTx"`
	Sign   string `json:"sign"`
	Pubkey string `json:"pubkey"`
	Ty     int32  `json:"ty"`
}

<<<<<<< HEAD
=======
// RawParm defines raw parameter command
>>>>>>> upstream/master
type RawParm struct {
	Token string `json:"token"`
	Data  string `json:"data"`
}

<<<<<<< HEAD
=======
// QueryParm Query parameter
>>>>>>> upstream/master
type QueryParm struct {
	Hash string `json:"hash"`
}

<<<<<<< HEAD
=======
// BlockParam block parameter
>>>>>>> upstream/master
type BlockParam struct {
	Start    int64 `json:"start"`
	End      int64 `json:"end"`
	Isdetail bool  `json:"isDetail"`
}

<<<<<<< HEAD
=======
// Header header parameter
>>>>>>> upstream/master
type Header struct {
	Version    int64      `json:"version"`
	ParentHash string     `json:"parentHash"`
	TxHash     string     `json:"txHash"`
	StateHash  string     `json:"stateHash"`
	Height     int64      `json:"height"`
	BlockTime  int64      `json:"blockTime"`
	TxCount    int64      `json:"txCount"`
	Hash       string     `json:"hash"`
	Difficulty uint32     `json:"difficulty"`
	Signature  *Signature `json:"signature,omitempty"`
}

<<<<<<< HEAD
=======
// Signature parameter
>>>>>>> upstream/master
type Signature struct {
	Ty        int32  `json:"ty"`
	Pubkey    string `json:"pubkey"`
	Signature string `json:"signature"`
}

<<<<<<< HEAD
=======
// Transaction parameter
>>>>>>> upstream/master
type Transaction struct {
	Execer     string          `json:"execer"`
	Payload    json.RawMessage `json:"payload"`
	RawPayload string          `json:"rawPayload"`
	Signature  *Signature      `json:"signature"`
	Fee        int64           `json:"fee"`
	FeeFmt     string          `json:"feefmt"`
	Expire     int64           `json:"expire"`
	Nonce      int64           `json:"nonce"`
	From       string          `json:"from,omitempty"`
	To         string          `json:"to"`
	Amount     int64           `json:"amount,omitempty"`
	AmountFmt  string          `json:"amountfmt,omitempty"`
	GroupCount int32           `json:"groupCount,omitempty"`
	Header     string          `json:"header,omitempty"`
	Next       string          `json:"next,omitempty"`
<<<<<<< HEAD
}

=======
	Hash       string          `json:"hash,omitempty"`
}

// ReceiptLog defines receipt log command
>>>>>>> upstream/master
type ReceiptLog struct {
	Ty  int32  `json:"ty"`
	Log string `json:"log"`
}

<<<<<<< HEAD
=======
// ReceiptData defines receipt data rpc command
>>>>>>> upstream/master
type ReceiptData struct {
	Ty   int32         `json:"ty"`
	Logs []*ReceiptLog `json:"logs"`
}

<<<<<<< HEAD
=======
// ReceiptDataResult receipt data result
>>>>>>> upstream/master
type ReceiptDataResult struct {
	Ty     int32               `json:"ty"`
	TyName string              `json:"tyName"`
	Logs   []*ReceiptLogResult `json:"logs"`
}

<<<<<<< HEAD
=======
// ReceiptLogResult receipt log result
>>>>>>> upstream/master
type ReceiptLogResult struct {
	Ty     int32           `json:"ty"`
	TyName string          `json:"tyName"`
	Log    json.RawMessage `json:"log"`
	RawLog string          `json:"rawLog"`
}

<<<<<<< HEAD
=======
// Block block information
>>>>>>> upstream/master
type Block struct {
	Version    int64          `json:"version"`
	ParentHash string         `json:"parentHash"`
	TxHash     string         `json:"txHash"`
	StateHash  string         `json:"stateHash"`
	Height     int64          `json:"height"`
	BlockTime  int64          `json:"blockTime"`
	Txs        []*Transaction `json:"txs"`
}

<<<<<<< HEAD
=======
// BlockDetail  block detail
>>>>>>> upstream/master
type BlockDetail struct {
	Block    *Block               `json:"block"`
	Receipts []*ReceiptDataResult `json:"recipts"`
}

<<<<<<< HEAD
=======
// BlockDetails block details
>>>>>>> upstream/master
type BlockDetails struct {
	Items []*BlockDetail `json:"items"`
}

<<<<<<< HEAD
=======
// TransactionDetail transaction detail
>>>>>>> upstream/master
type TransactionDetail struct {
	Tx         *Transaction       `json:"tx"`
	Receipt    *ReceiptDataResult `json:"receipt"`
	Proofs     []string           `json:"proofs"`
	Height     int64              `json:"height"`
	Index      int64              `json:"index"`
	Blocktime  int64              `json:"blockTime"`
	Amount     int64              `json:"amount"`
	Fromaddr   string             `json:"fromAddr"`
	ActionName string             `json:"actionName"`
	Assets     []*types.Asset     `json:"assets"`
}

<<<<<<< HEAD
=======
// ReplyTxInfos reply tx infos
>>>>>>> upstream/master
type ReplyTxInfos struct {
	TxInfos []*ReplyTxInfo `json:"txInfos"`
}

<<<<<<< HEAD
=======
// ReplyTxInfo reply tx information
>>>>>>> upstream/master
type ReplyTxInfo struct {
	Hash   string         `json:"hash"`
	Height int64          `json:"height"`
	Index  int64          `json:"index"`
	Assets []*types.Asset `json:"assets"`
}

<<<<<<< HEAD
=======
// TransactionDetails transaction details
>>>>>>> upstream/master
type TransactionDetails struct {
	//Txs []*Transaction `json:"txs"`
	Txs []*TransactionDetail `json:"txs"`
}

<<<<<<< HEAD
=======
// ReplyTxList reply tx list
>>>>>>> upstream/master
type ReplyTxList struct {
	Txs []*Transaction `json:"txs"`
}

<<<<<<< HEAD
=======
// ReplyHash reply hash string json
>>>>>>> upstream/master
type ReplyHash struct {
	Hash string `json:"hash"`
}

<<<<<<< HEAD
type ReplyHashes struct {
	Hashes []string `json:"hashes"`
}
type PeerList struct {
	Peers []*Peer `json:"peers"`
}
=======
// ReplyHashes reply hashes
type ReplyHashes struct {
	Hashes []string `json:"hashes"`
}

// PeerList peer list
type PeerList struct {
	Peers []*Peer `json:"peers"`
}

// Peer  information
>>>>>>> upstream/master
type Peer struct {
	Addr        string  `json:"addr"`
	Port        int32   `json:"port"`
	Name        string  `json:"name"`
	MempoolSize int32   `json:"mempoolSize"`
	Self        bool    `json:"self"`
	Header      *Header `json:"header"`
}

<<<<<<< HEAD
// Wallet Module
type WalletAccounts struct {
	Wallets []*WalletAccount `json:"wallets"`
}
=======
// WalletAccounts Wallet Module
type WalletAccounts struct {
	Wallets []*WalletAccount `json:"wallets"`
}

// WalletAccount  wallet account
>>>>>>> upstream/master
type WalletAccount struct {
	Acc   *Account `json:"acc"`
	Label string   `json:"label"`
}

<<<<<<< HEAD
=======
// Account account information
>>>>>>> upstream/master
type Account struct {
	Currency int32  `json:"currency"`
	Balance  int64  `json:"balance"`
	Frozen   int64  `json:"frozen"`
	Addr     string `json:"addr"`
}
<<<<<<< HEAD
=======

// Reply info
>>>>>>> upstream/master
type Reply struct {
	IsOk bool   `json:"isOK"`
	Msg  string `json:"msg"`
}
<<<<<<< HEAD
=======

// Headers defines headers rpc command
>>>>>>> upstream/master
type Headers struct {
	Items []*Header `json:"items"`
}

<<<<<<< HEAD
=======
// ReqAddr require address
>>>>>>> upstream/master
type ReqAddr struct {
	Addr string `json:"addr"`
}

<<<<<<< HEAD
=======
// ReqHashes require hashes
>>>>>>> upstream/master
type ReqHashes struct {
	Hashes        []string `json:"hashes"`
	DisableDetail bool     `json:"disableDetail"`
}

<<<<<<< HEAD
=======
// ReqWalletTransactionList require wallet transaction list
>>>>>>> upstream/master
type ReqWalletTransactionList struct {
	FromTx          string `json:"fromTx"`
	Count           int32  `json:"count"`
	Direction       int32  `json:"direction"`
	Mode            int32  `json:"mode,omitempty"`
	SendRecvPrivacy int32  `json:"sendRecvPrivacy,omitempty"`
	Address         string `json:"address,omitempty"`
	TokenName       string `json:"tokenname,omitempty"`
}

<<<<<<< HEAD
=======
// WalletTxDetails wallet tx details
>>>>>>> upstream/master
type WalletTxDetails struct {
	TxDetails []*WalletTxDetail `json:"txDetails"`
}

<<<<<<< HEAD
=======
// WalletTxDetail wallet tx detail
>>>>>>> upstream/master
type WalletTxDetail struct {
	Tx         *Transaction       `json:"tx"`
	Receipt    *ReceiptDataResult `json:"receipt"`
	Height     int64              `json:"height"`
	Index      int64              `json:"index"`
	BlockTime  int64              `json:"blockTime"`
	Amount     int64              `json:"amount"`
	FromAddr   string             `json:"fromAddr"`
	TxHash     string             `json:"txHash"`
	ActionName string             `json:"actionName"`
}

<<<<<<< HEAD
=======
// BlockOverview block overview
>>>>>>> upstream/master
type BlockOverview struct {
	Head     *Header  `json:"head"`
	TxCount  int64    `json:"txCount"`
	TxHashes []string `json:"txHashes"`
}

<<<<<<< HEAD
=======
// Query4Jrpc query jrpc
>>>>>>> upstream/master
type Query4Jrpc struct {
	Execer   string          `json:"execer"`
	FuncName string          `json:"funcName"`
	Payload  json.RawMessage `json:"payload"`
}

<<<<<<< HEAD
=======
// ChainExecutor chain executor
>>>>>>> upstream/master
type ChainExecutor struct {
	Driver    string          `json:"execer"`
	FuncName  string          `json:"funcName"`
	StateHash string          `json:"stateHash"`
	Payload   json.RawMessage `json:"payload"`
}

<<<<<<< HEAD
=======
// WalletStatus wallet status
>>>>>>> upstream/master
type WalletStatus struct {
	IsWalletLock bool `json:"isWalletLock"`
	IsAutoMining bool `json:"isAutoMining"`
	IsHasSeed    bool `json:"isHasSeed"`
	IsTicketLock bool `json:"isTicketLock"`
}

<<<<<<< HEAD
=======
// NodeNetinfo node net info
>>>>>>> upstream/master
type NodeNetinfo struct {
	Externaladdr string `json:"externalAddr"`
	Localaddr    string `json:"localAddr"`
	Service      bool   `json:"service"`
	Outbounds    int32  `json:"outbounds"`
	Inbounds     int32  `json:"inbounds"`
}

<<<<<<< HEAD
=======
// ReplyPrivacyPkPair   reply privekey pubkey pair
>>>>>>> upstream/master
type ReplyPrivacyPkPair struct {
	ShowSuccessful bool   `json:"showSuccessful,omitempty"`
	ViewPub        string `json:"viewPub,omitempty"`
	SpendPub       string `json:"spendPub,omitempty"`
}

<<<<<<< HEAD
=======
// ReplyCacheTxList reply cache tx list
>>>>>>> upstream/master
type ReplyCacheTxList struct {
	Txs []*Transaction `json:"txs,omitempty"`
}

<<<<<<< HEAD
=======
// TimeStatus time status
>>>>>>> upstream/master
type TimeStatus struct {
	NtpTime   string `json:"ntpTime"`
	LocalTime string `json:"localTime"`
	Diff      int64  `json:"diff"`
}
<<<<<<< HEAD
=======

// ReplyBlkSeqs reply block sequences
>>>>>>> upstream/master
type ReplyBlkSeqs struct {
	BlkSeqInfos []*ReplyBlkSeq `json:"blkseqInfos"`
}

<<<<<<< HEAD
=======
// ReplyBlkSeq reply block sequece
>>>>>>> upstream/master
type ReplyBlkSeq struct {
	Hash string `json:"hash"`
	Type int64  `json:"type"`
}

<<<<<<< HEAD
=======
// CreateTxIn create tx input
>>>>>>> upstream/master
type CreateTxIn struct {
	Execer     string          `json:"execer"`
	ActionName string          `json:"actionName"`
	Payload    json.RawMessage `json:"payload"`
}

<<<<<<< HEAD
=======
// AllExecBalance all exec balance
>>>>>>> upstream/master
type AllExecBalance struct {
	Addr        string         `json:"addr"`
	ExecAccount []*ExecAccount `json:"execAccount"`
}

<<<<<<< HEAD
=======
// ExecAccount exec account
>>>>>>> upstream/master
type ExecAccount struct {
	Execer  string   `json:"execer"`
	Account *Account `json:"account"`
}

<<<<<<< HEAD
type ExecNameParm struct {
	ExecName string `json:"execname"`
}
=======
// ExecNameParm exec name parameter
type ExecNameParm struct {
	ExecName string `json:"execname"`
}

//CreateTx 为了简化Note 的创建过程，在json rpc 中，note 采用string 格式
type CreateTx struct {
	To          string `json:"to,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Fee         int64  `json:"fee,omitempty"`
	Note        string `json:"note,omitempty"`
	IsWithdraw  bool   `json:"isWithdraw,omitempty"`
	IsToken     bool   `json:"isToken,omitempty"`
	TokenSymbol string `json:"tokenSymbol,omitempty"`
	ExecName    string `json:"execName,omitempty"` //TransferToExec and Withdraw 的执行器
	Execer      string `json:"execer,omitempty"`   //执行器名称
}
>>>>>>> upstream/master
