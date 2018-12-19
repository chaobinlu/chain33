// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package consensus

import (
	"errors"
	"math/rand"
	"sync"
	"sync/atomic"

	"github.com/33cn/chain33/client"
	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/common/merkle"
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
	"github.com/33cn/chain33/util"
)

var tlog = log.New("module", "consensus")

var (
	zeroHash [32]byte
)

var randgen *rand.Rand

func init() {
	randgen = rand.New(rand.NewSource(types.Now().UnixNano()))
	QueryData.Register("base", &BaseClient{})
}

<<<<<<< HEAD
=======
//Miner 矿工
>>>>>>> upstream/master
type Miner interface {
	CreateGenesisTx() []*types.Transaction
	GetGenesisBlockTime() int64
	CreateBlock()
	CheckBlock(parent *types.Block, current *types.BlockDetail) error
	ProcEvent(msg queue.Message) bool
}

<<<<<<< HEAD
=======
//BaseClient ...
>>>>>>> upstream/master
type BaseClient struct {
	client       queue.Client
	api          client.QueueProtocolAPI
	minerStart   int32
	once         sync.Once
	Cfg          *types.Consensus
	currentBlock *types.Block
	mulock       sync.Mutex
	child        Miner
	minerstartCB func()
	isCaughtUp   int32
}

<<<<<<< HEAD
=======
//NewBaseClient ...
>>>>>>> upstream/master
func NewBaseClient(cfg *types.Consensus) *BaseClient {
	var flag int32
	if cfg.Minerstart {
		flag = 1
	}
	client := &BaseClient{minerStart: flag, isCaughtUp: 0}
	client.Cfg = cfg
	log.Info("Enter consensus " + cfg.Name)
	return client
}

<<<<<<< HEAD
func (client *BaseClient) GetGenesisBlockTime() int64 {
	return client.Cfg.GenesisBlockTime
}

=======
//GetGenesisBlockTime 获取创世区块时间
func (bc *BaseClient) GetGenesisBlockTime() int64 {
	return bc.Cfg.GenesisBlockTime
}

//SetChild ...
>>>>>>> upstream/master
func (bc *BaseClient) SetChild(c Miner) {
	bc.child = c
}

<<<<<<< HEAD
=======
//GetAPI 获取api
>>>>>>> upstream/master
func (bc *BaseClient) GetAPI() client.QueueProtocolAPI {
	return bc.api
}

<<<<<<< HEAD
=======
//InitClient 初始化
>>>>>>> upstream/master
func (bc *BaseClient) InitClient(c queue.Client, minerstartCB func()) {
	log.Info("Enter SetQueueClient method of consensus")
	bc.client = c
	bc.minerstartCB = minerstartCB
	bc.api, _ = client.New(c, nil)
	bc.InitMiner()
}

<<<<<<< HEAD
=======
//GetQueueClient 获取客户端队列
>>>>>>> upstream/master
func (bc *BaseClient) GetQueueClient() queue.Client {
	return bc.client
}

<<<<<<< HEAD
=======
//RandInt64 随机数
>>>>>>> upstream/master
func (bc *BaseClient) RandInt64() int64 {
	return randgen.Int63()
}

<<<<<<< HEAD
=======
//InitMiner 初始化矿工
>>>>>>> upstream/master
func (bc *BaseClient) InitMiner() {
	bc.once.Do(bc.minerstartCB)
}

<<<<<<< HEAD
=======
//Wait wait for ready
func (bc *BaseClient) Wait() {}

//SetQueueClient 设置客户端队列
>>>>>>> upstream/master
func (bc *BaseClient) SetQueueClient(c queue.Client) {
	bc.InitClient(c, func() {
		//call init block
		bc.InitBlock()
	})
	go bc.EventLoop()
	go bc.child.CreateBlock()
}

<<<<<<< HEAD
//change init block
=======
//InitBlock change init block
>>>>>>> upstream/master
func (bc *BaseClient) InitBlock() {
	block, err := bc.RequestLastBlock()
	if err != nil {
		panic(err)
	}
	if block == nil {
		// 创世区块
		newblock := &types.Block{}
		newblock.Height = 0
		newblock.BlockTime = bc.child.GetGenesisBlockTime()
		// TODO: 下面这些值在创世区块中赋值nil，是否合理？
		newblock.ParentHash = zeroHash[:]
		tx := bc.child.CreateGenesisTx()
		newblock.Txs = tx
		newblock.TxHash = merkle.CalcMerkleRoot(newblock.Txs)
		if newblock.Height == 0 {
			newblock.Difficulty = types.GetP(0).PowLimitBits
		}
		bc.WriteBlock(zeroHash[:], newblock)
	} else {
		bc.SetCurrentBlock(block)
	}
}

<<<<<<< HEAD
=======
//Close 关闭
>>>>>>> upstream/master
func (bc *BaseClient) Close() {
	atomic.StoreInt32(&bc.minerStart, 0)
	bc.client.Close()
	log.Info("consensus base closed")
}

<<<<<<< HEAD
//为了不引起交易检查时候产生的无序
=======
//CheckTxDup 为了不引起交易检查时候产生的无序
>>>>>>> upstream/master
func (bc *BaseClient) CheckTxDup(txs []*types.Transaction) (transactions []*types.Transaction) {
	cacheTxs := types.TxsToCache(txs)
	var err error
	cacheTxs, err = util.CheckTxDup(bc.client, cacheTxs, 0)
	if err != nil {
		return txs
	}
	return types.CacheToTxs(cacheTxs)
}

<<<<<<< HEAD
=======
//IsMining 是否在挖矿
>>>>>>> upstream/master
func (bc *BaseClient) IsMining() bool {
	return atomic.LoadInt32(&bc.minerStart) == 1
}

<<<<<<< HEAD
=======
//IsCaughtUp 是否追上最新高度
>>>>>>> upstream/master
func (bc *BaseClient) IsCaughtUp() bool {
	if bc.client == nil {
		panic("bc not bind message queue.")
	}
	msg := bc.client.NewMessage("blockchain", types.EventIsSync, nil)
	bc.client.Send(msg, true)
	resp, err := bc.client.Wait(msg)
	if err != nil {
		return false
	}
	return resp.GetData().(*types.IsCaughtUp).GetIscaughtup()
}

<<<<<<< HEAD
=======
//ExecConsensus 执行共识
>>>>>>> upstream/master
func (bc *BaseClient) ExecConsensus(data *types.ChainExecutor) (types.Message, error) {
	param, err := QueryData.Decode(data.Driver, data.FuncName, data.Param)
	if err != nil {
		return nil, err
	}
	return QueryData.Call(data.Driver, data.FuncName, param)
}

<<<<<<< HEAD
// 准备新区块
=======
//EventLoop 准备新区块
>>>>>>> upstream/master
func (bc *BaseClient) EventLoop() {
	// 监听blockchain模块，获取当前最高区块
	bc.client.Sub("consensus")
	go func() {
		for msg := range bc.client.Recv() {
			tlog.Debug("consensus recv", "msg", msg)
			if msg.Ty == types.EventConsensusQuery {
				exec := msg.GetData().(*types.ChainExecutor)
				param, err := QueryData.Decode(exec.Driver, exec.FuncName, exec.Param)
				if err != nil {
					msg.Reply(bc.api.NewMessage("", 0, err))
					continue
				}
				reply, err := QueryData.Call(exec.Driver, exec.FuncName, param)
				if err != nil {
					msg.Reply(bc.api.NewMessage("", 0, err))
				} else {
					msg.Reply(bc.api.NewMessage("", 0, reply))
				}
			} else if msg.Ty == types.EventAddBlock {
				block := msg.GetData().(*types.BlockDetail).Block
				bc.SetCurrentBlock(block)
			} else if msg.Ty == types.EventCheckBlock {
				block := msg.GetData().(*types.BlockDetail)
				err := bc.CheckBlock(block)
				msg.ReplyErr("EventCheckBlock", err)
			} else if msg.Ty == types.EventMinerStart {
				if !atomic.CompareAndSwapInt32(&bc.minerStart, 0, 1) {
					msg.ReplyErr("EventMinerStart", types.ErrMinerIsStared)
				} else {
					bc.InitMiner()
					msg.ReplyErr("EventMinerStart", nil)
				}
			} else if msg.Ty == types.EventMinerStop {
				if !atomic.CompareAndSwapInt32(&bc.minerStart, 1, 0) {
					msg.ReplyErr("EventMinerStop", types.ErrMinerNotStared)
				} else {
					msg.ReplyErr("EventMinerStop", nil)
				}
			} else if msg.Ty == types.EventDelBlock {
				block := msg.GetData().(*types.BlockDetail).Block
				bc.UpdateCurrentBlock(block)
			} else {
				if !bc.child.ProcEvent(msg) {
					msg.ReplyErr("BaseClient.EventLoop() ", types.ErrActionNotSupport)
				}
			}
		}
	}()
}

<<<<<<< HEAD
=======
//CheckBlock 检查区块
>>>>>>> upstream/master
func (bc *BaseClient) CheckBlock(block *types.BlockDetail) error {
	//check parent
	if block.Block.Height <= 0 { //genesis block not check
		return nil
	}
	parent, err := bc.RequestBlock(block.Block.Height - 1)
	if err != nil {
		return err
	}
	//check base info
	if parent.Height+1 != block.Block.Height {
		return types.ErrBlockHeight
	}
	if types.IsFork(block.Block.Height, "ForkCheckBlockTime") && parent.BlockTime > block.Block.BlockTime {
		return types.ErrBlockTime
	}
	//check parent hash
	if string(block.Block.GetParentHash()) != string(parent.Hash()) {
		return types.ErrParentHash
	}
	//check by drivers
	err = bc.child.CheckBlock(parent, block)
	return err
}

<<<<<<< HEAD
// Mempool中取交易列表
=======
//RequestTx Mempool中取交易列表
>>>>>>> upstream/master
func (bc *BaseClient) RequestTx(listSize int, txHashList [][]byte) []*types.Transaction {
	if bc.client == nil {
		panic("bc not bind message queue.")
	}
	msg := bc.client.NewMessage("mempool", types.EventTxList, &types.TxHashList{Hashes: txHashList, Count: int64(listSize)})
	bc.client.Send(msg, true)
	resp, err := bc.client.Wait(msg)
	if err != nil {
		return nil
	}
	return resp.GetData().(*types.ReplyTxList).GetTxs()
}

<<<<<<< HEAD
=======
//RequestBlock 请求区块
>>>>>>> upstream/master
func (bc *BaseClient) RequestBlock(start int64) (*types.Block, error) {
	if bc.client == nil {
		panic("bc not bind message queue.")
	}
<<<<<<< HEAD
	msg := bc.client.NewMessage("blockchain", types.EventGetBlocks, &types.ReqBlocks{start, start, false, []string{""}})
=======
	reqblock := &types.ReqBlocks{Start: start, End: start, IsDetail: false, Pid: []string{""}}
	msg := bc.client.NewMessage("blockchain", types.EventGetBlocks, reqblock)
>>>>>>> upstream/master
	bc.client.Send(msg, true)
	resp, err := bc.client.Wait(msg)
	if err != nil {
		return nil, err
	}
	blocks := resp.GetData().(*types.BlockDetails)
	return blocks.Items[0].Block, nil
}

<<<<<<< HEAD
//获取最新的block从blockchain模块
=======
//RequestLastBlock 获取最新的block从blockchain模块
>>>>>>> upstream/master
func (bc *BaseClient) RequestLastBlock() (*types.Block, error) {
	if bc.client == nil {
		panic("client not bind message queue.")
	}
	msg := bc.client.NewMessage("blockchain", types.EventGetLastBlock, nil)
	bc.client.Send(msg, true)
	resp, err := bc.client.Wait(msg)
	if err != nil {
		return nil, err
	}
	block := resp.GetData().(*types.Block)
	return block, nil
}

//del mempool
func (bc *BaseClient) delMempoolTx(deltx []*types.Transaction) error {
	hashList := buildHashList(deltx)
	msg := bc.client.NewMessage("mempool", types.EventDelTxList, hashList)
	bc.client.Send(msg, true)
	resp, err := bc.client.Wait(msg)
	if err != nil {
		return err
	}
	if resp.GetData().(*types.Reply).GetIsOk() {
		return nil
	}
	return errors.New(string(resp.GetData().(*types.Reply).GetMsg()))
}

func buildHashList(deltx []*types.Transaction) *types.TxHashList {
	list := &types.TxHashList{}
	for i := 0; i < len(deltx); i++ {
		list.Hashes = append(list.Hashes, deltx[i].Hash())
	}
	return list
}

<<<<<<< HEAD
// 向blockchain写区块
func (bc *BaseClient) WriteBlock(prev []byte, block *types.Block) error {
=======
//WriteBlock 向blockchain写区块
func (bc *BaseClient) WriteBlock(prev []byte, block *types.Block) error {
	//保存block的原始信息用于删除mempool中的错误交易
	rawtxs := make([]*types.Transaction, len(block.Txs))
	copy(rawtxs, block.Txs)

>>>>>>> upstream/master
	blockdetail := &types.BlockDetail{Block: block}
	msg := bc.client.NewMessage("blockchain", types.EventAddBlockDetail, blockdetail)
	bc.client.Send(msg, true)
	resp, err := bc.client.Wait(msg)
	if err != nil {
		return err
	}
	blockdetail = resp.GetData().(*types.BlockDetail)
	//从mempool 中删除错误的交易
<<<<<<< HEAD
	deltx := diffTx(block.Txs, blockdetail.Block.Txs)
=======
	deltx := diffTx(rawtxs, blockdetail.Block.Txs)
>>>>>>> upstream/master
	if len(deltx) > 0 {
		bc.delMempoolTx(deltx)
	}
	if blockdetail != nil {
		bc.SetCurrentBlock(blockdetail.Block)
	} else {
		return errors.New("block detail is nil")
	}
	return nil
}

func diffTx(tx1, tx2 []*types.Transaction) (deltx []*types.Transaction) {
	txlist2 := make(map[string]bool)
	for _, tx := range tx2 {
		txlist2[string(tx.Hash())] = true
	}
	for _, tx := range tx1 {
		hash := string(tx.Hash())
		if _, ok := txlist2[hash]; !ok {
			deltx = append(deltx, tx)
		}
	}
	return deltx
}

<<<<<<< HEAD
=======
//SetCurrentBlock 设置当前区块
>>>>>>> upstream/master
func (bc *BaseClient) SetCurrentBlock(b *types.Block) {
	bc.mulock.Lock()
	bc.currentBlock = b
	bc.mulock.Unlock()
}

<<<<<<< HEAD
=======
//UpdateCurrentBlock 更新当前区块
>>>>>>> upstream/master
func (bc *BaseClient) UpdateCurrentBlock(b *types.Block) {
	bc.mulock.Lock()
	defer bc.mulock.Unlock()
	block, err := bc.RequestLastBlock()
	if err != nil {
		log.Error("UpdateCurrentBlock", "RequestLastBlock", err)
		return
	}
	bc.currentBlock = block
}

<<<<<<< HEAD
=======
//GetCurrentBlock 获取当前区块
>>>>>>> upstream/master
func (bc *BaseClient) GetCurrentBlock() (b *types.Block) {
	bc.mulock.Lock()
	defer bc.mulock.Unlock()
	return bc.currentBlock
}

<<<<<<< HEAD
=======
//GetCurrentHeight 获取当前高度
>>>>>>> upstream/master
func (bc *BaseClient) GetCurrentHeight() int64 {
	bc.mulock.Lock()
	start := bc.currentBlock.Height
	bc.mulock.Unlock()
	return start
}

<<<<<<< HEAD
=======
//Lock 上锁
>>>>>>> upstream/master
func (bc *BaseClient) Lock() {
	bc.mulock.Lock()
}

<<<<<<< HEAD
=======
//Unlock 解锁
>>>>>>> upstream/master
func (bc *BaseClient) Unlock() {
	bc.mulock.Unlock()
}

<<<<<<< HEAD
=======
//ConsensusTicketMiner ...
>>>>>>> upstream/master
func (bc *BaseClient) ConsensusTicketMiner(iscaughtup *types.IsCaughtUp) {
	if !atomic.CompareAndSwapInt32(&bc.isCaughtUp, 0, 1) {
		log.Info("ConsensusTicketMiner", "isCaughtUp", bc.isCaughtUp)
	} else {
		log.Info("ConsensusTicketMiner", "isCaughtUp", bc.isCaughtUp)
	}
}

<<<<<<< HEAD
=======
//AddTxsToBlock 添加交易到区块中
>>>>>>> upstream/master
func (bc *BaseClient) AddTxsToBlock(block *types.Block, txs []*types.Transaction) []*types.Transaction {
	size := block.Size()
	max := types.MaxBlockSize - 100000 //留下100K空间，添加其他的交易
	currentcount := int64(len(block.Txs))
	maxTx := types.GetP(block.Height).MaxTxNumber
	addedTx := make([]*types.Transaction, 0, len(txs))
	for i := 0; i < len(txs); i++ {
		txgroup, err := txs[i].GetTxGroup()
		if err != nil {
			continue
		}
		if txgroup == nil {
			if currentcount+1 > maxTx {
				return addedTx
			}
			size += txs[i].Size()
			if size > max {
				return addedTx
			}
			addedTx = append(addedTx, txs[i])
			block.Txs = append(block.Txs, txs[i])
		} else {
			if currentcount+int64(len(txgroup.Txs)) > maxTx {
				return addedTx
			}
			for i := 0; i < len(txgroup.Txs); i++ {
				size += txgroup.Txs[i].Size()
			}
			if size > max {
				return addedTx
			}
			addedTx = append(addedTx, txgroup.Txs...)
			block.Txs = append(block.Txs, txgroup.Txs...)
		}
	}
	return addedTx
}
