// Copyright 2014 The go-ethereum Authors
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

// Package xps implements the xPayments protocol.
package xps

import (
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/xpaymentsorg/go-xpayments/accounts"
	"github.com/xpaymentsorg/go-xpayments/accounts/keystore"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/consensus"
	"github.com/xpaymentsorg/go-xpayments/consensus/clique"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/core/bloombits"
	"github.com/xpaymentsorg/go-xpayments/core/rawdb"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/core/vm"
	"github.com/xpaymentsorg/go-xpayments/internal/xpsapi"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/miner"
	"github.com/xpaymentsorg/go-xpayments/node"
	"github.com/xpaymentsorg/go-xpayments/p2p"
	"github.com/xpaymentsorg/go-xpayments/params"
	"github.com/xpaymentsorg/go-xpayments/rpc"
	"github.com/xpaymentsorg/go-xpayments/xps/downloader"
	"github.com/xpaymentsorg/go-xpayments/xps/filters"
	"github.com/xpaymentsorg/go-xpayments/xps/gasprice"
)

type LxsServer interface {
	Start(srvr *p2p.Server)
	Stop()
	Protocols() []p2p.Protocol
	SetBloomBitsIndexer(bbIndexer *core.ChainIndexer)
}

// XPS implements the xPayments full node service.
type XPS struct {
	config      *Config
	chainConfig *params.ChainConfig

	// Channel for shutting down the service
	shutdownChan  chan bool    // Channel for shutting down the xpayments
	stopDbUpgrade func() error // stop chain db sequential key upgrade

	// Handlers
	txPool          *core.TxPool
	blockchain      *core.BlockChain
	protocolManager *ProtocolManager
	lxsServer       LxsServer

	// DB interfaces
	chainDb common.Database // Block chain database

	eventMux       *core.InterfaceFeed
	engine         consensus.Engine
	accountManager *accounts.Manager

	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer             // Bloom indexer operating during block imports

	ApiBackend *XpsApiBackend

	miner    *miner.Miner
	gasPrice *big.Int // nil for default/dynamic
	xpsbase  common.Address

	networkId     uint64
	netRPCService *xpsapi.PublicNetAPI

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and xpsbase)
}

func (xpayments *XPS) AddLxsServer(ls LxsServer) {
	xpayments.lxsServer = ls
	ls.SetBloomBitsIndexer(xpayments.bloomIndexer)
}

// New creates a new XPS object (including the
// initialisation of the common XPS object)
func New(sctx *node.ServiceContext, config *Config) (*XPS, error) {
	if config.SyncMode == downloader.LightSync {
		return nil, errors.New("can't run xps.XPS in light sync mode, use lxs.LightXPS")
	}
	if !config.SyncMode.IsValid() {
		return nil, fmt.Errorf("invalid sync mode %d", config.SyncMode)
	}
	chainDb, err := CreateDB(sctx, config, "chaindata")
	if err != nil {
		return nil, err
	}

	stopDbUpgrade := func() error { return nil } // upgradeDeduplicateData(chainDb)

	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlockWithOverride(chainDb, config.Genesis, config.ConstantinopleOverride)
	if _, ok := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !ok {
		return nil, genesisErr
	}
	if config.Genesis == nil {
		if genesisHash == params.MainnetGenesisHash {
			config.Genesis = core.DefaultGenesisBlock()
		}
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	if chainConfig.Clique == nil {
		return nil, fmt.Errorf("invalid configuration, clique is nil: %v", chainConfig)
	}
	xps := &XPS{
		config:         config,
		chainDb:        chainDb,
		chainConfig:    chainConfig,
		eventMux:       sctx.EventMux,
		accountManager: sctx.AccountManager,
		engine:         clique.New(chainConfig.Clique, chainDb),
		shutdownChan:   make(chan bool),
		stopDbUpgrade:  stopDbUpgrade,
		networkId:      config.NetworkId,
		gasPrice:       config.MinerGasPrice,
		xpsbase:        config.Xpsbase,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   NewBloomIndexer(chainDb, params.BloomBitsBlocks),
	}

	bcVersion := rawdb.ReadDatabaseVersion(chainDb.GlobalTable())
	var dbVer = "<nil>"
	if bcVersion != nil {
		dbVer = fmt.Sprintf("%d", *bcVersion)
	}
	log.Info("Initialising xPayments protocol", "versions", ProtocolVersions, "network", config.NetworkId, "dbversion", dbVer)

	if !config.SkipBcVersionCheck {
		if bcVersion != nil && *bcVersion > core.BlockChainVersion {
			return nil, fmt.Errorf("database version is v%d, xPayments %s only supports v%d", *bcVersion, params.Version, core.BlockChainVersion)
		} else if bcVersion == nil || *bcVersion < core.BlockChainVersion {
			log.Warn("Upgrade blockchain database version", "from", dbVer, "to", core.BlockChainVersion)
			rawdb.WriteDatabaseVersion(chainDb.GlobalTable(), core.BlockChainVersion)
		}
		rawdb.WriteDatabaseVersion(chainDb.GlobalTable(), core.BlockChainVersion)
	}
	var (
		vmConfig    = vm.Config{EnablePreimageRecording: config.EnablePreimageRecording}
		cacheConfig = &core.CacheConfig{Disabled: config.NoPruning, TrieNodeLimit: config.TrieCache, TrieTimeLimit: config.TrieTimeout}
	)
	xps.blockchain, err = core.NewBlockChain(chainDb, cacheConfig, xps.chainConfig, xps.engine, vmConfig)
	if err != nil {
		return nil, err
	}
	// Rewind the chain in case of an incompatible config upgrade.
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		if err := xps.blockchain.SetHead(compat.RewindTo); err != nil {
			log.Error("Cannot set head during chain rewind", "rewind_to", compat.RewindTo, "err", err)
		}
		rawdb.WriteChainConfig(chainDb.GlobalTable(), genesisHash, chainConfig)
	}
	xps.bloomIndexer.Start(xps.blockchain)

	if config.TxPool.Journal != "" {
		config.TxPool.Journal = sctx.ResolvePath(config.TxPool.Journal)
	}
	xps.txPool = core.NewTxPool(config.TxPool, xps.chainConfig, xps.blockchain)

	if xps.protocolManager, err = NewProtocolManager(xps.chainConfig, config.SyncMode, config.NetworkId, xps.eventMux, xps.txPool, xps.engine, xps.blockchain, chainDb); err != nil {
		return nil, err
	}
	xps.miner = miner.New(xps, xps.chainConfig, xps.EventMux(), xps.engine, config.MinerRecommit, config.MinerGasFloor, config.MinerGasCeil, xps.isLocalBlock)
	if err := xps.miner.SetExtra(makeExtraData(config.MinerExtraData)); err != nil {
		log.Error("Cannot set extra chain data", "err", err)
	}

	xps.ApiBackend = &XpsApiBackend{xps: xps}
	if g := xps.config.Genesis; g != nil {
		xps.ApiBackend.initialSupply = g.Alloc.Total()
	}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.MinerGasPrice
	}
	xps.ApiBackend.gpo = gasprice.NewOracle(xps.ApiBackend, gpoParams)

	return xps, nil
}

// Example: 2.0.73/linux-amd64/go1.10.2
var defaultExtraData []byte
var defaultExtraDataOnce sync.Once

func makeExtraData(extra []byte) []byte {
	if len(extra) == 0 {
		defaultExtraDataOnce.Do(func() {
			defaultExtraData = []byte(fmt.Sprintf("%s/%s-%s/%s", params.Version, runtime.GOOS, runtime.GOARCH, runtime.Version()))
			if uint64(len(defaultExtraData)) > params.MaximumExtraDataSize {
				defaultExtraData = defaultExtraData[:params.MaximumExtraDataSize]
			}
		})
		return defaultExtraData
	}
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		log.Warn("Miner extra data exceed limit", "extra", string(extra), "limit", params.MaximumExtraDataSize)
		extra = extra[:params.MaximumExtraDataSize]
	}
	return extra
}

// CreateDB creates the chain database.
func CreateDB(ctx *node.ServiceContext, config *Config, name string) (common.Database, error) {
	db, err := ctx.OpenDatabase(name, config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// APIs returns the collection of RPC services the xpayments package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (xpayments *XPS) APIs() []rpc.API {
	apis := xpsapi.GetAPIs(xpayments.ApiBackend)

	// Append any APIs exposed explicitly by the consensus engine
	apis = append(apis, xpayments.engine.APIs(xpayments.BlockChain())...)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicxPaymentsAPI(xpayments),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicMinerAPI(xpayments),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(xpayments.protocolManager.downloader, xpayments.eventMux),
			Public:    true,
		}, {
			Namespace: "miner",
			Version:   "1.0",
			Service:   NewPrivateMinerAPI(xpayments),
			Public:    false,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(xpayments.ApiBackend, false),
			Public:    true,
		}, {
			Namespace: "admin",
			Version:   "1.0",
			Service:   NewPrivateAdminAPI(xpayments),
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(xpayments),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(xpayments.chainConfig, xpayments),
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   xpayments.netRPCService,
			Public:    true,
		},
	}...)
}

func (xpayments *XPS) ResetWithGenesisBlock(gb *types.Block) {
	if err := xpayments.blockchain.ResetWithGenesisBlock(gb); err != nil {
		log.Error("Cannot reset with genesis block", "err", err)
	}
}

func (xpayments *XPS) Xpsbase() (eb common.Address, err error) {
	xpayments.lock.RLock()
	xpsbase := xpayments.xpsbase
	xpayments.lock.RUnlock()

	if xpsbase != (common.Address{}) {
		return xpsbase, nil
	}
	ks := xpayments.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	if wallets := ks.Wallets(); len(wallets) > 0 {
		if accounts := wallets[0].Accounts(); len(accounts) > 0 {
			xpsbase := accounts[0].Address

			xpayments.lock.Lock()
			xpayments.xpsbase = xpsbase
			xpayments.lock.Unlock()

			log.Info("Xpsbase automatically configured", "address", xpsbase)
			return xpsbase, nil
		}
	}
	return common.Address{}, fmt.Errorf("xpsbase must be explicitly specified")
}

// isLocalBlock checks whxps the specified block is mined
// by local miner accounts.
//
// We regard two types of accounts as local miner account: xpsbase
// and accounts specified via `txpool.locals` flag.
func (xpayments *XPS) isLocalBlock(block *types.Block) bool {
	author, err := xpayments.engine.Author(block.Header())
	if err != nil {
		log.Warn("Failed to retrieve block author", "number", block.NumberU64(), "hash", block.Hash(), "err", err)
		return false
	}
	// Check whxps the given address is xpsbase.
	xpayments.lock.RLock()
	xpsbase := xpayments.xpsbase
	xpayments.lock.RUnlock()
	if author == xpsbase {
		return true
	}
	// Check whxps the given address is specified by `txpool.local`
	// CLI flag.
	for _, account := range xpayments.config.TxPool.Locals {
		if account == author {
			return true
		}
	}
	return false
}

// SetXpsbase sets the mining reward address.
func (xpayments *XPS) SetXpsbase(xpsbase common.Address) {
	xpayments.lock.Lock()
	xpayments.xpsbase = xpsbase
	xpayments.lock.Unlock()

	xpayments.miner.SetXpsbase(xpsbase)
}

// StartMining starts the miner with the given number of CPU threads. If mining
// is already running, this method adjust the number of threads allowed to use
// and updates the minimum price required by the transaction pool.
func (xpayments *XPS) StartMining(threads int) error {
	// Update the thread count within the consensus engine
	type threaded interface {
		SetThreads(threads int)
	}
	if th, ok := xpayments.engine.(threaded); ok {
		log.Info("Updated mining threads", "threads", threads)
		if threads == 0 {
			threads = -1 // Disable the miner from within
		}
		th.SetThreads(threads)
	}
	// If the miner was not running, initialize it
	if !xpayments.IsMining() {
		// Propagate the initial price point to the transaction pool
		xpayments.lock.RLock()
		price := xpayments.gasPrice
		xpayments.lock.RUnlock()
		xpayments.txPool.SetGasPrice(price)

		// Configure the local mining address
		eb, err := xpayments.Xpsbase()
		if err != nil {
			log.Error("Cannot start mining without xpsbase", "err", err)
			return fmt.Errorf("xpsbase missing: %v", err)
		}
		if clique, ok := xpayments.engine.(*clique.Clique); ok {
			wallet, err := xpayments.accountManager.Find(accounts.Account{Address: eb})
			if wallet == nil || err != nil {
				log.Error("Xpsbase account unavailable locally", "err", err)
				return fmt.Errorf("signer missing: %v", err)
			}
			clique.Authorize(eb, wallet.SignData)
		}
		// If mining is started, we can disable the transaction rejection mechanism
		// introduced to speed sync times.
		atomic.StoreUint32(&xpayments.protocolManager.acceptTxs, 1)

		go xpayments.miner.Start(eb)
	}
	return nil
}

// StopMining terminates the miner, both at the consensus engine level as well as
// at the block creation level.
func (xpayments *XPS) StopMining() {
	// Update the thread count within the consensus engine
	type threaded interface {
		SetThreads(threads int)
	}
	if th, ok := xpayments.engine.(threaded); ok {
		th.SetThreads(-1)
	}
	// Stop the block creating itself
	xpayments.miner.Stop()
}

func (xpayments *XPS) IsMining() bool      { return xpayments.miner.Mining() }
func (xpayments *XPS) Miner() *miner.Miner { return xpayments.miner }

func (xpayments *XPS) AccountManager() *accounts.Manager { return xpayments.accountManager }
func (xpayments *XPS) BlockChain() *core.BlockChain      { return xpayments.blockchain }
func (xpayments *XPS) TxPool() *core.TxPool              { return xpayments.txPool }
func (xpayments *XPS) EventMux() *core.InterfaceFeed     { return xpayments.eventMux }
func (xpayments *XPS) Engine() consensus.Engine          { return xpayments.engine }
func (xpayments *XPS) ChainDb() common.Database          { return xpayments.chainDb }
func (xpayments *XPS) IsListening() bool                 { return true } // Always listening
func (xpayments *XPS) XpsVersion() int                   { return int(xpayments.protocolManager.SubProtocols[0].Version) }
func (xpayments *XPS) NetVersion() uint64                { return xpayments.networkId }
func (xpayments *XPS) Downloader() *downloader.Downloader {
	return xpayments.protocolManager.downloader
}

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (xpayments *XPS) Protocols() []p2p.Protocol {
	if xpayments.lxsServer == nil {
		return xpayments.protocolManager.SubProtocols
	}
	return append(xpayments.protocolManager.SubProtocols, xpayments.lxsServer.Protocols()...)
}

// Start implements node.Service, starting all internal goroutines needed by the
// xPayments protocol implementation.
func (xpayments *XPS) Start(srvr *p2p.Server) error {
	// Start the bloom bits servicing goroutines
	xpayments.startBloomHandlers()

	// Start the RPC service
	xpayments.netRPCService = xpsapi.NewPublicNetAPI(srvr, xpayments.NetVersion())

	// Figure out a max peers count based on the server limits
	maxPeers := srvr.MaxPeers
	if xpayments.config.LightServ > 0 {
		if xpayments.config.LightPeers >= srvr.MaxPeers {
			return fmt.Errorf("invalid peer config: light peer count (%d) >= total peer count (%d)", xpayments.config.LightPeers, srvr.MaxPeers)
		}
		maxPeers -= xpayments.config.LightPeers
	}
	// Start the networking layer and the light server if requested
	xpayments.protocolManager.Start(maxPeers)
	if xpayments.lxsServer != nil {
		xpayments.lxsServer.Start(srvr)
	}
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// xPayments protocol.
func (xpayments *XPS) Stop() error {
	xpayments.bloomIndexer.Close()
	xpayments.blockchain.Stop()
	xpayments.protocolManager.Stop()
	if xpayments.lxsServer != nil {
		xpayments.lxsServer.Stop()
	}
	xpayments.txPool.Stop()
	xpayments.miner.Stop()
	xpayments.eventMux.Close()

	xpayments.chainDb.Close()
	close(xpayments.shutdownChan)

	return nil
}
