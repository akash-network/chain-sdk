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

// GPUAttributeInterconnect is the on-chain GPU-attribute key emitted when
// an SDL compute profile declares gpu.attributes.interconnect: true.
// Providers advertising GPU-interconnect-capable hardware match this
// attribute via the standard GPU MatchResourcesRequirements path. The
// fabric (InfiniBand vs RoCE) is hidden from the SDL surface; the
// provider picks whichever it has.
const GPUAttributeInterconnect = "interconnect"

// GPUAttributeInterconnectGroup is the on-chain GPU-attribute key emitted
// when an SDL compute profile declares gpu.attributes.interconnect_group:
// <name>. Carries the peer-group label all the way into the on-chain
// Resources.GPU.attributes so the provider's bid engine can enforce
// per-group node separation at fit time (Service.InterconnectGroup
// off-chain alone would leave the bid step blind). The value is also
// lifted into Service.InterconnectGroup so the workload builder's pod
// anti-affinity rule keys off the same string.
const GPUAttributeInterconnectGroup = "interconnect_group"

type v2ResourceGPU struct {
	Units      gpuQuantity     `yaml:"units" json:"units"`
	Attributes v2GPUAttributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`

	// InterconnectGroup carries the parsed gpu.attributes.interconnect_group
	// value. The same value is also present in Attributes (as
	// GPUAttributeInterconnectGroup); this field exists so the
	// higher-level manifest builder can route it to
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

	// Lift the interconnect_group attribute value into the dedicated
	// InterconnectGroup field for downstream manifest builders, but KEEP
	// it in the attributes slice — the provider's bid engine consumes the
	// on-chain Resources.GPU.Attributes and needs interconnect_group
	// present there to enforce per-group node separation during
	// reservation.
	if len(res.Attributes) > 0 {
		for _, a := range res.Attributes {
			if a.Key == GPUAttributeInterconnectGroup {
				res.InterconnectGroup = a.Value
				break
			}
		}

		// v2GPUAttributes.UnmarshalYAML defers Validate to here so the
		// final attribute slice (including the interconnect_group key,
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

	// CS-5 invariant, enforced here so the SDL fails fast: interconnect
	// and interconnect_group are nonsense without an actual GPU to attach
	// an HCA to. A profile declaring interconnect: true or
	// interconnect_group: <name> with gpu.units == 0 would otherwise be
	// classified as interconnect-enabled by downstream validation passes
	// and the provider's reservation logic, then rejected much later (or,
	// worse, treated as a misconfiguration). Reject up front.
	if res.Units == 0 {
		if gpuAttributesHaveInterconnect(res.Attributes) {
			return fmt.Errorf("sdl: gpu.attributes.interconnect cannot be set when gpu.units == 0")
		}
		if res.InterconnectGroup != "" {
			return fmt.Errorf("sdl: gpu.attributes.interconnect_group=%q cannot be set when gpu.units == 0", res.InterconnectGroup)
		}
	}

	*sdl = res

	return nil
}

func (sdl *v2GPUAttributes) UnmarshalYAML(node *yaml.Node) error {
	var res types.Attributes

	var vendor *gpuVendor
	interconnectEnabled := false
	interconnectGroup := ""

	for i := 0; i < len(node.Content); i += 2 {
		switch node.Content[i].Value {
		case "vendor":
			if err := node.Content[i+1].Decode(&vendor); err != nil {
				return err
			}
		case "interconnect":
			// gpu.attributes.interconnect: bool (default false). When
			// true, emit an on-chain GPU attribute so providers
			// advertising GPU-interconnect-capable hardware can be
			// matched. Fabric (IB vs RoCE) is hidden from the SDL —
			// provider picks whichever it has.
			var ic bool
			if err := node.Content[i+1].Decode(&ic); err != nil {
				return fmt.Errorf("sdl: invalid value for gpu.attributes.interconnect: %w", err)
			}
			interconnectEnabled = ic
		case "interconnect_group":
			// gpu.attributes.interconnect_group: string (peer group
			// name). Emitted as an on-chain GPU attribute alongside
			// `interconnect=true` so the provider's bid engine can
			// track per-group node claims. Also lifted to
			// v2ResourceGPU.InterconnectGroup for the manifest builder.
			if err := node.Content[i+1].Decode(&interconnectGroup); err != nil {
				return fmt.Errorf("sdl: invalid value for gpu.attributes.interconnect_group: %w", err)
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

	if interconnectEnabled {
		res = append(res, types.Attribute{
			Key:   GPUAttributeInterconnect,
			Value: "true",
		})
	}

	// Emit interconnect_group directly as an on-chain GPU attribute. The
	// provider's reservation Adjust step reads this to enforce per-group
	// node separation; the parent v2ResourceGPU.UnmarshalYAML also lifts
	// the value into v2ResourceGPU.InterconnectGroup so the manifest
	// builder can route it to Service.InterconnectGroup for the off-chain
	// workload builder.
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
