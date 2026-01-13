package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSchemaValidation_AdditionalProperties tests that additionalProperties: false
// correctly rejects unknown fields in various structures
func TestSchemaValidation_AdditionalProperties(t *testing.T) {
	tests := []struct {
		name      string
		sdl       string
		shouldErr bool
		reason    string
	}{
		{
			name: "unknown field in service",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    unknown_field: value
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
`,
			shouldErr: true,
			reason:    "service should not allow unknown fields",
		},
		{
			name: "unknown field in credentials",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    credentials:
      host: docker.io
      username: user123
      password: secret123
      unknown_field: value
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
`,
			shouldErr: true,
			reason:    "credentials should not allow unknown fields",
		},
		{
			name: "unknown field in dependencies item",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    dependencies:
      - service: db
        unknown_field: value
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
			reason:    "dependencies items should not allow unknown fields",
		},
		{
			name: "unknown field in expose item",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        unknown_field: value
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
`,
			shouldErr: true,
			reason:    "expose items should not allow unknown fields",
		},
		{
			name: "unknown field in expose.to item",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - global: true
            unknown_field: value
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
`,
			shouldErr: true,
			reason:    "expose.to items should not allow unknown fields",
		},
		{
			name: "unknown field in http_options",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        http_options:
          max_body_size: 1048576
          unknown_field: value
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
`,
			shouldErr: true,
			reason:    "http_options should not allow unknown fields",
		},
		{
			name: "unknown field in cpu",
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
          unknown_field: value
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
`,
			shouldErr: true,
			reason:    "cpu should not allow unknown fields",
		},
		{
			name: "unknown field in memory",
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
          unknown_field: value
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
`,
			shouldErr: true,
			reason:    "memory should not allow unknown fields",
		},
		{
			name: "unknown field in gpu",
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
          - size: 1Gi
        gpu:
          units: 1
          unknown_field: value
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
`,
			shouldErr: true,
			reason:    "gpu should not allow unknown fields",
		},
		{
			name: "unknown field in gpu.attributes",
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
          - size: 1Gi
        gpu:
          units: 1
          attributes:
            unknown_field: value
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
`,
			shouldErr: true,
			reason:    "gpu.attributes should only allow vendor",
		},
		{
			name: "unknown field in nvidia gpu item",
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
          - size: 1Gi
        gpu:
          units: 1
          attributes:
            vendor:
              nvidia:
                - model: a100
                  unknown_field: value
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
`,
			shouldErr: true,
			reason:    "nvidia gpu items should not allow unknown fields",
		},
		{
			name: "unknown field in storage item",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    params:
      storage:
        data:
          mount: /data
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
            unknown_field: value
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
`,
			shouldErr: true,
			reason:    "storage items should not allow unknown fields",
		},
		{
			name: "unknown field in storage.attributes",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    params:
      storage:
        data:
          mount: /data
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
            attributes:
              persistent: true
              class: beta1
              unknown_field: value
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
`,
			shouldErr: true,
			reason:    "storage.attributes should not allow unknown fields",
		},
		{
			name: "unknown field in params.storage item",
			sdl: `version: "2.0"
services:
  web:
    image: nginx
    params:
      storage:
        data:
          mount: /data
          unknown_field: value
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
`,
			shouldErr: true,
			reason:    "params.storage items should not allow unknown fields",
		},
		{
			name: "unknown field in endpoint",
			sdl: `version: "2.0"
endpoints:
  myip:
    kind: ip
    unknown_field: value
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - ip: myip
            global: true
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
`,
			shouldErr: true,
			reason:    "endpoints should not allow unknown fields (already has additionalProperties: false)",
		},
		{
			name: "unknown field in placement",
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
          - size: 1Gi
  placement:
    dc:
      unknown_field: value
      pricing:
        web:
          denom: uakt
          amount: 1
deployment:
  web:
    dc:
      profile: web
      count: 1
`,
			shouldErr: true,
			reason:    "placement items should not allow unknown fields",
		},
		{
			name: "unknown field in pricing item",
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
          - size: 1Gi
  placement:
    dc:
      pricing:
        web:
          denom: uakt
          amount: 1
          unknown_field: value
deployment:
  web:
    dc:
      profile: web
      count: 1
`,
			shouldErr: true,
			reason:    "pricing items should not allow unknown fields",
		},
		{
			name: "unknown field in deployment item",
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
      unknown_field: value
`,
			shouldErr: true,
			reason:    "deployment items should not allow unknown fields",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputAgainstSchema([]byte(tt.sdl))
			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.reason)
				require.Contains(t, err.Error(), "Additional property", 
					"Error should mention Additional property")
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.reason)
			}
		})
	}
}

