// AiGen START

// below is the build tag for remote mode (two running binaries with RPC communication)
//go:build remote

package moduley_remote

import (
	"context"
	"fmt"
	"log"
	"net/http"

	moduleyv1 "gen-poc/gen/moduley/v1"
	moduleyv1connect "gen-poc/gen/moduley/v1/moduleyv1connect"

	"connectrpc.com/connect"
)

// Adapter invokes Module Y via ConnectRPC.
type Adapter struct {
	client moduleyv1connect.ModuleYServiceClient
}

// NewAdapter builds a remote adapter that targets the provided base URL.
func NewAdapter(baseURL string, httpClient *http.Client, opts ...connect.ClientOption) *Adapter {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := moduleyv1connect.NewModuleYServiceClient(httpClient, baseURL, opts...)
	return &Adapter{client: client}
}

// Compute converts the domain call into a ConnectRPC invocation.
func (a *Adapter) Compute(ctx context.Context, input string) (string, error) {
	log.Printf("[REMOTE_ADAPTER] Starting RPC Compute call with input=%q", input)

	req := connect.NewRequest(&moduleyv1.ComputeRequest{Input: input})

	resp, err := a.client.Compute(ctx, req)
	if err != nil {
		return "", fmt.Errorf("moduley_remote: compute failed: %w", err)
	}

	return resp.Msg.GetOutput(), nil
}

// AiGen END
