package query

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mv1beta5 "pkg.akt.dev/go/node/market/v1beta5"
)

type MarketQuery struct {
	dataDir string
	codec   codec.Codec
	data    *marketData
}

type marketData struct {
	Leases *mv1beta5.QueryLeasesResponse `json:"leases,omitempty"`
	Bids   *mv1beta5.QueryBidsResponse   `json:"bids,omitempty"`
}

func NewMarketQuery(dataDir string, codec codec.Codec) *MarketQuery {
	return &MarketQuery{
		dataDir: dataDir,
		codec:   codec,
	}
}

func (q *MarketQuery) loadData() error {
	if q.data != nil {
		return nil
	}

	dataPath := filepath.Join(q.dataDir, "market.json")
	dataBytes, err := os.ReadFile(dataPath)
	if err != nil {
		return fmt.Errorf("failed to read market.json: %w", err)
	}

	var data marketData
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return fmt.Errorf("failed to unmarshal market.json: %w", err)
	}

	if data.Leases != nil && data.Leases.Leases == nil {
		data.Leases.Leases = []mv1beta5.QueryLeaseResponse{}
	}
	if data.Bids != nil && data.Bids.Bids == nil {
		data.Bids.Bids = []mv1beta5.QueryBidResponse{}
	}

	q.data = &data
	return nil
}

func (q *MarketQuery) Leases(ctx context.Context, req *mv1beta5.QueryLeasesRequest) (*mv1beta5.QueryLeasesResponse, error) {
	if err := q.loadData(); err != nil {
		return nil, fmt.Errorf("failed to load market fixtures: %w", err)
	}

	resp := &mv1beta5.QueryLeasesResponse{
		Leases: []mv1beta5.QueryLeaseResponse{},
	}

	if q.data.Leases != nil {
		resp = q.data.Leases
		if resp.Leases == nil {
			resp.Leases = []mv1beta5.QueryLeaseResponse{}
		}
	}

	return resp, nil
}

func (q *MarketQuery) Lease(ctx context.Context, req *mv1beta5.QueryLeaseRequest) (*mv1beta5.QueryLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lease not implemented")
}

func (q *MarketQuery) Bids(ctx context.Context, req *mv1beta5.QueryBidsRequest) (*mv1beta5.QueryBidsResponse, error) {
	if err := q.loadData(); err != nil {
		return nil, fmt.Errorf("failed to load market fixtures: %w", err)
	}

	resp := &mv1beta5.QueryBidsResponse{
		Bids: []mv1beta5.QueryBidResponse{},
	}

	if q.data.Bids != nil {
		resp = q.data.Bids
		if resp.Bids == nil {
			resp.Bids = []mv1beta5.QueryBidResponse{}
		}
	}

	if req != nil && len(resp.Bids) > 0 {
		var filtered []mv1beta5.QueryBidResponse
		for _, bidResp := range resp.Bids {
			if req.Filters.Owner != "" && bidResp.Bid.ID.Owner != req.Filters.Owner {
				continue
			}
			if req.Filters.DSeq != 0 && bidResp.Bid.ID.DSeq != req.Filters.DSeq {
				continue
			}
			if req.Filters.GSeq != 0 && bidResp.Bid.ID.GSeq != req.Filters.GSeq {
				continue
			}
			if req.Filters.OSeq != 0 && bidResp.Bid.ID.OSeq != req.Filters.OSeq {
				continue
			}
			filtered = append(filtered, bidResp)
		}
		resp.Bids = filtered
	}

	if resp.Pagination == nil {
		resp.Pagination = &query.PageResponse{}
	}

	return resp, nil
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
