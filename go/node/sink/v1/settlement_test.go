package v1_test

import (
	"bytes"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
)

func account(b byte) string {
	return sdk.AccAddress(bytes.Repeat([]byte{b}, 20)).String()
}

func evidence() types.SolanaEvidence {
	return types.SolanaEvidence{
		Slot:                 7,
		TransactionSignature: bytes.Repeat([]byte{8}, 64),
	}
}

func TestSolanaEvidenceSignatureLength(t *testing.T) {
	for _, size := range []int{0, 1, solana.SignatureLength - 1, solana.SignatureLength, solana.SignatureLength + 1} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			e := types.SolanaEvidence{
				Slot:                 1,
				TransactionSignature: bytes.Repeat([]byte{1}, size),
			}
			if size == solana.SignatureLength {
				require.NoError(t, e.Validate())
				return
			}

			require.ErrorIs(t, e.Validate(), types.ErrInvalidObservation)
		})
	}
}

func observation(outcome types.Outcome) types.Observation {
	e := evidence()
	return types.Observation{
		Outcome:              outcome,
		Slot:                 e.Slot,
		TransactionSignature: e.TransactionSignature,
	}
}

func policy() types.ReporterPolicy {
	return types.ReporterPolicy{
		Version:   1,
		Members:   []string{account(1), account(2), account(3), account(4)},
		Threshold: 3,
	}
}

func quorum(outcome types.Outcome) []types.Vote {
	votes := make([]types.Vote, 3)
	for i := range votes {
		votes[i] = types.Vote{
			Reporter:      account(byte(i + 1)),
			PolicyVersion: 1,
			Observation:   observation(outcome),
		}
	}

	return votes
}

func settled(status types.Status) types.SettlementData {
	s := types.SettlementData{
		Round:  1,
		Policy: policy(),
	}
	if status == types.StatusMinted {
		s.Votes = quorum(types.OutcomeMinted)
	}

	if status == types.StatusRefundPending || status == types.StatusRefunded {
		s.Votes = quorum(types.OutcomeVoided)
	}

	if status == types.StatusRefundPending {
		s.ChallengeEnd = 3600
	}

	if status == types.StatusVoided {
		s.VoidAuthorization = &types.VoidAuthorization{
			RequestedHeight: 2,
			RequestedRound:  1,
		}
	}

	return s
}

func TestSettlementRetainsAuthorization(t *testing.T) {
	id := types.ID{
		Owner:    account(7),
		Height:   2,
		Sequence: 1,
	}
	policies := map[uint64]types.ReporterPolicy{1: policy(), 7: {
		Version:   7,
		Members:   []string{account(9)},
		Threshold: 1,
	}}
	for _, status := range []types.Status{types.StatusTransferPending, types.StatusVoided, types.StatusMinted, types.StatusRefundPending, types.StatusRefunded} {
		s := settled(status)
		require.NoError(t, s.Validate(id, status, policies))

		if status == types.StatusTransferPending || status == types.StatusVoided {
			continue
		}

		s.Votes = nil
		require.ErrorContains(t, s.Validate(id, status, policies), "retained quorum")
		o := observation(types.OutcomeVoided)
		if status == types.StatusMinted {
			o.Outcome = types.OutcomeMinted
		}

		s.GovernanceObservation, s.Round = &o, 2
		if status == types.StatusRefundPending {
			s.VoidAuthorization = &types.VoidAuthorization{
				RequestedHeight: 2,
				RequestedRound:  1,
			}
		}

		require.NoError(t, s.Validate(id, status, policies))
	}

	s := settled(types.StatusMinted)
	s.Votes = append(s.Votes, types.Vote{
		Reporter:      account(9),
		PolicyVersion: 7,
		Observation:   observation(types.OutcomeVoided),
	})
	require.ErrorContains(t, s.Validate(id, types.StatusMinted, policies), "conflict")
	s.Hold = types.HoldConflict
	require.NoError(t, s.Validate(id, types.StatusMinted, policies))
	require.ErrorContains(t, s.Validate(id, types.StatusMinted, map[uint64]types.ReporterPolicy{1: policy()}), "vote outside")
	minted, voided := s.Tallies()
	require.Equal(t, uint32(3), minted)
	require.Zero(t, voided)
	s.Votes = append(s.Votes, s.Votes[len(s.Votes)-1])
	require.ErrorContains(t, s.Validate(id, types.StatusMinted, policies), "bound")
}

func TestSettlementShapes(t *testing.T) {
	id := types.ID{
		Owner:    account(7),
		Height:   2,
		Sequence: 1,
	}
	policies := map[uint64]types.ReporterPolicy{1: policy()}
	for name, change := range map[string]func(*types.SettlementData){
		"unknown hold":         func(s *types.SettlementData) { s.Hold = 3 },
		"zero round":           func(s *types.SettlementData) { s.Round = 0 },
		"missing policy":       func(s *types.SettlementData) { s.Policy.Version = 2 },
		"different policy":     func(s *types.SettlementData) { s.Policy.Threshold = 4 },
		"missing deadline":     func(s *types.SettlementData) { s.ChallengeEnd = 0 },
		"negative deadline":    func(s *types.SettlementData) { s.ChallengeEnd = -1 },
		"timestamp overflow":   func(s *types.SettlementData) { s.ChallengeEnd = 253402300800 },
		"duplicate vote":       func(s *types.SettlementData) { s.Votes = append(s.Votes, s.Votes[0]) },
		"outsider":             func(s *types.SettlementData) { s.Votes[0].Reporter = account(9) },
		"vote policy":          func(s *types.SettlementData) { s.Votes[0].PolicyVersion = 2 },
		"veto missing":         func(s *types.SettlementData) { s.Hold = types.HoldVeto },
		"orphan veto":          func(s *types.SettlementData) { s.VetoBy, s.VetoReason = account(1), "verify" },
		"unsupported conflict": func(s *types.SettlementData) { s.Hold = types.HoldConflict },
		"opposing vote":        func(s *types.SettlementData) { s.Votes[0].Observation.Outcome = types.OutcomeMinted },
		"empty evidence":       func(s *types.SettlementData) { s.Votes[0].Observation.Slot = 0 },
		"governance round one": func(s *types.SettlementData) {
			o := observation(types.OutcomeVoided)
			s.GovernanceObservation = &o
		},
		"early receipt": func(s *types.SettlementData) {
			s.VoidAuthorization = &types.VoidAuthorization{
				RequestedHeight: 1,
				RequestedRound:  1,
			}
		},
		"future receipt round": func(s *types.SettlementData) {
			s.VoidAuthorization = &types.VoidAuthorization{
				RequestedHeight: 2,
				RequestedRound:  2,
			}
		},
	} {
		t.Run(name, func(t *testing.T) {
			s := settled(types.StatusRefundPending)
			change(&s)
			require.Error(t, s.Validate(id, types.StatusRefundPending, policies))
		})
	}

	s := settled(types.StatusRefundPending)
	require.True(t, s.Payable(types.StatusRefundPending, 1))
	require.False(t, s.Payable(types.StatusRefundPending, 7))
	o := observation(types.OutcomeVoided)
	s.GovernanceObservation, s.Round, s.Votes = &o, 2, nil
	s.VoidAuthorization = &types.VoidAuthorization{
		RequestedHeight: 2,
		RequestedRound:  1,
	}
	require.NoError(t, s.Validate(id, types.StatusRefundPending, policies))
	require.True(t, s.Payable(types.StatusRefundPending, 7))
	s.Votes = append(s.Votes, types.Vote{
		Reporter:      account(1),
		PolicyVersion: 1,
		Observation:   observation(types.OutcomeMinted),
	})
	s.Hold = types.HoldConflict
	require.NoError(t, s.Validate(id, types.StatusRefundPending, policies))
	require.False(t, s.Payable(types.StatusRefundPending, 1))
}

func TestSettlementVetoPolicyProvenance(t *testing.T) {
	id := types.ID{
		Owner:    account(7),
		Height:   2,
		Sequence: 1,
	}
	s := settled(types.StatusRefundPending)
	s.Hold, s.VetoBy, s.VetoReason, s.VetoPolicyVersion = types.HoldVeto, account(9), "verify", 7
	policies := map[uint64]types.ReporterPolicy{1: policy(), 7: {
		Version:   7,
		Members:   []string{account(9)},
		Threshold: 1,
	}}
	require.NoError(t, s.Validate(id, types.StatusRefundPending, policies))
	delete(policies, 7)
	require.ErrorContains(t, s.Validate(id, types.StatusRefundPending, policies), "veto outside")
}

func TestSettlementRejectsVetoOutsideRefundCandidate(t *testing.T) {
	id := types.ID{
		Owner:    account(7),
		Height:   2,
		Sequence: 1,
	}
	for _, status := range []types.Status{types.StatusTransferPending, types.StatusVoided} {
		s := settled(status)
		s.Votes = []types.Vote{
			{
				Reporter:      account(1),
				PolicyVersion: 1,
				Observation:   observation(types.OutcomeMinted),
			},
			{
				Reporter:      account(2),
				PolicyVersion: 1,
				Observation:   observation(types.OutcomeVoided),
			},
		}

		s.Hold, s.VetoBy, s.VetoReason, s.VetoPolicyVersion = types.HoldConflict, account(1), "retained veto", 1
		require.Error(t, s.Validate(id, status, map[uint64]types.ReporterPolicy{1: policy()}))
	}
}
