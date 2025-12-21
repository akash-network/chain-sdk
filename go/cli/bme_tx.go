package cli

import (
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	cflags "pkg.akt.dev/go/cli/flags"
	types "pkg.akt.dev/go/node/bme/v1"
)

// GetTxBMECmd returns the transaction commands for bme module
func GetTxBMECmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "BME transaction subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		GetTxBMEBurnMintCmd(),
	)

	return cmd
}

// GetTxBMEBurnMintCmd returns the command to burn one token and mint another
func GetTxBMEBurnMintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-mint [coins-to-burn] [denom-to-mint]",
		Short: "Burn tokens to mint another denomination",
		Long: `Burn tokens to mint another denomination.
This allows burning AKT to mint ACT, or burning unused ACT back to AKT.

Example:
  $ akash tx bme burn-mint 1000000uakt uact --from mykey
  $ akash tx bme burn-mint 500000uact uakt --from mykey`,
		Args:              cobra.ExactArgs(2),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)

			// Parse the coin to burn
			coinsToBurn, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			// Validate the denom to mint
			denomToMint := args[1]
			if err := sdk.ValidateDenom(denomToMint); err != nil {
				return err
			}

			// Get signer address from client context
			cctx, err := GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fromAddr := cctx.GetFromAddress().String()

			// Get optional 'to' address, defaults to owner
			toAddr, _ := cmd.Flags().GetString(cflags.FlagTo)
			if toAddr == "" {
				toAddr = fromAddr
			}

			msg := &types.MsgBurnMint{
				Owner:       fromAddr,
				To:          toAddr,
				CoinsToBurn: coinsToBurn,
				DenomToMint: denomToMint,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(cflags.FlagTo, "", "Destination address for minted coins (defaults to sender)")

	return cmd
}
