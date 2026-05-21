package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	collectionFile      = "collection.json"
	nonceFile           = "nonce.bin"
	snapshotPayloadFile = "snapshot_payload.pb"
	snapshotJSONFile    = "snapshot_payload.json"
	snapshotSigFile     = "snapshot_signature.bin"
)

func validateCollectedArtifactLayout(artifactDir string, evidence EvidenceDocument) error {
	raw, err := os.ReadFile(filepath.Join(artifactDir, collectionFile))
	if err != nil {
		return fmt.Errorf("read %s: %w", collectionFile, err)
	}

	var collection collectionArtifact
	if err := json.Unmarshal(raw, &collection); err != nil {
		return fmt.Errorf("decode %s: %w", collectionFile, err)
	}

	nonce, err := os.ReadFile(filepath.Join(artifactDir, nonceFile))
	if err != nil {
		return fmt.Errorf("read %s: %w", nonceFile, err)
	}
	nonceRef := base64.StdEncoding.EncodeToString(nonce)
	if evidence.InventoryNonce != nonceRef {
		return fmt.Errorf("inventory nonce mismatch between evidence and %s", nonceFile)
	}
	if collection.InventoryNonce != "" && collection.InventoryNonce != nonceRef {
		return fmt.Errorf("inventory nonce mismatch between %s and %s", collectionFile, nonceFile)
	}

	payload, err := os.ReadFile(filepath.Join(artifactDir, snapshotPayloadFile))
	if err != nil {
		return fmt.Errorf("read %s: %w", snapshotPayloadFile, err)
	}
	payloadHash := sha256Ref(sha256Bytes(payload))
	if evidence.SnapshotHash != payloadHash {
		return fmt.Errorf("snapshot hash mismatch between evidence and %s", snapshotPayloadFile)
	}
	if collection.SnapshotPayloadHash != "" && collection.SnapshotPayloadHash != payloadHash {
		return fmt.Errorf("snapshot hash mismatch between %s and %s", collectionFile, snapshotPayloadFile)
	}

	signature, err := os.ReadFile(filepath.Join(artifactDir, snapshotSigFile))
	if err != nil {
		return fmt.Errorf("read %s: %w", snapshotSigFile, err)
	}
	if len(signature) == 0 {
		return fmt.Errorf("%s is empty", snapshotSigFile)
	}
	if _, err := os.Stat(filepath.Join(artifactDir, snapshotJSONFile)); err != nil {
		return fmt.Errorf("read %s: %w", snapshotJSONFile, err)
	}

	if collection.Provider != "" && collection.Provider != evidence.Provider {
		return fmt.Errorf("provider mismatch between %s and evidence", collectionFile)
	}
	if collection.Auditor != "" && collection.Auditor != evidence.Auditor {
		return fmt.Errorf("auditor mismatch between %s and evidence", collectionFile)
	}
	if collection.AuditEscrowID != "" && !sameUint64String(collection.AuditEscrowID, evidence.AuditEscrowID) {
		return fmt.Errorf("audit escrow mismatch between %s and evidence", collectionFile)
	}
	if collection.ChainID != "" && collection.ChainID != evidence.ChainID {
		return fmt.Errorf("chain id mismatch between %s and evidence", collectionFile)
	}
	if collection.BlockHeight != "" && !sameUint64String(collection.BlockHeight, evidence.BlockHeight) {
		return fmt.Errorf("block height mismatch between %s and evidence", collectionFile)
	}

	return nil
}
