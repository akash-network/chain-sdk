package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"cosmossdk.io/core/address"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	cflags "pkg.akt.dev/go/cli/flags"
)

// GetTxVestingCmd returns vesting module's transaction commands.
func GetTxVestingCmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Vesting transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetTxVestingCreateAccountCmd(ac),
		GetTxVestingCreatePermanentLockedAccountCmd(ac),
		GetTxVestingCreatePeriodicAccountCmd(ac),
	)

	return cmd
}

// GetTxVestingCreateAccountCmd returns a CLI command handler for creating a
// MsgCreateVestingAccount transaction.
func GetTxVestingCreateAccountCmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vesting-account [to_address] [amount] [end_time]",
		Short: "Create a new vesting account funded with an allocation of tokens.",
		Long: `Create a new vesting account funded with an allocation of tokens. The
account can either be a delayed or continuous vesting account, which is determined
by the '--delayed' flag. All vesting accounts created will have their start time
set by the committed block's time. The end_time must be provided as a UNIX epoch
timestamp.`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			toAddr, err := ac.StringToBytes(args[0])
			if err != nil {
				return err
			}

			if args[1] == "" {
				return errors.New("amount is empty")
			}

			amount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			endTime, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			delayed, _ := cmd.Flags().GetBool(cflags.FlagDelayed)

			msg := types.NewMsgCreateVestingAccount(cctx.GetFromAddress(), toAddr, amount, endTime, delayed)

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cmd.Flags().Bool(cflags.FlagDelayed, false, "Create a delayed vesting account if true")
	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetTxVestingCreatePermanentLockedAccountCmd returns a CLI command handler for creating a
// MsgCreatePermanentLockedAccount transaction.
func GetTxVestingCreatePermanentLockedAccountCmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-permanent-locked-account [to_address] [amount]",
		Short: "Create a new permanently locked account funded with an allocation of tokens.",
		Long: `Create a new account funded with an allocation of permanently locked tokens. These
tokens may be used for staking but are non-transferable. Staking rewards will acrue as liquid and transferable
tokens.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			toAddr, err := ac.StringToBytes(args[0])
			if err != nil {
				return err
			}

			if args[1] == "" {
				return errors.New("amount is empty")
			}

			amount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePermanentLockedAccount(cctx.GetFromAddress(), toAddr, amount)

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

type VestingData struct {
	StartTime int64         `json:"start_time"`
	Periods   []InputPeriod `json:"periods"`
}

type InputPeriod struct {
	Coins  string `json:"coins"`
	Length int64  `json:"length_seconds"`
}

// GetTxVestingCreatePeriodicAccountCmd returns a CLI command handler for creating a
// MsgCreatePeriodicVestingAccountCmd transaction.
func GetTxVestingCreatePeriodicAccountCmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-periodic-vesting-account [to_address] [periods_json_file]",
		Short: "Create a new vesting account funded with an allocation of tokens.",
		Long: `A sequence of coins and period length in seconds. Periods are sequential, in that the duration of of a period only starts at the end of the previous period. The duration of the first period starts upon account creation. For instance, the following periods.json file shows 20 "test" coins vesting 30 days apart from each other.
		Where periods.json contains:

		An array of coin strings and unix epoch times for coins to vest
{ "start_time": 1625204910,
"periods":[
 {
  "coins": "10test",
  "length_seconds":2592000 //30 days
 },
 {
	"coins": "10test",
	"length_seconds":2592000 //30 days
 },
]
	}
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			toAddr, err := ac.StringToBytes(args[0])
			if err != nil {
				return err
			}

			contents, err := os.ReadFile(args[1])
			if err != nil {
				return err
			}

			var vestingData VestingData

			err = json.Unmarshal(contents, &vestingData)
			if err != nil {
				return err
			}

			var periods []types.Period

			for i, p := range vestingData.Periods {

				amount, err := sdk.ParseCoinsNormalized(p.Coins)
				if err != nil {
					return err
				}

				if p.Length < 0 {
					return fmt.Errorf("invalid period length of %d in period %d, length must be greater than 0", p.Length, i)
				}
				period := types.Period{Length: p.Length, Amount: amount}
				periods = append(periods, period)
			}

			msg := types.NewMsgCreatePeriodicVestingAccount(cctx.GetFromAddress(), toAddr, vestingData.StartTime, periods)

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
