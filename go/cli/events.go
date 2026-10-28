package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/cosmos/cosmos-sdk/client"

	cflags "pkg.akt.dev/go/cli/flags"

	"pkg.akt.dev/go/util/events"
	"pkg.akt.dev/go/util/pubsub"
)

// EventsCmd prints out events in real time
func EventsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "events",
		Short: "Prints out akash events in real time",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunForeverWithContext(cmd.Context(), func(ctx context.Context) error {
				return getEvents(ctx, cmd, args)
			})
		},
	}

	cmd.Flags().String(cflags.FlagNode, "tcp://localhost:26657", "The node address")
	if err := viper.BindPFlag(cflags.FlagNode, cmd.Flags().Lookup(cflags.FlagNode)); err != nil {
		panic(fmt.Sprintf("failed to bind flag %s: %v", cflags.FlagNode, err))
	}

	return cmd
}

func getEvents(ctx context.Context, cmd *cobra.Command, _ []string) error {
	cctx := client.GetClientContextFromCmd(cmd)

	// The events service dials the node itself, so it needs the node's address rather
	// than the client the context holds.
	nodeURI := cctx.NodeURI
	if nodeURI == "" {
		nodeURI, _ = cmd.Flags().GetString(cflags.FlagNode)
	}
	if nodeURI == "" {
		return errors.New("events: no node URI")
	}

	return streamEvents(ctx, cctx, nodeURI)
}

// streamEvents prints the chain events the node at nodeURI emits until ctx ends, or
// until the events service stops on its own, whose error it returns.
func streamEvents(ctx context.Context, cctx client.Context, nodeURI string, opts ...events.Option) error {
	bus := pubsub.NewBus()
	defer bus.Close()

	group, ctx := errgroup.WithContext(ctx)

	subscriber, err := bus.Subscribe()
	if err != nil {
		return err
	}

	evtSvc, err := events.NewEvents(ctx, nodeURI, "akash-cli", bus, opts...)
	if err != nil {
		return err
	}

	group.Go(func() error {
		return waitEvents(ctx, evtSvc)
	})

	group.Go(func() error {
		return printEvents(ctx, cctx, subscriber)
	})

	err = group.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

// waitEvents returns once ctx ends, after shutting the events service down, or once the
// service stops on its own, with the error it stopped on.
func waitEvents(ctx context.Context, svc events.Service) error {
	select {
	case <-ctx.Done():
		svc.Shutdown()
		return nil
	case <-svc.Done():
		return svc.Err()
	}
}

// printEvents prints every event sub receives until ctx ends or sub is closed.
func printEvents(ctx context.Context, cctx client.Context, sub pubsub.Subscriber) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-sub.Done():
			return nil
		case ev := <-sub.Events():
			if err := PrintJSON(cctx, ev); err != nil {
				return err
			}
		}
	}
}
