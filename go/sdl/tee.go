package sdl

import (
	"errors"
	"fmt"
)

// Supported TEE types. The provider determines the actual TEE technology
// (AMD SEV-SNP or Intel TDX) at deployment time based on node capabilities.
const (
	TEETypeCPU    = "cpu"     // CPU-only confidential compute
	TEETypeCPUGPU = "cpu-gpu" // CPU + GPU confidential compute (VFIO passthrough, CC-on mode)
)

var allowedTEETypes = map[string]bool{
	TEETypeCPU:    true,
	TEETypeCPUGPU: true,
}

var (
	errTEETypeRequired    = errors.New("sdl: tee type is required")
	errTEETypeUnsupported = errors.New("sdl: unsupported tee type")
	errTEEGPURequired     = errors.New("sdl: tee type requires gpu resources")
	errTEETypeMismatch    = errors.New("sdl: conflicting tee types in the same placement group")
)

// IsGPURequiredForTEEType returns true if the TEE type requires GPU confidential compute.
func IsGPURequiredForTEEType(t string) bool {
	return t == TEETypeCPUGPU
}

// parseTEEParam validates a raw tee string value.
func parseTEEParam(val string) (string, error) {
	if val == "" {
		return "", errTEETypeRequired
	}

	if !allowedTEETypes[val] {
		return "", fmt.Errorf("%w: %q (allowed: %s, %s)",
			errTEETypeUnsupported, val, TEETypeCPU, TEETypeCPUGPU)
	}

	return val, nil
}

// validateTEEWithGPU checks TEE + GPU consistency. Called during group building
// when GPU info is available.
func validateTEEWithGPU(teeType string, hasGPU bool) error {
	if teeType == "" {
		return nil
	}

	if IsGPURequiredForTEEType(teeType) && !hasGPU {
		return fmt.Errorf("%w: %q specified but no gpu resources defined", errTEEGPURequired, teeType)
	}

	return nil
}
