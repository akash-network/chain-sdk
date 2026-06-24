package v1beta4

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	attr "pkg.akt.dev/go/node/types/attributes/v1"
	rtypes "pkg.akt.dev/go/node/types/resources/v1beta4"
)

// CS-6: the provider reservation/commit path on the provider side rebuilds
// GroupSpec instances from parts and feeds the rebuilt value into the
// inventory.Adjust path. The provider's interconnect-detection logic reads
//
//   - GroupSpec.Requirements.Attributes  (carrying capabilities/gpu-interconnect=true)
//   - Resources.GPU.Attributes           (carrying interconnect=true per resource)
//
// from whatever ResourceGroup-typed value lands on its `reservation.Resources()`
// call. If anything in the chain SDK's Dup / accessor chain drops either of
// those attribute slices, the provider silently treats the order as
// non-interconnect. These tests pin down preservation for every concrete
// ResourceGroup the provider may hold:
//
//   - *Group         (the bid-engine path stores a *dtypes.Group)
//   - Group          (value form, dereferenced)
//   - *GroupSpec     (used by tests and by some legacy adapter paths)
//   - GroupSpec      (value form returned by Group.GroupSpec)
//
// We assert that a representative Group carrying capabilities/gpu-interconnect=true on
// Requirements and interconnect=true on the first resource's GPU.Attributes survives:
//
//   1. (Group).Dup()-equivalent — via GroupSpec.Dup().
//   2. Accessing the value via the ResourceGroup interface (GetResourceUnits).
//   3. Direct field reads of Requirements.Attributes from any of the four
//      concrete shapes.

func interconnectSampleGroupSpec() GroupSpec {
	return GroupSpec{
		Name: "ib",
		Requirements: attr.PlacementRequirements{
			Attributes: attr.Attributes{
				{Key: "capabilities/gpu-interconnect", Value: "true"},
				{Key: "capabilities/gpu-interconnect/fabric/infiniband", Value: "true"},
			},
		},
		Resources: ResourceUnits{
			{
				Resources: rtypes.Resources{
					ID: 1,
					CPU: &rtypes.CPU{
						Units: rtypes.NewResourceValue(1000),
					},
					Memory: &rtypes.Memory{
						Quantity: rtypes.NewResourceValue(1024 * 1024 * 1024),
					},
					GPU: &rtypes.GPU{
						Units: rtypes.NewResourceValue(8),
						Attributes: attr.Attributes{
							{Key: "interconnect", Value: "true"},
							{Key: "vendor/nvidia/model/a100", Value: "true"},
						},
					},
					Storage: rtypes.Volumes{},
				},
				Count: 1,
				Price: sdk.NewDecCoin("uact", sdkmath.NewInt(1)),
			},
		},
	}
}

func assertInterconnectSignalsPresent(t *testing.T, where string, reqs attr.Attributes, gpu attr.Attributes) {
	t.Helper()

	hasPlacementInterconnect := false
	for _, a := range reqs {
		if a.Key == "capabilities/gpu-interconnect" && a.Value == "true" {
			hasPlacementInterconnect = true
		}
	}
	require.True(t, hasPlacementInterconnect, "%s: Requirements lost capabilities/gpu-interconnect=true", where)

	hasGPUInterconnect := false
	for _, a := range gpu {
		if a.Key == "interconnect" && a.Value == "true" {
			hasGPUInterconnect = true
		}
	}
	require.True(t, hasGPUInterconnect, "%s: per-resource GPU.Attributes lost interconnect=true", where)
}

// TestCS6_InterconnectSignalsSurviveDup is the canonical regression test: drive the
// representative GroupSpec through Dup() and assert both attribute slices
// survive verbatim. If a future change to Resources.Dup() / Requirements.Dup()
// drops attributes, this fails loudly.
func TestCS6_InterconnectSignalsSurviveDup(t *testing.T) {
	src := interconnectSampleGroupSpec()
	dup := src.Dup()

	require.NotSame(t, &src, &dup)
	assertInterconnectSignalsPresent(t, "GroupSpec.Dup result",
		dup.Requirements.Attributes,
		dup.Resources[0].Resources.GPU.Attributes,
	)

	// And mutating the dup must not poison the source.
	dup.Requirements.Attributes[0].Value = "false"
	dup.Resources[0].Resources.GPU.Attributes[0].Value = "false"
	assertInterconnectSignalsPresent(t, "source after mutating dup",
		src.Requirements.Attributes,
		src.Resources[0].Resources.GPU.Attributes,
	)
}

// TestCS6_InterconnectSignalsAcrossConcreteTypes exercises the four concrete
// ResourceGroup-shaped values the provider's reservation/commit path can
// hold. Each row asserts that the interconnect signals are reachable via the
// type-specific accessors the provider uses.
func TestCS6_InterconnectSignalsAcrossConcreteTypes(t *testing.T) {
	specVal := interconnectSampleGroupSpec()
	specPtr := &specVal

	groupVal := Group{
		GroupSpec: specVal,
	}
	groupPtr := &groupVal

	tests := []struct {
		name string
		// readers translate the type-specific access pattern the provider
		// uses for that input shape into the two attribute slices we care
		// about. The provider's helpers do exactly this.
		readReqs func() attr.Attributes
		readGPU  func() attr.Attributes
	}{
		{
			name:     "*GroupSpec",
			readReqs: func() attr.Attributes { return specPtr.Requirements.Attributes },
			readGPU:  func() attr.Attributes { return specPtr.Resources[0].Resources.GPU.Attributes },
		},
		{
			name:     "GroupSpec value",
			readReqs: func() attr.Attributes { return specVal.Requirements.Attributes },
			readGPU:  func() attr.Attributes { return specVal.Resources[0].Resources.GPU.Attributes },
		},
		{
			name:     "*Group",
			readReqs: func() attr.Attributes { return groupPtr.GroupSpec.Requirements.Attributes },
			readGPU:  func() attr.Attributes { return groupPtr.GroupSpec.Resources[0].Resources.GPU.Attributes },
		},
		{
			name:     "Group value",
			readReqs: func() attr.Attributes { return groupVal.GroupSpec.Requirements.Attributes },
			readGPU:  func() attr.Attributes { return groupVal.GroupSpec.Resources[0].Resources.GPU.Attributes },
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assertInterconnectSignalsPresent(t, tc.name, tc.readReqs(), tc.readGPU())
		})
	}
}

// TestCS6_GetResourceUnitsPreservesGPUAttributes exercises the
// ResourceGroup-interface path the provider's commit code uses. The
// provider iterates GetResourceUnits() and reads each unit's GPU.Attributes;
// a regression where this collection silently drops attributes would
// translate to interconnect opt-in disappearing on the wire.
func TestCS6_GetResourceUnitsPreservesGPUAttributes(t *testing.T) {
	spec := interconnectSampleGroupSpec()

	units := spec.GetResourceUnits()
	require.NotEmpty(t, units, "GetResourceUnits returned empty")

	hasGPUInterconnect := false
	for _, u := range units {
		if u.Resources.GPU == nil {
			continue
		}
		for _, a := range u.Resources.GPU.Attributes {
			if a.Key == "interconnect" && a.Value == "true" {
				hasGPUInterconnect = true
			}
		}
	}
	require.True(t, hasGPUInterconnect,
		"GetResourceUnits() dropped GPU.Attributes — provider would treat as non-interconnect")
}
