package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaValidation_AdditionalProperties(t *testing.T) {
	tests := []struct {
		name    string
		builder sdlTestBuilder
		reason  string
	}{
		{
			name:    "unknown_field_in_service",
			builder: sdlTestBuilder{serviceBlock: "    unknown_field: value"},
			reason:  "service should not allow unknown fields",
		},
		{
			name: "unknown_field_in_credentials",
			builder: sdlTestBuilder{serviceBlock: `    credentials:
      host: docker.io
      username: user123
      password: secret123
      unknown_field: value`},
			reason: "credentials should not allow unknown fields",
		},
		{
			name: "unknown_field_in_dependencies_item",
			builder: sdlTestBuilder{
				hasTwoServices: true,
				serviceBlock: `    dependencies:
      - service: db
        unknown_field: value`},
			reason: "dependencies items should not allow unknown fields",
		},
		{
			name: "unknown_field_in_expose_item",
			builder: sdlTestBuilder{
				serviceBlock: `    expose:
      - port: 80
        unknown_field: value
        to:
          - global: true`},
			reason: "expose items should not allow unknown fields",
		},
		{
			name: "unknown_field_in_expose_to_item",
			builder: sdlTestBuilder{serviceBlock: `    expose:
      - port: 80
        to:
          - global: true
            unknown_field: value`},
			reason: "expose.to items should not allow unknown fields",
		},
		{
			name: "unknown_field_in_http_options",
			builder: sdlTestBuilder{serviceBlock: `    expose:
      - port: 80
        http_options:
          max_body_size: 1048576
          unknown_field: value
        to:
          - global: true`},
			reason: "http_options should not allow unknown fields",
		},
		{
			name: "unknown_field_in_cpu",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
          unknown_field: value
        memory:
          size: 1Gi
        storage:
          - size: 1Gi`},
			reason: "cpu should not allow unknown fields",
		},
		{
			name: "unknown_field_in_memory",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
          unknown_field: value
        storage:
          - size: 1Gi`},
			reason: "memory should not allow unknown fields",
		},
		{
			name: "unknown_field_in_gpu",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: 1
          unknown_field: value`},
			reason: "gpu should not allow unknown fields",
		},
		{
			name: "unknown_field_in_gpu_attributes",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: 1
          attributes:
            unknown_field: value`},
			reason: "gpu.attributes should only allow vendor",
		},
		{
			name: "unknown_field_in_nvidia_gpu_item",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: 1
          attributes:
            vendor:
              nvidia:
                - model: a100
                  unknown_field: value`},
			reason: "nvidia gpu items should not allow unknown fields",
		},
		{
			name: "unknown_field_in_storage_item",
			builder: sdlTestBuilder{
				serviceBlock: `    params:
      storage:
        data:
          mount: /data`,
				resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: data
            unknown_field: value`},
			reason: "storage items should not allow unknown fields",
		},
		{
			name: "unknown_field_in_storage_attributes",
			builder: sdlTestBuilder{
				serviceBlock: `    params:
      storage:
        data:
          mount: /data`,
				resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: data
            attributes:
              persistent: true
              class: beta1
              unknown_field: value`},
			reason: "storage.attributes should not allow unknown fields",
		},
		{
			name: "unknown_field_in_params_storage_item",
			builder: sdlTestBuilder{
				serviceBlock: `    params:
      storage:
        data:
          mount: /data
          unknown_field: value`,
				resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: data`},
			reason: "params.storage items should not allow unknown fields",
		},
		{
			name: "unknown_field_in_endpoint",
			builder: sdlTestBuilder{
				endpoints: `endpoints:
  myip:
    kind: ip
    unknown_field: value`,
				serviceBlock: `    expose:
      - port: 80
        to:
          - ip: myip
            global: true`},
			reason: "endpoints should not allow unknown fields",
		},
		{
			name:    "unknown_field_in_placement",
			builder: sdlTestBuilder{placementBlock: "      unknown_field: value"},
			reason:  "placement items should not allow unknown fields",
		},
		{
			name: "unknown_field_in_pricing_item",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi`,
				placementBlock: `      pricing:
        web:
          denom: uakt
          amount: 1
          unknown_field: value`},
			reason: "pricing items should not allow unknown fields",
		},
		{
			name:    "unknown_field_in_deployment_item",
			builder: sdlTestBuilder{deploymentBlock: "      unknown_field: value"},
			reason:  "deployment items should not allow unknown fields",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputAgainstSchema([]byte(tt.builder.build()))
			require.Error(t, err, "Schema should reject: %s", tt.reason)
			require.Contains(t, err.Error(), "Additional property",
				"Error should mention Additional property")
		})
	}
}
