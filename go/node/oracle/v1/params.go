package v1

import (
	"encoding/hex"
	"fmt"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var _ sdk.HasValidateBasic = (*PythContractParams)(nil)
var _ sdk.HasValidateBasic = (*WormholeContractParams)(nil)
var _ sdk.HasValidateBasic = (*Params)(nil)

// ValidateBasic validates PythContractParams
func (p *PythContractParams) ValidateBasic() error {
	if p.AktPriceFeedId == "" {
		return fmt.Errorf("akt_price_feed_id cannot be empty")
	}

	return nil
}

// ValidateBasic validates WormholeContractParams
func (p *WormholeContractParams) ValidateBasic() error {
	for i, addr := range p.GuardianAddresses {
		if len(addr) != 40 {
			return fmt.Errorf("guardian address %d must be 40 hex characters, got %d", i, len(addr))
		}
		if _, err := hex.DecodeString(addr); err != nil {
			return fmt.Errorf("guardian address %d is not valid hex: %w", i, err)
		}
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
		AktPriceFeedId: "0x4ea5bb4d2f5900cc2e97ba534240950740b4d3b89fe712a94a7304fd2fd92702",
	}
}

// DefaultWormholeContractParams returns default Wormhole contract params
// Guardian addresses are from Wormhole Mainnet Guardian Set 4
// Source: https://wormhole.com/docs/protocol/infrastructure/guardians/
func DefaultWormholeContractParams() *WormholeContractParams {
	return &WormholeContractParams{
		GuardianAddresses: []string{
			"5893B5A76c3f739645648885bDCcC06cd70a3Cd3",
			"fF6CB952589BDE862c25Ef4392132fb9D4A42157",
			"114De8460193bdf3A2fCf81f86a09765F4762fD1",
			"107A0086b32d7A0977926A205131d8731D39cbEB",
			"8C82B2fd82FaeD2711d59AF0F2499D16e726f6b2",
			"11b39756C042441BE6D8650b69b54EbE715E2343",
			"54Ce5B4D348fb74B958e8966e2ec3dBd4958a7cd",
			"15e7cAF07C4e3DC8e7C469f92C8Cd88FB8005a20",
			"74a3bf913953D695260D88BC1aA25A4eeE363ef0",
			"000aC0076727b35FBea2dAc28fEE5cCB0fEA768e",
			"AF45Ced136b9D9e24903464AE889F5C8a723FC14",
			"f93124b7c738843CBB89E864c862c38cddCccF95",
			"D2CC37A4dc036a8D232b48f62cDD4731412f4890",
			"DA798F6896A3331F64b48c12D1D57Fd9cbe70811",
			"71AA1BE1D36CaFE3867910F99C09e347899C19C3",
			"8192b6E7387CCd768277c17DAb1b7a5027c0b3Cf",
			"178e21ad2E77AE06711549CFBB1f9c7a9d8096e8",
			"5E1487F35515d02A92753504a8D75471b9f49EdB",
			"6FbEBc898F403E4773E95feB15E80C9A99c8348d",
		},
	}
}

// DefaultFeedContractsParams returns default feed contract params using Pyth
func DefaultFeedContractsParams() []sdk.Msg {
	return []sdk.Msg{
		DefaultPythContractParams(),
		DefaultWormholeContractParams(),
	}
}

func DefaultParams() Params {
	msgs, err := sdktx.SetMsgs(DefaultFeedContractsParams())
	if err != nil {
		panic(err.Error())
	}

	return Params{
		MinPriceSources:         1,
		MaxPriceStalenessBlocks: 60,
		MaxPriceDeviationBps:    150,
		TwapWindow:              180,
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
