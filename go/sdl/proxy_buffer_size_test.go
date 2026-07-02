package sdl

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	manifest "pkg.akt.dev/go/manifest/v2beta3"
)

// TestProxyBufferSize_NotDefaultedAtManifestLevel guards backward compatibility:
// proxy_buffer_size must NOT be given a non-zero default when translating the SDL
// to a manifest. A non-zero default would enter the manifest version hash and break
// deployment submission to providers that predate the field. The provider applies
// its own default when the value is 0.
func TestProxyBufferSize_NotDefaultedAtManifestLevel(t *testing.T) {
	// unset -> stays 0 (no manifest-level default)
	m, err := v2HTTPOptions{}.asManifest()
	require.NoError(t, err)
	require.Equal(t, uint32(0), m.ProxyBufferSize)

	// explicit value is preserved
	m, err = v2HTTPOptions{ProxyBufferSize: 32768}.asManifest()
	require.NoError(t, err)
	require.Equal(t, uint32(32768), m.ProxyBufferSize)

	// over the upper limit is rejected
	_, err = v2HTTPOptions{ProxyBufferSize: upperLimitProxyBufferSize + 1}.asManifest()
	require.Error(t, err)

	// at the upper limit is allowed
	m, err = v2HTTPOptions{ProxyBufferSize: upperLimitProxyBufferSize}.asManifest()
	require.NoError(t, err)
	require.Equal(t, upperLimitProxyBufferSize, m.ProxyBufferSize)
}

// TestProxyBufferSize_OmittedFromManifestJSONWhenZero is the core backward-compat
// guard. Manifest.Version() hashes the marshaled manifest JSON, so a zero
// proxy_buffer_size must be omitted entirely — otherwise manifests differ from
// those produced before the field existed and old providers reject the version.
func TestProxyBufferSize_OmittedFromManifestJSONWhenZero(t *testing.T) {
	b, err := json.Marshal(manifest.ServiceExposeHTTPOptions{MaxBodySize: 1})
	require.NoError(t, err)
	require.NotContains(t, string(b), "proxyBufferSize",
		"zero proxy_buffer_size must be omitted from manifest JSON for version-hash backward compatibility")

	b, err = json.Marshal(manifest.ServiceExposeHTTPOptions{MaxBodySize: 1, ProxyBufferSize: 32768})
	require.NoError(t, err)
	require.Contains(t, string(b), "proxyBufferSize",
		"explicitly set proxy_buffer_size must appear in manifest JSON")
}
