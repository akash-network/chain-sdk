package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"
)

// ResourcePair.Dup() previously dereferenced Capacity/Allocatable/Allocated
// unconditionally, which panics for a zero-value ResourcePair.
// Non-interconnect nodes legitimately leave NodeResources.GPUInterconnect
// at zero value (no quantity pointers populated), so NodeResources.Dup()
// ends up calling ResourcePair.Dup() on a zero value.
//
// Two assertions:
//  1. ResourcePair.Dup() returns an equivalent zero-valued ResourcePair
//     without panicking when called on the nil/zero receiver.
//  2. NodeResources.Dup() round-trips a non-interconnect NodeResources
//     (GPUInterconnect left at zero value) without panicking.

func TestResourcePair_Dup_ZeroValueDoesNotPanic(t *testing.T) {
	var zero ResourcePair
	require.True(t, zero.IsZero(), "fresh ResourcePair must report IsZero")

	got := zero.Dup()
	require.True(t, got.IsZero(), "Dup of zero ResourcePair must remain zero")
}

func TestResourcePair_Dup_NilReceiverDoesNotPanic(t *testing.T) {
	var nilRP *ResourcePair
	// Calling a method on a nil pointer receiver is only safe if the method
	// itself checks for nil; the new Dup() does.
	got := nilRP.Dup()
	require.True(t, got.IsZero())
}

func TestResourcePair_Dup_PopulatedRoundTrips(t *testing.T) {
	rp := NewResourcePair(63, 8, 8, resource.DecimalSI)
	got := rp.Dup()

	require.Equal(t, int64(63), got.Capacity.Value())
	require.Equal(t, int64(8), got.Allocatable.Value())
	require.Equal(t, int64(8), got.Allocated.Value())

	// Mutating the dup must not poison the source — Dup() is a deep copy.
	got.Allocated.Set(0)
	require.Equal(t, int64(8), rp.Allocated.Value())
}

// Regression-pins CodeRabbit review #3: previous Dup() always returned
// non-nil pointers for Capacity/Allocatable/Allocated even when the
// source pointers were nil. That changed protobuf field-presence on the
// copy and shifted the JSON serialization (which the manifest version
// hash is computed from). A partially-nil source must Dup to a
// structurally identical copy.
func TestResourcePair_Dup_PreservesNilPointers(t *testing.T) {
	q := resource.MustParse("8")

	// Capacity nil; Allocatable + Allocated populated.
	src := ResourcePair{
		Allocatable: &q,
		Allocated:   resource.NewQuantity(0, resource.DecimalSI),
	}
	require.False(t, src.IsZero(), "src has non-nil pointers — not zero")

	got := src.Dup()
	require.Nil(t, got.Capacity, "Dup must preserve nil Capacity")
	require.NotNil(t, got.Allocatable, "non-nil Allocatable must round-trip")
	require.NotNil(t, got.Allocated, "non-nil Allocated must round-trip")
	require.Equal(t, int64(8), got.Allocatable.Value())

	// Mutating the dup must not poison the source.
	got.Allocatable.Set(99)
	require.Equal(t, int64(8), src.Allocatable.Value(), "Dup must deep-copy")
}

// Mirrors the realistic non-interconnect-node case the bug would surface in.
func TestNodeResources_Dup_ZeroGPUInterconnect_DoesNotPanic(t *testing.T) {
	zero := NewResourcePair(0, 0, 0, resource.DecimalSI)

	// Every member is initialized except GPUInterconnect, which is left
	// at zero value.
	src := NodeResources{
		CPU:              CPU{Quantity: NewResourcePairMilli(0, 0, 0, resource.DecimalSI)},
		Memory:           Memory{Quantity: zero},
		GPU:              GPU{Quantity: zero},
		EphemeralStorage: zero,
		VolumesAttached:  zero,
		VolumesMounted:   zero,
		// GPUInterconnect: <intentionally zero value>
	}

	got := src.Dup()
	require.True(t, got.GPUInterconnect.IsZero(), "zero-value GPUInterconnect must Dup as zero")
}
