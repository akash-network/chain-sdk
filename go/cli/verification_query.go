package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	cflags "pkg.akt.dev/go/cli/flags"
	types "pkg.akt.dev/go/node/verification/v1"
)

const verificationModuleName = "verification"

func GetQueryVerificationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        verificationModuleName,
		Short:                      "Verification query commands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		GetQueryVerificationParamsCmd(),
		GetQueryVerificationAuditorCmd(),
		GetQueryVerificationAuditorsCmd(),
		GetQueryVerificationAttestationCmd(),
		GetQueryVerificationProviderAttestationsCmd(),
		GetQueryVerificationAuditorAttestationsCmd(),
		GetQueryVerificationDiscrepancyCmd(),
		GetQueryVerificationDiscrepanciesCmd(),
		GetQueryVerificationAuditEscrowCmd(),
		GetQueryVerificationProviderAuditEscrowsCmd(),
		GetQueryVerificationProviderGraceCmd(),
		GetQueryVerificationProviderBondCmd(),
		GetQueryVerificationProviderSnapshotCmd(),
	)

	return cmd
}

func GetQueryVerificationParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "params",
		Short:             "Query verification params",
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			res, err := cl.Query().Verification().Params(ctx, &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetQueryVerificationAuditorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "auditor [auditor]",
		Short:             "Query auditor",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			auditor, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().Auditor(ctx, &types.QueryAuditorRequest{Auditor: auditor.String()})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetQueryVerificationAuditorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "auditors",
		Short:             "Query auditors",
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			pageReq, err := ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().Auditors(ctx, &types.QueryAuditorsRequest{Pagination: pageReq})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)
	cflags.AddPaginationFlagsToCmd(cmd, "auditors")

	return cmd
}

func GetQueryVerificationAttestationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "attestation [provider] [auditor]",
		Short:             "Query attestation",
		Args:              cobra.ExactArgs(2),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			provider, auditor, err := parseProviderAuditorArgs(args[0], args[1])
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().Attestation(ctx, &types.QueryAttestationRequest{
				Provider: provider.String(),
				Auditor:  auditor.String(),
			})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetQueryVerificationProviderAttestationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "provider-attestations [provider]",
		Short:             "Query provider attestations",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			provider, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			pageReq, err := ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().ProviderAttestations(ctx, &types.QueryProviderAttestationsRequest{
				Provider:   provider.String(),
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)
	cflags.AddPaginationFlagsToCmd(cmd, "attestations")

	return cmd
}

func GetQueryVerificationAuditorAttestationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "auditor-attestations [auditor]",
		Short:             "Query auditor attestations",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			auditor, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			pageReq, err := ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().AuditorAttestations(ctx, &types.QueryAuditorAttestationsRequest{
				Auditor:    auditor.String(),
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)
	cflags.AddPaginationFlagsToCmd(cmd, "attestations")

	return cmd
}

func GetQueryVerificationDiscrepancyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "discrepancy [id]",
		Short:             "Query discrepancy",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().Discrepancy(ctx, &types.QueryDiscrepancyRequest{Id: id})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetQueryVerificationDiscrepanciesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "discrepancies",
		Short:             "Query discrepancies",
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			pageReq, err := ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().Discrepancies(ctx, &types.QueryDiscrepanciesRequest{Pagination: pageReq})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)
	cflags.AddPaginationFlagsToCmd(cmd, "discrepancies")

	return cmd
}

func GetQueryVerificationAuditEscrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "audit-escrow [id]",
		Short:             "Query audit escrow",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().AuditEscrow(ctx, &types.QueryAuditEscrowRequest{Id: id})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetQueryVerificationProviderAuditEscrowsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "provider-audit-escrows [provider]",
		Short:             "Query provider audit escrows",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			provider, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			pageReq, err := ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().ProviderAuditEscrows(ctx, &types.QueryProviderAuditEscrowsRequest{
				Provider:   provider.String(),
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)
	cflags.AddPaginationFlagsToCmd(cmd, "audit escrows")

	return cmd
}

func GetQueryVerificationProviderGraceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "provider-grace [provider]",
		Short:             "Query provider verification grace",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			provider, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().ProviderVerificationGrace(ctx, &types.QueryProviderVerificationGraceRequest{Provider: provider.String()})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetQueryVerificationProviderBondCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "provider-bond [provider]",
		Short:             "Query provider bond",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			provider, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().ProviderBond(ctx, &types.QueryProviderBondRequest{Provider: provider.String()})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetQueryVerificationProviderSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "provider-snapshot [provider]",
		Short:             "Query provider snapshot",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			provider, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := cl.Query().Verification().ProviderSnapshot(ctx, &types.QueryProviderSnapshotRequest{Provider: provider.String()})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func parseProviderAuditorArgs(providerVal, auditorVal string) (sdk.AccAddress, sdk.AccAddress, error) {
	provider, err := sdk.AccAddressFromBech32(providerVal)
	if err != nil {
		return nil, nil, err
	}
	auditor, err := sdk.AccAddressFromBech32(auditorVal)
	if err != nil {
		return nil, nil, err
	}
	return provider, auditor, nil
}
