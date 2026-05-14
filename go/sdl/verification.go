package sdl

import (
	"fmt"
	"strings"

	verificationv1 "pkg.akt.dev/go/node/verification/v1"
)

// v2Verification is the SDL representation of an
// `akash.verification.v1.VerificationRequirement`. It mirrors the YAML block
// described in AEP-86 README §SDL Syntax. The struct is decoded directly from
// the user's SDL and then converted to the on-chain proto via `toProto()`.
//
// There are two equivalent ways for an SDL to express "no verification
// filtering": omit the entire block, or write a block with `min_tier: 0` (or
// any other vacuous combination). `toProto()` collapses both to nil so the
// chain only ever sees one canonical shape for the no-filtering case.
type v2Verification struct {
	// MinTier is the minimum verification tier required of bidding providers.
	// Valid values are 0..4 inclusive. A value of 0 (TierUnspecified) combined
	// with no other filtering fields collapses the entire block to nil in
	// `toProto()` so the wire form matches an omitted block.
	MinTier int `yaml:"min_tier"`
	// Capabilities is the list of optional provider capability flags the
	// tenant requires bidders to assert. The SDL strings are the lowercase
	// names defined in `verification/v1.capabilities.go` (e.g.
	// `tee_hardware_attestation`).
	Capabilities []string `yaml:"capabilities,omitempty"`
	// Auditors is an optional list of specific auditor bech32 addresses whose
	// attestations the tenant requires; evaluated according to `AuditorMode`.
	Auditors []string `yaml:"auditors,omitempty"`
	// AuditorMode controls how the `Auditors` list is evaluated. Allowed
	// values: "" (defaults to "any"), "any", "all".
	AuditorMode string `yaml:"auditor_mode,omitempty"`
	// MinAuditorCount is the minimum number of independent qualifying
	// auditors that must have attested the provider, independent of
	// `Auditors`. Zero means "no count requirement".
	MinAuditorCount uint32 `yaml:"min_auditor_count,omitempty"`
}

// auditorBech32Prefix is the expected human-readable prefix on auditor
// addresses listed in an SDL `verification.auditors` block. The parser does
// not perform a full bech32 decode here — that is done by `x/market`
// validation on-chain — but it does enforce a non-empty value and a
// recognizable prefix so that obvious typos are rejected at SDL load time.
const auditorBech32Prefix = "akash1"

// UnmarshalYAML implements custom YAML decoding for the `verification:` block.
// It enforces the value constraints described in AEP-86 README §SDL Syntax:
//   - `min_tier` must be in [0, 4].
//   - `auditor_mode` must be "", "any", or "all".
//   - every entry of `capabilities` must be a recognized capability name
//     (see verificationv1.IsKnownCapabilityName).
//   - every entry of `auditors` must be a non-empty string with the expected
//     bech32 human-readable prefix.
//
// Constraints that the on-chain `x/market` filter applies (e.g. confirming
// each auditor address decodes as bech32, or that `min_auditor_count` is
// physically achievable) are intentionally NOT duplicated here — they live
// on-chain so the SDL parser stays cheap and dependency-light.
func (v *v2Verification) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type raw v2Verification
	var r raw
	if err := unmarshal(&r); err != nil {
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
		if !strings.HasPrefix(trimmed, auditorBech32Prefix) {
			return fmt.Errorf(
				"%w: verification.auditors[%d] %q does not have expected bech32 prefix %q",
				errSDLInvalid, i, trimmed, auditorBech32Prefix,
			)
		}
	}

	*v = v2Verification(r)
	return nil
}

// isVacuous reports whether this block expresses no actual filtering: no tier
// floor, no required capabilities, no named auditors, and no minimum auditor
// count. `AuditorMode` is intentionally ignored because it is only a modifier
// on `Auditors` — with no auditors listed, the mode has no effect.
//
// A vacuous block is functionally indistinguishable from omitting the
// `verification:` block entirely, so `toProto()` collapses both to nil. The
// AEP-86 spec (README §SDL Syntax / §3.5) explicitly defines omitted and
// `min_tier=0` as equivalent "no filtering" cases.
func (v *v2Verification) isVacuous() bool {
	return v.MinTier == 0 &&
		len(v.Capabilities) == 0 &&
		len(v.Auditors) == 0 &&
		v.MinAuditorCount == 0
}

// toProto converts the SDL representation into the on-chain
// `verificationv1.VerificationRequirement` message. Returns nil when:
//   - the receiver is nil (SDL omitted the block), or
//   - the block is vacuous (all filtering fields zero — see `isVacuous`).
//
// Collapsing the vacuous case to nil keeps the on-chain `PlacementRequirements`
// canonical: there is exactly one wire shape for "no verification filtering"
// rather than two semantically-equivalent ones, which prevents wasteful chain
// state and avoids latent bugs in BidFilter logic that might naively call
// `TierAtLeast(providerTier, TierUnspecified)` without short-circuiting.
//
// NOTE for L-8 (the on-chain BidFilter): the keeper MUST still short-circuit
// on vacuous-but-non-nil requirements that arrive through non-SDL paths
// (e.g. direct gRPC TX submission). The SDL collapse is a layer-1 convenience,
// not a chain-side invariant.
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
		// Should be unreachable: UnmarshalYAML validates the same set. Fall
		// back to "any" rather than panicking so a hand-constructed struct
		// produces sensible output.
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
