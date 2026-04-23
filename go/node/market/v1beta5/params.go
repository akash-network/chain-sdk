package v1beta5

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	v1 "pkg.akt.dev/go/node/market/v1"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	DefaultBidMinDeposit           = sdk.NewCoin("uakt", sdkmath.NewInt(500000))
	DefaultBidMinDepositACT        = sdk.NewCoin("uact", sdkmath.NewInt(500000))
	defaultOrderMaxBids     uint32 = 20
	maxOrderMaxBids         uint32 = 500

	DefaultMinReclamationWindow = 1 * time.Hour
	DefaultMaxReclamationWindow = 720 * time.Hour // 30 days
)

const (
	keyBidMinDeposit  = "BidMinDeposit"
	keyBidMinDeposits = "BidMinDeposits"
	keyOrderMaxBids   = "OrderMaxBids"
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair([]byte(keyBidMinDeposit), &p.BidMinDeposit, validateCoin),
		paramtypes.NewParamSetPair([]byte(keyBidMinDeposits), &p.BidMinDeposits, validateCoins),
		paramtypes.NewParamSetPair([]byte(keyOrderMaxBids), &p.OrderMaxBids, validateOrderMaxBids),
	}
}

func DefaultParams() Params {
	return Params{
		BidMinDeposit: DefaultBidMinDeposit,
		OrderMaxBids:  defaultOrderMaxBids,
		BidMinDeposits: sdk.NewCoins(
			DefaultBidMinDeposit,
			DefaultBidMinDepositACT,
		),
		MinReclamationWindow: DefaultMinReclamationWindow,
		MaxReclamationWindow: DefaultMaxReclamationWindow,
	}
}

func (p Params) Validate() error {
	if err := validateCoin(p.BidMinDeposit); err != nil {
		return err
	}

	if err := validateOrderMaxBids(p.OrderMaxBids); err != nil {
		return err
	}

	if err := validateCoins(p.BidMinDeposits); err != nil {
		return err
	}

	if err := validateReclamationWindows(p.MinReclamationWindow, p.MaxReclamationWindow); err != nil {
		return err
	}

	return nil
}

func validateCoin(i interface{}) error {
	_, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("%w: invalid type %T", v1.ErrInvalidParam, i)
	}

	return nil
}

func validateCoins(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("%w: invalid type %T", v1.ErrInvalidParam, i)
	}

	if err := v.Validate(); err != nil {
		return fmt.Errorf("%w: %s", v1.ErrInvalidParam, err)
	}

	return nil
}

func validateReclamationWindows(min, max time.Duration) error {
	if min < 0 {
		return fmt.Errorf("%w: min reclamation window must be >= 0", v1.ErrInvalidParam)
	}

	if max <= 0 {
		return fmt.Errorf("%w: max reclamation window must be > 0", v1.ErrInvalidParam)
	}

	if max <= min {
		return fmt.Errorf("%w: max reclamation window must be > min reclamation window", v1.ErrInvalidParam)
	}

	return nil
}

func validateOrderMaxBids(i interface{}) error {
	val, ok := i.(uint32)

	if !ok {
		return fmt.Errorf("%w: invalid type %T", v1.ErrInvalidParam, i)
	}

	if val == 0 {
		return fmt.Errorf("%w: order max bids too low", v1.ErrInvalidParam)
	}

	if val > maxOrderMaxBids {
		return fmt.Errorf("%w: order max bids too high", v1.ErrInvalidParam)
	}

	return nil
}
