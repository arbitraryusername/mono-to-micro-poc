// File: cmd/xserver/main.go
// AiGen START
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gen-poc/internal/adapters/moduley_local"
	"gen-poc/internal/adapters/moduley_remote"
	"gen-poc/internal/modulex"
	"gen-poc/internal/moduley"
	"gen-poc/internal/ports"
)

// main bootstraps Module X's HTTP surface and selects the Module Y adapter.
func main() {
	yAdapter, err := buildModuleYAdapter()
	if err != nil {
		log.Fatalf("configure module y: %v", err)
	}

	moduleX := modulex.NewService(yAdapter)

	mux := http.NewServeMux()
	mux.HandleFunc("/x", func(w http.ResponseWriter, r *http.Request) {
		input := r.URL.Query().Get("input")
		if input == "" {
			input = "demo"
		}

		result, err := moduleX.Execute(r.Context(), input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(result))
	})

	log.Println("xserver listening on port 8001")

	if err := http.ListenAndServe(":8001", mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// buildModuleYAdapter chooses the Module Y adapter based on configuration.
func buildModuleYAdapter() (ports.ModuleYPort, error) {
	mode := strings.ToLower(os.Getenv("MODULEY_MODE"))
	switch mode {
	case "remote":
		baseURL := os.Getenv("MODULEY_URL")
		if baseURL == "" {
			return nil, fmt.Errorf("MODULEY_URL required when MODULEY_MODE=remote")
		}

		httpClient := &http.Client{
			Timeout: 5 * time.Second,
		}

		return moduley_remote.NewAdapter(baseURL, httpClient), nil
	default:
		moduleY := moduley.NewService()
		return moduley_local.NewAdapter(moduleY), nil
	}
}

// AiGen END
