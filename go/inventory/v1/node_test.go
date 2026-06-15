package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeCapabilities_Dup_PreservesInterconnectFields(t *testing.T) {
	src := NodeCapabilities{
		StorageClasses:           []string{"beta3", "default"},
		InterconnectResourceName: "rdma/rdma_shared_device_ib",
		InterconnectFabric:       "infiniband",
		NCCLHCAPrefix:            "mlx5",
	}

	got := src.Dup()

	require.Equal(t, src.StorageClasses, got.StorageClasses)
	require.Equal(t, "rdma/rdma_shared_device_ib", got.InterconnectResourceName)
	require.Equal(t, "infiniband", got.InterconnectFabric)
	require.Equal(t, "mlx5", got.NCCLHCAPrefix)

	// mutating dup must not affect the source (Dup is a deep copy)
	got.StorageClasses[0] = "mutated"
	got.InterconnectResourceName = "rdma/rdma_shared_device_eth"
	require.Equal(t, "beta3", src.StorageClasses[0])
	require.Equal(t, "rdma/rdma_shared_device_ib", src.InterconnectResourceName)
}

func TestNodeCapabilities_Dup_ZeroValueInterconnect(t *testing.T) {
	// A non-interconnect node leaves the new fields at zero value.
	src := NodeCapabilities{
		StorageClasses: []string{"default"},
	}

	got := src.Dup()

	require.Empty(t, got.InterconnectResourceName)
	require.Empty(t, got.InterconnectFabric)
	require.Empty(t, got.NCCLHCAPrefix)
}
