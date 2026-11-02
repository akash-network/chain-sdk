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
)

func TestMsgSinkValidateBasic(t *testing.T) {
	pub := ed25519.NewKeyFromSeed(bytes.Repeat([]byte{7}, 32)).Public().(ed25519.PublicKey)
	valid := func() *types.MsgSink {
		return &types.MsgSink{DestType: 1, Dest: base58.Encode(pub), Amount: sdk.NewInt64Coin(types.Denom, 1),
			Sender: sdk.AccAddress(bytes.Repeat([]byte{7}, 20)).String()}
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
		"zero destination type":    {func(m *types.MsgSink) { m.DestType = 0 }, types.ErrInvalidDestination},
		"unknown destination type": {func(m *types.MsgSink) { m.DestType = 2 }, types.ErrInvalidDestination},
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
		{Guardian: guardian, Deposits: true},
		{Guardian: guardian, Refunds: true},
		{Guardian: guardian, Deposits: true, Refunds: true},
		{Guardian: sdk.AccAddress(bytes.Repeat([]byte{9}, 32)).String(), Deposits: true},
	} {
		require.NoError(t, msg.ValidateBasic())
	}
	require.ErrorIs(t, (&types.MsgGuardianPause{Guardian: guardian}).ValidateBasic(), types.ErrInvalidGuardianPause)
	for name, msg := range map[string]types.MsgGuardianPause{
		"empty guardian":   {Deposits: true},
		"malformed":        {Guardian: "akash1notanaddress", Deposits: true},
		"zero address":     {Guardian: sdk.AccAddress(make([]byte, 20)).String(), Deposits: true},
		"21-byte address":  {Guardian: sdk.AccAddress(bytes.Repeat([]byte{9}, 21)).String(), Deposits: true},
		"uppercase bech32": {Guardian: strings.ToUpper(guardian), Refunds: true},
	} {
		t.Run(name, func(t *testing.T) { require.ErrorIs(t, msg.ValidateBasic(), sdkerrors.ErrInvalidAddress) })
	}
	var msg sdk.Msg = &types.MsgGuardianPause{}
	_, ok := msg.(sdk.HasValidateBasic)
	require.True(t, ok)
}
