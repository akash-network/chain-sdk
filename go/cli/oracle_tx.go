package cli

import (
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	cflags "pkg.akt.dev/go/cli/flags"
	types "pkg.akt.dev/go/node/oracle/v1"
)

// GetTxOracleCmd returns the transaction commands for oracle module
func GetTxOracleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Oracle transaction subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		GetTxOracleFeedPriceCmd(),
	)

	return cmd
}

func GetTxOracleFeedPriceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "feed [asset-denom] [base-denom]",
		Short:             "Feed price for denom",
		Args:              cobra.ExactArgs(2),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)

			// Validate denoms
			if err := sdk.ValidateDenom(args[0]); err != nil {
				return err
			}
			if err := sdk.ValidateDenom(args[1]); err != nil {
				return err
			}

			// Get signer address from client context
			cctx, err := GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgAddPriceEntry{
				Signer: cctx.GetFromAddress().String(),
				ID: types.DataID{
					Denom:     args[0],
					BaseDenom: args[1],
				},
				Price: types.PriceDataState{},
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
