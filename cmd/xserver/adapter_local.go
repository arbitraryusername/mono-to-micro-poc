// File: cmd/xserver/adapter_local.go
// AiGen START

//go:build local

package main

import (
	"gen-poc/internal/adapters/moduley_local"
	"gen-poc/internal/moduley"
	"gen-poc/internal/ports"
)

// buildModuleYAdapter creates the local Module Y adapter.
// This version is compiled only when building with -tags=local.
func buildModuleYAdapter() (ports.ModuleYPort, error) {
	moduleY := moduley.NewService()
	return moduley_local.NewAdapter(moduleY), nil
}

// AiGen END
