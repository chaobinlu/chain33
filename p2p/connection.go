package p2p

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"code.aliyun.com/chain33/chain33/common/crypto"
	pb "code.aliyun.com/chain33/chain33/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type MConnection struct {
	nodeInfo      **NodeInfo
	gconn         *grpc.ClientConn
	conn          pb.P2PgserviceClient // source connection
	config        *MConnConfig
	key           string //pubkey
	quit          chan bool
	remoteAddress *NetAddress  //peer 的地址
	pingTimer     *RepeatTimer // send pings periodically
	versionTimer  *RepeatTimer
	peer          *peer
	sendMonitor   *Monitor
}

// MConnConfig is a MConnection configuration.
type MConnConfig struct {
	gconn *grpc.ClientConn
	conn  pb.P2PgserviceClient
}

// DefaultMConnConfig returns the default config.
func DefaultMConnConfig() *MConnConfig {
	return &MConnConfig{}
}

func NewTemMConnConfig(gconn *grpc.ClientConn, conn pb.P2PgserviceClient) *MConnConfig {
	return &MConnConfig{
		gconn: gconn,
		conn:  conn,
	}
}

// NewMConnection wraps net.Conn and creates multiplex connection
func NewMConnection(conn *grpc.ClientConn, remote *NetAddress, peer *peer) *MConnection {

	mconn := &MConnection{
		gconn:       conn,
		conn:        pb.NewP2PgserviceClient(conn),
		pingTimer:   NewRepeatTimer("ping", pingTimeout),
		sendMonitor: NewMonitor(),
		peer:        peer,
		quit:        make(chan bool),
	}
	mconn.nodeInfo = peer.nodeInfo
	mconn.remoteAddress = remote

	return mconn

}

func NewMConnectionWithConfig(cfg *MConnConfig) *MConnection {
	mconn := &MConnection{
		gconn: cfg.gconn,
		conn:  cfg.conn,
	}

	return mconn
}

func (c *MConnection) Signature(in *pb.P2PPing) (*pb.P2PPing, error) {
	data := pb.Encode(in)

	cr, err := crypto.New(pb.GetSignatureTypeName(pb.SECP256K1))
	if err != nil {
		log.Error("CryPto Error", "Error", err.Error())
		return nil, err
	}
	pribyts, err := hex.DecodeString(c.key)
	if err != nil {
		log.Error("DecodeString Error", "Error", err.Error())
		return nil, err
	}
	priv, err := cr.PrivKeyFromBytes(pribyts)
	if err != nil {
		log.Error("Load PrivKey", "Error", err.Error())
		return nil, err
	}
	in.Sign = new(pb.Signature)
	in.Sign.Signature = priv.Sign(data).Bytes()
	in.Sign.Ty = pb.SECP256K1
	in.Sign.Pubkey = priv.PubKey().Bytes()
	return in, nil
}

// sendRoutine polls for packets to send from channels.
func (c *MConnection) pingRoutine() {

	var pingtimes int64
FOR_LOOP:
	for {

		select {
		case <-c.pingTimer.Ch:
			randNonce := rand.Int31n(102040)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			in, err := c.Signature(&pb.P2PPing{Nonce: int64(randNonce), Addr: EXTERNALADDR, Port: int32((*c.nodeInfo).externalAddr.Port)})
			if err != nil {
				log.Error("Signature", "Error", err.Error())
				continue
			}
			log.Debug("SEND PING", "Peer", c.remoteAddress.String(), "nonce", randNonce)
			r, err := c.conn.Ping(ctx, in)
			if err != nil {
				c.sendMonitor.Update(false)
				log.Error("ERR PING", "ERROR", err.Error())
				if pingtimes == 0 {
					(*c.nodeInfo).monitorChan <- c.peer

				}
				continue
			}

			log.Debug("RECV PONG", "resp:", r.Nonce, "Ping nonce:", randNonce)
			c.sendMonitor.Update(true)
			pingtimes++
			//Send to Version,Once
			if pingtimes == 5 {
				c.ExChangeVersion(in)
			}
			c.pingTimer.Reset()

		case <-c.quit:
			break FOR_LOOP

		}

	}

}
func (c *MConnection) ExChangeVersion(in *pb.P2PPing) {
	var once sync.Once
	f := func() {
		//get blockheight
		client := (*c.nodeInfo).q.GetClient()
		msg := client.NewMessage("blockchain", pb.EventGetBlockHeight, nil)
		client.Send(msg, true)
		rsp, err := client.Wait(msg)
		if err != nil {
			log.Error("GetHeight", "Error", err.Error())
			return
		}

		blockheight := rsp.GetData().(*pb.ReplyBlockHeight).GetHeight()
		resp, err := c.conn.Version2(context.Background(), &pb.P2PVersion{Version: Version, Service: SERVICE, Timestamp: time.Now().Unix(),
			AddrRecv: c.remoteAddress.String(), AddrFrom: fmt.Sprintf("%v:%v", EXTERNALADDR, (*c.nodeInfo).externalAddr.Port), Nonce: int64(rand.Int31n(102040)),
			UserAgent: hex.EncodeToString(in.Sign.GetPubkey()), StartHeight: blockheight})
		if err != nil {
			c.peer.persistent = false
			c.Stop()
			log.Error("SendVersion2", "Error", err.Error())

			return
		}
		selfExternAddr := resp.AddrRecv
		log.Info("Version SelfAddr", "Addr", selfExternAddr)
		log.Debug("SHOW VERSION BACK", "VersionBack", resp)
		return

	}
	once.Do(f)
}

func (c *MConnection) GetAddr() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := c.conn.GetAddr(ctx, &pb.P2PGetAddr{Nonce: int64(rand.Int31n(102040))})
	if err != nil {
		log.Error("GetAddr", "err", err.Error())
		c.sendMonitor.Update(false)
		return nil, err
	}

	log.Debug("GetAddr Resp", "Resp", resp, "addrlist", resp.Addrlist)
	c.sendMonitor.Update(true)
	return resp.Addrlist, nil

}

// OnStart implements BaseService
func (c *MConnection) Start() error { //启动Mconnection，每一个MConnection 会在启动的时候启动SendRoutine,RecvRoutine

	go c.pingRoutine() //创建发送Routine
	return nil
}

func (c *MConnection) Close() {
	c.gconn.Close()
}
func (c *MConnection) Stop() {

	c.pingTimer.Stop()
	c.gconn.Close()
	c.quit <- false
	log.Debug("Mconnection", "Close", "^_^!")
}