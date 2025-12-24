package mock

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	dv1beta4 "pkg.akt.dev/go/node/deployment/v1beta4"
	mv1beta5 "pkg.akt.dev/go/node/market/v1beta5"
	"pkg.akt.dev/go/sdkutil"
	"pkg.akt.dev/go/testutil/mock/query"
	"pkg.akt.dev/go/testutil/mock/tx"
)

type Server struct {
	grpcAddr         string
	gatewayAddr      string
	grpcSrv          *grpc.Server
	gatewaySrv       *http.Server
	gatewayMux       *runtime.ServeMux
	grpcConn         *grpc.ClientConn
	txConfig         client.TxConfig
	group            *errgroup.Group
	ctx              context.Context
	cancel           context.CancelFunc
	lastDeploymentMu sync.Mutex
	lastDeployment   *dv1beta4.MsgCreateDeployment
	lastBidMu        sync.Mutex
	lastBid          *mv1beta5.MsgCreateBid
}

type Config struct {
	GRPCAddr    string
	GatewayAddr string
}

func NewServer(cfg Config) (*Server, error) {
	if cfg.GRPCAddr == "" {
		cfg.GRPCAddr = "127.0.0.1:0"
	}
	if cfg.GatewayAddr == "" {
		cfg.GatewayAddr = "127.0.0.1:0"
	}

	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)

	encCfg := setupCodec()
	grpcSrv := grpc.NewServer()
	deploymentQuery, marketQuery := registerGRPCServices(grpcSrv, encCfg.Codec)

	mux, err := registerGatewayHandlers(ctx, deploymentQuery, marketQuery)
	if err != nil {
		cancel()
		return nil, err
	}

	gatewaySrv := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	server := &Server{
		grpcAddr:    cfg.GRPCAddr,
		gatewayAddr: cfg.GatewayAddr,
		grpcSrv:     grpcSrv,
		gatewaySrv:  gatewaySrv,
		gatewayMux:  mux,
		txConfig:    encCfg.TxConfig,
		group:       group,
		ctx:         ctx,
		cancel:      cancel,
	}

	server.registerDebugHandlers()

	return server, nil
}

func setupCodec() sdkutil.EncodingConfig {
	encCfg := sdkutil.MakeEncodingConfig()
	dv1beta4.RegisterInterfaces(encCfg.InterfaceRegistry)
	mv1beta5.RegisterInterfaces(encCfg.InterfaceRegistry)
	dv1beta4.RegisterLegacyAminoCodec(encCfg.Amino)
	mv1beta5.RegisterLegacyAminoCodec(encCfg.Amino)
	return encCfg
}

func registerGRPCServices(grpcSrv *grpc.Server, codec codec.Codec) (*query.DeploymentQuery, *query.MarketQuery) {
	deploymentQuery := query.NewDeploymentQuery(codec)
	dv1beta4.RegisterQueryServer(grpcSrv, deploymentQuery)

	marketQuery := query.NewMarketQuery(codec)
	mv1beta5.RegisterQueryServer(grpcSrv, marketQuery)

	txService := tx.NewService()
	txv1beta1.RegisterServiceServer(grpcSrv, txService)

	return deploymentQuery, marketQuery
}

func registerGatewayHandlers(ctx context.Context, deploymentQuery *query.DeploymentQuery, marketQuery *query.MarketQuery) (*runtime.ServeMux, error) {
	jsonpbMarshaler := &runtime.JSONPb{
		OrigName:     true,
		EmitDefaults: true,
	}
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpbMarshaler),
	)

	if err := dv1beta4.RegisterQueryHandlerServer(ctx, mux, deploymentQuery); err != nil {
		return nil, fmt.Errorf("failed to register deployment query handler: %w", err)
	}

	if err := mv1beta5.RegisterQueryHandlerServer(ctx, mux, marketQuery); err != nil {
		return nil, fmt.Errorf("failed to register market query handler: %w", err)
	}

	return mux, nil
}

func (s *Server) Start() (err error) {
	grpcLis, gatewayLis, err := s.createListeners()
	if err != nil {
		return err
	}

	if err := s.startGRPCServer(grpcLis); err != nil {
		_ = gatewayLis.Close()
		return err
	}

	defer func() {
		if err != nil {
			_ = s.Stop()
			_ = gatewayLis.Close()
		}
	}()

	if err = s.waitForGRPCReady(); err != nil {
		return err
	}

	if err = s.setupTxHandlers(); err != nil {
		return err
	}

	if err = s.startGatewayServer(gatewayLis); err != nil {
		return err
	}

	return nil
}

func (s *Server) createListeners() (grpcLis, gatewayLis net.Listener, err error) {
	grpcLis, err = net.Listen("tcp", s.grpcAddr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen on grpc addr: %w", err)
	}
	s.grpcAddr = grpcLis.Addr().String()

	gatewayLis, err = net.Listen("tcp", s.gatewayAddr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen on gateway addr: %w", err)
	}
	s.gatewayAddr = gatewayLis.Addr().String()

	return grpcLis, gatewayLis, nil
}

func (s *Server) startGRPCServer(lis net.Listener) error {
	s.group.Go(func() error {
		err := s.grpcSrv.Serve(lis)
		if errors.Is(err, grpc.ErrServerStopped) {
			return nil
		}
		return err
	})
	return nil
}

func (s *Server) waitForGRPCReady() error {
	readyCtx, readyCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer readyCancel()

	backoff := 10 * time.Millisecond
	maxBackoff := 500 * time.Millisecond

	for {
		dialCtx, dialCancel := context.WithTimeout(readyCtx, 500*time.Millisecond)
		conn, err := grpc.DialContext(dialCtx, s.grpcAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock())
		dialCancel()

		if err == nil {
			s.grpcConn = conn
			return nil
		}

		if readyCtx.Err() != nil {
			return fmt.Errorf("grpc server readiness check timed out: %w", readyCtx.Err())
		}

		time.Sleep(backoff)
		backoff *= 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

func (s *Server) setupTxHandlers() error {
	txClient := txv1beta1.NewServiceClient(s.grpcConn)
	s.registerSimulateHandler(txClient)
	s.registerBroadcastHandler(txClient)
	return nil
}

func (s *Server) registerSimulateHandler(txClient txv1beta1.ServiceClient) {
	simulatePattern := runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"cosmos", "tx", "v1beta1", "simulate"}, "", runtime.AssumeColonVerbOpt(false)))
	s.gatewayMux.Handle("POST", simulatePattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		var jsonReq map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&jsonReq); err != nil {
			_, outboundMarshaler := runtime.MarshalerForRequest(s.gatewayMux, r)
			runtime.HTTPError(r.Context(), s.gatewayMux, outboundMarshaler, w, r, err)
			return
		}

		var req txv1beta1.SimulateRequest
		if txBytesStr, ok := jsonReq["tx_bytes"].(string); ok && txBytesStr != "" {
			txBytes, err := base64.StdEncoding.DecodeString(txBytesStr)
			if err != nil {
				_, outboundMarshaler := runtime.MarshalerForRequest(s.gatewayMux, r)
				runtime.HTTPError(r.Context(), s.gatewayMux, outboundMarshaler, w, r, fmt.Errorf("invalid tx_bytes: %w", err))
				return
			}
			req.TxBytes = txBytes

			if err := s.validateTxBytes(txBytes); err != nil {
				_, outboundMarshaler := runtime.MarshalerForRequest(s.gatewayMux, r)
				runtime.HTTPError(r.Context(), s.gatewayMux, outboundMarshaler, w, r, fmt.Errorf("transaction validation failed: %w", err))
				return
			}
		}

		resp, err := txClient.Simulate(r.Context(), &req)
		if err != nil {
			_, outboundMarshaler := runtime.MarshalerForRequest(s.gatewayMux, r)
			runtime.HTTPError(r.Context(), s.gatewayMux, outboundMarshaler, w, r, err)
			return
		}

		_, outboundMarshaler := runtime.MarshalerForRequest(s.gatewayMux, r)
		rctx, err := runtime.AnnotateIncomingContext(r.Context(), s.gatewayMux, r)
		if err != nil {
			runtime.HTTPError(r.Context(), s.gatewayMux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(rctx, s.gatewayMux, outboundMarshaler, w, r, resp, s.gatewayMux.GetForwardResponseOptions()...)
	})
}

func (s *Server) registerBroadcastHandler(txClient txv1beta1.ServiceClient) {
	broadcastPattern := runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"cosmos", "tx", "v1beta1", "txs"}, "", runtime.AssumeColonVerbOpt(false)))
	s.gatewayMux.Handle("POST", broadcastPattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		var jsonReq map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&jsonReq); err != nil {
			_, outboundMarshaler := runtime.MarshalerForRequest(s.gatewayMux, r)
			runtime.HTTPError(ctx, s.gatewayMux, outboundMarshaler, w, r, err)
			return
		}

		var req txv1beta1.BroadcastTxRequest
		if txBytesStr, ok := jsonReq["tx_bytes"].(string); ok {
			txBytes, err := base64.StdEncoding.DecodeString(txBytesStr)
			if err != nil {
				_, outboundMarshaler := runtime.MarshalerForRequest(s.gatewayMux, r)
				runtime.HTTPError(ctx, s.gatewayMux, outboundMarshaler, w, r, fmt.Errorf("invalid tx_bytes: %w", err))
				return
			}
			req.TxBytes = txBytes

			if err := s.validateTxBytes(txBytes); err != nil {
				_, outboundMarshaler := runtime.MarshalerForRequest(s.gatewayMux, r)
				runtime.HTTPError(ctx, s.gatewayMux, outboundMarshaler, w, r, fmt.Errorf("transaction validation failed: %w", err))
				return
			}
		}

		req.Mode = parseBroadcastMode(jsonReq)

		resp, err := txClient.BroadcastTx(ctx, &req)
		if err != nil {
			_, outboundMarshaler := runtime.MarshalerForRequest(s.gatewayMux, r)
			runtime.HTTPError(ctx, s.gatewayMux, outboundMarshaler, w, r, err)
			return
		}

		_, outboundMarshaler := runtime.MarshalerForRequest(s.gatewayMux, r)
		rctx, err := runtime.AnnotateIncomingContext(ctx, s.gatewayMux, r)
		if err != nil {
			runtime.HTTPError(ctx, s.gatewayMux, outboundMarshaler, w, r, err)
			return
		}
		runtime.ForwardResponseMessage(rctx, s.gatewayMux, outboundMarshaler, w, r, resp, s.gatewayMux.GetForwardResponseOptions()...)
	})
}

func (s *Server) startGatewayServer(lis net.Listener) error {
	s.group.Go(func() error {
		err := s.gatewaySrv.Serve(lis)
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	})
	return nil
}

func (s *Server) registerDebugHandlers() {
	pattern := runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"mock", "last-deployment"}, "", runtime.AssumeColonVerbOpt(false)))
	s.gatewayMux.Handle("GET", pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		deployment := s.getLastDeployment()
		if deployment == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		marshaler := &jsonpb.Marshaler{OrigName: true, EmitDefaults: true}
		if err := marshaler.Marshal(w, deployment); err != nil {
			http.Error(w, fmt.Sprintf("failed to marshal deployment: %v", err), http.StatusInternalServerError)
		}
	})

	bidPattern := runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"mock", "last-bid"}, "", runtime.AssumeColonVerbOpt(false)))
	s.gatewayMux.Handle("GET", bidPattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		bid := s.getLastBid()
		if bid == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		marshaler := &jsonpb.Marshaler{OrigName: true, EmitDefaults: true}
		if err := marshaler.Marshal(w, bid); err != nil {
			http.Error(w, fmt.Sprintf("failed to marshal bid: %v", err), http.StatusInternalServerError)
		}
	})
}

func (s *Server) setLastDeployment(msg *dv1beta4.MsgCreateDeployment) {
	s.lastDeploymentMu.Lock()
	defer s.lastDeploymentMu.Unlock()
	s.lastDeployment = msg
}

func (s *Server) getLastDeployment() *dv1beta4.MsgCreateDeployment {
	s.lastDeploymentMu.Lock()
	defer s.lastDeploymentMu.Unlock()
	if s.lastDeployment == nil {
		return nil
	}
	copy := *s.lastDeployment
	return &copy
}

func (s *Server) setLastBid(msg *mv1beta5.MsgCreateBid) {
	s.lastBidMu.Lock()
	defer s.lastBidMu.Unlock()
	s.lastBid = msg
}

func (s *Server) getLastBid() *mv1beta5.MsgCreateBid {
	s.lastBidMu.Lock()
	defer s.lastBidMu.Unlock()
	if s.lastBid == nil {
		return nil
	}
	copy := *s.lastBid
	return &copy
}

func parseBroadcastMode(jsonReq map[string]interface{}) txv1beta1.BroadcastMode {
	modeStr, ok := jsonReq["mode"].(string)
	if !ok {
		return txv1beta1.BroadcastMode_BROADCAST_MODE_SYNC
	}

	switch strings.ToUpper(modeStr) {
	case "BROADCAST_MODE_UNSPECIFIED", "BROADCAST_MODE_UNSPECIFIED_VALUE":
		return txv1beta1.BroadcastMode_BROADCAST_MODE_UNSPECIFIED
	case "BROADCAST_MODE_BLOCK":
		return txv1beta1.BroadcastMode_BROADCAST_MODE_BLOCK
	case "BROADCAST_MODE_SYNC":
		return txv1beta1.BroadcastMode_BROADCAST_MODE_SYNC
	case "BROADCAST_MODE_ASYNC":
		return txv1beta1.BroadcastMode_BROADCAST_MODE_ASYNC
	default:
		return txv1beta1.BroadcastMode_BROADCAST_MODE_SYNC
	}
}

func (s *Server) validateTxBytes(txBytes []byte) error {
	if len(txBytes) == 0 {
		return nil
	}

	txDecoder := s.txConfig.TxDecoder()
	decodedTx, err := txDecoder(txBytes)
	if err != nil {
		return fmt.Errorf("failed to decode transaction: %w", err)
	}

	msgs := decodedTx.GetMsgs()
	for i, msg := range msgs {
		if validator, ok := msg.(sdk.HasValidateBasic); ok {
			if err := validator.ValidateBasic(); err != nil {
				return fmt.Errorf("message %d validation failed: %w", i, err)
			}
		}

		if deploymentMsg, ok := msg.(*dv1beta4.MsgCreateDeployment); ok {
			s.setLastDeployment(deploymentMsg)
		}

		if bidMsg, ok := msg.(*mv1beta5.MsgCreateBid); ok {
			s.setLastBid(bidMsg)
		}
	}

	return nil
}

func (s *Server) GatewayURL() string {
	return fmt.Sprintf("http://%s", s.gatewayAddr)
}

func (s *Server) GRPCAddr() string {
	return s.grpcAddr
}

func (s *Server) Stop() error {
	s.cancel()

	if s.grpcSrv != nil {
		s.grpcSrv.Stop()
	}

	if s.gatewaySrv != nil {
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		_ = s.gatewaySrv.Shutdown(shutdownCtx)
	}

	if s.grpcConn != nil {
		_ = s.grpcConn.Close()
	}

	return s.group.Wait()
}
