package cli

import (
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/client"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/spf13/cobra"

	cflags "pkg.akt.dev/go/cli/flags"
)

// GetEncodeCommand returns the encode command to take a JSONified transaction and turn it into
// Amino-serialized bytes
func GetEncodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encode [file]",
		Short: "Encode transactions generated offline",
		Long: `Encode transactions created with the --generate-only flag or signed with the sign command.
Read a transaction from <file>, serialize it to the Protobuf wire protocol, and output it as base64.
If you supply a dash (-) argument in place of an input filename, the command reads from standard input.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			tx, err := authclient.ReadTxFromFile(clientCtx, args[0])
			if err != nil {
				return err
			}

			// re-encode it
			txBytes, err := clientCtx.TxConfig.TxEncoder()(tx)
			if err != nil {
				return err
			}

			// base64 encode the encoded tx bytes
			txBytesBase64 := base64.StdEncoding.EncodeToString(txBytes)

			return clientCtx.PrintString(txBytesBase64 + "\n")
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	_ = cmd.Flags().MarkHidden(cflags.FlagOutput) // encoding makes sense to output only json

	return cmd
}
