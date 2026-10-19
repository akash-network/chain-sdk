package cli_test

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"io"
	"math"
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtbytes "github.com/cometbft/cometbft/libs/bytes"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	rpcclientmock "github.com/cometbft/cometbft/rpc/client/mock"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"pkg.akt.dev/go/cli"
	cflags "pkg.akt.dev/go/cli/flags"
	clitestutil "pkg.akt.dev/go/cli/testutil"
	types "pkg.akt.dev/go/node/sink/v1"
	"pkg.akt.dev/go/sdkutil"
	"pkg.akt.dev/go/testutil"
)

func TestSinkQuerySDKCommand(t *testing.T) {
	enc := sdkutil.MakeEncodingConfig()
	types.RegisterInterfaces(enc.InterfaceRegistry)
	types.RegisterLegacyAminoCodec(enc.Amino)
	response := &types.QuerySinkResponse{Status: types.QuerySinkResponse_completed}
	bz, err := enc.Codec.Marshal(response)
	require.NoError(t, err)
	cctx := client.Context{}.WithCodec(enc.Codec).WithTxConfig(enc.TxConfig).
		WithLegacyAmino(enc.Amino).WithClient(testutil.NewMockCometRPC(abci.ResponseQuery{Value: bz}))
	out, err := clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), "sink", "record", "42", "--output", "json")
	require.NoError(t, err)
	var got types.QuerySinkResponse
	require.NoError(t, enc.Codec.UnmarshalJSON(out.Bytes(), &got))
	require.Equal(t, types.QuerySinkResponse_completed, got.Status)
}

// capturingCometRPC records the ABCI query the CLI sends.
type capturingCometRPC struct {
	testutil.MockCometRPC
	path *string
	data *[]byte
}

func (c capturingCometRPC) ABCIQueryWithOptions(ctx context.Context, path string, data cmtbytes.HexBytes, opts rpcclient.ABCIQueryOptions) (*coretypes.ResultABCIQuery, error) {
	*c.path, *c.data = path, data
	return c.MockCometRPC.ABCIQueryWithOptions(ctx, path, data, opts)
}

func TestSinkOpenTransfersQueryCommand(t *testing.T) {
	enc := sdkutil.MakeEncodingConfig()
	types.RegisterInterfaces(enc.InterfaceRegistry)
	types.RegisterLegacyAminoCodec(enc.Amino)
	response := &types.QueryOpenTransfersResponse{
		Transfers:      []types.TransferState{{Record: []byte{1}}, {Record: []byte{2}, VoidOrdered: true}},
		NextStartAfter: math.MaxUint64,
	}
	bz, err := enc.Codec.Marshal(response)
	require.NoError(t, err)
	var path string
	var data []byte
	rpc := capturingCometRPC{MockCometRPC: testutil.NewMockCometRPC(abci.ResponseQuery{Value: bz}), path: &path, data: &data}
	cctx := client.Context{}.WithCodec(enc.Codec).WithTxConfig(enc.TxConfig).
		WithLegacyAmino(enc.Amino).WithClient(rpc)
	out, err := clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), "sink", "open-transfers",
		"--start-after", "18446744073709551614", "--limit", "500", "--output", "json")
	require.NoError(t, err)
	require.Equal(t, "/akash.sink.v1.Query/OpenTransfers", path)
	var request types.QueryOpenTransfersRequest
	require.NoError(t, enc.Codec.Unmarshal(data, &request))
	require.Equal(t, types.QueryOpenTransfersRequest{StartAfter: math.MaxUint64 - 1, Limit: 500}, request)
	var got types.QueryOpenTransfersResponse
	require.NoError(t, enc.Codec.UnmarshalJSON(out.Bytes(), &got))
	require.Len(t, got.Transfers, 2)
	require.Equal(t, []byte{2}, got.Transfers[1].Record)
	require.True(t, got.Transfers[1].VoidOrdered)
	require.Equal(t, uint64(math.MaxUint64), got.NextStartAfter)

	// Defaults leave both fields zero; the node applies the default limit.
	_, err = clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), "sink", "open-transfers", "--output", "json")
	require.NoError(t, err)
	require.NoError(t, enc.Codec.Unmarshal(data, &request))
	require.Equal(t, types.QueryOpenTransfersRequest{}, request)
	_, err = clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), "sink", "open-transfers", "--limit", "-1")
	require.Error(t, err)
}

type SinkCLITestSuite struct {
	CLITestSuite

	from sdk.AccAddress
	dest string
}

func TestSinkCLITestSuite(t *testing.T) {
	suite.Run(t, new(SinkCLITestSuite))
}

func (s *SinkCLITestSuite) SetupSuite() {
	s.encCfg = sdkutil.MakeEncodingConfig()
	types.RegisterInterfaces(s.encCfg.InterfaceRegistry)
	types.RegisterLegacyAminoCodec(s.encCfg.Amino)
	s.kr = keyring.NewInMemory(s.encCfg.Codec)
	s.baseCtx = client.Context{}.
		WithKeyring(s.kr).
		WithTxConfig(s.encCfg.TxConfig).
		WithCodec(s.encCfg.Codec).
		WithLegacyAmino(s.encCfg.Amino).
		WithClient(testutil.MockCometRPC{Client: rpcclientmock.Client{}}).
		WithAccountRetriever(client.MockAccountRetriever{}).
		WithOutput(io.Discard).
		WithChainID("test-chain").
		WithSignModeStr(cflags.SignModeDirect)
	s.cctx = s.baseCtx
	ctx := context.WithValue(context.Background(), cli.ContextTypeAddressCodec, s.encCfg.SigningOptions.AddressCodec)
	s.ctx = context.WithValue(ctx, cli.ContextTypeValidatorCodec, s.encCfg.SigningOptions.ValidatorAddressCodec)
	s.from = sdktestutil.CreateKeyringAccounts(s.T(), s.kr, 1)[0].Address
	s.dest = base58.Encode(ed25519.NewKeyFromSeed(bytes.Repeat([]byte{7}, 32)).Public().(ed25519.PublicKey))
}

// Amounts are exact: AKT decimals convert to uakt, and fractional uakt fails
// instead of being truncated.
func (s *SinkCLITestSuite) TestDepositAmountIsExact() {
	for _, tc := range []struct {
		amount string
		want   int64
		err    string
	}{
		{"1500000uakt", 1_500_000, ""},
		{"1.5akt", 1_500_000, ""},
		{"0.000001akt", 1, ""},
		{"1.5uakt", 0, "not a whole number of uakt"},
		{"0.0000015akt", 0, "not a whole number of uakt"},
		{"0uakt", 0, "positive uint64 uakt"},
		{"-1uakt", 0, "unknown shorthand flag"},
		{"1e6uakt", 0, "positive uint64 uakt"},
		{"1other", 0, "positive uint64 uakt"},
		{"18446744073709551616uakt", 0, "positive uint64 uakt"},
	} {
		s.Run(tc.amount, func() {
			out, err := clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.GetTxSinkDepositCmd(),
				cli.TestFlags().With(s.dest, tc.amount).WithFrom(s.from.String()).WithGenerateOnly()...)
			if tc.err != "" {
				s.Require().ErrorContains(err, tc.err)
				return
			}
			s.Require().NoError(err)
			tx, err := s.encCfg.TxConfig.TxJSONDecoder()(out.Bytes())
			s.Require().NoError(err)
			s.Require().Len(tx.GetMsgs(), 1)
			msg, ok := tx.GetMsgs()[0].(*types.MsgSink)
			s.Require().True(ok)
			s.Require().Equal(sdk.NewInt64Coin(types.Denom, tc.want), msg.Amount)
			s.Require().Equal(s.from.String(), msg.Sender)
		})
	}
}

// ValidateBasic runs before generation or broadcast, so malformed destinations
// fail locally.
func (s *SinkCLITestSuite) TestDepositRejectsMalformedDestination() {
	for _, dest := range []string{"11111111111111111111111111111111", "not-base58", "1" + s.dest} {
		_, err := clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.GetTxSinkDepositCmd(),
			cli.TestFlags().With(dest, "1uakt").WithFrom(s.from.String()).WithGenerateOnly()...)
		s.Require().Error(err, dest)
	}
}

// guardian-pause signs from --from and requires at least one switch.
func (s *SinkCLITestSuite) TestGuardianPause() {
	for _, tc := range []struct {
		flags             []string
		deposits, refunds bool
	}{
		{[]string{"--deposits"}, true, false},
		{[]string{"--refunds"}, false, true},
		{[]string{"--deposits", "--refunds"}, true, true},
	} {
		out, err := clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.GetTxSinkGuardianPauseCmd(),
			cli.TestFlags().With(tc.flags...).WithFrom(s.from.String()).WithGenerateOnly()...)
		s.Require().NoError(err)
		tx, err := s.encCfg.TxConfig.TxJSONDecoder()(out.Bytes())
		s.Require().NoError(err)
		s.Require().Len(tx.GetMsgs(), 1)
		msg, ok := tx.GetMsgs()[0].(*types.MsgGuardianPause)
		s.Require().True(ok)
		s.Require().Equal(types.MsgGuardianPause{Guardian: s.from.String(), Deposits: tc.deposits, Refunds: tc.refunds}, *msg)
	}
	_, err := clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.GetTxSinkGuardianPauseCmd(),
		cli.TestFlags().WithFrom(s.from.String()).WithGenerateOnly()...)
	s.Require().ErrorContains(err, "pass --deposits, --refunds or both")
	_, err = clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.TxCmd(),
		cli.TestFlags().With("sink", "guardian-pause", "--deposits").WithFrom(s.from.String()).WithGenerateOnly()...)
	s.Require().NoError(err)
}
