package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func DefaultGenesis() *GenesisState {
	params := DefaultParams()

	return &GenesisState{
		Params:         params,
		Totals:         DefaultTotals(),
		PolicySequence: 0,
		Policies:       []ReporterPolicy{},
		SendUsage:      EpochUsage{Amount: ZeroAmount()},
		RefundUsage:    EpochUsage{Amount: ZeroAmount()},
	}
}

// ValidateBasic checks genesis state without store access. Beyond the SDK
// shape checks it requires the reporter policy that params describe: without
// a policy sequence the reporter set is empty, otherwise the policy at the
// sequence exists and matches params. SetParams and every reporter command
// load that policy, so a genesis without it could only be repaired by a chain
// upgrade.
func (g *GenesisState) ValidateBasic() error {
	var err error
	if err = g.Params.ValidateBasic(); err != nil {
		return err
	}

	if g.PolicySequence != uint64(len(g.Policies)) {
		return ErrInvalidGenesis.Wrap("invalid policies")
	}

	pHistory := make(map[uint64]ReporterPolicy, len(g.Policies))
	for _, policy := range g.Policies {
		if err = policy.Validate(); err != nil {
			return err
		}

		if _, exists := pHistory[policy.Version]; exists || policy.Version > g.PolicySequence {
			return ErrInvalidGenesis.Wrap("duplicate or future reporter policy")
		}

		pHistory[policy.Version] = policy
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
	ids := make(map[ID]bool, len(g.Records))

	type position struct {
		height   int64
		sequence uint64
	}

	positions := make(map[position]bool, len(g.Records))
	for _, transfer := range g.Records {
		if err = transfer.ID.ValidateBasic(); err != nil {
			return err
		}

		if err = transfer.Data.ValidateBasic(); err != nil {
			return err
		}

		if err = transfer.Settlement.Validate(transfer.ID, transfer.Status, pHistory); err != nil {
			return err
		}

		id := transfer.ID
		pos := position{id.Height, id.Sequence}
		if ids[id] || positions[pos] {
			return ErrInvalidGenesis.Wrap("duplicate transfer ID or block sequence")
		}

		ids[id], positions[pos] = true, true

		if !transfer.Status.Terminal() && transfer.Data.Amount.IsGT(g.Params.MaxEpochRefund) {
			return ErrInvalidGenesis.Wrap("unpaid transfer exceeds refund cap")
		}

		if err = totals.Add(transfer.Status, transfer.Data.Amount); err != nil {
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
		if _, err = sdk.AccAddressFromBech32(account.Owner); err != nil {
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

func policyHistory(policies []ReporterPolicy, sequence uint64) (map[uint64]ReporterPolicy, error) {
	history := make(map[uint64]ReporterPolicy, len(policies))

	for _, policy := range policies {
		if err := policy.Validate(); err != nil {
			return nil, err
		}

		if _, exists := history[policy.Version]; exists || policy.Version > sequence {
			return nil, ErrInvalidGenesis.Wrap("duplicate or future reporter policy")
		}

		history[policy.Version] = policy
	}

	return history, nil
}
