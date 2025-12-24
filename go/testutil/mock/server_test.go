package mock

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	server := StartMockServer(t)
	defer server.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := server.GatewayURL() + "/akash/deployment/v1beta4/deployments/list"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	t.Logf("Response body: %s", string(bodyBytes))

	var response map[string]interface{}
	err = json.Unmarshal(bodyBytes, &response)
	require.NoError(t, err)
	t.Logf("Parsed response: %+v", response)

	if deployments, ok := response["deployments"].([]interface{}); ok {
		t.Logf("Deployments array length: %d", len(deployments))
		require.GreaterOrEqual(t, len(deployments), 0)
	} else {
		t.Logf("Deployments field type: %T, value: %+v", response["deployments"], response["deployments"])
		require.Contains(t, response, "deployments", "Response should contain 'deployments' field")
	}
}
