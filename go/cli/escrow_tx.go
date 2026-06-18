package cli

import (
	"errors"
	"fmt"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	cflags "pkg.akt.dev/go/cli/flags"
	eid "pkg.akt.dev/go/node/escrow/id/v1"
	emodule "pkg.akt.dev/go/node/escrow/module"
	ev1 "pkg.akt.dev/go/node/escrow/v1"
	deposit "pkg.akt.dev/go/node/types/deposit/v1"
)

func GetTxEscrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        emodule.ModuleName,
		Short:                      "Escrow transaction commands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}
	cmd.AddCommand(
		GetTxEscrowDeposit(),
	)

	return cmd
}

func GetTxEscrowDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "deposit [deployment] [amount]",
		Short:             "deposit funds to escrow account",
		Args:              cobra.ExactArgs(2),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			var aid eid.Account

			switch args[0] {
			case "deployment":
				id, err := cflags.DeploymentIDFromFlags(cmd.Flags(), cflags.WithOwner(cctx.FromAddress))
				if err != nil {
					return err
				}
				aid = id.ToEscrowAccountID()
			default:
				return fmt.Errorf("invalid account scope. allowed values deployment")
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			sources, err := DepositSources(cmd.Flags())
			if err != nil {
				return err
			}

			msg := &ev1.MsgAccountDeposit{
				ID:     aid,
				Signer: cctx.FromAddress.String(),
				Deposit: deposit.Deposit{
					Amount:  amount,
					Sources: sources,
				},
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return annotateEscrowDepositError(err)
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	cflags.AddDeploymentIDFlags(cmd.Flags())
	cflags.AddDepositSourcesFlags(cmd.Flags())

	return cmd
}

// annotateEscrowDepositError adds a best-effort hint when the escrow account
// derived from the deployment flags does not exist. The escrow account ID is
// built from --owner (defaulting to the signer) and --dseq, so a missing,
// wrong, or mistyped value all surface as the same opaque "account not found".
// Matching the registered ErrAccountNotFound keeps this resilient to message
// rewording and scoped to the escrow module. Other errors are left untouched.
func annotateEscrowDepositError(err error) error {
	if errors.Is(err, emodule.ErrAccountNotFound) {
		return fmt.Errorf("%w: if this is another account's deployment, pass its owner with --%s", err, cflags.FlagOwner)
	}
	return err
}
