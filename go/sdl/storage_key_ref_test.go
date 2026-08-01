package sdl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testSDLVersion2   = "2.0"
	testSDLVersion2_1 = "2.1"
	testStorageKeyRef = "sealed.eyJhbGciOiJFUzI1NiJ9.eyJ2ZXJzaW9uIjoiMC4xLjAifQ.c2lnbmF0dXJl"
)

var testSDLVersions = []string{testSDLVersion2, testSDLVersion2_1}

func TestStorageKeyRefValidation(t *testing.T) {
	tests := []struct {
		name       string
		keyRef     string
		includeRef bool
		persistent bool
		wantErr    string
	}{
		{
			name:       "sealed reference on persistent storage",
			keyRef:     testStorageKeyRef,
			includeRef: true,
			persistent: true,
		},
		{
			name:       "missing reference remains backward compatible",
			persistent: true,
		},
		{
			name:       "empty reference is treated as absent",
			keyRef:     "",
			includeRef: true,
			persistent: false,
		},
		{
			name:       "plain KBS URI",
			keyRef:     "kbs:///default/storage-dek/example",
			includeRef: true,
			persistent: true,
			wantErr:    "storage keyRef must use sealed.<JWS> format",
		},
		{
			name:       "incomplete compact JWS",
			keyRef:     "sealed.header.payload",
			includeRef: true,
			persistent: true,
			wantErr:    "storage keyRef must use sealed.<JWS> format",
		},
		{
			name:       "oversized sealed reference",
			keyRef:     "sealed." + strings.Repeat("a", StorageKeyRefMaxBytes) + ".payload.signature",
			includeRef: true,
			persistent: true,
			wantErr:    "storage keyRef exceeds 65536 bytes",
		},
		{
			name:       "sealed reference on ephemeral storage",
			keyRef:     testStorageKeyRef,
			includeRef: true,
			persistent: false,
			wantErr:    "storage keyRef requires persistent storage",
		},
	}

	for _, version := range testSDLVersions {
		for _, test := range tests {
			t.Run(fmt.Sprintf("v%s/%s", version, test.name), func(t *testing.T) {
				sdl, err := Read(storageKeyRefSDL(version, test.keyRef, test.includeRef, test.persistent))
				if test.wantErr != "" {
					require.ErrorContains(t, err, test.wantErr)
					return
				}

				require.NoError(t, err)
				result, err := sdl.Manifest()
				require.NoError(t, err)
				require.Equal(t, test.keyRef, result.GetGroups()[0].Services[0].Params.Storage[0].KeyRef)
			})
		}
	}
}

func TestStorageKeyRefSchema(t *testing.T) {
	tests := []struct {
		name    string
		keyRef  string
		wantErr bool
	}{
		{name: "sealed compact JWS", keyRef: testStorageKeyRef},
		{name: "explicit empty value", keyRef: ""},
		{name: "plain KBS URI", keyRef: "kbs:///default/storage-dek/example", wantErr: true},
		{name: "incomplete compact JWS", keyRef: "sealed.header.payload", wantErr: true},
		{name: "oversized sealed reference", keyRef: "sealed." + strings.Repeat("a", StorageKeyRefMaxBytes) + ".payload.signature", wantErr: true},
	}

	for _, version := range testSDLVersions {
		for _, test := range tests {
			t.Run(fmt.Sprintf("v%s/%s", version, test.name), func(t *testing.T) {
				err := validateInputAgainstSchema(storageKeyRefSDL(version, test.keyRef, true, true))
				if test.wantErr {
					require.Error(t, err)
					return
				}

				require.NoError(t, err)
			})
		}
	}
}

func storageKeyRefSDL(version, keyRef string, includeKeyRef, persistent bool) []byte {
	keyRefLine := ""
	if includeKeyRef {
		keyRefLine = fmt.Sprintf("\n          keyRef: %q", keyRef)
	}

	attributes := ""
	if persistent {
		attributes = `
            attributes:
              class: default
              persistent: true`
	}

	return []byte(fmt.Sprintf(`version: %q
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - global: true
    params:
      storage:
        data:
          mount: /data%s
profiles:
  compute:
    web:
      resources:
        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - name: data
            size: 1Gi%s
  placement:
    dc:
      pricing:
        web:
          denom: uakt
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: 1
`, version, keyRefLine, attributes))
}
