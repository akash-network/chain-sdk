package sdl

import (
	"errors"
	"fmt"
	"sort"

	"gopkg.in/yaml.v3"

	types "pkg.akt.dev/go/node/types/attributes/v1"
)

const (
	TEEAttributeType        = "type"
	TEEAttributeAttestation = "attestation"
)

// Supported TEE types, each maps 1:1 to a Kata Containers runtime class.
// The -gpu variants indicate GPU confidential compute (VFIO passthrough, CC-on mode).
const (
	TEETypeSEVSNP    = "sev-snp"     // CPU-only AMD SEV-SNP -> kata-qemu-snp
	TEETypeSEVSNPGPU = "sev-snp-gpu" // CPU + GPU CC -> kata-qemu-nvidia-gpu-snp
	TEETypeTDX       = "tdx"         // CPU-only Intel TDX -> kata-qemu-tdx
	TEETypeTDXGPU    = "tdx-gpu"     // CPU + GPU CC -> kata-qemu-nvidia-gpu-tdx
)

var allowedTEETypes = map[string]bool{
	TEETypeSEVSNP:    true,
	TEETypeSEVSNPGPU: true,
	TEETypeTDX:       true,
	TEETypeTDXGPU:    true,
}

var (
	errTEETypeRequired    = errors.New("sdl: tee type is required")
	errTEETypeUnsupported = errors.New("sdl: unsupported tee type")
	errTEEGPURequired     = errors.New("sdl: tee type requires gpu resources")
)

type v2TEEAttributes types.Attributes

// v2TEEParams defines the TEE (Trusted Execution Environment) configuration
// for a compute resource. Follows the same pattern as v2ResourceStorage.
type v2TEEParams struct {
	Type        string `yaml:"type"`
	Attestation *bool  `yaml:"attestation,omitempty"`
}

func (tee *v2TEEParams) UnmarshalYAML(node *yaml.Node) error {
	type raw v2TEEParams
	var r raw
	if err := node.Decode(&r); err != nil {
		return err
	}
	*tee = v2TEEParams(r)
	return tee.validate()
}

func (tee *v2TEEParams) toAttributes() v2TEEAttributes {
	if tee == nil {
		return nil
	}

	attrs := make(v2TEEAttributes, 0, 2)
	attrs = append(attrs, types.Attribute{
		Key:   TEEAttributeType,
		Value: tee.Type,
	})

	// attestation defaults to true (provider injects sidecar)
	attestation := "true"
	if tee.Attestation != nil && !*tee.Attestation {
		attestation = "false"
	}
	attrs = append(attrs, types.Attribute{
		Key:   TEEAttributeAttestation,
		Value: attestation,
	})

	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Key < attrs[j].Key
	})

	return attrs
}

// IsGPUTEEType returns true if the TEE type requires GPU confidential compute.
func IsGPUTEEType(t string) bool {
	return t == TEETypeSEVSNPGPU || t == TEETypeTDXGPU
}

func (tee *v2TEEParams) validate() error {
	if tee == nil {
		return nil
	}

	if tee.Type == "" {
		return errTEETypeRequired
	}

	if !allowedTEETypes[tee.Type] {
		return fmt.Errorf("%w: %q (allowed: %s, %s, %s, %s)",
			errTEETypeUnsupported, tee.Type,
			TEETypeSEVSNP, TEETypeSEVSNPGPU, TEETypeTDX, TEETypeTDXGPU)
	}

	return nil
}

// validateWithGPU checks TEE + GPU consistency. Called during group building
// when GPU info is available.
func (tee *v2TEEParams) validateWithGPU(hasGPU bool) error {
	if tee == nil {
		return nil
	}

	if IsGPUTEEType(tee.Type) && !hasGPU {
		return fmt.Errorf("%w: %q specified but no gpu resources defined", errTEEGPURequired, tee.Type)
	}

	return nil
}

func (attr *v2TEEAttributes) UnmarshalYAML(node *yaml.Node) error {
	var tee v2TEEParams

	if err := node.Decode(&tee); err != nil {
		return err
	}

	if err := tee.validate(); err != nil {
		return err
	}

	*attr = tee.toAttributes()

	return nil
}
