package v1_test

import (
	"bytes"
	"math"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
)

func TestReporterPolicyBoundaries(t *testing.T) {
	members := make([]string, 4)
	for i := range members {
		members[i] = sdk.AccAddress(bytes.Repeat([]byte{byte(i) + 1}, 20)).String()
	}
	require.NoError(t, types.ValidateReporters(nil, 0))
	require.NoError(t, types.ValidateReporters(members, 3))
	require.NoError(t, types.ValidateReporters(members, 4))
	require.Error(t, types.ValidateReporters(members[:3], 2))
	require.Error(t, types.ValidateReporters(members, 2))
	require.Error(t, types.ValidateReporters(members, math.MaxUint32))
	require.Error(t, types.ValidateReporters(nil, 1))
	require.Error(t, types.ValidateReporters([]string{members[0], members[0]}, 2))
	require.Error(t, types.ValidateReporters([]string{"invalid"}, 1))
	p := types.DefaultParams()
	p.Enabled = true
	require.Error(t, p.Validate())
	p.Reporters, p.ReporterThreshold = members, 3
	require.ErrorContains(t, p.Validate(), "min transfer must be set") // deposits need a minimum
	p.MinTransfer = 1
	require.NoError(t, p.Validate())
	p.Enabled = false
	p.MinTransfer = 0
	require.NoError(t, p.Validate()) // disabling deposits never fails on the minimum
	p.Enabled, p.MinTransfer = true, 1
	p.ChallengePeriodSeconds = 0
	require.Error(t, p.Validate())
	p.ChallengePeriodSeconds = types.MinChallengePeriodSeconds - 1
	require.ErrorContains(t, p.Validate(), "challenge period")
	p.RefundsEnabled = false
	require.Error(t, p.Validate()) // the floor applies while refunds are off too
	p.ChallengePeriodSeconds = types.MinChallengePeriodSeconds
	require.NoError(t, p.Validate())
	p.ChallengePeriodSeconds = math.MaxUint64
	require.Error(t, p.Validate())
	p.ChallengePeriodSeconds = types.MinChallengePeriodSeconds
	p.DailyRefundCap = p.MaxTransfer - 1
	require.Error(t, p.Validate())
	require.Equal(t, types.DefaultChallengePeriodSeconds, types.DefaultParams().ChallengePeriodSeconds)
	require.NoError(t, types.DefaultParams().Validate())
}

func TestTalliesCountOnlyKnownOutcomes(t *testing.T) {
	vote := func(o types.Outcome) types.Vote {
		return types.Vote{Observation: types.Observation{Outcome: o}}
	}
	s := types.Settlement{Votes: []types.Vote{vote(types.Outcome_OUTCOME_MINTED), vote(types.Outcome_OUTCOME_VOIDED),
		vote(types.Outcome_OUTCOME_UNSPECIFIED), vote(types.Outcome(7))}}
	minted, voided := s.Tallies()
	require.Equal(t, uint32(1), minted)
	require.Equal(t, uint32(1), voided)
}

func TestSettlementValidationAndAccounting(t *testing.T) {
	member := sdk.AccAddress(bytes.Repeat([]byte{1}, 20)).String()
	o := types.Observation{Outcome: types.Outcome_OUTCOME_VOIDED, Slot: 100, TransactionSignature: bytes.Repeat([]byte{1}, 64)}
	valid := types.Settlement{Status: types.Status_STATUS_VOIDED, Round: 1, ChallengeEnd: 100,
		Policy: types.ReporterPolicy{Epoch: 1, Members: []string{member}, Threshold: 1},
		Votes:  []types.Vote{{Reporter: member, Observation: o}}}
	require.NoError(t, valid.Validate(1))
	for name, mutate := range map[string]func(*types.Settlement){
		"zero round":            func(s *types.Settlement) { s.Round = 0 },
		"missing epoch":         func(s *types.Settlement) { s.Policy.Epoch = 0 },
		"future epoch":          func(s *types.Settlement) { s.Policy.Epoch = 2 },
		"missing deadline":      func(s *types.Settlement) { s.ChallengeEnd = 0 },
		"missing authorization": func(s *types.Settlement) { s.Votes = nil },
		"duplicate vote":        func(s *types.Settlement) { s.Votes = append(s.Votes, s.Votes[0]) },
		"unknown reporter":      func(s *types.Settlement) { s.Policy.Members = nil; s.Policy.Threshold = 0 },
		"unknown status":        func(s *types.Settlement) { s.Status = 99 },
		"unknown hold":          func(s *types.Settlement) { s.Hold = 99 },
		"veto without signer":   func(s *types.Settlement) { s.Hold = types.Hold_HOLD_VETO },
		"conflict without hold": func(s *types.Settlement) {
			mint := o
			mint.Outcome = types.Outcome_OUTCOME_MINTED
			s.Votes = append(s.Votes, types.Vote{Reporter: member, Observation: mint})
		},
	} {
		t.Run(name, func(t *testing.T) {
			s := valid
			s.Votes = append([]types.Vote(nil), valid.Votes...)
			mutate(&s)
			require.Error(t, s.Validate(1))
		})
	}
	for _, o := range []types.Observation{
		{}, {Outcome: 99, Slot: 1, TransactionSignature: bytes.Repeat([]byte{1}, 64)},
		{Outcome: types.Outcome_OUTCOME_MINTED, TransactionSignature: bytes.Repeat([]byte{1}, 64)},
		{Outcome: types.Outcome_OUTCOME_MINTED, Slot: 1, TransactionSignature: make([]byte, 64)},
	} {
		require.Error(t, o.Validate())
	}
	totals := types.Totals{LaunchReserve: 100, PendingAmount: 10, MintedAmount: 20, VoidedAmount: 30, RefundedAmount: 40}
	escrow, err := totals.EscrowAmount()
	require.NoError(t, err)
	require.Equal(t, uint64(160), escrow)
	deposited, err := totals.DepositedAmount()
	require.NoError(t, err)
	require.Equal(t, uint64(100), deposited)
	totals.PendingAmount = math.MaxUint64
	_, err = totals.EscrowAmount()
	require.Error(t, err)
	_, err = totals.DepositedAmount()
	require.Error(t, err)
}
