package sdl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSchemaValidation_Credentials tests schema validation for credentials field
func TestSchemaValidation_Credentials(t *testing.T) {
	tests := []struct {
		name      string
		creds     string
		shouldErr bool
		reason    string
	}{
		{
			name: "valid credentials",
			creds: `host: docker.io
      username: user123
      password: secret123`,
			shouldErr: false,
		},
		{
			name: "email too short",
			creds: `host: docker.io
      username: user123
      password: secret123
      email: a@b`,
			shouldErr: true,
			reason:    "email must be at least 5 chars",
		},
		{
			name: "host missing",
			creds: `username: user123
      password: secret123`,
			shouldErr: true,
			reason:    "host is required",
		},
		{
			name: "password too short",
			creds: `host: docker.io
      username: user123
      password: short`,
			shouldErr: true,
			reason:    "password must be at least 6 chars",
		},
		{
			name: "username empty",
			creds: `host: docker.io
      username: ""
      password: secret123`,
			shouldErr: true,
			reason:    "username cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.1"
services:
  web:
    image: nginx
    credentials:
      %s
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
      count: 1
`, tt.creds)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept valid credentials")
			}
		})
	}
}

// TestSchemaValidation_Ports tests port number validation
func TestSchemaValidation_Ports(t *testing.T) {
	tests := []struct {
		name      string
		port      string
		shouldErr bool
	}{
		{"valid port 80", "80", false},
		{"valid port 1", "1", false},
		{"valid port 65535", "65535", false},
		{"invalid port 0", "0", true},
		{"invalid port 65536", "65536", true},
		{"invalid port -1", "-1", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: %s
        to:
          - global: true
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
      count: 1
`, tt.port)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject port: %s", tt.port)
			} else {
				require.NoError(t, err, "Schema should accept port: %s", tt.port)
			}
		})
	}
}

// TestSchemaValidation_Protocol tests protocol validation
func TestSchemaValidation_Protocol(t *testing.T) {
	tests := []struct {
		name      string
		proto     string
		shouldErr bool
	}{
		{"TCP uppercase", "TCP", false},
		{"tcp lowercase", "tcp", false},
		{"UDP uppercase", "UDP", false},
		{"udp lowercase", "udp", false},
		{"invalid http", "http", true},
		{"invalid HTTP", "HTTP", true},
		{"invalid SCTP", "SCTP", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        proto: %s
        to:
          - global: true
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
      count: 1
`, tt.proto)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject protocol: %s", tt.proto)
			} else {
				require.NoError(t, err, "Schema should accept protocol: %s", tt.proto)
			}
		})
	}
}

// TestSchemaValidation_Denom tests denom pattern validation
func TestSchemaValidation_Denom(t *testing.T) {
	tests := []struct {
		name      string
		denom     string
		shouldErr bool
	}{
		{"valid uakt", "uakt", false},
		{"valid ibc short", "ibc/ABC123", false},
		{"valid ibc long", "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", false},
		{"invalid akt", "akt", true},
		{"invalid empty", "", true},
		{"invalid usdc", "usdc", true},
		{"invalid ibc without slash", "ibcABC123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
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
          - size: 1Gi
  placement:
    dc:
      pricing:
        web:
          denom: %s
          amount: 100
deployment:
  web:
    dc:
      profile: web
      count: 1
`, tt.denom)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject denom: %s", tt.denom)
			} else {
				require.NoError(t, err, "Schema should accept denom: %s", tt.denom)
			}
		})
	}
}

// TestSchemaValidation_CPUUnits tests CPU units validation
func TestSchemaValidation_CPUUnits(t *testing.T) {
	tests := []struct {
		name      string
		units     string
		shouldErr bool
	}{
		{"valid integer", "1", false},
		{"valid decimal", "0.5", false},
		{"valid milli", "100m", false},
		{"valid string number", `"2"`, false},
		{"invalid negative", "-1", true},
		{"invalid negative decimal", "-0.5", true},
		{"invalid negative milli", "-100m", true},
		{"invalid zero", "0", true},
		{"invalid zero decimal", "0.0", true},
		{"invalid zero padded", "00", true},
		{"invalid zero padded decimal", "0.00", true},
		{"invalid zero mixed", "00.0", true},
		{"invalid zero milli", "0m", true},
		{"invalid zero padded milli", "00m", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
profiles:
  compute:
    web:
      resources:
        cpu:
          units: %s
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
      count: 1
`, tt.units)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject CPU units: %s", tt.units)
			} else {
				require.NoError(t, err, "Schema should accept CPU units: %s", tt.units)
			}
		})
	}
}

// TestSchemaValidation_EndpointNames tests endpoint name pattern
func TestSchemaValidation_EndpointNames(t *testing.T) {
	tests := []struct {
		name      string
		endpoint  string
		shouldErr bool
	}{
		{"valid simple", "myendpoint", false},
		{"valid with dash", "my-endpoint", false},
		{"valid with underscore", "my_endpoint", false},
		{"valid with numbers", "endpoint123", false},
		{"valid mixed", "my-endpoint_123", false},
		{"invalid uppercase", "MyEndpoint", true},
		{"invalid starts with number", "123endpoint", true},
		{"invalid starts with dash", "-endpoint", true},
		{"invalid starts with underscore", "_endpoint", true},
		{"invalid space", "my endpoint", true},
		{"invalid dot", "my.endpoint", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
endpoints:
  %s:
    kind: ip
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - global: true
            ip: %s
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
      count: 1
`, tt.endpoint, tt.endpoint)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject endpoint name: %s", tt.endpoint)
			} else {
				require.NoError(t, err, "Schema should accept endpoint name: %s", tt.endpoint)
			}
		})
	}
}

// TestSchemaValidation_Version tests version enum validation
func TestSchemaValidation_Version(t *testing.T) {
	tests := []struct {
		name      string
		version   string
		shouldErr bool
	}{
		{"valid 2.0", `"2.0"`, false},
		{"valid 2.1", `"2.1"`, false},
		{"invalid 1.0", `"1.0"`, true},
		{"invalid 3.0", `"3.0"`, true},
		{"invalid 2.2", `"2.2"`, true},
		{"invalid number", "2.0", true}, // without quotes
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: %s
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
      count: 1
`, tt.version)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject version: %s", tt.version)
			} else {
				require.NoError(t, err, "Schema should accept version: %s", tt.version)
			}
		})
	}
}

// TestSchemaValidation_RequiredFields tests that required fields are enforced
func TestSchemaValidation_RequiredFields(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		shouldErr bool
		reason    string
	}{
		{
			name: "missing version",
			sdl: `services:
  web:
    image: nginx`,
			shouldErr: true,
			reason:    "version is required",
		},
		{
			name: "missing deployment",
			sdl: `version: "2.0"
services:
  web:
    image: nginx`,
			shouldErr: true,
			reason:    "deployment is required",
		},
		{
			name: "missing image in service",
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
			name: "missing size in storage",
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
			name: "missing services",
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
			name: "missing profiles",
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
		gpu       string
		shouldErr bool
		reason    string
	}{
		{
			name: "gpu with units > 0 and attributes",
			gpu: `gpu:
          units: 1
          attributes:
            vendor:
              nvidia:
                - model: a100`,
			shouldErr: false,
		},
		{
			name: "gpu with units 0 without attributes",
			gpu: `gpu:
          units: 0`,
			shouldErr: false,
			reason:    "units=0 does not require attributes",
		},
		{
			name: "gpu with string units 0 without attributes",
			gpu: `gpu:
          units: "0"`,
			shouldErr: false,
			reason:    "units='0' does not require attributes",
		},
		{
			name: "gpu with padded zero without attributes",
			gpu: `gpu:
          units: "00"`,
			shouldErr: false,
			reason:    "units='00' is zero, does not require attributes",
		},
		{
			name: "gpu with decimal zero without attributes",
			gpu: `gpu:
          units: "0.00"`,
			shouldErr: false,
			reason:    "units='0.00' is zero, does not require attributes",
		},
		{
			name: "gpu with units > 0 without attributes",
			gpu: `gpu:
          units: 1`,
			shouldErr: true,
			reason:    "units > 0 requires attributes",
		},
		{
			name: "gpu with string units > 0 without attributes",
			gpu: `gpu:
          units: "1"`,
			shouldErr: true,
			reason:    "units > 0 requires attributes",
		},
		{
			name: "gpu with decimal units without attributes",
			gpu: `gpu:
          units: 0.5`,
			shouldErr: true,
			reason:    "units > 0 requires attributes",
		},
		{
			name: "gpu section empty",
			gpu: `gpu: {}`,
			shouldErr: false,
			reason:    "empty gpu defaults to units=0 in Go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
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
          - size: 1Gi
        %s
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
      count: 1
`, tt.gpu)

			err := validateInputAgainstSchema([]byte(sdl))
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
		gpu       string
		shouldErr bool
		reason    string
	}{
		{
			name: "attributes without units field",
			gpu: `gpu:
          attributes:
            vendor:
              nvidia:
                - model: a100`,
			shouldErr: true,
			reason:    "attributes present but units not specified",
		},
		{
			name: "attributes with units 0",
			gpu: `gpu:
          units: 0
          attributes:
            vendor:
              nvidia:
                - model: a100`,
			shouldErr: true,
			reason:    "attributes present with units=0",
		},
		{
			name: "attributes with string units 0",
			gpu: `gpu:
          units: "0"
          attributes:
            vendor:
              nvidia:
                - model: a100`,
			shouldErr: true,
			reason:    "attributes present with units='0'",
		},
		{
			name: "attributes with units > 0",
			gpu: `gpu:
          units: 1
          attributes:
            vendor:
              nvidia:
                - model: a100`,
			shouldErr: false,
			reason:    "valid: both units and attributes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
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
          - size: 1Gi
        %s
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
      count: 1
`, tt.gpu)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}
