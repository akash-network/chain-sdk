package sdkutil_test

import (
	"context"
	"testing"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"

	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txsigning "cosmossdk.io/x/tx/signing"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	gogoproto "github.com/cosmos/gogoproto/proto"

	dv1 "pkg.akt.dev/go/node/deployment/v1"
	dtypes "pkg.akt.dev/go/node/deployment/v1beta4"
	ptypes "pkg.akt.dev/go/node/provider/v1beta4"
	"pkg.akt.dev/go/sdkutil"
	"pkg.akt.dev/go/testutil"
)

// Simulation operations and third-party modules build their own tx configs.
// BaseApp then resolves the delivered tx's messages with the app's codec.
func TestTxSignersResolveAcrossEncodingConfigs(t *testing.T) {
	provider, sender, deployer := testutil.AccAddress(t), testutil.AccAddress(t), testutil.AccAddress(t)
	annotated := []sdk.Msg{
		providerMsg(provider, "ops@example.com"),
		&wasmtypes.MsgInstantiateContract{Sender: sender.String(), CodeID: 1, Label: "label", Msg: []byte(`{}`)},
	}
	global := globalDescriptorEncoding(t)
	akash := akashEncoding()

	for _, tc := range []struct {
		name               string
		producer, consumer sdkutil.EncodingConfig
		msgs               []sdk.Msg
		signers            [][]byte
	}{
		{"between Akash configs", akashEncoding(), akashEncoding(), append(annotated, deploymentMsg(deployer, "web")), [][]byte{provider, sender, deployer}},
		{"from global descriptors", global, akash, annotated, [][]byte{provider, sender}},
		{"to global descriptors", akash, global, annotated, [][]byte{provider, sender}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			builder := tc.producer.TxConfig.NewTxBuilder()
			require.NoError(t, builder.SetMsgs(tc.msgs...))
			produced := builder.GetTx()
			bz, err := tc.producer.TxConfig.TxEncoder()(produced)
			require.NoError(t, err)
			decoded, err := tc.consumer.TxConfig.TxDecoder()(bz)
			require.NoError(t, err)
			signers, err := decoded.(authsigning.SigVerifiableTx).GetSigners()
			require.NoError(t, err)
			require.Equal(t, tc.signers, signers)

			msgs, err := produced.GetMsgsV2()
			require.NoError(t, err)
			for i, msg := range msgs {
				require.NotPanics(t, func() { signers, err = tc.consumer.Codec.GetMsgV2Signers(msg) })
				require.NoError(t, err)
				require.Equal(t, tc.signers[i:i+1], signers)
			}
		})
	}
}

func TestSignaturesVerifyAcrossEncodingConfigs(t *testing.T) {
	producer, consumer := akashEncoding(), akashEncoding()
	handler := consumer.TxConfig.SignModeHandler()
	require.Equal(t, signingv1beta1.SignMode_SIGN_MODE_DIRECT, handler.DefaultMode())
	require.Equal(t, []signingv1beta1.SignMode{
		signingv1beta1.SignMode_SIGN_MODE_DIRECT,
		signingv1beta1.SignMode_SIGN_MODE_DIRECT_AUX,
		signingv1beta1.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
	}, handler.SupportedModes())

	key := testutil.Key(t)
	deployer := sdk.AccAddress(key.PubKey().Address())
	// The deployer is not the fee payer, so it may sign with every mode.
	msgs := []sdk.Msg{providerMsg(testutil.AccAddress(t), "ops@example.com"), deploymentMsg(deployer, "web")}
	signer := authsigning.SignerData{Address: deployer.String(), ChainID: "akash-signing", AccountNumber: 7, Sequence: 3, PubKey: key.PubKey()}
	packedKey, err := codectypes.NewAnyWithValue(key.PubKey())
	require.NoError(t, err)

	for _, mode := range handler.SupportedModes() {
		t.Run(mode.String(), func(t *testing.T) {
			builder := producer.TxConfig.NewTxBuilder()
			require.NoError(t, builder.SetMsgs(msgs...))
			builder.SetGasLimit(200_000)
			builder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin(sdkutil.DenomUakt, 5_000)))
			signMode := signingtypes.SignMode(mode)
			require.NoError(t, builder.SetSignatures(signingtypes.SignatureV2{
				PubKey: key.PubKey(), Data: &signingtypes.SingleSignatureData{SignMode: signMode}, Sequence: signer.Sequence,
			}))
			signature, err := clienttx.SignWithPrivKey(context.Background(), signMode, signer, builder, key, producer.TxConfig, signer.Sequence)
			require.NoError(t, err)
			require.NoError(t, builder.SetSignatures(signature))

			bz, err := producer.TxConfig.TxEncoder()(builder.GetTx())
			require.NoError(t, err)
			decoded, err := consumer.TxConfig.TxDecoder()(bz)
			require.NoError(t, err)
			signatures, err := decoded.(authsigning.Tx).GetSignaturesV2()
			require.NoError(t, err)
			require.Len(t, signatures, 1)
			require.NoError(t, authsigning.VerifySignature(context.Background(), key.PubKey(), txsigning.SignerData{
				Address: signer.Address, ChainID: signer.ChainID, AccountNumber: signer.AccountNumber, Sequence: signer.Sequence,
				PubKey: &anypb.Any{TypeUrl: packedKey.TypeUrl, Value: packedKey.Value},
			}, signatures[0].Data, handler, decoded.(authsigning.V2AdaptableTx).GetSigningTxData()))
		})
	}
}

// Akash files register before their same-package imports, so the global
// registry leaves these nested messages unresolved.
func TestLegacyAminoSignBytesBindNestedFields(t *testing.T) {
	encoding := akashEncoding()
	owner := testutil.AccAddress(t)
	for _, tc := range []struct {
		name             string
		signed, modified sdk.Msg
	}{
		{"top-level field", providerMsg(owner, "ops@example.com"), &ptypes.MsgCreateProvider{
			Owner: owner.String(), HostURI: "https://other.example:8443", Info: ptypes.Info{EMail: "ops@example.com"},
		}},
		{"nested message", providerMsg(owner, "ops@example.com"), providerMsg(owner, "abuse@example.com")},
		{"repeated nested message", deploymentMsg(owner, "web"), deploymentMsg(owner, "db")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.NotEqual(t, string(legacyAminoSignBytes(t, encoding, owner, tc.signed)), string(legacyAminoSignBytes(t, encoding, owner, tc.modified)),
				"changing a signed field must change the sign bytes")
		})
	}
}

func legacyAminoSignBytes(t *testing.T, encoding sdkutil.EncodingConfig, signer sdk.AccAddress, msg sdk.Msg) []byte {
	t.Helper()
	builder := encoding.TxConfig.NewTxBuilder()
	require.NoError(t, builder.SetMsgs(msg))
	signBytes, err := authsigning.GetSignBytesAdapter(context.Background(), encoding.TxConfig.SignModeHandler(),
		signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
		authsigning.SignerData{Address: signer.String(), ChainID: "akash-signing", AccountNumber: 1, Sequence: 1},
		builder.GetTx())
	require.NoError(t, err)
	return signBytes
}

func providerMsg(owner sdk.AccAddress, email string) *ptypes.MsgCreateProvider {
	return &ptypes.MsgCreateProvider{Owner: owner.String(), HostURI: "https://provider.example:8443", Info: ptypes.Info{EMail: email}}
}

func deploymentMsg(owner sdk.AccAddress, group string) *dtypes.MsgCreateDeployment {
	return &dtypes.MsgCreateDeployment{ID: dv1.DeploymentID{Owner: owner.String(), DSeq: 1}, Groups: dtypes.GroupSpecs{{Name: group}}}
}

func akashEncoding() sdkutil.EncodingConfig {
	encoding := sdkutil.MakeEncodingConfig()
	registerMsgs(encoding.InterfaceRegistry)
	return encoding
}

// globalDescriptorEncoding builds a tx config the way wasmd's simulation
// operations do, on the process-wide descriptors.
func globalDescriptorEncoding(t *testing.T) sdkutil.EncodingConfig {
	t.Helper()
	registry, err := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: gogoproto.HybridResolver,
		SigningOptions: txsigning.Options{
			AddressCodec:          address.NewBech32Codec(sdkutil.Bech32PrefixAccAddr),
			ValidatorAddressCodec: address.NewBech32Codec(sdkutil.Bech32PrefixValAddr),
		},
	})
	require.NoError(t, err)
	registerMsgs(registry)
	cdc := codec.NewProtoCodec(registry)
	return sdkutil.EncodingConfig{InterfaceRegistry: registry, Codec: cdc, TxConfig: authtx.NewTxConfig(cdc, authtx.DefaultSignModes)}
}

func registerMsgs(registry codectypes.InterfaceRegistry) {
	ptypes.RegisterInterfaces(registry)
	dtypes.RegisterInterfaces(registry)
	wasmtypes.RegisterInterfaces(registry)
}
