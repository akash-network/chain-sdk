package sdl

import (
	"fmt"
	"sort"

	"gopkg.in/yaml.v3"

	types "pkg.akt.dev/go/node/types/attributes/v1"
)

// cpuArchitectures is the set of CPU architectures an SDL may ask for. It is
// the Go half of the `arch` enum in sdl-input.schema.yaml — the two must stay
// in step, since the schema is what the TypeScript SDK validates against and
// the two SDKs have to accept and reject the same SDLs.
//
// There is deliberately no default. An SDL that omits `arch` produces zero CPU
// attributes; writing an implicit amd64 would change the group spec of every
// existing deployment and break bid matching.
var cpuArchitectures = map[string]bool{
	"amd64": true,
	"arm64": true,
}

type v2CPUAttributes types.Attributes

type v2ResourceCPU struct {
	Units      cpuQuantity     `yaml:"units"`
	Attributes v2CPUAttributes `yaml:"attributes,omitempty"`
}

func (sdl *v2ResourceCPU) UnmarshalYAML(node *yaml.Node) error {
	res := v2ResourceCPU{}

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
			return fmt.Errorf("sdl: unsupported field (%s) for CPU resource", node.Content[i].Value)
		}
	}

	*sdl = res
	return nil
}

func (sdl *v2CPUAttributes) UnmarshalYAML(node *yaml.Node) error {
	var attr v2CPUAttributes

	for i := 0; i+1 < len(node.Content); i += 2 {
		var value string
		switch node.Content[i].Value {
		case "arch":
			if err := node.Content[i+1].Decode(&value); err != nil {
				return err
			}

			if !cpuArchitectures[value] {
				return fmt.Errorf("unsupported cpu architecture \"%s\"", value)
			}
		default:
			return fmt.Errorf("unsupported cpu attribute \"%s\"", node.Content[i].Value)
		}

		attr = append(attr, types.Attribute{
			Key:   node.Content[i].Value,
			Value: value,
		})
	}

	sort.Sort(types.Attributes(attr))

	*sdl = attr

	return nil
}
