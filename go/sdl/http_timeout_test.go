package sdl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPTimeoutDurations(t *testing.T) {
	tests := []struct {
		name         string
		readTimeout  string
		sendTimeout  string
		expectedRead uint32
		expectedSend uint32
	}{
		{
			name:         "legacy millisecond integers",
			readTimeout:  "60001",
			sendTimeout:  "3600000",
			expectedRead: 60_001,
			expectedSend: 3_600_000,
		},
		{
			name:         "duration strings",
			readTimeout:  "60s",
			sendTimeout:  "1h",
			expectedRead: 60_000,
			expectedSend: 3_600_000,
		},
		{
			name:         "explicit millisecond durations",
			readTimeout:  "500ms",
			sendTimeout:  "1000ms",
			expectedRead: 500,
			expectedSend: 1_000,
		},
		{
			name:         "quoted values without units are milliseconds",
			readTimeout:  `"60000"`,
			sendTimeout:  `"3600000"`,
			expectedRead: 60_000,
			expectedSend: 3_600_000,
		},
		{
			name:         "manifest maximum",
			readTimeout:  "4294967295",
			sendTimeout:  `"4294967295"`,
			expectedRead: 4_294_967_295,
			expectedSend: 4_294_967_295,
		},
	}

	for _, version := range []string{"2.0", "2.1"} {
		for _, test := range tests {
			t.Run(version+"/"+test.name, func(t *testing.T) {
				sdl, err := Read(httpTimeoutSDL(version, test.readTimeout, test.sendTimeout))
				require.NoError(t, err)

				mani, err := sdl.Manifest()
				require.NoError(t, err)

				httpOptions := mani.GetGroups()[0].Services[0].Expose[0].HTTPOptions
				assert.Equal(t, test.expectedRead, httpOptions.ReadTimeout)
				assert.Equal(t, test.expectedSend, httpOptions.SendTimeout)
			})
		}
	}
}

func TestHTTPTimeoutRejectsValuesAboveManifestMaximum(t *testing.T) {
	_, err := Read(httpTimeoutSDL("2.0", "1194h", "60s"))
	require.Error(t, err)
	assert.ErrorContains(t, err, "read timeout cannot be greater than 4294967295 ms")
}

func httpTimeoutSDL(version, readTimeout, sendTimeout string) []byte {
	return []byte(sdlTestBuilder{
		version: fmt.Sprintf(`version: %q`, version),
		exposeBlock: fmt.Sprintf(`    expose:
      - port: 80
        http_options:
          read_timeout: %s
          send_timeout: %s
        to:
          - global: true`, readTimeout, sendTimeout),
	}.build())
}
