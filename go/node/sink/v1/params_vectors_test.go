package v1_test

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
	"pkg.akt.dev/go/sdkutil"
)

func TestParamsClientVectors(t *testing.T) {
	data, err := os.ReadFile("testdata/params_vectors.json")
	require.NoError(t, err)
	var fixtures []struct {
		Name          string          `json:"name"`
		Threshold     string          `json:"threshold"`
		DurationNanos string          `json:"duration_nanos"`
		JSON          json.RawMessage `json:"json"`
		ProtoHex      string          `json:"proto_hex"`
	}
	require.NoError(t, json.Unmarshal(data, &fixtures))
	require.Len(t, fixtures, 5)
	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())

	for _, fixture := range fixtures {
		t.Run(fixture.Name, func(t *testing.T) {
			duration, err := strconv.ParseInt(fixture.DurationNanos, 10, 64)
			require.NoError(t, err)
			want := types.DefaultParams()
			want.ReportersThreshold = sdkmath.LegacyMustNewDecFromStr(fixture.Threshold)
			want.ReporterChallengePeriod = time.Duration(duration)
			require.NoError(t, want.ValidateBasic())

			nativeJSON, err := cdc.MarshalJSON(&want)
			require.NoError(t, err)
			require.JSONEq(t, string(fixture.JSON), string(nativeJSON))

			binary, err := hex.DecodeString(fixture.ProtoHex)
			require.NoError(t, err)
			nativeBinary, err := want.Marshal()
			require.NoError(t, err)
			require.Equal(t, binary, nativeBinary)

			var fromJSON types.Params
			require.NoError(t, cdc.UnmarshalJSON(fixture.JSON, &fromJSON))
			encodedBinary, err := fromJSON.Marshal()
			require.NoError(t, err)
			require.Equal(t, binary, encodedBinary)

			var fromBinary types.Params
			require.NoError(t, fromBinary.Unmarshal(binary))
			encodedJSON, err := cdc.MarshalJSON(&fromBinary)
			require.NoError(t, err)
			require.JSONEq(t, string(fixture.JSON), string(encodedJSON))
		})
	}
}
