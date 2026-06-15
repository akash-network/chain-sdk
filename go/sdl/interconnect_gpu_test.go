package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// CS-2: gpu.attributes.interconnect: true flows to on-chain Resources.GPU.attributes
// as a free-form key=value pair, while interconnect: false (or unset) is absent.
func TestV2ResourceGPU_InterconnectFlag(t *testing.T) {
	tests := []struct {
		name           string
		yaml           string
		expectInterconnectAttr bool
		expectGroup    string
	}{
		{
			name: "interconnect true emits attribute",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: true`,
			expectInterconnectAttr: true,
			expectGroup:    "",
		},
		{
			name: "interconnect false does not emit attribute",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: false`,
			expectInterconnectAttr: false,
			expectGroup:    "",
		},
		{
			name: "no interconnect key behaves like interconnect false",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100`,
			expectInterconnectAttr: false,
			expectGroup:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var gpu v2ResourceGPU
			require.NoError(t, yaml.Unmarshal([]byte(tc.yaml), &gpu))

			hasInterconnect := false
			for _, a := range gpu.Attributes {
				if a.Key == GPUAttributeInterconnect && a.Value == "true" {
					hasInterconnect = true
				}
			}
			require.Equal(t, tc.expectInterconnectAttr, hasInterconnect,
				"unexpected presence of on-chain interconnect=true attribute")
			require.Equal(t, tc.expectGroup, gpu.InterconnectGroup)
		})
	}
}

// AKT-443: interconnect_group flows end-to-end. It appears in the on-chain GPU
// attribute slice (so the provider's bid engine can enforce per-group
// node separation during reservation) AND is lifted onto
// v2ResourceGPU.InterconnectGroup for the manifest builder to route into
// Service.InterconnectGroup (so the workload builder can label pods for
// anti-affinity). Both consumers see the same value.
func TestV2ResourceGPU_InterconnectGroupOnChainAndOffChain(t *testing.T) {
	yamlSrc := `units: 8
attributes:
  vendor:
    nvidia:
      - model: a100
        ram: 80Gi
        interface: sxm
  interconnect: true
  interconnect_group: pair1`

	var gpu v2ResourceGPU
	require.NoError(t, yaml.Unmarshal([]byte(yamlSrc), &gpu))

	// Off-chain: dedicated field for the manifest builder.
	require.Equal(t, "pair1", gpu.InterconnectGroup,
		"v2ResourceGPU.InterconnectGroup must hold the interconnect_group value")

	// On-chain: present in the GPU attribute slice alongside interconnect=true.
	keys := map[string]string{}
	for _, a := range gpu.Attributes {
		keys[a.Key] = a.Value
	}
	require.Equal(t, "true", keys[GPUAttributeInterconnect])
	require.Equal(t, "pair1", keys[GPUAttributeInterconnectGroup],
		"interconnect_group must appear in on-chain GPU attributes for bid-engine group tracking")
}

// AKT-443 (continued): omitting interconnect_group leaves both the field empty
// and the on-chain attribute absent — non-interconnect-group services produce
// the same byte-for-byte serialization they did before this feature.
func TestV2ResourceGPU_InterconnectGroupOmitted(t *testing.T) {
	yamlSrc := `units: 8
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: true`

	var gpu v2ResourceGPU
	require.NoError(t, yaml.Unmarshal([]byte(yamlSrc), &gpu))

	require.Empty(t, gpu.InterconnectGroup)
	for _, a := range gpu.Attributes {
		require.NotEqual(t, GPUAttributeInterconnectGroup, a.Key,
			"interconnect_group must not appear when SDL omits it")
	}
}

// CodeRabbit follow-up: a profile with gpu.units == 0 that declares
// interconnect: true or interconnect_group: <name> is a misconfiguration — there is no
// HCA to allocate, so the interconnect flags are meaningless. The parser must
// reject this fail-fast rather than letting downstream validation passes
// silently classify the profile as interconnect-enabled.
func TestV2ResourceGPU_InterconnectZeroUnitsRejected(t *testing.T) {
	t.Run("interconnect true with zero units", func(t *testing.T) {
		yamlSrc := `units: 0
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: true`
		var gpu v2ResourceGPU
		err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
		require.Error(t, err)
		require.Contains(t, err.Error(), "gpu.attributes.interconnect cannot be set when gpu.units == 0")
	})

	t.Run("interconnect_group with zero units", func(t *testing.T) {
		yamlSrc := `units: 0
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect_group: pair1`
		var gpu v2ResourceGPU
		err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
		require.Error(t, err)
		require.Contains(t, err.Error(), "interconnect_group")
		require.Contains(t, err.Error(), "gpu.units == 0")
	})

	t.Run("zero units without interconnect is fine", func(t *testing.T) {
		// Verifies the new guards don't accidentally break the existing
		// path where a profile has gpu.units == 0 and no attributes at all
		// (the parser leaves it alone).
		yamlSrc := `units: 0`
		var gpu v2ResourceGPU
		require.NoError(t, yaml.Unmarshal([]byte(yamlSrc), &gpu))
	})
}

// CS-2: an unsupported attribute key under gpu.attributes still errors,
// confirming the parser's strict-key behavior is preserved alongside the
// new interconnect_group handling.
func TestV2GPUAttributes_UnsupportedKeyRejected(t *testing.T) {
	yamlSrc := `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: true
  bogus: yes`

	var gpu v2ResourceGPU
	err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported attribute")
}
