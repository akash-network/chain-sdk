package v2

import (
	"fmt"
	"time"

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
		MaxPriceStalenessPeriod: 30 * time.Second,
		MaxPriceDeviationBps:    150,
		TwapWindow:              5 * time.Second,
		FeedContractsParams:     msgs,
		PriceRetention:          24 * time.Hour,
		PruneEpoch:              "hour",
		MaxPrunePerEpoch:        1000,
		MaxFutureTimeDrift:      10 * time.Second,
	}
}

func (p *Params) ValidateBasic() error {
	msgs, err := sdktx.GetMsgs(p.FeedContractsParams, "akash.oracle.v2.Params")
	if err != nil {
		return err
	}

	if p.MinPriceSources < 1 {
		return fmt.Errorf("min_price_sources must be at least 1")
	}
	if p.MaxPriceStalenessPeriod <= 0 {
		return fmt.Errorf("max_price_staleness_period must be greater than 0")
	}
	if p.MaxPriceDeviationBps == 0 {
		return fmt.Errorf("max_price_deviation_bps must be greater than 0")
	}
	if p.TwapWindow <= 0 {
		return fmt.Errorf("twap_window must be greater than 0")
	}

	if p.PriceRetention <= 0 {
		return fmt.Errorf("price_retention must be positive")
	}
	if p.PriceRetention < p.TwapWindow {
		return fmt.Errorf("price_retention (%s) must be at least twap_window (%s)", p.PriceRetention, p.TwapWindow)
	}
	if p.PruneEpoch == "" {
		return fmt.Errorf("prune_epoch must not be empty")
	}
	if p.MaxPrunePerEpoch < 0 {
		return fmt.Errorf("max_prune_per_epoch must be non-negative")
	}
	if p.MaxFutureTimeDrift < 0 {
		return fmt.Errorf("max_future_time_drift must be non-negative")
	}

	for _, src := range p.Sources {
		if _, err := sdk.AccAddressFromBech32(src); err != nil {
			return fmt.Errorf("invalid source address %q: %w", src, err)
		}
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
