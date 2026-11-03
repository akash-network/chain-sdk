package v1_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
	"pkg.akt.dev/go/sdkutil"
)

func TestStrictIDCodec(t *testing.T) {
	id := types.ID{Owner: account(17), Height: 9001, Sequence: 42}
	parsed, err := types.ParseID(id.String())
	require.NoError(t, err)
	require.Equal(t, id, parsed)
	b := make([]byte, types.IDKeySize)
	n, err := types.IDKey.Encode(b, id)
	require.NoError(t, err)
	require.Equal(t, types.IDKeySize, n)
	require.Equal(t, "14"+strings.Repeat("11", 20)+"0000000000002329000000000000002a", hex.EncodeToString(b))
	n, parsed, err = types.IDKey.Decode(b)
	require.NoError(t, err)
	require.Equal(t, id, parsed)
	require.Equal(t, types.IDKeySize, n)

	for length := 0; length < types.IDKeySize; length++ {
		_, _, err := types.IDKey.Decode(b[:length])
		require.Error(t, err)
		_, _, err = types.IDKey.DecodeNonTerminal(b[:length])
		require.Error(t, err)
		_, err = types.IDKey.Encode(b[:length], id)
		require.Error(t, err)
		_, err = types.IDKey.EncodeNonTerminal(b[:length], id)
		require.Error(t, err)
	}

	for _, change := range []func([]byte){
		func(b []byte) { b[0] = 19 }, func(b []byte) { clear(b[1:21]) },
		func(b []byte) { clear(b[21:29]) }, func(b []byte) { b[21] = 128 }, func(b []byte) { clear(b[29:37]) },
	} {
		bad := bytes.Clone(b)
		change(bad)
		_, _, err := types.IDKey.Decode(bad)
		require.Error(t, err)
	}

	for _, text := range []string{"42", id.Owner + "/01/1", id.Owner + "/+1/1", id.Owner + "/1/01", id.Owner + "/1/+1", id.Owner + "/0/1", id.Owner + "/1/0", id.Owner + "/-1/1", id.Owner + "/9223372036854775808/1", id.Owner + "/1/18446744073709551616", strings.ToUpper(id.Owner) + "/1/1", " " + id.String(), id.String() + "/1"} {
		_, err := types.ParseID(text)
		require.Error(t, err, text)
	}

	maximum := types.ID{Owner: account(1), Height: math.MaxInt64, Sequence: math.MaxUint64}
	_, err = types.IDKey.Encode(b, maximum)
	require.NoError(t, err)
	_, parsed, err = types.IDKey.DecodeNonTerminal(append(b, 1))
	require.NoError(t, err)
	require.Equal(t, maximum, parsed)
	require.Equal(t, types.IDKeySize, types.IDKey.Size(types.ID{}))
	require.Equal(t, types.IDKeySize, types.IDKey.SizeNonTerminal(types.ID{}))
}

func TestRecordAndVoidCodecsRejectEveryTruncation(t *testing.T) {
	recipient := types.SolanaAddress{2}
	data := types.RecordData{Dest: recipient.String(), Amount: sdk.NewInt64Coin(sdkutil.DenomUakt, 123456789)}
	encoded, err := types.RecordDataValue.Encode(data)
	require.NoError(t, err)
	require.Len(t, encoded, 41)
	require.Equal(t, "02"+hex.EncodeToString(recipient[:])+"15cd5b0700000000", hex.EncodeToString(encoded))

	for n := 0; n < len(encoded); n++ {
		_, err := types.RecordDataValue.Decode(encoded[:n])
		require.Error(t, err)
	}

	for _, change := range []func([]byte){func(b []byte) { b[0] = 1 }, func(b []byte) { clear(b[1:33]) }, func(b []byte) { clear(b[33:]) }} {
		bad := bytes.Clone(encoded)
		change(bad)
		_, err := types.RecordDataValue.Decode(bad)
		require.Error(t, err)
	}

	_, err = types.RecordDataValue.Decode(append(encoded, 0))
	require.Error(t, err)

	for _, bad := range [][]byte{nil, {1}, {2, 0}} {
		_, err := types.VoidOrderValue.Decode(bad)
		require.Error(t, err)
	}

	_, err = types.VoidOrderValue.Decode([]byte{2})
	require.NoError(t, err)
}

type proofVector struct {
	Name         string `json:"name"`
	Owner        string `json:"owner"`
	OwnerHex     string `json:"owner_hex"`
	Height       string `json:"height"`
	Sequence     string `json:"sequence"`
	Recipient    string `json:"recipient"`
	RecipientHex string `json:"recipient_hex"`
	Amount       string `json:"amount"`
	IDText       string `json:"id_text"`
	TransferKey  string `json:"transfer_key_hex"`
	RecordValue  string `json:"record_value_hex"`
	VoidKey      string `json:"void_order_key_hex"`
	VoidValue    string `json:"void_order_value_hex"`
	MintNonce    string `json:"mint_proof_nonce_hex"`
	VoidNonce    string `json:"void_proof_nonce_hex"`
	HeightLE     string `json:"height_le_hex"`
	SequenceLE   string `json:"sequence_le_hex"`
}

func TestProofVectors(t *testing.T) {
	vectors := make([]proofVector, 0, 3)
	for _, v := range []struct {
		name             string
		owner            byte
		height           int64
		sequence, amount uint64
	}{
		{"ordinary", 17, 9001, 42, 123456789}, {"minimum", 1, 1, 1, 1}, {"maximum", 255, math.MaxInt64, math.MaxUint64, math.MaxUint64},
	} {
		id := types.ID{Owner: account(v.owner), Height: v.height, Sequence: v.sequence}
		recipient := types.SolanaAddress{}
		copy(recipient[:], bytes.Repeat([]byte{0x22}, 32))
		data := types.RecordData{Dest: recipient.String(), Amount: sdk.NewCoin(sdkutil.DenomUakt, sdkmath.NewIntFromUint64(v.amount))}
		key, err := types.TransferKey(id)
		require.NoError(t, err)
		value, err := types.RecordDataValue.Encode(data)
		require.NoError(t, err)
		order, err := types.VoidOrderKey(id)
		require.NoError(t, err)
		void, err := types.VoidOrderValue.Encode(types.VoidOrder{})
		require.NoError(t, err)
		mintNonce, voidNonce := sha256.Sum256(key), sha256.Sum256(order)
		require.NotEqual(t, mintNonce, voidNonce)
		heightLE := binary.LittleEndian.AppendUint64(nil, uint64(v.height)) //nolint:gosec // Fixture heights are positive int64 values.
		vectors = append(vectors, proofVector{
			Name:         v.name,
			Owner:        id.Owner,
			OwnerHex:     hex.EncodeToString(bytes.Repeat([]byte{v.owner}, 20)),
			Height:       strconv.FormatInt(v.height, 10),
			Sequence:     strconv.FormatUint(v.sequence, 10),
			Recipient:    data.Dest,
			RecipientHex: hex.EncodeToString(recipient[:]),
			Amount:       strconv.FormatUint(v.amount, 10),
			IDText:       id.String(),
			TransferKey:  hex.EncodeToString(key),
			RecordValue:  hex.EncodeToString(value),
			VoidKey:      hex.EncodeToString(order),
			VoidValue:    hex.EncodeToString(void),
			MintNonce:    hex.EncodeToString(mintNonce[:]),
			VoidNonce:    hex.EncodeToString(voidNonce[:]),
			HeightLE:     hex.EncodeToString(heightLE),
			SequenceLE:   hex.EncodeToString(binary.LittleEndian.AppendUint64(nil, v.sequence)),
		})
	}

	encoded, err := json.MarshalIndent(vectors, "", "  ")
	require.NoError(t, err)
	encoded = append(encoded, '\n')
	path := "testdata/proof_vectors.json"
	if os.Getenv("UPDATE_PROOF_VECTORS") == "1" {
		require.NoError(t, os.MkdirAll("testdata", 0755))
		require.NoError(t, os.WriteFile(path, encoded, 0600))
	}

	want, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, string(want), string(encoded))
}
