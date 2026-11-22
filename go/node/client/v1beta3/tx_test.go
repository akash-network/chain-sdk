package v1beta3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	abci "github.com/cometbft/cometbft/abci/types"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	tmtypes "github.com/cometbft/cometbft/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.akt.dev/go/sdkutil"
)

const testBroadcastCommitMethod = "broadcast_tx_commit"

func TestBroadcastTxConfirmsAfterBadGateway(t *testing.T) {
	tests := []struct {
		name              string
		mode              string
		txCode            uint32
		firstQueryFailure bool
		firstQueryMissing bool
		broadcastTimedOut bool
	}{
		{
			name: "block",
			mode: BroadcastBlock,
		},
		{
			name: "sync",
			mode: BroadcastSync,
		},
		{
			name: "async",
			mode: BroadcastAsync,
		},
		{
			name:   "failed transaction",
			mode:   BroadcastBlock,
			txCode: 5,
		},
		{
			name:              "interrupted confirmation",
			mode:              BroadcastBlock,
			firstQueryFailure: true,
		},
		{
			name:              "existing RPC timeout recovery",
			mode:              BroadcastBlock,
			broadcastTimedOut: true,
		},
		{
			name:              "not yet indexed",
			mode:              BroadcastBlock,
			firstQueryMissing: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var broadcasts, queries atomic.Int32
			cctx, tx := newBroadcastTestClient(t, func(w http.ResponseWriter, _ *http.Request, request rpctypes.RPCRequest, encoded []byte) {
				hash := tmtypes.Tx(encoded).Hash()
				var result any
				switch request.Method {
				case testBroadcastCommitMethod, "broadcast_tx_sync", "broadcast_tx_async":
					broadcasts.Add(1)
					if tt.broadcastTimedOut {
						writeBroadcastRPCResponse(t, w, rpctypes.NewRPCErrorResponse(request.ID, -32603, "Internal error", timeoutErrorMessage))
						return
					}
					http.Error(w, "<html>Bad Gateway</html>", http.StatusBadGateway)
					return
				case "tx":
					count := queries.Add(1)
					var params struct {
						Hash []byte `json:"hash"`
					}
					assert.NoError(t, json.Unmarshal(request.Params, &params))
					assert.Equal(t, []byte(hash), params.Hash)
					if count == 1 && tt.firstQueryFailure {
						http.Error(w, "<html>Bad Gateway</html>", http.StatusBadGateway)
						return
					}
					if count == 1 && tt.firstQueryMissing {
						writeBroadcastRPCResponse(t, w, rpctypes.NewRPCErrorResponse(request.ID, -32603, "Internal error", fmt.Sprintf("tx (%X) not found", hash)))
						return
					}
					result = &coretypes.ResultTx{
						Hash:   hash,
						Height: 42,
						Tx:     encoded,
						TxResult: abci.ExecTxResult{
							Code: tt.txCode,
						},
					}
				case "block":
					result = &coretypes.ResultBlock{
						Block: &tmtypes.Block{
							Header: tmtypes.Header{
								Height: 42,
								Time:   time.Unix(1, 0),
							},
						},
					}
				default:
					t.Errorf("unexpected RPC method %s", request.Method)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				writeBroadcastRPCResponse(t, w, rpctypes.NewRPCSuccessResponse(request.ID, result))
			})

			cctx = cctx.WithBroadcastMode(tt.mode)
			broadcaster, _, err := newSerialTx(context.Background(), cctx.WithSignModeStr("direct"), nil)
			require.NoError(t, err)
			t.Cleanup(func() { broadcaster.lc.Shutdown(nil) })
			result, err := broadcaster.BroadcastTx(context.Background(), tx, WithBroadcastTimeout(5*time.Second), WithResultCodeAsError())
			if tt.txCode == 0 {
				require.NoError(t, err)
			} else {
				require.Error(t, err, "committed execution failures must remain failures")
			}
			response, ok := result.(*sdk.TxResponse)
			require.True(t, ok)
			require.NotNil(t, response)
			encoded, err := cctx.TxConfig.TxEncoder()(tx)
			require.NoError(t, err)
			require.Equal(t, fmt.Sprintf("%X", tmtypes.Tx(encoded).Hash()), response.TxHash)
			require.Equal(t, tt.txCode, response.Code)
			require.EqualValues(t, 1, broadcasts.Load(), "confirmation must not rebroadcast")
			expectedQueries := int32(1)
			if tt.firstQueryFailure || tt.firstQueryMissing {
				expectedQueries++
			}
			require.Equal(t, expectedQueries, queries.Load())
		})
	}
}

func TestBroadcastTxPreservesRPCResults(t *testing.T) {
	tests := []struct {
		name                 string
		rpcError             bool
		checkCode            uint32
		confirmationRejected bool
	}{
		{name: "success"},
		{
			name:      "check tx rejection",
			checkCode: 5,
		},
		{
			name:     "rpc rejection",
			rpcError: true,
		},
		{
			name:                 "confirmation RPC rejection",
			rpcError:             true,
			confirmationRejected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var calls atomic.Int32
			cctx, tx := newBroadcastTestClient(t, func(w http.ResponseWriter, _ *http.Request, request rpctypes.RPCRequest, encoded []byte) {
				count := calls.Add(1)
				if tt.confirmationRejected && count == 1 {
					assert.Equal(t, testBroadcastCommitMethod, request.Method)
					http.Error(w, "<html>Bad Gateway</html>", http.StatusBadGateway)
					return
				}
				expectedMethod := testBroadcastCommitMethod
				if tt.confirmationRejected {
					expectedMethod = "tx"
				}
				assert.Equal(t, expectedMethod, request.Method)
				if tt.rpcError {
					writeBroadcastRPCResponse(t, w, rpctypes.NewRPCErrorResponse(request.ID, -32602, "Invalid params", "rejected request"))
					return
				}
				result := &coretypes.ResultBroadcastTxCommit{
					Hash: tmtypes.Tx(encoded).Hash(),
					CheckTx: abci.ResponseCheckTx{
						Code: tt.checkCode,
					},
					Height: 42,
				}

				writeBroadcastRPCResponse(t, w, rpctypes.NewRPCSuccessResponse(request.ID, result))
			})

			broadcaster := serialBroadcaster{ctx: context.Background()}
			response, err := broadcaster.broadcastTx(context.Background(), cctx, 3*time.Second, tx)
			if tt.rpcError {
				require.Error(t, err)
				var rpcErr *rpctypes.RPCError
				require.ErrorAs(t, err, &rpcErr)
				require.Equal(t, -32602, rpcErr.Code)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.checkCode, response.Code)
			}
			expectedCalls := int32(1)
			if tt.confirmationRejected {
				expectedCalls++
			}
			require.Equal(t, expectedCalls, calls.Load(), "definitive responses must not be retried")
		})
	}
}

func TestBroadcastTxMissingTransactionDoesNotSucceed(t *testing.T) {
	var broadcasts, queries atomic.Int32
	cctx, tx := newBroadcastTestClient(t, func(w http.ResponseWriter, _ *http.Request, request rpctypes.RPCRequest, encoded []byte) {
		if request.Method == testBroadcastCommitMethod {
			broadcasts.Add(1)
			http.Error(w, "<html>Bad Gateway</html>", http.StatusBadGateway)
			return
		}
		assert.Equal(t, "tx", request.Method)
		queries.Add(1)
		writeBroadcastRPCResponse(t, w, rpctypes.NewRPCErrorResponse(request.ID, -32603, "Internal error", fmt.Sprintf("tx (%X) not found", tmtypes.Tx(encoded).Hash())))
	})
	broadcaster := serialBroadcaster{ctx: context.Background()}
	response, err := broadcaster.broadcastTx(context.Background(), cctx, 2500*time.Millisecond, tx)
	require.Error(t, err)
	require.Nil(t, response)
	require.EqualValues(t, 1, broadcasts.Load())
	require.Greater(t, queries.Load(), int32(0))
}

func TestBroadcastTxConfirmationStops(t *testing.T) {
	for _, cause := range []string{"deadline", "caller canceled", "client stopped"} {
		t.Run(cause, func(t *testing.T) {
			t.Parallel()

			queryStarted := make(chan struct{})
			queryCanceled := make(chan struct{})
			cctx, tx := newBroadcastTestClient(t, func(w http.ResponseWriter, r *http.Request, request rpctypes.RPCRequest, _ []byte) {
				if request.Method == testBroadcastCommitMethod {
					http.Error(w, "<html>Bad Gateway</html>", http.StatusBadGateway)
					return
				}
				assert.Equal(t, "tx", request.Method)
				close(queryStarted)
				<-r.Context().Done()
				close(queryCanceled)
			})

			callerCtx, callerCancel := context.WithCancel(context.Background())
			defer callerCancel()
			clientCtx, clientCancel := context.WithCancel(context.Background())
			defer clientCancel()
			broadcaster := serialBroadcaster{ctx: clientCtx}
			timeout := 10 * time.Second
			if cause == "deadline" {
				timeout = 2 * time.Second
			}
			done := make(chan error, 1)
			go func() {
				_, err := broadcaster.broadcastTx(callerCtx, cctx, timeout, tx)
				done <- err
			}()
			select {
			case <-queryStarted:
			case <-time.After(3 * time.Second):
				t.Fatal("transaction confirmation did not start")
			}
			switch cause {
			case "caller canceled":
				callerCancel()
			case "client stopped":
				clientCancel()
			}
			select {
			case err := <-done:
				require.Error(t, err)
			case <-time.After(3 * time.Second):
				t.Fatal("transaction confirmation did not stop")
			}
			select {
			case <-queryCanceled:
			case <-time.After(time.Second):
				t.Fatal("confirmation HTTP request was left running")
			}
		})
	}
}

func newBroadcastTestClient(t *testing.T, handle func(http.ResponseWriter, *http.Request, rpctypes.RPCRequest, []byte)) (sdkclient.Context, sdk.Tx) {
	t.Helper()

	txConfig := sdkutil.MakeEncodingConfig().TxConfig
	tx := txConfig.NewTxBuilder().GetTx()
	encoded, err := txConfig.TxEncoder()(tx)
	require.NoError(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request rpctypes.RPCRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handle(w, r, request, encoded)
	}))
	t.Cleanup(server.Close)

	rpc, err := rpchttp.New(server.URL, "/websocket")
	require.NoError(t, err)
	cctx := sdkclient.Context{}.
		WithClient(rpc).
		WithTxConfig(txConfig).
		WithBroadcastMode(BroadcastBlock)
	return cctx, tx
}

func writeBroadcastRPCResponse(t *testing.T, w http.ResponseWriter, response rpctypes.RPCResponse) {
	t.Helper()
	assert.NoError(t, json.NewEncoder(w).Encode(response))
}
