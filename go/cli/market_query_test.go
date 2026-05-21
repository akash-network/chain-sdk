package cli

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQueryMarketIncludesProviderLeaseStats(t *testing.T) {
	root := GetQueryMarketCmds()

	var statsCmdFound bool
	for _, cmd := range root.Commands() {
		if cmd.Name() == "provider-lease-stats" {
			statsCmdFound = true
			require.NotNil(t, cmd.Flags().Lookup("since"))
			break
		}
	}

	require.True(t, statsCmdFound)
}

func TestReadSinceFlag(t *testing.T) {
	cmd := GetQueryMarketProviderLeaseStatsCmd()

	since, err := readSinceFlag(cmd)
	require.NoError(t, err)
	require.True(t, since.IsZero())

	expected := "2026-05-21T12:34:56Z"
	require.NoError(t, cmd.Flags().Set("since", expected))

	since, err = readSinceFlag(cmd)
	require.NoError(t, err)
	require.Equal(t, expected, since.UTC().Format(time.RFC3339))

	require.NoError(t, cmd.Flags().Set("since", "not-a-time"))
	_, err = readSinceFlag(cmd)
	require.Error(t, err)
}
