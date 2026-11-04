package v1

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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
