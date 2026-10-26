package v1

import (
	"fmt"
	"math"
)

func DefaultGenesis() *GenesisState {
	return &GenesisState{Params: DefaultParams(), Totals: Totals{ReporterEpoch: 1}}
}

func (g GenesisState) Validate() error {
	if err := g.Params.Validate(); err != nil {
		return err
	}
	if g.Totals.ReporterEpoch == 0 || g.Totals.RefundedToday > g.Totals.RefundedAmount {
		return fmt.Errorf("invalid reporter epoch or refund accounting")
	}
	seen := make(map[uint64]bool)
	pending := make(map[[20]byte]uint32)
	var total, pendingAmount, mintedAmount, voidedAmount, refundedAmount, pendingCount uint64
	for _, entry := range g.Transfers {
		t, err := DecodeTransfer(entry.Record)
		if err != nil {
			return err
		}
		if seen[t.ID] || t.ID > g.Totals.LastTransferId {
			return fmt.Errorf("duplicate ID or invalid last transfer ID")
		}
		seen[t.ID] = true
		if err := entry.Settlement.Validate(g.Totals.ReporterEpoch); err != nil {
			return err
		}
		policy := entry.Settlement.Policy
		if policy.Epoch == g.Totals.ReporterEpoch && !g.Params.SameReporterSet(Params{Reporters: policy.Members, ReporterThreshold: policy.Threshold}) {
			return fmt.Errorf("current reporter snapshot does not match parameters")
		}
		if t.Amount > math.MaxUint64-total {
			return fmt.Errorf("genesis counter overflow")
		}
		total += t.Amount
		if !entry.Settlement.Terminal() {
			if pending[t.Sender] == math.MaxUint32 {
				return fmt.Errorf("genesis pending count overflow")
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
		return fmt.Errorf("genesis settlement totals do not match records")
	}
	accounts := make(map[[20]byte]bool)
	for _, a := range g.Accounts {
		if len(a.Sender) != 20 {
			return fmt.Errorf("invalid account address")
		}
		key := [20]byte(a.Sender)
		if accounts[key] || a.PendingCount != pending[key] {
			return fmt.Errorf("invalid account pending count or duplicate account")
		}
		accounts[key] = true
	}
	for sender := range pending {
		if !accounts[sender] {
			return fmt.Errorf("missing account usage")
		}
	}
	return nil
}
