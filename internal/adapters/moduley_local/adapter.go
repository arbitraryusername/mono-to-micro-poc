// File: internal/adapters/moduley_local/adapter.go
// AiGen START
package moduley_local

import (
	"context"
	"gen-poc/internal/moduley"
	"log"
)

// Adapter bridges Module X to the local Module Y implementation.
type Adapter struct {
	svc *moduley.Service
}

// NewAdapter creates the local adapter around Module Y.
func NewAdapter(svc *moduley.Service) *Adapter {
	return &Adapter{svc: svc}
}

// Compute proxies calls into Module Y.
func (a *Adapter) Compute(ctx context.Context, input string) (string, error) {
	log.Printf("[LOCAL_ADAPTER] Starting Compute call with input=%q", input)
	return a.svc.Compute(ctx, input)
}

// AiGen END
