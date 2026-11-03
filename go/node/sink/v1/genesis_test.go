package v1_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
	"pkg.akt.dev/go/sdkutil"
)

func fullGenesis(t *testing.T) *types.GenesisState {
	g := types.DefaultGenesis()
	g.Params.Reporters = policy().Members
	g.Policies = []types.ReporterPolicy{g.Params.Policy(1)}
	g.SendUsage.Epoch, g.RefundUsage.Epoch = 10, 9
	statuses := []types.Status{types.StatusTransferPending, types.StatusMinted, types.StatusRefundPending, types.StatusRefunded, types.StatusTransferPending, types.StatusRefundPending, types.StatusMinted, types.StatusRefunded, types.StatusRefundPending}
	amount := g.Params.MinAmountPerTransfer
	owner := account(7)
	total := types.ZeroAmount()
	for i, status := range statuses {
		s := settled(status)
		if i == 4 || i == 5 {
			s.Hold, s.Veto = types.HoldVeto, &types.Veto{Signer: account(1), Reason: "check"}
		}

		if i == 6 || i == 7 {
			s.Hold = types.HoldConflict
		}

		if i == 8 {
			e := evidence()
			s.GovernanceEvidence, s.Round, s.Votes = &e, 2, nil
		}

		g.Transfers = append(g.Transfers, types.TransferState{ID: types.ID{Owner: owner, Height: 2, Sequence: uint64(i + 1)}, Data: types.RecordData{Dest: base58.Encode(makeRecipient()), Amount: amount}, Settlement: s, VoidOrdered: i == 8})
		require.NoError(t, g.Totals.Add(status, amount))
		var err error
		total, err = types.AddAmount(total, amount)
		require.NoError(t, err)
	}

	g.Accounts = []types.GenesisAccount{
		{Owner: owner, Usage: types.AccountUsage{TotalSent: total, EpochSent: amount, LastEpoch: 10}},
		{Owner: account(8), Usage: types.AccountUsage{TotalSent: types.ZeroAmount(), EpochSent: types.ZeroAmount(), LastEpoch: 0}},
	}

	g.SendUsage.Amount, g.RefundUsage.Amount = amount, amount
	require.NoError(t, g.Validate())

	return g
}

func makeRecipient() []byte {
	a := make([]byte, 32)
	a[0] = 2
	return a
}

func TestFullGenesisRoundTrip(t *testing.T) {
	g := fullGenesis(t)
	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())
	encoded, err := cdc.MarshalJSON(g)
	require.NoError(t, err)
	var restored types.GenesisState
	require.NoError(t, cdc.UnmarshalJSON(encoded, &restored))
	require.NoError(t, restored.Validate())
	again, err := cdc.MarshalJSON(&restored)
	require.NoError(t, err)
	require.JSONEq(t, string(encoded), string(again))
	binary, err := g.Marshal()
	require.NoError(t, err)
	restored = types.GenesisState{}
	require.NoError(t, restored.Unmarshal(binary))
	require.NoError(t, restored.Validate())
}

func TestGenesisRejectsLostOrContradictoryState(t *testing.T) {
	for name, change := range map[string]func(*types.GenesisState){
		"old schema":     func(g *types.GenesisState) { g.SchemaVersion = 1 },
		"fixed days":     func(g *types.GenesisState) { g.SchemaVersion = 2 },
		"missing schema": func(g *types.GenesisState) { g.SchemaVersion = 0 },
		"duplicate ID":   func(g *types.GenesisState) { g.Transfers = append(g.Transfers, g.Transfers[0]) },
		"cross owner reused sequence": func(g *types.GenesisState) {
			g.Transfers[1].ID = g.Transfers[0].ID
			g.Transfers[1].ID.Owner = account(9)
		},
		"zero height":             func(g *types.GenesisState) { g.Transfers[0].ID.Height = 0 },
		"missing settlement":      func(g *types.GenesisState) { g.Transfers[0].Settlement = types.SettlementData{} },
		"missing minted quorum":   func(g *types.GenesisState) { g.Transfers[1].Settlement.Votes = nil },
		"missing refund quorum":   func(g *types.GenesisState) { g.Transfers[2].Settlement.Votes = nil },
		"missing paid quorum":     func(g *types.GenesisState) { g.Transfers[3].Settlement.Votes = nil },
		"missing void order":      func(g *types.GenesisState) { g.Transfers[8].VoidOrdered = false },
		"duplicate owner":         func(g *types.GenesisState) { g.Accounts = append(g.Accounts, g.Accounts[0]) },
		"missing owner":           func(g *types.GenesisState) { g.Accounts = g.Accounts[1:] },
		"lifetime mismatch":       func(g *types.GenesisState) { g.Accounts[0].Usage.TotalSent = g.Params.MinAmountPerTransfer },
		"future account stamp":    func(g *types.GenesisState) { g.Accounts[0].Usage.LastEpoch++ },
		"negative account stamp":  func(g *types.GenesisState) { g.Accounts[0].Usage.LastEpoch = -1 },
		"send consumption":        func(g *types.GenesisState) { g.SendUsage.Amount = types.ZeroAmount() },
		"refund consumption":      func(g *types.GenesisState) { g.RefundUsage.Amount = g.Params.MaxEpochRefund },
		"missing policies":        func(g *types.GenesisState) { g.Policies = nil },
		"duplicate policies":      func(g *types.GenesisState) { g.Policies = append(g.Policies, g.Policies[0]) },
		"missing policy version":  func(g *types.GenesisState) { g.Policies[0].Version = 2 },
		"current policy mismatch": func(g *types.GenesisState) { g.Params.Reporters = nil },
		"wrong status count":      func(g *types.GenesisState) { g.Totals.PendingCount++ },
		"nil Coin":                func(g *types.GenesisState) { g.Totals = types.Totals{} },
	} {
		t.Run(name, func(t *testing.T) {
			g := fullGenesis(t)
			change(g)
			require.Error(t, g.Validate())
		})
	}
}

func TestGenesisTerminalAuthorizationSurvivesRotation(t *testing.T) {
	g := fullGenesis(t)
	g.Params.Reporters = []string{account(9)}
	g.Policies = append(g.Policies, g.Params.Policy(2))
	require.NoError(t, g.Validate())

	for _, i := range []int{1, 3, 6, 7} {
		g.Transfers[i].Settlement.Votes = append(g.Transfers[i].Settlement.Votes, types.Vote{Reporter: account(9), PolicyVersion: 2, Observation: types.Observation{Outcome: types.OutcomeMinted, Evidence: evidence()}})
		g.Transfers[i].Settlement.Hold = types.HoldConflict
	}

	require.NoError(t, g.Validate())
	g.Transfers[6].Settlement.Votes = g.Transfers[6].Settlement.Votes[3:]
	require.ErrorContains(t, g.Validate(), "retained quorum")
}

func TestGenesisRejectsOldShapesAndKeepsRawUsage(t *testing.T) {
	cdc := codec.NewProtoCodec(sdkutil.NewCodecOptions().NewInterfaceRegistry())
	var old types.GenesisState
	require.NoError(t, cdc.UnmarshalJSON([]byte(`{}`), &old))
	require.ErrorContains(t, old.Validate(), "schema_version")
	g := fullGenesis(t)
	g.Params.MaxAmountPerTransfer = g.Params.MinAmountPerTransfer
	g.Params.MaxEpochAmountPerAccount = g.Params.MinAmountPerTransfer
	g.Params.MaxTransferAmountPerEpoch = g.Params.MinAmountPerTransfer
	g.Params.MaxEpochRefund = g.Params.MinAmountPerTransfer
	amount, err := types.AddAmount(g.SendUsage.Amount, g.SendUsage.Amount)
	require.NoError(t, err)
	g.SendUsage.Amount, g.Accounts[0].Usage.EpochSent = amount, amount
	require.NoError(t, g.Validate())
	g.Accounts[0].Usage.LastEpoch--
	g.SendUsage.Amount = types.ZeroAmount()
	require.NoError(t, g.Validate())
	require.False(t, g.Accounts[0].Usage.EpochSent.IsZero())
}
