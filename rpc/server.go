// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"net"
	"net/rpc"
	"time"

	"github.com/33cn/chain33/client"
	"github.com/33cn/chain33/pluginmgr"
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
	"golang.org/x/net/context"
<<<<<<< HEAD

	// register gzip
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

var (
	remoteIpWhitelist = make(map[string]bool)
	rpcCfg            *types.Rpc
	jrpcFuncWhitelist = make(map[string]bool)
	grpcFuncWhitelist = make(map[string]bool)
	jrpcFuncBlacklist = make(map[string]bool)
	grpcFuncBlacklist = make(map[string]bool)
)

=======
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip" // register gzip
)

var (
	remoteIPWhitelist           = make(map[string]bool)
	rpcCfg                      *types.RPC
	jrpcFuncWhitelist           = make(map[string]bool)
	grpcFuncWhitelist           = make(map[string]bool)
	jrpcFuncBlacklist           = make(map[string]bool)
	grpcFuncBlacklist           = make(map[string]bool)
	rpcFilterPrintFuncBlacklist = make(map[string]bool)
)

// Chain33  a channel client
>>>>>>> upstream/master
type Chain33 struct {
	cli channelClient
}

<<<<<<< HEAD
=======
// Grpc a channelClient
>>>>>>> upstream/master
type Grpc struct {
	cli channelClient
}

<<<<<<< HEAD
type Grpcserver struct {
	grpc Grpc
	s    *grpc.Server
	l    net.Listener
	//addr string
}

type JSONRPCServer struct {
	jrpc Chain33
	s    *rpc.Server
	l    net.Listener
	//addr string
}

=======
// Grpcserver a object
type Grpcserver struct {
	grpc *Grpc
	s    *grpc.Server
	l    net.Listener
}

// NewGrpcServer new  GrpcServer object
func NewGrpcServer() *Grpcserver {
	return &Grpcserver{grpc: &Grpc{}}
}

// JSONRPCServer  a json rpcserver object
type JSONRPCServer struct {
	jrpc *Chain33
	s    *rpc.Server
	l    net.Listener
}

// Close json rpcserver close
>>>>>>> upstream/master
func (s *JSONRPCServer) Close() {
	if s.l != nil {
		s.l.Close()
	}
<<<<<<< HEAD
	s.jrpc.cli.Close()
}

func checkIpWhitelist(addr string) bool {
=======
	if s.jrpc != nil {
		s.jrpc.cli.Close()
	}
}

func checkIPWhitelist(addr string) bool {
>>>>>>> upstream/master
	//回环网络直接允许
	ip := net.ParseIP(addr)
	if ip.IsLoopback() {
		return true
	}
	ipv4 := ip.To4()
	if ipv4 != nil {
		addr = ipv4.String()
	}
<<<<<<< HEAD
	if _, ok := remoteIpWhitelist["0.0.0.0"]; ok {
		return true
	}
	if _, ok := remoteIpWhitelist[addr]; ok {
=======
	if _, ok := remoteIPWhitelist["0.0.0.0"]; ok {
		return true
	}
	if _, ok := remoteIPWhitelist[addr]; ok {
>>>>>>> upstream/master
		return true
	}
	return false
}

func checkJrpcFuncWhitelist(funcName string) bool {

	if _, ok := jrpcFuncWhitelist["*"]; ok {
		return true
	}

	if _, ok := jrpcFuncWhitelist[funcName]; ok {
		return true
	}
	return false
}
func checkGrpcFuncWhitelist(funcName string) bool {

	if _, ok := grpcFuncWhitelist["*"]; ok {
		return true
	}

	if _, ok := grpcFuncWhitelist[funcName]; ok {
		return true
	}
	return false
}
func checkJrpcFuncBlacklist(funcName string) bool {
	if _, ok := jrpcFuncBlacklist[funcName]; ok {
		return true
	}
	return false
}
func checkGrpcFuncBlacklist(funcName string) bool {
	if _, ok := grpcFuncBlacklist[funcName]; ok {
		return true
	}
	return false
}

<<<<<<< HEAD
=======
// Close grpcserver close
>>>>>>> upstream/master
func (j *Grpcserver) Close() {
	if j == nil {
		return
	}
	if j.l != nil {
		j.l.Close()
	}
<<<<<<< HEAD
	j.grpc.cli.Close()
}

func NewGRpcServer(c queue.Client, api client.QueueProtocolAPI) *Grpcserver {
	s := &Grpcserver{}
=======
	if j.grpc != nil {
		j.grpc.cli.Close()
	}
}

// NewGRpcServer new grpcserver object
func NewGRpcServer(c queue.Client, api client.QueueProtocolAPI) *Grpcserver {
	s := &Grpcserver{grpc: &Grpc{}}
>>>>>>> upstream/master
	s.grpc.cli.Init(c, api)
	var opts []grpc.ServerOption
	//register interceptor
	//var interceptor grpc.UnaryServerInterceptor
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if err := auth(ctx, info); err != nil {
			return nil, err
		}
		// Continue processing the request
		return handler(ctx, req)
	}
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
<<<<<<< HEAD
	server := grpc.NewServer(opts...)
	s.s = server
	types.RegisterChain33Server(server, &s.grpc)
	return s
}

func NewJSONRPCServer(c queue.Client, api client.QueueProtocolAPI) *JSONRPCServer {
	j := &JSONRPCServer{}
	j.jrpc.cli.Init(c, api)
	server := rpc.NewServer()
	j.s = server
	server.RegisterName("Chain33", &j.jrpc)
	return j
}

type RPC struct {
	cfg  *types.Rpc
=======
	if rpcCfg.EnableTLS {
		creds, err := credentials.NewServerTLSFromFile(rpcCfg.CertFile, rpcCfg.KeyFile)
		if err != nil {
			panic(err)
		}
		credsOps := grpc.Creds(creds)
		opts = append(opts, credsOps)
	}
	server := grpc.NewServer(opts...)
	s.s = server
	types.RegisterChain33Server(server, s.grpc)
	return s
}

// NewJSONRPCServer new json rpcserver object
func NewJSONRPCServer(c queue.Client, api client.QueueProtocolAPI) *JSONRPCServer {
	j := &JSONRPCServer{jrpc: &Chain33{}}
	j.jrpc.cli.Init(c, api)
	server := rpc.NewServer()
	j.s = server
	server.RegisterName("Chain33", j.jrpc)
	return j
}

// RPC a type object
type RPC struct {
	cfg  *types.RPC
>>>>>>> upstream/master
	gapi *Grpcserver
	japi *JSONRPCServer
	c    queue.Client
	api  client.QueueProtocolAPI
}

<<<<<<< HEAD
func InitCfg(cfg *types.Rpc) {
	rpcCfg = cfg
	InitIpWhitelist(cfg)
=======
// InitCfg  interfaces
func InitCfg(cfg *types.RPC) {
	rpcCfg = cfg
	InitIPWhitelist(cfg)
>>>>>>> upstream/master
	InitJrpcFuncWhitelist(cfg)
	InitGrpcFuncWhitelist(cfg)
	InitJrpcFuncBlacklist(cfg)
	InitGrpcFuncBlacklist(cfg)
<<<<<<< HEAD
}

func New(cfg *types.Rpc) *RPC {
=======
	InitFilterPrintFuncBlacklist()
}

// New produce a rpc by cfg
func New(cfg *types.RPC) *RPC {
>>>>>>> upstream/master
	InitCfg(cfg)
	return &RPC{cfg: cfg}
}

<<<<<<< HEAD
=======
// SetAPI set api of rpc
>>>>>>> upstream/master
func (r *RPC) SetAPI(api client.QueueProtocolAPI) {
	r.api = api
}

<<<<<<< HEAD
=======
// SetQueueClient set queue client
>>>>>>> upstream/master
func (r *RPC) SetQueueClient(c queue.Client) {
	gapi := NewGRpcServer(c, r.api)
	japi := NewJSONRPCServer(c, r.api)
	r.gapi = gapi
	r.japi = japi
	r.c = c
	//注册系统rpc
	pluginmgr.AddRPC(r)
	r.Listen()
}

<<<<<<< HEAD
=======
// SetQueueClientNoListen  set queue client with  no listen
>>>>>>> upstream/master
func (r *RPC) SetQueueClientNoListen(c queue.Client) {
	gapi := NewGRpcServer(c, r.api)
	japi := NewJSONRPCServer(c, r.api)
	r.gapi = gapi
	r.japi = japi
	r.c = c
}

<<<<<<< HEAD
func (rpc *RPC) Listen() (port1 int, port2 int) {
	var err error
	for i := 0; i < 10; i++ {
		port1, err = rpc.gapi.Listen()
=======
// Listen rpc listen
func (r *RPC) Listen() (port1 int, port2 int) {
	var err error
	for i := 0; i < 10; i++ {
		port1, err = r.gapi.Listen()
>>>>>>> upstream/master
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		break
	}
	for i := 0; i < 10; i++ {
<<<<<<< HEAD
		port2, err = rpc.japi.Listen()
=======
		port2, err = r.japi.Listen()
>>>>>>> upstream/master
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		break
	}
	//sleep for a while
	time.Sleep(time.Millisecond)
	return port1, port2
}

<<<<<<< HEAD
func (rpc *RPC) GetQueueClient() queue.Client {
	return rpc.c
}

func (rpc *RPC) GRPC() *grpc.Server {
	return rpc.gapi.s
}

func (rpc *RPC) JRPC() *rpc.Server {
	return rpc.japi.s
}

func (rpc *RPC) Close() {
	if rpc.gapi != nil {
		rpc.gapi.Close()
	}
	if rpc.japi != nil {
		rpc.japi.Close()
	}
}

func InitIpWhitelist(cfg *types.Rpc) {
	if len(cfg.Whitelist) == 0 && len(cfg.Whitlist) == 0 {
		remoteIpWhitelist["127.0.0.1"] = true
		return
	}
	if len(cfg.Whitelist) == 1 && cfg.Whitelist[0] == "*" {
		remoteIpWhitelist["0.0.0.0"] = true
		return
	}
	if len(cfg.Whitlist) == 1 && cfg.Whitlist[0] == "*" {
		remoteIpWhitelist["0.0.0.0"] = true
=======
// GetQueueClient get queue client
func (r *RPC) GetQueueClient() queue.Client {
	return r.c
}

// GRPC return grpc rpc
func (r *RPC) GRPC() *grpc.Server {
	return r.gapi.s
}

// JRPC return jrpc
func (r *RPC) JRPC() *rpc.Server {
	return r.japi.s
}

// Close rpc close
func (r *RPC) Close() {
	if r.gapi != nil {
		r.gapi.Close()
	}
	if r.japi != nil {
		r.japi.Close()
	}
}

// InitIPWhitelist init ip whitelist
func InitIPWhitelist(cfg *types.RPC) {
	if len(cfg.Whitelist) == 0 && len(cfg.Whitlist) == 0 {
		remoteIPWhitelist["127.0.0.1"] = true
		return
	}
	if len(cfg.Whitelist) == 1 && cfg.Whitelist[0] == "*" {
		remoteIPWhitelist["0.0.0.0"] = true
		return
	}
	if len(cfg.Whitlist) == 1 && cfg.Whitlist[0] == "*" {
		remoteIPWhitelist["0.0.0.0"] = true
>>>>>>> upstream/master
		return
	}
	if len(cfg.Whitelist) != 0 {
		for _, addr := range cfg.Whitelist {
<<<<<<< HEAD
			remoteIpWhitelist[addr] = true
=======
			remoteIPWhitelist[addr] = true
>>>>>>> upstream/master
		}
		return
	}
	if len(cfg.Whitlist) != 0 {
		for _, addr := range cfg.Whitlist {
<<<<<<< HEAD
			remoteIpWhitelist[addr] = true
=======
			remoteIPWhitelist[addr] = true
>>>>>>> upstream/master
		}
		return
	}

}

<<<<<<< HEAD
func InitJrpcFuncWhitelist(cfg *types.Rpc) {
=======
// InitJrpcFuncWhitelist init jrpc function whitelist
func InitJrpcFuncWhitelist(cfg *types.RPC) {
>>>>>>> upstream/master
	if len(cfg.JrpcFuncWhitelist) == 0 {
		jrpcFuncWhitelist["*"] = true
		return
	}
	if len(cfg.JrpcFuncWhitelist) == 1 && cfg.JrpcFuncWhitelist[0] == "*" {
		jrpcFuncWhitelist["*"] = true
		return
	}
	for _, funcName := range cfg.JrpcFuncWhitelist {
		jrpcFuncWhitelist[funcName] = true
	}
}

<<<<<<< HEAD
func InitGrpcFuncWhitelist(cfg *types.Rpc) {
=======
// InitGrpcFuncWhitelist init grpc function whitelist
func InitGrpcFuncWhitelist(cfg *types.RPC) {
>>>>>>> upstream/master
	if len(cfg.GrpcFuncWhitelist) == 0 {
		grpcFuncWhitelist["*"] = true
		return
	}
	if len(cfg.GrpcFuncWhitelist) == 1 && cfg.GrpcFuncWhitelist[0] == "*" {
		grpcFuncWhitelist["*"] = true
		return
	}
	for _, funcName := range cfg.GrpcFuncWhitelist {
		grpcFuncWhitelist[funcName] = true
	}
}

<<<<<<< HEAD
func InitJrpcFuncBlacklist(cfg *types.Rpc) {
=======
// InitJrpcFuncBlacklist init jrpc function blacklist
func InitJrpcFuncBlacklist(cfg *types.RPC) {
>>>>>>> upstream/master
	if len(cfg.JrpcFuncBlacklist) == 0 {
		jrpcFuncBlacklist["CloseQueue"] = true
		return
	}
	for _, funcName := range cfg.JrpcFuncBlacklist {
		jrpcFuncBlacklist[funcName] = true
	}

}

<<<<<<< HEAD
func InitGrpcFuncBlacklist(cfg *types.Rpc) {
=======
// InitGrpcFuncBlacklist init grpc function blacklist
func InitGrpcFuncBlacklist(cfg *types.RPC) {
>>>>>>> upstream/master
	if len(cfg.GrpcFuncBlacklist) == 0 {
		grpcFuncBlacklist["CloseQueue"] = true
		return
	}
	for _, funcName := range cfg.GrpcFuncBlacklist {
		grpcFuncBlacklist[funcName] = true
	}
}
<<<<<<< HEAD
=======

// InitFilterPrintFuncBlacklist rpc模块打印requet信息时需要过滤掉一些敏感接口的入参打印，比如钱包密码相关的
func InitFilterPrintFuncBlacklist() {
	rpcFilterPrintFuncBlacklist["UnLock"] = true
	rpcFilterPrintFuncBlacklist["SetPasswd"] = true
	rpcFilterPrintFuncBlacklist["GetSeed"] = true
	rpcFilterPrintFuncBlacklist["SaveSeed"] = true
	rpcFilterPrintFuncBlacklist["ImportPrivkey"] = true
}

func checkFilterPrintFuncBlacklist(funcName string) bool {
	if _, ok := rpcFilterPrintFuncBlacklist[funcName]; ok {
		return true
	}
	return false
}
>>>>>>> upstream/master
