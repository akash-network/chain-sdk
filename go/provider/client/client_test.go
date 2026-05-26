package rest

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/gorilla/websocket"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	mtypes "pkg.akt.dev/go/node/market/v1"
	_ "pkg.akt.dev/go/sdkutil"
)

func TestNewClientWithProviderURL(t *testing.T) {
	ctx := context.Background()
	providerURL := "https://example.com:8443"
	addr, err := sdk.AccAddressFromBech32("akash1365yvmc4s7awdyj3n2sav7xfx76adc6dnmlx63")
	require.NoError(t, err)

	t.Run("basic functionality", func(t *testing.T) {
		cl, err := NewClient(ctx, addr, WithProviderURL(providerURL))
		require.NoError(t, err)
		require.NotNil(t, cl)

		c := cl.(*client)
		require.Nil(t, c.cclient)
		require.Equal(t, ctx, c.ctx)
		require.Equal(t, addr, c.addr)
		require.NotNil(t, c.host)
		require.NotNil(t, c.tlsCfg)
	})

	t.Run("options are executed", func(t *testing.T) {
		token := "test-token"

		cl, err := NewClient(ctx, addr, WithProviderURL(providerURL), WithAuthToken(token))
		require.NoError(t, err)

		c := cl.(*client)
		require.Nil(t, c.cclient)
		require.Equal(t, token, c.opts.token)
		require.Equal(t, providerURL, c.opts.providerURL)
	})

	t.Run("invalid URL", func(t *testing.T) {
		invalidURL := "://invalid-url"
		_, err := NewClient(ctx, addr, WithProviderURL(invalidURL))
		require.Error(t, err)
	})

	t.Run("option error handling", func(t *testing.T) {
		testError := errors.New("test error")
		errorOption := func(*clientOptions) error {
			return testError
		}

		_, err := NewClient(ctx, addr, WithProviderURL(providerURL), errorOption)
		require.Error(t, err)
		require.Equal(t, testError, err)
	})

	t.Run("JWT auth without cert querier", func(t *testing.T) {
		// Test that client can be created without cert querier (for JWT auth)
		cl, err := NewClient(ctx, addr, WithProviderURL(providerURL))
		require.NoError(t, err)

		c := cl.(*client)
		require.Nil(t, c.opts.signer)  // Should be nil when no signer provided
		require.Empty(t, c.opts.token) // Should be empty when no token provided
	})

}

// captureQueryServer spins up an httptest.NewTLSServer that records the query
// string of the first websocket upgrade request, then closes the connection.
// The provider gateway uses singular `service` and `tail` query params; these
// tests assert the client sends exactly those keys (and not the legacy plural
// `services`).
//
// Cross-reference: akash-network/provider gateway/rest/middleware.go
// requestStreamParams reads vars.Get("service") and vars.Get("tail").
func captureQueryServer(t *testing.T) (*httptest.Server, func() url.Values) {
	t.Helper()
	var (
		mu sync.Mutex
		q  url.Values
	)
	upgrader := websocket.Upgrader{
		CheckOrigin: func(*http.Request) bool { return true },
	}
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		q = r.URL.Query()
		mu.Unlock()

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		_ = conn.Close()
	}))
	return srv, func() url.Values {
		mu.Lock()
		defer mu.Unlock()
		return q
	}
}

// newTestClient builds a *client pointed at the test server's URL with a TLS
// config that trusts the httptest self-signed certificate. Same-package
// access lets us replace tlsCfg directly, avoiding any production-surface
// change.
func newTestClient(t *testing.T, ctx context.Context, providerURL string) *client {
	t.Helper()
	addr, err := sdk.AccAddressFromBech32("akash1365yvmc4s7awdyj3n2sav7xfx76adc6dnmlx63")
	require.NoError(t, err)
	cl, err := NewClient(ctx, addr, WithProviderURL(providerURL))
	require.NoError(t, err)
	c := cl.(*client)
	c.tlsCfg = &tls.Config{InsecureSkipVerify: true} // nolint: gosec // httptest self-signed cert
	return c
}

func testLeaseID() mtypes.LeaseID {
	const akashAddr = "akash1365yvmc4s7awdyj3n2sav7xfx76adc6dnmlx63"
	return mtypes.LeaseID{
		Owner:    akashAddr,
		DSeq:     1,
		GSeq:     1,
		OSeq:     1,
		Provider: akashAddr,
	}
}

func TestLeaseLogs_ForwardsServiceAndTailToGateway(t *testing.T) {
	srv, getQuery := captureQueryServer(t)
	defer srv.Close()

	ctx := context.Background()
	c := newTestClient(t, ctx, srv.URL)
	lid := testLeaseID()

	out, err := c.LeaseLogs(ctx, lid, "db-replica,api", true, 50)
	require.NoError(t, err)
	require.NotNil(t, out)
	<-out.OnClose // wait for reader goroutine to drain after server close

	q := getQuery()
	require.Equal(t, "db-replica,api", q.Get("service"),
		"gateway requestStreamParams reads vars.Get(\"service\") — must be singular")
	require.Equal(t, "50", q.Get("tail"),
		"tailLines must be forwarded as ?tail=N")
	require.Equal(t, "true", q.Get("follow"))
	require.Empty(t, q.Get("services"),
		"legacy plural 'services' key is silently dropped by the gateway — must not be sent")
}

func TestLeaseLogs_OmitsEmptyServiceAndZeroTail(t *testing.T) {
	srv, getQuery := captureQueryServer(t)
	defer srv.Close()

	ctx := context.Background()
	c := newTestClient(t, ctx, srv.URL)
	lid := testLeaseID()

	out, err := c.LeaseLogs(ctx, lid, "", false, 0)
	require.NoError(t, err)
	<-out.OnClose

	q := getQuery()
	require.Empty(t, q.Get("service"), "empty service filter must not be sent")
	require.Empty(t, q.Get("tail"), "tail <= 0 must not be sent (gateway default is -1)")
	require.Equal(t, "false", q.Get("follow"))
}

func TestLeaseEvents_ForwardsServiceToGateway(t *testing.T) {
	srv, getQuery := captureQueryServer(t)
	defer srv.Close()

	ctx := context.Background()
	c := newTestClient(t, ctx, srv.URL)
	lid := testLeaseID()

	out, err := c.LeaseEvents(ctx, lid, "db-replica", true)
	require.NoError(t, err)
	require.NotNil(t, out)
	<-out.OnClose

	q := getQuery()
	require.Equal(t, "db-replica", q.Get("service"))
	require.Equal(t, "true", q.Get("follow"))
	require.Empty(t, q.Get("services"))
}

func TestLeaseEvents_OmitsEmptyService(t *testing.T) {
	srv, getQuery := captureQueryServer(t)
	defer srv.Close()

	ctx := context.Background()
	c := newTestClient(t, ctx, srv.URL)
	lid := testLeaseID()

	out, err := c.LeaseEvents(ctx, lid, "", false)
	require.NoError(t, err)
	<-out.OnClose

	q := getQuery()
	require.Empty(t, q.Get("service"))
}
