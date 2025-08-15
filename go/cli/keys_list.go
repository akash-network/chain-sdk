package cli

import (
	"github.com/spf13/cobra"

	cflags "pkg.akt.dev/go/cli/flags"
)

const flagListNames = "list-names"

// ListKeysCmd lists all keys in the key store.
func ListKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all keys",
		Long: `Return a list of all public keys stored by this key manager
along with their associated name and address.`,
		PreRunE: keysPreRunE,
		RunE:    runListCmd,
	}

	cmd.Flags().BoolP(flagListNames, "n", false, "List names only")
	return cmd
}

func runListCmd(cmd *cobra.Command, _ []string) error {
	clientCtx, err := GetClientQueryContext(cmd)
	if err != nil {
		return err
	}

	records, err := clientCtx.Keyring.List()
	if err != nil {
		return err
	}

	if len(records) == 0 && clientCtx.OutputFormat == cflags.OutputFormatText {
		cmd.Println("No records were found in keyring")
		return nil
	}

	if ok, _ := cmd.Flags().GetBool(flagListNames); !ok {
		return printKeyringRecords(cmd.OutOrStdout(), records, clientCtx.OutputFormat)
	}

	for _, k := range records {
		cmd.Println(k.Name)
	}

	return nil
}

// ListKeyTypesCmd lists all key types.
func ListKeyTypesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list-key-types",
		Short:   "List all key types",
		Long:    `Return a list of all supported key types (also known as algos)`,
		PreRunE: keysPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			cmd.Println("Supported key types/algos:")
			keyring, _ := clientCtx.Keyring.SupportedAlgorithms()
			cmd.Printf("%+q\n", keyring)
			return nil
		},
	}
}
