package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

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
			destType, err := cmd.Flags().GetUint32("dest-type")
			if err != nil {
				return err
			}

			if destType != types.DestTypeSolana {
				return fmt.Errorf("destination type must be 1 (Solana)")
			}

			amount, err := parseSinkAmount(args[1])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)

			msg := &types.MsgSink{
				DestType: destType,
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

	cmd.Flags().Uint32("dest-type", types.DestTypeSolana, "Destination network (1 = Solana)")
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

			if !deposits && !refunds {
				return fmt.Errorf("pass --deposits, --refunds or both")
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

	cmd.Flags().Bool("deposits", false, "Disable deposits (params.enabled)")
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
		return sdk.Coin{}, fmt.Errorf("amount %s is not a whole number of uakt", value)
	}

	if amount.Denom != types.Denom || !amount.Amount.IsPositive() || !amount.Amount.IsUint64() {
		return sdk.Coin{}, fmt.Errorf("amount must be positive uint64 uakt")
	}

	return amount, nil
}
