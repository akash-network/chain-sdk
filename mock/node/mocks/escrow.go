package mocks

import (
	"context"

	escrowtypesv1 "pkg.akt.dev/go/node/escrow/types/v1"
	escrowv1 "pkg.akt.dev/go/node/escrow/v1"
)

// EscrowData holds the escrow data
type EscrowData struct {
	Accounts []escrowtypesv1.Account
	Payments []escrowtypesv1.Payment
}

// MockEscrowQueryServer implements the escrow query server
type MockEscrowQueryServer struct {
	escrowv1.UnimplementedQueryServer
	Data EscrowData
}

// NewMockEscrowQueryServer creates a new mock escrow query server
func NewMockEscrowQueryServer(data EscrowData) *MockEscrowQueryServer {
	return &MockEscrowQueryServer{
		Data: data,
	}
}

// Accounts implements the Accounts query
func (m MockEscrowQueryServer) Accounts(ctx context.Context, req *escrowv1.QueryAccountsRequest) (*escrowv1.QueryAccountsResponse, error) {
	accounts := m.Data.Accounts

	// Filter by XID if specified
	if req.XID != "" {
		var filtered []escrowtypesv1.Account
		for _, a := range accounts {
			if a.ID.XID == req.XID {
				filtered = append(filtered, a)
			}
		}
		accounts = filtered
	}

	return &escrowv1.QueryAccountsResponse{
		Accounts: accounts,
	}, nil
}

// Payments implements the Payments query
func (m MockEscrowQueryServer) Payments(ctx context.Context, req *escrowv1.QueryPaymentsRequest) (*escrowv1.QueryPaymentsResponse, error) {
	payments := m.Data.Payments

	// Filter by XID if specified
	if req.XID != "" {
		var filtered []escrowtypesv1.Payment
		for _, p := range payments {
			if p.ID.AID.XID == req.XID {
				filtered = append(filtered, p)
			}
		}
		payments = filtered
	}

	return &escrowv1.QueryPaymentsResponse{
		Payments: payments,
	}, nil
}
