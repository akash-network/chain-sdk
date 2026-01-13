package sdl

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/types/attributes/v1"
	"pkg.akt.dev/go/node/types/unit"
)

func TestStorage_Parse(t *testing.T) {
	tests := []struct {
		name      string
		yaml      string
		shouldErr bool
		checkFunc func(*testing.T, v2ResourceStorageArray)
	}{
		{
			name: "legacy format",
			yaml: `size: 1Gi`,
			checkFunc: func(t *testing.T, p v2ResourceStorageArray) {
				require.Len(t, p, 1)
				require.Equal(t, byteQuantity(1*unit.Gi), p[0].Quantity)
				require.Len(t, p[0].Attributes, 0)
			},
		},
		{
			name: "array single element",
			yaml: `- size: 1Gi`,
			checkFunc: func(t *testing.T, p v2ResourceStorageArray) {
				require.Len(t, p, 1)
				require.Equal(t, byteQuantity(1*unit.Gi), p[0].Quantity)
				require.Len(t, p[0].Attributes, 0)
			},
		},
		{
			name: "persistent with class",
			yaml: `- size: 1Gi
  attributes:
    persistent: true
    class: default`,
			checkFunc: func(t *testing.T, p v2ResourceStorageArray) {
				require.Len(t, p, 1)
				require.Equal(t, byteQuantity(1*unit.Gi), p[0].Quantity)
				require.Len(t, p[0].Attributes, 2)
				attr := types.Attributes(p[0].Attributes)
				require.Equal(t, "class", attr[0].Key)
				require.Equal(t, "default", attr[0].Value)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p v2ResourceStorageArray
			err := yaml.Unmarshal([]byte(tt.yaml), &p)

			if tt.shouldErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			if tt.checkFunc != nil {
				tt.checkFunc(t, p)
			}
		})
	}
}

func TestStorage_Attributes(t *testing.T) {
	tests := []struct {
		name      string
		yaml      string
		shouldErr bool
		errType   error
		checkFunc func(*testing.T, v2ResourceStorageArray)
	}{
		{
			name: "unknown attribute",
			yaml: `- size: 1Gi
  attributes:
    somefield: foo`,
			shouldErr: true,
			errType:   errUnsupportedStorageAttribute,
		},
		{
			name: "multiple unnamed ephemeral",
			yaml: `- size: 1Gi
- size: 2Gi`,
			shouldErr: true,
			errType:   errStorageDuplicatedVolumeName,
		},
		{
			name: "ephemeral no class",
			yaml: `- size: 1Gi`,
		},
		{
			name: "ephemeral with class",
			yaml: `- size: 1Gi
  attributes:
    class: foo`,
			shouldErr: true,
			errType:   errStorageEphemeralClass,
		},
		{
			name: "beta1 class with ephemeral",
			yaml: `- size: 1Gi
  attributes:
    class: beta1
    persistent: false`,
			shouldErr: true,
			errType:   errStorageEphemeralClass,
		},
		{
			name: "persistent default class",
			yaml: `- size: 1Gi
  attributes:
    persistent: true`,
			checkFunc: func(t *testing.T, p v2ResourceStorageArray) {
				require.Len(t, p[0].Attributes, 2)
				require.Equal(t, "class", p[0].Attributes[0].Key)
				require.Equal(t, "default", p[0].Attributes[0].Value)
			},
		},
		{
			name: "persistent custom class",
			yaml: `- size: 1Gi
  attributes:
    persistent: true
    class: beta1`,
			checkFunc: func(t *testing.T, p v2ResourceStorageArray) {
				require.Len(t, p[0].Attributes, 2)
				require.Equal(t, "class", p[0].Attributes[0].Key)
				require.Equal(t, "beta1", p[0].Attributes[0].Value)
			},
		},
		{
			name: "RAM class valid",
			yaml: `- size: 1Gi
  attributes:
    persistent: false
    class: ram`,
			checkFunc: func(t *testing.T, p v2ResourceStorageArray) {
				require.Len(t, p[0].Attributes, 2)
				require.Equal(t, "class", p[0].Attributes[0].Key)
				require.Equal(t, "ram", p[0].Attributes[0].Value)
			},
		},
		{
			name: "RAM class with persistent true",
			yaml: `- size: 1Gi
  attributes:
    persistent: true
    class: ram`,
			shouldErr: true,
			errType:   errStorageRAMClass,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p v2ResourceStorageArray
			err := yaml.Unmarshal([]byte(tt.yaml), &p)

			if tt.shouldErr {
				require.Error(t, err)
				if tt.errType != nil {
					require.ErrorIs(t, err, tt.errType)
				}
				return
			}

			require.NoError(t, err)
			if tt.checkFunc != nil {
				tt.checkFunc(t, p)
			}
		})
	}
}

func TestStorage_StableSort(t *testing.T) {
	storage := v2ResourceStorageArray{
		{
			Quantity: 2 * unit.Gi,
			Attributes: v2StorageAttributes{
				types.Attribute{
					Key:   "persistent",
					Value: "true",
				},
			},
		},
		{
			Quantity: 1 * unit.Gi,
		},
		{
			Quantity: 10 * unit.Gi,
		},
	}

	storage.sort()

	require.Equal(t, byteQuantity(1*unit.Gi), storage[0].Quantity)
	require.Equal(t, byteQuantity(2*unit.Gi), storage[1].Quantity)
	require.Equal(t, byteQuantity(10*unit.Gi), storage[2].Quantity)
}

func TestStorage_Invalid_InvalidMount(t *testing.T) {
	_, err := ReadFile("./_testdata/storageClass1.yaml")
	require.Error(t, err)
	require.Contains(t, err.Error(), "expected absolute path")
}

func TestStorage_Invalid_MountNotAbsolute(t *testing.T) {
	_, err := ReadFile("./_testdata/storageClass2.yaml")
	require.Error(t, err)
	require.Contains(t, err.Error(), "expected absolute path")
}

func TestStorage_Invalid_VolumeReference(t *testing.T) {
	_, err := ReadFile("./_testdata/storageClass3.yaml")
	require.Error(t, err)
	require.Contains(t, err.Error(), "references to no-existing compute volume")
}

func TestStorage_Invalid_DuplicatedMount(t *testing.T) {
	_, err := ReadFile("./_testdata/storageClass4.yaml")
	require.Error(t, err)
	require.Contains(t, err.Error(), "already in use by volume")
}

func TestStorage_Invalid_NoMount(t *testing.T) {
	_, err := ReadFile("./_testdata/storageClass5.yaml")
	require.Error(t, err)
	require.Contains(t, err.Error(), "to have mount")
}

func TestStorage_SchemaValidation_Mount(t *testing.T) {
	const sdlTemplate = `version: "2.0"
services:
  web:
    image: nginx
    params:
      storage:
        data:
          mount: %s
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
`

	tests := []struct {
		name      string
		mount     string
		shouldErr bool
	}{
		{"absolute path valid", "/data", false},
		{"absolute path with subdirs", "/var/lib/data", false},
		{"relative path invalid", "data", true},
		{"relative path with slash", "data/path", true},
		{"empty string invalid", `""`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl := fmt.Sprintf(sdlTemplate, tt.mount)
			err := validateInputAgainstSchema([]byte(sdl))

			if tt.shouldErr {
				require.Error(t, err, "Schema should reject: %s", tt.mount)
			} else {
				require.NoError(t, err, "Schema should accept: %s", tt.mount)
			}
		})
	}
}

func TestStorage_SchemaValidation_Classes(t *testing.T) {
	tests := []struct {
		name        string
		attributes  string
		shouldErr   bool
		description string
	}{
		{
			name: "RAM with persistent false",
			attributes: `class: ram
              persistent: false`,
			shouldErr:   false,
			description: "RAM class with persistent=false should be valid",
		},
		{
			name: "RAM with persistent true",
			attributes: `class: ram
              persistent: true`,
			shouldErr:   true,
			description: "RAM class with persistent=true should be invalid",
		},
		{
			name: "beta1 with persistent true",
			attributes: `class: beta1
              persistent: true`,
			shouldErr:   false,
			description: "Non-RAM class with persistent=true should be valid",
		},
		{
			name: "beta1 with persistent false",
			attributes: `class: beta1
              persistent: false`,
			shouldErr:   true,
			description: "Non-RAM class with persistent=false should be invalid",
		},
		{
			name: "default with persistent true",
			attributes: `class: default
              persistent: true`,
			shouldErr:   false,
			description: "Default class with persistent=true should be valid",
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
            name: data
            attributes:
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
`, tt.attributes)

			err := validateInputAgainstSchema([]byte(sdl))

			if tt.shouldErr {
				require.Error(t, err, tt.description)
			} else {
				require.NoError(t, err, tt.description)
			}
		})
	}
}
