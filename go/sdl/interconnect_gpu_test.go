package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// AKT-492: `gpu.attributes.interconnect: []` parses as an implicit group
// named `auto`, surfaced on Resources.GPU.Attributes and lifted to the
// off-chain v2ResourceGPU.interconnectGroup field. No separate
// `interconnect=true` marker is emitted — the group key alone is the
// opt-in signal.
func TestV2ResourceGPU_InterconnectImplicitAuto(t *testing.T) {
	yamlSrc := `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: []`

	var gpu v2ResourceGPU
	require.NoError(t, yaml.Unmarshal([]byte(yamlSrc), &gpu))

	require.Equal(t, InterconnectGroupAuto, gpu.interconnectGroup,
		"implicit `interconnect: []` must lift to the literal `auto` group name")

	keys := map[string]string{}
	for _, a := range gpu.Attributes {
		keys[a.Key] = a.Value
	}
	require.Equal(t, "auto", keys[GPUAttributeInterconnectGroup],
		"on-chain interconnect/group must be the literal `auto`")
}

// AKT-492: `gpu.attributes.interconnect: { group: pair0 }` parses as an
// explicit named group. The same string appears in the on-chain attribute
// slice (under the new `interconnect/group` key) and the off-chain
// v2ResourceGPU.interconnectGroup field.
func TestV2ResourceGPU_InterconnectExplicitGroup(t *testing.T) {
	yamlSrc := `units: 8
attributes:
  vendor:
    nvidia:
      - model: a100
        ram: 80Gi
        interface: sxm
  interconnect:
    group: pair1`

	var gpu v2ResourceGPU
	require.NoError(t, yaml.Unmarshal([]byte(yamlSrc), &gpu))

	require.Equal(t, "pair1", gpu.interconnectGroup,
		"v2ResourceGPU.interconnectGroup must hold the explicit group value")

	keys := map[string]string{}
	for _, a := range gpu.Attributes {
		keys[a.Key] = a.Value
	}
	require.Equal(t, "pair1", keys[GPUAttributeInterconnectGroup],
		"on-chain interconnect/group must carry the explicit name")
}

// AKT-492: omitting `interconnect` leaves the off-chain field empty and
// the on-chain key absent — non-interconnect deployments serialize the
// same as before this feature.
func TestV2ResourceGPU_InterconnectOmitted(t *testing.T) {
	yamlSrc := `units: 8
attributes:
  vendor:
    nvidia:
      - model: a100`

	var gpu v2ResourceGPU
	require.NoError(t, yaml.Unmarshal([]byte(yamlSrc), &gpu))

	require.Empty(t, gpu.interconnectGroup)
	for _, a := range gpu.Attributes {
		require.NotEqual(t, GPUAttributeInterconnectGroup, a.Key,
			"interconnect/group must not appear when SDL omits it")
	}
}

// AKT-492: bare scalar `interconnect: true` is the retired rc4 shape.
// The parser must reject it with a clear message rather than silently
// downgrading to "no opt-in" — a tenant migrating from rc4 should see
// the failure immediately.
func TestV2ResourceGPU_InterconnectBareBoolRejected(t *testing.T) {
	yamlSrc := `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: true`

	var gpu v2ResourceGPU
	err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
	require.Error(t, err)
	require.Contains(t, err.Error(), "expected `[]` or `{group: <name>}`")
}

// AKT-492: the reserved group name `auto` cannot be written explicitly
// under `{group: ...}`. Tenants who want the implicit form must use the
// empty-sequence shape.
func TestV2ResourceGPU_InterconnectAutoNameReserved(t *testing.T) {
	yamlSrc := `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect:
    group: auto`

	var gpu v2ResourceGPU
	err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
	require.Error(t, err)
	require.Contains(t, err.Error(), "reserved name")
}

// AKT-492: `{group: ...}` must carry a non-empty group name; the empty
// mapping form is ambiguous (did the tenant mean implicit? Did they
// forget the name?) so the parser rejects it.
func TestV2ResourceGPU_InterconnectEmptyMappingRejected(t *testing.T) {
	yamlSrc := `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: {}`

	var gpu v2ResourceGPU
	err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
	require.Error(t, err)
	require.Contains(t, err.Error(), "`group` is required")
}

// AKT-492: any interconnect opt-in with gpu.units == 0 is rejected at
// parse time (CS-5). Since the group is the sole opt-in signal, this
// single guard covers both implicit and explicit forms.
func TestV2ResourceGPU_InterconnectZeroUnitsRejected(t *testing.T) {
	t.Run("implicit form with zero units", func(t *testing.T) {
		yamlSrc := `units: 0
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: []`
		var gpu v2ResourceGPU
		err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
		require.Error(t, err)
		require.Contains(t, err.Error(), "gpu.units == 0")
	})

	t.Run("explicit form with zero units", func(t *testing.T) {
		yamlSrc := `units: 0
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect:
    group: pair1`
		var gpu v2ResourceGPU
		err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
		require.Error(t, err)
		require.Contains(t, err.Error(), "gpu.units == 0")
	})

	t.Run("zero units without interconnect is fine", func(t *testing.T) {
		yamlSrc := `units: 0`
		var gpu v2ResourceGPU
		require.NoError(t, yaml.Unmarshal([]byte(yamlSrc), &gpu))
	})
}

// AKT-492: the rc4 flat `interconnect_group:` key is no longer a
// top-level SDL key — the group lives under `interconnect.group` now.
// Writing the old shape should error out under the parser's strict-key
// behavior.
func TestV2ResourceGPU_LegacyInterconnectGroupKeyRejected(t *testing.T) {
	yamlSrc := `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: []
  interconnect_group: pair1`

	var gpu v2ResourceGPU
	err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported attribute")
}

// CS-2: an unsupported attribute key under gpu.attributes still errors,
// confirming the parser's strict-key behavior is preserved alongside the
// new interconnect shape.
func TestV2GPUAttributes_UnsupportedKeyRejected(t *testing.T) {
	yamlSrc := `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
  interconnect: []
  bogus: yes`

	var gpu v2ResourceGPU
	err := yaml.Unmarshal([]byte(yamlSrc), &gpu)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported attribute")
}
