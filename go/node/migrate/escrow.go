package migrate

import (
	"bytes"
	"strings"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	dv1beta3 "pkg.akt.dev/go/node/deployment/v1beta3"
	eid "pkg.akt.dev/go/node/escrow/id/v1"
	etypes "pkg.akt.dev/go/node/escrow/types/v1"
	ev1 "pkg.akt.dev/go/node/escrow/v1"
	"pkg.akt.dev/go/node/escrow/v1beta3"
)

func AccountV1beta3Prefix() []byte {
	return v1beta3.AccountKeyPrefix()
}

func PaymentV1beta3Prefix() []byte {
	return v1beta3.PaymentKeyPrefix()
}

func AccountIDFromV1beta3(key []byte) eid.Account {
	prefix := v1beta3.AccountKeyPrefix()

	if len(key) < len(prefix)+1 {
		panic("invalid escrow.v1beta3 key")
	}

	if !bytes.Equal(prefix, key[:len(prefix)]) {
		panic("invalid escrow.v1beta3 account prefix")
	}

	key = key[len(prefix):]
	if key[0] != '/' {
		panic("invalid escrow.v1beta3 account separator")
	}

	key = key[1:]

	parts := strings.Split(string(key), "/")
	if len(parts) != 3 {
		panic("invalid escrow.v1beta3 account xid")
	}

	return eid.Account{
		Scope: eid.Scope(eid.Scope_value[parts[0]]),
		XID:   strings.Join(parts[1:], "/"),
	}
}

func PaymentIDFromV1beta3(key []byte) eid.Payment {
	prefix := v1beta3.PaymentKeyPrefix()

	if len(key) < len(prefix)+1 {
		panic("invalid escrow.v1beta3 key")
	}

	if !bytes.Equal(prefix, key[:len(prefix)]) {
		panic("invalid escrow.v1beta3 account prefix")
	}

	key = key[len(prefix):]
	if key[0] != '/' {
		panic("invalid escrow.v1beta3 account separator")
	}

	key = key[1:]

	parts := strings.Split(string(key), "/")
	if len(parts) != 6 {
		panic("invalid escrow.v1beta3 account xid")
	}

	return eid.Payment{
		AID: eid.Account{
			Scope: eid.Scope(eid.Scope_value[parts[0]]),
			XID:   strings.Join(parts[1:2], "/"),
		},
		XID: strings.Join(parts[3:5], "/"),
	}
}

func AccountFromV1beta3(cdc codec.BinaryCodec, key []byte, val []byte) etypes.Account {
	id := AccountIDFromV1beta3(key)

	var from v1beta3.Account
	cdc.MustUnmarshal(val, &from)

	deposits := make([]etypes.Depositor, 0)

	if from.Balance.IsPositive() {
		deposits = append(deposits, etypes.Depositor{
			Owner:   from.Owner,
			Height:  0,
			Balance: from.Balance,
		})
	}

	if from.Funds.IsPositive() {
		deposits = append(deposits, etypes.Depositor{
			Owner:   from.Depositor,
			Height:  0,
			Balance: from.Funds,
		})
	}

	to := etypes.Account{
		ID: id,
		State: etypes.AccountState{
			Owner: from.Owner,
			State: etypes.State(from.State),
			Funds: []etypes.Balance{
				{
					Denom:  from.Balance.Denom,
					Amount: from.Balance.Add(from.Funds).Amount,
				},
			},
			Transferred: sdk.DecCoins{
				from.Transferred,
			},
			SettledAt: from.SettledAt,
			Deposits:  deposits,
		},
	}

	return to
}

func PaymentFromV1beta3(cdc codec.BinaryCodec, key []byte, val []byte) etypes.Payment {
	id := PaymentIDFromV1beta3(key)

	var from v1beta3.FractionalPayment
	cdc.MustUnmarshal(val, &from)

	to := etypes.Payment{
		ID: id,
		State: etypes.PaymentState{
			Owner:     from.Owner,
			State:     etypes.State(from.State),
			Rate:      from.Rate,
			Balance:   from.Balance,
			Unsettled: sdk.NewDecCoin(from.Balance.Denom, sdkmath.ZeroInt()),
			Withdrawn: from.Withdrawn,
		},
	}

	return to
}

func DepositAuthorizationFromV1beta3(from dv1beta3.DepositDeploymentAuthorization) ev1.DepositAuthorization {
	return ev1.DepositAuthorization{
		Scopes: ev1.DepositAuthorizationScopes{
			ev1.DepositScopeDeployment,
		},
		SpendLimit: from.SpendLimit,
	}
}
