// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"fmt"
	"time"
)

<<<<<<< HEAD
=======
// 定义key值
>>>>>>> upstream/master
var (
	LocalPrefix       = []byte("LODB")
	FlagTxQuickIndex  = []byte("FLAG:FlagTxQuickIndex")
	FlagKeyMVCC       = []byte("FLAG:keyMVCCFlag")
	TxHashPerfix      = []byte("TX:")
	TxShortHashPerfix = []byte("STX:")
	TxAddrHash        = []byte("TxAddrHash:")
	TxAddrDirHash     = []byte("TxAddrDirHash:")
	AddrTxsCount      = []byte("AddrTxsCount:")
)

<<<<<<< HEAD
=======
// GetLocalDBKeyList 获取localdb的key列表
>>>>>>> upstream/master
func GetLocalDBKeyList() [][]byte {
	return [][]byte{
		FlagTxQuickIndex, FlagKeyMVCC, TxHashPerfix, TxShortHashPerfix,
	}
}

//CalcTxKey local db中保存交易的方法
func CalcTxKey(hash []byte) []byte {
	if IsEnable("quickIndex") {
		return append(TxHashPerfix, hash...)
	}
	return hash
}

<<<<<<< HEAD
=======
//CalcTxShortKey local db中保存交易的方法
>>>>>>> upstream/master
func CalcTxShortKey(hash []byte) []byte {
	return append(TxShortHashPerfix, hash[0:8]...)
}

<<<<<<< HEAD
//用于存储地址相关的hash列表，key=TxAddrHash:addr:height*100000 + index
=======
//CalcTxAddrHashKey 用于存储地址相关的hash列表，key=TxAddrHash:addr:height*100000 + index
>>>>>>> upstream/master
//地址下面所有的交易
func CalcTxAddrHashKey(addr string, heightindex string) []byte {
	return append(TxAddrHash, []byte(fmt.Sprintf("%s:%s", addr, heightindex))...)
}

<<<<<<< HEAD
//用于存储地址相关的hash列表，key=TxAddrHash:addr:flag:height*100000 + index
=======
//CalcTxAddrDirHashKey 用于存储地址相关的hash列表，key=TxAddrHash:addr:flag:height*100000 + index
>>>>>>> upstream/master
//地址下面某个分类的交易
func CalcTxAddrDirHashKey(addr string, flag int32, heightindex string) []byte {
	return append(TxAddrDirHash, []byte(fmt.Sprintf("%s:%d:%s", addr, flag, heightindex))...)
}

<<<<<<< HEAD
//存储地址参与的交易数量。add时加一，del时减一
=======
//CalcAddrTxsCountKey 存储地址参与的交易数量。add时加一，del时减一
>>>>>>> upstream/master
func CalcAddrTxsCountKey(addr string) []byte {
	return append(AddrTxsCount, []byte(addr)...)
}

<<<<<<< HEAD
=======
//StatisticFlag 用于记录统计的key
>>>>>>> upstream/master
func StatisticFlag() []byte {
	return []byte("Statistics:Flag")
}

<<<<<<< HEAD
func StatisticTicketInfoKey(ticketId string) []byte {
	return []byte("Statistics:TicketInfo:TicketId:" + ticketId)
}

func StatisticTicketInfoOrderKey(minerAddr string, createTime int64, ticketId string) []byte {
	return []byte("Statistics:TicketInfoOrder:Addr:" + minerAddr + ":CreateTime:" + time.Unix(createTime, 0).Format("20060102150405") + ":TicketId:" + ticketId)
}

=======
//StatisticTicketInfoKey 统计ticket的key
func StatisticTicketInfoKey(ticketID string) []byte {
	return []byte("Statistics:TicketInfo:TicketId:" + ticketID)
}

//StatisticTicketInfoOrderKey 统计ticket的key
func StatisticTicketInfoOrderKey(minerAddr string, createTime int64, ticketID string) []byte {
	return []byte("Statistics:TicketInfoOrder:Addr:" + minerAddr + ":CreateTime:" + time.Unix(createTime, 0).Format("20060102150405") + ":TicketId:" + ticketID)
}

//StatisticTicketKey 统计ticket的key
>>>>>>> upstream/master
func StatisticTicketKey(minerAddr string) []byte {
	return []byte("Statistics:TicketStat:Addr:" + minerAddr)
}

<<<<<<< HEAD
=======
//TotalFeeKey 统计所有费用的key
>>>>>>> upstream/master
func TotalFeeKey(hash []byte) []byte {
	key := []byte("TotalFeeKey:")
	return append(key, hash...)
}
