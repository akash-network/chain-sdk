package v1

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"pkg.akt.dev/go/sdkutil"
)

func (r *RecordData) ValidateBasic() error {
	if _, err := ParseSolanaAddress(r.Dest); err != nil {
		return err
	}

	if err := r.Amount.Validate(); err != nil {
		return sdkerrors.ErrInvalidCoins.Wrap(err.Error())
	}

	if r.Amount.Denom != sdkutil.DenomUakt {
		return sdkerrors.ErrInvalidCoins.Wrap("invalid denomination. expected uakt")
	}

	//if err := ValidateAmount(r.Amount); err != nil {
	//	return err
	//}
	//
	//if !r.Amount.Amount.IsPositive() || !r.Amount.Amount.IsUint64() {
	//	return sdkerrors.ErrInvalidCoins.Wrap("transfer amount must be a positive uint64")
	//}

	return nil
}

func (t *TransferState) SettlementID() SettlementID {
	return NewSettlementID(t.Status, t.ID)
}
