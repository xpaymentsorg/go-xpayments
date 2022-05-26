package XPSxlending

import (
	"context"
	"errors"
	"sync"
	"time"
)

// List of errors
var (
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicXPSXLendingAPI provides the XPSX RPC service that can be
// use publicly without security implications.
type PublicXPSXLendingAPI struct {
	t        *Lending
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicXPSXLendingAPI create a new RPC XPSX service.
func NewPublicXPSXLendingAPI(t *Lending) *PublicXPSXLendingAPI {
	api := &PublicXPSXLendingAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Lending sub-protocol version.
func (api *PublicXPSXLendingAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
