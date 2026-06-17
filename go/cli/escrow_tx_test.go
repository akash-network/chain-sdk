package cli

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	cflags "pkg.akt.dev/go/cli/flags"
)

func TestAnnotateEscrowDepositError(t *testing.T) {
	notFound := errors.New("rpc error: code = Unknown desc = failed to execute message; message index: 0: account not found")

	tests := []struct {
		name     string
		err      error
		ownerSet bool
		wantHint bool
	}{
		{
			name:     "nil error is passed through",
			err:      nil,
			ownerSet: false,
			wantHint: false,
		},
		{
			name:     "account not found without --owner gets the hint",
			err:      notFound,
			ownerSet: false,
			wantHint: true,
		},
		{
			name:     "account not found with --owner is left untouched",
			err:      notFound,
			ownerSet: true,
			wantHint: false,
		},
		{
			name:     "unrelated error without --owner is left untouched",
			err:      errors.New("insufficient funds"),
			ownerSet: false,
			wantHint: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := annotateEscrowDepositError(tt.err, tt.ownerSet)

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
