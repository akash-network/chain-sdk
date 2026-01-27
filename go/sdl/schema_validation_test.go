package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaValidation_Credentials(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "valid_credentials",
			builder: sdlTestBuilder{
				version: `version: "2.1"`,
				serviceBlock: `    image: nginx
    credentials:
      host: docker.io
      username: user123
      password: secret123`},
			shouldErr: false,
		},
		{
			name: "email_too_short",
			builder: sdlTestBuilder{
				version: `version: "2.1"`,
				serviceBlock: `    image: nginx
    credentials:
      host: docker.io
      username: user123
      password: secret123
      email: a@b`},
			shouldErr: true,
			reason:    "email must be at least 5 chars",
		},
		{
			name: "host_missing",
			builder: sdlTestBuilder{
				version: `version: "2.1"`,
				serviceBlock: `    image: nginx
    credentials:
      username: user123
      password: secret123`},
			shouldErr: true,
			reason:    "host is required",
		},
		{
			name: "password_too_short",
			builder: sdlTestBuilder{
				version: `version: "2.1"`,
				serviceBlock: `    image: nginx
    credentials:
      host: docker.io
      username: user123
      password: short`},
			shouldErr: true,
			reason:    "password must be at least 6 chars",
		},
		{
			name: "username_empty",
			builder: sdlTestBuilder{
				version: `version: "2.1"`,
				serviceBlock: `    image: nginx
    credentials:
      host: docker.io
      username: ""
      password: secret123`},
			shouldErr: true,
			reason:    "username cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputAgainstSchema([]byte(tt.builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept valid credentials")
			}
		})
	}
}

func TestSchemaValidation_Ports(t *testing.T) {
	tests := []struct {
		name      string
		port      string
		shouldErr bool
	}{
		{"valid_port_80", "80", false},
		{"valid_port_1", "1", false},
		{"valid_port_65535", "65535", false},
		{"invalid_port_0", "0", true},
		{"invalid_port_65536", "65536", true},
		{"invalid_port_-1", "-1", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := sdlTestBuilder{exposeBlock: `    expose:
      - port: ` + tt.port + `
        to:
          - global: true`}

			err := validateInputAgainstSchema([]byte(builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject port: %s", tt.port)
			} else {
				require.NoError(t, err, "Schema should accept port: %s", tt.port)
			}
		})
	}
}

func TestSchemaValidation_Protocol(t *testing.T) {
	tests := []struct {
		name      string
		proto     string
		shouldErr bool
	}{
		{"TCP_uppercase", "TCP", false},
		{"tcp_lowercase", "tcp", false},
		{"UDP_uppercase", "UDP", false},
		{"udp_lowercase", "udp", false},
		{"invalid_http", "http", true},
		{"invalid_HTTP", "HTTP", true},
		{"invalid_SCTP", "SCTP", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        proto: ` + tt.proto + `
        to:
          - global: true`}

			err := validateInputAgainstSchema([]byte(builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject protocol: %s", tt.proto)
			} else {
				require.NoError(t, err, "Schema should accept protocol: %s", tt.proto)
			}
		})
	}
}

func TestSchemaValidation_Denom(t *testing.T) {
	tests := []struct {
		name      string
		denom     string
		shouldErr bool
	}{
		{"valid_uakt", "uakt", false},
		{"valid_ibc_short", "ibc/ABC123", false},
		{"valid_ibc_long", "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", false},
		{"invalid_akt", "akt", true},
		{"invalid_empty", "", true},
		{"invalid_usdc", "usdc", true},
		{"invalid_ibc_without_slash", "ibcABC123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := sdlTestBuilder{placementBlock: `    dc:
      pricing:
        web:
          denom: ` + tt.denom + `
          amount: 100`}

			err := validateInputAgainstSchema([]byte(builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject denom: %s", tt.denom)
			} else {
				require.NoError(t, err, "Schema should accept denom: %s", tt.denom)
			}
		})
	}
}

func TestSchemaValidation_CPUUnits(t *testing.T) {
	tests := []struct {
		name      string
		units     string
		shouldErr bool
	}{
		{"valid_integer", "1", false},
		{"valid_decimal", "0.5", false},
		{"valid_milli", "100m", false},
		{"valid_string_number", `"2"`, false},
		{"invalid_negative", "-1", true},
		{"invalid_negative_decimal", "-0.5", true},
		{"invalid_negative_milli", "-100m", true},
		{"invalid_zero", "0", true},
		{"invalid_zero_decimal", "0.0", true},
		{"invalid_zero_padded", "00", true},
		{"invalid_zero_padded_decimal", "0.00", true},
		{"invalid_zero_mixed", "00.0", true},
		{"invalid_zero_milli", "0m", true},
		{"invalid_zero_padded_milli", "00m", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := sdlTestBuilder{resourcesBlock: `        cpu:
          units: ` + tt.units + `
        memory:
          size: 1Gi
        storage:
          - size: 1Gi`}

			err := validateInputAgainstSchema([]byte(builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject CPU units: %s", tt.units)
			} else {
				require.NoError(t, err, "Schema should accept CPU units: %s", tt.units)
			}
		})
	}
}

func TestSchemaValidation_EndpointNames(t *testing.T) {
	tests := []struct {
		name      string
		endpoint  string
		shouldErr bool
	}{
		{"valid_simple", "myendpoint", false},
		{"valid_with_dash", "my-endpoint", false},
		{"valid_with_underscore", "my_endpoint", false},
		{"valid_with_numbers", "endpoint123", false},
		{"valid_mixed", "my-endpoint_123", false},
		{"invalid_uppercase", "MyEndpoint", true},
		{"invalid_starts_with_number", "123endpoint", true},
		{"invalid_starts_with_dash", "-endpoint", true},
		{"invalid_starts_with_underscore", "_endpoint", true},
		{"invalid_space", "my endpoint", true},
		{"invalid_dot", "my.endpoint", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := sdlTestBuilder{
				endpoints: "endpoints:\n  " + tt.endpoint + ":\n    kind: ip",
				exposeBlock: `    expose:
      - port: 80
        to:
          - global: true
            ip: ` + tt.endpoint}

			err := validateInputAgainstSchema([]byte(builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject endpoint name: %s", tt.endpoint)
			} else {
				require.NoError(t, err, "Schema should accept endpoint name: %s", tt.endpoint)
			}
		})
	}
}

func TestSchemaValidation_Version(t *testing.T) {
	tests := []struct {
		name      string
		version   string
		shouldErr bool
	}{
		{"valid_2_0", `"2.0"`, false},
		{"valid_2_1", `"2.1"`, false},
		{"invalid_1_0", `"1.0"`, true},
		{"invalid_3_0", `"3.0"`, true},
		{"invalid_2_2", `"2.2"`, true},
		{"invalid_number", "2.0", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := sdlTestBuilder{version: "version: " + tt.version}

			err := validateInputAgainstSchema([]byte(builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject version: %s", tt.version)
			} else {
				require.NoError(t, err, "Schema should accept version: %s", tt.version)
			}
		})
	}
}

func TestSchemaValidation_RequiredFields(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		shouldErr bool
		reason    string
	}{
		{
			name: "missing_version",
			sdl: `services:
  web:
    image: nginx`,
			shouldErr: true,
			reason:    "version is required",
		},
		{
			name: "missing_deployment",
			sdl: `version: "2.0"
services:
  web:
    image: nginx`,
			shouldErr: true,
			reason:    "deployment is required",
		},
		{
			name: "missing_image_in_service",
			sdl: `version: "2.0"
services:
  web:
    expose:
      - port: 80
profiles:
  compute:
    web:
      resources:
        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
  placement:
    dc:
      pricing:
        web:
          denom: uakt
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: 1`,
			shouldErr: true,
			reason:    "image is required in service",
		},
		{
			name: "missing_size_in_storage",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
profiles:
  compute:
    web:
      resources:
        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - name: data
  placement:
    dc:
      pricing:
        web:
          denom: uakt
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: 1`,
			shouldErr: true,
			reason:    "size is required in storage",
		},
		{
			name: "missing_services",
			sdl: `version: "2.0"
profiles:
  compute:
    web:
      resources:
        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
  placement:
    dc:
      pricing:
        web:
          denom: uakt
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: 1`,
			shouldErr: true,
			reason:    "services is required",
		},
		{
			name: "missing_profiles",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - global: true
deployment:
  web:
    dc:
      profile: web
      count: 1`,
			shouldErr: true,
			reason:    "profiles is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputAgainstSchema([]byte(tt.sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept valid SDL")
			}
		})
	}
}

func TestSchemaValidation_GPUUnitsRequireAttributes(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "gpu_with_units_gt_0_and_attributes",
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
                - model: a100`},
			shouldErr: false,
		},
		{
			name: "gpu_with_units_0_without_attributes",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: 0`},
			shouldErr: false,
			reason:    "units=0 does not require attributes",
		},
		{
			name: "gpu_with_string_units_0_without_attributes",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: "0"`},
			shouldErr: false,
			reason:    "units='0' does not require attributes",
		},
		{
			name: "gpu_with_padded_zero_without_attributes",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: "00"`},
			shouldErr: false,
			reason:    "units='00' is zero, does not require attributes",
		},
		{
			name: "gpu_with_decimal_zero_without_attributes",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: "0.00"`},
			shouldErr: false,
			reason:    "units='0.00' is zero, does not require attributes",
		},
		{
			name: "gpu_with_units_gt_0_without_attributes",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: 1`},
			shouldErr: true,
			reason:    "units > 0 requires attributes",
		},
		{
			name: "gpu_with_string_units_gt_0_without_attributes",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: "1"`},
			shouldErr: true,
			reason:    "units > 0 requires attributes",
		},
		{
			name: "gpu_with_decimal_units_without_attributes",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: 0.5`},
			shouldErr: true,
			reason:    "units > 0 requires attributes",
		},
		{
			name: "gpu_section_empty",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu: {}`},
			shouldErr: false,
			reason:    "empty gpu defaults to units=0 in Go",
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

func TestSchemaValidation_GPUAttributesRequireUnits(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "attributes_without_units_field",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          attributes:
            vendor:
              nvidia:
                - model: a100`},
			shouldErr: true,
			reason:    "attributes present but units not specified",
		},
		{
			name: "attributes_with_units_0",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: 0
          attributes:
            vendor:
              nvidia:
                - model: a100`},
			shouldErr: true,
			reason:    "attributes present with units=0",
		},
		{
			name: "attributes_with_string_units_0",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: "0"
          attributes:
            vendor:
              nvidia:
                - model: a100`},
			shouldErr: true,
			reason:    "attributes present with units='0'",
		},
		{
			name: "attributes_with_units_gt_0",
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
                - model: a100`},
			shouldErr: false,
			reason:    "valid: both units and attributes",
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
