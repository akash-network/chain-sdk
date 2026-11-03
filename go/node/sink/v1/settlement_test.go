package v1_test

import (
	"bytes"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
)

func account(b byte) string {
	return sdk.AccAddress(bytes.Repeat([]byte{b}, 20)).String()
}

func evidence() types.SolanaEvidence {
	return types.SolanaEvidence{Slot: 7, TransactionSignature: bytes.Repeat([]byte{8}, 64)}
}

func policy() types.ReporterPolicy {
	return types.ReporterPolicy{Version: 1, Members: []string{account(1), account(2), account(3), account(4)}, Threshold: 3}
}

func quorum(outcome types.Outcome) []types.Vote {
	votes := make([]types.Vote, 3)
	for i := range votes {
		votes[i] = types.Vote{Reporter: account(byte(i + 1)), PolicyVersion: 1, Observation: types.Observation{Outcome: outcome, Evidence: evidence()}}
	}

	return votes
}

func settled(status types.Status) types.SettlementData {
	s := types.SettlementData{Status: status, Round: 1, PolicyVersion: 1}
	if status == types.StatusMinted {
		s.Votes = quorum(types.OutcomeMinted)
	}

	if status == types.StatusRefundPending || status == types.StatusRefunded {
		s.Votes = quorum(types.OutcomeVoided)
	}

	if status == types.StatusRefundPending {
		deadline := time.Unix(3600, 0).UTC()
		s.ChallengeEnd = &deadline
	}

	return s
}

func TestSettlementRetainsAuthorization(t *testing.T) {
	policies := map[uint64]types.ReporterPolicy{1: policy(), 2: {Version: 2, Members: []string{account(9)}, Threshold: 1}}
	for _, status := range []types.Status{types.StatusTransferPending, types.StatusMinted, types.StatusRefundPending, types.StatusRefunded} {
		s := settled(status)
		require.NoError(t, s.Validate(policies))

		if status != types.StatusTransferPending {
			s.Votes = nil
			require.ErrorContains(t, s.Validate(policies), "retained quorum")
			e := evidence()
			s.GovernanceEvidence, s.Round = &e, 2
			require.NoError(t, s.Validate(policies))
		}
	}

	s := settled(types.StatusMinted)
	s.Votes = append(s.Votes, types.Vote{Reporter: account(9), PolicyVersion: 2, Observation: types.Observation{Outcome: types.OutcomeVoided, Evidence: evidence()}})
	require.ErrorContains(t, s.Validate(policies), "conflict")
	s.Hold = types.HoldConflict
	require.NoError(t, s.Validate(policies))
	minted, voided := s.Tallies()
	require.Equal(t, uint32(3), minted)
	require.Zero(t, voided)
	s.PolicyVersion = 2
	require.ErrorContains(t, s.Validate(policies), "retained quorum")
}

func TestSettlementShapes(t *testing.T) {
	policies := map[uint64]types.ReporterPolicy{1: policy()}
	for name, change := range map[string]func(*types.SettlementData){
		"unknown status":   func(s *types.SettlementData) { s.Status = 3 },
		"unknown hold":     func(s *types.SettlementData) { s.Hold = 3 },
		"zero round":       func(s *types.SettlementData) { s.Round = 0 },
		"missing policy":   func(s *types.SettlementData) { s.PolicyVersion = 2 },
		"missing deadline": func(s *types.SettlementData) { s.ChallengeEnd = nil },
		"zero deadline":    func(s *types.SettlementData) { s.ChallengeEnd = &time.Time{} },
		"duplicate vote":   func(s *types.SettlementData) { s.Votes = append(s.Votes, s.Votes[0]) },
		"outsider":         func(s *types.SettlementData) { s.Votes[0].Reporter = account(9) },
		"vote policy":      func(s *types.SettlementData) { s.Votes[0].PolicyVersion = 2 },
		"veto missing":     func(s *types.SettlementData) { s.Hold = types.HoldVeto },
		"orphan veto":      func(s *types.SettlementData) { s.Veto = &types.Veto{Signer: account(1), Reason: "verify"} },
		"opposing vote":    func(s *types.SettlementData) { s.Votes[0].Observation.Outcome = types.OutcomeMinted },
		"empty evidence":   func(s *types.SettlementData) { s.Votes[0].Observation.Evidence = types.SolanaEvidence{} },
		"governance round one": func(s *types.SettlementData) {
			e := evidence()
			s.GovernanceEvidence = &e
		},
	} {
		t.Run(name, func(t *testing.T) {
			s := settled(types.StatusRefundPending)
			change(&s)
			require.Error(t, s.Validate(policies))
		})
	}

	s := settled(types.StatusTransferPending)
	s.Hold, s.Veto = types.HoldVeto, &types.Veto{Signer: account(1), Reason: "retained through rotation"}
	require.NoError(t, s.Validate(policies))
	s.Hold = types.HoldConflict
	require.NoError(t, s.Validate(policies))
	s = settled(types.StatusRefundPending)
	require.True(t, s.Payable(policy(), 1))
	require.False(t, s.Payable(policy(), 2))
	scope, ok := s.RefundScope()
	require.True(t, ok)
	require.Equal(t, uint64(1), scope)
	e := evidence()
	s.GovernanceEvidence, s.Round = &e, 2
	require.True(t, s.Payable(types.ReporterPolicy{Version: 2}, 2))
	scope, ok = s.RefundScope()
	require.True(t, ok)
	require.Zero(t, scope)
	s.Hold = types.HoldConflict
	require.False(t, s.Payable(policy(), 1))
	_, ok = s.RefundScope()
	require.False(t, ok)
}

func TestSettlementValidatesReferencedPolicies(t *testing.T) {
	rotated := types.ReporterPolicy{Version: 7, Members: []string{account(9)}, Threshold: 1}
	s := settled(types.StatusMinted)
	require.NoError(t, s.Validate(map[uint64]types.ReporterPolicy{1: policy()}), "unreferenced versions are not required")
	require.ErrorContains(t, s.Validate(map[uint64]types.ReporterPolicy{7: rotated}), "missing authorization policy")
	s.Hold = types.HoldConflict
	s.Votes = append(s.Votes, types.Vote{Reporter: account(9), PolicyVersion: 7, Observation: types.Observation{Outcome: types.OutcomeVoided, Evidence: evidence()}})
	require.NoError(t, s.Validate(map[uint64]types.ReporterPolicy{1: policy(), 7: rotated}))
	require.ErrorContains(t, s.Validate(map[uint64]types.ReporterPolicy{1: policy()}), "vote outside its reporter policy")
}
