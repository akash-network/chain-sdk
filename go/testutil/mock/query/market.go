package query

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mv1beta5 "pkg.akt.dev/go/node/market/v1beta5"
)

type MarketQuery struct {
	codec codec.Codec
}

func NewMarketQuery(codec codec.Codec) *MarketQuery {
	return &MarketQuery{
		codec: codec,
	}
}

func (q *MarketQuery) Leases(ctx context.Context, req *mv1beta5.QueryLeasesRequest) (*mv1beta5.QueryLeasesResponse, error) {
	return &mv1beta5.QueryLeasesResponse{
		Leases:     []mv1beta5.QueryLeaseResponse{},
		Pagination: &query.PageResponse{Total: 0},
	}, nil
}

func (q *MarketQuery) Lease(ctx context.Context, req *mv1beta5.QueryLeaseRequest) (*mv1beta5.QueryLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lease not implemented")
}

func (q *MarketQuery) Bids(ctx context.Context, req *mv1beta5.QueryBidsRequest) (*mv1beta5.QueryBidsResponse, error) {
	return &mv1beta5.QueryBidsResponse{
		Bids:       []mv1beta5.QueryBidResponse{},
		Pagination: &query.PageResponse{Total: 0},
	}, nil
}

func (q *MarketQuery) Bid(ctx context.Context, req *mv1beta5.QueryBidRequest) (*mv1beta5.QueryBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}

func (q *MarketQuery) Orders(ctx context.Context, req *mv1beta5.QueryOrdersRequest) (*mv1beta5.QueryOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Orders not implemented")
}

func (q *MarketQuery) Order(ctx context.Context, req *mv1beta5.QueryOrderRequest) (*mv1beta5.QueryOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Order not implemented")
}

func (q *MarketQuery) Params(ctx context.Context, req *mv1beta5.QueryParamsRequest) (*mv1beta5.QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
