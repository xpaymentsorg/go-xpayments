package engine_v2

import (
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/consensus"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/ethdb"
	"github.com/xpaymentsorg/go-xpayments/params"
)

type XPoS_v2 struct {
	config *params.XPoSConfig // Consensus engine configuration parameters
	db     ethdb.Database     // Database to store and retrieve snapshot checkpoints
}

func New(config *params.XPoSConfig, db ethdb.Database) *XPoS_v2 {
	return &XPoS_v2{
		config: config,
		db:     db,
	}
}

func NewFaker(db ethdb.Database, config *params.XPoSConfig) *XPoS_v2 {
	var fakeEngine *XPoS_v2
	// Set any missing consensus parameters to their defaults
	conf := config

	// Allocate the snapshot caches and create the engine
	fakeEngine = &XPoS_v2{
		config: conf,
		db:     db,
	}
	return fakeEngine
}

func (consensus *XPoS_v2) Author(header *types.Header) (common.Address, error) {
	return common.Address{}, nil
}

func (consensus *XPoS_v2) VerifyHeader(chain consensus.ChainReader, header *types.Header, fullVerify bool) error {
	return nil
}
