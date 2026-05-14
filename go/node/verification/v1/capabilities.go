package v1

// This file (capabilities.go) is NOT generated. The codegen pipeline only
// removes files matching *.pb.go and *.pb.gw.go, so this file is preserved
// across `make codegen` runs.

// Lowercase SDL names for the capability flags. These match the strings
// accepted in the v2.1 SDL `verification.capabilities` block and the strings
// produced when serializing a capability back to YAML.
const (
	capabilityNameTEEHardwareAttestation = "tee_hardware_attestation"
	capabilityNameConfidentialComputing  = "confidential_computing"
	capabilityNamePersistentStorage      = "persistent_storage"
	capabilityNameBareMetal              = "bare_metal"
)

// knownCapabilityNameToFlag maps the lowercase SDL capability name to its
// generated CapabilityFlag enum value. Only the four spec-defined flags are
// included; CapabilityUnspecified is deliberately omitted because it is not a
// real capability — it is the zero default and must not be accepted from SDL
// input.
var knownCapabilityNameToFlag = map[string]CapabilityFlag{
	capabilityNameTEEHardwareAttestation: CapabilityTEEHardwareAttestation,
	capabilityNameConfidentialComputing:  CapabilityConfidentialComputing,
	capabilityNamePersistentStorage:      CapabilityPersistentStorage,
	capabilityNameBareMetal:              CapabilityBareMetal,
}

// knownCapabilityFlagToName is the inverse of knownCapabilityNameToFlag.
// CapabilityUnspecified intentionally has no entry.
var knownCapabilityFlagToName = map[CapabilityFlag]string{
	CapabilityTEEHardwareAttestation: capabilityNameTEEHardwareAttestation,
	CapabilityConfidentialComputing:  capabilityNameConfidentialComputing,
	CapabilityPersistentStorage:      capabilityNamePersistentStorage,
	CapabilityBareMetal:              capabilityNameBareMetal,
}

// KnownCapabilities returns the spec-defined CapabilityFlag values in
// declaration order: TEE hardware attestation, confidential computing,
// persistent storage, bare metal. CapabilityUnspecified is excluded.
//
// The returned slice is a fresh copy on each call; callers may mutate it
// freely without affecting subsequent invocations.
func KnownCapabilities() []CapabilityFlag {
	return []CapabilityFlag{
		CapabilityTEEHardwareAttestation,
		CapabilityConfidentialComputing,
		CapabilityPersistentStorage,
		CapabilityBareMetal,
	}
}

// IsKnownCapabilityName looks up a CapabilityFlag by its lowercase SDL name.
// It returns the matching flag and true on hit, or CapabilityUnspecified and
// false on miss. The four recognized names are: tee_hardware_attestation,
// confidential_computing, persistent_storage, bare_metal.
func IsKnownCapabilityName(s string) (CapabilityFlag, bool) {
	c, ok := knownCapabilityNameToFlag[s]
	return c, ok
}

// CapabilityName returns the lowercase SDL name for the given CapabilityFlag.
// For unknown values (including CapabilityUnspecified) it returns the empty
// string so that callers can distinguish "no name" from any real name.
func CapabilityName(c CapabilityFlag) string {
	return knownCapabilityFlagToName[c]
}
