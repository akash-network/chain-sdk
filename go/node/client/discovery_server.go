package client

import (
	"context"

	gogogrpc "github.com/cosmos/gogoproto/grpc"
)

type discoveryServer struct {
	UnimplementedDiscoveryServer
	registry *VersionRegistry
}

var _ DiscoveryServer = (*discoveryServer)(nil)

// NewDiscoveryServer creates a new gRPC DiscoveryServer backed by the given registry.
// Passing a nil registry panics.
func NewDiscoveryServer(registry *VersionRegistry) DiscoveryServer {
	if registry == nil {
		panic("client: NewDiscoveryServer called with nil registry")
	}
	return &discoveryServer{registry: registry}
}

func (s *discoveryServer) GetInfo(_ context.Context, _ *GetInfoRequest) (*GetInfoResponse, error) {
	return &GetInfoResponse{
		Info: s.registry.ToAkash(),
	}, nil
}

// RegisterDiscoveryService registers the Discovery gRPC service on the given server.
// Accepts gogogrpc.Server so it can be called from within Cosmos SDK's
// RegisterGRPCServerWithSkipCheckHeader or with a concrete *grpc.Server.
func RegisterDiscoveryService(server gogogrpc.Server, registry *VersionRegistry) {
	RegisterDiscoveryServer(server, NewDiscoveryServer(registry))
}
