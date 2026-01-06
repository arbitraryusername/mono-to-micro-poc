// File: cmd/xserver/adapter_remote.go
// AiGen START

//go:build remote

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"gen-poc/internal/adapters/moduley_remote"
	"gen-poc/internal/ports"
)

// buildModuleYAdapter creates the remote Module Y adapter.
// This version is compiled only when building with -tags=remote.
func buildModuleYAdapter() (ports.ModuleYPort, error) {
	baseURL := os.Getenv("MODULEY_URL")
	if baseURL == "" {
		return nil, fmt.Errorf("MODULEY_URL environment variable required for remote mode")
	}

	// TODO: this doesn't have any retry or circuit breaker, etc that would be needed for a real deployed environment
	httpClient := &http.Client{
		Timeout: 1000 * time.Second, // very large timeout to account for stopping on breakpoints
	}

	return moduley_remote.NewAdapter(baseURL, httpClient), nil
}

// AiGen END
