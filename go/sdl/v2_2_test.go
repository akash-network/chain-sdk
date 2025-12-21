package sdl

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV2_2_Parse_Simple_MultiDenom(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.2-simple-multi-denom.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	assert.Len(t, groups, 1)

	group := groups[0]
	assert.Len(t, group.Resources, 1)

	resource := group.Resources[0]
	assert.Equal(t, uint32(2), resource.Count)

	// Check multi-denom prices
	assert.Len(t, resource.Prices, 2)

	expectedPrices := sdk.NewDecCoins(
		sdk.NewInt64DecCoin("uact", 100),
		sdk.NewInt64DecCoin("uakt", 50),
	)

	assert.True(t, expectedPrices.Equal(resource.Prices), "prices should match: expected %v, got %v", expectedPrices, resource.Prices)
}

func TestV2_2_Parse_Simple_SingleDenom(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.2-simple-single-denom.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	assert.Len(t, groups, 1)

	group := groups[0]
	assert.Len(t, group.Resources, 1)

	resource := group.Resources[0]
	assert.Equal(t, uint32(1), resource.Count)

	// Check single denom price (backward compatibility)
	assert.Len(t, resource.Prices, 1)

	expectedPrice := sdk.NewInt64DecCoin("uakt", 50)
	assert.True(t, resource.Prices.Equal(sdk.NewDecCoins(expectedPrice)), "prices should match: expected %v, got %v", expectedPrice, resource.Prices)
}

func TestV2_2_MultiDenom_Parsing(t *testing.T) {
	tests := []struct {
		name          string
		yaml          string
		expectedCoins sdk.DecCoins
		expectError   bool
	}{
		{
			name: "single denom",
			yaml: `
version: "2.2"
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - global: true
profiles:
  compute:
    web:
      resources:
        cpu:
          units: "100m"
        memory:
          size: "128Mi"
        storage:
        - size: "1Gi"
  placement:
    westcoast:
      pricing:
        web:
          denom: uakt
          amount: 50
deployment:
  web:
    westcoast:
      profile: web
      count: 1
`,
			expectedCoins: sdk.NewDecCoins(sdk.NewInt64DecCoin("uakt", 50)),
			expectError:   false,
		},
		{
			name: "multiple denoms",
			yaml: `
version: "2.2"
services:
  web:
    image: nginx
    expose:
      - port: 80
        to:
          - global: true
profiles:
  compute:
    web:
      resources:
        cpu:
          units: "100m"
        memory:
          size: "128Mi"
        storage:
        - size: "1Gi"
  placement:
    westcoast:
      pricing:
        web:
          - denom: uakt
            amount: 50
          - denom: uact
            amount: 100
          - denom: usdc
            amount: 0.1
deployment:
  web:
    westcoast:
      profile: web
      count: 1
`,
			expectedCoins: sdk.NewDecCoins(
				sdk.NewInt64DecCoin("uact", 100),
				sdk.NewInt64DecCoin("uakt", 50),
				sdk.NewDecCoinFromDec("usdc", math.LegacyMustNewDecFromStr("0.1")),
			),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdl, err := Read([]byte(tt.yaml))

			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			groups, err := sdl.DeploymentGroups()
			require.NoError(t, err)
			require.Len(t, groups, 1)

			resource := groups[0].Resources[0]
			assert.True(t, tt.expectedCoins.Equal(resource.Prices), "prices should match: expected %v, got %v", tt.expectedCoins, resource.Prices)
		})
	}
}
