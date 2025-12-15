// File: internal/modulex/service.go
// AiGen START
package modulex

import (
	"context"
	"fmt"

	"gen-poc/internal/ports"
)

// Service encapsulates Module X domain behavior.
type Service struct {
	y ports.ModuleYPort
}

// NewService wires the Module Y dependency into Module X.
func NewService(y ports.ModuleYPort) *Service {
	return &Service{y: y}
}

// Execute demonstrates orchestration across Module X and Module Y.
func (s *Service) Execute(ctx context.Context, input string) (string, error) {
	prep := fmt.Sprintf("x-pre(%s)", input)

	resultFromY, err := s.y.Compute(ctx, prep)
	if err != nil {
		return "", fmt.Errorf("modulex: module y call failed: %w", err)
	}

	response := fmt.Sprintf("%s -> %s -> x-post", prep, resultFromY)
	return response, nil
}

// AiGen END
