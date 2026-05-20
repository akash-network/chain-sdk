package main

import (
	"bytes"
	"fmt"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	inventoryv1 "pkg.akt.dev/go/inventory/v1"
)

type verifiedSnapshot struct {
	PayloadBytes          []byte
	Payload               *inventoryv1.SnapshotPayload
	PayloadHash           []byte
	Provider              string
	SignatureVerified     bool
	SignatureSkipped      bool
	ProviderPubKeyAddress string
}

func verifySnapshotEnvelope(resp *inventoryv1.GetInventorySnapshotResponse, nonce []byte) (*verifiedSnapshot, error) {
	if resp == nil {
		return nil, fmt.Errorf("empty inventory snapshot response")
	}
	if len(resp.GetSnapshotPayload()) == 0 {
		return nil, fmt.Errorf("empty inventory snapshot payload")
	}
	if len(resp.GetSignature()) == 0 {
		return nil, fmt.Errorf("empty inventory snapshot signature")
	}

	payload, err := protoUnmarshalSnapshotPayload(resp.GetSnapshotPayload())
	if err != nil {
		return nil, fmt.Errorf("decode snapshot payload: %w", err)
	}
	if !bytes.Equal(payload.GetNonce(), nonce) {
		return nil, fmt.Errorf("snapshot nonce mismatch")
	}
	if payload.GetProvider() == "" {
		return nil, fmt.Errorf("snapshot payload missing provider")
	}
	if resp.GetProvider() != "" && resp.GetProvider() != payload.GetProvider() {
		return nil, fmt.Errorf("snapshot response provider %q does not match payload provider %q", resp.GetProvider(), payload.GetProvider())
	}
	if payload.GetChainID() == "" {
		return nil, fmt.Errorf("snapshot payload missing chain_id")
	}
	if payload.GetSchemaVersion() == 0 {
		return nil, fmt.Errorf("snapshot payload missing schema_version")
	}

	return &verifiedSnapshot{
		PayloadBytes: append([]byte(nil), resp.GetSnapshotPayload()...),
		Payload:      payload,
		PayloadHash:  sha256Bytes(resp.GetSnapshotPayload()),
		Provider:     payload.GetProvider(),
	}, nil
}

func verifyProviderSignature(payload, signature []byte, pubKey cryptotypes.PubKey, provider string) error {
	if pubKey == nil {
		return fmt.Errorf("missing provider public key")
	}

	pubKeyAddress := sdk.AccAddress(pubKey.Address()).String()
	if pubKeyAddress != provider {
		return fmt.Errorf("provider public key address %q does not match snapshot provider %q", pubKeyAddress, provider)
	}

	if !pubKey.VerifySignature(payload, signature) {
		return fmt.Errorf("provider signature verification failed")
	}

	return nil
}
