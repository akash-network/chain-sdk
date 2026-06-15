package v1

import (
	"math"

	sdkmath "cosmossdk.io/math"
	"k8s.io/apimachinery/pkg/api/resource"

	types "pkg.akt.dev/go/node/types/resources/v1beta4"
)

func NewResourcePair(capacity, allocatable, allocated int64, format resource.Format) ResourcePair {
	res := ResourcePair{
		Capacity:    resource.NewQuantity(capacity, format),
		Allocatable: resource.NewQuantity(allocatable, format),
		Allocated:   resource.NewQuantity(allocated, format),
		Attributes:  nil,
	}

	return res
}

func NewResourcePairMilli(capacity, allocatable, allocated int64, format resource.Format) ResourcePair {
	res := ResourcePair{
		Capacity:    resource.NewMilliQuantity(capacity, format),
		Allocatable: resource.NewMilliQuantity(allocatable, format),
		Allocated:   resource.NewMilliQuantity(allocated, format),
		Attributes:  nil,
	}

	return res
}

func (m *ResourcePair) Equal(rhs ResourcePair) bool {
	if m == nil {
		return false
	}

	return (m.Allocatable.Cmp(*rhs.Allocatable) == 0) && (m.Allocated.Cmp(*rhs.Allocated) == 0)
}

func (m *ResourcePair) LT(rhs ResourcePair) bool {
	if m == nil {
		return false
	}

	return m.Allocatable.Cmp(*rhs.Allocatable) == -1
}

// IsZero reports whether the ResourcePair has been initialized. A
// zero-valued ResourcePair has nil quantity pointers and is the natural
// state for, e.g., a node that does not have GPU interconnect capacity
// (and therefore leaves `NodeResources.GPUInterconnect` untouched).
func (m *ResourcePair) IsZero() bool {
	if m == nil {
		return true
	}
	return m.Capacity == nil && m.Allocatable == nil && m.Allocated == nil && len(m.Attributes) == 0
}

func (m *ResourcePair) Dup() ResourcePair {
	// A zero-valued ResourcePair (all quantity pointers nil) must round-trip
	// through Dup() without panicking. Without this guard, calling Dup()
	// against e.g. an unpopulated `NodeResources.GPUInterconnect` on a
	// non-interconnect node nil-derefs Capacity.DeepCopy().
	if m == nil || m.IsZero() {
		return ResourcePair{}
	}

	// Preserve the nil/non-nil shape of each quantity pointer. Returning
	// `&zeroQuantity` for an originally-nil field would change protobuf
	// field-presence semantics (and the JSON serialization the manifest
	// version hash is computed from), so a partially-populated source
	// must Dup to a structurally identical copy.
	var capacity, allocatable, allocated *resource.Quantity
	if m.Capacity != nil {
		c := m.Capacity.DeepCopy()
		capacity = &c
	}
	if m.Allocatable != nil {
		a := m.Allocatable.DeepCopy()
		allocatable = &a
	}
	if m.Allocated != nil {
		al := m.Allocated.DeepCopy()
		allocated = &al
	}

	return ResourcePair{
		Capacity:    capacity,
		Allocatable: allocatable,
		Allocated:   allocated,
		Attributes:  m.Attributes.Dup(),
	}
}

func (m *ResourcePair) SubMilliNLZ(val types.ResourceValue) bool {
	avail := m.Available()

	res := sdkmath.NewInt(avail.MilliValue())
	res = res.Sub(val.Val)
	if res.IsNegative() {
		return false
	}

	allocated := m.Allocated.DeepCopy()
	allocated.Add(*resource.NewMilliQuantity(int64(val.Value()), resource.DecimalSI)) // nolint: gosec

	allocatable := m.Allocatable.DeepCopy()
	capacity := m.Capacity.DeepCopy()

	*m = ResourcePair{
		Capacity:    &capacity,
		Allocatable: &allocatable,
		Allocated:   &allocated,
	}

	return true
}

func (m *ResourcePair) SubNLZ(val types.ResourceValue) bool {
	avail := m.Available()

	res := sdkmath.NewInt(avail.Value())
	res = res.Sub(val.Val)

	if res.IsNegative() {
		return false
	}

	allocated := m.Allocated.DeepCopy()
	allocated.Add(*resource.NewQuantity(int64(val.Value()), resource.DecimalSI)) // nolint: gosec

	allocatable := m.Allocatable.DeepCopy()
	capacity := m.Capacity.DeepCopy()

	*m = ResourcePair{
		Capacity:    &capacity,
		Allocatable: &allocatable,
		Allocated:   &allocated,
	}

	return true
}

func (m *ResourcePair) Available() *resource.Quantity {
	result := m.Allocatable.DeepCopy()

	if result.Value() == -1 {
		result = *resource.NewQuantity(math.MaxInt64, resource.DecimalSI)
	}

	// Modifies the value in place
	(&result).Sub(*m.Allocated)

	// Clamp to zero to prevent underflow when Allocated > Allocatable
	if result.Value() < 0 {
		result = *resource.NewQuantity(0, resource.DecimalSI)
	}

	return &result
}
