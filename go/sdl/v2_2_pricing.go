package sdl

import (
	"errors"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gopkg.in/yaml.v3"

	types "pkg.akt.dev/go/node/types/attributes/v1"
)

// v2_2PlacementPricing supports multiple denoms per compute profile
type v2_2PlacementPricing map[string]v2_2Coins

// v2_2Coins represents multiple coin prices for a compute profile
type v2_2Coins struct {
	Values sdk.DecCoins `yaml:"-"`
}

var errInvalidMultiCoinAmount = errors.New("invalid multi-coin amount")

func (sdl *v2_2Coins) UnmarshalYAML(node *yaml.Node) error {
	// Support both single coin (legacy) and multiple coins
	if node.Kind == yaml.MappingNode {
		// Single coin: { denom: uakt, amount: 50 }
		parsedCoin := struct {
			Amount string `yaml:"amount"`
			Denom  string `yaml:"denom"`
		}{}

		if err := node.Decode(&parsedCoin); err != nil {
			return err
		}

		amount, err := math.LegacyNewDecFromStr(parsedCoin.Amount)
		if err != nil {
			return err
		}

		if amount.IsZero() {
			return fmt.Errorf("%w: amount is zero", errInvalidMultiCoinAmount)
		}

		if amount.IsNegative() {
			return fmt.Errorf("%w: amount %q is negative", errNegativeValue, amount.String())
		}

		coin := sdk.NewDecCoinFromDec(parsedCoin.Denom, amount)

		*sdl = v2_2Coins{
			Values: sdk.NewDecCoins(coin),
		}

		return nil
	}

	// Multiple coins as sequence: [{ denom: uakt, amount: 50 }, { denom: uact, amount: 100 }]
	var parsedCoins []struct {
		Amount string `yaml:"amount"`
		Denom  string `yaml:"denom"`
	}

	if err := node.Decode(&parsedCoins); err != nil {
		return err
	}

	if len(parsedCoins) == 0 {
		return fmt.Errorf("%w: no coins specified", errInvalidMultiCoinAmount)
	}

	coins := make(sdk.DecCoins, 0, len(parsedCoins))

	for _, parsedCoin := range parsedCoins {
		amount, err := math.LegacyNewDecFromStr(parsedCoin.Amount)
		if err != nil {
			return err
		}

		if amount.IsZero() {
			return fmt.Errorf("%w: amount is zero for denom %s", errInvalidMultiCoinAmount, parsedCoin.Denom)
		}

		if amount.IsNegative() {
			return fmt.Errorf("%w: amount %q is negative for denom %s", errNegativeValue, amount.String(), parsedCoin.Denom)
		}

		coin := sdk.NewDecCoinFromDec(parsedCoin.Denom, amount)
		coins = append(coins, coin)
	}

	// Sort and validate no duplicates
	coins = coins.Sort()

	*sdl = v2_2Coins{
		Values: coins,
	}

	return nil
}

// v2_2ProfilePlacement represents placement profile with multi-denom pricing
type v2_2ProfilePlacement struct {
	Attributes v2PlacementAttributes `yaml:"attributes"`
	SignedBy   types.SignedBy        `yaml:"signedBy"`
	Pricing    v2_2PlacementPricing  `yaml:"pricing"`
}

// v2_2profiles contains compute and placement profiles for SDL 2.2
type v2_2profiles struct {
	Compute   map[string]v2ProfileCompute     `yaml:"compute"`
	Placement map[string]v2_2ProfilePlacement `yaml:"placement"`
}
