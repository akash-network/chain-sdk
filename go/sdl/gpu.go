package sdl

import (
	"fmt"
	"sort"

	"gopkg.in/yaml.v3"

	types "pkg.akt.dev/go/node/types/attributes/v1"
)

type gpuInterface string

type v2GPUNvidia struct {
	Model     string          `yaml:"model"`
	RAM       *memoryQuantity `yaml:"ram,omitempty"`
	Interface *gpuInterface   `yaml:"interface,omitempty"`
}

func (sdl *v2GPUNvidia) String() string {
	key := sdl.Model
	if sdl.RAM != nil {
		key = fmt.Sprintf("%s/ram/%s", key, sdl.RAM.StringWithSuffix("Gi"))
	}

	if sdl.Interface != nil {
		key = fmt.Sprintf("%s/interface/%s", key, *sdl.Interface)
	}

	return key
}

type v2GPUsNvidia []v2GPUNvidia

type gpuVendor struct {
	Nvidia v2GPUsNvidia `yaml:"nvidia,omitempty"`
}

type v2GPUAttributes types.Attributes

// GPUAttributeInterconnectGroup is the on-chain GPU-attribute key emitted
// for every interconnect-enabled resource. Its presence (regardless of
// value) signals "this resource wants GPU interconnect"; its value is the
// peer-group label the provider's bid engine uses to enforce per-group
// node separation. Path-separated key matches the existing
// `capabilities/gpu-interconnect/fabric/...` convention.
//
// Replaces the rc4 pair (`interconnect=true` + `interconnect_group=<name>`)
// with a single key. See docs/sdl-interconnect-spec.md.
const GPUAttributeInterconnectGroup = "interconnect/group"

// InterconnectGroupAuto is the reserved group name the SDL parser assigns
// to every `interconnect: []` resource within one placement. Tenants
// cannot write this name explicitly under `interconnect: { group: ... }`.
const InterconnectGroupAuto = "auto"

type v2ResourceGPU struct {
	Units      gpuQuantity     `yaml:"units" json:"units"`
	Attributes v2GPUAttributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`

	// InterconnectGroup carries the parsed group name from
	// gpu.attributes.interconnect (the implicit `[]` form resolves to
	// the literal "auto", the explicit `{group: <name>}` form carries
	// the tenant-chosen name). The same value is also present in
	// Attributes under the GPUAttributeInterconnectGroup key; this
	// field exists so the higher-level manifest builder can route it to
	// Service.InterconnectGroup without re-walking the slice.
	InterconnectGroup string `yaml:"-" json:"-"`
}

func (sdl *v2ResourceGPU) UnmarshalYAML(node *yaml.Node) error {
	res := v2ResourceGPU{}

	for i := 0; i < len(node.Content); i += 2 {
		switch node.Content[i].Value {
		case "units":
			if err := node.Content[i+1].Decode(&res.Units); err != nil {
				return err
			}
		case "attributes":
			if err := node.Content[i+1].Decode(&res.Attributes); err != nil {
				return err
			}
		default:
			return fmt.Errorf("sdl: unsupported field (%s) for GPU resource", node.Content[i].Value)
		}
	}

	// Lift the interconnect/group attribute value into the dedicated
	// InterconnectGroup field for downstream manifest builders, but KEEP
	// it in the attributes slice — the provider's bid engine consumes the
	// on-chain Resources.GPU.Attributes and needs the group key present
	// there to enforce per-group node separation during reservation.
	if len(res.Attributes) > 0 {
		for _, a := range res.Attributes {
			if a.Key == GPUAttributeInterconnectGroup {
				res.InterconnectGroup = a.Value
				break
			}
		}

		// v2GPUAttributes.UnmarshalYAML defers Validate to here so the
		// final attribute slice (including the interconnect/group key,
		// which matches the on-chain attribute key regex) gets one
		// validate pass.
		final := types.Attributes(res.Attributes)
		if err := final.Validate(); err != nil {
			return fmt.Errorf("sdl: invalid GPU attributes: %w", err)
		}
	}

	if res.Units > 0 && len(res.Attributes) == 0 {
		return fmt.Errorf("sdl: GPU attributes must be present if units > 0")
	}

	// CS-5 invariant, enforced here so the SDL fails fast: any
	// `interconnect:` opt-in is nonsense without an actual GPU to attach
	// an HCA to. A profile opting in with gpu.units == 0 would otherwise
	// be classified as interconnect-enabled by downstream validation passes
	// and the provider's reservation logic, then rejected much later (or,
	// worse, treated as a misconfiguration). Reject up front. Since the
	// group key is now the sole opt-in signal, checking InterconnectGroup
	// alone covers both implicit and explicit forms.
	if res.Units == 0 && res.InterconnectGroup != "" {
		return fmt.Errorf("sdl: gpu.attributes.interconnect cannot be set when gpu.units == 0 (group=%q)", res.InterconnectGroup)
	}

	*sdl = res

	return nil
}

func (sdl *v2GPUAttributes) UnmarshalYAML(node *yaml.Node) error {
	var res types.Attributes

	var vendor *gpuVendor
	interconnectGroup := ""

	for i := 0; i < len(node.Content); i += 2 {
		switch node.Content[i].Value {
		case "vendor":
			if err := node.Content[i+1].Decode(&vendor); err != nil {
				return err
			}
		case "interconnect":
			// gpu.attributes.interconnect accepts two forms:
			//   - empty sequence `[]`            → implicit group, named `auto`
			//   - mapping `{group: <name>}`      → explicit named group
			// Bare boolean and other shapes are rejected. The group
			// string (whether the literal "auto" sentinel or the
			// tenant-chosen name) becomes the on-chain
			// `interconnect/group` attribute value, which is the sole
			// opt-in signal — no separate "interconnect=true" marker.
			val := node.Content[i+1]
			switch val.Kind {
			case yaml.SequenceNode:
				if len(val.Content) != 0 {
					return fmt.Errorf("sdl: gpu.attributes.interconnect: only the empty sequence form `[]` is accepted; use `{group: <name>}` for explicit groups")
				}
				interconnectGroup = InterconnectGroupAuto
			case yaml.MappingNode:
				var m struct {
					Group string `yaml:"group"`
				}
				if err := val.Decode(&m); err != nil {
					return fmt.Errorf("sdl: invalid value for gpu.attributes.interconnect: %w", err)
				}
				if m.Group == "" {
					return fmt.Errorf("sdl: gpu.attributes.interconnect: `group` is required when using the mapping form")
				}
				if m.Group == InterconnectGroupAuto {
					return fmt.Errorf("sdl: gpu.attributes.interconnect.group: %q is a reserved name (parser auto-assigns it to `interconnect: []` resources); pick a different name", InterconnectGroupAuto)
				}
				interconnectGroup = m.Group
			default:
				return fmt.Errorf("sdl: gpu.attributes.interconnect: expected `[]` or `{group: <name>}`; bare scalar (including `true`/`false`) is no longer accepted")
			}
		default:
			return fmt.Errorf("sdl: unsupported attribute (%s) for GPU resource", node.Content[i].Value)
		}
	}

	if vendor == nil {
		return fmt.Errorf("sdl: invalid GPU attributes. at least one vendor must be set")
	}

	res = make(types.Attributes, 0, len(vendor.Nvidia)+2)

	for _, model := range vendor.Nvidia {
		res = append(res, types.Attribute{
			Key:   fmt.Sprintf("vendor/nvidia/model/%s", model.String()),
			Value: "true",
		})
	}

	if len(res) == 0 {
		res = append(res, types.Attribute{
			Key:   "vendor/nvidia/model/*",
			Value: "true",
		})
	}

	// Emit `interconnect/group = <name>` as the sole opt-in signal — no
	// separate boolean marker. The provider's bid engine reads this key
	// during reservation Adjust to enforce per-group node separation; the
	// parent v2ResourceGPU.UnmarshalYAML lifts the value into
	// v2ResourceGPU.InterconnectGroup so the manifest builder can route it
	// to Service.InterconnectGroup for the off-chain workload builder.
	if interconnectGroup != "" {
		res = append(res, types.Attribute{
			Key:   GPUAttributeInterconnectGroup,
			Value: interconnectGroup,
		})
	}

	sort.Sort(res)

	// Validate() is deferred to v2ResourceGPU.UnmarshalYAML so the parent
	// hook can lift interconnect_group into the dedicated field before
	// the final attribute-key regex runs across the slice.

	*sdl = v2GPUAttributes(res)

	return nil
}

func (sdl *gpuInterface) UnmarshalYAML(node *yaml.Node) error {
	switch node.Value {
	case "pcie":
	case "sxm":
	default:
		return fmt.Errorf("sdl: invalid GPU interface %s. expected \"pcie|sxm\"", node.Value)
	}

	*sdl = gpuInterface(node.Value)

	return nil
}
