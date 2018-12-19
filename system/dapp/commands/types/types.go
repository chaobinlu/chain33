// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// Package types commands中结构体定义
>>>>>>> upstream/master
package types

import (
	rpctypes "github.com/33cn/chain33/rpc/types"
	"github.com/33cn/chain33/types"
)

<<<<<<< HEAD
=======
// AccountsResult defines accountsresult command
>>>>>>> upstream/master
type AccountsResult struct {
	Wallets []*WalletResult `json:"wallets"`
}

<<<<<<< HEAD
=======
// WalletResult defines walletresult command
>>>>>>> upstream/master
type WalletResult struct {
	Acc   *AccountResult `json:"acc,omitempty"`
	Label string         `json:"label,omitempty"`
}

<<<<<<< HEAD
=======
// AccountResult defines account result command
>>>>>>> upstream/master
type AccountResult struct {
	Currency int32  `json:"currency,omitempty"`
	Balance  string `json:"balance,omitempty"`
	Frozen   string `json:"frozen,omitempty"`
	Addr     string `json:"addr,omitempty"`
}

<<<<<<< HEAD
=======
// TokenAccountResult defines accounts result of token command
>>>>>>> upstream/master
type TokenAccountResult struct {
	Token    string `json:"Token,omitempty"`
	Currency int32  `json:"currency,omitempty"`
	Balance  string `json:"balance,omitempty"`
	Frozen   string `json:"frozen,omitempty"`
	Addr     string `json:"addr,omitempty"`
}

<<<<<<< HEAD
=======
// TxListResult defines txlist result command
>>>>>>> upstream/master
type TxListResult struct {
	Txs []*TxResult `json:"txs"`
}

<<<<<<< HEAD
=======
// TxResult defines txresult command
>>>>>>> upstream/master
type TxResult struct {
	Execer     string              `json:"execer"`
	Payload    interface{}         `json:"payload"`
	RawPayload string              `json:"rawpayload"`
	Signature  *rpctypes.Signature `json:"signature"`
	Fee        string              `json:"fee"`
	Expire     int64               `json:"expire"`
	Nonce      int64               `json:"nonce"`
	To         string              `json:"to"`
	Amount     string              `json:"amount,omitempty"`
	From       string              `json:"from,omitempty"`
	GroupCount int32               `json:"groupCount,omitempty"`
	Header     string              `json:"header,omitempty"`
	Next       string              `json:"next,omitempty"`
<<<<<<< HEAD
}

=======
	Hash       string              `json:"hash,omitempty"`
}

// ReceiptAccountTransfer defines receipt account transfer
>>>>>>> upstream/master
type ReceiptAccountTransfer struct {
	Prev    *AccountResult `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	Current *AccountResult `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

<<<<<<< HEAD
=======
// ReceiptExecAccountTransfer defines account transfer of exec command
>>>>>>> upstream/master
type ReceiptExecAccountTransfer struct {
	ExecAddr string         `protobuf:"bytes,1,opt,name=execAddr" json:"execAddr,omitempty"`
	Prev     *AccountResult `protobuf:"bytes,2,opt,name=prev" json:"prev,omitempty"`
	Current  *AccountResult `protobuf:"bytes,3,opt,name=current" json:"current,omitempty"`
}

<<<<<<< HEAD
=======
// TxDetailResult defines txdetail result command
>>>>>>> upstream/master
type TxDetailResult struct {
	Tx         *TxResult                   `json:"tx"`
	Receipt    *rpctypes.ReceiptDataResult `json:"receipt"`
	Proofs     []string                    `json:"proofs,omitempty"`
	Height     int64                       `json:"height"`
	Index      int64                       `json:"index"`
	Blocktime  int64                       `json:"blocktime"`
	Amount     string                      `json:"amount"`
	Fromaddr   string                      `json:"fromaddr"`
	ActionName string                      `json:"actionname"`
	Assets     []*types.Asset              `json:"assets"`
}

<<<<<<< HEAD
=======
// TxDetailsResult defines txdetails result command
>>>>>>> upstream/master
type TxDetailsResult struct {
	Txs []*TxDetailResult `json:"txs"`
}

<<<<<<< HEAD
=======
// BlockResult defines blockresult rpc command
>>>>>>> upstream/master
type BlockResult struct {
	Version    int64       `json:"version"`
	ParentHash string      `json:"parenthash"`
	TxHash     string      `json:"txhash"`
	StateHash  string      `json:"statehash"`
	Height     int64       `json:"height"`
	BlockTime  int64       `json:"blocktime"`
	Txs        []*TxResult `json:"txs"`
}

<<<<<<< HEAD
=======
// BlockDetailResult defines blockdetailresult rpc command
>>>>>>> upstream/master
type BlockDetailResult struct {
	Block    *BlockResult                  `json:"block"`
	Receipts []*rpctypes.ReceiptDataResult `json:"receipts"`
}

<<<<<<< HEAD
=======
// BlockDetailsResult defines blockdetails result rpc command
>>>>>>> upstream/master
type BlockDetailsResult struct {
	Items []*BlockDetailResult `json:"items"`
}

<<<<<<< HEAD
=======
// WalletTxDetailsResult defines walletexdetails result rpc command
>>>>>>> upstream/master
type WalletTxDetailsResult struct {
	TxDetails []*WalletTxDetailResult `json:"txDetails"`
}

<<<<<<< HEAD
=======
// WalletTxDetailResult  defines wallettxdetail result rpc command
>>>>>>> upstream/master
type WalletTxDetailResult struct {
	Tx         *TxResult                   `json:"tx"`
	Receipt    *rpctypes.ReceiptDataResult `json:"receipt"`
	Height     int64                       `json:"height"`
	Index      int64                       `json:"index"`
	Blocktime  int64                       `json:"blocktime"`
	Amount     string                      `json:"amount"`
	Fromaddr   string                      `json:"fromaddr"`
	Txhash     string                      `json:"txhash"`
	ActionName string                      `json:"actionname"`
}

<<<<<<< HEAD
=======
// AddrOverviewResult defines address overview result rpc command
>>>>>>> upstream/master
type AddrOverviewResult struct {
	Receiver string `json:"receiver"`
	Balance  string `json:"balance"`
	TxCount  int64  `json:"txCount"`
}

<<<<<<< HEAD
=======
// GetTotalCoinsResult defines totalcoinsresult rpc command
>>>>>>> upstream/master
type GetTotalCoinsResult struct {
	TxCount          int64  `json:"txCount"`
	AccountCount     int64  `json:"accountCount"`
	TotalAmount      string `json:"totalAmount"`
	ActualAmount     string `json:"actualAmount,omitempty"`
	DifferenceAmount string `json:"differenceAmount,omitempty"`
}

<<<<<<< HEAD
=======
// GetTicketStatisticResult defines ticketstatistic result rpc command
>>>>>>> upstream/master
type GetTicketStatisticResult struct {
	CurrentOpenCount int64 `json:"currentOpenCount"`
	TotalMinerCount  int64 `json:"totalMinerCount"`
	TotalCloseCount  int64 `json:"totalCloseCount"`
}

<<<<<<< HEAD
type GetTicketMinerInfoResult struct {
	TicketId     string `json:"ticketId"`
=======
// GetTicketMinerInfoResult defines ticker minerinformation result rpc command
type GetTicketMinerInfoResult struct {
	TicketID     string `json:"ticketId"`
>>>>>>> upstream/master
	Status       string `json:"status"`
	PrevStatus   string `json:"prevStatus"`
	IsGenesis    bool   `json:"isGenesis"`
	CreateTime   string `json:"createTime"`
	MinerTime    string `json:"minerTime"`
	CloseTime    string `json:"closeTime"`
	MinerValue   int64  `json:"minerValue,omitempty"`
	MinerAddress string `json:"minerAddress,omitempty"`
}

<<<<<<< HEAD
=======
// UTXOGlobalIndex defines  utxo globalindex command
>>>>>>> upstream/master
type UTXOGlobalIndex struct {
	// Height   int64  `json:"height,omitempty"`
	// Txindex  int32  `json:"txindex,omitempty"`
	Outindex int32  `json:"outindex,omitempty"`
	Txhash   string `json:"txhash,omitempty"`
}

<<<<<<< HEAD
=======
// KeyInput defines keyinput info command
>>>>>>> upstream/master
type KeyInput struct {
	Amount          string             `json:"amount,omitempty"`
	UtxoGlobalIndex []*UTXOGlobalIndex `json:"utxoGlobalIndex,omitempty"`
	KeyImage        string             `json:"keyImage,omitempty"`
}

<<<<<<< HEAD
=======
// PrivacyInput defines privacy input command
>>>>>>> upstream/master
type PrivacyInput struct {
	Keyinput []*KeyInput `json:"keyinput,omitempty"`
}

<<<<<<< HEAD
// privacy output
=======
// KeyOutput privacy output
>>>>>>> upstream/master
type KeyOutput struct {
	Amount        string `json:"amount,omitempty"`
	Onetimepubkey string `json:"onetimepubkey,omitempty"`
}

<<<<<<< HEAD
=======
// ReceiptPrivacyOutput defines receipt privacy output command
>>>>>>> upstream/master
type ReceiptPrivacyOutput struct {
	Token     string       `json:"token,omitempty"`
	Keyoutput []*KeyOutput `json:"keyoutput,omitempty"`
}

<<<<<<< HEAD
=======
// AllExecBalance defines all balance of exec command
>>>>>>> upstream/master
type AllExecBalance struct {
	Addr        string         `json:"addr"`
	ExecAccount []*ExecAccount `json:"execAccount"`
}

<<<<<<< HEAD
=======
// ExecAccount defines account of exec command
>>>>>>> upstream/master
type ExecAccount struct {
	Execer  string         `json:"execer"`
	Account *AccountResult `json:"account"`
}

<<<<<<< HEAD
// decodetx
=======
// CoinsTransferCLI  decodetx
>>>>>>> upstream/master
type CoinsTransferCLI struct {
	Cointoken string `json:"cointoken,omitempty"`
	Amount    string `json:"amount,omitempty"`
	Note      string `json:"note,omitempty"`
	To        string `json:"to,omitempty"`
}

<<<<<<< HEAD
=======
// CoinsWithdrawCLI defines coins withdrawcli command
>>>>>>> upstream/master
type CoinsWithdrawCLI struct {
	Cointoken string `json:"cointoken,omitempty"`
	Amount    string `json:"amount,omitempty"`
	Note      string `json:"note,omitempty"`
	ExecName  string `json:"execName,omitempty"`
	To        string `json:"to,omitempty"`
}

<<<<<<< HEAD
=======
// CoinsGenesisCLI defines coins genesis cli command
>>>>>>> upstream/master
type CoinsGenesisCLI struct {
	Amount        string `json:"amount,omitempty"`
	ReturnAddress string `json:"returnAddress,omitempty"`
}

<<<<<<< HEAD
=======
// CoinsTransferToExecCLI defines coins transfertoexec cli command
>>>>>>> upstream/master
type CoinsTransferToExecCLI struct {
	Cointoken string `json:"cointoken,omitempty"`
	Amount    string `json:"amount,omitempty"`
	Note      string `json:"note,omitempty"`
	ExecName  string `json:"execName,omitempty"`
	To        string `json:"to,omitempty"`
}

<<<<<<< HEAD
=======
// HashlockLockCLI defines hashlocklockcli rpc command
>>>>>>> upstream/master
type HashlockLockCLI struct {
	Amount        string `json:"amount,omitempty"`
	Time          int64  `json:"time,omitempty"`
	Hash          []byte `json:"hash,omitempty"`
	ToAddress     string `json:"toAddress,omitempty"`
	ReturnAddress string `json:"returnAddress,omitempty"`
}

<<<<<<< HEAD
type TicketMinerCLI struct {
	Bits     uint32 `json:"bits,omitempty"`
	Reward   string `json:"reward,omitempty"`
	TicketId string `json:"ticketId,omitempty"`
	Modify   []byte `json:"modify,omitempty"`
}

=======
// TicketMinerCLI defines ticket minercli command
type TicketMinerCLI struct {
	Bits     uint32 `json:"bits,omitempty"`
	Reward   string `json:"reward,omitempty"`
	TicketID string `json:"ticketId,omitempty"`
	Modify   []byte `json:"modify,omitempty"`
}

// TokenPreCreateCLI defines token precreatecli command
>>>>>>> upstream/master
type TokenPreCreateCLI struct {
	Name         string `json:"name,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Total        int64  `json:"total,omitempty"`
	Price        string `json:"price,omitempty"`
	Owner        string `json:"owner,omitempty"`
}

<<<<<<< HEAD
=======
// Public2PrivacyCLI defines public to privacy cli command
>>>>>>> upstream/master
type Public2PrivacyCLI struct {
	Tokenname string         `json:"tokenname,omitempty"`
	Amount    string         `json:"amount,omitempty"`
	Note      string         `json:"note,omitempty"`
	Output    *PrivacyOutput `json:"output,omitempty"`
}

<<<<<<< HEAD
=======
// Privacy2PrivacyCLI defines privacy to privacy cli command
>>>>>>> upstream/master
type Privacy2PrivacyCLI struct {
	Tokenname string         `json:"tokenname,omitempty"`
	Amount    string         `json:"amount,omitempty"`
	Note      string         `json:"note,omitempty"`
	Input     *PrivacyInput  `json:"input,omitempty"`
	Output    *PrivacyOutput `json:"output,omitempty"`
}

<<<<<<< HEAD
=======
// Privacy2PublicCLI defines privacy to public cli command
>>>>>>> upstream/master
type Privacy2PublicCLI struct {
	Tokenname string         `json:"tokenname,omitempty"`
	Amount    string         `json:"amount,omitempty"`
	Note      string         `json:"note,omitempty"`
	Input     *PrivacyInput  `json:"input,omitempty"`
	Output    *PrivacyOutput `json:"output,omitempty"`
}

<<<<<<< HEAD
=======
// PrivacyOutput defines Privacy output command
>>>>>>> upstream/master
type PrivacyOutput struct {
	RpubKeytx string       `protobuf:"bytes,1,opt,name=RpubKeytx,proto3" json:"RpubKeytx,omitempty"`
	Keyoutput []*KeyOutput `protobuf:"bytes,2,rep,name=keyoutput" json:"keyoutput,omitempty"`
}
<<<<<<< HEAD
=======

// GetExecBalanceResult  defines balance of exec result rpc command
type GetExecBalanceResult struct {
	Amount       string         `json:"totalAmount"`
	AmountFrozen string         `json:"frozenAmount"`
	AmountActive string         `json:"activeAmount"`
	ExecBalances []*ExecBalance `json:"execBalances,omitempty"`
}

// ExecBalance defines exec balance rpc command
type ExecBalance struct {
	ExecAddr string `json:"execAddr,omitempty"`
	Frozen   string `json:"frozen"`
	Active   string `json:"active"`
}
>>>>>>> upstream/master
