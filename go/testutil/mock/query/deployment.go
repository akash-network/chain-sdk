package query

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/codec"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dv1beta4 "pkg.akt.dev/go/node/deployment/v1beta4"
)

type DeploymentQuery struct {
	dataDir string
	codec   codec.Codec
	data    *deploymentData
}

type deploymentData struct {
	Deployments *dv1beta4.QueryDeploymentsResponse `json:"deployments,omitempty"`
}

func NewDeploymentQuery(dataDir string, codec codec.Codec) *DeploymentQuery {
	return &DeploymentQuery{
		dataDir: dataDir,
		codec:   codec,
	}
}

func (q *DeploymentQuery) loadData() error {
	if q.data != nil {
		return nil
	}

	dataPath := filepath.Join(q.dataDir, "deployments.json")
	dataBytes, err := os.ReadFile(dataPath)
	if err != nil {
		return fmt.Errorf("failed to read deployments.json: %w", err)
	}

	var data deploymentData
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return fmt.Errorf("failed to unmarshal deployments.json: %w", err)
	}

	if data.Deployments != nil && data.Deployments.Deployments == nil {
		data.Deployments.Deployments = []dv1beta4.QueryDeploymentResponse{}
	}

	q.data = &data
	return nil
}

func (q *DeploymentQuery) Deployments(ctx context.Context, req *dv1beta4.QueryDeploymentsRequest) (*dv1beta4.QueryDeploymentsResponse, error) {
	if err := q.loadData(); err != nil {
		return nil, fmt.Errorf("failed to load deployment fixtures: %w", err)
	}

	resp := &dv1beta4.QueryDeploymentsResponse{
		Deployments: dv1beta4.DeploymentResponses{},
	}

	if q.data.Deployments != nil {
		resp = q.data.Deployments
		if resp.Deployments == nil {
			resp.Deployments = dv1beta4.DeploymentResponses{}
		}
	}

	for _, depResp := range resp.Deployments {
		for _, group := range depResp.Groups {
			for _, resource := range group.GroupSpec.Resources {
				if err := resource.Price.Validate(); err != nil {
					return nil, fmt.Errorf("invalid deployment resource price: %w", err)
				}
			}
		}
	}

	return resp, nil
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
