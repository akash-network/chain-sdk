package v1

import (
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

func (a SolanaAddress) Validate() error {
	if err := requireNonZero(a); err != nil {
		return err
	}

	return nil
}

// ParseSolanaAddress performs the stateless format checks: canonical base58 of
// 32 nonzero bytes. The curve check depends on the governed vault allowlist.
func ParseSolanaAddress(address string) (SolanaAddress, error) {
	var recipient SolanaAddress
	if len(address) > SolanaAddressSizeStr {
		return SolanaAddress{}, ErrInvalidDestination.Wrap("Solana address too long")
	}

	key, err := solana.PublicKeyFromBase58(address)
	if err != nil {
		return SolanaAddress{}, ErrInvalidDestination.Wrap("Solana address must be canonical base58 of 32 bytes")
	}

	recipient = SolanaAddress(key)

	if err = recipient.Validate(); err != nil {
		return SolanaAddress{}, err
	}

	return recipient, nil
}

func requireNonZero(recipient SolanaAddress) error {
	if solana.PublicKey(recipient).IsZero() {
		return ErrInvalidDestination.Wrap("zero Solana recipient")
	}

	return nil
}
