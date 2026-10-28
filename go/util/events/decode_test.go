package events

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/cometbft/cometbft/abci/types"

	mtypes "pkg.akt.dev/go/node/market/v1"
)

// A node that answers with another height's results must not have them published as
// the requested height.
func TestBlockEventsRejectsMismatchedHeight(t *testing.T) {
	node := newFakeNode()
	node.set(func(f *fakeNode) { f.blocks[101] = txBlock(100, leaseClosedEvent(t, 100)) })
	opts, err := newOptions()
	require.NoError(t, err)
	e := &events{client: node, log: newRecLogger(), opts: opts}

	_, err = e.blockEvents(context.Background(), 101)
	require.ErrorContains(t, err, "height 100")
	require.ErrorContains(t, err, "101")
}

// A recognized block-level event that cannot be decoded is left out of its height; the
// height's other events are still returned.
func TestBlockEventsSkipsUndecodableBlockLevelEvent(t *testing.T) {
	node := newFakeNode()

	bad := withAttrs(abci.Event{Type: leaseClosedEvent(t, 2).Type, Attributes: []abci.EventAttribute{{Key: "id", Value: undecodableAttr}}},
		abci.EventAttribute{Key: modeKey, Value: modeEndBlock})
	blk := txBlock(101, leaseClosedEvent(t, 1))
	blk.FinalizeBlockEvents = []abci.Event{bad}
	node.set(func(f *fakeNode) { f.blocks[101] = blk })

	opts, err := newOptions()
	require.NoError(t, err)
	rec := newRecLogger()
	e := &events{client: node, log: rec, opts: opts}

	evs, err := e.blockEvents(context.Background(), 101)
	require.NoError(t, err, "an undecodable event is not a failure of its height")
	require.Len(t, evs, 1)
	closed, ok := evs[0].(*mtypes.EventLeaseClosed)
	require.True(t, ok)
	require.Equal(t, uint64(1), closed.ID.DSeq)
	require.Equal(t, 1, rec.count(levelError, "skipping undecodable event"))
}

func TestProcessEventFiltersUnrecognized(t *testing.T) {
	_, ok, err := processEvent(leaseClosedEvent(t, 1))
	require.NoError(t, err)
	require.True(t, ok, "a lease-closed event must be recognized")

	other := abci.Event{Type: "coin_received", Attributes: []abci.EventAttribute{{Key: "amount", Value: "1uakt"}}}
	_, ok, err = processEvent(other)
	require.NoError(t, err, "an event the service does not publish is not a decode failure")
	require.False(t, ok, "a non-Akash event must be ignored")
}

// The SDK appends bookkeeping attributes to the events it records: msg_index to every
// message event, authz_msg_index to events of messages run through authz MsgExec, and
// mode to every BeginBlock and EndBlock event. None of them may stop a recognized
// event from being decoded. mode in particular is a bare word, not JSON.
func TestProcessEventAcceptsSDKBookkeepingAttributes(t *testing.T) {
	for _, tc := range []struct {
		name  string
		attrs []abci.EventAttribute
	}{
		{name: "tx", attrs: []abci.EventAttribute{{Key: "msg_index", Value: "0"}}},
		{name: "authz", attrs: []abci.EventAttribute{{Key: "authz_msg_index", Value: "1"}, {Key: "msg_index", Value: "0"}}},
		{name: "begin block", attrs: []abci.EventAttribute{{Key: modeKey, Value: modeBeginBlock}}},
		{name: "end block", attrs: []abci.EventAttribute{{Key: modeKey, Value: modeEndBlock}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ev := withAttrs(leaseClosedEvent(t, 7), tc.attrs...)

			pev, ok, err := processEvent(ev)
			require.NoError(t, err)
			require.True(t, ok)
			closed, isClosed := pev.(*mtypes.EventLeaseClosed)
			require.True(t, isClosed)
			require.Equal(t, uint64(7), closed.ID.DSeq)
		})
	}
}

// A recognized event that cannot be decoded is reported, never treated as an event
// the service does not publish.
func TestProcessEventReportsUndecodableRecognizedEvent(t *testing.T) {
	ev := leaseClosedEvent(t, 1)
	ev.Attributes = []abci.EventAttribute{{Key: "id", Value: undecodableAttr}}

	_, ok, err := processEvent(ev)
	require.Error(t, err)
	require.False(t, ok)
}
