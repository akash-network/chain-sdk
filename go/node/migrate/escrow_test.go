package migrate

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	eid "pkg.akt.dev/go/node/escrow/id/v1"
	"pkg.akt.dev/go/node/escrow/v1beta3"
)

func TestAccountIDFromV1beta3(t *testing.T) {
	// Get the prefix directly from v1beta3 module
	prefix := v1beta3.AccountKeyPrefix()

	// Sample data
	akashAddr := "akash1keydahz9uv8fs8u4lk6q3cluaprjpnm7dd3cf0"
	dseq := "1000"

	// Helper function to create valid account keys
	createAccountKey := func(scope string, xid string) []byte {
		return append(prefix, []byte("/"+scope+"/"+xid)...)
	}

	t.Run("valid deployment account", func(t *testing.T) {
		key := createAccountKey("deployment", akashAddr+"/"+dseq)
		expected := eid.Account{
			Scope: eid.ScopeDeployment,
			XID:   akashAddr + "/" + dseq,
		}

		result := AccountIDFromV1beta3(key)
		require.Equal(t, expected, result)
	})

	t.Run("valid bid account", func(t *testing.T) {
		key := createAccountKey("bid", akashAddr+"/"+dseq)
		expected := eid.Account{
			Scope: eid.ScopeBid,
			XID:   akashAddr + "/" + dseq,
		}

		result := AccountIDFromV1beta3(key)
		require.Equal(t, expected, result)
	})

	t.Run("prefix check", func(t *testing.T) {
		invalidPrefix := []byte("/wrong_prefix/deployment/" + akashAddr + "/" + dseq)
		require.Panics(t, func() {
			AccountIDFromV1beta3(invalidPrefix)
		})
	})

	t.Run("separator check", func(t *testing.T) {
		// Missing separator after prefix
		invalidKey := append(bytes.Clone(prefix), []byte("deployment/"+akashAddr+"/"+dseq)...)
		require.Panics(t, func() {
			AccountIDFromV1beta3(invalidKey)
		})
	})
	t.Run("empty XID after separator", func(t *testing.T) {
		// Key has prefix and separator but no XID parts
		shortKey := append(bytes.Clone(prefix), '/')
		require.Panics(t, func() {
			AccountIDFromV1beta3(shortKey)
		})
	})

	t.Run("invalid scope", func(t *testing.T) {
		// Scope is not "deployment" or "bid"
		invalidScope := createAccountKey("invalid", akashAddr+"/"+dseq)
		require.Panics(t, func() {
			AccountIDFromV1beta3(invalidScope)
		})
	})

	t.Run("invalid parts count", func(t *testing.T) {
		// Not enough parts
		tooFewParts := createAccountKey("deployment", akashAddr)
		require.Panics(t, func() {
			AccountIDFromV1beta3(tooFewParts)
		})

		// Too many parts that don't match 3 or 6
		tooManyParts := createAccountKey("deployment", akashAddr+"/"+dseq+"/extra/random")
		require.Panics(t, func() {
			AccountIDFromV1beta3(tooManyParts)
		})
	})

	t.Run("complex XID with 6 parts", func(t *testing.T) {
		// Valid XID with 6 parts
		complexXID := akashAddr + "/1/2/3/4"
		key := createAccountKey("bid", complexXID)
		expected := eid.Account{
			Scope: eid.ScopeBid,
			XID:   complexXID,
		}

		result := AccountIDFromV1beta3(key)
		require.Equal(t, expected, result)
	})
}
