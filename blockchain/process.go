package blockchain

import (
	"bytes"
	"container/list"
	"math/big"
	"sync/atomic"
	"time"

	"code.aliyun.com/chain33/chain33/common"
	"code.aliyun.com/chain33/chain33/types"
	"code.aliyun.com/chain33/chain33/util"
)

// 处理共识模块过来的blockdetail，peer广播过来的block，以及从peer同步过来的block
// 共识模块和peer广播过来的block需要广播出去
//共识模块过来的Receipts不为空,广播和同步过来的Receipts为空
// 返回参数说明：是否主链，是否孤儿节点，具体err
func (b *BlockChain) ProcessBlock(broadcast bool, block *types.BlockDetail) (bool, bool, error) {

	b.chainLock.Lock()
	defer b.chainLock.Unlock()
	//blockchain close 时不再处理block
	if atomic.LoadInt32(&b.isclosed) == 1 {
		return false, false, types.ErrIsClosed
	}
	blockHash := block.Block.Hash()
	chainlog.Debug("ProcessBlock Processing block", "height", block.Block.Height, "blockHash", common.ToHex(blockHash))

	// 判断本block是否已经存在主链或者侧链中
	exists := b.blockExists(blockHash)
	if exists {
		chainlog.Debug("ProcessBlock already have block", "blockHash", common.ToHex(blockHash))
		return false, false, types.ErrBlockExist
	}

	// 判断本节点是否已经存在孤儿链中
	exists = b.orphanPool.IsKnownOrphan(blockHash)
	if exists {
		chainlog.Debug("ProcessBlock already have block(orphan)", "blockHash", common.ToHex(blockHash))
		return false, false, types.ErrBlockExist
	}

	//checkpoint 的处理流程，block的时间必须晚于上一次的checkpoint点，以后再增加处理

	// 判断本block的父block是否存在，如果不存在就将此block添加到孤儿链中
	var prevHashExists bool
	prevHash := block.Block.GetParentHash()
	//创世块0需要做一些特殊的判断
	if 0 == block.Block.GetHeight() {
		if bytes.Equal(prevHash, common.Hash{}.Bytes()) {
			prevHashExists = true
		}
	} else {
		prevHashExists = b.blockExists(prevHash)
	}
	if !prevHashExists {
		chainlog.Debug("ProcessBlock addOrphanBlock", "height", block.Block.GetHeight(), "blockHash", common.ToHex(blockHash), "prevHash", common.ToHex(prevHash))
		b.orphanPool.addOrphanBlock(broadcast, block.Block)
		return false, true, nil
	}

	// 尝试将此block添加到主链上
	isMainChain, err := b.maybeAcceptBlock(broadcast, block)
	if err != nil {
		return false, false, err
	}
	// 尝试处理blockHash对应的孤儿子节点
	err = b.processOrphans(blockHash)
	if err != nil {
		return false, false, err
	}

	chainlog.Debug("ProcessBlock", "Accepted block", common.ToHex(blockHash))

	return isMainChain, false, nil
}

//检查block是否已经存在index或者数据库中
func (b *BlockChain) blockExists(hash []byte) bool {
	// Check block index first (could be main chain or side chain blocks).
	if b.index.HaveBlock(hash) {
		return true
	}

	// 检测数据库中是否存在，通过hash获取blcok，并通过hash获取height
	block, _ := b.blockStore.LoadBlockByHash(hash)
	if block == nil {
		return false
	}
	height, _ := b.blockStore.GetHeightByBlockHash(hash)
	if height != -1 {
		return true
	}
	return false
}

//孤儿链的处理,将本hash对应的子block插入chain中
func (b *BlockChain) processOrphans(hash []byte) error {
	chainlog.Debug("processOrphans parent", "hash", common.ToHex(hash))

	processHashes := make([]string, 0, 100)
	processHashes = append(processHashes, string(hash))
	for len(processHashes) > 0 {
		// Pop the first hash to process from the slice.
		processHash := processHashes[0]
		//chainlog.Debug("processOrphans", "processHash", common.ToHex([]byte(processHash)))

		processHashes[0] = "" // Prevent GC leak.
		processHashes = processHashes[1:]

		//  处理以processHash为父hash的所有子block
		count := b.orphanPool.GetChildOrphanCount(processHash)
		for i := 0; i < count; i++ {
			orphan := b.orphanPool.GetChildOrphan(processHash, i)
			if orphan == nil {
				chainlog.Debug("processOrphans", "Found a nil entry at index", i, "orphan dependency list for block", common.ToHex([]byte(processHash)))
				continue
			}
			//chainlog.Debug("processOrphans", "orphan.block.height", orphan.block.Height)

			// 从孤儿池中删除此孤儿节点
			orphanHash := orphan.block.Hash()
			b.orphanPool.RemoveOrphanBlock(orphan)
			i--

			chainlog.Debug("processOrphans  maybeAcceptBlock", "height", orphan.block.GetHeight(), "hash", common.ToHex(orphan.block.Hash()))
			// 尝试将此孤儿节点添加到主链
			_, err := b.maybeAcceptBlock(orphan.broadcast, &types.BlockDetail{Block: orphan.block})
			if err != nil {
				return err
			}
			processHashes = append(processHashes, string(orphanHash))
			///chainlog.Debug("processOrphans", "orphanHash", common.ToHex(orphanHash))
			//chainlog.Debug("processOrphans", "processHashes[0]", common.ToHex([]byte(processHashes[0])))
		}
	}
	return nil
}

// 尝试接受此block
func (b *BlockChain) maybeAcceptBlock(broadcast bool, block *types.BlockDetail) (bool, error) {
	// 首先判断本block的Parent block是否存在index中
	prevHash := block.Block.GetParentHash()
	prevNode := b.index.LookupNode(prevHash)
	if prevNode == nil {
		chainlog.Debug("maybeAcceptBlock", "previous block is unknown", common.ToHex(prevHash))
		return false, types.ErrParentBlockNoExist
	}

	blockHeight := block.Block.GetHeight()
	if blockHeight != prevNode.height+1 {
		chainlog.Debug("maybeAcceptBlock height err", "blockHeight", blockHeight, "prevHeight", prevNode.height)
		return false, types.ErrBlockHeightNoMatch
	}

	//将此block存储到db中，方便后面blockchain重组时使用，加入到主链saveblock时通过hash重新覆盖即可
	b.blockStore.dbMaybeStoreBlock(block)

	// 创建一个node并添加到内存中index
	newNode := newBlockNode(broadcast, block.Block)
	if prevNode != nil {
		newNode.parent = prevNode
	}
	b.index.AddNode(newNode)

	// 将本block添加到主链中
	isMainChain, err := b.connectBestChain(newNode, block)
	if err != nil {
		return false, err
	}

	// Notify the caller that the new block was accepted into the block
	// chain.  The caller would typically want to react by relaying the
	// inventory to other peers.

	//b.sendNotification(NTBlockAccepted, block)

	return isMainChain, nil
}

//将block添加到主链中
func (b *BlockChain) connectBestChain(node *blockNode, block *types.BlockDetail) (bool, error) {

	// 将此block插入到主链
	parentHash := block.Block.GetParentHash()
	if bytes.Equal(parentHash, b.bestChain.Tip().hash) {

		// 将此block添加到主链中,tip节点刚好是插入block的父节点.
		err := b.connectBlock(node, block)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	chainlog.Debug("connectBestChain", "parentHash", common.ToHex(parentHash), "bestChain.Tip().hash", common.ToHex(b.bestChain.Tip().hash))

	// 获取tip节点的block总难度tipid
	tiptd, _ := b.blockStore.GetTdByBlockHash(b.bestChain.Tip().hash)
	parenttd, _ := b.blockStore.GetTdByBlockHash(parentHash)
	if parenttd == nil {
		chainlog.Error("connectBestChain parenttd is not exits!", "hieght", block.Block.Height, "parentHash", common.ToHex(parentHash), "block.Block.hash", common.ToHex(block.Block.Hash()))
		return false, types.ErrParentTdNoExist
	}
	blocktd := new(big.Int).Add(node.Difficulty, parenttd)

	chainlog.Debug("connectBestChain tip:", "hash", common.ToHex(b.bestChain.Tip().hash), "height", b.bestChain.Tip().height, "TD", common.BigToCompact(tiptd))
	chainlog.Debug("connectBestChain node:", "hash", common.ToHex(node.hash), "height", node.height, "TD", common.BigToCompact(blocktd))

	if blocktd.Cmp(tiptd) <= 0 {
		fork := b.bestChain.FindFork(node)
		if bytes.Equal(parentHash, fork.hash) {
			chainlog.Info("connectBestChain FORK:", "Block hash", common.ToHex(node.hash), "fork.height", fork.height, "fork.hash", common.ToHex(fork.hash))
		} else {
			chainlog.Info("connectBestChain extends a side chain:", "Block hash", common.ToHex(node.hash), "fork.height", fork.height, "fork.hash", common.ToHex(fork.hash))
		}
		return false, nil
	}

	//print
	chainlog.Debug("connectBestChain tip", "height", b.bestChain.Tip().height, "hash", common.ToHex(b.bestChain.Tip().hash))
	chainlog.Debug("connectBestChain node", "height", node.height, "hash", common.ToHex(node.hash), "parentHash", common.ToHex(parentHash))
	chainlog.Debug("connectBestChain block", "height", block.Block.Height, "hash", common.ToHex(block.Block.Hash()))

	// 获取需要重组的block node
	detachNodes, attachNodes := b.getReorganizeNodes(node)

	// Reorganize the chain.
	//chainlog.Info("connectBestChain REORGANIZE:", "block height", node.height, "block hash", common.ToHex(node.hash))
	err := b.reorganizeChain(detachNodes, attachNodes)
	if err != nil {
		return false, err
	}

	return true, nil
}

//将本block信息存储到数据库中，并更新bestchain的tip节点
func (b *BlockChain) connectBlock(node *blockNode, blockdetail *types.BlockDetail) error {

	//blockchain close 时不再处理block
	if atomic.LoadInt32(&b.isclosed) == 1 {
		return types.ErrIsClosed
	}

	// Make sure it's extending the end of the best chain.
	parentHash := blockdetail.Block.GetParentHash()
	if !bytes.Equal(parentHash, b.bestChain.Tip().hash) {
		chainlog.Error("connectBlock hash err", "height", blockdetail.Block.Height, "Tip.height", b.bestChain.Tip().height)
		return types.ErrBlockHashNoMatch
	}
	var err error
	block := blockdetail.Block
	prevStateHash := b.bestChain.Tip().statehash
	//广播或者同步过来的blcok需要调用执行模块

	if !isStrongConsistency || blockdetail.Receipts == nil {
		blockdetail, err = util.ExecBlock(b.q, prevStateHash, block, true)
		if err != nil {
			chainlog.Error("connectBlock ExecBlock is err!", "height", block.Height, "err", err)
			return err
		}
	}

	beg := time.Now()
	// 写入磁盘
	//批量将block信息写入磁盘
	newbatch := b.blockStore.NewBatch(true)
	//保存tx信息到db中 (newbatch, blockdetail)
	err = b.blockStore.AddTxs(newbatch, blockdetail)
	if err != nil {
		chainlog.Error("connectBlock indexTxs:", "height", block.Height, "err", err)
		return err
	}
	chainlog.Debug("connectBlock AddTxs!", "height", block.Height)

	//保存block信息到db中
	err = b.blockStore.SaveBlock(newbatch, blockdetail)
	if err != nil {
		chainlog.Error("connectBlock SaveBlock:", "height", block.Height, "err", err)
		return err
	}

	//保存block的总难度到db中
	difficulty := common.CalcWork(block.Difficulty)
	var blocktd *big.Int
	if block.Height == 0 {
		blocktd = difficulty
	} else {
		parenttd, _ := b.blockStore.GetTdByBlockHash(parentHash)
		blocktd = new(big.Int).Add(difficulty, parenttd)
		//chainlog.Error("connectBlock Difficulty", "height", block.Height, "parenttd.td", common.BigToCompact(parenttd))
		//chainlog.Error("connectBlock Difficulty", "height", block.Height, "self.td", common.BigToCompact(blocktd))
	}

	err = b.blockStore.SaveTdByBlockHash(newbatch, blockdetail.Block.Hash(), blocktd)
	if err != nil {
		chainlog.Error("connectBlock SaveTdByBlockHash:", "height", block.Height, "err", err)
		return err
	}
	newbatch.Write()

	chainlog.Debug("connectBlock write db", "height", block.Height, "cost", time.Now().Sub(beg))

	b.blockStore.UpdateHeight()

	// 更新 best chain的tip节点
	b.bestChain.SetTip(node)

	b.query.updateStateHash(blockdetail.GetBlock().GetStateHash())

	b.SendAddBlockEvent(blockdetail)

	// 通知此block已经处理完，主要处理孤儿节点时需要设置
	b.task.Done(blockdetail.Block.GetHeight())

	//广播此block到全网络
	if node.broadcast {
		b.SendBlockBroadcast(blockdetail)
	}
	return nil
}

//从主链中删除blocks
func (b *BlockChain) disconnectBlock(node *blockNode, blockdetail *types.BlockDetail) error {
	// 只能从 best chain tip节点开始删除
	if !bytes.Equal(node.hash, b.bestChain.Tip().hash) {
		chainlog.Error("disconnectBlock:", "height", blockdetail.Block.Height, "node.hash", common.ToHex(node.hash), "bestChain.top.hash", common.ToHex(b.bestChain.Tip().hash))
		return types.ErrBlockHashNoMatch
	}

	//批量删除block的信息从磁盘中
	newbatch := b.blockStore.NewBatch(true)

	//从db中删除tx相关的信息
	err := b.blockStore.DelTxs(newbatch, blockdetail)
	if err != nil {
		chainlog.Error("disconnectBlock DelTxs:", "height", blockdetail.Block.Height, "err", err)
		return err
	}

	//从db中删除block相关的信息
	err = b.blockStore.DelBlock(newbatch, blockdetail)
	if err != nil {
		chainlog.Error("disconnectBlock DelBlock:", "height", blockdetail.Block.Height, "err", err)
		return err
	}
	newbatch.Write()

	b.blockStore.UpdateHeight()

	// 删除主链的tip节点，将其父节点升级成tip节点
	b.bestChain.DelTip(node)

	//通知共识，mempool和钱包删除block
	b.SendDelBlockEvent(blockdetail)

	b.query.updateStateHash(node.parent.statehash)

	//确定node的父节点升级成tip节点
	newtipnode := b.bestChain.Tip()

	//删除缓存中的block信息
	b.DelBlockFromCache(blockdetail.Block.Height)

	if newtipnode != node.parent {
		chainlog.Error("disconnectBlock newtipnode err:", "newtipnode.height", newtipnode.height, "node.parent.height", node.parent.height)
	}
	if !bytes.Equal(blockdetail.Block.GetParentHash(), b.bestChain.Tip().hash) {
		chainlog.Error("disconnectBlock", "newtipnode.height", newtipnode.height, "node.parent.height", node.parent.height)
		chainlog.Error("disconnectBlock", "newtipnode.hash", common.ToHex(newtipnode.hash), "delblock.parent.hash", common.ToHex(blockdetail.Block.GetParentHash()))
	}

	chainlog.Debug("disconnectBlock success", "newtipnode.height", newtipnode.height, "node.parent.height", node.parent.height)
	chainlog.Debug("disconnectBlock success", "newtipnode.hash", common.ToHex(newtipnode.hash), "delblock.parent.hash", common.ToHex(blockdetail.Block.GetParentHash()))

	return nil
}

//获取重组blockchain需要删除和添加节点
func (b *BlockChain) getReorganizeNodes(node *blockNode) (*list.List, *list.List) {
	attachNodes := list.New()
	detachNodes := list.New()

	// 查找到分叉的节点，并将分叉之后的block从index链push到attachNodes中
	forkNode := b.bestChain.FindFork(node)
	for n := node; n != nil && n != forkNode; n = n.parent {
		attachNodes.PushFront(n)
	}

	// 查找到分叉的节点，并将分叉之后的block从bestchain链push到attachNodes中
	for n := b.bestChain.Tip(); n != nil && n != forkNode; n = n.parent {
		detachNodes.PushBack(n)
	}

	return detachNodes, attachNodes
}

//重组blockchain
func (b *BlockChain) reorganizeChain(detachNodes, attachNodes *list.List) error {
	detachBlocks := make([]*types.BlockDetail, 0, detachNodes.Len())
	attachBlocks := make([]*types.BlockDetail, 0, attachNodes.Len())

	//通过node中的blockhash获取block信息从db中
	for e := detachNodes.Front(); e != nil; e = e.Next() {
		n := e.Value.(*blockNode)
		var block *types.BlockDetail
		block, _ = b.blockStore.LoadBlockByHash(n.hash)

		// 需要删除的blocks
		detachBlocks = append(detachBlocks, block)
		chainlog.Debug("reorganizeChain detachBlocks ", "height", block.Block.Height, "hash", common.ToHex(block.Block.Hash()))
	}

	for e := attachNodes.Front(); e != nil; e = e.Next() {
		n := e.Value.(*blockNode)
		var block *types.BlockDetail
		block, _ = b.blockStore.LoadBlockByHash(n.hash)

		// 需要加载到db的blocks
		attachBlocks = append(attachBlocks, block)
		chainlog.Debug("reorganizeChain attachBlocks ", "height", block.Block.Height, "hash", common.ToHex(block.Block.Hash()))
	}

	// Disconnect blocks from the main chain.
	for i, e := 0, detachNodes.Front(); e != nil; i, e = i+1, e.Next() {
		n := e.Value.(*blockNode)
		block := detachBlocks[i]

		// Update the database and chain state.
		err := b.disconnectBlock(n, block)
		if err != nil {
			return err
		}
	}

	// Connect the new best chain blocks.
	for i, e := 0, attachNodes.Front(); e != nil; i, e = i+1, e.Next() {
		n := e.Value.(*blockNode)
		block := attachBlocks[i]

		// Update the database and chain state.
		err := b.connectBlock(n, block)
		if err != nil {
			return err
		}
	}

	// Log the point where the chain forked and old and new best chain
	// heads.
	if attachNodes.Front() != nil {
		firstAttachNode := attachNodes.Front().Value.(*blockNode)
		chainlog.Debug("REORGANIZE: Chain forks at hash", "hash", common.ToHex(firstAttachNode.parent.hash), "height", firstAttachNode.parent.height)

	}
	if detachNodes.Front() != nil {
		firstDetachNode := detachNodes.Front().Value.(*blockNode)
		chainlog.Debug("REORGANIZE: Old best chain head was hash", "hash", common.ToHex(firstDetachNode.hash), "height", firstDetachNode.parent.height)

	}
	if attachNodes.Back() != nil {
		lastAttachNode := attachNodes.Back().Value.(*blockNode)
		chainlog.Debug("REORGANIZE: New best chain head is hash", "hash", common.ToHex(lastAttachNode.hash), "height", lastAttachNode.parent.height)
	}
	return nil
}