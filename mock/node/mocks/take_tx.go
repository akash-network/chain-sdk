package mocks

import (
	"context"

	takev1 "pkg.akt.dev/go/node/take/v1"
)

// MockTakeMsgServer implements the take message server for transactions
type MockTakeMsgServer struct {
	takev1.UnimplementedMsgServer
	UpdateParamsResponses []takev1.MsgUpdateParamsResponse
}

// TakeTxData holds mock data for take transactions
type TakeTxData struct {
	UpdateParamsResponses []takev1.MsgUpdateParamsResponse `json:"update_params_responses"`
}

// NewMockTakeMsgServer creates a new mock take message server
func NewMockTakeMsgServer(data TakeTxData) *MockTakeMsgServer {
	return &MockTakeMsgServer{
		UpdateParamsResponses: data.UpdateParamsResponses,
	}
}

// UpdateParams implements the UpdateParams transaction
func (m *MockTakeMsgServer) UpdateParams(ctx context.Context, req *takev1.MsgUpdateParams) (*takev1.MsgUpdateParamsResponse, error) {
	if len(m.UpdateParamsResponses) == 0 {
		return &takev1.MsgUpdateParamsResponse{}, nil
	}
	return &m.UpdateParamsResponses[0], nil
}
