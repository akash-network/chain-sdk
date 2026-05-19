package sdl

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
	verificationv1 "pkg.akt.dev/go/node/verification/v1"
)

// v2Verification is the SDL verification block.
type v2Verification struct {
	MinTier         int32    `yaml:"min_tier"`
	Capabilities    []string `yaml:"capabilities,omitempty"`
	Auditors        []string `yaml:"auditors,omitempty"`
	AuditorMode     string   `yaml:"auditor_mode,omitempty"`
	MinAuditorCount uint32   `yaml:"min_auditor_count,omitempty"`
}

const auditorBech32Prefix = "akash1"

func (v *v2Verification) UnmarshalYAML(node *yaml.Node) error {
	type raw v2Verification
	var r raw
	if err := node.Decode(&r); err != nil {
		return err
	}

	if r.MinTier < 0 || r.MinTier > 4 {
		return fmt.Errorf("%w: verification.min_tier must be in [0,4], got %d", errSDLInvalid, r.MinTier)
	}

	switch r.AuditorMode {
	case "", "any", "all":
		// valid
	default:
		return fmt.Errorf("%w: verification.auditor_mode must be \"any\" or \"all\" (got %q)", errSDLInvalid, r.AuditorMode)
	}

	for _, c := range r.Capabilities {
		if _, ok := verificationv1.IsKnownCapabilityName(c); !ok {
			return fmt.Errorf("%w: verification.capabilities: unknown capability %q", errSDLInvalid, c)
		}
	}

	for i, a := range r.Auditors {
		trimmed := strings.TrimSpace(a)
		if trimmed == "" {
			return fmt.Errorf("%w: verification.auditors[%d] is empty", errSDLInvalid, i)
		}
		if trimmed != a {
			return fmt.Errorf("%w: verification.auditors[%d] has surrounding whitespace", errSDLInvalid, i)
		}
		if !strings.HasPrefix(trimmed, auditorBech32Prefix) {
			return fmt.Errorf(
				"%w: verification.auditors[%d] %q does not have expected bech32 prefix %q",
				errSDLInvalid, i, trimmed, auditorBech32Prefix,
			)
		}
	}

	if r.MinTier == 0 && (len(r.Capabilities) != 0 || len(r.Auditors) != 0 || r.MinAuditorCount != 0) {
		return fmt.Errorf(
			"%w: verification.min_tier must be greater than 0 when capabilities, auditors, or min_auditor_count are set",
			errSDLInvalid,
		)
	}

	*v = v2Verification(r)
	return nil
}

func (v *v2Verification) isVacuous() bool {
	return v.MinTier == 0 &&
		len(v.Capabilities) == 0 &&
		len(v.Auditors) == 0 &&
		v.MinAuditorCount == 0
}

// toProto converts the SDL verification block to the chain type.
func (v *v2Verification) toProto() *verificationv1.VerificationRequirement {
	if v == nil || v.isVacuous() {
		return nil
	}

	caps := make([]verificationv1.CapabilityFlag, 0, len(v.Capabilities))
	for _, c := range v.Capabilities {
		if flag, ok := verificationv1.IsKnownCapabilityName(c); ok {
			caps = append(caps, flag)
		}
	}

	var mode verificationv1.AuditorSelectionMode
	switch v.AuditorMode {
	case "", "any":
		mode = verificationv1.AuditorSelectionModeAny
	case "all":
		mode = verificationv1.AuditorSelectionModeAll
	default:
		mode = verificationv1.AuditorSelectionModeAny
	}

	auditors := append([]string(nil), v.Auditors...)

	return &verificationv1.VerificationRequirement{
		MinTier:              verificationv1.VerificationTier(v.MinTier),
		RequiredCapabilities: caps,
		RequiredAuditors:     auditors,
		AuditorMode:          mode,
		MinAuditorCount:      v.MinAuditorCount,
	}
}
