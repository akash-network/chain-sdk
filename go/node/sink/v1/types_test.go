package v1_test

import (
	"bytes"
	"crypto/ed25519"
	"fmt"
	"testing"

	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
)

func TestDestination(t *testing.T) {
	pub := ed25519.NewKeyFromSeed(bytes.Repeat([]byte{7}, 32)).Public().(ed25519.PublicKey)
	addr := base58.Encode(pub)
	got, err := types.ParseDestination(addr, nil)
	require.NoError(t, err)
	require.Equal(t, pub, ed25519.PublicKey(got[:]))

	for name, address := range map[string]string{
		"empty":  "",
		"zero":   "11111111111111111111111111111111",
		"base58": "0OIl",
		"short":  "123",
		"long":   "1" + addr,
		"space":  " " + addr,
		"memo":   fmt.Sprintf(`{"v":1,"solana_address":%q}`, addr),
	} {
		t.Run(name, func(t *testing.T) {
			_, err := types.ParseDestination(address, nil)
			require.ErrorIs(t, err, types.ErrInvalidDestination)
		})
	}

	offCurve := types.SolanaAddress{2}
	require.ErrorIs(t, types.ValidateRecipient(offCurve, nil), types.ErrInvalidDestination)
	got, err = types.ParseDestination(base58.Encode(offCurve[:]), []types.SolanaAddress{offCurve})
	require.NoError(t, err)
	require.Equal(t, offCurve, got)
	require.ErrorIs(t, types.ValidateRecipient(types.SolanaAddress{}, []types.SolanaAddress{{}}), types.ErrInvalidDestination)

	// Small-order points decode but no wallet holds their key: the identity
	// (y = 1) and the point of order 2 (y = -1).
	identity := types.SolanaAddress{1}
	require.ErrorContains(t, types.ValidateRecipient(identity, nil), "small-order")
	var order2 types.SolanaAddress
	copy(order2[:], []byte{0xec})

	for i := 1; i < 31; i++ {
		order2[i] = 0xff
	}

	order2[31] = 0x7f
	require.ErrorContains(t, types.ValidateRecipient(order2, nil), "small-order")

	// A listed vault is exempt: vaults may be any 32-byte address.
	require.NoError(t, types.ValidateRecipient(identity, []types.SolanaAddress{identity}))

	// y + p encodes the same point as y for y < 19. Find such a point of large
	// order and check its non-canonical spelling is rejected.
	found := false
	for y := byte(2); y < 19 && !found; y++ {
		var canonical, alias types.SolanaAddress
		canonical[0] = y
		alias[0] = 0xed + y // p = 2^255 - 19: low byte 0xed, then 0xff..., top 0x7f
		for i := 1; i < 31; i++ {
			alias[i] = 0xff
		}

		alias[31] = 0x7f

		if types.ValidateRecipient(canonical, nil) != nil {
			continue
		}

		found = true
		require.ErrorContains(t, types.ValidateRecipient(alias, nil), "canonical")
	}

	require.True(t, found, "no large-order point with a small y")
}
