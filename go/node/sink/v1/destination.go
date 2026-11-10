package v1

import (
	"bytes"

	"filippo.io/edwards25519"
	"github.com/gagliardetto/solana-go"
)

const (
	SolanaAddressSize    = solana.PublicKeyLength
	SolanaAddressSizeStr = 44
)

type SolanaAddress [SolanaAddressSize]byte

func (a SolanaAddress) String() string {
	return solana.PublicKey(a).String()
}

// ParseDestination accepts the canonical base58 public key carried by MsgSink.
// Unfunded wallets are valid; account ownership is checked on Solana at minting.
func ParseDestination(address string, vaults []SolanaAddress) (SolanaAddress, error) {
	recipient, err := ParseSolanaAddress(address)
	if err != nil {
		return SolanaAddress{}, err
	}

	if err = ValidateRecipient(recipient, vaults); err != nil {
		return SolanaAddress{}, err
	}

	return recipient, nil
}

// ParseSolanaAddress performs the stateless format checks: canonical base58 of
// 32 nonzero bytes. The curve check depends on the governed vault allowlist.
func ParseSolanaAddress(address string) (SolanaAddress, error) {
	var recipient SolanaAddress
	if len(address) > SolanaAddressSizeStr {
		return recipient, ErrInvalidDestination.Wrap("Solana address too long")
	}

	key, err := solana.PublicKeyFromBase58(address)
	if err != nil {
		return recipient, ErrInvalidDestination.Wrap("Solana address must be canonical base58 of 32 bytes")
	}

	recipient = SolanaAddress(key)

	if err := requireNonZero(recipient); err != nil {
		return SolanaAddress{}, err
	}

	return recipient, nil
}

func ValidateRecipient(recipient SolanaAddress, vaults []SolanaAddress) error {
	if err := requireNonZero(recipient); err != nil {
		return err
	}

	for _, vault := range vaults {
		if vault == recipient {
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

func requireNonZero(recipient SolanaAddress) error {
	if solana.PublicKey(recipient).IsZero() {
		return ErrInvalidDestination.Wrap("zero Solana recipient")
	}

	return nil
}
