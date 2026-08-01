package v2beta3_test

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	manifest "pkg.akt.dev/go/manifest/v2beta3"
)

const testStorageKeyRef = "sealed.eyJhbGciOiJFUzI1NiJ9.eyJ2ZXJzaW9uIjoiMC4xLjAifQ.c2lnbmF0dXJl"

func TestStorageParamsKeyRefRoundTrip(t *testing.T) {
	want := &manifest.StorageParams{
		Name:     "data",
		Mount:    "/data",
		ReadOnly: true,
		KeyRef:   testStorageKeyRef,
	}

	wire, err := proto.Marshal(want)
	require.NoError(t, err)

	got := &manifest.StorageParams{}
	require.NoError(t, proto.Unmarshal(wire, got))
	require.Equal(t, want, got)
}

func TestStorageParamsKeyRefSerialization(t *testing.T) {
	params := manifest.StorageParams{KeyRef: testStorageKeyRef}

	jsonValue, err := json.Marshal(params)
	require.NoError(t, err)
	jsonObject := make(map[string]any)
	require.NoError(t, json.Unmarshal(jsonValue, &jsonObject))
	require.Equal(t, testStorageKeyRef, jsonObject["keyRef"])

	yamlValue, err := yaml.Marshal(params)
	require.NoError(t, err)
	require.Contains(t, string(yamlValue), "keyRef: "+testStorageKeyRef)

	emptyJSON, err := json.Marshal(manifest.StorageParams{})
	require.NoError(t, err)
	emptyJSONObject := make(map[string]any)
	require.NoError(t, json.Unmarshal(emptyJSON, &emptyJSONObject))
	require.NotContains(t, emptyJSONObject, "keyRef")

	emptyYAML, err := yaml.Marshal(manifest.StorageParams{})
	require.NoError(t, err)
	require.NotContains(t, string(emptyYAML), "keyRef")
}
