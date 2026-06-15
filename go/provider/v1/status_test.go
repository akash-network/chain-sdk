package v1

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protowire"
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

func TestStatusReportsReclamationWindow(t *testing.T) {
	reclamationWindow := 24 * time.Hour
	status := Status{
		ReclamationWindow: &reclamationWindow,
	}

	data, err := status.Marshal()
	require.NoError(t, err)

	var decoded Status
	require.NoError(t, decoded.Unmarshal(data))

	require.NotNil(t, decoded.GetReclamationWindow())
	require.Equal(t, reclamationWindow, *decoded.GetReclamationWindow())
}

func TestStatusOmitsUnsetReclamationWindow(t *testing.T) {
	status := Status{}

	data, err := status.Marshal()
	require.NoError(t, err)
	require.False(t, hasProtoField(t, data, 7))

	var decoded Status
	require.NoError(t, decoded.Unmarshal(data))

	require.Nil(t, decoded.GetReclamationWindow())
}

func hasProtoField(t *testing.T, data []byte, fieldNumber protowire.Number) bool {
	t.Helper()

	for len(data) > 0 {
		num, typ, n := protowire.ConsumeTag(data)
		require.GreaterOrEqual(t, n, 0)

		data = data[n:]
		if num == fieldNumber {
			return true
		}

		n = protowire.ConsumeFieldValue(num, typ, data)
		require.GreaterOrEqual(t, n, 0)
		data = data[n:]
	}

	return false
}
