package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/spf13/cobra"

	inventoryv1 "pkg.akt.dev/go/inventory/v1"
)

const (
	nonceSize      = 32
	evidenceSchema = "akash.audit.evidence.v1"
)

type collectConfig struct {
	providerGRPC            string
	providerInsecure        bool
	providerSkipTLSVerify   bool
	providerTLSServerName   string
	chainGRPC               string
	chainGRPCTLS            bool
	chainGRPCSkipTLSVerify  bool
	chainGRPCTLSServerName  string
	auditor                 string
	auditEscrowID           string
	targetTier              string
	attestedTier            string
	attestedCapabilities    []string
	softwareBinaryHash      string
	outputDir               string
	timeout                 time.Duration
	allowMissingChainPubKey bool
}

type collectionArtifact struct {
	SchemaVersion         string            `json:"schema_version"`
	CollectedAt           string            `json:"collected_at"`
	ProviderEndpoint      string            `json:"provider_endpoint"`
	ChainGRPCEndpoint     string            `json:"chain_grpc_endpoint"`
	Provider              string            `json:"provider"`
	Auditor               string            `json:"auditor"`
	AuditEscrowID         string            `json:"audit_escrow_id"`
	ChainID               string            `json:"chain_id"`
	BlockHeight           string            `json:"block_height"`
	SnapshotPayloadHash   string            `json:"snapshot_payload_hash"`
	InventoryNonce        string            `json:"inventory_nonce"`
	Signature             string            `json:"signature"`
	SignatureVerified     bool              `json:"signature_verified"`
	SignatureSkipped      bool              `json:"signature_skipped,omitempty"`
	ProviderPubKeyAddress string            `json:"provider_pubkey_address,omitempty"`
	Payload               payloadSummary    `json:"payload"`
	Checks                []EvidenceCheck   `json:"checks"`
	Warnings              []string          `json:"warnings,omitempty"`
	Files                 map[string]string `json:"files"`
	ChainFacts            map[string]any    `json:"chain_facts,omitempty"`
}

type payloadSummary struct {
	SchemaVersion    uint32 `json:"schema_version"`
	Provider         string `json:"provider"`
	ChainID          string `json:"chain_id"`
	Timestamp        string `json:"timestamp"`
	TotalGPUs        uint32 `json:"total_gpus"`
	TotalVCPUs       uint32 `json:"total_vcpus"`
	TotalMemoryMB    uint64 `json:"total_memory_mb"`
	TotalStorageMB   uint64 `json:"total_storage_mb"`
	ActiveLeases     uint32 `json:"active_leases"`
	SoftwareVersion  string `json:"software_version"`
	EvidenceSections int    `json:"evidence_sections"`
}

func newCollectCmd() *cobra.Command {
	cfg := collectConfig{}

	cmd := &cobra.Command{
		Use:   "collect",
		Short: "Collect and verify a nonce-bound provider inventory snapshot",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCollect(cmd, cfg)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&cfg.providerGRPC, "provider-grpc", "", "Provider daemon gRPC endpoint that serves akash.inventory.v1.InventoryService")
	flags.BoolVar(&cfg.providerInsecure, "provider-insecure", false, "Use plaintext for provider gRPC")
	flags.BoolVar(&cfg.providerSkipTLSVerify, "provider-skip-tls-verify", false, "Skip provider TLS certificate verification")
	flags.StringVar(&cfg.providerTLSServerName, "provider-tls-server-name", "", "Provider TLS server name override")
	flags.StringVar(&cfg.chainGRPC, "chain-grpc", "", "Akash chain gRPC endpoint used to query the provider account public key")
	flags.BoolVar(&cfg.chainGRPCTLS, "chain-grpc-tls", false, "Use TLS for chain gRPC")
	flags.BoolVar(&cfg.chainGRPCSkipTLSVerify, "chain-grpc-skip-tls-verify", false, "Skip chain gRPC TLS certificate verification")
	flags.StringVar(&cfg.chainGRPCTLSServerName, "chain-grpc-tls-server-name", "", "Chain gRPC TLS server name override")
	flags.StringVar(&cfg.auditor, "auditor", "", "Auditor account address to include in draft evidence")
	flags.StringVar(&cfg.auditEscrowID, "audit-escrow-id", "0", "Audit escrow id to include in draft evidence")
	flags.StringVar(&cfg.targetTier, "target-tier", "L1", "Target verification tier for draft evidence")
	flags.StringVar(&cfg.attestedTier, "attested-tier", "", "Attested verification tier for draft evidence; defaults to target tier")
	flags.StringSliceVar(&cfg.attestedCapabilities, "capability", nil, "Attested capability for draft evidence; repeat or comma-separate")
	flags.StringVar(&cfg.softwareBinaryHash, "software-binary-hash", "", "Observed provider software binary hash in sha256:<hex> form")
	flags.StringVar(&cfg.outputDir, "output-dir", "", "Directory for raw artifacts and draft evidence")
	flags.DurationVar(&cfg.timeout, "timeout", 30*time.Second, "Collection timeout")
	flags.BoolVar(&cfg.allowMissingChainPubKey, "allow-missing-chain-pubkey", false, "Write artifacts even when provider account public key cannot be queried")

	_ = cmd.MarkFlagRequired("provider-grpc")
	_ = cmd.MarkFlagRequired("chain-grpc")
	_ = cmd.MarkFlagRequired("auditor")

	return cmd
}

func runCollect(cmd *cobra.Command, cfg collectConfig) error {
	if cfg.attestedTier == "" {
		cfg.attestedTier = cfg.targetTier
	}

	if err := validateEvidenceInputs(cfg); err != nil {
		return err
	}

	ctx, cancel := timeoutContext(cmd, cfg.timeout)
	defer cancel()

	nonce, err := randomNonce()
	if err != nil {
		return err
	}

	providerConn, err := dialGRPC(ctx, grpcDialConfig{
		endpoint:      cfg.providerGRPC,
		insecure:      cfg.providerInsecure,
		skipTLSVerify: cfg.providerSkipTLSVerify,
		serverName:    cfg.providerTLSServerName,
	})
	if err != nil {
		return err
	}
	defer providerConn.Close()

	snapshotResp, err := inventoryv1.NewInventoryServiceClient(providerConn).GetInventorySnapshot(ctx, &inventoryv1.GetInventorySnapshotRequest{
		Nonce: nonce,
	})
	if err != nil {
		return err
	}

	verified, err := verifySnapshotEnvelope(snapshotResp, nonce)
	if err != nil {
		return err
	}

	chainFacts, err := collectChainFacts(ctx, cfg, verified.Provider)
	if err != nil {
		if !cfg.allowMissingChainPubKey {
			return err
		}
		chainFacts = &chainFactsResult{
			BlockHeight: "0",
			Warnings:    []string{err.Error()},
		}
	}

	if chainFacts.ProviderPubKey != nil {
		if err := verifyProviderSignature(verified.PayloadBytes, snapshotResp.GetSignature(), chainFacts.ProviderPubKey, verified.Provider); err != nil {
			return err
		}
		verified.SignatureVerified = true
		verified.ProviderPubKeyAddress = chainFacts.ProviderPubKeyAddress
	} else {
		verified.SignatureSkipped = true
	}

	outDir := cfg.outputDir
	if outDir == "" {
		outDir = defaultOutputDir(verified.Provider)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}

	files, err := writeArtifacts(outDir, snapshotResp, verified.Payload, nonce)
	if err != nil {
		return err
	}

	collectedAt := time.Now().UTC()
	checks := evidenceChecks(verified, chainFacts)
	warnings := append([]string(nil), chainFacts.Warnings...)
	if cfg.softwareBinaryHash == "" {
		warnings = append(warnings, "software_binary_hash not supplied; draft evidence is shape-complete but not schema-valid for final submission")
	}

	evidence := buildEvidence(cfg, verified, chainFacts, collectedAt, checks)
	evidenceBytes, evidenceHash, err := marshalEvidenceCanonical(evidence)
	if err != nil {
		return err
	}

	evidencePath := filepath.Join(outDir, "evidence.draft.json")
	if err := os.WriteFile(evidencePath, evidenceBytes, 0o644); err != nil {
		return err
	}
	files["evidence_draft"] = evidencePath

	hashPath := filepath.Join(outDir, "evidence.draft.sha256")
	if err := os.WriteFile(hashPath, []byte(evidenceHash+"\n"), 0o644); err != nil {
		return err
	}
	files["evidence_draft_hash"] = hashPath

	artifact := collectionArtifact{
		SchemaVersion:         evidenceSchema + ".collection",
		CollectedAt:           collectedAt.Format(time.RFC3339Nano),
		ProviderEndpoint:      cfg.providerGRPC,
		ChainGRPCEndpoint:     cfg.chainGRPC,
		Provider:              verified.Provider,
		Auditor:               cfg.auditor,
		AuditEscrowID:         cfg.auditEscrowID,
		ChainID:               verified.Payload.GetChainID(),
		BlockHeight:           chainFacts.BlockHeight,
		SnapshotPayloadHash:   sha256Ref(verified.PayloadHash),
		InventoryNonce:        base64.StdEncoding.EncodeToString(nonce),
		Signature:             base64.StdEncoding.EncodeToString(snapshotResp.GetSignature()),
		SignatureVerified:     verified.SignatureVerified,
		SignatureSkipped:      verified.SignatureSkipped,
		ProviderPubKeyAddress: verified.ProviderPubKeyAddress,
		Payload:               summarizePayload(verified.Payload),
		Checks:                checks,
		Warnings:              warnings,
		Files:                 files,
		ChainFacts:            chainFacts.Facts,
	}

	collectionPath := filepath.Join(outDir, "collection.json")
	if err := writeJSON(collectionPath, artifact); err != nil {
		return err
	}

	fmt.Fprintf(cmd.OutOrStdout(), "wrote collection artifacts to %s\n", outDir)
	fmt.Fprintf(cmd.OutOrStdout(), "snapshot_hash=%s\n", sha256Ref(verified.PayloadHash))
	fmt.Fprintf(cmd.OutOrStdout(), "evidence_hash=%s\n", evidenceHash)

	return nil
}

func validateEvidenceInputs(cfg collectConfig) error {
	if cfg.providerGRPC == "" {
		return fmt.Errorf("provider-grpc is required")
	}
	if cfg.chainGRPC == "" {
		return fmt.Errorf("chain-grpc is required")
	}
	if cfg.auditor == "" {
		return fmt.Errorf("auditor is required")
	}
	if cfg.auditEscrowID == "" {
		return fmt.Errorf("audit-escrow-id is required")
	}
	if _, err := strconv.ParseUint(cfg.auditEscrowID, 10, 64); err != nil {
		return fmt.Errorf("invalid audit-escrow-id: %w", err)
	}
	if !isTier(cfg.targetTier) {
		return fmt.Errorf("invalid target-tier %q", cfg.targetTier)
	}
	if !isTier(cfg.attestedTier) {
		return fmt.Errorf("invalid attested-tier %q", cfg.attestedTier)
	}
	if cfg.softwareBinaryHash != "" && !isSHA256Ref(cfg.softwareBinaryHash) {
		return fmt.Errorf("software-binary-hash must use sha256:<64 hex> form")
	}

	return nil
}

func randomNonce() ([]byte, error) {
	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	return nonce, nil
}

func writeArtifacts(dir string, resp *inventoryv1.GetInventorySnapshotResponse, payload *inventoryv1.SnapshotPayload, nonce []byte) (map[string]string, error) {
	files := map[string]string{
		"nonce":            filepath.Join(dir, "nonce.bin"),
		"snapshot_payload": filepath.Join(dir, "snapshot_payload.pb"),
		"signature":        filepath.Join(dir, "snapshot_signature.bin"),
		"payload_json":     filepath.Join(dir, "snapshot_payload.json"),
	}

	if err := os.WriteFile(files["nonce"], nonce, 0o644); err != nil {
		return nil, err
	}
	if err := os.WriteFile(files["snapshot_payload"], resp.GetSnapshotPayload(), 0o644); err != nil {
		return nil, err
	}
	if err := os.WriteFile(files["signature"], resp.GetSignature(), 0o644); err != nil {
		return nil, err
	}
	if err := writeJSON(files["payload_json"], payload); err != nil {
		return nil, err
	}

	return files, nil
}

func writeJSON(path string, val any) error {
	raw, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		return err
	}
	raw = append(raw, '\n')

	return os.WriteFile(path, raw, 0o644)
}

func summarizePayload(payload *inventoryv1.SnapshotPayload) payloadSummary {
	summary := payload.GetResourceSummary()

	return payloadSummary{
		SchemaVersion:    payload.GetSchemaVersion(),
		Provider:         payload.GetProvider(),
		ChainID:          payload.GetChainID(),
		Timestamp:        payload.GetTimestamp().UTC().Format(time.RFC3339Nano),
		TotalGPUs:        summary.GetTotalGPUs(),
		TotalVCPUs:       summary.GetTotalVCPUs(),
		TotalMemoryMB:    summary.GetTotalMemoryMB(),
		TotalStorageMB:   summary.GetTotalStorageMB(),
		ActiveLeases:     summary.GetActiveLeases(),
		SoftwareVersion:  summary.GetSoftwareVersion(),
		EvidenceSections: len(payload.GetEvidenceSections()),
	}
}

func defaultOutputDir(provider string) string {
	cleanProvider := strings.ReplaceAll(provider, string(os.PathSeparator), "_")
	if cleanProvider == "" {
		cleanProvider = "unknown-provider"
	}

	return fmt.Sprintf("aep86-audit-%s-%s", cleanProvider, time.Now().UTC().Format("20060102T150405Z"))
}

func sha256Bytes(payload []byte) []byte {
	hash := sha256.Sum256(payload)
	return hash[:]
}

func sha256Ref(hash []byte) string {
	return "sha256:" + hex.EncodeToString(hash)
}

func isSHA256Ref(val string) bool {
	if !strings.HasPrefix(val, "sha256:") {
		return false
	}
	raw := strings.TrimPrefix(val, "sha256:")
	if len(raw) != sha256.Size*2 {
		return false
	}
	_, err := hex.DecodeString(raw)
	return err == nil
}

func isTier(val string) bool {
	switch val {
	case "L1", "L2", "L3", "L4":
		return true
	default:
		return false
	}
}

func protoUnmarshalSnapshotPayload(raw []byte) (*inventoryv1.SnapshotPayload, error) {
	var payload inventoryv1.SnapshotPayload
	if err := proto.Unmarshal(raw, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
