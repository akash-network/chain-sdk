package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/url"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcDialConfig struct {
	endpoint      string
	insecure      bool
	skipTLSVerify bool
	serverName    string
}

func dialGRPC(ctx context.Context, cfg grpcDialConfig) (*grpc.ClientConn, error) {
	target, err := normalizeEndpoint(cfg.endpoint)
	if err != nil {
		return nil, err
	}

	var opts []grpc.DialOption
	if cfg.insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		tlsCfg := &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: cfg.skipTLSVerify, //nolint:gosec // explicit devnet flag
			ServerName:         cfg.serverName,
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg)))
	}

	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func normalizeEndpoint(endpoint string) (string, error) {
	endpoint = strings.TrimSpace(endpoint)
	if endpoint == "" {
		return "", fmt.Errorf("empty grpc endpoint")
	}

	if strings.Contains(endpoint, "://") {
		parsed, err := url.Parse(endpoint)
		if err != nil {
			return "", err
		}
		if parsed.Host == "" {
			return "", fmt.Errorf("endpoint %q is missing host", endpoint)
		}
		return parsed.Host, nil
	}

	return endpoint, nil
}
