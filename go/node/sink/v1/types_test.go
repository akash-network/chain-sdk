package v1_test

import (
	"crypto/sha256"
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
	for _, text := range []string{"42", id.Owner + "/1/+1", id.Owner + "/0/1", id.Owner + "/1/0", id.Owner + "/-1/1", id.Owner + "/9223372036854775808/1", id.Owner + "/1/18446744073709551616", " " + id.String(), id.String() + "/1"} {
		_, err := types.ParseID(text)
		require.Error(t, err, text)
	}

	for _, text := range []string{
		id.Owner + "/01/1",
		id.Owner + "/+1/1",
		id.Owner + "/1/01",
		strings.ToUpper(id.Owner) + "/1/1",
	} {
		parsed, err := types.ParseID(text)
		require.NoError(t, err, text)
		owner, err := sdk.AccAddressFromBech32(parsed.Owner)
		require.NoError(t, err)
		require.Equal(t, id.Owner, owner.String())
		require.Equal(t, int64(1), parsed.Height)
		require.Equal(t, uint64(1), parsed.Sequence)
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
		nonce := sha256.Sum256(key)
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
		nonce := sha256.Sum256(key)
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

//func TestDestination(t *testing.T) {
//	pub := ed25519.NewKeyFromSeed(bytes.Repeat([]byte{7}, 32)).Public().(ed25519.PublicKey)
//	addr := base58.Encode(pub)
//	got, err := types.ParseDestination(addr, nil)
//	require.NoError(t, err)
//	require.Equal(t, pub, ed25519.PublicKey(got[:]))
//	require.Equal(t, addr, got.String())
//
//	for name, address := range map[string]string{
//		"empty":  "",
//		"zero":   "11111111111111111111111111111111",
//		"base58": "0OIl",
//		"short":  "123",
//		"long":   "1" + addr,
//		"space":  " " + addr,
//		"memo":   fmt.Sprintf(`{"v":1,"solana_address":%q}`, addr),
//	} {
//		t.Run(name, func(t *testing.T) {
//			_, err := types.ParseDestination(address, nil)
//			require.ErrorIs(t, err, types.ErrInvalidDestination)
//		})
//	}
//
//	offCurve := types.SolanaAddress{2}
//	require.ErrorIs(t, types.ValidateRecipient(offCurve, nil), types.ErrInvalidDestination)
//	got, err = types.ParseDestination(base58.Encode(offCurve[:]), []types.SolanaAddress{offCurve})
//	require.NoError(t, err)
//	require.Equal(t, offCurve, got)
//	require.ErrorIs(t, types.ValidateRecipient(types.SolanaAddress{}, []types.SolanaAddress{{}}), types.ErrInvalidDestination)
//
//	// Small-order points decode but no wallet holds their key: the identity
//	// (y = 1) and the point of order 2 (y = -1).
//	identity := types.SolanaAddress{1}
//	require.ErrorContains(t, types.ValidateRecipient(identity, nil), "small-order")
//	var order2 types.SolanaAddress
//	copy(order2[:], []byte{0xec})
//
//	for i := 1; i < 31; i++ {
//		order2[i] = 0xff
//	}
//
//	order2[31] = 0x7f
//	require.ErrorContains(t, types.ValidateRecipient(order2, nil), "small-order")
//
//	// A listed vault is exempt: vaults may be any 32-byte address.
//	require.NoError(t, types.ValidateRecipient(identity, []types.SolanaAddress{identity}))
//
//	// y + p encodes the same point as y for y < 19. Find such a point of large
//	// order and check its non-canonical spelling is rejected.
//	found := false
//
//	for y := byte(2); y < 19 && !found; y++ {
//		var canonical, alias types.SolanaAddress
//		canonical[0] = y
//		alias[0] = 0xed + y // p = 2^255 - 19: low byte 0xed, then 0xff..., top 0x7f
//		for i := 1; i < 31; i++ {
//			alias[i] = 0xff
//		}
//
//		alias[31] = 0x7f
//
//		if types.ValidateRecipient(canonical, nil) != nil {
//			continue
//		}
//
//		found = true
//		require.ErrorContains(t, types.ValidateRecipient(alias, nil), "canonical")
//	}
//
//	require.True(t, found, "no large-order point with a small y")
//}
