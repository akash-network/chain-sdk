package v1_test

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
	"pkg.akt.dev/go/sdkutil"
)

func TestDestination(t *testing.T) {
	pub := ed25519.NewKeyFromSeed(bytes.Repeat([]byte{7}, 32)).Public().(ed25519.PublicKey)
	addr := base58.Encode(pub)
	got, err := types.ParseDestination(addr, nil)
	require.NoError(t, err)
	require.Equal(t, pub, ed25519.PublicKey(got[:]))
	for name, address := range map[string]string{
		"empty": "", "zero": "11111111111111111111111111111111", "base58": "0OIl",
		"short": "123", "long": "1" + addr, "space": " " + addr, "memo": fmt.Sprintf(`{"v":1,"solana_address":%q}`, addr),
	} {
		t.Run(name, func(t *testing.T) { _, err := types.ParseDestination(address, nil); require.Error(t, err) })
	}
	offCurve := [32]byte{2}
	require.Error(t, types.ValidateRecipient(offCurve, nil))
	got, err = types.ParseDestination(base58.Encode(offCurve[:]), [][]byte{offCurve[:]})
	require.NoError(t, err)
	require.Equal(t, offCurve, got)
	require.Error(t, types.ValidateRecipient([32]byte{}, [][]byte{make([]byte, 32)}))

	// Small-order points decode but no wallet holds their key: the identity
	// (y = 1) and the point of order 2 (y = -1).
	identity := [32]byte{1}
	require.ErrorContains(t, types.ValidateRecipient(identity, nil), "small-order")
	var order2 [32]byte
	copy(order2[:], []byte{0xec})
	for i := 1; i < 31; i++ {
		order2[i] = 0xff
	}
	order2[31] = 0x7f
	require.ErrorContains(t, types.ValidateRecipient(order2, nil), "small-order")
	// A listed vault is exempt: vaults may be any 32-byte address.
	require.NoError(t, types.ValidateRecipient(identity, [][]byte{identity[:]}))
	// y + p encodes the same point as y for y < 19. Find such a point of large
	// order and check its non-canonical spelling is rejected.
	found := false
	for y := byte(2); y < 19 && !found; y++ {
		var canonical, alias [32]byte
		canonical[0] = y
		alias[0] = 0xed + y // p = 2^255 - 19: low byte 0xed, then 0xff..., top 0x7f
		for i := 1; i < 31; i++ {
			alias[i] = 0xff
		}
		alias[31] = 0x7f
		if types.ValidateRecipient(canonical, nil) != nil {
			continue
		}
		found = true
		require.ErrorContains(t, types.ValidateRecipient(alias, nil), "canonical")
	}
	require.True(t, found, "no large-order point with a small y")
}

func TestTransferWireVector(t *testing.T) {
	// The Solana verifier parses these exact bytes: no protobuf envelope,
	// 20-byte sender, little-endian numbers and ASCII decimal keys.
	transfer := types.Transfer{ID: 42, Amount: 123456789, Height: 9001}
	copy(transfer.Sender[:], bytes.Repeat([]byte{0x11}, 20))
	copy(transfer.Recipient[:], bytes.Repeat([]byte{0x22}, 32))
	b, err := transfer.Marshal()
	require.NoError(t, err)
	want := "012a00000000000000" + "1111111111111111111111111111111111111111" +
		"2222222222222222222222222222222222222222222222222222222222222222" + "15cd5b0700000000" + "2923000000000000"
	require.Equal(t, want, hex.EncodeToString(b))
	got, err := types.DecodeTransfer(b)
	require.NoError(t, err)
	require.Equal(t, transfer, got)
	require.Equal(t, "0200000000000000002a", hex.EncodeToString(types.TransferKey(42)))
	require.Equal(t, "0201000000000000002a", hex.EncodeToString(types.VoidKey(42)))
	require.Equal(t, "0200ffffffffffffffff", hex.EncodeToString(types.TransferKey(math.MaxUint64)))
	require.Equal(t, "0201ffffffffffffffff", hex.EncodeToString(types.VoidKey(math.MaxUint64)))
	require.Equal(t, "012a00000000000000", hex.EncodeToString(types.VoidValue(42)))
	for _, bad := range [][]byte{nil, b[:76], append(bytes.Clone(b), 0), append([]byte{2}, b[1:]...)} {
		_, err := types.DecodeTransfer(bad)
		require.Error(t, err)
	}
}

func TestParamsAndGenesis(t *testing.T) {
	g := types.DefaultGenesis()
	require.NoError(t, g.Validate())
	g.Params.Enabled = true
	g.Params.Reporters = []string{sdk.AccAddress(bytes.Repeat([]byte{7}, 20)).String()}
	g.Params.ReporterThreshold = 1
	g.Params.MinTransfer = 1
	// No launch reserve: deposits need only reporters and a minimum.
	require.NoError(t, g.Validate())
	g.Params.Reporters, g.Params.ReporterThreshold = nil, 0
	require.ErrorContains(t, g.Validate(), "reporters must be configured")
	g.Params.Reporters = []string{sdk.AccAddress(bytes.Repeat([]byte{7}, 20)).String()}
	g.Params.ReporterThreshold = 1
	g.Params.MaxTransfer = g.Params.SolanaDailyMintCap + 1
	require.Error(t, g.Validate())
	g.Params = types.DefaultParams()
	g.Params.VaultAllowlist = [][]byte{make([]byte, 32)}
	require.Error(t, g.Validate())
	g.Params = types.DefaultParams()
	g.Totals.PendingAmount = 1
	require.Error(t, g.Validate())
}

func TestDefaultParams(t *testing.T) {
	p := types.DefaultParams()
	require.NoError(t, p.Validate())
	require.False(t, p.Enabled)
	require.Equal(t, uint64(1_000_000_000_000), p.MaxTransfer)         // 1M AKT
	require.Equal(t, uint64(10_000_000_000_000), p.SolanaDailyMintCap) // 10M AKT
	require.Equal(t, uint64(1_000_000_000_000), p.DailyRefundCap)      // 1M AKT
	require.GreaterOrEqual(t, p.DailyRefundCap, p.MaxTransfer)
	require.Equal(t, uint32(20), p.MaxDailyPerAccount)
	require.Equal(t, uint32(3), p.MaxPendingPerAccount)
	require.Zero(t, p.MinTransfer)
	require.Equal(t, types.DefaultChallengePeriodSeconds, p.ChallengePeriodSeconds)
	p.DailyRefundCap = p.MaxTransfer - 1
	require.ErrorContains(t, p.Validate(), "daily refund cap")
}

func TestParamsMinTransfer(t *testing.T) {
	p := types.DefaultParams()
	require.Zero(t, p.MinTransfer) // no minimum unless governance sets one
	require.NoError(t, p.Validate())
	// Parameters stored before min_transfer existed carry no field 12 and must
	// decode to the same zero minimum.
	stored, err := p.Marshal()
	require.NoError(t, err)
	var decoded types.Params
	require.NoError(t, decoded.Unmarshal(stored))
	require.Zero(t, decoded.MinTransfer)
	require.Equal(t, p, decoded)
	p.MinTransfer = p.MaxTransfer
	require.NoError(t, p.Validate())
	p.MinTransfer = p.MaxTransfer + 1
	require.ErrorContains(t, p.Validate(), "min transfer must not exceed max transfer")

	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())
	g := types.DefaultGenesis()
	g.Params.MinTransfer = 1_000_000
	require.NoError(t, g.Validate())
	bz, err := cdc.MarshalJSON(g)
	require.NoError(t, err)
	require.Contains(t, string(bz), `"min_transfer":"1000000"`)
	var imported types.GenesisState
	require.NoError(t, cdc.UnmarshalJSON(bz, &imported))
	require.Equal(t, uint64(1_000_000), imported.Params.MinTransfer)
	require.NoError(t, imported.Validate())
	again, err := cdc.MarshalJSON(&imported)
	require.NoError(t, err)
	require.JSONEq(t, string(bz), string(again))
	g.Params.MinTransfer = g.Params.MaxTransfer + 1
	require.Error(t, g.Validate())
}

func TestParamsMaxDailyAmountPerAccount(t *testing.T) {
	p := types.DefaultParams()
	require.Zero(t, p.MaxDailyAmountPerAccount) // no amount cap unless governance sets one
	require.NoError(t, p.Validate())
	// Parameters and usage stored before these fields decode to zero.
	stored, err := p.Marshal()
	require.NoError(t, err)
	var decoded types.Params
	require.NoError(t, decoded.Unmarshal(stored))
	require.Equal(t, p, decoded)
	usage := types.AccountUsage{Sender: bytes.Repeat([]byte{1}, 20), Day: 3, DailyCount: 2, PendingCount: 1}
	stored, err = usage.Marshal()
	require.NoError(t, err)
	var decodedUsage types.AccountUsage
	require.NoError(t, decodedUsage.Unmarshal(stored))
	require.Zero(t, decodedUsage.DailyAmount)

	p.MinTransfer = 100
	p.MaxDailyAmountPerAccount = 99
	require.ErrorContains(t, p.Validate(), "max daily amount per account must be zero or at least min transfer")
	p.MaxDailyAmountPerAccount = 100
	require.NoError(t, p.Validate())
	// The cap may be below max_transfer; larger deposits then always fail.
	p.MaxDailyAmountPerAccount = p.MaxTransfer + 1
	require.NoError(t, p.Validate())
	p.MaxDailyAmountPerAccount = 0
	require.NoError(t, p.Validate())

	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())
	g := types.DefaultGenesis()
	g.Params.MaxDailyAmountPerAccount = 18446744073709551615
	require.NoError(t, g.Validate())
	bz, err := cdc.MarshalJSON(g)
	require.NoError(t, err)
	require.Contains(t, string(bz), `"max_daily_amount_per_account":"18446744073709551615"`)
	var imported types.GenesisState
	require.NoError(t, cdc.UnmarshalJSON(bz, &imported))
	require.Equal(t, uint64(18446744073709551615), imported.Params.MaxDailyAmountPerAccount)
	require.NoError(t, imported.Validate())
}

func TestParamsPauseGuardian(t *testing.T) {
	p := types.DefaultParams()
	require.Empty(t, p.PauseGuardian) // no guardian unless governance sets one
	for _, guardian := range []string{
		sdk.AccAddress(bytes.Repeat([]byte{9}, 20)).String(),
		sdk.AccAddress(bytes.Repeat([]byte{9}, 32)).String(), // e.g. a group policy account
	} {
		p.PauseGuardian = guardian
		require.NoError(t, p.Validate(), guardian)
	}
	for _, guardian := range []string{
		"invalid",
		sdk.AccAddress(make([]byte, 20)).String(),
		sdk.AccAddress(bytes.Repeat([]byte{9}, 19)).String(),
		strings.ToUpper(sdk.AccAddress(bytes.Repeat([]byte{9}, 20)).String()),
		" " + sdk.AccAddress(bytes.Repeat([]byte{9}, 20)).String(),
	} {
		p.PauseGuardian = guardian
		require.ErrorContains(t, p.Validate(), "pause guardian", guardian)
	}

	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())
	g := types.DefaultGenesis()
	g.Params.PauseGuardian = sdk.AccAddress(bytes.Repeat([]byte{9}, 20)).String()
	bz, err := cdc.MarshalJSON(g)
	require.NoError(t, err)
	require.Contains(t, string(bz), `"pause_guardian":"`+g.Params.PauseGuardian+`"`)
	var imported types.GenesisState
	require.NoError(t, cdc.UnmarshalJSON(bz, &imported))
	require.Equal(t, g.Params.PauseGuardian, imported.Params.PauseGuardian)
	require.NoError(t, imported.Validate())
}
