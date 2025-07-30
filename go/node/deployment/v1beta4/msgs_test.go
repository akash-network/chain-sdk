package v1beta4_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v1 "pkg.akt.dev/go/node/deployment/v1"
	types "pkg.akt.dev/go/node/deployment/v1beta4"
	"pkg.akt.dev/go/testutil"
)

type testMsg struct {
	msg sdk.Msg
	err error
}

func TestVersionValidation(t *testing.T) {
	tests := []testMsg{
		{
			msg: &types.MsgCreateDeployment{
				ID:   testutil.DeploymentID(t),
				Hash: testutil.DeploymentVersion(t),
				Groups: types.GroupSpecs{
					testutil.GroupSpec(t),
				},
				Depositor: testutil.AccAddress(t).String(),
				Deposit:   testutil.AkashCoin(t, 0),
			},
			err: nil,
		},
		{
			msg: &types.MsgCreateDeployment{
				ID:   testutil.DeploymentID(t),
				Hash: []byte(""),
				Groups: []types.GroupSpec{
					testutil.GroupSpec(t),
				},
				Depositor: testutil.AccAddress(t).String(),
				Deposit:   testutil.AkashCoin(t, 0),
			},
			err: v1.ErrEmptyHash,
		},
		{
			msg: &types.MsgCreateDeployment{
				ID:   testutil.DeploymentID(t),
				Hash: []byte("invalidversion"),
				Groups: []types.GroupSpec{
					testutil.GroupSpec(t),
				},
				Depositor: testutil.AccAddress(t).String(),
				Deposit:   testutil.AkashCoin(t, 0),
			},
			err: v1.ErrInvalidHash,
		},
		{
			msg: &types.MsgUpdateDeployment{
				ID:   testutil.DeploymentID(t),
				Hash: testutil.DeploymentVersion(t),
			},
			err: nil,
		},
		{
			msg: &types.MsgUpdateDeployment{
				ID:   testutil.DeploymentID(t),
				Hash: []byte(""),
			},
			err: v1.ErrEmptyHash,
		},
		{
			msg: &types.MsgUpdateDeployment{
				ID:   testutil.DeploymentID(t),
				Hash: []byte("invalidversion"),
			},
			err: v1.ErrInvalidHash,
		},
	}

	for _, test := range tests {
		m, ok := test.msg.(sdk.HasValidateBasic)
		require.True(t, ok)
		require.Equal(t, test.err, m.ValidateBasic())
	}
}
