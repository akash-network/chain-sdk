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
	"github.com/cosmos/cosmos-sdk/codec"
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

func TestNativeRecordCodec(t *testing.T) {
	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())
	value := codec.CollValue[types.RecordData](cdc)
	data := types.RecordData{Dest: (types.SolanaAddress{2}).String(), Amount: sdk.NewInt64Coin(sdkutil.DenomUakt, 123456789)}
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
	MintNonce    string `json:"mint_proof_nonce_hex"`
	HeightLE     string `json:"height_le_hex"`
	SequenceLE   string `json:"sequence_le_hex"`
}

func TestProofVectors(t *testing.T) {
	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())
	vectors := make([]proofVector, 0, 5)
	for _, v := range []struct {
		name             string
		owner            byte
		height           int64
		sequence, amount uint64
	}{
		{"ordinary", 17, 9001, 42, 123456789}, {"minimum", 1, 1, 1, 1},
		{"amount127", 2, 127, 127, 127}, {"amount128", 3, 128, 128, 128},
		{"maximum", 255, math.MaxInt64, math.MaxUint64, math.MaxUint64},
	} {
		id := types.ID{Owner: account(v.owner), Height: v.height, Sequence: v.sequence}
		recipient := types.SolanaAddress{}
		copy(recipient[:], bytes.Repeat([]byte{0x22}, 32))

		if v.name == "minimum" {
			clear(recipient[:])
			recipient[31] = 1
		}

		data := types.RecordData{Dest: recipient.String(), Amount: sdk.NewCoin(sdkutil.DenomUakt, sdkmath.NewIntFromUint64(v.amount))}
		require.NoError(t, data.ValidateBasic())
		key, err := types.TransferKey(id)
		require.NoError(t, err)
		require.Len(t, key, 39)
		value, err := codec.CollValue[types.RecordData](cdc).Encode(data)
		require.NoError(t, err)
		mintNonce := types.ProofNonce(key)
		require.Equal(t, sha256.Sum256(key), mintNonce)
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
			MintNonce:    hex.EncodeToString(mintNonce[:]),
			HeightLE:     hex.EncodeToString(heightLE),
			SequenceLE:   hex.EncodeToString(binary.LittleEndian.AppendUint64(nil, v.sequence)),
		})
	}

	checkVectors(t, "testdata/proof_vectors.json", vectors)
}

type settlementVector struct {
	Name       string          `json:"name"`
	Owner      string          `json:"owner"`
	OwnerHex   string          `json:"owner_hex"`
	Height     string          `json:"height"`
	Sequence   string          `json:"sequence"`
	Status     int32           `json:"status"`
	Settlement json.RawMessage `json:"settlement"`
	Policies   json.RawMessage `json:"policies"`
	Key        string          `json:"settlement_key_hex"`
	Value      string          `json:"settlement_value_hex"`
	Nonce      string          `json:"proof_nonce_hex"`
}

func TestSettlementProofVectors(t *testing.T) {
	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())
	id := types.ID{Owner: account(17), Height: 9001, Sequence: 42}
	policies := map[uint64]types.ReporterPolicy{1: policy()}
	var vectors []settlementVector
	add := func(name string, id types.ID, status types.Status, settlement types.SettlementData, policies map[uint64]types.ReporterPolicy) {
		t.Helper()
		require.NoError(t, settlement.Validate(id, status, policies))
		key, err := types.SettlementKey(types.NewSettlementID(status, id))
		require.NoError(t, err)
		require.Len(t, key, 40)
		value, err := codec.CollValue[types.SettlementData](cdc).Encode(settlement)
		require.NoError(t, err)

		if name == "maximum_evidence" {
			require.Len(t, value, 6500)
		}

		js, err := cdc.MarshalJSON(&settlement)
		require.NoError(t, err)
		pj, err := json.Marshal(policies)
		require.NoError(t, err)
		owner, err := types.ParseAccount(id.Owner, types.AccountLength)
		require.NoError(t, err)
		nonce := types.ProofNonce(key)
		vectors = append(vectors, settlementVector{Name: name, Owner: id.Owner, OwnerHex: hex.EncodeToString(owner), Height: strconv.FormatInt(id.Height, 10), Sequence: strconv.FormatUint(id.Sequence, 10), Status: int32(status), Settlement: js, Policies: pj, Key: hex.EncodeToString(key), Value: hex.EncodeToString(value), Nonce: hex.EncodeToString(nonce[:])})
	}

	for _, status := range []types.Status{types.StatusTransferPending, types.StatusMinted, types.StatusVoided, types.StatusRefundPending, types.StatusRefunded} {
		s := settled(status)
		if status != types.StatusTransferPending {
			s.VoidAuthorization = &types.VoidAuthorization{RequestedHeight: 9001, RequestedRound: 1}
		}

		add(status.String(), id, status, s, policies)
	}

	members := make([]string, types.MaxReporters)
	for i := range members {
		members[i] = account(byte(i + 1))
	}

	primary := types.ReporterPolicy{Version: math.MaxUint64 - 1, Members: members, Threshold: types.MaxReporters}
	witness := types.ReporterPolicy{Version: math.MaxUint64, Members: []string{account(33), account(34)}, Threshold: 2}
	s := types.SettlementData{
		Round:             math.MaxUint64,
		Hold:              types.HoldConflict,
		Policy:            primary,
		ChallengeEnd:      253402300799,
		VetoBy:            account(34),
		VetoReason:        strings.Repeat("v", 256),
		VetoPolicyVersion: witness.Version,
		VoidAuthorization: &types.VoidAuthorization{RequestedHeight: math.MaxInt64, RequestedRound: math.MaxUint64},
	}

	for _, member := range members {
		s.Votes = append(s.Votes, types.Vote{Reporter: member, PolicyVersion: primary.Version, Observation: types.Observation{Outcome: types.OutcomeVoided, Slot: math.MaxUint64, TransactionSignature: bytes.Repeat([]byte{0xff}, 64)}})
	}

	s.Votes = append(s.Votes, types.Vote{Reporter: account(33), PolicyVersion: witness.Version, Observation: types.Observation{Outcome: types.OutcomeMinted, Slot: math.MaxUint64, TransactionSignature: bytes.Repeat([]byte{0xff}, 64)}})
	add("maximum_evidence", types.ID{Owner: account(255), Height: math.MaxInt64, Sequence: math.MaxUint64}, types.StatusRefundPending, s, map[uint64]types.ReporterPolicy{primary.Version: primary, witness.Version: witness})
	checkVectors(t, "testdata/settlement_vectors.json", vectors)
}

func checkVectors(t *testing.T, path string, vectors any) {
	t.Helper()
	encoded, err := json.MarshalIndent(vectors, "", "  ")
	require.NoError(t, err)
	encoded = append(encoded, '\n')
	if os.Getenv("UPDATE_PROOF_VECTORS") == "1" {
		require.NoError(t, os.MkdirAll("testdata", 0755))
		require.NoError(t, os.WriteFile(path, encoded, 0600))
	}

	want, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, string(want), string(encoded))
}
