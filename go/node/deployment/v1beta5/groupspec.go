package v1beta5

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	atypes "pkg.akt.dev/go/node/audit/v1"
	attr "pkg.akt.dev/go/node/types/attributes/v1"
)

type ResourceGroup interface {
	GetName() string
	GetResourceUnits() ResourceUnits
}

var _ ResourceGroup = (*GroupSpec)(nil)

type GroupSpecs []GroupSpec

func (gspecs GroupSpecs) Dup() GroupSpecs {
	res := make(GroupSpecs, 0, len(gspecs))

	for _, gspec := range gspecs {
		gs := gspec.Dup()
		res = append(res, gs)
	}
	return res
}

func (g GroupSpec) Dup() GroupSpec {
	res := GroupSpec{
		Name:         g.Name,
		Requirements: g.Requirements.Dup(),
		Resources:    g.Resources.Dup(),
	}

	return res
}

// ValidateBasic asserts non-zero values
func (g GroupSpec) ValidateBasic() error {
	return g.validate()
}

// GetResourceUnits method returns resources list in group
func (g GroupSpec) GetResourceUnits() ResourceUnits {
	resources := make(ResourceUnits, 0, len(g.Resources))

	resources = append(resources, g.Resources...)

	return resources
}

// GetName method returns group name
func (g GroupSpec) GetName() string {
	return g.Name
}

// Prices method returns prices group can be paid in
func (g GroupSpec) Prices() sdk.DecCoins {
	var prices sdk.DecCoins

	prices.Add()
	for idx, resource := range g.Resources {
		if idx == 0 {
			prices = resource.FullPrice()
			continue
		}
		prices = prices.Add(resource.FullPrice()...)
	}

	prices.Sort()
	return prices
}

// MatchResourcesRequirements check if resources attributes match provider's capabilities
func (g GroupSpec) MatchResourcesRequirements(pattr attr.Attributes) bool {
	for _, rgroup := range g.GetResourceUnits() {
		pgroup := pattr.GetCapabilitiesGroup("storage")
		for _, storage := range rgroup.Storage {
			if len(storage.Attributes) == 0 {
				continue
			}

			if !storage.Attributes.IN(pgroup) {
				return false
			}
		}
		if gpu := rgroup.GPU; gpu.Units.Val.Uint64() > 0 {
			attr := gpu.Attributes
			if len(attr) == 0 {
				continue
			}

			pgroup = pattr.GetCapabilitiesMap("gpu")

			if !gpu.Attributes.AnyIN(pgroup) {
				return false
			}
		}
	}

	return true
}

// MatchRequirements method compares provided attributes with specific group attributes.
// Argument provider is a bit cumbersome. First element is attributes from x/provider store
// in case tenant does not need signed attributes at all
// rest of elements (if any) are attributes signed by various auditors
func (g GroupSpec) MatchRequirements(provider []atypes.AuditedProvider) bool {
	if (len(g.Requirements.SignedBy.AnyOf) != 0) || (len(g.Requirements.SignedBy.AllOf) != 0) {
		// we cannot match if there is no signed attributes
		if len(provider) < 2 {
			return false
		}

		existingRequirements := make(attributesMatching)

		for _, existing := range provider[1:] {
			existingRequirements[existing.Auditor] = existing.Attributes
		}

		if len(g.Requirements.SignedBy.AllOf) != 0 {
			for _, validator := range g.Requirements.SignedBy.AllOf {
				// if at least one signature does not exist or no match on attributes - requirements cannot match
				if existingAttr, exists := existingRequirements[validator]; !exists ||
					!attr.AttributesSubsetOf(g.Requirements.Attributes, existingAttr) {
					return false
				}
			}
		}

		if len(g.Requirements.SignedBy.AnyOf) != 0 {
			for _, validator := range g.Requirements.SignedBy.AnyOf {
				if existingAttr, exists := existingRequirements[validator]; exists &&
					attr.AttributesSubsetOf(g.Requirements.Attributes, existingAttr) {
					return true
				}
			}

			return false
		}

		return true
	}

	return attr.AttributesSubsetOf(g.Requirements.Attributes, provider[0].Attributes)
}

// validate does validation for provided deployment group
func (g *GroupSpec) validate() error {
	if g.Name == "" {
		return fmt.Errorf("empty group spec name denomination")
	}

	if err := g.GetResourceUnits().Validate(); err != nil {
		return err
	}

	if err := g.validatePricing(); err != nil {
		return err
	}

	return nil
}

func (g *GroupSpec) validatePricing() error {
	var prices sdk.DecCoins

	for idx, resource := range g.Resources {
		if err := resource.validatePricing(); err != nil {
			return fmt.Errorf("group %v: %w", g.GetName(), err)
		}

		// all groups must be the same denominations
		if idx == 0 {
			prices = resource.FullPrice()

			if len(prices) == 0 {
				return errors.New("invalid price object")
			}

		} else {
			rprice := resource.FullPrice()

			if len(prices) != len(rprice) {
				return errors.New("mismatching group denominations")
			}

			// FullPrice call returns sorted Coins
			for i := range prices {
				if prices[i].Denom != rprice[i].Denom {
					return fmt.Errorf("mismatching group denominations: (%v == %v fails)", prices[i].Denom, rprice[i].Denom)
				}
			}
		}
	}

	return nil
}
