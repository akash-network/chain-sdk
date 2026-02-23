package v1

import (
	"math"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"

	types "pkg.akt.dev/go/node/types/resources/v1beta4"
)

func TestAvailable_Normal(t *testing.T) {
	rp := NewResourcePair(10, 10, 3, resource.DecimalSI)
	avail := rp.Available()
	assert.Equal(t, int64(7), avail.Value())
}

func TestAvailable_ExactMatch(t *testing.T) {
	rp := NewResourcePair(5, 5, 5, resource.DecimalSI)
	avail := rp.Available()
	assert.Equal(t, int64(0), avail.Value())
}

func TestAvailable_AllocatedExceedsAllocatable(t *testing.T) {
	rp := NewResourcePair(0, 0, 8, resource.DecimalSI)
	avail := rp.Available()
	assert.Equal(t, int64(0), avail.Value(), "Available must be 0 when Allocated > Allocatable, not underflow")

	// Verify casting to uint64 doesn't produce a huge number
	val := uint64(avail.Value()) // nolint: gosec
	assert.Equal(t, uint64(0), val, "uint64 cast must also be 0")
}

func TestAvailable_PartialUnderflow(t *testing.T) {
	rp := NewResourcePair(10, 3, 10, resource.DecimalSI)
	avail := rp.Available()
	assert.Equal(t, int64(0), avail.Value(), "Available must clamp to 0 when Allocated > Allocatable")
}

func TestAvailable_Unlimited(t *testing.T) {
	rp := NewResourcePair(0, -1, 100, resource.DecimalSI)
	avail := rp.Available()
	expected := int64(math.MaxInt64) - 100
	assert.Equal(t, expected, avail.Value())
}

func TestSubNLZ_Normal(t *testing.T) {
	rp := NewResourcePair(10, 10, 3, resource.DecimalSI)
	ok := rp.SubNLZ(types.ResourceValue{Val: sdkmath.NewInt(5)})
	assert.True(t, ok)
	assert.Equal(t, int64(8), rp.Allocated.Value())
}

func TestSubNLZ_ExceedsAvailable(t *testing.T) {
	rp := NewResourcePair(10, 10, 3, resource.DecimalSI)
	ok := rp.SubNLZ(types.ResourceValue{Val: sdkmath.NewInt(8)})
	assert.False(t, ok)
	assert.Equal(t, int64(3), rp.Allocated.Value())
}

func TestSubNLZ_WhenOverAllocated(t *testing.T) {
	rp := NewResourcePair(0, 0, 8, resource.DecimalSI)
	original := rp.Dup()

	ok := rp.SubNLZ(types.ResourceValue{Val: sdkmath.NewInt(1)})
	assert.False(t, ok, "SubNLZ must return false when already over-allocated")

	require.True(t, rp.Equal(original))
}

func TestSubMilliNLZ_Normal(t *testing.T) {
	rp := NewResourcePairMilli(10000, 10000, 3000, resource.DecimalSI)
	ok := rp.SubMilliNLZ(types.ResourceValue{Val: sdkmath.NewInt(5000)})
	assert.True(t, ok)
	assert.Equal(t, int64(8000), rp.Allocated.MilliValue())
}

func TestSubMilliNLZ_WhenOverAllocated(t *testing.T) {
	rp := NewResourcePairMilli(0, 0, 8000, resource.DecimalSI)
	original := rp.Dup()

	ok := rp.SubMilliNLZ(types.ResourceValue{Val: sdkmath.NewInt(1000)})
	assert.False(t, ok, "SubMilliNLZ must return false when already over-allocated")

	require.True(t, rp.Equal(original))
}
