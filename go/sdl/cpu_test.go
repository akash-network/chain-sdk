package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestV2ResourceCPU_Valid(t *testing.T) {
	var stream = `
units: 0.1
attributes:
  arch: amd64
`
	var p v2ResourceCPU

	err := yaml.Unmarshal([]byte(stream), &p)
	require.NoError(t, err)
	require.Equal(t, cpuQuantity(100), p.Units)
	require.Equal(t, 1, len(p.Attributes))
	require.Equal(t, "arch", p.Attributes[0].Key)
	require.Equal(t, "amd64", p.Attributes[0].Value)
}

func TestV2ResourceCPU_NoArch(t *testing.T) {
	var stream = `
units: 0.1
`
	var p v2ResourceCPU

	err := yaml.Unmarshal([]byte(stream), &p)
	require.NoError(t, err)
	require.Equal(t, cpuQuantity(100), p.Units)
	require.Equal(t, 0, len(p.Attributes))
}

func TestV2ResourceCPU_Arm64(t *testing.T) {
	var stream = `
units: 0.1
attributes:
  arch: arm64
`
	var p v2ResourceCPU

	err := yaml.Unmarshal([]byte(stream), &p)
	require.NoError(t, err)
	require.Equal(t, 1, len(p.Attributes))
	require.Equal(t, "arch", p.Attributes[0].Key)
	require.Equal(t, "arm64", p.Attributes[0].Value)
}

// The architecture whitelist has to reject the same values the `arch` enum in
// sdl-input.schema.yaml rejects, so both SDKs answer the same for one SDL.
func TestV2ResourceCPU_UnsupportedArch(t *testing.T) {
	var stream = `
units: 0.1
attributes:
  arch: sparc64
`
	var p v2ResourceCPU

	err := yaml.Unmarshal([]byte(stream), &p)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported cpu architecture")
}

func TestV2ResourceCPU_UnsupportedAttribute(t *testing.T) {
	var stream = `
units: 0.1
attributes:
  vendor: intel
`
	var p v2ResourceCPU

	err := yaml.Unmarshal([]byte(stream), &p)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported cpu attribute")
}
