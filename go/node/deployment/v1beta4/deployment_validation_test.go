package v1beta4_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1 "pkg.akt.dev/go/node/deployment/v1"
	types "pkg.akt.dev/go/node/deployment/v1beta4"
	attr "pkg.akt.dev/go/node/types/attributes/v1"
	akashtypes "pkg.akt.dev/go/node/types/resources/v1beta4"
	tutil "pkg.akt.dev/go/testutil"
)

const (
	regexInvalidUnitBoundaries = `^.*invalid unit count|CPU|GPU|memory|storage \(\d+ > 0 > \d+ fails\)$`
)

func TestZeroValueGroupSpec(t *testing.T) {
	did := tutil.DeploymentID(t)

	dgroup := tutil.DeploymentGroup(t, did, uint32(6))
	gspec := dgroup.GroupSpec

	t.Run("assert nominal test success", func(t *testing.T) {
		err := gspec.ValidateBasic()
		require.NoError(t, err)
	})
}

func TestZeroValueGroupSpecs(t *testing.T) {
	did := tutil.DeploymentID(t)
	dgroups := tutil.DeploymentGroups(t, did, uint32(6))
	gspecs := make([]types.GroupSpec, 0)
	for _, d := range dgroups {
		gspecs = append(gspecs, d.GroupSpec)
	}

	t.Run("assert nominal test success", func(t *testing.T) {
		err := types.ValidateDeploymentGroups(gspecs)
		require.NoError(t, err)
	})

	gspecZeroed := make([]types.GroupSpec, len(gspecs))
	gspecZeroed = append(gspecZeroed, gspecs...)
	t.Run("assert error for zero value bid duration", func(t *testing.T) {
		err := types.ValidateDeploymentGroups(gspecZeroed)
		require.Error(t, err)
	})
}

func TestEmptyGroupSpecIsInvalid(t *testing.T) {
	err := types.ValidateDeploymentGroups(make([]types.GroupSpec, 0))
	require.Equal(t, v1.ErrInvalidGroups, err)
}

func validSimpleGroupSpec() types.GroupSpec {
	resources := make(types.ResourceUnits, 1)
	resources[0] = types.ResourceUnit{
		Resources: akashtypes.Resources{
			ID: 1,
			CPU: &akashtypes.CPU{
				Units: akashtypes.ResourceValue{
					Val: sdkmath.NewInt(10),
				},
				Attributes: nil,
			},
			GPU: &akashtypes.GPU{
				Units: akashtypes.ResourceValue{
					Val: sdkmath.NewInt(0),
				},
				Attributes: nil,
			},
			Memory: &akashtypes.Memory{
				Quantity: akashtypes.ResourceValue{
					Val: sdkmath.NewIntFromUint64(types.GetValidationConfig().Unit.Min.Memory),
				},
				Attributes: nil,
			},
			Storage: akashtypes.Volumes{
				akashtypes.Storage{
					Quantity: akashtypes.ResourceValue{
						Val: sdkmath.NewIntFromUint64(types.GetValidationConfig().Unit.Min.Storage),
					},
					Attributes: nil,
				},
			},
			Endpoints: akashtypes.Endpoints{},
		},
		Count: 1,
		Price: sdk.NewInt64DecCoin(tutil.CoinDenom, 1),
	}
	return types.GroupSpec{
		Name:         "testGroup",
		Requirements: attr.PlacementRequirements{},
		Resources:    resources,
	}
}

func validSimpleGroupSpecs() []types.GroupSpec {
	result := make([]types.GroupSpec, 1)
	result[0] = validSimpleGroupSpec()

	return result
}

func TestSimpleGroupSpecIsValid(t *testing.T) {
	groups := validSimpleGroupSpecs()
	err := types.ValidateDeploymentGroups(groups)
	require.NoError(t, err)
}

func TestDuplicateSimpleGroupSpecIsInvalid(t *testing.T) {
	groups := validSimpleGroupSpecs()
	groupsDuplicate := make([]types.GroupSpec, 2)
	groupsDuplicate[0] = groups[0]
	groupsDuplicate[1] = groups[0]
	err := types.ValidateDeploymentGroups(groupsDuplicate)
	require.Error(t, err) // TODO - specific error
	require.Regexp(t, "^.*duplicate.*$", err)
}

func TestGroupWithZeroCount(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Count = 0
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, regexInvalidUnitBoundaries, err)
}

func TestGroupWithZeroCPU(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].CPU.Units.Val = sdkmath.NewInt(0)
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, regexInvalidUnitBoundaries, err)
}

func TestGroupWithZeroMemory(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Memory.Quantity.Val = sdkmath.NewInt(0)
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, regexInvalidUnitBoundaries, err)
}

func TestGroupWithZeroStorage(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Storage[0].Quantity.Val = sdkmath.NewInt(0)
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, regexInvalidUnitBoundaries, err)
}

func TestGroupWithNilCPU(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].CPU = nil
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid unit CPU.*$", err)
}

func TestGroupWithNilGPU(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].GPU = nil
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid unit GPU.*$", err)
}

func TestGroupWithNilMemory(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Memory = nil
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid unit memory.*$", err)
}

func TestGroupWithNilStorage(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Storage = nil
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid unit storage.*$", err)
}

func TestGroupWithInvalidPrice(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Price = sdk.DecCoin{}
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid price object.*$", err)
}

func TestGroupWithNegativePrice(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Price.Amount = sdkmath.LegacyNewDec(-1)
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid price object.*$", err)
}
