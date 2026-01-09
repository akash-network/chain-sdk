package sdl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSchemaValidation_ExtremeValues tests extreme numeric values
func TestSchemaValidation_ExtremeValues(t *testing.T) {
	tests := []struct {
		name      string
		field     string
		value     string
		shouldErr bool
		reason    string
	}{
		{
			name:      "CPU units extremely large",
			field:     "cpu:\n          units: 999999999",
			shouldErr: false, // Go code might accept but is it reasonable?
		},
		{
			name:      "memory size invalid unit",
			field:     "memory:\n          size: 1ZB",
			shouldErr: true,
			reason:    "Invalid memory unit",
		},
		{
			name:      "storage size zero",
			field:     "storage:\n          - size: 0",
			shouldErr: true,
			reason:    "Zero storage should be invalid",
		},
		{
			name:      "pricing amount string overflow",
			field:     "pricing:\n        web:\n          denom: uakt\n          amount: '99999999999999999999999999999'",
			shouldErr: false, // String numbers can be very large
		},
		{
			name:      "port exactly 65535",
			field:     "port: 65535",
			shouldErr: false,
		},
		{
			name:      "count extremely large",
			field:     "count: 999999",
			shouldErr: false, // Should schema limit this?
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sdl string
			if strings.Contains(tt.field, "port:") {
				sdl = fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
    expose:
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
`, tt.field)
			} else if strings.Contains(tt.field, "count:") {
				sdl = fmt.Sprintf(`version: "2.0"
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
      %s
`, tt.field)
			} else if strings.Contains(tt.field, "pricing:") {
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
      %s
deployment:
  web:
    dc:
      profile: web
      count: 1
`, tt.field)
				err := validateInputAgainstSchema([]byte(sdl))
				if tt.shouldErr {
					require.Error(t, err, "Schema should reject: %s", tt.reason)
				} else {
					require.NoError(t, err, "Schema should accept: %s", tt.reason)
				}
				return
			} else {
				sdl = fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
profiles:
  compute:
    web:
      resources:
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
`, tt.field)
			}

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_SpecialCharacters tests special characters in various fields
func TestSchemaValidation_SpecialCharacters(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		shouldErr bool
		reason    string
	}{
		{
			name: "service name with special chars",
			sdl: `version: "2.0"
services:
  web@service:
    image: nginx
profiles:
  compute:
    web@service:
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
        web@service:
          denom: uakt
          amount: 1
deployment:
  web@service:
    dc:
      profile: web@service
      count: 1`,
			shouldErr: true, // Should service names allow @?
			reason:    "Service names should be restricted",
		},
		{
			name: "image with spaces",
			sdl: `version: "2.0"
services:
  web:
    image: "nginx latest"
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
			shouldErr: false, // Spaces in image names are valid (image tags)
		},
		{
			name: "accept with unicode",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        accept:
          - "こんにちは.com"
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
      count: 1`,
			shouldErr: false, // Unicode domains are valid
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputAgainstSchema([]byte(tt.sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_EmptyAndNullValues tests empty and null handling
func TestSchemaValidation_EmptyAndNullValues(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		shouldErr bool
		reason    string
	}{
		{
			name: "null args",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    args: null
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
			shouldErr: false,
			reason:    "null args is allowed",
		},
		{
			name: "empty array args",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    args: []
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
			shouldErr: false,
			reason:    "empty args array is allowed",
		},
		{
			name: "empty storage array",
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
        storage: []
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
			reason:    "storage is required",
		},
		{
			name: "null expose",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    expose: null
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
			shouldErr: false,
			reason:    "expose is optional",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputAgainstSchema([]byte(tt.sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_DuplicateNames tests duplicate name handling
func TestSchemaValidation_DuplicateNames(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		shouldErr bool
		reason    string
	}{
		{
			name: "duplicate service names",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
  web:
    image: apache
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
			shouldErr: false, // YAML will overwrite, schema won't catch
			reason:    "YAML handles duplicates by overwriting",
		},
		{
			name: "duplicate endpoint names",
			sdl: `version: "2.0"
endpoints:
  myip:
    kind: ip
  myip:
    kind: ip
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
      count: 1`,
			shouldErr: false, // YAML handles duplicates
			reason:    "YAML parser handles this",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputAgainstSchema([]byte(tt.sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_WhitespaceHandling tests whitespace in various contexts
func TestSchemaValidation_WhitespaceHandling(t *testing.T) {
	tests := []struct {
		name      string
		image     string
		shouldErr bool
		reason    string
	}{
		{
			name:      "image with leading space",
			image:     `" nginx"`,
			shouldErr: false, // Spaces in strings are allowed
		},
		{
			name:      "image with trailing space",
			image:     `"nginx "`,
			shouldErr: false,
		},
		{
			name:      "image with only spaces",
			image:     `"   "`,
			shouldErr: true, // minLength should catch this
			reason:    "Only whitespace should be rejected",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
services:
  web:
    image: %s
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
`, tt.image)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_CaseSensitivity tests case sensitivity
func TestSchemaValidation_CaseSensitivity(t *testing.T) {
	tests := []struct {
		name      string
		field     string
		value     string
		shouldErr bool
		reason    string
	}{
		{
			name:      "version uppercase",
			field:     "VERSION",
			value:     `"2.0"`,
			shouldErr: true,
			reason:    "Field names are case sensitive",
		},
		{
			name:      "denom uppercase UAKT",
			field:     "denom",
			value:     "UAKT",
			shouldErr: true,
			reason:    "denom must be lowercase uakt",
		},
		{
			name:      "protocol TCP",
			field:     "proto",
			value:     "TCP",
			shouldErr: false,
			reason:    "TCP is valid",
		},
		{
			name:      "protocol Tcp",
			field:     "proto",
			value:     "Tcp",
			shouldErr: true,
			reason:    "Mixed case protocol invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sdl string
			if tt.field == "VERSION" {
				sdl = fmt.Sprintf(`%s: %s
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
      count: 1`, tt.field, tt.value)
			} else if tt.field == "denom" {
				sdl = fmt.Sprintf(`version: "2.0"
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
          %s: %s
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: 1`, tt.field, tt.value)
			} else if tt.field == "proto" {
				sdl = fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        %s: %s
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
      count: 1`, tt.field, tt.value)
			}

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_IPWithoutGlobal tests IP endpoint requires global
func TestSchemaValidation_IPWithoutGlobal(t *testing.T) {
	sdl := `version: "2.0"
endpoints:
  myip:
    kind: ip
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - ip: myip
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
      count: 1`

	err := validateInputAgainstSchema([]byte(sdl))
	require.Error(t, err, "IP without global: true should be rejected")
}

// TestSchemaValidation_StorageNameConflicts tests storage name handling
func TestSchemaValidation_StorageNameConflicts(t *testing.T) {
	tests := []struct {
		name      string
		storage   string
		shouldErr bool
		reason    string
	}{
		{
			name: "duplicate storage names",
			storage: `- size: 1Gi
            name: data
          - size: 2Gi
            name: data`,
			shouldErr: false, // Schema doesn't validate this, Go does
			reason:    "Schema doesn't catch duplicate names",
		},
		{
			name: "reserved name 'default'",
			storage: `- size: 1Gi
            name: default`,
			shouldErr: false, // 'default' is actually valid
		},
		{
			name: "storage name with special chars",
			storage: `- size: 1Gi
            name: "my-data_123"`,
			shouldErr: false, // Names are flexible
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
`, tt.storage)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_ZeroValues tests zero value handling
func TestSchemaValidation_ZeroValues(t *testing.T) {
	tests := []struct {
		name      string
		field     string
		value     string
		shouldErr bool
		reason    string
	}{
		{
			name:      "GPU units zero",
			field:     "gpu:\n          units: 0",
			shouldErr: false, // Zero GPU is valid (no GPU)
		},
		{
			name:      "CPU units zero",
			field:     "cpu:\n          units: 0",
			shouldErr: true,
			reason:    "CPU cannot be zero",
		},
		{
			name:      "pricing amount zero",
			field:     "amount: 0",
			shouldErr: false, // Zero price is technically valid
		},
		{
			name:      "next_tries zero",
			field:     "next_tries: 0",
			shouldErr: false, // Zero might mean no retries
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sdl string
			if strings.Contains(tt.field, "amount:") {
				sdl = fmt.Sprintf(`version: "2.0"
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
          %s
deployment:
  web:
    dc:
      profile: web
      count: 1`, tt.field)
			} else if strings.Contains(tt.field, "next_tries:") {
				sdl = fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - global: true
            http_options:
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
      count: 1`, tt.field)
			} else {
				sdl = fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
profiles:
  compute:
    web:
      resources:
        %s
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
      count: 1`, tt.field)
			}

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

