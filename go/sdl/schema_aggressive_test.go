package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaValidation_ExtremeValues(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "cpu_units_extremely_large",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 999999999
        memory:
          size: 1Gi
        storage:
          - size: 1Gi`},
			shouldErr: false,
		},
		{
			name: "memory_size_invalid_unit",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1ZB
        storage:
          - size: 1Gi`},
			shouldErr: false,
			reason:    "Schema doesn't validate memory units strictly",
		},
		{
			name: "storage_size_zero",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 0`},
			shouldErr: true,
			reason:    "Zero storage should be invalid",
		},
		{
			name: "pricing_amount_string_overflow",
			builder: sdlTestBuilder{placementBlock: `    dc:
      pricing:
        web:
          denom: uakt
          amount: '99999999999999999999999999999'`},
			shouldErr: false,
		},
		{
			name: "port_exactly_65535",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 65535
        to:
          - global: true`},
			shouldErr: false,
		},
		{
			name: "count_extremely_large",
			builder: sdlTestBuilder{deploymentBlock: `  web:
    dc:
      profile: web
      count: 999999`},
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

func TestSchemaValidation_SpecialCharacters(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "service_name_with_special_chars",
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
			shouldErr: false,
			reason:    "Service names with @ are allowed by Go parser",
		},
		{
			name:      "image_with_spaces",
			builder:   sdlTestBuilder{serviceBlock: `    image: "nginx latest"`},
			shouldErr: false,
		},
		{
			name: "accept_with_unicode",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        accept:
          - "こんにちは.com"
        to:
          - global: true`},
			shouldErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := tt.sdl
			if sdl == "" {
				sdl = tt.builder.build()
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

func TestSchemaValidation_EmptyAndNullValues(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name:      "null_args",
			builder:   sdlTestBuilder{serviceBlock: "    args: null"},
			shouldErr: false,
			reason:    "null args is allowed",
		},
		{
			name:      "empty_array_args",
			builder:   sdlTestBuilder{serviceBlock: "    args: []"},
			shouldErr: false,
			reason:    "empty args array is allowed",
		},
		{
			name: "empty_storage_array",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage: []`},
			shouldErr: false,
			reason:    "Empty storage array passes schema but fails Go validation",
		},
		{
			name:      "null_expose",
			builder:   sdlTestBuilder{serviceBlock: "    expose: null"},
			shouldErr: false,
			reason:    "expose is optional",
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

func TestSchemaValidation_DuplicateNames(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "duplicate_service_names",
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
			shouldErr: true,
			reason:    "YAML v3 parser rejects duplicate mapping keys",
		},
		{
			name: "duplicate_endpoint_names",
			builder: sdlTestBuilder{
				endpoints: `endpoints:
  myip:
    kind: ip
  myip:
    kind: ip`,
			},
			shouldErr: true,
			reason:    "YAML v3 parser rejects duplicate mapping keys",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := tt.sdl
			if sdl == "" {
				sdl = tt.builder.build()
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

func TestSchemaValidation_WhitespaceHandling(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name:      "image_with_leading_space",
			builder:   sdlTestBuilder{serviceBlock: `    image: " nginx"`},
			shouldErr: false,
		},
		{
			name:      "image_with_trailing_space",
			builder:   sdlTestBuilder{serviceBlock: `    image: "nginx "`},
			shouldErr: false,
		},
		{
			name:      "image_with_only_spaces",
			builder:   sdlTestBuilder{serviceBlock: `    image: "   "`},
			shouldErr: false,
			reason:    "Whitespace-only strings pass schema but fail Go validation",
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

func TestSchemaValidation_CaseSensitivity(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "version_uppercase",
			sdl: `VERSION: "2.0"
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
			shouldErr: true,
			reason:    "Field names are case sensitive",
		},
		{
			name: "denom_uppercase_UAKT",
			builder: sdlTestBuilder{placementBlock: `    dc:
      pricing:
        web:
          denom: UAKT
          amount: 1`},
			shouldErr: true,
			reason:    "denom must be lowercase uakt",
		},
		{
			name: "protocol_TCP",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        proto: TCP
        to:
          - global: true`},
			shouldErr: false,
			reason:    "TCP is valid",
		},
		{
			name: "protocol_Tcp",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        proto: Tcp
        to:
          - global: true`},
			shouldErr: true,
			reason:    "Mixed case protocol invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := tt.sdl
			if sdl == "" {
				sdl = tt.builder.build()
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

func TestSchemaValidation_IPWithoutGlobal(t *testing.T) {
	sdl := sdlTestBuilder{
		endpoints: `endpoints:
  myip:
    kind: ip`,
		exposeBlock: `    expose:
      - port: 80
        to:
          - ip: myip`,
	}.build()

	err := validateInputAgainstSchema([]byte(sdl))
	require.Error(t, err, "IP without global: true should be rejected")
}

func TestSchemaValidation_StorageNameConflicts(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "duplicate_storage_names",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: data
          - size: 2Gi
            name: data`},
			shouldErr: false,
			reason:    "Schema doesn't catch duplicate names",
		},
		{
			name: "reserved_name_default",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: default`},
			shouldErr: false,
		},
		{
			name: "storage_name_with_special_chars",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: "my-data_123"`},
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

func TestSchemaValidation_ZeroValues(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "gpu_units_zero",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: 0`},
			shouldErr: false,
		},
		{
			name: "cpu_units_zero",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 0
        memory:
          size: 1Gi
        storage:
          - size: 1Gi`},
			shouldErr: true,
			reason:    "CPU cannot be zero",
		},
		{
			name: "pricing_amount_zero",
			builder: sdlTestBuilder{placementBlock: `    dc:
      pricing:
        web:
          denom: uakt
          amount: 0`},
			shouldErr: false,
		},
		{
			name: "next_tries_zero",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        to:
          - global: true
        http_options:
          next_tries: 0`},
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
