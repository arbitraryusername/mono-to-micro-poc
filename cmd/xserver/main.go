// File: cmd/xserver/main.go
// AiGen START
package main

import (
	"log"
	"net/http"
	"os"

	"gen-poc/internal/modulex"
)

// main bootstraps Module X's HTTP surface and selects the Module Y adapter.
func main() {
	yAdapter, err := buildModuleYAdapter()
	if err != nil {
		log.Fatalf("configure module y: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/x", func(w http.ResponseWriter, r *http.Request) {
		input := r.URL.Query().Get("input")
		if input == "" {
			input = "demo"
		}

		moduleX := modulex.NewService(yAdapter)

		result, err := moduleX.Execute(r.Context(), input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(result))
	})

	addr := os.Getenv("XSERVER_ADDR")
	if addr == "" {
		log.Fatalf("XSERVER_ADDR environment variable must be set")
	}

	log.Printf("xserver listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// buildModuleYAdapter is implemented in adapter_local.go or adapter_remote.go
// depending on the build tag used (-tags=local or -tags=remote).

// AiGen END
