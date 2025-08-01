package cli

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/x/upgrade/types"

	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"

	cflags "pkg.akt.dev/go/cli/flags"
)

func TestParsePlan(t *testing.T) {
	fs := NewCmdSubmitUpgradeProposal(addresscodec.NewBech32Codec("akash")).Flags()

	proposal := types.MsgSoftwareUpgrade{
		Plan: types.Plan{
			Name:   "plan name",
			Height: 123456,
			Info:   "plan info",
		},
	}

	require.NoError(t, fs.Set(cflags.FlagUpgradeHeight, strconv.FormatInt(proposal.Plan.Height, 10)))
	require.NoError(t, fs.Set(cflags.FlagUpgradeInfo, proposal.Plan.Info))

	p, err := parsePlan(fs, proposal.Plan.Name)
	require.NoError(t, err)
	require.Equal(t, p.Name, proposal.Plan.Name)
	require.Equal(t, p.Height, proposal.Plan.Height)
	require.Equal(t, p.Info, proposal.Plan.Info)
}
