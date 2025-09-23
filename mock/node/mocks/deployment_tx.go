package mocks

import (
	"context"
	"fmt"

	deploymentv1beta4 "pkg.akt.dev/go/node/deployment/v1beta4"
)

// MockDeploymentMsgServer implements the deployment message server for transactions
type MockDeploymentMsgServer struct {
	deploymentv1beta4.UnimplementedMsgServer
	CreateDeploymentResponses []deploymentv1beta4.MsgCreateDeploymentResponse
	UpdateDeploymentResponses []deploymentv1beta4.MsgUpdateDeploymentResponse
	CloseDeploymentResponses  []deploymentv1beta4.MsgCloseDeploymentResponse
}

// DeploymentTxData holds mock data for deployment transactions
type DeploymentTxData struct {
	CreateDeploymentResponses []deploymentv1beta4.MsgCreateDeploymentResponse `json:"create_deployment_responses"`
	UpdateDeploymentResponses []deploymentv1beta4.MsgUpdateDeploymentResponse `json:"update_deployment_responses"`
	CloseDeploymentResponses  []deploymentv1beta4.MsgCloseDeploymentResponse  `json:"close_deployment_responses"`
}

// NewMockDeploymentMsgServer creates a new mock deployment message server
func NewMockDeploymentMsgServer(data DeploymentTxData) *MockDeploymentMsgServer {
	return &MockDeploymentMsgServer{
		CreateDeploymentResponses: data.CreateDeploymentResponses,
		UpdateDeploymentResponses: data.UpdateDeploymentResponses,
		CloseDeploymentResponses:  data.CloseDeploymentResponses,
	}
}

// CreateDeployment implements the CreateDeployment transaction
func (m *MockDeploymentMsgServer) CreateDeployment(ctx context.Context, req *deploymentv1beta4.MsgCreateDeployment) (*deploymentv1beta4.MsgCreateDeploymentResponse, error) {
	if len(m.CreateDeploymentResponses) == 0 {
		return &deploymentv1beta4.MsgCreateDeploymentResponse{}, nil
	}

	// Return the first available response for simplicity
	// In a more sophisticated mock, you might match based on request parameters
	return &m.CreateDeploymentResponses[0], nil
}

// UpdateDeployment implements the UpdateDeployment transaction
func (m *MockDeploymentMsgServer) UpdateDeployment(ctx context.Context, req *deploymentv1beta4.MsgUpdateDeployment) (*deploymentv1beta4.MsgUpdateDeploymentResponse, error) {
	if len(m.UpdateDeploymentResponses) == 0 {
		return &deploymentv1beta4.MsgUpdateDeploymentResponse{}, nil
	}

	return &m.UpdateDeploymentResponses[0], nil
}

// CloseDeployment implements the CloseDeployment transaction
func (m *MockDeploymentMsgServer) CloseDeployment(ctx context.Context, req *deploymentv1beta4.MsgCloseDeployment) (*deploymentv1beta4.MsgCloseDeploymentResponse, error) {
	if len(m.CloseDeploymentResponses) == 0 {
		return &deploymentv1beta4.MsgCloseDeploymentResponse{}, nil
	}

	return &m.CloseDeploymentResponses[0], nil
}

// CloseGroup implements the CloseGroup transaction (not implemented for this example)
func (m *MockDeploymentMsgServer) CloseGroup(ctx context.Context, req *deploymentv1beta4.MsgCloseGroup) (*deploymentv1beta4.MsgCloseGroupResponse, error) {
	return nil, fmt.Errorf("CloseGroup not implemented in this mock")
}

// PauseGroup implements the PauseGroup transaction (not implemented for this example)
func (m *MockDeploymentMsgServer) PauseGroup(ctx context.Context, req *deploymentv1beta4.MsgPauseGroup) (*deploymentv1beta4.MsgPauseGroupResponse, error) {
	return nil, fmt.Errorf("PauseGroup not implemented in this mock")
}

// StartGroup implements the StartGroup transaction (not implemented for this example)
func (m *MockDeploymentMsgServer) StartGroup(ctx context.Context, req *deploymentv1beta4.MsgStartGroup) (*deploymentv1beta4.MsgStartGroupResponse, error) {
	return nil, fmt.Errorf("StartGroup not implemented in this mock")
}

// UpdateParams implements the UpdateParams transaction (not implemented for this example)
func (m *MockDeploymentMsgServer) UpdateParams(ctx context.Context, req *deploymentv1beta4.MsgUpdateParams) (*deploymentv1beta4.MsgUpdateParamsResponse, error) {
	return nil, fmt.Errorf("UpdateParams not implemented in this mock")
}
