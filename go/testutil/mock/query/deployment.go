package query

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dv1beta4 "pkg.akt.dev/go/node/deployment/v1beta4"
)

type DeploymentQuery struct {
	codec codec.Codec
}

func NewDeploymentQuery(codec codec.Codec) *DeploymentQuery {
	return &DeploymentQuery{
		codec: codec,
	}
}

func (q *DeploymentQuery) Deployments(ctx context.Context, req *dv1beta4.QueryDeploymentsRequest) (*dv1beta4.QueryDeploymentsResponse, error) {
	return &dv1beta4.QueryDeploymentsResponse{
		Deployments: dv1beta4.DeploymentResponses{},
		Pagination:  &sdkquery.PageResponse{Total: 0},
	}, nil
}

func (q *DeploymentQuery) Deployment(ctx context.Context, req *dv1beta4.QueryDeploymentRequest) (*dv1beta4.QueryDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deployment not implemented")
}

func (q *DeploymentQuery) Group(ctx context.Context, req *dv1beta4.QueryGroupRequest) (*dv1beta4.QueryGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Group not implemented")
}

func (q *DeploymentQuery) Params(ctx context.Context, req *dv1beta4.QueryParamsRequest) (*dv1beta4.QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
