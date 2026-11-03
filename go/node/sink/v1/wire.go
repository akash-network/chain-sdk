package v1

import (
	"crypto/sha256"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	TransferKeyPrefix   = []byte{0x02, 0x00}
	SettlementKeyPrefix = []byte{0x03, 0x01}
)

func TransferKey(id ID) ([]byte, error) {
	b := make([]byte, len(TransferKeyPrefix)+IDKeySize)
	copy(b, TransferKeyPrefix)

	if _, err := IDKey.Encode(b[len(TransferKeyPrefix):], id); err != nil {
		return nil, err
	}

	return b, nil
}

func SettlementKey(id SettlementID) ([]byte, error) {
	b := make([]byte, len(SettlementKeyPrefix)+SettlementIDKeySize)
	copy(b, SettlementKeyPrefix)

	if _, err := SettlementIDKey.Encode(b[len(SettlementKeyPrefix):], id); err != nil {
		return nil, err
	}

	return b, nil
}

// ProofNonce binds every byte of the full store key, including settlement status.
func ProofNonce(key []byte) [32]byte {
	return sha256.Sum256(key)
}

func (r RecordData) ValidateBasic() error {
	if _, err := ParseSolanaAddress(r.Dest); err != nil {
		return err
	}

	if err := ValidateAmount(r.Amount); err != nil {
		return err
	}

	if !r.Amount.Amount.IsPositive() || !r.Amount.Amount.IsUint64() {
		return sdkerrors.ErrInvalidCoins.Wrap("transfer amount must be a positive uint64")
	}

	return nil
}
