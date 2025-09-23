package mocks

import (
	"context"

	certv1 "pkg.akt.dev/go/node/cert/v1"
)

// MockCertMsgServer implements the certificate message server for transactions
type MockCertMsgServer struct {
	certv1.UnimplementedMsgServer
	CreateCertificateResponses []certv1.MsgCreateCertificateResponse
	RevokeCertificateResponses []certv1.MsgRevokeCertificateResponse
}

// CertTxData holds mock data for certificate transactions
type CertTxData struct {
	CreateCertificateResponses []certv1.MsgCreateCertificateResponse `json:"create_certificate_responses"`
	RevokeCertificateResponses []certv1.MsgRevokeCertificateResponse `json:"revoke_certificate_responses"`
}

// NewMockCertMsgServer creates a new mock certificate message server
func NewMockCertMsgServer(data CertTxData) *MockCertMsgServer {
	return &MockCertMsgServer{
		CreateCertificateResponses: data.CreateCertificateResponses,
		RevokeCertificateResponses: data.RevokeCertificateResponses,
	}
}

// CreateCertificate implements the CreateCertificate transaction
func (m *MockCertMsgServer) CreateCertificate(ctx context.Context, req *certv1.MsgCreateCertificate) (*certv1.MsgCreateCertificateResponse, error) {
	if len(m.CreateCertificateResponses) == 0 {
		return &certv1.MsgCreateCertificateResponse{}, nil
	}
	return &m.CreateCertificateResponses[0], nil
}

// RevokeCertificate implements the RevokeCertificate transaction
func (m *MockCertMsgServer) RevokeCertificate(ctx context.Context, req *certv1.MsgRevokeCertificate) (*certv1.MsgRevokeCertificateResponse, error) {
	if len(m.RevokeCertificateResponses) == 0 {
		return &certv1.MsgRevokeCertificateResponse{}, nil
	}
	return &m.RevokeCertificateResponses[0], nil
}
