package p2p

import (
	"encoding/hex"
	"sync"
	"time"

	pb "code.aliyun.com/chain33/chain33/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (p *peer) Start() {

	log.Debug("Peer", "Start", p.Addr())
	go p.heartBeat()

	return
}
func (p *peer) Close() {
	p.SetRunning(false)
	p.mconn.Close()
	close(p.taskPool)
	pub.Unsub(p.taskChan, "block", "tx")
}

type peer struct {
	wg         sync.WaitGroup
	pmutx      sync.Mutex
	nodeInfo   **NodeInfo
	conn       *grpc.ClientConn // source connection
	persistent bool
	isrunning  bool
	version    *Version
	key        string
	mconn      *MConnection
	peerAddr   *NetAddress
	peerStat   *Stat
	taskPool   chan struct{}
	taskChan   chan interface{} //tx block
}

func NewPeer(conn *grpc.ClientConn, nodeinfo **NodeInfo, remote *NetAddress) *peer {
	p := &peer{
		conn:     conn,
		taskPool: make(chan struct{}, 50),
		nodeInfo: nodeinfo,
	}

	p.peerStat = new(Stat)
	p.version = new(Version)
	p.version.SetSupport(true)
	p.SetRunning(true)
	p.key = (*nodeinfo).addrBook.GetKey()
	p.mconn = NewMConnection(conn, remote, p)
	return p
}

type Version struct {
	mtx            sync.Mutex
	versionSupport bool
}
type Stat struct {
	mtx sync.Mutex
	ok  bool
}

func (st *Stat) Ok() {
	st.mtx.Lock()
	defer st.mtx.Unlock()
	st.ok = true
}

func (st *Stat) NotOk() {
	st.mtx.Lock()
	defer st.mtx.Unlock()
	st.ok = false
}

func (st *Stat) IsOk() bool {
	st.mtx.Lock()
	defer st.mtx.Unlock()
	return st.ok
}

func (v *Version) SetSupport(ok bool) {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	v.versionSupport = ok
}

func (v *Version) IsSupport() bool {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	return v.versionSupport
}

func (p *peer) heartBeat() {
	for {
		if p.GetRunning() == false {
			return
		}

		if (*p.nodeInfo).IsNatDone() { //如果nat 没有结束，在nat 重试的过程中，exter port 是在随机变化，
			//此时对连接的远程节点公布自己的外端端口将是不准确的,导致外网无法获取其nat结束后真正的端口。
			break
		}
		time.Sleep(time.Second) //wait for natwork done
	}

	pcli := NewP2pCli(nil)
	for {
		if p.GetRunning() == false {
			return
		}
		err := pcli.SendVersion(p, *p.nodeInfo)
		P2pComm.CollectPeerStat(err, p)
		if err == nil {
			p.taskChan = pub.Sub("block", "tx")
			go p.sendStream()
			go p.readStream()
			break
		} else {
			time.Sleep(time.Second * 5)
			continue
		}
	}

	ticker := time.NewTicker(PingTimeout)
	defer ticker.Stop()
	for {
		if p.GetRunning() == false {
			return
		}
		select {
		case <-ticker.C:
			err := pcli.SendPing(p, *p.nodeInfo)
			P2pComm.CollectPeerStat(err, p)

		}

	}

}

func (p *peer) GetPeerInfo(version int32) (*pb.P2PPeerInfo, error) {
	return p.mconn.gcli.GetPeerInfo(context.Background(), &pb.P2PGetPeerInfo{Version: version})
}

func (p *peer) sendStream() {
	//Stream Send data
	for {
		if p.GetRunning() == false {
			log.Info("sendStream peer is not running")
			return
		}
		ctx, cancel := context.WithCancel(context.Background())
		resp, err := p.mconn.gcli.ServerStreamRead(ctx)
		P2pComm.CollectPeerStat(err, p)
		if err != nil {
			cancel()
			log.Error("sendStream", "CollectPeerStat", err)
			time.Sleep(time.Second * 5)
			continue
		}
		//send ping package
		ping, err := P2pComm.NewPingData(p)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		p2pdata := new(pb.BroadCastData)
		p2pdata.Value = &pb.BroadCastData_Ping{Ping: ping}
		if err := resp.Send(p2pdata); err != nil {
			resp.CloseSend()
			cancel()
			log.Error("sendStream", "sendping", err)
			time.Sleep(time.Second)
			continue
		}

		timeout := time.NewTimer(time.Second * 2)
		defer timeout.Stop()
	SEND_LOOP:
		for {

			select {
			case task := <-p.taskChan:
				if p.GetRunning() == false {
					resp.CloseSend()
					cancel()
					log.Error("sendStream peer is not running")
					return
				}
				p2pdata := new(pb.BroadCastData)
				if block, ok := task.(*pb.P2PBlock); ok {
					height := block.GetBlock().GetHeight()
					blockhash := hex.EncodeToString(block.GetBlock().GetTxHash())
					log.Debug("sendStream", "will send block", blockhash)
					pinfo, err := p.GetPeerInfo((*p.nodeInfo).cfg.GetVersion())
					P2pComm.CollectPeerStat(err, p)
					if err == nil {
						if pinfo.GetHeader().GetHeight() >= height {
							log.Debug("sendStream", "find peer height>this broadblock ,send process", "break")
							continue
						}
					}

					p2pdata.Value = &pb.BroadCastData_Block{Block: block}

				} else if tx, ok := task.(*pb.P2PTx); ok {
					txhash := hex.EncodeToString(tx.GetTx().Hash())
					log.Debug("sendStream", "will send tx", txhash)
					p2pdata.Value = &pb.BroadCastData_Tx{Tx: tx}
				}

				err := resp.Send(p2pdata)
				P2pComm.CollectPeerStat(err, p)
				if err != nil {
					log.Error("sendStream", "send", err)
					if grpc.Code(err) == codes.Unimplemented { //maybe order peers delete peer to BlackList
						(*p.nodeInfo).blacklist.Add(p.Addr())
					}
					time.Sleep(time.Second) //have a rest
					resp.CloseSend()
					cancel()

					break SEND_LOOP //下一次外循环重新获取stream
				}
				log.Debug("sendStream", "send data", "ok")

			case <-timeout.C:
				if p.GetRunning() == false {
					log.Error("sendStream timeout")
					resp.CloseSend()
					cancel()
					return
				}
				timeout.Reset(time.Second * 2)

			}
		}

	}
}

func (p *peer) readStream() {

	pcli := NewP2pCli(nil)

	for {
		if p.GetRunning() == false {
			return
		}
		ping, err := P2pComm.NewPingData(p)
		if err != nil {
			log.Error("readStream", "err:", err.Error())
			continue
		}
		resp, err := p.mconn.gcli.ServerStreamSend(context.Background(), ping)
		P2pComm.CollectPeerStat(err, p)
		if err != nil {
			log.Error("readStream", "serverstreamsend,err:", err)
			time.Sleep(time.Second * 5)
			continue
		}
		log.Debug("SubStreamBlock", "Start", p.Addr())

		for {
			if p.GetRunning() == false {
				return
			}
			data, err := resp.Recv()
			P2pComm.CollectPeerStat(err, p)
			if err != nil {
				log.Error("readStream", "recv,err:", err)
				resp.CloseSend()
				if grpc.Code(err) == codes.Unimplemented { //maybe order peers delete peer to BlackList
					(*p.nodeInfo).blacklist.Add(p.Addr())
				}
				time.Sleep(time.Second) //have a rest
				break
			}

			if block := data.GetBlock(); block != nil {
				if block.GetBlock() != nil {
					//如果已经有登记过的消息记录，则不发送给本地blockchain
					blockhash := hex.EncodeToString(block.GetBlock().GetTxHash())
					if Filter.QueryRecvData(blockhash) == true {
						continue
					}

					//判断比自己低的区块，则不发送给blockchain

					height, err := pcli.GetBlockHeight((*p.nodeInfo))
					if err == nil {
						if height >= block.GetBlock().GetHeight() {
							continue
						}
					}
					log.Info("readStream", "block==+======+====+=>Height", block.GetBlock().GetHeight(), "from peer", p.Addr())
					msg := (*p.nodeInfo).client.NewMessage("blockchain", pb.EventBroadcastAddBlock, block.GetBlock())
					err = (*p.nodeInfo).client.Send(msg, false)
					if err != nil {
						log.Error("readStream", "send to blockchain Error", err.Error())
						continue
					}
					Filter.RegRecvData(blockhash) //添加发送登记，下次通过stream 接收同样的消息的时候可以过滤
				}

			} else if tx := data.GetTx(); tx != nil {

				if tx.GetTx() != nil {
					txhash := hex.EncodeToString(tx.Tx.Hash())
					log.Debug("readStream", "tx", "0x"+txhash)
					if Filter.QueryRecvData(txhash) == true {
						continue //处理方式同上
					}
					msg := (*p.nodeInfo).client.NewMessage("mempool", pb.EventTx, tx.GetTx())
					(*p.nodeInfo).client.Send(msg, false)
					Filter.RegRecvData(txhash) //登记
				}
			}
		}
	}
}

func (p *peer) SetRunning(run bool) {
	p.pmutx.Lock()
	defer p.pmutx.Unlock()
	p.isrunning = run
}
func (p *peer) GetRunning() bool {
	p.pmutx.Lock()
	defer p.pmutx.Unlock()
	return p.isrunning
}

// makePersistent marks the peer as persistent.
func (p *peer) makePersistent() {

	p.persistent = true
}

// Addr returns peer's remote network address.
func (p *peer) Addr() string {
	return p.peerAddr.String()

}

// IsPersistent returns true if the peer is persitent, false otherwise.
func (p *peer) IsPersistent() bool {
	return p.persistent
}
