package v1_test

import (
	"slices"
	"testing"

	sdkmath "cosmossdk.io/math"
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
