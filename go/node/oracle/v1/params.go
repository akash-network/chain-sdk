package v1

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable for take module
// Deprecated: now params can be accessed on key `0x01` on the take store.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func DefaultParams() Params {
	return Params{
		AktPriceFeedId: "0x1c5d745dc0e0c8a0034b6c3d3a8e5d34e4e9b79c9ab2f4b3e6a8e7f0c9e8a5b4",
	}
}

func (p Params) Validate() error {
	return nil
}
