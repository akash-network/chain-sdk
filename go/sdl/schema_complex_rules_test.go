package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaValidation_NextCases_OffLogic(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "off_alone",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          next_cases: ["off"]
        to:
          - global: true`},
			shouldErr: false,
			reason:    "off can be alone",
		},
		{
			name: "off_with_other_codes",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          next_cases: ["off", "500"]
        to:
          - global: true`},
			shouldErr: true,
			reason:    "off cannot be combined with other codes",
		},
		{
			name: "off_with_error",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          next_cases: ["off", "error"]
        to:
          - global: true`},
			shouldErr: true,
			reason:    "off cannot be combined with error",
		},
		{
			name: "multiple_valid_codes_without_off",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          next_cases: ["500", "502", "503"]
        to:
          - global: true`},
			shouldErr: false,
			reason:    "multiple codes valid without off",
		},
		{
			name: "error_and_timeout",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          next_cases: ["error", "timeout"]
        to:
          - global: true`},
			shouldErr: false,
			reason:    "error and timeout are valid together",
		},
		{
			name: "empty_array",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        http_options:
          next_cases: []
        to:
          - global: true`},
			shouldErr: false,
			reason:    "empty array might be valid (use defaults)",
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

func TestSchemaValidation_SignedBy(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "valid_anyOf",
			builder: sdlTestBuilder{placementBlock: `    dc:
      attributes:
        region: us-west
      signedBy:
        anyOf:
          - akash1address1
          - akash1address2
      pricing:
        web:
          denom: uakt
          amount: 1`},
			shouldErr: false,
		},
		{
			name: "valid_allOf",
			builder: sdlTestBuilder{placementBlock: `    dc:
      attributes:
        region: us-west
      signedBy:
        allOf:
          - akash1address1
          - akash1address2
      pricing:
        web:
          denom: uakt
          amount: 1`},
			shouldErr: false,
		},
		{
			name: "both_anyOf_and_allOf",
			builder: sdlTestBuilder{placementBlock: `    dc:
      attributes:
        region: us-west
      signedBy:
        anyOf:
          - akash1address1
        allOf:
          - akash1address2
      pricing:
        web:
          denom: uakt
          amount: 1`},
			shouldErr: false,
			reason:    "both fields can coexist",
		},
		{
			name: "empty_anyOf_array",
			builder: sdlTestBuilder{placementBlock: `    dc:
      attributes:
        region: us-west
      signedBy:
        anyOf: []
      pricing:
        web:
          denom: uakt
          amount: 1`},
			shouldErr: false,
			reason:    "empty array might be valid",
		},
		{
			name: "anyOf_with_non_string",
			builder: sdlTestBuilder{placementBlock: `    dc:
      attributes:
        region: us-west
      signedBy:
        anyOf:
          - 123
      pricing:
        web:
          denom: uakt
          amount: 1`},
			shouldErr: true,
			reason:    "addresses must be strings",
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

func TestSchemaValidation_GPUAttributes(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "valid_nvidia_without_models",
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
              nvidia: []`},
			shouldErr: false,
			reason:    "empty nvidia array is wildcard",
		},
		{
			name: "nvidia_with_model",
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
			name: "nvidia_with_model_and_ram",
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
                  ram: 80Gi`},
			shouldErr: false,
		},
		{
			name: "invalid_RAM_format",
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
                  ram: 80GB`},
			shouldErr: false,
			reason:    "RAM format not validated by schema",
		},
		{
			name: "multiple_vendors",
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
              amd:
                - model: mi250`},
			shouldErr: true,
			reason:    "additionalProperties: false should reject AMD",
		},
		{
			name: "empty_nvidia_array",
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
              nvidia: []`},
			shouldErr: false,
			reason:    "empty array might default to wildcard",
		},
		{
			name: "completely_empty_vendor",
			builder: sdlTestBuilder{resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
        gpu:
          units: 1
          attributes:
            vendor: {}`},
			shouldErr: true,
			reason:    "vendor must have at least one property (minProperties: 1)",
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

func TestSchemaValidation_IPRequiresGlobal(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "IP_with_global_true",
			builder: sdlTestBuilder{
				endpoints: `endpoints:
  myip:
    kind: ip`,
				exposeBlock: `    expose:
      - port: 80
        to:
          - ip: myip
            global: true`},
			shouldErr: false,
		},
		{
			name: "IP_with_global_false",
			builder: sdlTestBuilder{
				endpoints: `endpoints:
  myip:
    kind: ip`,
				exposeBlock: `    expose:
      - port: 80
        to:
          - ip: myip
            global: false`},
			shouldErr: true,
			reason:    "IP requires global: true",
		},
		{
			name: "IP_without_global",
			builder: sdlTestBuilder{
				endpoints: `endpoints:
  myip:
    kind: ip`,
				exposeBlock: `    expose:
      - port: 80
        to:
          - ip: myip`},
			shouldErr: true,
			reason:    "IP requires global field",
		},
		{
			name: "no_IP_with_global_false",
			builder: sdlTestBuilder{exposeBlock: `    expose:
      - port: 80
        to:
          - global: false`},
			shouldErr: false,
			reason:    "global: false valid without IP",
		},
		{
			name: "empty_IP_string",
			builder: sdlTestBuilder{
				endpoints: `endpoints:
  myip:
    kind: ip`,
				exposeBlock: `    expose:
      - port: 80
        to:
          - ip: ""
            global: true`},
			shouldErr: true,
			reason:    "IP cannot be empty",
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

func TestSchemaValidation_Dependencies(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		shouldErr bool
		reason    string
	}{
		{
			name: "valid_dependency",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    dependencies:
      - service: db
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
`,
			shouldErr: false,
		},
		{
			name: "multiple_dependencies",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    dependencies:
      - service: db
      - service: cache
  db:
    image: postgres
  cache:
    image: redis
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
    cache:
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
        cache:
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
  cache:
    dc:
      profile: cache
      count: 1
`,
			shouldErr: false,
		},
		{
			name: "empty_dependencies_array",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    dependencies: []
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
`,
			shouldErr: false,
		},
		{
			name: "dependency_without_service_field",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    dependencies:
      - name: db
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
`,
			shouldErr: true,
			reason:    "only service field is allowed in dependencies",
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

func TestSchemaValidation_EndpointKind(t *testing.T) {
	tests := []struct {
		name      string
		kind      string
		shouldErr bool
		reason    string
	}{
		{
			name:      "valid_ip_kind",
			kind:      "ip",
			shouldErr: false,
		},
		{
			name:      "invalid_http_kind",
			kind:      "http",
			shouldErr: true,
			reason:    "only 'ip' kind is supported",
		},
		{
			name:      "invalid_lb_kind",
			kind:      "loadbalancer",
			shouldErr: true,
			reason:    "only 'ip' kind is supported",
		},
		{
			name:      "empty_kind",
			kind:      `""`,
			shouldErr: true,
			reason:    "kind is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := sdlTestBuilder{
				endpoints: "endpoints:\n  myendpoint:\n    kind: " + tt.kind,
			}
			err := validateInputAgainstSchema([]byte(builder.build()))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

func TestSchemaValidation_EnvArgsCommand(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "valid_env_array",
			builder: sdlTestBuilder{serviceBlock: `    image: nginx
    env:
      - NODE_ENV=production
      - PORT=3000`},
			shouldErr: false,
		},
		{
			name: "valid_args_array",
			builder: sdlTestBuilder{serviceBlock: `    image: nginx
    args:
      - --verbose
      - --config=/etc/app.conf`},
			shouldErr: false,
		},
		{
			name: "valid_command_array",
			builder: sdlTestBuilder{serviceBlock: `    image: nginx
    command:
      - /bin/sh
      - -c
      - "echo hello"`},
			shouldErr: false,
		},
		{
			name: "all_three_together",
			builder: sdlTestBuilder{serviceBlock: `    image: nginx
    env:
      - DEBUG=true
    args:
      - --port=8080
    command:
      - /app/start.sh`},
			shouldErr: false,
		},
		{
			name: "env_as_null",
			builder: sdlTestBuilder{serviceBlock: `    image: nginx
    env: null`},
			shouldErr: false,
			reason:    "null is valid for optional arrays",
		},
		{
			name: "empty_env_array",
			builder: sdlTestBuilder{serviceBlock: `    image: nginx
    env: []`},
			shouldErr: false,
			reason:    "empty array is valid",
		},
		{
			name: "env_as_object",
			builder: sdlTestBuilder{serviceBlock: `    image: nginx
    env:
      DEBUG: true`},
			shouldErr: true,
			reason:    "env must be array, not object",
		},
		{
			name: "env_items_as_numbers",
			builder: sdlTestBuilder{serviceBlock: `    image: nginx
    env:
      - 123`},
			shouldErr: true,
			reason:    "env items must be strings",
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

func TestSchemaValidation_ExposeToService(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		shouldErr bool
		reason    string
	}{
		{
			name: "valid_service_routing",
			sdl: `version: "2.0"
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
`,
			shouldErr: false,
		},
		{
			name: "multiple_to_targets",
			sdl: `version: "2.0"
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
`,
			shouldErr: false,
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

func TestSchemaValidation_StorageParams(t *testing.T) {
	tests := []struct {
		name      string
		builder   sdlTestBuilder
		shouldErr bool
		reason    string
	}{
		{
			name: "valid_mount_and_readOnly",
			builder: sdlTestBuilder{
				serviceBlock: `    image: nginx
    params:
      storage:
        data:
          mount: /data
          readOnly: false`,
				resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: data
          - size: 1Gi
            name: data1
          - size: 1Gi
            name: data2`},
			shouldErr: false,
		},
		{
			name: "readOnly_true",
			builder: sdlTestBuilder{
				serviceBlock: `    image: nginx
    params:
      storage:
        data:
          mount: /data
          readOnly: true`,
				resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: data
          - size: 1Gi
            name: data1
          - size: 1Gi
            name: data2`},
			shouldErr: false,
		},
		{
			name: "only_mount_no_readOnly",
			builder: sdlTestBuilder{
				serviceBlock: `    image: nginx
    params:
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
          - size: 1Gi
            name: data1
          - size: 1Gi
            name: data2`},
			shouldErr: false,
			reason:    "readOnly is optional",
		},
		{
			name: "mount_without_storage_definition",
			builder: sdlTestBuilder{
				serviceBlock: `    image: nginx
    params:
      storage:
        nonexistent:
          mount: /data`,
				resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: data
          - size: 1Gi
            name: data1
          - size: 1Gi
            name: data2`},
			shouldErr: false,
			reason:    "Schema validates structure only, not references",
		},
		{
			name: "multiple_storage_params",
			builder: sdlTestBuilder{
				serviceBlock: `    image: nginx
    params:
      storage:
        data1:
          mount: /data1
        data2:
          mount: /data2`,
				resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: data
          - size: 1Gi
            name: data1
          - size: 1Gi
            name: data2`},
			shouldErr: false,
		},
		{
			name: "readOnly_as_string",
			builder: sdlTestBuilder{
				serviceBlock: `    image: nginx
    params:
      storage:
        data:
          mount: /data
          readOnly: "true"`,
				resourcesBlock: `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi
            name: data
          - size: 1Gi
            name: data1
          - size: 1Gi
            name: data2`},
			shouldErr: true,
			reason:    "readOnly must be boolean",
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
