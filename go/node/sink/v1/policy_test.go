package v1_test

import (
	"slices"
	"strings"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
)

func TestReporterPolicyEqual(t *testing.T) {
	a, b := policy(), policy()
	b.Version = 2
	b.Threshold = sdkmath.LegacyMustNewDecFromStr("0.750000000000000000")
	require.True(t, a.Equal(b))

	b.Members[0], b.Members[3] = b.Members[3], b.Members[0]
	require.False(t, a.Equal(b))
	slices.Sort(b.Members)
	require.True(t, a.Equal(b))

	b.Threshold = sdkmath.LegacyMustNewDecFromStr("0.7")
	require.False(t, a.Equal(b))
	b.Threshold = sdkmath.LegacyDec{}
	require.False(t, a.Equal(b))
	require.False(t, b.Equal(a))
}

func TestReporterPolicyDecimalValidation(t *testing.T) {
	for _, threshold := range []sdkmath.LegacyDec{
		{},
		sdkmath.LegacyMustNewDecFromStr("-0.1"),
		sdkmath.LegacyMustNewDecFromStr("1.000000000000000001"),
		sdkmath.LegacyZeroDec(),
		sdkmath.LegacyMustNewDecFromStr("0.5"),
	} {
		p := policy()
		p.Threshold = threshold
		require.Error(t, p.Validate())
	}

	for _, threshold := range []string{"0.500000000000000001", "0.66", "0.75", "1"} {
		p := policy()
		p.Threshold = sdkmath.LegacyMustNewDecFromStr(threshold)
		require.NoError(t, p.Validate(), threshold)
	}

	require.NoError(t, types.ValidateReporters(nil, sdkmath.LegacyZeroDec()))
}

// Bech32 also decodes the all-upper-case spelling of an address, so two
// spellings of one account are a duplicate reporter.
func TestValidateReportersDedupesDecodedAddresses(t *testing.T) {
	members := policy().Members
	upper := strings.ToUpper(members[0])
	alias, err := sdk.AccAddressFromBech32(upper)
	require.NoError(t, err)
	require.Equal(t, members[0], alias.String(), "both spellings name one account")

	// Four strings at 0.66 need three votes, which two accounts would cast.
	threshold := sdkmath.LegacyMustNewDecFromStr("0.66")
	duplicate := []string{members[0], upper, members[1], members[2]}
	require.ErrorIs(t, types.ValidateReporters(duplicate, threshold), types.ErrInvalidReporters)

	p := policy()
	p.Members = duplicate
	p.Threshold = threshold
	require.ErrorIs(t, p.Validate(), types.ErrInvalidReporters)

	// One spelling per account is valid in either case.
	require.NoError(t, types.ValidateReporters([]string{upper, members[1], members[2], members[3]}, threshold))
}
