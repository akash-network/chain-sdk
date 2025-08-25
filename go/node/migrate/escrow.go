package migrate

import (
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1 "pkg.akt.dev/go/node/escrow/v1"
	"pkg.akt.dev/go/node/escrow/v1beta3"
)

func AccountV1beta3Prefix() []byte {
	return v1beta3.AccountKeyPrefix()
}

func PaymentV1beta3Prefix() []byte {
	return v1beta3.PaymentKeyPrefix()
}

func AccountIDFromV1beta3(from v1beta3.AccountID) v1.AccountID {
	return v1.AccountID{
		Scope: from.Scope,
		XID:   from.XID,
	}
}

func AccountFromV1beta3(cdc codec.BinaryCodec, fromBz []byte) v1.Account {
	var from v1beta3.Account
	cdc.MustUnmarshal(fromBz, &from)

	deposits := make([]v1.Deposit, 0)

	if from.Balance.IsPositive() {
		deposits = append(deposits, v1.Deposit{
			Depositor: from.Owner,
			Height:    0,
			Amount:    sdk.NewCoin(from.Balance.Denom, from.Balance.Amount.TruncateInt()),
			Balance:   from.Balance,
		})
	}

	if from.Funds.IsPositive() {
		deposits = append(deposits, v1.Deposit{
			Depositor: from.Depositor,
			Height:    0,
			Amount:    sdk.NewCoin(from.Funds.Denom, from.Funds.Amount.TruncateInt()),
			Balance:   from.Funds,
		})
	}

	to := v1.Account{
		ID:    AccountIDFromV1beta3(from.ID),
		Owner: from.Owner,
		State: v1.State(from.State),
		Funds: []v1.Funds{
			{
				Balance:   from.Balance.Add(from.Funds),
				Overdraft: sdkmath.LegacyZeroDec(),
			},
		},
		Transferred: sdk.DecCoins{
			from.Transferred,
		},
		SettledAt: from.SettledAt,
		Deposits:  deposits,
	}

	return to
}

func FractionalPaymentFromV1beta3(cdc codec.BinaryCodec, fromBz []byte) v1.FractionalPayment {
	var from v1beta3.FractionalPayment
	cdc.MustUnmarshal(fromBz, &from)

	to := v1.FractionalPayment{
		AccountID: AccountIDFromV1beta3(from.AccountID),
		PaymentID: from.PaymentID,
		Owner:     from.Owner,
		State:     v1.State(from.State),
		Rate:      from.Rate,
		Balance:   from.Balance,
		Withdrawn: from.Withdrawn,
	}

	return to
}
