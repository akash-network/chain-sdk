package v1

import (
	"fmt"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/gogoproto/proto"
)

const (
	// FeedContractNamePyth is the name identifier for Pyth price feed contracts
	FeedContractNamePyth = "pyth"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// FeedContractConfig is the interface that all feed contract configurations must implement
type FeedContractConfig interface {
	proto.Message
	ValidateBasic() error
}

// Ensure PythContractParams implements FeedContractConfig
var _ FeedContractConfig = (*PythContractParams)(nil)

// ValidateBasic validates PythContractParams
func (p *PythContractParams) ValidateBasic() error {
	if p.AktPriceFeedId == "" {
		return fmt.Errorf("akt_price_feed_id cannot be empty")
	}
	return nil
}

// ParamKeyTable for oracle module
// Deprecated: now params can be accessed on key `0x01` on the oracle store.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// DefaultPythContractParams returns default Pyth contract params
func DefaultPythContractParams() *PythContractParams {
	return &PythContractParams{
		AktPriceFeedId: "0x1c5d745dc0e0c8a0034b6c3d3a8e5d34e4e9b79c9ab2f4b3e6a8e7f0c9e8a5b4",
	}
}

// NewFeedContractParams creates a new FeedContractParams with the given name and config
func NewFeedContractParams(name string, config FeedContractConfig) (FeedContractParams, error) {
	any, err := cdctypes.NewAnyWithValue(config)
	if err != nil {
		return FeedContractParams{}, err
	}
	return FeedContractParams{
		Name:   name,
		Config: any,
	}, nil
}

// DefaultFeedContractsParams returns default feed contract params using Pyth
func DefaultFeedContractsParams() []FeedContractParams {
	params, _ := NewFeedContractParams(FeedContractNamePyth, DefaultPythContractParams())
	return []FeedContractParams{params}
}

func DefaultParams() Params {
	return Params{
		FeedContractsParams: DefaultFeedContractsParams(),
	}
}

// GetFeedContractConfig returns the feed contract config from the Any field
func (p *FeedContractParams) GetFeedContractConfig() (FeedContractConfig, error) {
	if p == nil || p.Config == nil {
		return nil, fmt.Errorf("feed contract config is nil")
	}

	cachedValue := p.Config.GetCachedValue()
	if cachedValue == nil {
		return nil, fmt.Errorf("feed contract config cached value is nil")
	}

	config, ok := cachedValue.(FeedContractConfig)
	if !ok {
		return nil, fmt.Errorf("expected FeedContractConfig, got %T", cachedValue)
	}
	return config, nil
}

// Validate validates FeedContractParams
func (p *FeedContractParams) Validate() error {
	if p == nil {
		return fmt.Errorf("feed contract params cannot be nil")
	}

	if p.Name == "" {
		return fmt.Errorf("feed contract name cannot be empty")
	}

	if p.Config == nil {
		return fmt.Errorf("feed contract config cannot be nil")
	}

	config, err := p.GetFeedContractConfig()
	if err != nil {
		return fmt.Errorf("failed to get feed contract config: %w", err)
	}

	return config.ValidateBasic()
}

func (p Params) Validate() error {
	for _, feed := range p.FeedContractsParams {
		if err := feed.Validate(); err != nil {
			return fmt.Errorf("invalid feed contract params: %w", err)
		}
	}

	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage
func (p *Params) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	for i := range p.FeedContractsParams {
		if err := p.FeedContractsParams[i].UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}

	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage for FeedContractParams
func (p *FeedContractParams) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	if p.Config == nil {
		return nil
	}
	var config FeedContractConfig
	return unpacker.UnpackAny(p.Config, &config)
}
