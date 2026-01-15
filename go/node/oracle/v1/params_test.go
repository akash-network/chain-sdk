package v1

import (
	"testing"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
)

func TestPythContractParams_ValidateBasic(t *testing.T) {
	tests := []struct {
		name    string
		params  *PythContractParams
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid pyth params",
			params: &PythContractParams{
				AktPriceFeedId: "0x1c5d745dc0e0c8a0034b6c3d3a8e5d34e4e9b79c9ab2f4b3e6a8e7f0c9e8a5b4",
			},
			wantErr: false,
		},
		{
			name: "empty akt_price_feed_id",
			params: &PythContractParams{
				AktPriceFeedId: "",
			},
			wantErr: true,
			errMsg:  "akt_price_feed_id cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.ValidateBasic()
			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestParams_Validate(t *testing.T) {
	tests := []struct {
		name    string
		params  Params
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid default params",
			params:  DefaultParams(),
			wantErr: false,
		},
		{
			name: "valid params with nil feed contract params",
			params: Params{
				FeedContractsParams: nil,
			},
			wantErr: false,
		},
		{
			name: "invalid feed contract params",
			params: Params{
				FeedContractsParams: []FeedContractParams{
					{
						Name:   "",
						Config: nil,
					},
				},
			},
			wantErr: true,
			errMsg:  "invalid feed contract params",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDefaultParams(t *testing.T) {
	params := DefaultParams()

	// Verify default params are valid
	err := params.Validate()
	require.NoError(t, err)

	// Verify feed contract params are set
	require.NotNil(t, params.FeedContractsParams)
	require.Equal(t, FeedContractNamePyth, params.FeedContractsParams[0].Name)
	require.NotNil(t, params.FeedContractsParams[0].Config)
}

func TestDefaultPythContractParams(t *testing.T) {
	params := DefaultPythContractParams()

	require.NotNil(t, params)
	require.NotEmpty(t, params.AktPriceFeedId)
	require.Equal(t, "0x1c5d745dc0e0c8a0034b6c3d3a8e5d34e4e9b79c9ab2f4b3e6a8e7f0c9e8a5b4", params.AktPriceFeedId)

	// Validate
	err := params.ValidateBasic()
	require.NoError(t, err)
}

func TestNewFeedContractParams(t *testing.T) {
	pythParams := &PythContractParams{
		AktPriceFeedId: "test-feed-id",
	}

	params, err := NewFeedContractParams(FeedContractNamePyth, pythParams)
	require.NoError(t, err)
	require.NotNil(t, params)
	require.Equal(t, FeedContractNamePyth, params.Name)
	require.NotNil(t, params.Config)
}

func TestFeedContractParams_GetFeedContractConfig(t *testing.T) {
	t.Run("valid config", func(t *testing.T) {
		pythParams := &PythContractParams{
			AktPriceFeedId: "test-feed-id",
		}
		params, err := NewFeedContractParams(FeedContractNamePyth, pythParams)
		require.NoError(t, err)

		config, err := params.GetFeedContractConfig()
		require.NoError(t, err)
		require.NotNil(t, config)

		// Cast to PythContractParams and verify
		pyth, ok := config.(*PythContractParams)
		require.True(t, ok)
		require.Equal(t, "test-feed-id", pyth.AktPriceFeedId)
	})

	t.Run("nil params", func(t *testing.T) {
		var params *FeedContractParams
		config, err := params.GetFeedContractConfig()
		require.Error(t, err)
		require.Nil(t, config)
		require.Contains(t, err.Error(), "feed contract config is nil")
	})

	t.Run("nil config", func(t *testing.T) {
		params := &FeedContractParams{
			Name:   FeedContractNamePyth,
			Config: nil,
		}
		config, err := params.GetFeedContractConfig()
		require.Error(t, err)
		require.Nil(t, config)
		require.Contains(t, err.Error(), "feed contract config is nil")
	})
}

func TestParams_UnpackInterfaces(t *testing.T) {
	// Create a registry and register the interface
	registry := cdctypes.NewInterfaceRegistry()
	RegisterInterfaces(registry)

	// Create default params
	params := DefaultParams()

	// Unpack should succeed
	err := params.UnpackInterfaces(registry)
	require.NoError(t, err)

	// Nil feed contract params should also succeed
	paramsWithNil := Params{}
	err = paramsWithNil.UnpackInterfaces(registry)
	require.NoError(t, err)
}

func TestFeedContractParams_UnpackInterfaces(t *testing.T) {
	// Create a registry and register the interface
	registry := cdctypes.NewInterfaceRegistry()
	RegisterInterfaces(registry)

	pythParams := &PythContractParams{
		AktPriceFeedId: "test-feed-id",
	}
	params, err := NewFeedContractParams(FeedContractNamePyth, pythParams)
	require.NoError(t, err)

	// Unpack should succeed
	err = params.UnpackInterfaces(registry)
	require.NoError(t, err)

	// Nil config should also succeed
	paramsWithNil := &FeedContractParams{
		Name:   FeedContractNamePyth,
		Config: nil,
	}
	err = paramsWithNil.UnpackInterfaces(registry)
	require.NoError(t, err)
}
