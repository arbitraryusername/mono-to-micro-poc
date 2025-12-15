// File: cmd/xserver/main.go
// AiGen START
package main

import (
	"log"
	"net/http"

	"gen-poc/internal/adapters/moduley_local"
	"gen-poc/internal/modulex"
	"gen-poc/internal/moduley"
)

func main() {
	moduleY := moduley.NewService()
	yAdapter := moduley_local.NewAdapter(moduleY)
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

	err := http.ListenAndServe(":8001", mux)
	if err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// AiGen END
