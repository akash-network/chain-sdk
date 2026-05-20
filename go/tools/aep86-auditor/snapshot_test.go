package main

import (
	"strings"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	inventoryv1 "pkg.akt.dev/go/inventory/v1"
)

func TestVerifySnapshotEnvelopeAndSignature(t *testing.T) {
	priv := secp256k1.GenPrivKey()
	provider := sdk.AccAddress(priv.PubKey().Address()).String()
	nonce := []byte("12345678901234567890123456789012")

	payloadBytes := mustSnapshotPayload(t, provider, nonce)
	signature, err := priv.Sign(payloadBytes)
	require.NoError(t, err)

	resp := &inventoryv1.GetInventorySnapshotResponse{
		SnapshotPayload: payloadBytes,
		Signature:       signature,
		Provider:        provider,
	}

	verified, err := verifySnapshotEnvelope(resp, nonce)
	require.NoError(t, err)
	require.Equal(t, provider, verified.Provider)
	require.Equal(t, nonce, verified.Payload.GetNonce())

	err = verifyProviderSignature(verified.PayloadBytes, resp.GetSignature(), priv.PubKey(), provider)
	require.NoError(t, err)
}

func TestVerifySnapshotEnvelopeRejectsNonceMismatch(t *testing.T) {
	priv := secp256k1.GenPrivKey()
	provider := sdk.AccAddress(priv.PubKey().Address()).String()
	payloadBytes := mustSnapshotPayload(t, provider, []byte("12345678901234567890123456789012"))

	resp := &inventoryv1.GetInventorySnapshotResponse{
		SnapshotPayload: payloadBytes,
		Signature:       []byte("signature"),
		Provider:        provider,
	}

	_, err := verifySnapshotEnvelope(resp, []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	require.ErrorContains(t, err, "snapshot nonce mismatch")
}

func TestMarshalEvidenceCanonicalIsStable(t *testing.T) {
	evidence := EvidenceDocument{
		SchemaVersion:        evidenceSchema,
		ChainID:              "akash-local",
		Provider:             "akash1provider",
		Auditor:              "akash1auditor",
		AuditEscrowID:        "7",
		TargetTier:           "L1",
		AttestedTier:         "L1",
		AttestedCapabilities: []string{"persistent_storage"},
		CollectedAt:          "2026-05-19T00:00:00Z",
		BlockHeight:          "123",
		SnapshotHash:         "sha256:0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
		InventoryNonce:       "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI=",
		Software: SoftwareEvidence{
			Version:            "test",
			BinaryHash:         "sha256:abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
			VerificationStatus: "observed_only",
		},
		NetworkBaseline: NetworkBaseline{
			ProofRef: "sha256:0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
		},
		SustainedValidation: SustainedValidation{
			BaselineID:    "sha256:0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
			Window:        "attestation_ttl",
			LastCheckedAt: "2026-05-19T00:00:00Z",
			Status:        "not_evaluated",
			ProofRef:      "sha256:0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
		},
		Checks: []EvidenceCheck{{
			Name:     "inventory_signature_valid",
			Status:   "pass",
			ProofRef: "sha256:0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
			Details: map[string]any{
				"z": true,
				"a": "first",
			},
		}},
		FaultContext: FaultContext{
			FaultAttribution: "unspecified",
			Reason:           "unspecified",
		},
	}

	first, firstHash, err := marshalEvidenceCanonical(evidence)
	require.NoError(t, err)
	second, secondHash, err := marshalEvidenceCanonical(evidence)
	require.NoError(t, err)

	require.Equal(t, string(first), string(second))
	require.Equal(t, firstHash, secondHash)
	require.True(t, strings.HasPrefix(firstHash, "sha256:"))
	require.NotContains(t, string(first), "\n")
}

func TestMarshalEvidenceCanonicalEmptyAttestedCapabilitiesIsArray(t *testing.T) {
	evidence := EvidenceDocument{}

	raw, _, err := marshalEvidenceCanonical(evidence)
	require.NoError(t, err)
	require.Contains(t, string(raw), `"attested_capabilities":[]`)
	require.NotContains(t, string(raw), `"attested_capabilities":null`)
}

func TestValidateEvidenceInputsRequiresSoftwareBinaryHash(t *testing.T) {
	cfg := validCollectConfig()
	cfg.softwareBinaryHash = ""

	err := validateEvidenceInputs(cfg)
	require.ErrorContains(t, err, "software-binary-hash is required")
}

func TestValidateEvidenceInputsRejectsInvalidSoftwareBinaryHash(t *testing.T) {
	cfg := validCollectConfig()
	cfg.softwareBinaryHash = "sha256:not-hex"

	err := validateEvidenceInputs(cfg)
	require.ErrorContains(t, err, "software-binary-hash must use sha256:<64 hex> form")
}

func TestValidateEvidenceInputsAcceptsSoftwareBinaryHash(t *testing.T) {
	cfg := validCollectConfig()

	err := validateEvidenceInputs(cfg)
	require.NoError(t, err)
}

func mustSnapshotPayload(t *testing.T, provider string, nonce []byte) []byte {
	t.Helper()

	payload := &inventoryv1.SnapshotPayload{
		SchemaVersion: 1,
		Provider:      provider,
		ChainID:       "akash-local",
		Nonce:         nonce,
		Timestamp:     time.Unix(1, 0).UTC(),
		ResourceSummary: inventoryv1.SnapshotResourceSummary{
			SoftwareVersion: "provider-services-test",
		},
	}

	raw, err := proto.Marshal(payload)
	require.NoError(t, err)

	return raw
}

func validCollectConfig() collectConfig {
	return collectConfig{
		providerGRPC:       "provider.example.com:8443",
		chainGRPC:          "rpc.example.com:9090",
		auditor:            "akash1auditor",
		auditEscrowID:      "7",
		targetTier:         "L1",
		attestedTier:       "L1",
		softwareBinaryHash: "sha256:0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	}
}
