package tx

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	abciv1beta1 "cosmossdk.io/api/cosmos/base/abci/v1beta1"
	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
)

type Service struct {
	txv1beta1.UnimplementedServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Simulate(ctx context.Context, req *txv1beta1.SimulateRequest) (*txv1beta1.SimulateResponse, error) {
	// Return a mock gas estimate
	// Typical gas for a deployment transaction is around 200000-500000
	return &txv1beta1.SimulateResponse{
		GasInfo: &abciv1beta1.GasInfo{
			GasWanted: 300000,
			GasUsed:  250000,
		},
		Result: &abciv1beta1.Result{
			Data: []byte{},
			Log:  "mock simulation successful",
		},
	}, nil
}

func (s *Service) BroadcastTx(ctx context.Context, req *txv1beta1.BroadcastTxRequest) (*txv1beta1.BroadcastTxResponse, error) {
	// Generate a fake tx hash from the tx bytes
	hash := sha256.Sum256(req.TxBytes)
	txHash := hex.EncodeToString(hash[:])

	// Return success response
	return &txv1beta1.BroadcastTxResponse{
		TxResponse: &abciv1beta1.TxResponse{
			Height:    1000,
			Txhash:    txHash,
			Code:      0,
			Data:      "",
			RawLog:    `[{"msg_index":0,"events":[{"type":"message","attributes":[{"key":"action","value":"create_deployment"}]}]}]`,
			Logs:      []*abciv1beta1.ABCIMessageLog{},
			Info:      "",
			GasWanted: 300000,
			GasUsed:   250000,
			Tx:        nil,
			Timestamp: "",
		},
	}, nil
}

func (s *Service) GetTx(ctx context.Context, req *txv1beta1.GetTxRequest) (*txv1beta1.GetTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTx not implemented")
}

func (s *Service) GetTxsEvent(ctx context.Context, req *txv1beta1.GetTxsEventRequest) (*txv1beta1.GetTxsEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxsEvent not implemented")
}

func (s *Service) GetBlockWithTxs(ctx context.Context, req *txv1beta1.GetBlockWithTxsRequest) (*txv1beta1.GetBlockWithTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockWithTxs not implemented")
}

func (s *Service) TxDecode(ctx context.Context, req *txv1beta1.TxDecodeRequest) (*txv1beta1.TxDecodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxDecode not implemented")
}

func (s *Service) TxEncode(ctx context.Context, req *txv1beta1.TxEncodeRequest) (*txv1beta1.TxEncodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxEncode not implemented")
}

func (s *Service) TxEncodeAmino(ctx context.Context, req *txv1beta1.TxEncodeAminoRequest) (*txv1beta1.TxEncodeAminoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxEncodeAmino not implemented")
}

func (s *Service) TxDecodeAmino(ctx context.Context, req *txv1beta1.TxDecodeAminoRequest) (*txv1beta1.TxDecodeAminoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxDecodeAmino not implemented")
}

