package mock

import (
	"testing"
	"time"
)

func StartMockServer(t testing.TB, dataDir string) *Server {
	t.Helper()

	cfg := Config{
		GRPCAddr:    "127.0.0.1:0",
		GatewayAddr: "127.0.0.1:0",
		DataDir:     dataDir,
	}

	server, err := NewServer(cfg)
	if err != nil {
		t.Fatalf("failed to create mock server: %v", err)
	}

	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	return server
}

