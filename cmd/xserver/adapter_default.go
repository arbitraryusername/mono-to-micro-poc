// File: cmd/xserver/adapter_default.go
// AiGen START

//  NOTE: this file only exists to prevent IDE/linter errors

// below is the build tag for default mode.
//go:build !local && !remote

package main

import (
	"fmt"

	"gen-poc/internal/ports"
)

// buildModuleYAdapter is a stub that should never be called at runtime.
// This file exists only to satisfy the IDE/linter when no build tag is specified.
// At build time, you MUST use either -tags=local or -tags=remote, otherwise you will get a build error.
func buildModuleYAdapter() (ports.ModuleYPort, error) {
	return nil, fmt.Errorf("FATAL: xserver must be built with either -tags=local or -tags=remote")
}

// AiGen END
