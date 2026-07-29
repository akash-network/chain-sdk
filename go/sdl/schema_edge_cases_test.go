package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaValidation_HTTPOptions_Limits(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "max_body_size_within_limit",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          max_body_size: 1048576
        to:
          - global: true`},
			shouldErr: false,
		},
		{
			name: "max_body_size_at_upper_limit",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          max_body_size: 104857600
        to:
          - global: true`},
			shouldErr: false,
			reason:    "100 MB should be accepted",
		},
		{
			name: "max_body_size_exceeds_limit",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          max_body_size: 104857601
        to:
          - global: true`},
			shouldErr: true,
			reason:    "Should reject > 100 MB",
		},
		{
			name: "read_timeout_within_limit",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          read_timeout: 60000
        to:
          - global: true`},
			shouldErr: false,
		},
		{
			name: "read_timeout_exceeds_limit",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          read_timeout: 60001
        to:
          - global: true`},
			shouldErr: true,
			reason:    "Should reject > 60000 ms",
		},
		{
			name: "send_timeout_within_limit",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          send_timeout: 60000
        to:
          - global: true`},
			shouldErr: false,
		},
		{
			name: "send_timeout_exceeds_limit",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          send_timeout: 60001
        to:
          - global: true`},
			shouldErr: true,
			reason:    "Should reject > 60000 ms",
		},
		{
			name: "negative_max_body_size",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          max_body_size: -1
        to:
          - global: true`},
			shouldErr: true,
			reason:    "Negative values should be rejected",
		},
		{
			name: "negative_read_timeout",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          read_timeout: -1
        to:
          - global: true`},
			shouldErr: true,
			reason:    "Negative values should be rejected",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputAgainstSchema([]byte(tt.builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

func TestSchemaValidation_Accept_Hostnames(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "valid_hostname",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        accept:
          - example.com
        to:
          - global: true`},
			shouldErr: false,
		},
		{
			name: "valid_subdomain",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        accept:
          - api.example.com
        to:
          - global: true`},
			shouldErr: false,
		},
		{
			name: "wildcard_subdomain",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        accept:
          - "*.example.com"
        to:
          - global: true`},
			shouldErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputAgainstSchema([]byte(tt.builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

func TestSchemaValidation_GPUInterface(t *testing.T) {
	tests := []struct {
		name      string
		iface     string
		shouldErr bool
	}{
		{"valid_pcie", "pcie", false},
		{"valid_sxm", "sxm", false},
		{"invalid_PCIe", "PCIe", true},
		{"invalid_nvlink", "nvlink", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := sdlTestBuilder{resourcesBlock: `        cpu:
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
                  interface: ` + tt.iface}

			err := validateInputAgainstSchema([]byte(builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject interface: %s", tt.iface)
			} else {
				require.NoError(t, err, "Schema should accept interface: %s", tt.iface)
			}
		})
	}
}

func TestSchemaValidation_DeploymentCount(t *testing.T) {
	tests := []struct {
		name      string
		count     string
		shouldErr bool
	}{
		{"valid_count_1", "1", false},
		{"valid_count_10", "10", false},
		{"invalid_count_0", "0", true},
		{"invalid_count_-1", "-1", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := sdlTestBuilder{deploymentBlock: `  web:
    dc:
      profile: web
      count: ` + tt.count}

			err := validateInputAgainstSchema([]byte(builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject count: %s", tt.count)
			} else {
				require.NoError(t, err, "Schema should accept count: %s", tt.count)
			}
		})
	}
}

func TestSchemaValidation_PricingAmount(t *testing.T) {
	tests := []struct {
		name      string
		amount    string
		shouldErr bool
	}{
		{"valid_positive", "100", false},
		{"valid_zero", "0", false},
		{"valid_decimal", "0.5", false},
		{"invalid_negative", "-1", true},
		{"invalid_negative_decimal", "-0.5", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := sdlTestBuilder{placementBlock: `    dc:
      pricing:
        web:
          denom: uakt
          amount: ` + tt.amount}

			err := validateInputAgainstSchema([]byte(builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject amount: %s", tt.amount)
			} else {
				require.NoError(t, err, "Schema should accept amount: %s", tt.amount)
			}
		})
	}
}

func TestSchemaValidation_EmptyImage(t *testing.T) {
	builder := sdlTestBuilder{serviceBlock: `    image: ""`}
	err := validateInputAgainstSchema([]byte(builder.build()))
	require.Error(t, err, "Schema should reject empty image")
}
