package mock

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	dtypes "pkg.akt.dev/go/node/deployment/v1"
	dv1beta4 "pkg.akt.dev/go/node/deployment/v1beta4"
	depositv1 "pkg.akt.dev/go/node/types/deposit/v1"
	"pkg.akt.dev/go/testutil"
)

func startTestServer(t *testing.T) *Server {
	t.Helper()
	server, err := NewServer(Config{})
	require.NoError(t, err)

	err = server.Start()
	require.NoError(t, err)
	t.Cleanup(func() { _ = server.Stop() })

	time.Sleep(100 * time.Millisecond)
	return server
}

func buildTxBytes(t *testing.T, server *Server, msg sdk.Msg) []byte {
	t.Helper()
	txBuilder := server.txConfig.NewTxBuilder()
	require.NoError(t, txBuilder.SetMsgs(msg))
	txEncoder := server.txConfig.TxEncoder()
	txBytes, err := txEncoder(txBuilder.GetTx())
	require.NoError(t, err)
	return txBytes
}

func postJSON(t *testing.T, url string, body any) *http.Response {
	t.Helper()
	jsonBody, err := json.Marshal(body)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(jsonBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	return resp
}

func getJSON(t *testing.T, url string) (int, map[string]any) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var result map[string]any
	if len(body) > 0 {
		require.NoError(t, json.Unmarshal(body, &result))
	}
	return resp.StatusCode, result
}

func createValidDeploymentMsg(t *testing.T) *dv1beta4.MsgCreateDeployment {
	t.Helper()
	did := testutil.DeploymentID(t)
	groups := []dv1beta4.GroupSpec{testutil.GroupSpec(t)}
	deposit := depositv1.Deposit{
		Amount:  testutil.AkashCoin(t, 5000),
		Sources: depositv1.Sources{depositv1.SourceBalance},
	}

	return &dv1beta4.MsgCreateDeployment{
		ID:      dtypes.DeploymentID{Owner: did.Owner, DSeq: did.DSeq},
		Groups:  groups,
		Hash:    testutil.DefaultDeploymentHash[:],
		Deposit: deposit,
	}
}

func TestServer_DeploymentQuery(t *testing.T) {
	server := startTestServer(t)

	status, result := getJSON(t, server.GatewayURL()+"/akash/deployment/v1beta4/deployments/list")
	require.Equal(t, http.StatusOK, status)
	require.Contains(t, result, "deployments")
}

func TestServer_SimulateValidTx(t *testing.T) {
	server := startTestServer(t)
	msg := createValidDeploymentMsg(t)
	txBytes := buildTxBytes(t, server, msg)

	resp := postJSON(t, server.GatewayURL()+"/cosmos/tx/v1beta1/simulate", map[string]any{
		"tx_bytes": base64.StdEncoding.EncodeToString(txBytes),
	})
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]any
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&result))
	gasInfo, ok := result["gas_info"].(map[string]any)
	require.True(t, ok)
	require.NotEmpty(t, gasInfo["gas_wanted"])
}

func TestServer_SimulateDoesNotRecord(t *testing.T) {
	server := startTestServer(t)
	msg := createValidDeploymentMsg(t)
	txBytes := buildTxBytes(t, server, msg)

	resp := postJSON(t, server.GatewayURL()+"/cosmos/tx/v1beta1/simulate", map[string]any{
		"tx_bytes": base64.StdEncoding.EncodeToString(txBytes),
	})
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	status, _ := getJSON(t, server.GatewayURL()+"/mock/last-deployment")
	require.Equal(t, http.StatusNotFound, status)
}

func TestServer_BroadcastRecordsDeployment(t *testing.T) {
	server := startTestServer(t)
	msg := createValidDeploymentMsg(t)
	txBytes := buildTxBytes(t, server, msg)

	resp := postJSON(t, server.GatewayURL()+"/cosmos/tx/v1beta1/txs", map[string]any{
		"tx_bytes": base64.StdEncoding.EncodeToString(txBytes),
		"mode":     "BROADCAST_MODE_SYNC",
	})
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]any
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&result))
	txResp, ok := result["tx_response"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, float64(0), txResp["code"])
	require.NotEmpty(t, txResp["txhash"])

	status, deployment := getJSON(t, server.GatewayURL()+"/mock/last-deployment")
	require.Equal(t, http.StatusOK, status)

	id, ok := deployment["id"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, msg.ID.Owner, id["owner"])
}

func TestServer_RejectsInvalidDeployment(t *testing.T) {
	server := startTestServer(t)
	msg := &dv1beta4.MsgCreateDeployment{
		ID:     dtypes.DeploymentID{Owner: testutil.AccAddress(t).String(), DSeq: 1},
		Groups: []dv1beta4.GroupSpec{},
		Hash:   testutil.DefaultDeploymentHash[:],
		Deposit: depositv1.Deposit{
			Amount:  testutil.AkashCoin(t, 5000),
			Sources: depositv1.Sources{depositv1.SourceBalance},
		},
	}
	txBytes := buildTxBytes(t, server, msg)

	resp := postJSON(t, server.GatewayURL()+"/cosmos/tx/v1beta1/txs", map[string]any{
		"tx_bytes": base64.StdEncoding.EncodeToString(txBytes),
		"mode":     "BROADCAST_MODE_SYNC",
	})
	defer resp.Body.Close()

	require.NotEqual(t, http.StatusOK, resp.StatusCode)
}

func TestServer_RejectsInvalidTxBytes(t *testing.T) {
	server := startTestServer(t)

	resp := postJSON(t, server.GatewayURL()+"/cosmos/tx/v1beta1/txs", map[string]any{
		"tx_bytes": base64.StdEncoding.EncodeToString([]byte("invalid-tx-bytes")),
		"mode":     "BROADCAST_MODE_SYNC",
	})
	defer resp.Body.Close()

	require.NotEqual(t, http.StatusOK, resp.StatusCode)
}

func TestServer_DebugEndpointsReturnNotFoundInitially(t *testing.T) {
	server := startTestServer(t)

	endpoints := []string{
		"/mock/last-deployment",
		"/mock/last-bid",
		"/mock/last-lease",
		"/mock/last-close-bid",
	}

	for _, ep := range endpoints {
		status, _ := getJSON(t, server.GatewayURL()+ep)
		require.Equal(t, http.StatusNotFound, status, "expected 404 for %s", ep)
	}
}
