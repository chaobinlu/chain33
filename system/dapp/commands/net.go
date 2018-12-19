// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commands

import (
	"github.com/spf13/cobra"

	"github.com/33cn/chain33/rpc/jsonclient"
	rpctypes "github.com/33cn/chain33/rpc/types"
)

<<<<<<< HEAD
=======
// NetCmd net command
>>>>>>> upstream/master
func NetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "net",
		Short: "Net operation",
		Args:  cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(
		GetPeerInfoCmd(),
		IsClockSyncCmd(),
		IsSyncCmd(),
		GetNetInfoCmd(),
		GetFatalFailureCmd(),
		GetTimeStausCmd(),
	)

	return cmd
}

<<<<<<< HEAD
// get peers connected info
=======
// GetPeerInfoCmd  get peers connected info
>>>>>>> upstream/master
func GetPeerInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "peer_info",
		Short: "Get remote peer nodes",
		Run:   peerInfo,
	}
	return cmd
}

func peerInfo(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res rpctypes.PeerList
<<<<<<< HEAD
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.GetPeerInfo", nil, &res)
	ctx.Run()
}

// get ntp clock sync status
=======
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.GetPeerInfo", nil, &res)
	ctx.Run()
}

// IsClockSyncCmd  get ntp clock sync status
>>>>>>> upstream/master
func IsClockSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is_clock_sync",
		Short: "Get ntp clock synchronization status",
		Run:   isClockSync,
	}
	return cmd
}

func isClockSync(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res bool
<<<<<<< HEAD
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.IsNtpClockSync", nil, &res)
	ctx.Run()
}

// get local db sync status
=======
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.IsNtpClockSync", nil, &res)
	ctx.Run()
}

// IsSyncCmd get local db sync status
>>>>>>> upstream/master
func IsSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is_sync",
		Short: "Get blockchain synchronization status",
		Run:   isSync,
	}
	return cmd
}

func isSync(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res bool
<<<<<<< HEAD
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.IsSync", nil, &res)
	ctx.Run()
}

// get net info
=======
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.IsSync", nil, &res)
	ctx.Run()
}

// GetNetInfoCmd get net info
>>>>>>> upstream/master
func GetNetInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Get net information",
		Run:   netInfo,
	}
	return cmd
}

func netInfo(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res rpctypes.NodeNetinfo
<<<<<<< HEAD
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.GetNetInfo", nil, &res)
	ctx.Run()
}

// get FatalFailure
=======
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.GetNetInfo", nil, &res)
	ctx.Run()
}

// GetFatalFailureCmd get FatalFailure
>>>>>>> upstream/master
func GetFatalFailureCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fault",
		Short: "Get system fault",
		Run:   fatalFailure,
	}
	return cmd
}

func fatalFailure(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res int64
<<<<<<< HEAD
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.GetFatalFailure", nil, &res)
	ctx.Run()
}

// get time status
=======
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.GetFatalFailure", nil, &res)
	ctx.Run()
}

// GetTimeStausCmd get time status
>>>>>>> upstream/master
func GetTimeStausCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "time",
		Short: "Get time status",
		Run:   timestatus,
	}
	return cmd
}

func timestatus(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res rpctypes.TimeStatus
<<<<<<< HEAD
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.GetTimeStatus", nil, &res)
=======
	ctx := jsonclient.NewRPCCtx(rpcLaddr, "Chain33.GetTimeStatus", nil, &res)
>>>>>>> upstream/master
	ctx.Run()
}
