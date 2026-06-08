package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// CS-2: gpu.attributes.rdma: true flows to on-chain Resources.GPU.attributes
// as a free-form key=value pair, while rdma: false (or unset) is absent.
func TestV2ResourceGPU_RDMAFlag(t *testing.T) {
	tests := []struct {
		name           string
		yaml           string
		expectRDMAAttr bool
		expectGroup    string
	}{
		{
			name: "rdma true emits attribute",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  rdma: true`,
			expectRDMAAttr: true,
			expectGroup:    "",
		},
		{
			name: "rdma false does not emit attribute",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  rdma: false`,
			expectRDMAAttr: false,
			expectGroup:    "",
		},
		{
			name: "no rdma key behaves like rdma false",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100`,
			expectRDMAAttr: false,
			expectGroup:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var gpu v2ResourceGPU
			require.NoError(t, yaml.Unmarshal([]byte(tc.yaml), &gpu))

			hasRDMA := false
			for _, a := range gpu.Attributes {
				if a.Key == GPUAttributeRDMA && a.Value == "true" {
					hasRDMA = true
				}
			}
			require.Equal(t, tc.expectRDMAAttr, hasRDMA,
				"unexpected presence of on-chain rdma=true attribute")
			require.Equal(t, tc.expectGroup, gpu.RDMAGroup)
		})
	}
}

// AKT-443: rdma_group flows end-to-end. It appears in the on-chain GPU
// attribute slice (so the provider's bid engine can enforce per-group
// node separation during reservation) AND is lifted onto
// v2ResourceGPU.RDMAGroup for the manifest builder to route into
// Service.RDMAGroup (so the workload builder can label pods for
// anti-affinity). Both consumers see the same value.
func TestV2ResourceGPU_RDMAGroupOnChainAndOffChain(t *testing.T) {
	yamlSrc := `units: 8
attributes:
  vendor:
    nvidia:
      - model: a100
        ram: 80Gi
        interface: sxm
  rdma: true
  rdma_group: pair1`

	var gpu v2ResourceGPU
	require.NoError(t, yaml.Unmarshal([]byte(yamlSrc), &gpu))

	// Off-chain: dedicated field for the manifest builder.
	require.Equal(t, "pair1", gpu.RDMAGroup,
		"v2ResourceGPU.RDMAGroup must hold the rdma_group value")

	// On-chain: present in the GPU attribute slice alongside rdma=true.
	keys := map[string]string{}
	for _, a := range gpu.Attributes {
		keys[a.Key] = a.Value
	}
	require.Equal(t, "true", keys[GPUAttributeRDMA])
	require.Equal(t, "pair1", keys[GPUAttributeRDMAGroup],
		"rdma_group must appear in on-chain GPU attributes for bid-engine group tracking")
}

// AKT-443 (continued): omitting rdma_group leaves both the field empty
// and the on-chain attribute absent — non-RDMA-group services produce
// the same byte-for-byte serialization they did before this feature.
func TestV2ResourceGPU_RDMAGroupOmitted(t *testing.T) {
	yamlSrc := `units: 8
attributes:
  vendor:
    nvidia:
      - model: a100
  rdma: true`

	var gpu v2ResourceGPU
	require.NoError(t, yaml.Unmarshal([]byte(yamlSrc), &gpu))

	require.Empty(t, gpu.RDMAGroup)
	for _, a := range gpu.Attributes {
		require.NotEqual(t, GPUAttributeRDMAGroup, a.Key,
			"rdma_group must not appear when SDL omits it")
	}
}

// CodeRabbit follow-up: a profile with gpu.units == 0 that declares
// rdma: true or rdma_group: <name> is a misconfiguration — there is no
// HCA to allocate, so the RDMA flags are meaningless. The parser must
// reject this fail-fast rather than letting downstream validation passes
// silently classify the profile as RDMA-enabled.
func TestV2ResourceGPU_RDMAZeroUnitsRejected(t *testing.T) {
	t.Run("rdma true with zero units", func(t *testing.T) {
		yamlSrc := `units: 0
attributes:
  vendor:
    nvidia:
      - model: a100
  rdma: true`
		var gpu v2ResourceGPU
		err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
		require.Error(t, err)
		require.Contains(t, err.Error(), "gpu.attributes.rdma cannot be set when gpu.units == 0")
	})

	t.Run("rdma_group with zero units", func(t *testing.T) {
		yamlSrc := `units: 0
attributes:
  vendor:
    nvidia:
      - model: a100
  rdma_group: pair1`
		var gpu v2ResourceGPU
		err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
		require.Error(t, err)
		require.Contains(t, err.Error(), "rdma_group")
		require.Contains(t, err.Error(), "gpu.units == 0")
	})

	t.Run("zero units without rdma is fine", func(t *testing.T) {
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
// new rdma/rdma_group handling.
func TestV2GPUAttributes_UnsupportedKeyRejected(t *testing.T) {
	yamlSrc := `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  rdma: true
  bogus: yes`

	var gpu v2ResourceGPU
	err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported attribute")
}
