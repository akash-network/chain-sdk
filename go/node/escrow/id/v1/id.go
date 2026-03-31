package v1

import (
	"bytes"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.akt.dev/go/node/escrow/module"
)

type ID interface {
	Key() string
}

var (
	_ ID = (*Account)(nil)
	_ ID = (*Payment)(nil)
)

func ParseAccount(key string) (Account, error) {
	parts := strings.Split(key, "/")
	if len(parts) < 3 {
		return Account{}, module.ErrMalformedKey.Wrapf("malformed account key %s", key)
	}

	scopeVal, valid := Scope_value[parts[0]]
	if !valid {
		return Account{}, module.ErrMalformedKey.Wrapf("invalid account scope \"%s\"", parts[0])
	}

	parts = parts[1:]

	scope := Scope(scopeVal)

	switch scope {
	case ScopeDeployment:
		if len(parts) != 2 {
			return Account{}, module.ErrMalformedKey.Wrapf("malformed account key for deployment scope \"%s\"", key)
		}
	case ScopeBid:
		if len(parts) != 5 {
			return Account{}, module.ErrMalformedKey.Wrapf("malformed account key for bid scope \"%s\"", key)
		}
	default:
		return Account{}, module.ErrMalformedKey.Wrapf("invalid account scope \"%s\"", key)
	}

	acc := Account{
		Scope: scope,
		XID:   strings.Join(parts, "/"),
	}

	if err := acc.ValidateBasic(); err != nil {
		return Account{}, module.ErrMalformedKey.Wrap(err.Error())
	}

	return acc, nil
}

func ParsePayment(key string) (Payment, error) {
	parts := strings.Split(key, "/")

	if len(parts) < 6 {
		return Payment{}, module.ErrMalformedKey.Wrapf("malformed payment key %s", key)
	}

	scope, valid := Scope_value[parts[0]]
	if !valid || Scope(scope) == ScopeInvalid {
		return Payment{}, module.ErrMalformedKey.Wrapf("invalid payment scope \"%s\"", parts[0])
	}

	var aidParts int

	switch Scope(scope) {
	case ScopeDeployment:
		// scope/owner/dseq + gseq/oseq/provider = 6 parts
		if len(parts) != 6 {
			return Payment{}, module.ErrMalformedKey.Wrapf("malformed payment key for deployment scope \"%s\"", key)
		}
		aidParts = 2
	case ScopeBid:
		// scope/owner/dseq/gseq/oseq/provider + gseq/oseq/provider = 9 parts
		if len(parts) != 9 {
			return Payment{}, module.ErrMalformedKey.Wrapf("malformed payment key for bid scope \"%s\"", key)
		}
		aidParts = 5
	default:
		return Payment{}, module.ErrMalformedKey.Wrapf("invalid payment scope \"%s\"", key)
	}

	pmt := Payment{
		AID: Account{
			Scope: Scope(scope),
			XID:   strings.Join(parts[1:1+aidParts], "/"),
		},
		XID: strings.Join(parts[1+aidParts:], "/"),
	}

	if err := pmt.ValidateBasic(); err != nil {
		return Payment{}, module.ErrMalformedKey.Wrap(err.Error())
	}

	return pmt, nil
}

func (obj *Account) Key() string {
	buf := &bytes.Buffer{}

	buf.WriteString(Scope_name[int32(obj.Scope)])
	buf.WriteRune('/')
	buf.WriteString(obj.XID)

	return buf.String()
}

func (obj *Payment) Key() string {
	buf := &bytes.Buffer{}

	buf.WriteString(obj.AID.Key())
	buf.WriteRune('/')
	buf.WriteString(obj.XID)

	return buf.String()
}

func (obj *Account) ValidateBasic() error {
	parts := strings.Split(obj.XID, "/")

	switch obj.Scope {
	case ScopeDeployment:
		if len(parts) != 2 {
			return module.ErrInvalidID.Wrap("invalid xid format")
		}
	case ScopeBid:
		if len(parts) != 5 {
			return module.ErrInvalidID.Wrap("invalid xid format")
		}
	default:
		return module.ErrInvalidID.Wrap("invalid scope")
	}

	_, err := sdk.AccAddressFromBech32(parts[0])
	if err != nil {
		return module.ErrInvalidID.Wrapf("invalid xid/owner: %s", err.Error())
	}

	_, err = strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return module.ErrInvalidID.Wrapf("invalid xid/dseq: %s", err.Error())
	}

	if obj.Scope == ScopeBid {
		parts = parts[2:]
		err = validateBidScope(parts)
		if err != nil {
			return err
		}
	}

	return nil
}

func (obj *Payment) ValidateBasic() error {
	err := obj.AID.ValidateBasic()
	if err != nil {
		return err
	}

	parts := strings.Split(obj.XID, "/")
	if len(parts) != 3 {
		return module.ErrInvalidID.Wrap("invalid xid format")
	}

	err = validateBidScope(parts)
	if err != nil {
		return err
	}

	return nil
}

func validateBidScope(parts []string) error {
	_, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return module.ErrInvalidID.Wrapf("invalid xid/gseq: %s", err.Error())
	}

	_, err = strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return module.ErrInvalidID.Wrapf("invalid xid/oseq: %s", err.Error())
	}

	_, err = sdk.AccAddressFromBech32(parts[2])
	if err != nil {
		return module.ErrInvalidID.Wrapf("invalid xid/provider: %s", err.Error())
	}

	return nil
}

func (obj *Payment) Account() Account {
	return obj.AID
}
