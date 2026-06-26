package rest

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestClientAttestationQuote(t *testing.T) {
	ctx := context.Background()
	addr, err := sdk.AccAddressFromBech32("akash1365yvmc4s7awdyj3n2sav7xfx76adc6dnmlx63")
	require.NoError(t, err)

	expectedBody := []byte(`{"nonce":"test-nonce"}`)
	expectedResponse := []byte(`{"report":"test-report"}`)
	lid := mtypes.LeaseID{
		DSeq: 1,
		GSeq: 2,
		OSeq: 3,
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/lease/1/2/3/attestation/quote", r.URL.Path)
		require.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		require.Equal(t, contentTypeJSON, r.Header.Get("Content-Type"))

		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.Equal(t, expectedBody, body)

		w.Header().Set("Content-Type", contentTypeJSON)
		_, err = w.Write(expectedResponse)
		require.NoError(t, err)
	}))
	defer srv.Close()

	cl, err := NewClient(ctx, addr, WithProviderURL(srv.URL), WithAuthToken("test-token"))
	require.NoError(t, err)

	resp, err := cl.AttestationQuote(ctx, lid, expectedBody)
	require.NoError(t, err)
	require.Equal(t, expectedResponse, resp)
}
