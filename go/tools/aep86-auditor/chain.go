package main

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	gogogrpc "github.com/cosmos/gogoproto/grpc"
	"google.golang.org/grpc"

	verificationv1 "pkg.akt.dev/go/node/verification/v1"
	"pkg.akt.dev/go/sdkutil"
)

type chainFactsResult struct {
	BlockHeight           string
	ProviderPubKey        cryptotypes.PubKey
	ProviderPubKeyAddress string
	Facts                 map[string]any
	Warnings              []string
}

func collectChainFacts(ctx context.Context, cfg collectConfig, provider string) (*chainFactsResult, error) {
	conn, err := dialGRPC(ctx, grpcDialConfig{
		endpoint:      cfg.chainGRPC,
		insecure:      !cfg.chainGRPCTLS,
		skipTLSVerify: cfg.chainGRPCSkipTLSVerify,
		serverName:    cfg.chainGRPCTLSServerName,
	})
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	result := &chainFactsResult{
		BlockHeight: "0",
		Facts:       map[string]any{},
	}

	nodeInfo, err := cmtservice.NewServiceClient(conn).GetNodeInfo(ctx, &cmtservice.GetNodeInfoRequest{})
	if err != nil {
		result.Warnings = append(result.Warnings, "chain node info query failed: "+err.Error())
	} else if info := nodeInfo.GetDefaultNodeInfo(); info != nil {
		result.Facts["node_network"] = info.Network
	}

	latestBlock, err := cmtservice.NewServiceClient(conn).GetLatestBlock(ctx, &cmtservice.GetLatestBlockRequest{})
	if err != nil {
		result.Warnings = append(result.Warnings, "chain latest block query failed: "+err.Error())
	} else if block := latestBlock.GetBlock(); block != nil {
		result.BlockHeight = fmt.Sprintf("%d", block.Header.Height)
	}

	pubKey, err := queryAccountPubKey(ctx, conn, provider)
	if err != nil {
		return result, err
	}
	result.ProviderPubKey = pubKey
	result.ProviderPubKeyAddress = sdk.AccAddress(pubKey.Address()).String()
	result.Facts["provider_pubkey_address"] = result.ProviderPubKeyAddress
	result.Facts["provider_pubkey_type"] = fmt.Sprintf("%T", pubKey)

	queryVerificationFacts(ctx, conn, provider, result)

	return result, nil
}

func queryAccountPubKey(ctx context.Context, conn grpc.ClientConnInterface, provider string) (cryptotypes.PubKey, error) {
	resp, err := authtypes.NewQueryClient(conn).Account(ctx, &authtypes.QueryAccountRequest{Address: provider})
	if err != nil {
		return nil, fmt.Errorf("query provider account %q: %w", provider, err)
	}
	if resp.GetAccount() == nil {
		return nil, fmt.Errorf("provider account %q not found", provider)
	}

	encCfg := sdkutil.MakeEncodingConfig(auth.AppModuleBasic{})
	var account authtypes.AccountI
	if err := encCfg.InterfaceRegistry.UnpackAny(resp.GetAccount(), &account); err != nil {
		return nil, fmt.Errorf("unpack provider account %q: %w", provider, err)
	}
	if account.GetPubKey() == nil {
		return nil, fmt.Errorf("provider account %q has no on-chain public key", provider)
	}

	return account.GetPubKey(), nil
}

func queryVerificationFacts(ctx context.Context, conn gogogrpc.ClientConn, provider string, result *chainFactsResult) {
	query := verificationv1.NewQueryClient(conn)

	escrows, err := query.ProviderAuditEscrows(ctx, &verificationv1.QueryProviderAuditEscrowsRequest{
		Provider: provider,
	})
	if err != nil {
		result.Warnings = append(result.Warnings, "verification ProviderAuditEscrows query failed: "+err.Error())
	} else {
		result.Facts["provider_audit_escrow_count"] = len(escrows.GetEscrows())
	}

	bond, err := query.ProviderBond(ctx, &verificationv1.QueryProviderBondRequest{
		Provider: provider,
	})
	if err != nil {
		result.Warnings = append(result.Warnings, "verification ProviderBond query failed: "+err.Error())
	} else {
		result.Facts["provider_bond"] = bond.GetBond()
		result.Facts["provider_bond_required_for_current_tier"] = bond.GetRequiredForCurrentTier()
	}

	snapshot, err := query.ProviderSnapshot(ctx, &verificationv1.QueryProviderSnapshotRequest{
		Provider: provider,
	})
	if err != nil {
		result.Warnings = append(result.Warnings, "verification ProviderSnapshot query failed: "+err.Error())
	} else {
		result.Facts["provider_snapshot"] = snapshot.GetSnapshot()
	}
}
