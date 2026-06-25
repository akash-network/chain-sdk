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

func TestV2ResourceCPU_TopLevelArch(t *testing.T) {
	var stream = `
units: 0.5
arch: arm64
`
	var p v2ResourceCPU

	err := yaml.Unmarshal([]byte(stream), &p)
	require.NoError(t, err)
	require.Equal(t, cpuQuantity(500), p.Units)
	require.Equal(t, 1, len(p.Attributes))
	require.Equal(t, "arch", p.Attributes[0].Key)
	require.Equal(t, "arm64", p.Attributes[0].Value)
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

func TestV2ResourceCPU_InvalidArch(t *testing.T) {
	var stream = `
units: 0.1
arch: x86_64
`
	var p v2ResourceCPU

	err := yaml.Unmarshal([]byte(stream), &p)
	require.Error(t, err)
}

func TestV2ResourceCPU_ArchConflict(t *testing.T) {
	var stream = `
units: 0.1
arch: arm64
attributes:
  arch: amd64
`
	var p v2ResourceCPU

	err := yaml.Unmarshal([]byte(stream), &p)
	require.Error(t, err)
}
