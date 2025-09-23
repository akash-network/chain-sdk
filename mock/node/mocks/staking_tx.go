package mocks

import (
	"context"

	stakingv1beta3 "pkg.akt.dev/go/node/staking/v1beta3"
)

// MockStakingMsgServer implements the staking message server for transactions
type MockStakingMsgServer struct {
	stakingv1beta3.UnimplementedMsgServer
	UpdateParamsResponses []stakingv1beta3.MsgUpdateParamsResponse
}

// StakingTxData holds mock data for staking transactions
type StakingTxData struct {
	UpdateParamsResponses []stakingv1beta3.MsgUpdateParamsResponse `json:"update_params_responses"`
}

// NewMockStakingMsgServer creates a new mock staking message server
func NewMockStakingMsgServer(data StakingTxData) *MockStakingMsgServer {
	return &MockStakingMsgServer{
		UpdateParamsResponses: data.UpdateParamsResponses,
	}
}

// UpdateParams implements the UpdateParams transaction
func (m *MockStakingMsgServer) UpdateParams(ctx context.Context, req *stakingv1beta3.MsgUpdateParams) (*stakingv1beta3.MsgUpdateParamsResponse, error) {
	if len(m.UpdateParamsResponses) == 0 {
		return &stakingv1beta3.MsgUpdateParamsResponse{}, nil
	}
	return &m.UpdateParamsResponses[0], nil
}
