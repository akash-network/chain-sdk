package v1

import (
	"math"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func DefaultGenesis() *GenesisState {
	return &GenesisState{Params: DefaultParams(), Totals: Totals{ReporterEpoch: 1}}
}

func (g *GenesisState) Validate() error {
	if err := g.Params.Validate(); err != nil {
		return err
	}

	if g.Totals.ReporterEpoch == 0 || g.Totals.RefundedToday > g.Totals.RefundedAmount {
		return ErrInvalidGenesis.Wrap("zero reporter epoch or refunded today above refunded total")
	}

	seen := make(map[uint64]bool)
	pending := make(map[[AccountLength]byte]uint32)
	var total, pendingAmount, mintedAmount, voidedAmount, refundedAmount, pendingCount uint64

	for _, entry := range g.Transfers {
		t, err := DecodeTransfer(entry.Record)
		if err != nil {
			return err
		}
		if seen[t.ID] || t.ID > g.Totals.LastTransferId {
			return ErrInvalidGenesis.Wrapf("transfer %d duplicated or above last transfer id %d", t.ID, g.Totals.LastTransferId)
		}
		seen[t.ID] = true
		if err := entry.Settlement.Validate(g.Totals.ReporterEpoch); err != nil {
			return err
		}
		policy := entry.Settlement.Policy
		if policy.Epoch == g.Totals.ReporterEpoch && !g.Params.SameReporterSet(Params{Reporters: policy.Members, ReporterThreshold: policy.Threshold}) {
			return ErrInvalidGenesis.Wrapf("transfer %d reporter snapshot does not match parameters", t.ID)
		}
		if t.Amount > math.MaxUint64-total {
			return ErrOverflow.Wrap("transfer amount total")
		}
		total += t.Amount
		if !entry.Settlement.Terminal() {
			if pending[t.Sender] == math.MaxUint32 {
				return ErrOverflow.Wrap("account pending count")
			}
			pending[t.Sender]++
			pendingCount++
		}
		switch entry.Settlement.Status {
		case Status_STATUS_PENDING:
			pendingAmount += t.Amount
		case Status_STATUS_MINTED:
			mintedAmount += t.Amount
		case Status_STATUS_VOIDED:
			voidedAmount += t.Amount
		case Status_STATUS_REFUNDED:
			refundedAmount += t.Amount
		}
	}

	if pendingAmount != g.Totals.PendingAmount || pendingCount != g.Totals.PendingCount || mintedAmount != g.Totals.MintedAmount || voidedAmount != g.Totals.VoidedAmount || refundedAmount != g.Totals.RefundedAmount {
		return ErrInvalidGenesis.Wrap("totals do not match transfer records")
	}

	accounts := make(map[[AccountLength]byte]bool)

	for _, a := range g.Accounts {
		sender := sdk.AccAddress(a.Sender)
		if err := ValidateAccount(sender, AccountLength); err != nil {
			return errorsmod.Wrap(err, "account usage sender")
		}
		key := [AccountLength]byte(sender)
		if accounts[key] || a.PendingCount != pending[key] {
			return ErrInvalidGenesis.Wrapf("account %s duplicated or pending count mismatch", sender)
		}
		accounts[key] = true
	}

	for sender := range pending {
		if !accounts[sender] {
			return ErrInvalidGenesis.Wrapf("missing account usage for %s", sdk.AccAddress(sender[:]))
		}
	}

	return nil
}
