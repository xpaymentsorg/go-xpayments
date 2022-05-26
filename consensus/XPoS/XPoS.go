// Copyright (c) 2021 XDPoSChain
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Package XPoS is the adaptor for different consensus engine.
package XPoS

import (
	"math/big"

	lru "github.com/hashicorp/golang-lru"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/consensus"
	"github.com/xpaymentsorg/go-xpayments/consensus/XPoS/engines/engine_v1"
	"github.com/xpaymentsorg/go-xpayments/consensus/XPoS/engines/engine_v2"
	"github.com/xpaymentsorg/go-xpayments/consensus/XPoS/utils"
	"github.com/xpaymentsorg/go-xpayments/consensus/clique"
	"github.com/xpaymentsorg/go-xpayments/core/state"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/ethdb"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/params"
	"github.com/xpaymentsorg/go-xpayments/rpc"
)

func SigHash(header *types.Header) (hash common.Hash) {
	return utils.SigHash(header)
}

// XPoS is the delegated-proof-of-stake consensus engine proposed to support the
// Ethereum testnet following the Ropsten attacks.
type XPoS struct {
	config *params.XPoSConfig // Consensus engine configuration parameters
	db     ethdb.Database     // Database to store and retrieve snapshot checkpoints

	// Transaction cache, only make sense for adaptor level
	signingTxsCache *lru.Cache

	// Trading and lending service
	GetXPSXService    func() utils.TradingService
	GetLendingService func() utils.LendingService

	// The exact consensus engine with different versions
	EngineV1 engine_v1.XPoS_v1
	EngineV2 engine_v2.XPoS_v2
}

// New creates a XPoS delegated-proof-of-stake consensus engine with the initial
// signers set to the ones provided by the user.
func New(config *params.XPoSConfig, db ethdb.Database) *XPoS {
	// Set any missing consensus parameters to their defaults
	conf := *config
	if conf.Epoch == 0 {
		conf.Epoch = utils.EpochLength
	}

	// Allocate the snapshot caches and create the engine
	signingTxsCache, _ := lru.New(utils.BlockSignersCacheLimit)

	return &XPoS{
		config: &conf,
		db:     db,

		signingTxsCache: signingTxsCache,
		EngineV1:        *engine_v1.New(&conf, db),
		EngineV2:        *engine_v2.New(&conf, db),
	}
}

// NewFullFaker creates an ethash consensus engine with a full fake scheme that
// accepts all blocks as valid, without checking any consensus rules whatsoever.
func NewFaker(db ethdb.Database, chainConfig *params.ChainConfig) *XPoS {
	var fakeEngine *XPoS
	// Set any missing consensus parameters to their defaults
	conf := params.TestXPoSMockChainConfig.XPoS
	if chainConfig != nil {
		conf = chainConfig.XPoS
	}

	// Allocate the snapshot caches and create the engine
	signingTxsCache, _ := lru.New(utils.BlockSignersCacheLimit)

	fakeEngine = &XPoS{
		config: conf,
		db:     db,

		signingTxsCache: signingTxsCache,
		EngineV1:        *engine_v1.NewFaker(db, conf),
		EngineV2:        *engine_v2.NewFaker(db, conf),
	}
	return fakeEngine
}

/*
	Eth Consensus engine interface implementation
*/
// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (x *XPoS) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "XPoS",
		Version:   "1.0",
		Service:   &API{chain: chain, XPoS: x},
		Public:    true,
	}}
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (x *XPoS) Author(header *types.Header) (common.Address, error) {
	switch x.config.BlockConsensusVersion(header.Number) {
	case params.ConsensusEngineVersion2:
		return x.EngineV2.Author(header)
	default: // Default "v1"
		return x.EngineV1.Author(header)
	}
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (x *XPoS) VerifyHeader(chain consensus.ChainReader, header *types.Header, fullVerify bool) error {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.VerifyHeader(chain, header, fullVerify)
	}
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (x *XPoS) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, fullVerifies []bool) (chan<- struct{}, <-chan error) {
	// TODO: (Hashlab) This funciton is a special case
	return x.EngineV1.VerifyHeaders(chain, headers, fullVerifies)
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (x *XPoS) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	switch x.config.BlockConsensusVersion(block.Number()) {
	default: // Default "v1"
		return x.EngineV1.VerifyUncles(chain, block)
	}
}

// VerifySeal implements consensus.Engine, checking whether the signature contained
// in the header satisfies the consensus protocol requirements.
func (x *XPoS) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.VerifySeal(chain, header)
	}
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (x *XPoS) Prepare(chain consensus.ChainReader, header *types.Header) error {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.Prepare(chain, header)
	}
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (x *XPoS) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, parentState *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.Finalize(chain, header, state, parentState, txs, uncles, receipts)
	}
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (x *XPoS) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	switch x.config.BlockConsensusVersion(block.Number()) {
	default: // Default "v1"
		return x.EngineV1.Seal(chain, block, stop)
	}
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func (x *XPoS) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	switch x.config.BlockConsensusVersion(parent.Number) {
	default: // Default "v1"
		return x.EngineV1.CalcDifficulty(chain, time, parent)
	}
}

/*
	XPS specific methods
*/

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (x *XPoS) Authorize(signer common.Address, signFn clique.SignerFn) {
	// Authorize each consensus individually
	x.EngineV1.Authorize(signer, signFn)
}

func (x *XPoS) GetPeriod() uint64 {
	return x.config.Period
}

func (x *XPoS) IsAuthorisedAddress(header *types.Header, chain consensus.ChainReader, address common.Address) bool {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.IsAuthorisedAddress(header, chain, address)
	}
}

func (x *XPoS) GetMasternodes(chain consensus.ChainReader, header *types.Header) []common.Address {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.GetMasternodes(chain, header)
	}
}

func (x *XPoS) YourTurn(chain consensus.ChainReader, parent *types.Header, signer common.Address) (int, int, int, bool, error) {
	switch x.config.BlockConsensusVersion(parent.Number) {
	default: // Default "v1"
		return x.EngineV1.YourTurn(chain, parent, signer)
	}
}

func (x *XPoS) GetValidator(creator common.Address, chain consensus.ChainReader, header *types.Header) (common.Address, error) {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.GetValidator(creator, chain, header)
	}
}

func (x *XPoS) UpdateMasternodes(chain consensus.ChainReader, header *types.Header, ms []utils.Masternode) error {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.UpdateMasternodes(chain, header, ms)
	}
}

func (x *XPoS) RecoverSigner(header *types.Header) (common.Address, error) {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.RecoverSigner(header)
	}
}

func (x *XPoS) RecoverValidator(header *types.Header) (common.Address, error) {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.RecoverValidator(header)
	}
}

// Get master nodes over extra data of previous checkpoint block.
func (x *XPoS) GetMasternodesFromCheckpointHeader(preCheckpointHeader *types.Header, n, e uint64) []common.Address {
	switch x.config.BlockConsensusVersion(preCheckpointHeader.Number) {
	default: // Default "v1"
		return x.EngineV1.GetMasternodesFromCheckpointHeader(preCheckpointHeader, n, e)
	}
}

// Same DB across all consensus engines
func (x *XPoS) GetDb() ethdb.Database {
	return x.db
}

func (x *XPoS) GetSnapshot(chain consensus.ChainReader, header *types.Header) (*utils.PublicApiSnapshot, error) {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		sp, err := x.EngineV1.GetSnapshot(chain, header)
		// Convert to a standard PublicApiSnapshot type, otherwise it's a breaking change to API
		return &utils.PublicApiSnapshot{
			Number:  sp.Number,
			Hash:    sp.Hash,
			Signers: sp.Signers,
			Recents: sp.Recents,
			Votes:   sp.Votes,
			Tally:   sp.Tally,
		}, err
	}
}

func (x *XPoS) GetAuthorisedSignersFromSnapshot(chain consensus.ChainReader, header *types.Header) ([]common.Address, error) {
	switch x.config.BlockConsensusVersion(header.Number) {
	default: // Default "v1"
		return x.EngineV1.GetAuthorisedSignersFromSnapshot(chain, header)
	}
}

/**
Caching
*/

// Cache signing transaction data into BlockSingers cache object
func (x *XPoS) CacheNoneTIPSigningTxs(header *types.Header, txs []*types.Transaction, receipts []*types.Receipt) []*types.Transaction {
	signTxs := []*types.Transaction{}
	for _, tx := range txs {
		if tx.IsSigningTransaction() {
			var b uint
			for _, r := range receipts {
				if r.TxHash == tx.Hash() {
					if len(r.PostState) > 0 {
						b = types.ReceiptStatusSuccessful
					} else {
						b = r.Status
					}
					break
				}
			}

			if b == types.ReceiptStatusFailed {
				continue
			}

			signTxs = append(signTxs, tx)
		}
	}

	log.Debug("Save tx signers to cache", "hash", header.Hash().String(), "number", header.Number, "len(txs)", len(signTxs))
	x.signingTxsCache.Add(header.Hash(), signTxs)

	return signTxs
}

// Cache
func (x *XPoS) CacheSigningTxs(hash common.Hash, txs []*types.Transaction) []*types.Transaction {
	signTxs := []*types.Transaction{}
	for _, tx := range txs {
		if tx.IsSigningTransaction() {
			signTxs = append(signTxs, tx)
		}
	}
	log.Debug("Save tx signers to cache", "hash", hash.String(), "len(txs)", len(signTxs))
	x.signingTxsCache.Add(hash, signTxs)
	return signTxs
}

func (x *XPoS) GetCachedSigningTxs(hash common.Hash) (interface{}, bool) {
	return x.signingTxsCache.Get(hash)
}
