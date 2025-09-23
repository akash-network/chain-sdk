package mocks

import (
	"context"

	marketv1beta5 "pkg.akt.dev/go/node/market/v1beta5"
)

// MockMarketMsgServer implements the market message server for transactions
type MockMarketMsgServer struct {
	marketv1beta5.UnimplementedMsgServer
	CreateBidResponses     []marketv1beta5.MsgCreateBidResponse
	CloseBidResponses      []marketv1beta5.MsgCloseBidResponse
	WithdrawLeaseResponses []marketv1beta5.MsgWithdrawLeaseResponse
	CreateLeaseResponses   []marketv1beta5.MsgCreateLeaseResponse
	CloseLeaseResponses    []marketv1beta5.MsgCloseLeaseResponse
}

// MarketTxData holds mock data for market transactions
type MarketTxData struct {
	CreateBidResponses     []marketv1beta5.MsgCreateBidResponse     `json:"create_bid_responses"`
	CloseBidResponses      []marketv1beta5.MsgCloseBidResponse      `json:"close_bid_responses"`
	WithdrawLeaseResponses []marketv1beta5.MsgWithdrawLeaseResponse `json:"withdraw_lease_responses"`
	CreateLeaseResponses   []marketv1beta5.MsgCreateLeaseResponse   `json:"create_lease_responses"`
	CloseLeaseResponses    []marketv1beta5.MsgCloseLeaseResponse    `json:"close_lease_responses"`
}

// NewMockMarketMsgServer creates a new mock market message server
func NewMockMarketMsgServer(data MarketTxData) *MockMarketMsgServer {
	return &MockMarketMsgServer{
		CreateBidResponses:     data.CreateBidResponses,
		CloseBidResponses:      data.CloseBidResponses,
		WithdrawLeaseResponses: data.WithdrawLeaseResponses,
		CreateLeaseResponses:   data.CreateLeaseResponses,
		CloseLeaseResponses:    data.CloseLeaseResponses,
	}
}

// CreateBid implements the CreateBid transaction
func (m *MockMarketMsgServer) CreateBid(ctx context.Context, req *marketv1beta5.MsgCreateBid) (*marketv1beta5.MsgCreateBidResponse, error) {
	if len(m.CreateBidResponses) == 0 {
		return &marketv1beta5.MsgCreateBidResponse{}, nil
	}
	return &m.CreateBidResponses[0], nil
}

// CloseBid implements the CloseBid transaction
func (m *MockMarketMsgServer) CloseBid(ctx context.Context, req *marketv1beta5.MsgCloseBid) (*marketv1beta5.MsgCloseBidResponse, error) {
	if len(m.CloseBidResponses) == 0 {
		return &marketv1beta5.MsgCloseBidResponse{}, nil
	}
	return &m.CloseBidResponses[0], nil
}

// WithdrawLease implements the WithdrawLease transaction
func (m *MockMarketMsgServer) WithdrawLease(ctx context.Context, req *marketv1beta5.MsgWithdrawLease) (*marketv1beta5.MsgWithdrawLeaseResponse, error) {
	if len(m.WithdrawLeaseResponses) == 0 {
		return &marketv1beta5.MsgWithdrawLeaseResponse{}, nil
	}
	return &m.WithdrawLeaseResponses[0], nil
}

// CreateLease implements the CreateLease transaction
func (m *MockMarketMsgServer) CreateLease(ctx context.Context, req *marketv1beta5.MsgCreateLease) (*marketv1beta5.MsgCreateLeaseResponse, error) {
	if len(m.CreateLeaseResponses) == 0 {
		return &marketv1beta5.MsgCreateLeaseResponse{}, nil
	}
	return &m.CreateLeaseResponses[0], nil
}

// CloseLease implements the CloseLease transaction
func (m *MockMarketMsgServer) CloseLease(ctx context.Context, req *marketv1beta5.MsgCloseLease) (*marketv1beta5.MsgCloseLeaseResponse, error) {
	if len(m.CloseLeaseResponses) == 0 {
		return &marketv1beta5.MsgCloseLeaseResponse{}, nil
	}
	return &m.CloseLeaseResponses[0], nil
}
