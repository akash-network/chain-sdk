package sdl

import (
	"sort"

	"gopkg.in/yaml.v3"

	types "pkg.akt.dev/go/node/types/attributes/v1"
)

type v2PlacementAttributes types.Attributes

func (sdl *v2PlacementAttributes) UnmarshalYAML(node *yaml.Node) error {
	var attr v2PlacementAttributes

	var res map[string]string

	if err := node.Decode(&res); err != nil {
		return err
	}

	for k, v := range res {
		attr = append(attr, types.Attribute{
			Key:   k,
			Value: v,
		})
	}

	// keys are unique in attributes parsed from sdl so don't need to use sort.SliceStable
	sort.Slice(attr, func(i, j int) bool {
		return attr[i].Key < attr[j].Key
	})

	*sdl = attr

	return nil
}
