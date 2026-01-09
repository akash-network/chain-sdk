package sdl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSchemaValidation_HTTPOptions_Limits tests that schema enforces HTTP option limits
func TestSchemaValidation_HTTPOptions_Limits(t *testing.T) {
	tests := []struct {
		name        string
		httpOptions string
		shouldErr   bool
		reason      string
	}{
		{
			name: "max_body_size within limit",
			httpOptions: `http_options:
                max_body_size: 1048576`,
			shouldErr: false,
		},
		{
			name: "max_body_size at upper limit",
			httpOptions: `http_options:
                max_body_size: 104857600`,
			shouldErr: false,
			reason:    "100 MB should be accepted",
		},
		{
			name: "max_body_size exceeds limit",
			httpOptions: `http_options:
                max_body_size: 104857601`,
			shouldErr: true,
			reason:    "Should reject > 100 MB",
		},
		{
			name: "read_timeout within limit",
			httpOptions: `http_options:
                read_timeout: 60000`,
			shouldErr: false,
		},
		{
			name: "read_timeout exceeds limit",
			httpOptions: `http_options:
                read_timeout: 60001`,
			shouldErr: true,
			reason:    "Should reject > 60000 ms",
		},
		{
			name: "send_timeout within limit",
			httpOptions: `http_options:
                send_timeout: 60000`,
			shouldErr: false,
		},
		{
			name: "send_timeout exceeds limit",
			httpOptions: `http_options:
                send_timeout: 60001`,
			shouldErr: true,
			reason:    "Should reject > 60000 ms",
		},
		{
			name: "negative max_body_size",
			httpOptions: `http_options:
                max_body_size: -1`,
			shouldErr: true,
			reason:    "Negative values should be rejected",
		},
		{
			name: "negative read_timeout",
			httpOptions: `http_options:
                read_timeout: -1`,
			shouldErr: true,
			reason:    "Negative values should be rejected",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        %s
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
`, tt.httpOptions)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_Accept_Hostnames tests accept item validation
func TestSchemaValidation_Accept_Hostnames(t *testing.T) {
	tests := []struct {
		name      string
		accept    string
		shouldErr bool
		reason    string
	}{
		{
			name:      "valid hostname",
			accept:    "example.com",
			shouldErr: false,
		},
		{
			name:      "valid subdomain",
			accept:    "api.example.com",
			shouldErr: false,
		},
		{
			name:      "wildcard subdomain",
			accept:    `"*.example.com"`,
			shouldErr: false,
		},
		// Note: Go код використовує url.ParseRequestURI("http://" + item)
		// Це означає що валідація досить ліберальна
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        accept:
          - %s
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
`, tt.accept)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_GPUInterface tests GPU interface enum
func TestSchemaValidation_GPUInterface(t *testing.T) {
	tests := []struct {
		name      string
		interface_ string
		shouldErr bool
	}{
		{"valid pcie", "pcie", false},
		{"valid sxm", "sxm", false},
		{"invalid PCIe", "PCIe", true},
		{"invalid nvlink", "nvlink", true},
		// Note: empty interface is actually allowed in schema (optional field)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interfaceStr := fmt.Sprintf("interface: %s", tt.interface_)
			
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
        gpu:
          units: 1
          attributes:
            vendor:
              nvidia:
                - model: a100
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
`, interfaceStr)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject interface: %s", tt.interface_)
			} else {
				require.NoError(t, err, "Schema should accept interface: %s", tt.interface_)
			}
		})
	}
}

// TestSchemaValidation_DeploymentCount tests count minimum validation
func TestSchemaValidation_DeploymentCount(t *testing.T) {
	tests := []struct {
		name      string
		count     string
		shouldErr bool
	}{
		{"valid count 1", "1", false},
		{"valid count 10", "10", false},
		{"invalid count 0", "0", true},
		{"invalid count -1", "-1", true},
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
          denom: uakt
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: %s
`, tt.count)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject count: %s", tt.count)
			} else {
				require.NoError(t, err, "Schema should accept count: %s", tt.count)
			}
		})
	}
}

// TestSchemaValidation_PricingAmount tests negative amount rejection
func TestSchemaValidation_PricingAmount(t *testing.T) {
	tests := []struct {
		name      string
		amount    string
		shouldErr bool
	}{
		{"valid positive", "100", false},
		{"valid zero", "0", false},
		{"valid decimal", "0.5", false},
		{"invalid negative", "-1", true},
		{"invalid negative decimal", "-0.5", true},
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
          denom: uakt
          amount: %s
deployment:
  web:
    dc:
      profile: web
      count: 1
`, tt.amount)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject amount: %s", tt.amount)
			} else {
				require.NoError(t, err, "Schema should accept amount: %s", tt.amount)
			}
		})
	}
}

// TestSchemaValidation_EmptyImage tests that empty image is rejected
func TestSchemaValidation_EmptyImage(t *testing.T) {
	sdl := `version: "2.0"
services:
  web:
    image: ""
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
`
	err := validateInputAgainstSchema([]byte(sdl))
	require.Error(t, err, "Schema should reject empty image")
}

