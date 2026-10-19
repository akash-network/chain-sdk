package v1_test

import (
	"bytes"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
	"pkg.akt.dev/go/sdkutil"
)

func TestSinkSDKCodec(t *testing.T) {
	registry := sdkutil.NewCodecOptions().NewInterfaceRegistry()
	types.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)
	amino := codec.NewLegacyAmino()
	std.RegisterLegacyAminoCodec(amino)
	types.RegisterLegacyAminoCodec(amino)
	msg := &types.MsgSink{DestType: 1, Dest: "destination", Amount: sdk.NewInt64Coin(types.Denom, 17), Sender: sdk.AccAddress(bytes.Repeat([]byte{7}, 20)).String()}
	packed, err := codectypes.NewAnyWithValue(msg)
	require.NoError(t, err)
	require.Equal(t, "/akash.sink.v1.MsgSink", packed.TypeUrl)
	var decoded sdk.Msg
	require.NoError(t, cdc.UnpackAny(packed, &decoded))
	require.Equal(t, msg, decoded)
	encoded, err := amino.MarshalJSON(msg)
	require.NoError(t, err)
	require.Contains(t, string(encoded), `"type":"sink/MsgSink"`)
	var decodedAmino sdk.Msg
	require.NoError(t, amino.UnmarshalJSON(encoded, &decodedAmino))
	require.Equal(t, msg, decodedAmino)

	response := &types.MsgSinkResponse{Id: "18446744073709551615"}
	packedResponse, err := codectypes.NewAnyWithValue(response)
	require.NoError(t, err)
	require.Equal(t, "/akash.sink.v1.MsgSinkResponse", packedResponse.TypeUrl)
	resolved, err := registry.Resolve(packedResponse.TypeUrl)
	require.NoError(t, err)
	require.NoError(t, proto.Unmarshal(packedResponse.Value, resolved))
	require.True(t, proto.Equal(response, resolved))
}

func TestSinkGuardianPauseCodec(t *testing.T) {
	registry := sdkutil.NewCodecOptions().NewInterfaceRegistry()
	types.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)
	amino := codec.NewLegacyAmino()
	std.RegisterLegacyAminoCodec(amino)
	types.RegisterLegacyAminoCodec(amino)
	msg := &types.MsgGuardianPause{Guardian: sdk.AccAddress(bytes.Repeat([]byte{9}, 20)).String(), Deposits: true}
	packed, err := codectypes.NewAnyWithValue(msg)
	require.NoError(t, err)
	require.Equal(t, "/akash.sink.v1.MsgGuardianPause", packed.TypeUrl)
	var decoded sdk.Msg
	require.NoError(t, cdc.UnpackAny(packed, &decoded))
	require.Equal(t, msg, decoded)
	encoded, err := amino.MarshalJSON(msg)
	require.NoError(t, err)
	require.Contains(t, string(encoded), `"type":"sink/MsgGuardianPause"`)
	var decodedAmino sdk.Msg
	require.NoError(t, amino.UnmarshalJSON(encoded, &decodedAmino))
	require.Equal(t, msg, decodedAmino)
	resolved, err := registry.Resolve("/akash.sink.v1.MsgGuardianPauseResponse")
	require.NoError(t, err)
	require.IsType(t, &types.MsgGuardianPauseResponse{}, resolved)
}
