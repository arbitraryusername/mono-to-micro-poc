// File: internal/ports/moduley.go
// AiGen START
package ports

import "context"

// ModuleYPort defines the behavior Module X relies on from Module Y.
type ModuleYPort interface {
	Compute(ctx context.Context, input string) (string, error)
}

// AiGen END
