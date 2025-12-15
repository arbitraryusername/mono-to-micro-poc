// AiGen START
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	moduleyv1 "gen-poc/gen/moduley/v1"
	moduleyv1connect "gen-poc/gen/moduley/v1/moduleyv1connect"
	"gen-poc/internal/moduley"
)

// main starts Module Y's ConnectRPC server.
func main() {
	moduleY := moduley.NewService()
	handler := &rpcHandler{svc: moduleY}

	mux := http.NewServeMux()
	path, connectHandler := moduleyv1connect.NewModuleYServiceHandler(handler)
	mux.Handle(path, connectHandler)

	addr := os.Getenv("YSERVER_ADDR")
	if addr == "" {
		addr = ":8081"
	}

	log.Printf("yserver listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("yserver failed: %v", err)
	}
}

// rpcHandler adapts Module Y's domain service to the RPC contract.
type rpcHandler struct {
	svc *moduley.Service
}

// Compute satisfies the ConnectRPC handler interface by delegating to Module Y.
func (h *rpcHandler) Compute(ctx context.Context, req *connect.Request[moduleyv1.ComputeRequest]) (*connect.Response[moduleyv1.ComputeResponse], error) {
	result, err := h.svc.Compute(ctx, req.Msg.GetInput())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&moduleyv1.ComputeResponse{
		Output: result,
	}), nil
}

// AiGen END
