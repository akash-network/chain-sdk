package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const SchemaVersion uint32 = 2

func DefaultGenesis() *GenesisState {
	params := DefaultParams()
	return &GenesisState{
		SchemaVersion: SchemaVersion,
		Params:        params,
		Totals:        DefaultTotals(),
		Policies:      []ReporterPolicy{params.Policy(1)},
		SendUsage:     EpochUsage{Amount: ZeroAmount()},
		RefundUsage:   EpochUsage{Amount: ZeroAmount()},
	}
}

func (g *GenesisState) Validate() error {
	if g == nil || g.SchemaVersion != SchemaVersion {
		return ErrInvalidGenesis.Wrapf("schema_version must be %d; old sink state cannot be imported", SchemaVersion)
	}

	if err := g.Params.ValidateBasic(); err != nil {
		return err
	}

	history, current, err := policyHistory(g.Policies)
	if err != nil {
		return err
	}

	if !current.SameRule(g.Params.Policy(current.Version)) {
		return ErrInvalidGenesis.Wrap("current reporter policy differs from params")
	}

	if err = g.Totals.Validate(); err != nil {
		return err
	}

	if _, err = g.Totals.DepositedAmount(); err != nil {
		return err
	}

	if err = g.SendUsage.Validate(); err != nil {
		return err
	}

	if err = g.RefundUsage.Validate(); err != nil {
		return err
	}

	totals := DefaultTotals()
	amounts := make(map[string]sdk.Coin)
	ids := make(map[ID]bool, len(g.Transfers))
	type position struct {
		height   int64
		sequence uint64
	}

	positions := make(map[position]bool, len(g.Transfers))
	for _, transfer := range g.Transfers {
		if err = transfer.ID.ValidateBasic(); err != nil {
			return err
		}

		if err = transfer.Data.ValidateBasic(); err != nil {
			return err
		}

		if err = transfer.Settlement.Validate(history); err != nil {
			return err
		}

		id := transfer.ID
		pos := position{id.Height, id.Sequence}
		if ids[id] || positions[pos] {
			return ErrInvalidGenesis.Wrap("duplicate transfer ID or block sequence")
		}

		ids[id], positions[pos] = true, true

		if transfer.Settlement.Status == StatusRefundPending && transfer.Settlement.GovernanceEvidence != nil && !transfer.VoidOrdered {
			return ErrInvalidGenesis.Wrap("governance refund authorization requires a void order")
		}

		if !transfer.Settlement.Terminal() && transfer.Data.Amount.IsGT(g.Params.MaxEpochRefund) {
			return ErrInvalidGenesis.Wrap("unpaid transfer exceeds refund cap")
		}

		if err = totals.Add(transfer.Settlement.Status, transfer.Data.Amount); err != nil {
			return err
		}

		amount, exists := amounts[id.Owner]
		if !exists {
			amount = ZeroAmount()
		}

		amounts[id.Owner], err = AddAmount(amount, transfer.Data.Amount)
		if err != nil {
			return err
		}
	}

	if !totals.Equal(g.Totals) {
		return ErrInvalidGenesis.Wrap("totals do not match financial status populations")
	}

	owners := make(map[string]bool, len(g.Accounts))
	epochAmount := ZeroAmount()
	for _, account := range g.Accounts {
		if _, err = ParseAccount(account.Owner, AccountLength); err != nil {
			return err
		}

		if owners[account.Owner] {
			return ErrInvalidGenesis.Wrap("duplicate owner usage")
		}

		owners[account.Owner] = true

		usage, err := account.Usage.At(g.SendUsage.Epoch)
		if err != nil {
			return err
		}

		amount, exists := amounts[account.Owner]
		if !exists {
			amount = ZeroAmount()
		}

		if !usage.TotalSent.Equal(amount) {
			return ErrInvalidGenesis.Wrap("account lifetime amount differs from transfer records")
		}

		epochAmount, err = AddAmount(epochAmount, usage.EpochSent)
		if err != nil {
			return err
		}
	}

	for owner := range amounts {
		if !owners[owner] {
			return ErrInvalidGenesis.Wrap("missing owner usage")
		}
	}

	if !epochAmount.Equal(g.SendUsage.Amount) {
		return ErrInvalidGenesis.Wrap("global send usage differs from account epoch amounts")
	}

	if g.RefundUsage.Amount.IsGT(totals.RefundedAmount) {
		return ErrInvalidGenesis.Wrap("refund usage exceeds total paid refunds")
	}

	return nil
}
