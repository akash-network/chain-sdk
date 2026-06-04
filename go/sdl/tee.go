package sdl

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
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
	errTEETypeMismatch    = errors.New("sdl: conflicting tee types in the same placement group")
)

// v2TEEParams defines the TEE (Trusted Execution Environment) configuration
// as a service parameter. Uses *bool for Attestation so we can distinguish
// "not set" (nil -> default true) from "explicitly false". The group builder
// always sets the manifest field explicitly because proto3 bool defaults to
// false, which would silently disable attestation for non-Go producers.
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

// IsGPURequiredForTEEType returns true if the TEE type requires GPU confidential compute.
func IsGPURequiredForTEEType(t string) bool {
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

	if IsGPURequiredForTEEType(tee.Type) && !hasGPU {
		return fmt.Errorf("%w: %q specified but no gpu resources defined", errTEEGPURequired, tee.Type)
	}

	return nil
}
