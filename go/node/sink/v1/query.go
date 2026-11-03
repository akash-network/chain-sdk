package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (q *QueryTransferRequest) ValidateBasic() error {
	return q.ID.ValidateBasic()
}

func (q *QueryTransfersRequest) ValidateBasic() error {
	if q.Owner != "" {
		if _, err := sdk.AccAddressFromBech32(q.Owner); err != nil {
			return err
		}
	}

	if q.Status != StatusInvalid && !q.Status.IsValid() {
		return ErrInvalidSettlement.Wrap("invalid status filter")
	}

	return nil
}

func (q *QueryDisputesRequest) ValidateBasic() error {
	return nil
}

func (q *QueryAccountRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(q.Owner)

	return err
}
