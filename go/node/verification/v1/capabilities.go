package v1

const (
	capabilityNameTEEHardwareAttestation = "tee_hardware_attestation"
	capabilityNameConfidentialComputing  = "confidential_computing"
	capabilityNamePersistentStorage      = "persistent_storage"
	capabilityNameBareMetal              = "bare_metal"
)

var knownCapabilityNameToFlag = map[string]CapabilityFlag{
	capabilityNameTEEHardwareAttestation: CapabilityTEEHardwareAttestation,
	capabilityNameConfidentialComputing:  CapabilityConfidentialComputing,
	capabilityNamePersistentStorage:      CapabilityPersistentStorage,
	capabilityNameBareMetal:              CapabilityBareMetal,
}

var knownCapabilityFlagToName = map[CapabilityFlag]string{
	CapabilityTEEHardwareAttestation: capabilityNameTEEHardwareAttestation,
	CapabilityConfidentialComputing:  capabilityNameConfidentialComputing,
	CapabilityPersistentStorage:      capabilityNamePersistentStorage,
	CapabilityBareMetal:              capabilityNameBareMetal,
}

// KnownCapabilities returns known capability flags.
func KnownCapabilities() []CapabilityFlag {
	return []CapabilityFlag{
		CapabilityTEEHardwareAttestation,
		CapabilityConfidentialComputing,
		CapabilityPersistentStorage,
		CapabilityBareMetal,
	}
}

// IsKnownCapabilityName returns a CapabilityFlag for name.
func IsKnownCapabilityName(s string) (CapabilityFlag, bool) {
	c, ok := knownCapabilityNameToFlag[s]
	return c, ok
}

// CapabilityName returns the SDL name for c.
func CapabilityName(c CapabilityFlag) string {
	return knownCapabilityFlagToName[c]
}
