package mocks

import (
	"context"

	escrowv1 "pkg.akt.dev/go/node/escrow/v1"
)

// MockEscrowMsgServer implements the escrow message server for transactions
type MockEscrowMsgServer struct {
	escrowv1.UnimplementedMsgServer
	AccountDepositResponses []escrowv1.MsgAccountDepositResponse
}

// EscrowTxData holds mock data for escrow transactions
type EscrowTxData struct {
	AccountDepositResponses []escrowv1.MsgAccountDepositResponse `json:"account_deposit_responses"`
}

// NewMockEscrowMsgServer creates a new mock escrow message server
func NewMockEscrowMsgServer(data EscrowTxData) *MockEscrowMsgServer {
	return &MockEscrowMsgServer{
		AccountDepositResponses: data.AccountDepositResponses,
	}
}

// AccountDeposit implements the AccountDeposit transaction
func (m *MockEscrowMsgServer) AccountDeposit(ctx context.Context, req *escrowv1.MsgAccountDeposit) (*escrowv1.MsgAccountDepositResponse, error) {
	if len(m.AccountDepositResponses) == 0 {
		return &escrowv1.MsgAccountDepositResponse{}, nil
	}
	return &m.AccountDepositResponses[0], nil
}
