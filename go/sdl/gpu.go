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

// gpuAttributeRDMAGroupSentinel is an internal-only key the SDL parser uses to
// transport the value of gpu.attributes.rdma_group between
// v2GPUAttributes.UnmarshalYAML and the parent v2ResourceGPU.UnmarshalYAML.
// It is stripped out of the GPU attribute slice before those attributes ever
// reach the on-chain Resources.GPU.attributes — rdma_group is a tenant
// scheduling directive carried in the off-chain manifest, not a hardware
// capability claim.
const gpuAttributeRDMAGroupSentinel = "__rdma_group__"

type v2ResourceGPU struct {
	Units      gpuQuantity     `yaml:"units" json:"units"`
	Attributes v2GPUAttributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`

	// RDMAGroup carries the parsed gpu.attributes.rdma_group value. The SDL
	// parser strips this key from on-chain GPU attributes (see
	// gpuAttributeRDMAGroupSentinel) and lifts it here so the higher-level
	// manifest builder can route it to the per-service Service.RDMAGroup
	// off-chain manifest field.
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

	// Extract the rdma_group sentinel into the dedicated field, then strip
	// it out of the on-chain attribute slice. After this step Attributes
	// contains only attributes that are safe to flow to Resources.GPU.attributes.
	// We must do this *before* Validate() runs against the slice, because the
	// sentinel key (with leading underscores) does not match the on-chain
	// attribute key regex.
	if len(res.Attributes) > 0 {
		filtered := make(types.Attributes, 0, len(res.Attributes))
		for _, a := range res.Attributes {
			if a.Key == gpuAttributeRDMAGroupSentinel {
				res.RDMAGroup = a.Value
				continue
			}
			filtered = append(filtered, a)
		}
		res.Attributes = v2GPUAttributes(filtered)

		// Validate now that the on-chain-bound slice is clean of any sentinel.
		// (v2GPUAttributes.UnmarshalYAML intentionally skips Validate so this
		// hook can run post-strip.)
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

	// Carry rdma_group as a sentinel; the parent v2ResourceGPU.UnmarshalYAML
	// peels it off before these attributes ever become on-chain Resources.
	if rdmaGroup != "" {
		res = append(res, types.Attribute{
			Key:   gpuAttributeRDMAGroupSentinel,
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
