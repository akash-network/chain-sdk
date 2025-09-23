package mocks

import (
	"context"

	providerv1beta4 "pkg.akt.dev/go/node/provider/v1beta4"
)

// MockProviderMsgServer implements the provider message server for transactions
type MockProviderMsgServer struct {
	providerv1beta4.UnimplementedMsgServer
	CreateProviderResponses []providerv1beta4.MsgCreateProviderResponse
	UpdateProviderResponses []providerv1beta4.MsgUpdateProviderResponse
	DeleteProviderResponses []providerv1beta4.MsgDeleteProviderResponse
}

// ProviderTxData holds mock data for provider transactions
type ProviderTxData struct {
	CreateProviderResponses []providerv1beta4.MsgCreateProviderResponse `json:"create_provider_responses"`
	UpdateProviderResponses []providerv1beta4.MsgUpdateProviderResponse `json:"update_provider_responses"`
	DeleteProviderResponses []providerv1beta4.MsgDeleteProviderResponse `json:"delete_provider_responses"`
}

// NewMockProviderMsgServer creates a new mock provider message server
func NewMockProviderMsgServer(data ProviderTxData) *MockProviderMsgServer {
	return &MockProviderMsgServer{
		CreateProviderResponses: data.CreateProviderResponses,
		UpdateProviderResponses: data.UpdateProviderResponses,
		DeleteProviderResponses: data.DeleteProviderResponses,
	}
}

// CreateProvider implements the CreateProvider transaction
func (m *MockProviderMsgServer) CreateProvider(ctx context.Context, req *providerv1beta4.MsgCreateProvider) (*providerv1beta4.MsgCreateProviderResponse, error) {
	if len(m.CreateProviderResponses) == 0 {
		return &providerv1beta4.MsgCreateProviderResponse{}, nil
	}
	return &m.CreateProviderResponses[0], nil
}

// UpdateProvider implements the UpdateProvider transaction
func (m *MockProviderMsgServer) UpdateProvider(ctx context.Context, req *providerv1beta4.MsgUpdateProvider) (*providerv1beta4.MsgUpdateProviderResponse, error) {
	if len(m.UpdateProviderResponses) == 0 {
		return &providerv1beta4.MsgUpdateProviderResponse{}, nil
	}
	return &m.UpdateProviderResponses[0], nil
}

// DeleteProvider implements the DeleteProvider transaction
func (m *MockProviderMsgServer) DeleteProvider(ctx context.Context, req *providerv1beta4.MsgDeleteProvider) (*providerv1beta4.MsgDeleteProviderResponse, error) {
	if len(m.DeleteProviderResponses) == 0 {
		return &providerv1beta4.MsgDeleteProviderResponse{}, nil
	}
	return &m.DeleteProviderResponses[0], nil
}
