package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"

	inventoryV1 "pkg.akt.dev/go/inventory/v1"
)

func TestInventoryReportsLeasedIPResourcePair(t *testing.T) {
	inventory := Inventory{
		LeasedIP: inventoryV1.NewResourcePair(10, 10, 6, resource.DecimalSI),
	}

	status := inventory.GetLeasedIP()
	require.Equal(t, int64(10), status.GetCapacity().Value())
	require.Equal(t, int64(10), status.GetAllocatable().Value())
	require.Equal(t, int64(6), status.GetAllocated().Value())
	require.Equal(t, int64(4), status.Available().Value())

	data, err := inventory.Marshal()
	require.NoError(t, err)

	var decoded Inventory
	require.NoError(t, decoded.Unmarshal(data))

	status = decoded.GetLeasedIP()
	require.Equal(t, int64(10), status.GetCapacity().Value())
	require.Equal(t, int64(10), status.GetAllocatable().Value())
	require.Equal(t, int64(6), status.GetAllocated().Value())
	require.Equal(t, int64(4), status.Available().Value())
}
