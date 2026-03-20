package v1

import (
	"fmt"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var _ sdk.HasValidateBasic = (*Params)(nil)

// ParamKeyTable for oracle module
// Deprecated: now params can be accessed on key `0x01` on the oracle store.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// DefaultFeedContractsParams returns default feed contract params using Pyth
func DefaultFeedContractsParams() []sdk.Msg {
	return []sdk.Msg{}
}

func DefaultParams() Params {
	msgs, err := sdktx.SetMsgs(DefaultFeedContractsParams())
	if err != nil {
		panic(err.Error())
	}

	return Params{
		MinPriceSources:         1,
		MaxPriceStalenessBlocks: 2,
		MaxPriceDeviationBps:    150,
		TwapWindow:              5,
		FeedContractsParams:     msgs,
	}
}

func (p *Params) ValidateBasic() error {
	msgs, err := sdktx.GetMsgs(p.FeedContractsParams, "akash.oracle.v1.Params")
	if err != nil {
		return err
	}

	if p.MinPriceSources < 1 {
		return fmt.Errorf("min_price_sources must be at least 1")
	}
	if p.MaxPriceStalenessBlocks == 0 {
		return fmt.Errorf("max_price_staleness_blocks must be greater than 0")
	}
	if p.MaxPriceDeviationBps == 0 {
		return fmt.Errorf("max_price_deviation_bps must be greater than 0")
	}
	if p.TwapWindow == 0 {
		return fmt.Errorf("twap_window must be greater than 0")
	}

	for _, msg := range msgs {
		if m, ok := msg.(sdk.HasValidateBasic); ok {
			if err := m.ValidateBasic(); err != nil {
				return fmt.Errorf("invalid feed contract params: %w", err)
			}
		}
	}

	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage
func (p Params) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	return sdktx.UnpackInterfaces(unpacker, p.FeedContractsParams)
}
