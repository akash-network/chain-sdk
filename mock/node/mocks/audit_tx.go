package mocks

import (
	"context"

	auditv1 "pkg.akt.dev/go/node/audit/v1"
)

// MockAuditMsgServer implements the audit message server for transactions
type MockAuditMsgServer struct {
	auditv1.UnimplementedMsgServer
	SignProviderAttributesResponses   []auditv1.MsgSignProviderAttributesResponse
	DeleteProviderAttributesResponses []auditv1.MsgDeleteProviderAttributesResponse
}

// AuditTxData holds mock data for audit transactions
type AuditTxData struct {
	SignProviderAttributesResponses   []auditv1.MsgSignProviderAttributesResponse   `json:"sign_provider_attributes_responses"`
	DeleteProviderAttributesResponses []auditv1.MsgDeleteProviderAttributesResponse `json:"delete_provider_attributes_responses"`
}

// NewMockAuditMsgServer creates a new mock audit message server
func NewMockAuditMsgServer(data AuditTxData) *MockAuditMsgServer {
	return &MockAuditMsgServer{
		SignProviderAttributesResponses:   data.SignProviderAttributesResponses,
		DeleteProviderAttributesResponses: data.DeleteProviderAttributesResponses,
	}
}

// SignProviderAttributes implements the SignProviderAttributes transaction
func (m *MockAuditMsgServer) SignProviderAttributes(ctx context.Context, req *auditv1.MsgSignProviderAttributes) (*auditv1.MsgSignProviderAttributesResponse, error) {
	if len(m.SignProviderAttributesResponses) == 0 {
		return &auditv1.MsgSignProviderAttributesResponse{}, nil
	}
	return &m.SignProviderAttributesResponses[0], nil
}

// DeleteProviderAttributes implements the DeleteProviderAttributes transaction
func (m *MockAuditMsgServer) DeleteProviderAttributes(ctx context.Context, req *auditv1.MsgDeleteProviderAttributes) (*auditv1.MsgDeleteProviderAttributesResponse, error) {
	if len(m.DeleteProviderAttributesResponses) == 0 {
		return &auditv1.MsgDeleteProviderAttributesResponse{}, nil
	}
	return &m.DeleteProviderAttributesResponses[0], nil
}
