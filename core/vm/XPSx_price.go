package vm

import (
	"github.com/xpaymentsorg/go-xpayments/XPSx/tradingstate"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/params"
)

const XPSXPriceNumberOfBytesReturn = 32

// XPSxPrice implements a pre-compile contract to get token price in XPSx

type XPSxLastPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}
type XPSxEpochPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}

func (t *XPSxLastPrice) RequiredGas(input []byte) uint64 {
	return params.XPSXPriceGas
}

func (t *XPSxLastPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:])  // 20 bytes from 45-64
		price := t.tradingStateDB.GetLastPrice(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetLastPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), XPSXPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, XPSXPriceNumberOfBytesReturn), nil
}

func (t *XPSxLastPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}

func (t *XPSxEpochPrice) RequiredGas(input []byte) uint64 {
	return params.XPSXPriceGas
}

func (t *XPSxEpochPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:])  // 20 bytes from 45-64
		price := t.tradingStateDB.GetMediumPriceBeforeEpoch(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetEpochPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), XPSXPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, XPSXPriceNumberOfBytesReturn), nil
}

func (t *XPSxEpochPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}
