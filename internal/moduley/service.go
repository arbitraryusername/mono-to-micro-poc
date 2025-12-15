// File: internal/moduley/service.go
// AiGen START
package moduley

import (
	"context"
	"fmt"
	"strings"
)

// Service contains Module Y's domain logic.
type Service struct{}

// NewService constructs Module Y's service.
func NewService() *Service {
	return &Service{}
}

// Compute applies Module Y-specific rules to the input.
func (s *Service) Compute(_ context.Context, input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("moduley: empty input")
	}

	transformed := strings.ToUpper(input)
	return transformed, nil
}

// AiGen END
