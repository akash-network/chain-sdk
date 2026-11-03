package v1_test

import (
	"bytes"
	"context"
	"encoding/json"
	"math"
	"strings"
	"testing"
	"time"

	basev1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	sdkmath "cosmossdk.io/math"
	txsigning "cosmossdk.io/x/tx/signing"
	"cosmossdk.io/x/tx/signing/aminojson"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	gogoproto "github.com/cosmos/gogoproto/proto"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	types "pkg.akt.dev/go/node/sink/v1"
	"pkg.akt.dev/go/sdkutil"
)

func TestSinkAminoJSONParity(t *testing.T) {
	encoding := sdkutil.MakeEncodingConfig()
	types.RegisterInterfaces(encoding.InterfaceRegistry)
	testSinkAminoJSONParity(t, encoding.TxConfig.SignModeHandler())
}

func TestSinkAminoJSONParityWithResolvedDescriptors(t *testing.T) {
	// Rebuilding the descriptor graph isolates field encoding from init order.
	registry, err := gogoproto.MergedRegistry()
	require.NoError(t, err)
	handler := aminojson.NewSignModeHandler(aminojson.SignModeHandlerOptions{FileResolver: registry})
	testSinkAminoJSONParity(t, txsigning.NewHandlerMap(handler))
}

func TestSinkChallengePeriodSigningBoundary(t *testing.T) {
	encoding := sdkutil.MakeEncodingConfig()
	lastSecond := time.Duration(aminojson.MaxDurationSeconds) * time.Second
	for _, tc := range []struct {
		name     string
		duration time.Duration
		valid    bool
	}{
		{"last whole second", lastSecond, true},
		{"last nanosecond", lastSecond + time.Second - time.Nanosecond, true},
		{"first unsupported nanosecond", lastSecond + time.Second, false},
		{"maximum Go duration", time.Duration(math.MaxInt64), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			params := types.DefaultParams()
			params.ReporterChallengePeriod = tc.duration
			msg := &types.MsgUpdateParams{Authority: account(8), Params: params}
			genesis := types.DefaultGenesis()
			genesis.Params = params
			for _, err := range []error{params.ValidateBasic(), msg.ValidateBasic(), genesis.Validate()} {
				if tc.valid {
					require.NoError(t, err)
				} else {
					require.ErrorIs(t, err, types.ErrInvalidParam)
				}
			}

			builder := encoding.TxConfig.NewTxBuilder()
			require.NoError(t, builder.SetMsgs(msg))
			_, err := authsigning.GetSignBytesAdapter(context.Background(), encoding.TxConfig.SignModeHandler(),
				signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
				authsigning.SignerData{Address: msg.Authority},
				builder.GetTx())
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, "seconds would overflow an int64")
			}
		})
	}
}

func testSinkAminoJSONParity(t *testing.T, handler *txsigning.HandlerMap) {
	t.Helper()

	amino := codec.NewLegacyAmino()
	std.RegisterLegacyAminoCodec(amino)
	types.RegisterLegacyAminoCodec(amino)

	// StdSignBytes is the SDK's legacy regression oracle and requires this codec.
	previous := legacytx.RegressionTestingAminoCodec
	legacytx.RegressionTestingAminoCodec = amino
	t.Cleanup(func() { legacytx.RegressionTestingAminoCodec = previous })

	authority := sdk.AccAddress(bytes.Repeat([]byte{8}, types.DerivedAccountLength)).String()
	guardian := sdk.AccAddress(bytes.Repeat([]byte{9}, types.DerivedAccountLength)).String()
	recipient := base58.Encode(bytes.Repeat([]byte{7}, len(types.SolanaAddress{})))
	signature := make([]byte, 64)
	for i := range signature {
		signature[i] = byte(i * 4)
	}

	evidence := types.SolanaEvidence{Slot: math.MaxUint64, TransactionSignature: signature}
	maximum := sdk.NewCoin(sdkutil.DenomUakt, sdkmath.NewIntFromUint64(math.MaxUint64))
	params := types.DefaultParams()
	params.Enabled, params.RefundsEnabled = true, true
	params.Reporters = []string{account(1), account(2), account(3), account(4)}
	params.AllowedDestinationVaults = []string{recipient, base58.Encode(bytes.Repeat([]byte{6}, len(types.SolanaAddress{})))}
	params.PauseGuardian = guardian
	params.MinAmountPerTransfer = sdk.NewInt64Coin(sdkutil.DenomUakt, 1)
	params.MaxAmountPerTransfer = maximum
	params.MaxEpochAmountPerAccount = maximum.AddAmount(sdkmath.OneInt())
	params.MaxTransferAmountPerEpoch = maximum.Add(maximum)
	params.MaxEpochRefund = maximum.AddAmount(sdkmath.NewInt(2))
	params.ReporterChallengePeriod += 123456789 * time.Nanosecond

	type messageCase struct {
		name string
		msg  interface {
			sdk.Msg
			sdk.HasValidateBasic
		}
		signer string
	}

	cases := make([]messageCase, 0, 43)
	cases = append(cases,
		messageCase{"MsgSink/minimum", &types.MsgSink{Sender: account(7), DestType: types.DestTypeSolana, Dest: recipient, Amount: sdk.NewInt64Coin(sdkutil.DenomUakt, 1)}, account(7)},
		messageCase{"MsgSink/maximum", &types.MsgSink{Sender: account(7), DestType: types.DestTypeSolana, Dest: recipient, Amount: maximum}, account(7)},
		messageCase{"MsgUpdateParams/populated", &types.MsgUpdateParams{Authority: authority, Params: params}, authority},
		messageCase{"MsgUpdateParams/defaults", &types.MsgUpdateParams{Authority: authority, Params: types.DefaultParams()}, authority},
		messageCase{"MsgGuardianPause/deposits", &types.MsgGuardianPause{Guardian: guardian, Deposits: true}, guardian},
		messageCase{"MsgGuardianPause/refunds", &types.MsgGuardianPause{Guardian: account(9), Refunds: true}, account(9)},
		messageCase{"MsgGuardianPause/both", &types.MsgGuardianPause{Guardian: guardian, Deposits: true, Refunds: true}, guardian},
	)

	for _, ratio := range []string{"0.666666666666666667", "1"} {
		p := params
		p.ReportersThreshold = sdkmath.LegacyMustNewDecFromStr(ratio)
		cases = append(cases, messageCase{"MsgUpdateParams/ratio_" + ratio, &types.MsgUpdateParams{Authority: authority, Params: p}, authority})
	}

	for name, duration := range map[string]time.Duration{
		"challenge_floor": time.Duration(types.MinChallengePeriodSeconds) * time.Second,
		"max_duration":    types.MaxChallengePeriod,
	} {
		p := params
		p.ReporterChallengePeriod = duration
		cases = append(cases, messageCase{"MsgUpdateParams/" + name, &types.MsgUpdateParams{Authority: authority, Params: p}, authority})
	}

	for name, change := range map[string]func(*types.Params){
		"deposits_disabled": func(p *types.Params) { p.Enabled = false },
		"refunds_disabled":  func(p *types.Params) { p.RefundsEnabled = false },
		"no_guardian":       func(p *types.Params) { p.PauseGuardian = "" },
		"nil_vaults":        func(p *types.Params) { p.AllowedDestinationVaults = nil },
		"empty_vaults":      func(p *types.Params) { p.AllowedDestinationVaults = []string{} },
		"nil_reporters": func(p *types.Params) {
			p.Enabled, p.Reporters = false, nil
		},
		"empty_reporters": func(p *types.Params) {
			p.Enabled, p.Reporters = false, []string{}
		},
		"zero_ratio_without_reporters": func(p *types.Params) {
			p.Enabled, p.Reporters = false, nil
			p.ReportersThreshold = sdkmath.LegacyZeroDec()
		},
	} {
		p := params
		change(&p)
		cases = append(cases, messageCase{"MsgUpdateParams/" + name, &types.MsgUpdateParams{Authority: authority, Params: p}, authority})
	}

	for _, fixture := range []struct {
		name     string
		id       types.ID
		round    uint64
		reason   string
		evidence types.SolanaEvidence
	}{
		{"first", types.ID{Owner: account(7), Height: 1, Sequence: 1}, 1, "verify", types.SolanaEvidence{Slot: 1, TransactionSignature: append([]byte{1}, make([]byte, 63)...)}},
		{"maximum", types.ID{Owner: account(7), Height: math.MaxInt64, Sequence: math.MaxUint64}, math.MaxUint64 - 1, strings.Repeat("é", 128), evidence},
		{"escaped_reason", types.ID{Owner: account(7), Height: 9007199254740993, Sequence: 9007199254740993}, 9007199254740993, "verify <mint> & \"void\" \\ path\n\t雪\u2028\u2029", evidence},
	} {
		id, round, reason, ev := fixture.id, fixture.round, fixture.reason, fixture.evidence
		cases = append(cases,
			messageCase{"MsgRequestVoid/" + fixture.name, &types.MsgRequestVoid{Authority: authority, ID: id}, authority},
			messageCase{"MsgVetoVoid/" + fixture.name, &types.MsgVetoVoid{Signer: account(1), ID: id, Reason: reason}, account(1)},
			messageCase{"MsgRestartTransfer/" + fixture.name, &types.MsgRestartTransfer{Authority: authority, ID: id, Reason: reason, ExpectedRound: round}, authority},
			messageCase{"MsgConfirmMint/" + fixture.name, &types.MsgConfirmMint{Authority: authority, ID: id, Reason: reason, ExpectedRound: round, Evidence: ev}, authority},
			messageCase{"MsgConfirmVoid/" + fixture.name, &types.MsgConfirmVoid{Authority: authority, ID: id, Reason: reason, ExpectedRound: round, Evidence: ev}, authority},
			messageCase{"MsgFinalizeRefund/" + fixture.name, &types.MsgFinalizeRefund{Executor: account(9), ID: id}, account(9)},
		)

		for _, outcome := range []types.Outcome{types.OutcomeMinted, types.OutcomeVoided} {
			cases = append(cases, messageCase{"MsgReportOutcome/" + fixture.name + "/" + outcome.String(), &types.MsgReportOutcome{
				Reporter:      account(1),
				ID:            id,
				PolicyVersion: math.MaxUint64,
				Round:         round,
				Observation:   types.Observation{Outcome: outcome, Evidence: ev},
			}, account(1)})
		}
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.NoError(t, tc.msg.ValidateBasic(), "signing fixtures must be valid under this build's challenge floor")
			packed, err := codectypes.NewAnyWithValue(tc.msg)
			require.NoError(t, err)
			message := &anypb.Any{TypeUrl: packed.TypeUrl, Value: packed.Value}
			t.Run("message_json", func(t *testing.T) {
				legacyJSON, err := amino.MarshalJSON(tc.msg)
				require.NoError(t, err)

				// Legacy JSON preserves struct order; signing sorts its keys.
				require.JSONEq(t, string(legacyJSON), string(signedMessageJSON(t, handler, tc.signer, message)))
			})

			t.Run("sign_bytes", func(t *testing.T) {
				signer := txsigning.SignerData{Address: tc.signer, ChainID: "sink-amino-parity", AccountNumber: math.MaxUint64, Sequence: 9007199254740993}
				body := &txv1beta1.TxBody{Messages: []*anypb.Any{message}, Memo: "Sink <review> & \"sign\"\n雪", TimeoutHeight: math.MaxUint64}
				bodyBytes, err := proto.Marshal(body)
				require.NoError(t, err)
				fee := legacytx.StdFee{Amount: sdk.NewCoins(sdk.NewInt64Coin(sdkutil.DenomUakt, 1234)), Gas: 200000, Payer: account(5), Granter: account(6)}
				legacyBytes := legacytx.StdSignBytes(signer.ChainID, signer.AccountNumber, signer.Sequence, body.TimeoutHeight, fee, []sdk.Msg{tc.msg}, body.Memo) //nolint:staticcheck // Compare the SDK's actual legacy signing path.
				signBytes, err := handler.GetSignBytes(context.Background(), signingv1beta1.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, signer, txsigning.TxData{
					Body:      body,
					BodyBytes: bodyBytes,
					AuthInfo: &txv1beta1.AuthInfo{
						Fee: &txv1beta1.Fee{
							Amount:   []*basev1beta1.Coin{{Denom: sdkutil.DenomUakt, Amount: "1234"}},
							GasLimit: fee.Gas,
							Payer:    fee.Payer,
							Granter:  fee.Granter,
						},
					},
				})
				require.NoError(t, err)
				require.Equal(t, string(legacyBytes), string(signBytes))
			})
		})
	}
}

func signedMessageJSON(t *testing.T, handler *txsigning.HandlerMap, signer string, message *anypb.Any) json.RawMessage {
	t.Helper()
	body := &txv1beta1.TxBody{Messages: []*anypb.Any{message}}
	bodyBytes, err := proto.Marshal(body)
	require.NoError(t, err)
	signBytes, err := handler.GetSignBytes(context.Background(), signingv1beta1.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
		txsigning.SignerData{Address: signer},
		txsigning.TxData{Body: body, BodyBytes: bodyBytes, AuthInfo: &txv1beta1.AuthInfo{Fee: &txv1beta1.Fee{}}})
	require.NoError(t, err)
	var document struct {
		Msgs []json.RawMessage `json:"msgs"`
	}

	require.NoError(t, json.Unmarshal(signBytes, &document))
	require.Len(t, document.Msgs, 1)

	return document.Msgs[0]
}

func TestSinkAminoSignBytesUseParamsEnabled(t *testing.T) {
	encoding := sdkutil.MakeEncodingConfig()
	types.RegisterInterfaces(encoding.InterfaceRegistry)
	params := types.DefaultParams()
	params.Reporters = policy().Members
	params.Enabled = true
	msg := &types.MsgUpdateParams{Authority: account(8), Params: params}
	require.NoError(t, msg.ValidateBasic())
	enabled := legacyAminoSignBytes(t, encoding, msg)
	require.Contains(t, string(enabled), `"enabled":true`)
	require.NotContains(t, string(enabled), `"deposits_enabled"`)
	msg.Params.Enabled = false
	require.NoError(t, msg.ValidateBasic())
	disabled := legacyAminoSignBytes(t, encoding, msg)
	require.NotContains(t, string(disabled), `"enabled"`)
	require.NotEqual(t, string(enabled), string(disabled))
}

func TestSinkAminoSignBytesBindNestedFields(t *testing.T) {
	encoding := sdkutil.MakeEncodingConfig()
	types.RegisterInterfaces(encoding.InterfaceRegistry)
	id := types.ID{Owner: account(7), Height: 1, Sequence: 1}
	request := &types.MsgRequestVoid{Authority: account(8), ID: id}
	report := &types.MsgReportOutcome{
		Reporter:      account(1),
		ID:            id,
		PolicyVersion: 1,
		Round:         1,
		Observation:   types.Observation{Outcome: types.OutcomeMinted, Evidence: evidence()},
	}

	params := &types.MsgUpdateParams{Authority: account(8), Params: types.DefaultParams()}
	params.Params.Enabled, params.Params.RefundsEnabled = true, true
	params.Params.Reporters = policy().Members
	params.Params.PauseGuardian = account(9)
	params.Params.AllowedDestinationVaults = []string{base58.Encode(bytes.Repeat([]byte{7}, len(types.SolanaAddress{})))}
	confirmMint := &types.MsgConfirmMint{Authority: account(8), ID: id, Evidence: evidence(), Reason: "confirm", ExpectedRound: 1}
	confirmVoid := &types.MsgConfirmVoid{Authority: account(8), ID: id, Evidence: evidence(), Reason: "confirm", ExpectedRound: 1}
	sink := sinkMsg()
	for _, tc := range []struct {
		name   string
		msg    sdk.Msg
		change func(sdk.Msg)
	}{
		{"ID owner", request, func(m sdk.Msg) { m.(*types.MsgRequestVoid).ID.Owner = account(9) }},
		{"ID height", request, func(m sdk.Msg) { m.(*types.MsgRequestVoid).ID.Height++ }},
		{"ID sequence", request, func(m sdk.Msg) { m.(*types.MsgRequestVoid).ID.Sequence++ }},
		{"outcome", report, func(m sdk.Msg) { m.(*types.MsgReportOutcome).Observation.Outcome = types.OutcomeVoided }},
		{"evidence slot", report, func(m sdk.Msg) { m.(*types.MsgReportOutcome).Observation.Evidence.Slot++ }},
		{"evidence signature", report, func(m sdk.Msg) { m.(*types.MsgReportOutcome).Observation.Evidence.TransactionSignature[63] ^= 1 }},
		{"report round", report, func(m sdk.Msg) { m.(*types.MsgReportOutcome).Round++ }},
		{"report policy", report, func(m sdk.Msg) { m.(*types.MsgReportOutcome).PolicyVersion++ }},
		{"threshold", params, func(m sdk.Msg) { m.(*types.MsgUpdateParams).Params.ReportersThreshold = sdkmath.LegacyOneDec() }},
		{"challenge", params, func(m sdk.Msg) { m.(*types.MsgUpdateParams).Params.ReporterChallengePeriod += time.Nanosecond }},
		{"coin", params, func(m sdk.Msg) {
			m.(*types.MsgUpdateParams).Params.MaxTransferAmountPerEpoch = sdk.NewInt64Coin(sdkutil.DenomUakt, 20_000_000_000_000)
		}},
		{"minimum coin", params, func(m sdk.Msg) {
			m.(*types.MsgUpdateParams).Params.MinAmountPerTransfer.Amount = params.Params.MinAmountPerTransfer.Amount.AddRaw(1)
		}},
		{"maximum coin", params, func(m sdk.Msg) {
			m.(*types.MsgUpdateParams).Params.MaxAmountPerTransfer.Amount = params.Params.MaxAmountPerTransfer.Amount.SubRaw(1)
		}},
		{"account coin", params, func(m sdk.Msg) {
			m.(*types.MsgUpdateParams).Params.MaxEpochAmountPerAccount.Amount = params.Params.MaxEpochAmountPerAccount.Amount.AddRaw(1)
		}},
		{"refund coin", params, func(m sdk.Msg) {
			m.(*types.MsgUpdateParams).Params.MaxEpochRefund.Amount = params.Params.MaxEpochRefund.Amount.AddRaw(1)
		}},
		{"deposits enabled", params, func(m sdk.Msg) { m.(*types.MsgUpdateParams).Params.Enabled = false }},
		{"refunds enabled", params, func(m sdk.Msg) { m.(*types.MsgUpdateParams).Params.RefundsEnabled = false }},
		{"reporter", params, func(m sdk.Msg) { m.(*types.MsgUpdateParams).Params.Reporters[0] = account(9) }},
		{"guardian", params, func(m sdk.Msg) { m.(*types.MsgUpdateParams).Params.PauseGuardian = account(8) }},
		{"destination vault", params, func(m sdk.Msg) {
			m.(*types.MsgUpdateParams).Params.AllowedDestinationVaults[0] = base58.Encode(bytes.Repeat([]byte{8}, len(types.SolanaAddress{})))
		}},
		{"mint evidence slot", confirmMint, func(m sdk.Msg) { m.(*types.MsgConfirmMint).Evidence.Slot++ }},
		{"mint evidence signature", confirmMint, func(m sdk.Msg) { m.(*types.MsgConfirmMint).Evidence.TransactionSignature[63] ^= 1 }},
		{"void evidence slot", confirmVoid, func(m sdk.Msg) { m.(*types.MsgConfirmVoid).Evidence.Slot++ }},
		{"void evidence signature", confirmVoid, func(m sdk.Msg) { m.(*types.MsgConfirmVoid).Evidence.TransactionSignature[63] ^= 1 }},
		{"destination", sink, func(m sdk.Msg) {
			m.(*types.MsgSink).Dest = base58.Encode(bytes.Repeat([]byte{8}, len(types.SolanaAddress{})))
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			encoded, err := encoding.Codec.Marshal(tc.msg)
			require.NoError(t, err)
			changed, err := encoding.InterfaceRegistry.Resolve(sdk.MsgTypeURL(tc.msg))
			require.NoError(t, err)
			require.NoError(t, encoding.Codec.Unmarshal(encoded, changed))
			tc.change(changed)
			var signBytes [2][]byte
			for i, msg := range []sdk.Msg{tc.msg, changed} {
				require.NoError(t, msg.(sdk.HasValidateBasic).ValidateBasic())
				signBytes[i] = legacyAminoSignBytes(t, encoding, msg)
			}

			require.NotEqual(t, string(signBytes[0]), string(signBytes[1]), "changing a signed field must change the sign bytes")
		})
	}
}

// Only Solana validates today, so the changed message is invalid. A later
// destination type must still not reuse a Solana signature.
func TestSinkAminoSignBytesBindDestinationType(t *testing.T) {
	encoding := sdkutil.MakeEncodingConfig()
	types.RegisterInterfaces(encoding.InterfaceRegistry)
	solana := sinkMsg()
	other := sinkMsg()
	other.DestType = types.DestTypeSolana + 1
	require.NoError(t, solana.ValidateBasic())
	require.ErrorIs(t, other.ValidateBasic(), types.ErrInvalidDestination)
	require.NotEqual(t, string(legacyAminoSignBytes(t, encoding, solana)), string(legacyAminoSignBytes(t, encoding, other)))
}

func sinkMsg() *types.MsgSink {
	return &types.MsgSink{
		DestType: types.DestTypeSolana,
		Dest:     base58.Encode(bytes.Repeat([]byte{7}, len(types.SolanaAddress{}))),
		Amount:   sdk.NewInt64Coin(sdkutil.DenomUakt, 1),
		Sender:   account(7),
	}
}

func legacyAminoSignBytes(t *testing.T, encoding sdkutil.EncodingConfig, msg sdk.Msg) []byte {
	t.Helper()
	builder := encoding.TxConfig.NewTxBuilder()
	require.NoError(t, builder.SetMsgs(msg))
	signers, _, err := encoding.Codec.GetMsgV1Signers(msg)
	require.NoError(t, err)
	require.Len(t, signers, 1)
	signBytes, err := authsigning.GetSignBytesAdapter(context.Background(), encoding.TxConfig.SignModeHandler(),
		signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
		authsigning.SignerData{Address: sdk.AccAddress(signers[0]).String(), ChainID: "sink-amino-parity", AccountNumber: 1, Sequence: 1},
		builder.GetTx())
	require.NoError(t, err)

	return signBytes
}
