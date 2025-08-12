package cli

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"cosmossdk.io/core/address"
	"github.com/spf13/cobra"

	aclient "pkg.akt.dev/go/node/client/discovery"
	"pkg.akt.dev/go/node/client/v1beta3"

	cflags "pkg.akt.dev/go/cli/flags"
)

type ContextType string

const (
	ContextTypeClient         = ContextType("context-client")
	ContextTypeQueryClient    = ContextType("context-query-client")
	ContextTypeAddressCodec   = ContextType("address-codec")
	ContextTypeValidatorCodec = ContextType("validator-codec")
)

func TxPersistentPreRunE(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()

	cctx, err := GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	if cctx.Codec == nil {
		return errors.New("codec is not initialized")
	}

	if cctx.LegacyAmino == nil {
		return errors.New("legacy amino codec is not initialized")
	}

	opts, err := cflags.ClientOptionsFromFlags(cmd.Flags())
	if err != nil {
		return err
	}

	cl, err := aclient.DiscoverClient(ctx, cctx, opts...)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, ContextTypeClient, cl)

	cmd.SetContext(ctx)

	return nil
}

func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}

	cmd.AddCommand(
		GetTxAuthzCmd(),
		GetTxBankCmd(),
		GetTxCrisisCmd(),
		getTxDistributionCmd(),
		GetTxEvidenceCmd([]*cobra.Command{}),
		GetTxFeegrantCmd(),
		GetSignCommand(),
		GetSignBatchCommand(),
		GetAuthMultiSignCmd(),
		GetValidateSignaturesCommand(),
		GetBroadcastCommand(),
		GetEncodeCommand(),
		GetDecodeCommand(),
		GetTxVestingCmd(),
		cflags.LineBreak,
		GetTxAuditCmd(),
		GetTxCertCmd(),
		GetTxDeploymentCmds(),
		GetTxMarketCmds(),
		GetTxProviderCmd(),
		GetTxGovCmd(
			[]*cobra.Command{
				GetTxParamsSubmitParamChangeProposalCmd(),
			},
		),
		GetTxSlashingCmd(),
		GetTxStakingCmd(),
		GetTxUpgradeCmd(),
	)

	cmd.PersistentFlags().String(cflags.FlagChainID, "", "The network chain ID")

	return cmd
}

func MustClientFromContext(ctx context.Context) v1beta3.Client {
	val := ctx.Value(ContextTypeClient)
	if val == nil {
		panic("context does not have client set")
	}

	res, valid := val.(v1beta3.Client)
	if !valid {
		panic("invalid context value")
	}

	return res
}

func MustLightClientFromContext(ctx context.Context) v1beta3.LightClient {
	val := ctx.Value(ContextTypeQueryClient)
	if val == nil {
		panic("context does not have client set")
	}

	switch cl := val.(type) {
	case v1beta3.LightClient:
		return cl
	case v1beta3.Client:
		return cl
	default:
		panic(fmt.Sprintf("invalid context value. actual %s", reflect.TypeOf(val).String()))
	}
}

func MustAddressCodecFromContext(ctx context.Context) address.Codec {
	val := ctx.Value(ContextTypeAddressCodec)
	if val == nil {
		panic("context does not have address codec set")
	}

	res, valid := val.(address.Codec)
	if !valid {
		panic("invalid context value")
	}

	return res
}

func MustValidatorCodecFromContext(ctx context.Context) address.Codec {
	val := ctx.Value(ContextTypeValidatorCodec)
	if val == nil {
		panic("context does not have validator codec set")
	}

	res, valid := val.(address.Codec)
	if !valid {
		panic("invalid context value")
	}

	return res
}
