package sdl

import (
	"fmt"
	"sort"

	"gopkg.in/yaml.v3"

	types "pkg.akt.dev/go/node/types/attributes/v1"
)

type v2CPUAttributes types.Attributes

type v2ResourceCPU struct {
	Units      cpuQuantity     `yaml:"units"`
	Attributes v2CPUAttributes `yaml:"attributes,omitempty"`
}

func (sdl *v2ResourceCPU) UnmarshalYAML(node *yaml.Node) error {
	res := v2ResourceCPU{}
	var arch string

	for i := 0; i < len(node.Content); i += 2 {
		switch node.Content[i].Value {
		case "units":
			if err := node.Content[i+1].Decode(&res.Units); err != nil {
				return err
			}
		case "arch":
			if err := node.Content[i+1].Decode(&arch); err != nil {
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

	if arch != "" {
		switch arch {
		case "amd64", "arm64":
		default:
			return fmt.Errorf("sdl: unsupported CPU arch (%s), expected amd64|arm64", arch)
		}
		for _, attr := range res.Attributes {
			if attr.Key == "arch" {
				return fmt.Errorf("sdl: cpu arch specified at both top-level and in attributes")
			}
		}
		res.Attributes = append(res.Attributes, types.Attribute{Key: "arch", Value: arch})
		sort.Sort(types.Attributes(res.Attributes))
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
