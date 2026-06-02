package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeCapabilities_Dup_PreservesRDMAFields(t *testing.T) {
	src := NodeCapabilities{
		StorageClasses:   []string{"beta3", "default"},
		RDMAResourceName: "rdma/rdma_shared_device_ib",
		RDMAFabric:       "infiniband",
		NCCLHCAPrefix:    "mlx5",
	}

	got := src.Dup()

	require.Equal(t, src.StorageClasses, got.StorageClasses)
	require.Equal(t, "rdma/rdma_shared_device_ib", got.RDMAResourceName)
	require.Equal(t, "infiniband", got.RDMAFabric)
	require.Equal(t, "mlx5", got.NCCLHCAPrefix)

	// mutating dup must not affect the source (Dup is a deep copy)
	got.StorageClasses[0] = "mutated"
	got.RDMAResourceName = "rdma/rdma_shared_device_eth"
	require.Equal(t, "beta3", src.StorageClasses[0])
	require.Equal(t, "rdma/rdma_shared_device_ib", src.RDMAResourceName)
}

func TestNodeCapabilities_Dup_ZeroValueRDMA(t *testing.T) {
	// A non-RDMA node leaves the new fields at zero value.
	src := NodeCapabilities{
		StorageClasses: []string{"default"},
	}

	got := src.Dup()

	require.Empty(t, got.RDMAResourceName)
	require.Empty(t, got.RDMAFabric)
	require.Empty(t, got.NCCLHCAPrefix)
}
