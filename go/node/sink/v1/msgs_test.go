package v1_test

import (
	"bytes"
	"crypto/ed25519"
	"strings"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
	"pkg.akt.dev/go/sdkutil"
)

func TestMsgSinkValidateBasic(t *testing.T) {
	pub := ed25519.NewKeyFromSeed(bytes.Repeat([]byte{7}, 32)).Public().(ed25519.PublicKey)
	valid := func() *types.MsgSink {
		return &types.MsgSink{
			DestType: types.DestTypeSolana,
			Dest:     base58.Encode(pub),
			Amount:   sdk.NewInt64Coin(sdkutil.DenomUakt, 1),
			Sender:   sdk.AccAddress(bytes.Repeat([]byte{7}, 20)).String(),
		}
	}

	require.NoError(t, valid().ValidateBasic())

	// Off-curve keys may be allowlisted vaults; only the keeper knows the allowlist.
	offCurve := valid()
	offCurve.Dest = base58.Encode(append([]byte{2}, make([]byte, 31)...))
	require.NoError(t, offCurve.ValidateBasic())
	overflow, ok := sdkmath.NewIntFromString("18446744073709551616")
	require.True(t, ok)

	for name, tc := range map[string]struct {
		change func(*types.MsgSink)
		err    error
	}{
		"zero destination type":    {func(m *types.MsgSink) { m.DestType = types.DestTypeInvalid }, types.ErrInvalidDestination},
		"unknown destination type": {func(m *types.MsgSink) { m.DestType = types.DestTypeSolana + 1 }, types.ErrInvalidDestination},
		"empty destination":        {func(m *types.MsgSink) { m.Dest = "" }, types.ErrInvalidDestination},
		"zero destination":         {func(m *types.MsgSink) { m.Dest = "11111111111111111111111111111111" }, types.ErrInvalidDestination},
		"noncanonical destination": {func(m *types.MsgSink) { m.Dest = "1" + m.Dest }, types.ErrInvalidDestination},
		"memo destination":         {func(m *types.MsgSink) { m.Dest = `{"solana_address":"x"}` }, types.ErrInvalidDestination},
		"nil amount":               {func(m *types.MsgSink) { m.Amount.Amount = sdkmath.Int{} }, sdkerrors.ErrInvalidCoins},
		"zero amount":              {func(m *types.MsgSink) { m.Amount.Amount = sdkmath.ZeroInt() }, sdkerrors.ErrInvalidCoins},
		"negative amount":          {func(m *types.MsgSink) { m.Amount.Amount = sdkmath.NewInt(-1) }, sdkerrors.ErrInvalidCoins},
		"overflow amount":          {func(m *types.MsgSink) { m.Amount.Amount = overflow }, sdkerrors.ErrInvalidCoins},
		"wrong denomination":       {func(m *types.MsgSink) { m.Amount.Denom = "akt" }, sdkerrors.ErrInvalidCoins},
		"empty sender":             {func(m *types.MsgSink) { m.Sender = "" }, sdkerrors.ErrInvalidAddress},
		"malformed sender":         {func(m *types.MsgSink) { m.Sender = "akash1notanaddress" }, sdkerrors.ErrInvalidAddress},
		"long sender":              {func(m *types.MsgSink) { m.Sender = sdk.AccAddress(bytes.Repeat([]byte{7}, 32)).String() }, sdkerrors.ErrInvalidAddress},
	} {
		t.Run(name, func(t *testing.T) {
			msg := valid()
			tc.change(msg)
			require.ErrorIs(t, msg.ValidateBasic(), tc.err)
		})
	}

	var msg sdk.Msg = valid()
	_, ok = msg.(sdk.HasValidateBasic)
	require.True(t, ok)
}

func TestMsgGuardianPauseValidateBasic(t *testing.T) {
	guardian := sdk.AccAddress(bytes.Repeat([]byte{9}, 20)).String()
	for _, msg := range []types.MsgGuardianPause{
		{
			Guardian: guardian,
			Deposits: true,
		},
		{
			Guardian: guardian,
			Refunds:  true,
		},
		{
			Guardian: guardian,
			Deposits: true,
			Refunds:  true,
		},
		{
			Guardian: sdk.AccAddress(bytes.Repeat([]byte{9}, 32)).String(),
			Deposits: true,
		},
	} {
		require.NoError(t, msg.ValidateBasic())
	}

	require.ErrorIs(t, (&types.MsgGuardianPause{Guardian: guardian}).ValidateBasic(), types.ErrInvalidGuardianPause)

	for name, msg := range map[string]types.MsgGuardianPause{
		"empty guardian": {Deposits: true},
		"malformed": {
			Guardian: "akash1notanaddress",
			Deposits: true,
		},
		"zero address": {
			Guardian: sdk.AccAddress(make([]byte, 20)).String(),
			Deposits: true,
		},
		"21-byte address": {
			Guardian: sdk.AccAddress(bytes.Repeat([]byte{9}, 21)).String(),
			Deposits: true,
		},
		"uppercase bech32": {
			Guardian: strings.ToUpper(guardian),
			Refunds:  true,
		},
	} {
		t.Run(name, func(t *testing.T) { require.ErrorIs(t, msg.ValidateBasic(), sdkerrors.ErrInvalidAddress) })
	}

	var msg sdk.Msg = &types.MsgGuardianPause{}
	_, ok := msg.(sdk.HasValidateBasic)
	require.True(t, ok)
}

func TestEveryCommandValidatesStatelessRules(t *testing.T) {
	id := types.ID{
		Owner:    account(7),
		Height:   8,
		Sequence: 9,
	}
	authority := account(8)
	for _, msg := range []sdk.HasValidateBasic{
		&types.MsgUpdateParams{
			Authority: authority,
			Params:    types.DefaultParams(),
		},
		&types.MsgRequestVoid{
			Authority:     authority,
			ID:            id,
			ExpectedRound: 1,
		},
		&types.MsgReportOutcome{
			Reporter:      account(1),
			ID:            id,
			PolicyVersion: 1,
			Round:         1,
			Observation: types.Observation{
				Outcome:              types.OutcomeMinted,
				Slot:                 evidence().Slot,
				TransactionSignature: evidence().TransactionSignature,
			},
		},
		&types.MsgVetoVoid{
			Signer: authority,
			ID:     id,
			Reason: "verify",
		},
		&types.MsgRestartTransfer{
			Authority:     authority,
			ID:            id,
			Reason:        "retry",
			ExpectedRound: 1,
		},
		&types.MsgConfirmMint{
			Authority:     authority,
			ID:            id,
			Evidence:      evidence(),
			Reason:        "minted",
			ExpectedRound: 1,
		},
		&types.MsgConfirmVoid{
			Authority:     authority,
			ID:            id,
			Evidence:      evidence(),
			Reason:        "voided",
			ExpectedRound: 1,
		},
		&types.MsgFinalizeRefund{
			Executor: authority,
			ID:       id,
		},
	} {
		require.NoError(t, msg.ValidateBasic(), "%T", msg)
	}

	for _, msg := range []sdk.HasValidateBasic{
		&types.MsgUpdateParams{}, &types.MsgRequestVoid{}, &types.MsgReportOutcome{}, &types.MsgVetoVoid{},
		&types.MsgRestartTransfer{}, &types.MsgConfirmMint{}, &types.MsgConfirmVoid{}, &types.MsgFinalizeRefund{},
	} {
		require.Error(t, msg.ValidateBasic(), "%T", msg)
	}

	for _, round := range []uint64{0, ^uint64(0)} {
		require.Error(t, (&types.MsgRestartTransfer{
			Authority:     authority,
			ID:            id,
			Reason:        "retry",
			ExpectedRound: round,
		}).ValidateBasic())
		require.Error(t, (&types.MsgConfirmMint{
			Authority:     authority,
			ID:            id,
			Evidence:      evidence(),
			Reason:        "minted",
			ExpectedRound: round,
		}).ValidateBasic())
		require.Error(t, (&types.MsgConfirmVoid{
			Authority:     authority,
			ID:            id,
			Evidence:      evidence(),
			Reason:        "voided",
			ExpectedRound: round,
		}).ValidateBasic())
	}

	for _, e := range []types.SolanaEvidence{{}, {
		Slot:                 1,
		TransactionSignature: make([]byte, 64),
	}, {
		Slot:                 1,
		TransactionSignature: []byte{1},
	}} {
		require.Error(t, (&types.MsgConfirmMint{
			Authority:     authority,
			ID:            id,
			Evidence:      e,
			Reason:        "minted",
			ExpectedRound: 1,
		}).ValidateBasic())
	}

	for _, reason := range []string{"", " \t", strings.Repeat("a", 257), string([]byte{0xff})} {
		require.Error(t, types.ValidateReason(reason))
	}

	require.NoError(t, types.ValidateReason(strings.Repeat("a", 256)))
}
