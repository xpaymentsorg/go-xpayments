// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package xps

import (
	"context"
	"math/big"

	"github.com/xpaymentsorg/go-xpayments/accounts"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/common/math"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/core/bloombits"
	"github.com/xpaymentsorg/go-xpayments/core/state"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/core/vm"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/params"
	"github.com/xpaymentsorg/go-xpayments/rpc"
	"github.com/xpaymentsorg/go-xpayments/xps/downloader"
	"github.com/xpaymentsorg/go-xpayments/xps/gasprice"
)

// XpsApiBackend implements xpsapi.Backend for full nodes
type XpsApiBackend struct {
	xps           *XPS
	initialSupply *big.Int
	gpo           *gasprice.Oracle
}

func (b *XpsApiBackend) ChainConfig() *params.ChainConfig {
	return b.xps.chainConfig
}

func (b *XpsApiBackend) InitialSupply() *big.Int {
	return b.initialSupply
}

func (b *XpsApiBackend) GenesisAlloc() core.GenesisAlloc {
	if g := b.xps.config.Genesis; g != nil {
		return g.Alloc
	}
	return nil
}

func (b *XpsApiBackend) CurrentBlock() *types.Block {
	return b.xps.blockchain.CurrentBlock()
}

func (b *XpsApiBackend) SetHead(number uint64) {
	b.xps.protocolManager.downloader.Cancel()
	if err := b.xps.blockchain.SetHead(number); err != nil {
		log.Error("Cannot set xps api backend head", "number", number, "err", err)
	}
}

func (b *XpsApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.xps.miner.PendingBlock()
		if block == nil {
			return nil, nil
		}
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.xps.blockchain.CurrentBlock().Header(), nil
	}
	return b.xps.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *XpsApiBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.xps.blockchain.GetHeaderByHash(hash), nil
}

func (b *XpsApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.xps.miner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.xps.blockchain.CurrentBlock(), nil
	}
	return b.xps.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *XpsApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, state := b.xps.miner.Pending()
		var header *types.Header
		if block != nil {
			header = block.Header()
		}
		return state, header, nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.xps.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *XpsApiBackend) GetBlock(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return b.xps.blockchain.GetBlockByHash(hash), nil
}

func (b *XpsApiBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	return b.xps.blockchain.GetReceiptsByHash(hash), nil
}

func (b *XpsApiBackend) GetLogs(ctx context.Context, hash common.Hash) ([][]*types.Log, error) {
	receipts := b.xps.blockchain.GetReceiptsByHash(hash)
	if receipts == nil {
		return nil, nil
	}
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

func (b *XpsApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.xps.blockchain.GetTdByHash(blockHash)
}

func (b *XpsApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, error) {
	state.SetBalance(msg.From(), math.MaxBig256)

	context := core.NewEVMContext(msg, header, b.xps.BlockChain(), nil)
	return vm.NewEVM(context, state, b.xps.chainConfig, vmCfg), nil
}

func (b *XpsApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent, name string) {
	b.xps.BlockChain().SubscribeRemovedLogsEvent(ch, name)
}

func (b *XpsApiBackend) UnsubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) {
	b.xps.BlockChain().UnsubscribeRemovedLogsEvent(ch)
}

func (b *XpsApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent, name string) {
	b.xps.BlockChain().SubscribeChainEvent(ch, name)
}

func (b *XpsApiBackend) UnsubscribeChainEvent(ch chan<- core.ChainEvent) {
	b.xps.BlockChain().UnsubscribeChainEvent(ch)
}

func (b *XpsApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent, name string) {
	b.xps.BlockChain().SubscribeChainHeadEvent(ch, name)
}

func (b *XpsApiBackend) UnsubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) {
	b.xps.BlockChain().UnsubscribeChainHeadEvent(ch)
}

func (b *XpsApiBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent, name string) {
	b.xps.BlockChain().SubscribeChainSideEvent(ch, name)
}

func (b *XpsApiBackend) UnsubscribeChainSideEvent(ch chan<- core.ChainSideEvent) {
	b.xps.BlockChain().UnsubscribeChainSideEvent(ch)
}

func (b *XpsApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log, name string) {
	b.xps.BlockChain().SubscribeLogsEvent(ch, name)
}

func (b *XpsApiBackend) UnsubscribeLogsEvent(ch chan<- []*types.Log) {
	b.xps.BlockChain().UnsubscribeLogsEvent(ch)
}

func (b *XpsApiBackend) SubscribePendingLogsEvent(ch chan<- core.PendingLogsEvent, name string) {
	b.xps.BlockChain().SubscribePendingLogsEvent(ch, name)
}

func (b *XpsApiBackend) UnsubscribePendingLogsEvent(ch chan<- core.PendingLogsEvent) {
	b.xps.BlockChain().UnsubscribePendingLogsEvent(ch)
}

func (b *XpsApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.xps.txPool.AddLocal(signedTx)
}

func (b *XpsApiBackend) GetPoolTransactions() types.Transactions {
	return b.xps.txPool.PendingList()
}

func (b *XpsApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.xps.txPool.Get(hash)
}

func (b *XpsApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.xps.txPool.State().GetNonce(addr), nil
}

func (b *XpsApiBackend) Stats() (pending int, queued int) {
	return b.xps.txPool.Stats()
}

func (b *XpsApiBackend) TxPoolContent(ctx context.Context) (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.xps.TxPool().Content()
}

func (b *XpsApiBackend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent, name string) {
	b.xps.TxPool().SubscribeNewTxsEvent(ch, name)
}

func (b *XpsApiBackend) UnsubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) {
	b.xps.TxPool().UnsubscribeNewTxsEvent(ch)
}

func (b *XpsApiBackend) Downloader() *downloader.Downloader {
	return b.xps.Downloader()
}

func (b *XpsApiBackend) ProtocolVersion() int {
	return b.xps.XpsVersion()
}

func (b *XpsApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *XpsApiBackend) ChainDb() common.Database {
	return b.xps.ChainDb()
}

func (b *XpsApiBackend) AccountManager() *accounts.Manager {
	return b.xps.AccountManager()
}

func (b *XpsApiBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.xps.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *XpsApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.xps.bloomRequests)
	}
}
