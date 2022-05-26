package XPSx

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	LimitThresholdOrderNonceInQueue = 100
)

// List of errors
var (
	ErrNoTopics          = errors.New("missing topic(s)")
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicXPSXAPI provides the XPSX RPC service that can be
// use publicly without security implications.
type PublicXPSXAPI struct {
	t        *XPSX
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicXPSXAPI create a new RPC XPSX service.
func NewPublicXPSXAPI(t *XPSX) *PublicXPSXAPI {
	api := &PublicXPSXAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the XPSX sub-protocol version.
func (api *PublicXPSXAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
