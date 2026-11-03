package cli

import (
	"os"
	"strconv"

	"github.com/mr-tron/base58"
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	cflags "pkg.akt.dev/go/cli/flags"
	types "pkg.akt.dev/go/node/sink/v1"
)

// GetTxSinkCmd returns the transaction commands for the sink module
func GetTxSinkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Sink transaction subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		GetTxSinkDepositCmd(),
		GetTxSinkGuardianPauseCmd(),
		GetTxSinkRequestVoidCmd(),
		GetTxSinkReportOutcomeCmd(),
		GetTxSinkVetoVoidCmd(),
		GetTxSinkRestartTransferCmd(),
		GetTxSinkConfirmMintCmd(),
		GetTxSinkConfirmVoidCmd(),
		GetTxSinkFinalizeRefundCmd(),
		GetTxSinkUpdateParamsCmd(),
	)

	return cmd
}

// GetTxSinkDepositCmd locks AKT for a Solana destination through MsgSink
func GetTxSinkDepositCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [destination] [amount]",
		Short: "Lock AKT for a Solana destination and create a sink transfer record",
		Long: `Lock AKT for a Solana destination and create a sink transfer record.
The amount is a whole number of uakt (1500000uakt) or an exact AKT decimal (1.5akt).
The record ID is returned in MsgSinkResponse once the transaction is included.

Example:
  $ akash tx sink deposit <solana-address> 1000000uakt --from mykey`,
		Args:              cobra.ExactArgs(2),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			amount, err := parseSinkAmount(args[1])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)

			msg := &types.MsgSink{
				DestType: types.DestTypeSolana,
				Dest:     args[0],
				Amount:   amount,
				Sender:   cl.ClientContext().GetFromAddress().String(),
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetTxSinkGuardianPauseCmd disables deposits, refunds or both. The signer
// must be params.pause_guardian; only governance can enable either again.
func GetTxSinkGuardianPauseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "guardian-pause",
		Short: "Disable sink deposits, refunds or both as the pause guardian",
		Long: `Sign MsgGuardianPause from the pause guardian account.
Pass --deposits, --refunds or both. Only governance can enable deposits or refunds again.

Example:
  $ akash tx sink guardian-pause --deposits --refunds --from guardian`,
		Args:              cobra.NoArgs,
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deposits, err := cmd.Flags().GetBool("deposits")
			if err != nil {
				return err
			}

			refunds, err := cmd.Flags().GetBool("refunds")
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)

			msg := &types.MsgGuardianPause{
				Guardian: cl.ClientContext().GetFromAddress().String(),
				Deposits: deposits,
				Refunds:  refunds,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cmd.Flags().Bool("deposits", false, "Disable deposits (params.deposits_enabled)")
	cmd.Flags().Bool("refunds", false, "Disable refunds (params.refunds_enabled)")
	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

// parseSinkAmount converts uakt or AKT amounts exactly. sdk.ParseCoinNormalized
// truncates fractions, which would silently turn 1.5uakt into 1uakt.
func parseSinkAmount(value string) (sdk.Coin, error) {
	dec, err := sdk.ParseDecCoin(value)
	if err != nil {
		return sdk.Coin{}, err
	}

	amount, fraction := sdk.NormalizeDecCoin(dec).TruncateDecimal()
	if !fraction.IsZero() {
		return sdk.Coin{}, sdkerrors.ErrInvalidCoins.Wrapf("amount %s is not a whole number of uakt", value)
	}

	return amount, nil
}

// GetTxSinkRequestVoidCmd orders a transfer void through governance
func GetTxSinkRequestVoidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-void [owner/height/sequence]",
		Short: "Order a transfer void through governance",
		Long: `Generate MsgRequestVoid for a governance proposal. Solana can prove the order.
It does not confirm the void on Akash and cannot release escrow.

Example:
  $ akash tx sink request-void <owner>/1200/1 --generate-only --from mykey`,
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.ParseID(args[0])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			signer, err := cmd.Flags().GetString(cflags.FlagAuthority)
			if err != nil {
				return err
			}

			msg := &types.MsgRequestVoid{Authority: signer, ID: id}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cmd.Flags().String(cflags.FlagAuthority, authtypes.NewModuleAddress(govtypes.ModuleName).String(), "Governance authority; use --generate-only when preparing a proposal")
	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetTxSinkReportOutcomeCmd reports a finalized Solana outcome as a current reporter
func GetTxSinkReportOutcomeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-outcome [owner/height/sequence] [policy-version] [round] [minted|voided] [slot] [signature]",
		Short: "Report a finalized Solana outcome",
		Long: `Report a finalized Solana outcome as a member of the current reporter policy.
The policy version must be current and the round must match the transfer's settlement round.
The slot and base58 transaction signature identify the Solana evidence.

Example:
  $ akash tx sink report-outcome <owner>/1200/1 1 1 minted 312345678 <signature> --from reporter`,
		Args:              cobra.ExactArgs(6),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.ParseID(args[0])
			if err != nil {
				return err
			}

			version, err := parseSinkPositiveUint(args[1])
			if err != nil {
				return err
			}

			round, err := parseSinkPositiveUint(args[2])
			if err != nil {
				return err
			}

			var outcome types.Outcome
			switch args[3] {
			case "minted":
				outcome = types.OutcomeMinted
			case "voided":
				outcome = types.OutcomeVoided
			default:
				return types.ErrInvalidObservation.Wrap("expected minted or voided")
			}

			evidence, err := parseSinkEvidence(args[4], args[5])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			signer := cl.ClientContext().GetFromAddress().String()
			msg := &types.MsgReportOutcome{
				Reporter:      signer,
				ID:            id,
				PolicyVersion: version,
				Round:         round,
				Observation:   types.Observation{Outcome: outcome, Evidence: evidence},
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetTxSinkVetoVoidCmd holds a refund-pending transfer for governance review
func GetTxSinkVetoVoidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "veto-void [owner/height/sequence] [reason]",
		Short: "Hold a refund for governance review",
		Long: `Hold a refund-pending transfer for governance review.
Current reporters can veto until the challenge window ends. Governance can veto at any time.
The reason must contain 1 to 256 UTF-8 bytes.

Example:
  $ akash tx sink veto-void <owner>/1200/1 "void signature not finalized" --from reporter`,
		Args:              cobra.ExactArgs(2),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.ParseID(args[0])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			signer := cl.ClientContext().GetFromAddress().String()
			msg := &types.MsgVetoVoid{Signer: signer, ID: id, Reason: args[1]}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetTxSinkRestartTransferCmd generates a governance restart of a settlement round
func GetTxSinkRestartTransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart-transfer [owner/height/sequence] [expected-round] [reason]",
		Short: "Generate a governance restart for the expected settlement round",
		Long: `Generate MsgRestartTransfer for a governance proposal.
A transfer that is neither minted nor refunded returns to pending with no votes or holds.
Its settlement round increments and it moves to the current reporter policy.
The expected round must equal the transfer's current round.

Example:
  $ akash tx sink restart-transfer <owner>/1200/1 2 "reporter policy rotated" --generate-only --from mykey`,
		Args:              cobra.ExactArgs(3),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.ParseID(args[0])
			if err != nil {
				return err
			}

			round, err := parseSinkPositiveUint(args[1])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			signer, err := cmd.Flags().GetString(cflags.FlagAuthority)
			if err != nil {
				return err
			}

			msg := &types.MsgRestartTransfer{Authority: signer, ID: id, ExpectedRound: round, Reason: args[2]}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cmd.Flags().String(cflags.FlagAuthority, authtypes.NewModuleAddress(govtypes.ModuleName).String(), "Governance authority; use --generate-only when preparing a proposal")
	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetTxSinkConfirmMintCmd generates governance confirmation of a Solana mint
func GetTxSinkConfirmMintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-mint [owner/height/sequence] [expected-round] [slot] [signature] [reason]",
		Short: "Generate governance mint confirmation for the expected round",
		Long: `Generate MsgConfirmMint for a governance proposal.
It records the mint evidence, marks the transfer minted and increments its settlement round.
The expected round must equal the transfer's current round.
A refunded transfer cannot be confirmed as minted.

Example:
  $ akash tx sink confirm-mint <owner>/1200/1 2 312345678 <signature> "mint finalized on Solana" --generate-only --from mykey`,
		Args:              cobra.ExactArgs(5),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.ParseID(args[0])
			if err != nil {
				return err
			}

			round, err := parseSinkPositiveUint(args[1])
			if err != nil {
				return err
			}

			evidence, err := parseSinkEvidence(args[2], args[3])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			signer, err := cmd.Flags().GetString(cflags.FlagAuthority)
			if err != nil {
				return err
			}

			msg := &types.MsgConfirmMint{Authority: signer, ID: id, ExpectedRound: round, Evidence: evidence, Reason: args[4]}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cmd.Flags().String(cflags.FlagAuthority, authtypes.NewModuleAddress(govtypes.ModuleName).String(), "Governance authority; use --generate-only when preparing a proposal")
	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetTxSinkConfirmVoidCmd generates governance confirmation of a Solana void
func GetTxSinkConfirmVoidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-void [owner/height/sequence] [expected-round] [slot] [signature] [reason]",
		Short: "Generate governance void confirmation for the expected round",
		Long: `Generate MsgConfirmVoid for a governance proposal.
It records the void evidence and increments the settlement round.
Unless already refunded, the transfer moves to refund pending with a new challenge window.
That move requires a void order from request-void and an amount within the epoch refund cap.
A minted transfer without a conflict hold cannot be voided.
The expected round must equal the transfer's current round.

Example:
  $ akash tx sink confirm-void <owner>/1200/1 2 312345678 <signature> "void finalized on Solana" --generate-only --from mykey`,
		Args:              cobra.ExactArgs(5),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.ParseID(args[0])
			if err != nil {
				return err
			}

			round, err := parseSinkPositiveUint(args[1])
			if err != nil {
				return err
			}

			evidence, err := parseSinkEvidence(args[2], args[3])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			signer, err := cmd.Flags().GetString(cflags.FlagAuthority)
			if err != nil {
				return err
			}

			msg := &types.MsgConfirmVoid{Authority: signer, ID: id, ExpectedRound: round, Evidence: evidence, Reason: args[4]}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cmd.Flags().String(cflags.FlagAuthority, authtypes.NewModuleAddress(govtypes.ModuleName).String(), "Governance authority; use --generate-only when preparing a proposal")
	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetTxSinkFinalizeRefundCmd pays the eligible refund queue head to its original owner
func GetTxSinkFinalizeRefundCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "finalize-refund [owner/height/sequence]",
		Short: "Pay the eligible refund queue head to its original owner",
		Long: `Pay a refund-pending transfer back to its original owner. Any account can submit it.
The transfer must be the refund queue head, with current authorization and no hold.
Its challenge window must have elapsed, and the refund must fit the epoch refund cap.

Example:
  $ akash tx sink finalize-refund <owner>/1200/1 --from mykey`,
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.ParseID(args[0])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			signer := cl.ClientContext().GetFromAddress().String()
			msg := &types.MsgFinalizeRefund{Executor: signer, ID: id}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetTxSinkUpdateParamsCmd generates a governance update of the sink params
func GetTxSinkUpdateParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-params [params.json]",
		Short: "Generate a governance parameter update",
		Long: `Generate MsgUpdateParams for a governance proposal.
The file holds the complete sink params as JSON.
A changed reporter set or threshold creates the next reporter policy version.

Example:
  $ akash tx sink update-params params.json --generate-only --from mykey`,
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			data, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			var params types.Params
			if err = cl.ClientContext().Codec.UnmarshalJSON(data, &params); err != nil {
				return err
			}

			signer, err := cmd.Flags().GetString(cflags.FlagAuthority)
			if err != nil {
				return err
			}

			msg := &types.MsgUpdateParams{Authority: signer, Params: params}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cmd.Flags().String(cflags.FlagAuthority, authtypes.NewModuleAddress(govtypes.ModuleName).String(), "Governance authority; use --generate-only when preparing a proposal")
	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

func parseSinkPositiveUint(value string) (uint64, error) {
	number, err := strconv.ParseUint(value, 10, 64)
	if err != nil || number == 0 || strconv.FormatUint(number, 10) != value {
		return 0, types.ErrInvalidResolution.Wrap("expected a canonical positive decimal integer")
	}

	return number, nil
}

func parseSinkEvidence(slot, signature string) (types.SolanaEvidence, error) {
	number, err := parseSinkPositiveUint(slot)
	if err != nil {
		return types.SolanaEvidence{}, err
	}

	bytes, err := base58.Decode(signature)
	if err != nil || base58.Encode(bytes) != signature {
		return types.SolanaEvidence{}, types.ErrInvalidObservation.Wrap("expected a canonical base58 signature")
	}

	evidence := types.SolanaEvidence{Slot: number, TransactionSignature: bytes}
	return evidence, evidence.Validate()
}
