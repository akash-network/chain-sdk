package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestNodeResources_Dup_PreservesGPUInterconnect(t *testing.T) {
	// NodeResources.Dup() reads each member's Dup() so every member must be
	// initialized to a non-zero ResourcePair to avoid nil-quantity panics.
	zeroPair := NewResourcePair(0, 0, 0, resource.DecimalSI)
	src := NodeResources{
		CPU:              CPU{Quantity: NewResourcePairMilli(0, 0, 0, resource.DecimalSI)},
		Memory:           Memory{Quantity: zeroPair},
		GPU:              GPU{Quantity: zeroPair},
		EphemeralStorage: zeroPair,
		VolumesAttached:  zeroPair,
		VolumesMounted:   zeroPair,
		GPUInterconnect:  NewResourcePair(63, 8, 8, resource.DecimalSI),
	}

	got := src.Dup()

	require.Equal(t, int64(63), got.GPUInterconnect.Capacity.Value())
	require.Equal(t, int64(8), got.GPUInterconnect.Allocatable.Value())
	require.Equal(t, int64(8), got.GPUInterconnect.Allocated.Value())

	// Source unaffected when mutating dup
	got.GPUInterconnect.Allocated.Set(0)
	require.Equal(t, int64(8), src.GPUInterconnect.Allocated.Value())
}
