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
// inventory.Adjust path. The provider's RDMA-detection logic reads
//
//   - GroupSpec.Requirements.Attributes  (carrying capabilities/rdma=true)
//   - Resources.GPU.Attributes           (carrying rdma=true per resource)
//
// from whatever ResourceGroup-typed value lands on its `reservation.Resources()`
// call. If anything in the chain SDK's Dup / accessor chain drops either of
// those attribute slices, the provider silently treats the order as
// non-RDMA. These tests pin down preservation for every concrete
// ResourceGroup the provider may hold:
//
//   - *Group         (the bid-engine path stores a *dtypes.Group)
//   - Group          (value form, dereferenced)
//   - *GroupSpec     (used by tests and by some legacy adapter paths)
//   - GroupSpec      (value form returned by Group.GroupSpec)
//
// We assert that a representative Group carrying capabilities/rdma=true on
// Requirements and rdma=true on the first resource's GPU.Attributes survives:
//
//   1. (Group).Dup()-equivalent — via GroupSpec.Dup().
//   2. Accessing the value via the ResourceGroup interface (GetResourceUnits).
//   3. Direct field reads of Requirements.Attributes from any of the four
//      concrete shapes.

func rdmaSampleGroupSpec() GroupSpec {
	return GroupSpec{
		Name: "ib",
		Requirements: attr.PlacementRequirements{
			Attributes: attr.Attributes{
				{Key: "capabilities/rdma", Value: "true"},
				{Key: "capabilities/rdma/fabric/infiniband", Value: "true"},
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
							{Key: "rdma", Value: "true"},
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

func assertRDMASignalsPresent(t *testing.T, where string, reqs attr.Attributes, gpu attr.Attributes) {
	t.Helper()

	hasPlacementRDMA := false
	for _, a := range reqs {
		if a.Key == "capabilities/rdma" && a.Value == "true" {
			hasPlacementRDMA = true
		}
	}
	require.True(t, hasPlacementRDMA, "%s: Requirements lost capabilities/rdma=true", where)

	hasGPURDMA := false
	for _, a := range gpu {
		if a.Key == "rdma" && a.Value == "true" {
			hasGPURDMA = true
		}
	}
	require.True(t, hasGPURDMA, "%s: per-resource GPU.Attributes lost rdma=true", where)
}

// TestCS6_RDMASignalsSurviveDup is the canonical regression test: drive the
// representative GroupSpec through Dup() and assert both attribute slices
// survive verbatim. If a future change to Resources.Dup() / Requirements.Dup()
// drops attributes, this fails loudly.
func TestCS6_RDMASignalsSurviveDup(t *testing.T) {
	src := rdmaSampleGroupSpec()
	dup := src.Dup()

	require.NotSame(t, &src, &dup)
	assertRDMASignalsPresent(t, "GroupSpec.Dup result",
		dup.Requirements.Attributes,
		dup.Resources[0].Resources.GPU.Attributes,
	)

	// And mutating the dup must not poison the source.
	dup.Requirements.Attributes[0].Value = "false"
	dup.Resources[0].Resources.GPU.Attributes[0].Value = "false"
	assertRDMASignalsPresent(t, "source after mutating dup",
		src.Requirements.Attributes,
		src.Resources[0].Resources.GPU.Attributes,
	)
}

// TestCS6_RDMASignalsAcrossConcreteTypes exercises the four concrete
// ResourceGroup-shaped values the provider's reservation/commit path can
// hold. Each row asserts that the RDMA signals are reachable via the
// type-specific accessors the provider uses.
func TestCS6_RDMASignalsAcrossConcreteTypes(t *testing.T) {
	specVal := rdmaSampleGroupSpec()
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
			assertRDMASignalsPresent(t, tc.name, tc.readReqs(), tc.readGPU())
		})
	}
}

// TestCS6_GetResourceUnitsPreservesGPUAttributes exercises the
// ResourceGroup-interface path the provider's commit code uses. The
// provider iterates GetResourceUnits() and reads each unit's GPU.Attributes;
// a regression where this collection silently drops attributes would
// translate to RDMA opt-in disappearing on the wire.
func TestCS6_GetResourceUnitsPreservesGPUAttributes(t *testing.T) {
	spec := rdmaSampleGroupSpec()

	units := spec.GetResourceUnits()
	require.NotEmpty(t, units, "GetResourceUnits returned empty")

	hasGPURDMA := false
	for _, u := range units {
		if u.Resources.GPU == nil {
			continue
		}
		for _, a := range u.Resources.GPU.Attributes {
			if a.Key == "rdma" && a.Value == "true" {
				hasGPURDMA = true
			}
		}
	}
	require.True(t, hasGPURDMA,
		"GetResourceUnits() dropped GPU.Attributes — provider would treat as non-RDMA")
}
