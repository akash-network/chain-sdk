package v1_test

import (
	"encoding/hex"
	"encoding/json"
	"math"
	"os"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
	"pkg.akt.dev/go/sdkutil"
)

func TestIDText(t *testing.T) {
	id := types.ID{
		Owner:    account(17),
		Height:   9001,
		Sequence: 42,
	}
	parsed, err := types.ParseID(id.String())
	require.NoError(t, err)
	require.Equal(t, id, parsed)
	for _, text := range []string{"42", id.Owner + "/01/1", id.Owner + "/+1/1", id.Owner + "/1/01", id.Owner + "/1/+1", id.Owner + "/0/1", id.Owner + "/1/0", id.Owner + "/-1/1", id.Owner + "/9223372036854775808/1", id.Owner + "/1/18446744073709551616", strings.ToUpper(id.Owner) + "/1/1", " " + id.String(), id.String() + "/1"} {
		_, err := types.ParseID(text)
		require.Error(t, err, text)
	}

	maximum := types.ID{
		Owner:    account(255),
		Height:   math.MaxInt64,
		Sequence: math.MaxUint64,
	}
	parsed, err = types.ParseID(maximum.String())
	require.NoError(t, err)
	require.Equal(t, maximum, parsed)
}

func TestNativeRecordCodec(t *testing.T) {
	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())
	value := codec.CollValue[types.RecordData](cdc)
	data := types.RecordData{
		Dest:   (types.SolanaAddress{2}).String(),
		Amount: sdk.NewInt64Coin(sdkutil.DenomUakt, 123456789),
	}
	encoded, err := value.Encode(data)
	require.NoError(t, err)
	generated, err := data.Marshal()
	require.NoError(t, err)
	require.Equal(t, generated, encoded)
	decoded, err := value.Decode(encoded)
	require.NoError(t, err)
	require.NoError(t, decoded.ValidateBasic())
	require.Equal(t, data, decoded)

	for n := 0; n < len(encoded); n++ {
		decoded, err := value.Decode(encoded[:n])
		if err == nil {
			err = decoded.ValidateBasic()
		}

		require.Error(t, err, "truncation at %d", n)
	}
}

func TestProofFixtureContracts(t *testing.T) {
	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())
	var records []struct {
		IDText string `json:"id_text"`
		Key    string `json:"transfer_key_hex"`
		Value  string `json:"record_value_hex"`
		Nonce  string `json:"mint_proof_nonce_hex"`
	}
	data, err := os.ReadFile("testdata/proof_vectors.json")
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &records))
	require.Len(t, records, 5)
	for _, vector := range records {
		_, err = types.ParseID(vector.IDText)
		require.NoError(t, err)
		key, err := hex.DecodeString(vector.Key)
		require.NoError(t, err)
		require.Equal(t, types.TransferKeyPrefix, key[:len(types.TransferKeyPrefix)])
		nonce := types.ProofNonce(key)
		require.Equal(t, vector.Nonce, hex.EncodeToString(nonce[:]))
		value, err := hex.DecodeString(vector.Value)
		require.NoError(t, err)
		record, err := codec.CollValue[types.RecordData](cdc).Decode(value)
		require.NoError(t, err)
		require.NoError(t, record.ValidateBasic())
		encoded, err := record.Marshal()
		require.NoError(t, err)
		require.Equal(t, value, encoded)
	}

	var settlements []struct {
		Owner    string                          `json:"owner"`
		Height   int64                           `json:"height,string"`
		Sequence uint64                          `json:"sequence,string"`
		Status   types.Status                    `json:"status"`
		Policies map[uint64]types.ReporterPolicy `json:"policies"`
		Key      string                          `json:"settlement_key_hex"`
		Value    string                          `json:"settlement_value_hex"`
		Nonce    string                          `json:"proof_nonce_hex"`
	}
	data, err = os.ReadFile("testdata/settlement_vectors.json")
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &settlements))
	require.Len(t, settlements, 6)
	for _, vector := range settlements {
		id := types.ID{
			Owner:    vector.Owner,
			Height:   vector.Height,
			Sequence: vector.Sequence,
		}
		require.NoError(t, id.ValidateBasic())
		key, err := hex.DecodeString(vector.Key)
		require.NoError(t, err)
		require.Equal(t, types.SettlementKeyPrefix, key[:len(types.SettlementKeyPrefix)])
		nonce := types.ProofNonce(key)
		require.Equal(t, vector.Nonce, hex.EncodeToString(nonce[:]))
		value, err := hex.DecodeString(vector.Value)
		require.NoError(t, err)
		settlement, err := codec.CollValue[types.SettlementData](cdc).Decode(value)
		require.NoError(t, err)
		require.NoError(t, settlement.Validate(id, vector.Status, vector.Policies))
		encoded, err := settlement.Marshal()
		require.NoError(t, err)
		require.Equal(t, value, encoded)
	}
}
