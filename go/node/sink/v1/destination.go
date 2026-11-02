package v1

import (
	"bytes"

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
		return recipient, ErrInvalidDestination.Wrap("Solana address too long")
	}
	b, err := base58.Decode(address)
	if err != nil || len(b) != 32 || base58.Encode(b) != address {
		return recipient, ErrInvalidDestination.Wrap("Solana address must be canonical base58 of 32 bytes")
	}
	copy(recipient[:], b)
	if err := requireNonZero(recipient); err != nil {
		return [32]byte{}, err
	}
	return recipient, nil
}

func ValidateRecipient(recipient [32]byte, vaults [][]byte) error {
	if err := requireNonZero(recipient); err != nil {
		return err
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
		return ErrInvalidDestination.Wrap("Solana recipient must be a canonical on-curve key or an allowlisted vault")
	}
	if new(edwards25519.Point).MultByCofactor(point).Equal(edwards25519.NewIdentityPoint()) == 1 {
		return ErrInvalidDestination.Wrap("Solana recipient is a small-order point")
	}
	return nil
}

// requireNonZero rejects the all-zero address, which no Solana account can hold.
func requireNonZero(recipient [32]byte) error {
	if recipient == [32]byte{} {
		return ErrInvalidDestination.Wrap("zero Solana recipient")
	}
	return nil
}
