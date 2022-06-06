// Copyright 2016 The go-ethereum Authors
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

package lxs

import (
	"context"
	"math/big"

	"github.com/xpaymentsorg/go-xpayments/accounts"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/common/math"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/core/bloombits"
	"github.com/xpaymentsorg/go-xpayments/core/rawdb"
	"github.com/xpaymentsorg/go-xpayments/core/state"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/core/vm"
	"github.com/xpaymentsorg/go-xpayments/light"
	"github.com/xpaymentsorg/go-xpayments/params"
	"github.com/xpaymentsorg/go-xpayments/rpc"
	"github.com/xpaymentsorg/go-xpayments/xps/downloader"
	"github.com/xpaymentsorg/go-xpayments/xps/gasprice"
)

type LxsApiBackend struct {
	xps           *LightXPS
	initialSupply *big.Int
	gpo           *gasprice.Oracle
}

func (b *LxsApiBackend) ChainConfig() *params.ChainConfig {
	return b.xps.chainConfig
}

func (b *LxsApiBackend) InitialSupply() *big.Int {
	return b.initialSupply
}

func (b *LxsApiBackend) GenesisAlloc() core.GenesisAlloc {
	if g := b.xps.config.Genesis; g != nil {
		return g.Alloc
	}
	return nil
}

func (b *LxsApiBackend) CurrentBlock() *types.Block {
	return types.NewBlockWithHeader(b.xps.BlockChain().CurrentHeader())
}

func (b *LxsApiBackend) SetHead(number uint64) {
	b.xps.protocolManager.downloader.Cancel()
	b.xps.blockchain.SetHead(number)
}

func (b *LxsApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	if blockNr == rpc.LatestBlockNumber || blockNr == rpc.PendingBlockNumber {
		return b.xps.blockchain.CurrentHeader(), nil
	}

	return b.xps.blockchain.GetHeaderByNumberOdr(ctx, uint64(blockNr))
}

func (b *LxsApiBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.xps.blockchain.GetHeaderByHash(hash), nil
}

func (b *LxsApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, err
	}
	return b.GetBlock(ctx, header.Hash())
}

func (b *LxsApiBackend) StateQuery(ctx context.Context, blockNr rpc.BlockNumber, fn func(*state.StateDB) error) error {
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return err
	}
	return fn(light.NewState(ctx, header, b.xps.odr))
}

func (b *LxsApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	return light.NewState(ctx, header, b.xps.odr), header, nil
}

func (b *LxsApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.xps.blockchain.GetBlockByHash(ctx, blockHash)
}

func (b *LxsApiBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	if number := rawdb.ReadHeaderNumber(b.xps.chainDb.GlobalTable(), hash); number != nil {
		return light.GetBlockReceipts(ctx, b.xps.odr, hash, *number)
	}
	return nil, nil
}

func (b *LxsApiBackend) GetLogs(ctx context.Context, hash common.Hash) ([][]*types.Log, error) {
	if number := rawdb.ReadHeaderNumber(b.xps.chainDb.GlobalTable(), hash); number != nil {
		return light.GetBlockLogs(ctx, b.xps.odr, hash, *number)
	}
	return nil, nil
}

func (b *LxsApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.xps.blockchain.GetTdByHash(blockHash)
}

func (b *LxsApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	context := core.NewEVMContext(msg, header, b.xps.blockchain, nil)
	return vm.NewEVM(context, state, b.xps.chainConfig, vmCfg), nil
}

func (b *LxsApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.xps.txPool.Add(ctx, signedTx)
}

func (b *LxsApiBackend) RemoveTx(txHash common.Hash) {
	b.xps.txPool.RemoveTx(txHash)
}

func (b *LxsApiBackend) GetPoolTransactions() types.Transactions {
	return b.xps.txPool.GetTransactions()
}

func (b *LxsApiBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	return b.xps.txPool.GetTransaction(txHash)
}

func (b *LxsApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.xps.txPool.GetNonce(ctx, addr)
}

func (b *LxsApiBackend) Stats() (pending int, queued int) {
	return b.xps.txPool.Stats(), 0
}

func (b *LxsApiBackend) TxPoolContent(ctx context.Context) (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.xps.txPool.Content(ctx)
}

func (b *LxsApiBackend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent, name string) {
	b.xps.txPool.SubscribeNewTxsEvent(ch, name)
}

func (b *LxsApiBackend) UnsubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) {
	b.xps.txPool.UnsubscribeNewTxsEvent(ch)
}

func (b *LxsApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent, name string) {
	b.xps.blockchain.SubscribeChainEvent(ch, name)
}

func (b *LxsApiBackend) UnsubscribeChainEvent(ch chan<- core.ChainEvent) {
	b.xps.blockchain.UnsubscribeChainEvent(ch)
}

func (b *LxsApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent, name string) {
	b.xps.blockchain.SubscribeChainHeadEvent(ch, name)
}

func (b *LxsApiBackend) UnsubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) {
	b.xps.blockchain.UnsubscribeChainHeadEvent(ch)
}

func (b *LxsApiBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent, name string) {
	b.xps.blockchain.SubscribeChainSideEvent(ch, name)
}

func (b *LxsApiBackend) UnsubscribeChainSideEvent(ch chan<- core.ChainSideEvent) {
	b.xps.blockchain.UnsubscribeChainSideEvent(ch)
}

func (b *LxsApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log, name string) {}

func (b *LxsApiBackend) UnsubscribeLogsEvent(ch chan<- []*types.Log) {}

func (b *LxsApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent, name string) {}

func (b *LxsApiBackend) UnsubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) {}

func (b *LxsApiBackend) SubscribePendingLogsEvent(ch chan<- core.PendingLogsEvent, name string) {}

func (b *LxsApiBackend) UnsubscribePendingLogsEvent(ch chan<- core.PendingLogsEvent) {}

func (b *LxsApiBackend) Downloader() *downloader.Downloader {
	return b.xps.Downloader()
}

func (b *LxsApiBackend) ProtocolVersion() int {
	return b.xps.LxsVersion() + 10000
}

func (b *LxsApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *LxsApiBackend) ChainDb() common.Database {
	return b.xps.chainDb
}

func (b *LxsApiBackend) AccountManager() *accounts.Manager {
	return b.xps.accountManager
}

func (b *LxsApiBackend) BloomStatus() (uint64, uint64) {
	if b.xps.bloomIndexer == nil {
		return 0, 0
	}
	sections, _, _ := b.xps.bloomIndexer.Sections()
	return light.BloomTrieFrequency, sections
}

func (b *LxsApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.xps.bloomRequests)
	}
}
