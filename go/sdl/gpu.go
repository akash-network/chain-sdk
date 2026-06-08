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

// GPUAttributeRDMA is the on-chain GPU-attribute key emitted when an SDL
// compute profile declares gpu.attributes.rdma: true. Providers advertising
// RDMA-capable GPU hardware match this attribute via the standard GPU
// MatchResourcesRequirements path.
const GPUAttributeRDMA = "rdma"

// GPUAttributeRDMAGroup is the on-chain GPU-attribute key emitted when an
// SDL compute profile declares gpu.attributes.rdma_group: <name>. Carries
// the peer-group label all the way into the on-chain Resources.GPU.attributes
// so the provider's bid engine can enforce per-group node separation at fit
// time (it cannot otherwise — Service.RDMAGroup is off-chain only, and the
// bid engine consumes Resources, not the manifest). The value is also lifted
// into Service.RDMAGroup so the workload builder's pod anti-affinity rule
// continues to work the same way.
const GPUAttributeRDMAGroup = "rdma_group"

type v2ResourceGPU struct {
	Units      gpuQuantity     `yaml:"units" json:"units"`
	Attributes v2GPUAttributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`

	// RDMAGroup carries the parsed gpu.attributes.rdma_group value. The
	// same value is also present in Attributes (as GPUAttributeRDMAGroup);
	// this field exists so the higher-level manifest builder can route it
	// to Service.RDMAGroup without re-walking the slice.
	RDMAGroup string `yaml:"-" json:"-"`
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

	// Lift the rdma_group attribute value into the dedicated RDMAGroup
	// field for downstream manifest builders, but KEEP it in the
	// attributes slice — the provider's bid engine consumes the on-chain
	// Resources.GPU.Attributes and needs rdma_group present there to
	// enforce per-group node separation during reservation. (Was a
	// sentinel-stripped off-chain-only field; now flows end-to-end.)
	if len(res.Attributes) > 0 {
		for _, a := range res.Attributes {
			if a.Key == GPUAttributeRDMAGroup {
				res.RDMAGroup = a.Value
				break
			}
		}

		// v2GPUAttributes.UnmarshalYAML defers Validate to here so the
		// final attribute slice (including the rdma_group key, which now
		// matches the on-chain attribute key regex) gets one validate pass.
		final := types.Attributes(res.Attributes)
		if err := final.Validate(); err != nil {
			return fmt.Errorf("sdl: invalid GPU attributes: %w", err)
		}
	}

	if res.Units > 0 && len(res.Attributes) == 0 {
		return fmt.Errorf("sdl: GPU attributes must be present if units > 0")
	}

	// CS-5 invariant, enforced here so the SDL fails fast: rdma / rdma_group
	// are nonsense without an actual GPU to attach an HCA to. A profile
	// declaring rdma: true or rdma_group: <name> with gpu.units == 0 would
	// otherwise be classified as RDMA-enabled by downstream validation
	// passes and the provider's reservation logic, then rejected much later
	// (or, worse, treated as a misconfiguration). Reject up front.
	if res.Units == 0 {
		if gpuAttributesHaveRDMA(res.Attributes) {
			return fmt.Errorf("sdl: gpu.attributes.rdma cannot be set when gpu.units == 0")
		}
		if res.RDMAGroup != "" {
			return fmt.Errorf("sdl: gpu.attributes.rdma_group=%q cannot be set when gpu.units == 0", res.RDMAGroup)
		}
	}

	*sdl = res

	return nil
}

func (sdl *v2GPUAttributes) UnmarshalYAML(node *yaml.Node) error {
	var res types.Attributes

	var vendor *gpuVendor
	rdmaEnabled := false
	rdmaGroup := ""

	for i := 0; i < len(node.Content); i += 2 {
		switch node.Content[i].Value {
		case "vendor":
			if err := node.Content[i+1].Decode(&vendor); err != nil {
				return err
			}
		case "rdma":
			// gpu.attributes.rdma: bool (default false). When true, emit an
			// on-chain GPU attribute so providers advertising RDMA-capable
			// GPU hardware can be matched.
			var rdma bool
			if err := node.Content[i+1].Decode(&rdma); err != nil {
				return fmt.Errorf("sdl: invalid value for gpu.attributes.rdma: %w", err)
			}
			rdmaEnabled = rdma
		case "rdma_group":
			// gpu.attributes.rdma_group: string (peer group name). Captured
			// here and emitted into the slice as a sentinel attribute that
			// v2ResourceGPU.UnmarshalYAML strips before it reaches chain
			// state. See gpuAttributeRDMAGroupSentinel.
			if err := node.Content[i+1].Decode(&rdmaGroup); err != nil {
				return fmt.Errorf("sdl: invalid value for gpu.attributes.rdma_group: %w", err)
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

	if rdmaEnabled {
		res = append(res, types.Attribute{
			Key:   GPUAttributeRDMA,
			Value: "true",
		})
	}

	// Emit rdma_group directly as an on-chain GPU attribute. The provider's
	// reservation Adjust step reads this to enforce per-group node
	// separation; the parent v2ResourceGPU.UnmarshalYAML also lifts the
	// value into v2ResourceGPU.RDMAGroup so the manifest builder can
	// route it to Service.RDMAGroup for the off-chain workload builder.
	if rdmaGroup != "" {
		res = append(res, types.Attribute{
			Key:   GPUAttributeRDMAGroup,
			Value: rdmaGroup,
		})
	}

	sort.Sort(res)

	// Validate() is deferred to v2ResourceGPU.UnmarshalYAML so the
	// rdma_group sentinel can be stripped from the slice before the
	// attribute-key regex runs against it.

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
