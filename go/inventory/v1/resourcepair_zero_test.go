package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"
)

// CodeRabbit follow-up: ResourcePair.Dup() previously dereferenced
// Capacity/Allocatable/Allocated unconditionally, which panics for a
// zero-value ResourcePair. Non-RDMA nodes legitimately leave
// NodeResources.RDMA at zero value (no quantity pointers populated), so
// NodeResources.Dup() ends up calling ResourcePair.Dup() on a zero value.
//
// Two assertions:
//  1. ResourcePair.Dup() returns an equivalent zero-valued ResourcePair
//     without panicking when called on the nil/zero receiver.
//  2. NodeResources.Dup() round-trips a non-RDMA NodeResources (RDMA
//     left at zero value) without panicking.

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

// Mirrors the realistic non-RDMA-node case the bug would surface in.
func TestNodeResources_Dup_ZeroRDMA_DoesNotPanic(t *testing.T) {
	zero := NewResourcePair(0, 0, 0, resource.DecimalSI)

	// Every member is initialized except RDMA, which is left at zero value.
	src := NodeResources{
		CPU:              CPU{Quantity: NewResourcePairMilli(0, 0, 0, resource.DecimalSI)},
		Memory:           Memory{Quantity: zero},
		GPU:              GPU{Quantity: zero},
		EphemeralStorage: zero,
		VolumesAttached:  zero,
		VolumesMounted:   zero,
		// RDMA: <intentionally zero value>
	}

	got := src.Dup()
	require.True(t, got.RDMA.IsZero(), "zero-value RDMA must Dup as zero")
}
