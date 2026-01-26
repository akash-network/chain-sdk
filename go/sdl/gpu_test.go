package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestV2ResourceGPU(t *testing.T) {
	tests := []struct {
		name          string
		yaml          string
		shouldErr     bool
		expectedUnits gpuQuantity
		expectedAttrs []struct{ key, value string }
	}{
		{
			name: "empty vendor",
			yaml: `units: 1
attributes:
  vendor:`,
			shouldErr: true,
		},
		{
			name: "wildcard",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:`,
			expectedUnits: 1,
			expectedAttrs: []struct{ key, value string }{
				{"vendor/nvidia/model/*", "true"},
			},
		},
		{
			name: "single model",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100`,
			expectedUnits: 1,
			expectedAttrs: []struct{ key, value string }{
				{"vendor/nvidia/model/a100", "true"},
			},
		},
		{
			name: "model with RAM",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
        ram: 80Gi`,
			expectedUnits: 1,
			expectedAttrs: []struct{ key, value string }{
				{"vendor/nvidia/model/a100/ram/80Gi", "true"},
			},
		},
		{
			name: "invalid RAM unit",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
        ram: 80G`,
			shouldErr: true,
		},
		{
			name: "invalid interface",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
        interface: pciex`,
			shouldErr: true,
		},
		{
			name: "RAM with interface",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
        ram: 80Gi
        interface: pcie`,
			expectedUnits: 1,
			expectedAttrs: []struct{ key, value string }{
				{"vendor/nvidia/model/a100/ram/80Gi/interface/pcie", "true"},
			},
		},
		{
			name: "multiple models same type different RAM",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
        ram: 80Gi
      - model: a100
        ram: 40Gi`,
			expectedUnits: 1,
			expectedAttrs: []struct{ key, value string }{
				{"vendor/nvidia/model/a100/ram/40Gi", "true"},
				{"vendor/nvidia/model/a100/ram/80Gi", "true"},
			},
		},
		{
			name: "multiple models with and without RAM",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a100
        ram: 80Gi
      - model: a100`,
			expectedUnits: 1,
			expectedAttrs: []struct{ key, value string }{
				{"vendor/nvidia/model/a100", "true"},
				{"vendor/nvidia/model/a100/ram/80Gi", "true"},
			},
		},
		{
			name: "multiple different models",
			yaml: `units: 1
attributes:
  vendor:
    nvidia:
      - model: a6000
      - model: a40`,
			expectedUnits: 1,
			expectedAttrs: []struct{ key, value string }{
				{"vendor/nvidia/model/a40", "true"},
				{"vendor/nvidia/model/a6000", "true"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p v2ResourceGPU
			err := yaml.Unmarshal([]byte(tt.yaml), &p)

			if tt.shouldErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expectedUnits, p.Units)
			require.Len(t, p.Attributes, len(tt.expectedAttrs))

			for i, expected := range tt.expectedAttrs {
				require.Equal(t, expected.key, p.Attributes[i].Key)
				require.Equal(t, expected.value, p.Attributes[i].Value)
			}
		})
	}
}
