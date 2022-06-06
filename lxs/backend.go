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

// Package lxs implements the Light xPayments Subprotocol.
package lxs

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/xpaymentsorg/go-xpayments/accounts"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/common/hexutil"
	"github.com/xpaymentsorg/go-xpayments/consensus"
	"github.com/xpaymentsorg/go-xpayments/consensus/clique"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/core/bloombits"
	"github.com/xpaymentsorg/go-xpayments/core/rawdb"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/internal/xpsapi"
	"github.com/xpaymentsorg/go-xpayments/light"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/node"
	"github.com/xpaymentsorg/go-xpayments/p2p"
	"github.com/xpaymentsorg/go-xpayments/p2p/discv5"
	"github.com/xpaymentsorg/go-xpayments/params"
	"github.com/xpaymentsorg/go-xpayments/rpc"
	"github.com/xpaymentsorg/go-xpayments/xps"
	"github.com/xpaymentsorg/go-xpayments/xps/downloader"
	"github.com/xpaymentsorg/go-xpayments/xps/filters"
	"github.com/xpaymentsorg/go-xpayments/xps/gasprice"
)

type LightXPS struct {
	config *xps.Config

	odr         *LxsOdr
	relay       *LxsTxRelay
	chainConfig *params.ChainConfig
	// Channel for shutting down the service
	shutdownChan chan bool
	// Handlers
	peers           *peerSet
	txPool          *light.TxPool
	blockchain      *light.LightChain
	protocolManager *ProtocolManager
	serverPool      *serverPool
	reqDist         *requestDistributor
	retriever       *retrieveManager
	// DB interfaces
	chainDb common.Database // Block chain database

	bloomRequests                              chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer, chtIndexer, bloomTrieIndexer *core.ChainIndexer

	ApiBackend *LxsApiBackend

	eventMux       *core.InterfaceFeed
	engine         consensus.Engine
	accountManager *accounts.Manager

	networkId     uint64
	netRPCService *xpsapi.PublicNetAPI

	wg sync.WaitGroup
}

func New(ctx context.Context, sctx *node.ServiceContext, config *xps.Config) (*LightXPS, error) {
	chainDb, err := xps.CreateDB(sctx, config, "lightchaindata")
	if err != nil {
		return nil, err
	}
	if config.Genesis == nil {
		config.Genesis = core.DefaultGenesisBlock()
	}
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlockWithOverride(chainDb, config.Genesis, config.ConstantinopleOverride)
	if _, isCompat := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !isCompat {
		return nil, genesisErr
	}
	if config.Genesis == nil {
		if genesisHash == params.MainnetGenesisHash {
			config.Genesis = core.DefaultGenesisBlock()
		}
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	peers := newPeerSet()
	quitSync := make(chan struct{})

	if chainConfig.Clique == nil {
		return nil, fmt.Errorf("invalid configuration, clique is nil: %v", chainConfig)
	}
	lxps := &LightXPS{
		config:           config,
		chainConfig:      chainConfig,
		chainDb:          chainDb,
		eventMux:         sctx.EventMux,
		peers:            peers,
		reqDist:          newRequestDistributor(peers, quitSync),
		accountManager:   sctx.AccountManager,
		engine:           clique.New(chainConfig.Clique, chainDb),
		shutdownChan:     make(chan bool),
		networkId:        config.NetworkId,
		bloomRequests:    make(chan chan *bloombits.Retrieval),
		bloomIndexer:     xps.NewBloomIndexer(chainDb, light.BloomTrieFrequency),
		chtIndexer:       light.NewChtIndexer(chainDb, true),
		bloomTrieIndexer: light.NewBloomTrieIndexer(chainDb, true),
	}

	lxps.relay = NewLxsTxRelay(peers, lxps.reqDist)
	lxps.serverPool = newServerPool(chainDb, quitSync, &lxps.wg)
	lxps.retriever = newRetrieveManager(peers, lxps.reqDist, lxps.serverPool)
	lxps.odr = NewLxsOdr(chainDb, lxps.chtIndexer, lxps.bloomTrieIndexer, lxps.bloomIndexer, lxps.retriever)
	if lxps.blockchain, err = light.NewLightChain(lxps.odr, lxps.chainConfig, lxps.engine); err != nil {
		return nil, err
	}
	lxps.bloomIndexer.Start(lxps.blockchain)
	// Rewind the chain in case of an incompatible config upgrade.
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		lxps.blockchain.SetHead(compat.RewindTo)
		rawdb.WriteChainConfig(chainDb.GlobalTable(), genesisHash, chainConfig)
	}

	lxps.txPool = light.NewTxPool(lxps.chainConfig, lxps.blockchain, lxps.relay)
	if lxps.protocolManager, err = NewProtocolManager(lxps.chainConfig, true, ClientProtocolVersions, config.NetworkId, lxps.eventMux, lxps.peers, lxps.blockchain, nil, chainDb, lxps.odr, lxps.relay, quitSync, &lxps.wg); err != nil {
		return nil, err
	}
	lxps.ApiBackend = &LxsApiBackend{
		xps: lxps,
		gpo: nil,
	}
	if g := lxps.config.Genesis; g != nil {
		lxps.ApiBackend.initialSupply = g.Alloc.Total()
	}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.MinerGasPrice
	}
	lxps.ApiBackend.gpo = gasprice.NewOracle(lxps.ApiBackend, gpoParams)
	return lxps, nil
}

func lxsTopic(genesisHash common.Hash, protocolVersion uint) discv5.Topic {
	var name string
	switch protocolVersion {
	case lpv1:
		name = "LXS"
	case lpv2:
		name = "LXS2"
	default:
		panic(nil)
	}
	return discv5.Topic(name + "@" + common.Bytes2Hex(genesisHash.Bytes()[0:8]))
}

type LightDummyAPI struct{}

// Xpsbase is the address that mining rewards will be send to
func (s *LightDummyAPI) Xpsbase() (common.Address, error) {
	return common.Address{}, fmt.Errorf("not supported")
}

// Coinbase is the address that mining rewards will be send to (alias for Xpsbase)
func (s *LightDummyAPI) Coinbase() (common.Address, error) {
	return common.Address{}, fmt.Errorf("not supported")
}

// Hashrate returns the POW hashrate
func (s *LightDummyAPI) Hashrate() hexutil.Uint {
	return 0
}

// Mining returns an indication if this node is currently mining.
func (s *LightDummyAPI) Mining() bool {
	return false
}

// APIs returns the collection of RPC services the xpayments package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *LightXPS) APIs() []rpc.API {
	return append(xpsapi.GetAPIs(s.ApiBackend), []rpc.API{
		{
			Namespace: "xps",
			Version:   "1.0",
			Service:   &LightDummyAPI{},
			Public:    true,
		}, {
			Namespace: "xps",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "xps",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.ApiBackend, true),
			Public:    true,
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},
	}...)
}

func (s *LightXPS) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *LightXPS) BlockChain() *light.LightChain      { return s.blockchain }
func (s *LightXPS) TxPool() *light.TxPool              { return s.txPool }
func (s *LightXPS) Engine() consensus.Engine           { return s.engine }
func (s *LightXPS) LxsVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *LightXPS) Downloader() *downloader.Downloader { return s.protocolManager.downloader }
func (s *LightXPS) EventMux() *core.InterfaceFeed      { return s.eventMux }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *LightXPS) Protocols() []p2p.Protocol {
	return s.protocolManager.SubProtocols
}

// Start implements node.Service, starting all internal goroutines needed by the
// xPayments protocol implementation.
func (s *LightXPS) Start(srvr *p2p.Server) error {
	s.startBloomHandlers()
	log.Warn("Light client mode is an experimental feature")
	s.netRPCService = xpsapi.NewPublicNetAPI(srvr, s.networkId)
	// clients are searching for the first advertised protocol in the list
	protocolVersion := AdvertiseProtocolVersions[0]
	s.serverPool.start(srvr, lxsTopic(s.blockchain.Genesis().Hash(), protocolVersion))
	s.protocolManager.Start(s.config.LightPeers)
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// xPayments protocol.
func (s *LightXPS) Stop() error {
	s.odr.Stop()
	if s.bloomIndexer != nil {
		if err := s.bloomIndexer.Close(); err != nil {
			log.Error("cannot close bloom indexer", "err", err)
		}
	}
	if s.chtIndexer != nil {
		if err := s.chtIndexer.Close(); err != nil {
			log.Error("cannot close chain indexer", "err", err)
		}
	}
	if s.bloomTrieIndexer != nil {
		if err := s.bloomTrieIndexer.Close(); err != nil {
			log.Error("cannot close bloom trie indexer", "err", err)
		}
	}
	s.blockchain.Stop()
	s.protocolManager.Stop()
	s.txPool.Stop()

	s.eventMux.Close()

	time.Sleep(time.Millisecond * 200)
	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}
