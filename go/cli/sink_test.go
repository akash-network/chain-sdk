package cli_test

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"io"
	"strconv"
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
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
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
	id := types.ID{
		Owner:    sdk.AccAddress(bytes.Repeat([]byte{7}, 20)).String(),
		Height:   9001,
		Sequence: 42,
	}
	response := &types.QueryTransferResponse{Transfer: types.TransferState{
		ID:     id,
		Data:   types.RecordData{Amount: sdk.NewInt64Coin(sdkutil.DenomUakt, 1)},
		Status: types.StatusMinted,
	}}
	bz, err := enc.Codec.Marshal(response)
	require.NoError(t, err)
	cctx := client.Context{}.WithCodec(enc.Codec).WithTxConfig(enc.TxConfig).WithLegacyAmino(enc.Amino).WithClient(testutil.NewMockCometRPC(abci.ResponseQuery{Value: bz}))
	out, err := clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), "sink", "transfer", id.String(), "--output", "json")
	require.NoError(t, err)
	var got types.QueryTransferResponse
	require.NoError(t, enc.Codec.UnmarshalJSON(out.Bytes(), &got))
	require.Equal(t, id, got.Transfer.ID)
	_, err = clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), "sink", "transfer", "42")
	require.Error(t, err)
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

func TestSinkTransfersQueryCommand(t *testing.T) {
	enc := sdkutil.MakeEncodingConfig()
	types.RegisterInterfaces(enc.InterfaceRegistry)
	response := &types.QueryTransfersResponse{Pagination: &query.PageResponse{NextKey: []byte{4}}}
	bz, err := enc.Codec.Marshal(response)
	require.NoError(t, err)
	var path string
	var data []byte
	rpc := capturingCometRPC{
		MockCometRPC: testutil.NewMockCometRPC(abci.ResponseQuery{Value: bz}),
		path:         &path,
		data:         &data,
	}
	cctx := client.Context{}.WithCodec(enc.Codec).WithTxConfig(enc.TxConfig).WithLegacyAmino(enc.Amino).WithClient(rpc)
	owner := sdk.AccAddress(bytes.Repeat([]byte{7}, 20)).String()
	out, err := clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), "sink", "transfers", "--owner", owner, "--status", "STATUS_REFUND_PENDING", "--page-key", base64.StdEncoding.EncodeToString([]byte{1, 2}), "--limit", "500", "--height", "123", "--output", "json")
	require.NoError(t, err)
	require.Equal(t, "/akash.sink.v1.Query/Transfers", path)
	var request types.QueryTransfersRequest
	require.NoError(t, enc.Codec.Unmarshal(data, &request))
	require.Equal(t, owner, request.Owner)
	require.Equal(t, types.StatusRefundPending, request.Status)
	require.Equal(t, []byte{1, 2}, request.Pagination.Key)
	var got types.QueryTransfersResponse
	require.NoError(t, enc.Codec.UnmarshalJSON(out.Bytes(), &got))
	require.Equal(t, []byte{4}, got.Pagination.NextKey)

	for _, args := range [][]string{{"--offset", "1"}, {"--count-total"}, {"--status", "STATUS_UNKNOWN"}} {
		path, data = "", nil
		_, err = clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), append([]string{"sink", "transfers"}, args...)...)
		require.Error(t, err)
		require.Empty(t, path)
		require.Empty(t, data)
	}

	_, err = clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), "sink", "disputes")
	require.NoError(t, err)
	require.Equal(t, "/akash.sink.v1.Query/Disputes", path)
}

func TestSinkQueryPagination(t *testing.T) {
	enc := sdkutil.MakeEncodingConfig()
	types.RegisterInterfaces(enc.InterfaceRegistry)
	var path string
	var data []byte
	rpc := capturingCometRPC{
		MockCometRPC: testutil.NewMockCometRPC(abci.ResponseQuery{Value: []byte{}}),
		path:         &path,
		data:         &data,
	}
	cctx := client.Context{}.WithCodec(enc.Codec).WithTxConfig(enc.TxConfig).WithLegacyAmino(enc.Amino).WithClient(rpc)

	for _, command := range []string{"transfers", "disputes"} {
		t.Run(command, func(t *testing.T) {
			for _, args := range [][]string{
				{"--page", "2"},
				{"--page", "2", "--limit", "0"},
				{"--offset", "0"},
				{"--count-total=false"},
			} {
				path, data = "", nil
				_, err := clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), append([]string{"sink", command}, args...)...)
				require.ErrorContains(t, err, "unknown flag", args)
				require.Empty(t, path)
				require.Empty(t, data)
			}

			for _, tc := range []struct {
				args []string
				want query.PageRequest
			}{
				{
					want: query.PageRequest{Limit: 100},
				},
				{
					args: []string{"--limit", "0"},
					want: query.PageRequest{},
				},
				{
					args: []string{"--page-key", "AQI=", "--limit", "500", "--reverse", "--height", "123"},
					want: query.PageRequest{
						Key:     []byte{1, 2},
						Limit:   500,
						Reverse: true,
					},
				},
			} {
				_, err := clitestutil.ExecTestCLICmd(context.Background(), cctx, cli.QueryCmd(), append([]string{"sink", command}, tc.args...)...)
				require.NoError(t, err)

				var pagination *query.PageRequest
				if command == "transfers" {
					var request types.QueryTransfersRequest
					require.NoError(t, enc.Codec.Unmarshal(data, &request))
					pagination = request.Pagination
				} else {
					var request types.QueryDisputesRequest
					require.NoError(t, enc.Codec.Unmarshal(data, &request))
					pagination = request.Pagination
				}

				require.Equal(t, &tc.want, pagination)
			}
		})
	}
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
		err    error
	}{
		{"1500000uakt", 1_500_000, nil},
		{"1.5akt", 1_500_000, nil},
		{"0.000001akt", 1, nil},
		{"1.5uakt", 0, sdkerrors.ErrInvalidCoins},
		{"0.0000015akt", 0, sdkerrors.ErrInvalidCoins},
		{"0uakt", 0, sdkerrors.ErrInvalidCoins},
		{"1e6uakt", 0, sdkerrors.ErrInvalidCoins},
		{"1other", 0, sdkerrors.ErrInvalidCoins},
		{"18446744073709551616uakt", 0, sdkerrors.ErrInvalidCoins},
	} {
		s.Run(tc.amount, func() {
			out, err := clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.GetTxSinkDepositCmd(),
				cli.TestFlags().With(s.dest, tc.amount).WithFrom(s.from.String()).WithGenerateOnly()...)
			if tc.err != nil {
				s.Require().ErrorIs(err, tc.err)
				return
			}

			s.Require().NoError(err)
			tx, err := s.encCfg.TxConfig.TxJSONDecoder()(out.Bytes())
			s.Require().NoError(err)
			s.Require().Len(tx.GetMsgs(), 1)
			msg, ok := tx.GetMsgs()[0].(*types.MsgSink)
			s.Require().True(ok)
			s.Require().Equal(sdk.NewInt64Coin(sdkutil.DenomUakt, tc.want), msg.Amount)
			s.Require().Equal(s.from.String(), msg.Sender)
			s.Require().Equal(types.DestTypeSolana, msg.DestType)
			s.Require().Equal(s.dest, msg.Dest)
		})
	}

	// A negative amount parses as a flag, so it never reaches the command.
	_, err := clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.GetTxSinkDepositCmd(),
		cli.TestFlags().With(s.dest, "-1uakt").WithFrom(s.from.String()).WithGenerateOnly()...)
	s.Require().ErrorContains(err, "unknown shorthand flag")
}

// ValidateBasic runs before generation or broadcast, so malformed destinations
// fail locally.
func (s *SinkCLITestSuite) TestDepositRejectsMalformedDestination() {
	for _, dest := range []string{"11111111111111111111111111111111", "not-base58", "1" + s.dest} {
		_, err := clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.GetTxSinkDepositCmd(),
			cli.TestFlags().With(dest, "1uakt").WithFrom(s.from.String()).WithGenerateOnly()...)
		s.Require().ErrorIs(err, types.ErrInvalidDestination, dest)
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
		s.Require().Equal(types.MsgGuardianPause{
			Guardian: s.from.String(),
			Deposits: tc.deposits,
			Refunds:  tc.refunds,
		}, *msg)
	}

	_, err := clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.GetTxSinkGuardianPauseCmd(),
		cli.TestFlags().WithFrom(s.from.String()).WithGenerateOnly()...)
	s.Require().ErrorIs(err, types.ErrInvalidGuardianPause)
	_, err = clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.TxCmd(),
		cli.TestFlags().With("sink", "guardian-pause", "--deposits").WithFrom(s.from.String()).WithGenerateOnly()...)
	s.Require().NoError(err)
}

func (s *SinkCLITestSuite) TestGovernanceCommandsBindExpectedRound() {
	id := types.ID{
		Owner:    s.from.String(),
		Height:   5,
		Sequence: 7,
	}
	signature := base58.Encode(bytes.Repeat([]byte{8}, 64))
	for _, command := range []string{"restart-transfer", "confirm-mint", "confirm-void"} {
		for _, round := range []string{"1", "0", "01", "+1", "18446744073709551615", "18446744073709551616"} {
			args := []string{"sink", command, id.String(), round}
			if command != "restart-transfer" {
				args = append(args, "9", signature)
			}

			args = append(args, "reviewed", "--authority", s.from.String())
			out, err := clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.TxCmd(), cli.TestFlags().With(args...).WithFrom(s.from.String()).WithGenerateOnly()...)
			if round != "1" {
				s.Require().Error(err)
				continue
			}

			s.Require().NoError(err)
			tx, err := s.encCfg.TxConfig.TxJSONDecoder()(out.Bytes())
			s.Require().NoError(err)
			expected, _ := strconv.ParseUint(round, 10, 64)
			switch msg := tx.GetMsgs()[0].(type) {
			case *types.MsgRestartTransfer:
				s.Require().Equal(expected, msg.ExpectedRound)
				s.Require().Equal(id, msg.ID)
			case *types.MsgConfirmMint:
				s.Require().Equal(expected, msg.ExpectedRound)
				s.Require().Equal(uint64(9), msg.Evidence.Slot)
			case *types.MsgConfirmVoid:
				s.Require().Equal(expected, msg.ExpectedRound)
				s.Require().Equal(uint64(9), msg.Evidence.Slot)
			default:
				s.T().Fatalf("unexpected message %T", msg)
			}
		}
	}
}

func (s *SinkCLITestSuite) TestRequestVoidBindsExpectedRound() {
	id := types.ID{
		Owner:    s.from.String(),
		Height:   5,
		Sequence: 7,
	}
	for _, round := range []string{"1", "18446744073709551615", "0", "01", "+1", "18446744073709551616"} {
		out, err := clitestutil.ExecTestCLICmd(s.ctx, s.cctx, cli.TxCmd(),
			cli.TestFlags().With("sink", "request-void", id.String(), round, "--authority", s.from.String()).WithFrom(s.from.String()).WithGenerateOnly()...)
		if round != "1" && round != "18446744073709551615" {
			s.Require().Error(err)
			continue
		}

		s.Require().NoError(err)
		tx, err := s.encCfg.TxConfig.TxJSONDecoder()(out.Bytes())
		s.Require().NoError(err)
		msg, ok := tx.GetMsgs()[0].(*types.MsgRequestVoid)
		s.Require().True(ok)
		expected, err := strconv.ParseUint(round, 10, 64)
		s.Require().NoError(err)
		s.Require().Equal(types.MsgRequestVoid{
			Authority:     s.from.String(),
			ID:            id,
			ExpectedRound: expected,
		}, *msg)
	}
}
