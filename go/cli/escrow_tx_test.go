package cli

import (
	"errors"
	"testing"

	cerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	cflags "pkg.akt.dev/go/cli/flags"
	emodule "pkg.akt.dev/go/node/escrow/module"
)

// abciErr rebuilds an error the way the tx client does when a broadcast fails,
// so the test exercises the real codespace+code reconstruction rather than a
// hand-written string. See node/client/v1beta3/tx.go.
func abciErr(e *cerrors.Error) error {
	return cerrors.ABCIError(e.Codespace(), e.ABCICode(), "failed to execute message; message index: 0: "+e.Error())
}

func TestAnnotateEscrowDepositError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantHint bool
	}{
		{
			name:     "nil error is passed through",
			err:      nil,
			wantHint: false,
		},
		{
			name:     "escrow account-not-found gets the hint",
			err:      abciErr(emodule.ErrAccountNotFound),
			wantHint: true,
		},
		{
			name:     "a different escrow error is left untouched",
			err:      abciErr(emodule.ErrAccountOverdrawn),
			wantHint: false,
		},
		{
			name:     "unrelated error is left untouched",
			err:      errors.New("insufficient funds"),
			wantHint: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := annotateEscrowDepositError(tt.err)

			if tt.err == nil {
				require.NoError(t, got)
				return
			}

			require.Error(t, got)
			// the original error is always preserved in the chain
			require.ErrorIs(t, got, tt.err)

			if tt.wantHint {
				require.Contains(t, got.Error(), "--"+cflags.FlagOwner)
			} else {
				require.Equal(t, tt.err.Error(), got.Error())
			}
		})
	}
}
