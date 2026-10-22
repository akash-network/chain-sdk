package v1

import (
	"bytes"
	"fmt"

	"filippo.io/edwards25519"
	"github.com/mr-tron/base58"
)

// ParseDestination accepts the canonical base58 public key carried by MsgSink.
// Unfunded wallets are valid; account ownership is checked on Solana at minting.
func ParseDestination(address string, vaults [][]byte) ([32]byte, error) {
	recipient, err := DecodeDestination(address)
	if err != nil {
		return [32]byte{}, err
	}
	if err = ValidateRecipient(recipient, vaults); err != nil {
		return [32]byte{}, err
	}
	return recipient, nil
}

// DecodeDestination performs the stateless format checks: canonical base58 of
// 32 nonzero bytes. The curve check depends on the governed vault allowlist.
func DecodeDestination(address string) ([32]byte, error) {
	var recipient [32]byte
	if len(address) > 44 {
		return recipient, fmt.Errorf("Solana address too long") //nolint:staticcheck // Solana is a proper noun
	}
	b, err := base58.Decode(address)
	if err != nil || len(b) != 32 || base58.Encode(b) != address {
		return recipient, fmt.Errorf("Solana address must be canonical base58 of 32 bytes") //nolint:staticcheck // Solana is a proper noun
	}
	copy(recipient[:], b)
	if recipient == [32]byte{} {
		return [32]byte{}, fmt.Errorf("zero Solana recipient")
	}
	return recipient, nil
}

func ValidateRecipient(recipient [32]byte, vaults [][]byte) error {
	if recipient == [32]byte{} {
		return fmt.Errorf("zero Solana recipient")
	}
	for _, vault := range vaults {
		if bytes.Equal(vault, recipient[:]) {
			return nil
		}
	}
	// SetBytes also accepts non-canonical encodings and small-order points. No
	// wallet holds a key for either, so tokens minted there could never move.
	point, err := new(edwards25519.Point).SetBytes(recipient[:])
	if err != nil || !bytes.Equal(point.Bytes(), recipient[:]) {
		return fmt.Errorf("Solana recipient must be a canonical on-curve key or an allowlisted vault") //nolint:staticcheck // Solana is a proper noun
	}
	if new(edwards25519.Point).MultByCofactor(point).Equal(edwards25519.NewIdentityPoint()) == 1 {
		return fmt.Errorf("Solana recipient is a small-order point") //nolint:staticcheck // Solana is a proper noun
	}
	return nil
}
