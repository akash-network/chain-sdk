package v1beta5

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1 "pkg.akt.dev/go/node/deployment/v1"
)

// ValidateDeploymentGroups does validation for all deployment groups
func ValidateDeploymentGroups(gspecs []GroupSpec) error {
	if len(gspecs) == 0 {
		return v1.ErrInvalidGroups
	}

	var prices sdk.DecCoins

	names := make(map[string]int, len(gspecs)) // Used as set
	for idx, group := range gspecs {
		// all must be the same denomination
		if idx == 0 {
			prices = group.Prices()
			if len(prices) == 0 {
				return errors.New("group pricing cannot be empty")
			}
		} else {
			rprice := group.Prices()

			if len(prices) != len(rprice) {
				return errors.New("inconsistent denominations")
			}

			// FullPrice call returns sorted Coins
			for i := range prices {
				if prices[i].Denom != rprice[i].Denom {
					return fmt.Errorf("inconsistent denominations: (%v == %v)", prices[i].Denom, rprice[i].Denom)
				}
			}
		}

		if err := group.ValidateBasic(); err != nil {
			return err
		}

		if _, exists := names[group.GetName()]; exists {
			return fmt.Errorf("duplicate deployment group name %q", group.GetName())
		}

		names[group.GetName()] = 0 // Value stored does not matter
	}

	return nil
}
