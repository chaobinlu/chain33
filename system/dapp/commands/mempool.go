// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commands

import (
<<<<<<< HEAD
	"github.com/spf13/cobra"
	"github.com/33cn/chain33/rpc/jsonclient"
	rpctypes "github.com/33cn/chain33/rpc/types"
	. "github.com/33cn/chain33/system/dapp/commands/types"
)

=======
	"github.com/33cn/chain33/rpc/jsonclient"
	rpctypes "github.com/33cn/chain33/rpc/types"
	"github.com/33cn/chain33/system/dapp/commands/types"
	"github.com/spf13/cobra"
)

// MempoolCmd mempool command
>>>>>>> upstream/master
func MempoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mempool",
		Short: "Mempool management",
		Args:  cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(
		GetMempoolCmd(),
		GetLastMempoolCmd(),
	)

	return cmd
}

<<<<<<< HEAD
// get mempool
=======
// GetMempoolCmd get mempool
>>>>>>> upstream/master
func GetMempoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List mempool txs",
		Run:   listMempoolTxs,
	}
	return cmd
}

func listMempoolTxs(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res rpctypes.ReplyTxList
<<<<<<< HEAD
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.GetMempool", nil, &res)
=======
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.GetMempool", nil, &res)
>>>>>>> upstream/master
	ctx.SetResultCb(parseListMempoolTxsRes)
	ctx.Run()
}

func parseListMempoolTxsRes(arg interface{}) (interface{}, error) {
	res := arg.(*rpctypes.ReplyTxList)
<<<<<<< HEAD
	var result TxListResult
	for _, v := range res.Txs {
		result.Txs = append(result.Txs, DecodeTransaction(v))
=======
	var result types.TxListResult
	for _, v := range res.Txs {
		result.Txs = append(result.Txs, types.DecodeTransaction(v))
>>>>>>> upstream/master
	}
	return result, nil
}

<<<<<<< HEAD
// get last 10 txs of mempool
=======
// GetLastMempoolCmd  get last 10 txs of mempool
>>>>>>> upstream/master
func GetLastMempoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "last_txs",
		Short: "Get latest mempool txs",
		Run:   lastMempoolTxs,
	}
	return cmd
}

func lastMempoolTxs(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res rpctypes.ReplyTxList
<<<<<<< HEAD
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.GetLastMemPool", nil, &res)
=======
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.GetLastMemPool", nil, &res)
>>>>>>> upstream/master
	ctx.SetResultCb(parselastMempoolTxsRes)
	ctx.Run()
}

func parselastMempoolTxsRes(arg interface{}) (interface{}, error) {
	res := arg.(*rpctypes.ReplyTxList)
<<<<<<< HEAD
	var result TxListResult
	for _, v := range res.Txs {
		result.Txs = append(result.Txs, DecodeTransaction(v))
=======
	var result types.TxListResult
	for _, v := range res.Txs {
		result.Txs = append(result.Txs, types.DecodeTransaction(v))
>>>>>>> upstream/master
	}
	return result, nil
}
