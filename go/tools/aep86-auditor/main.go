package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	if err := newRootCmd().ExecuteContext(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "aep86-auditor",
		Short:        "AEP-86 reference auditor tooling",
		SilenceUsage: true,
	}

	cmd.AddCommand(newCollectCmd(), newEvidenceCmd(), newVerifyCmd())

	return cmd
}

func timeoutContext(cmd *cobra.Command, timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx := cmd.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithTimeout(ctx, timeout)
}
