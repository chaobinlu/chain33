// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
package commands

import (
	"github.com/spf13/cobra"
	"github.com/33cn/chain33/rpc/jsonclient"
)

=======
// Package commands 系统级dapp相关命令包
package commands

import (
	"github.com/33cn/chain33/rpc/jsonclient"
	"github.com/33cn/chain33/types"
	"github.com/spf13/cobra"
)

// VersionCmd version command
>>>>>>> upstream/master
func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Get node version",
		Run:   version,
	}

	return cmd
}

func version(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
<<<<<<< HEAD

	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.Version", nil, nil)
	ctx.RunWithoutMarshal()
=======
	var res types.VersionInfo
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.Version", nil, &res)
	ctx.Run()

>>>>>>> upstream/master
}
