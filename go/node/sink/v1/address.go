package v1

import (
	"bytes"
	"slices"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// AccountLength is the length of an ordinary account address.
	AccountLength = 20

	// DerivedAccountLength is the length of a derived account address, such as a group policy.
	DerivedAccountLength = 32
)

// ParseAccount parses a canonical bech32 account address and checks it with
// ValidateAccount. Canonical means lower case, so accounts compare as strings.
func ParseAccount(addr string, lengths ...int) (sdk.AccAddress, error) {
	acc, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("%s: %s", addr, err)
	}

	if acc.String() != addr {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("%s is not canonical", addr)
	}

	if err := ValidateAccount(acc, lengths...); err != nil {
		return nil, err
	}

	return acc, nil
}

// ValidateAccount checks that an account address is nonzero and one of the given lengths.
func ValidateAccount(acc sdk.AccAddress, lengths ...int) error {
	if !slices.Contains(lengths, len(acc)) {
		return sdkerrors.ErrInvalidAddress.Wrapf("%d-byte address, expected %v bytes", len(acc), lengths)
	}

	if bytes.Equal(acc, make([]byte, len(acc))) {
		return sdkerrors.ErrInvalidAddress.Wrap("zero address")
	}

	return nil
}
