package sdl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSchemaValidation_NextCases_OffLogic tests complex next_cases validation
// Rule: "off" can only appear alone, not with other error codes
func TestSchemaValidation_NextCases_OffLogic(t *testing.T) {
	tests := []struct {
		name      string
		nextCases string
		shouldErr bool
		reason    string
	}{
		{
			name:      "off alone",
			nextCases: `["off"]`,
			shouldErr: false,
			reason:    "off can be alone",
		},
		{
			name:      "off with other codes",
			nextCases: `["off", "500"]`,
			shouldErr: true,
			reason:    "off cannot be combined with other codes",
		},
		{
			name:      "off with error",
			nextCases: `["off", "error"]`,
			shouldErr: true,
			reason:    "off cannot be combined with error",
		},
		{
			name:      "multiple valid codes without off",
			nextCases: `["500", "502", "503"]`,
			shouldErr: false,
			reason:    "multiple codes valid without off",
		},
		{
			name:      "error and timeout",
			nextCases: `["error", "timeout"]`,
			shouldErr: false,
			reason:    "error and timeout are valid together",
		},
		{
			name:      "empty array",
			nextCases: `[]`,
			shouldErr: false,
			reason:    "empty array might be valid (use defaults)",
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
        http_options:
          next_cases: %s
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
`, tt.nextCases)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_SignedBy tests signedBy allOf/anyOf structure
func TestSchemaValidation_SignedBy(t *testing.T) {
	tests := []struct {
		name      string
		signedBy  string
		shouldErr bool
		reason    string
	}{
		{
			name: "valid anyOf",
			signedBy: `signedBy:
        anyOf:
          - akash1address1
          - akash1address2`,
			shouldErr: false,
		},
		{
			name: "valid allOf",
			signedBy: `signedBy:
        allOf:
          - akash1address1
          - akash1address2`,
			shouldErr: false,
		},
		{
			name: "both anyOf and allOf",
			signedBy: `signedBy:
        anyOf:
          - akash1address1
        allOf:
          - akash1address2`,
			shouldErr: false,
			reason:    "both fields can coexist",
		},
		{
			name: "empty anyOf array",
			signedBy: `signedBy:
        anyOf: []`,
			shouldErr: false,
			reason:    "empty array might be valid",
		},
		{
			name: "anyOf with non-string",
			signedBy: `signedBy:
        anyOf:
          - 123`,
			shouldErr: true,
			reason:    "addresses must be strings",
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
  placement:
    dc:
      attributes:
        region: us-west
      %s
      pricing:
        web:
          denom: uakt
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: 1
`, tt.signedBy)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_GPUAttributes tests GPU vendor/model/ram structure
func TestSchemaValidation_GPUAttributes(t *testing.T) {
	tests := []struct {
		name      string
		gpu       string
		shouldErr bool
		reason    string
	}{
		{
			name: "valid nvidia without models",
			gpu: `gpu:
          units: 1
          attributes:
            vendor:
              nvidia: []`,
			shouldErr: false,
			reason:    "empty nvidia array is wildcard",
		},
		{
			name: "nvidia with model",
			gpu: `gpu:
          units: 1
          attributes:
            vendor:
              nvidia:
                - model: a100`,
			shouldErr: false,
		},
		{
			name: "nvidia with model and ram",
			gpu: `gpu:
          units: 1
          attributes:
            vendor:
              nvidia:
                - model: a100
                  ram: 80Gi`,
			shouldErr: false,
		},
		{
			name: "invalid RAM format",
			gpu: `gpu:
          units: 1
          attributes:
            vendor:
              nvidia:
                - model: a100
                  ram: 80GB`,
			shouldErr: false, // Schema doesn't validate RAM format (Go does)
			reason:    "RAM format not validated by schema",
		},
		{
			name: "multiple vendors",
			gpu: `gpu:
          units: 1
          attributes:
            vendor:
              nvidia:
                - model: a100
              amd:
                - model: mi250`,
			shouldErr: true,
			reason:    "additionalProperties: false should reject AMD",
		},
		{
			name: "empty nvidia array",
			gpu: `gpu:
          units: 1
          attributes:
            vendor:
              nvidia: []`,
			shouldErr: false,
			reason:    "empty array might default to wildcard",
		},
		{
			name: "completely empty vendor",
			gpu: `gpu:
          units: 1
          attributes:
            vendor: {}`,
			shouldErr: true,
			reason:    "vendor must have at least one property (minProperties: 1)",
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

// TestSchemaValidation_IPRequiresGlobal tests that IP endpoints require global: true
func TestSchemaValidation_IPRequiresGlobal(t *testing.T) {
	tests := []struct {
		name      string
		expose    string
		shouldErr bool
		reason    string
	}{
		{
			name: "IP with global true",
			expose: `- port: 80
        to:
          - ip: myip
            global: true`,
			shouldErr: false,
		},
		{
			name: "IP with global false",
			expose: `- port: 80
        to:
          - ip: myip
            global: false`,
			shouldErr: true,
			reason:    "IP requires global: true",
		},
		{
			name: "IP without global",
			expose: `- port: 80
        to:
          - ip: myip`,
			shouldErr: true,
			reason:    "IP requires global field",
		},
		{
			name: "no IP with global false",
			expose: `- port: 80
        to:
          - global: false`,
			shouldErr: false,
			reason:    "global: false valid without IP",
		},
		{
			name: "empty IP string",
			expose: `- port: 80
        to:
          - ip: ""
            global: true`,
			shouldErr: true,
			reason:    "IP cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
endpoints:
  myip:
    kind: ip
services:
  web:
    image: nginx
    expose:
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
`, tt.expose)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_Dependencies tests dependencies structure
func TestSchemaValidation_Dependencies(t *testing.T) {
	tests := []struct {
		name         string
		dependencies string
		shouldErr    bool
		reason       string
	}{
		{
			name: "valid dependency",
			dependencies: `dependencies:
      - service: db`,
			shouldErr: false,
		},
		{
			name: "multiple dependencies",
			dependencies: `dependencies:
      - service: db
      - service: cache`,
			shouldErr: false,
		},
		{
			name:         "empty dependencies array",
			dependencies: `dependencies: []`,
			shouldErr:    false,
		},
		{
			name: "dependency without service field",
			dependencies: `dependencies:
      - name: db`,
			shouldErr: true, // Schema has additionalProperties: false
			reason:    "only service field is allowed in dependencies",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
    %s
  db:
    image: postgres
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
    db:
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
        db:
          denom: uakt
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: 1
  db:
    dc:
      profile: db
      count: 1
`, tt.dependencies)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_EndpointKind tests endpoint kind validation
func TestSchemaValidation_EndpointKind(t *testing.T) {
	tests := []struct {
		name      string
		kind      string
		shouldErr bool
		reason    string
	}{
		{
			name:      "valid ip kind",
			kind:      "ip",
			shouldErr: false,
		},
		{
			name:      "invalid http kind",
			kind:      "http",
			shouldErr: true,
			reason:    "only 'ip' kind is supported",
		},
		{
			name:      "invalid lb kind",
			kind:      "loadbalancer",
			shouldErr: true,
			reason:    "only 'ip' kind is supported",
		},
		{
			name:      "empty kind",
			kind:      `""`,
			shouldErr: true,
			reason:    "kind is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
endpoints:
  myendpoint:
    kind: %s
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
`, tt.kind)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_EnvArgsCommand tests env, args, command arrays
func TestSchemaValidation_EnvArgsCommand(t *testing.T) {
	tests := []struct {
		name      string
		service   string
		shouldErr bool
		reason    string
	}{
		{
			name: "valid env array",
			service: `image: nginx
    env:
      - NODE_ENV=production
      - PORT=3000`,
			shouldErr: false,
		},
		{
			name: "valid args array",
			service: `image: nginx
    args:
      - --verbose
      - --config=/etc/app.conf`,
			shouldErr: false,
		},
		{
			name: "valid command array",
			service: `image: nginx
    command:
      - /bin/sh
      - -c
      - "echo hello"`,
			shouldErr: false,
		},
		{
			name: "all three together",
			service: `image: nginx
    env:
      - DEBUG=true
    args:
      - --port=8080
    command:
      - /app/start.sh`,
			shouldErr: false,
		},
		{
			name: "env as null",
			service: `image: nginx
    env: null`,
			shouldErr: false,
			reason:    "null is valid for optional arrays",
		},
		{
			name: "empty env array",
			service: `image: nginx
    env: []`,
			shouldErr: false,
			reason:    "empty array is valid",
		},
		{
			name: "env as object",
			service: `image: nginx
    env:
      DEBUG: true`,
			shouldErr: true,
			reason:    "env must be array, not object",
		},
		{
			name: "env items as numbers",
			service: `image: nginx
    env:
      - 123`,
			shouldErr: true,
			reason:    "env items must be strings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
services:
  web:
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
`, tt.service)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

// TestSchemaValidation_ExposeToService tests expose.to.service for internal routing
func TestSchemaValidation_ExposeToService(t *testing.T) {
	tests := []struct {
		name      string
		shouldErr bool
		reason    string
	}{
		{
			name:      "valid service routing",
			shouldErr: false,
		},
		{
			name:      "multiple to targets",
			shouldErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sdl string
			if tt.name == "valid service routing" {
				sdl = `version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - service: api
  api:
    image: node:18
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
    api:
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
        api:
          denom: uakt
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: 1
  api:
    dc:
      profile: api
      count: 1
`
			} else {
				sdl = `version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - service: api
          - global: true
  api:
    image: node:18
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
    api:
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
        api:
          denom: uakt
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: 1
  api:
    dc:
      profile: api
      count: 1
`
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

// TestSchemaValidation_StorageParams tests storage params validation
func TestSchemaValidation_StorageParams(t *testing.T) {
	tests := []struct {
		name      string
		params    string
		shouldErr bool
		reason    string
	}{
		{
			name: "valid mount and readOnly",
			params: `params:
      storage:
        data:
          mount: /data
          readOnly: false`,
			shouldErr: false,
		},
		{
			name: "readOnly true",
			params: `params:
      storage:
        data:
          mount: /data
          readOnly: true`,
			shouldErr: false,
		},
		{
			name: "only mount no readOnly",
			params: `params:
      storage:
        data:
          mount: /data`,
			shouldErr: false,
			reason:    "readOnly is optional",
		},
		{
			name: "mount without storage definition",
			params: `params:
      storage:
        nonexistent:
          mount: /data`,
			shouldErr: false, // Schema doesn't validate cross-references
			reason:    "Schema validates structure only, not references",
		},
		{
			name: "multiple storage params",
			params: `params:
      storage:
        data1:
          mount: /data1
        data2:
          mount: /data2`,
			shouldErr: false,
		},
		{
			name: "readOnly as string",
			params: `params:
      storage:
        data:
          mount: /data
          readOnly: "true"`,
			shouldErr: true,
			reason:    "readOnly must be boolean",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(`version: "2.0"
services:
  web:
    image: nginx
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
            name: data
          - size: 1Gi
            name: data1
          - size: 1Gi
            name: data2
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
`, tt.params)

			err := validateInputAgainstSchema([]byte(sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}
