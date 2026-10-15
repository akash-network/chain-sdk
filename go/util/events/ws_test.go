package events

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	cmtlog "github.com/cometbft/cometbft/libs/log"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	rpcserver "github.com/cometbft/cometbft/rpc/jsonrpc/server"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	cmtypes "github.com/cometbft/cometbft/types"

	mtypes "pkg.akt.dev/go/node/market/v1"
	"pkg.akt.dev/go/util/pubsub"
)

// trackingListener records every accepted connection so a test can drop them all
// abruptly, the way a load balancer or a NAT timeout does, without a websocket close
// frame.
type trackingListener struct {
	net.Listener

	mu    sync.Mutex
	conns []net.Conn
}

func (l *trackingListener) Accept() (net.Conn, error) {
	c, err := l.Listener.Accept()
	if err == nil {
		l.mu.Lock()
		l.conns = append(l.conns, c)
		l.mu.Unlock()
	}
	return c, err
}

// dropAll closes every accepted connection. With reset set it sends a TCP RST
// instead of a FIN, which the client reads as a connection reset rather than EOF.
func (l *trackingListener) dropAll(reset bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, c := range l.conns {
		if tc, ok := c.(*net.TCPConn); ok && reset {
			_ = tc.SetLinger(0)
		}
		_ = c.Close()
	}
	l.conns = nil
}

type wsSub struct {
	conn  rpctypes.WSRPCConnection
	req   *rpctypes.RPCRequest
	query string
}

// rpcNode is a minimal CometBFT RPC server: status and block_results over HTTP, and
// subscribe over the websocket, driven by the real cometbft jsonrpc server so the
// client under test is the real rpchttp client.
type rpcNode struct {
	t  *testing.T
	ln *trackingListener

	mu         sync.Mutex
	height     int64
	subs       []wsSub
	subscribes int
	// frozen, when non-zero, is the height status reports regardless of the chain, so
	// the watchdog's catch-up cannot reach anything above it.
	frozen int64

	// lagEvery, when non-zero, makes block_results for every height divisible by it
	// fail once the way a pooled backend one block behind answers, and lagged records
	// which heights did.
	lagEvery int64
	lagged   map[int64]bool
	// rejectPosts is how many of the next HTTP POSTs the proxy refuses with 403, as a
	// rate limiter in front of the backends does; rejected counts the refusals.
	rejectPosts int
	rejected    int
}

func newRPCNode(t *testing.T, height int64) *rpcNode {
	t.Helper()

	n := &rpcNode{t: t, height: height, lagged: map[int64]bool{}}

	funcs := map[string]*rpcserver.RPCFunc{
		"status":          rpcserver.NewRPCFunc(n.status, ""),
		"block_results":   rpcserver.NewRPCFunc(n.blockResults, "height"),
		"subscribe":       rpcserver.NewWSRPCFunc(n.subscribe, "query"),
		"unsubscribe":     rpcserver.NewWSRPCFunc(n.unsubscribe, "query"),
		"unsubscribe_all": rpcserver.NewWSRPCFunc(n.unsubscribeAll, ""),
	}

	mux := http.NewServeMux()
	rpcserver.RegisterRPCFuncs(mux, funcs, cmtlog.NewNopLogger())
	wm := rpcserver.NewWebsocketManager(funcs)
	wm.SetLogger(cmtlog.NewNopLogger())
	mux.HandleFunc("/websocket", wm.WebsocketHandler)

	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	n.ln = &trackingListener{Listener: l}

	srv := &http.Server{Handler: n.proxy(mux), ReadHeaderTimeout: time.Second}
	go func() { _ = srv.Serve(n.ln) }()
	t.Cleanup(func() { _ = srv.Close() })

	return n
}

func (n *rpcNode) addr() string { return "tcp://" + n.ln.Addr().String() }

func (n *rpcNode) status(*rpctypes.Context) (*coretypes.ResultStatus, error) {
	n.mu.Lock()
	defer n.mu.Unlock()
	latest := n.height
	if n.frozen != 0 {
		latest = n.frozen
	}
	return &coretypes.ResultStatus{SyncInfo: coretypes.SyncInfo{LatestBlockHeight: latest, EarliestBlockHeight: 1}}, nil
}

// freezeStatus pins the height status reports to the current one and returns it.
func (n *rpcNode) freezeStatus() int64 {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.frozen = n.height
	return n.frozen
}

func (n *rpcNode) blockResults(_ *rpctypes.Context, height *int64) (*coretypes.ResultBlockResults, error) {
	n.mu.Lock()
	defer n.mu.Unlock()
	if h := *height; n.lagEvery > 0 && h%n.lagEvery == 0 && !n.lagged[h] {
		n.lagged[h] = true
		return nil, fmt.Errorf("height %d must be less than or equal to the current blockchain height %d", h, h-1)
	}
	return txBlock(*height, leaseClosedEvent(n.t, dseqAt(*height))), nil
}

// proxy stands in front of the RPC server the way a rate-limiting load balancer does,
// refusing POSTs with 403 while rejectPosts lasts. The websocket upgrade is a GET and
// passes.
func (n *rpcNode) proxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			n.mu.Lock()
			reject := n.rejectPosts > 0
			if reject {
				n.rejectPosts--
				n.rejected++
			}
			n.mu.Unlock()
			if reject {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (n *rpcNode) set(mutate func(*rpcNode)) {
	n.mu.Lock()
	defer n.mu.Unlock()
	mutate(n)
}

func (n *rpcNode) get(read func(*rpcNode) int) int {
	n.mu.Lock()
	defer n.mu.Unlock()
	return read(n)
}

func (n *rpcNode) subscribe(ctx *rpctypes.Context, query string) (*coretypes.ResultSubscribe, error) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.subscribes++
	n.subs = append(n.subs, wsSub{conn: ctx.WSConn, req: ctx.JSONReq, query: query})
	return &coretypes.ResultSubscribe{}, nil
}

func (n *rpcNode) unsubscribe(*rpctypes.Context, string) (*coretypes.ResultUnsubscribe, error) {
	return &coretypes.ResultUnsubscribe{}, nil
}

func (n *rpcNode) unsubscribeAll(*rpctypes.Context) (*coretypes.ResultUnsubscribe, error) {
	return &coretypes.ResultUnsubscribe{}, nil
}

func (n *rpcNode) subscribeCount() int {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.subscribes
}

// commit advances the chain one block and pushes its header to every live subscriber.
func (n *rpcNode) commit() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.height++
	for _, s := range n.subs {
		_ = s.conn.TryWriteRPCResponse(rpctypes.NewRPCSuccessResponse(s.req.ID, coretypes.ResultEvent{
			Query: s.query,
			Data:  cmtypes.EventDataNewBlockHeader{Header: cmtypes.Header{Height: n.height}},
		}))
	}
}

// drop severs every connection without a close frame and forgets the subscriptions
// carried on them, as the node does when its side of the socket goes away.
func (n *rpcNode) drop(reset bool) {
	n.mu.Lock()
	n.subs = nil
	n.mu.Unlock()
	n.ln.dropAll(reset)
}

// The websocket can drop at any moment and in more than one way. A FIN reads as an
// abnormal close and the cometbft client redials on its own. A reset (or a keepalive
// timeout) reads as a plain network error, and the client's read loop then exits
// without scheduling a reconnect: the subscription channel stays open and goes silent
// forever. Either way the service must notice through the watchdog, get the
// subscription re-established, and publish every block in between, in order.
func TestEventsRecoversWebsocketDrop(t *testing.T) {
	if testing.Short() {
		t.Skip("drives the real cometbft websocket client, which backs off at least a second before redialing")
	}

	for _, tc := range []struct {
		name  string
		reset bool
	}{
		{name: "fin", reset: false},
		{name: "reset", reset: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testWebsocketDrop(t, tc.reset)
		})
	}
}

// startRPCService runs the service against node through the real rpchttp client and
// waits for its block subscription to reach the node.
func startRPCService(t *testing.T, node *rpcNode, tune tuning) pubsub.Subscriber {
	t.Helper()

	cl, err := rpchttp.New(node.addr(), "/websocket")
	require.NoError(t, err)
	require.NoError(t, cl.Start())
	t.Cleanup(func() { _ = cl.Stop() })

	bus := pubsub.NewBus()
	sub, err := bus.Subscribe()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	svc, err := NewEvents(ctx, cl, "test", bus, tune.options()...)
	require.NoError(t, err)
	t.Cleanup(func() {
		svc.Shutdown()
		bus.Close()
	})

	require.Eventually(t, func() bool { return node.subscribeCount() == 1 }, 2*time.Second, 10*time.Millisecond)

	return sub
}

// produceBlocks commits a block every interval until the test ends.
func produceBlocks(t *testing.T, node *rpcNode, interval time.Duration) {
	t.Helper()

	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		tick := time.NewTicker(interval)
		defer tick.Stop()
		for {
			select {
			case <-stop:
				return
			case <-tick.C:
				node.commit()
			}
		}
	}()
	t.Cleanup(func() {
		close(stop)
		wg.Wait()
	})
}

// requireInOrderThrough reads lease-closed events and requires DSeqs from, from+1, ...
// through last with none missing and none repeated.
func requireInOrderThrough(t *testing.T, sub pubsub.Subscriber, from, last uint64) {
	t.Helper()

	want := from
	deadline := time.After(5 * time.Second)
	for want <= last {
		select {
		case ev := <-sub.Events():
			closed, ok := ev.(*mtypes.EventLeaseClosed)
			if !ok {
				continue
			}
			require.Equal(t, want, closed.ID.DSeq, "events must be published in block order with no gaps or repeats")
			want++
		case <-deadline:
			t.Fatalf("stopped receiving events at height %d", want)
		}
	}
}

// The incident behind rpc.akt.dev: the websocket and the per-block POSTs reach
// different backends, so block_results for a height a header just announced can land
// on a backend one block behind, or be refused with 403 by the proxy's rate limiter.
// The service used to drop such a height silently. Each failure must delay the height
// by a retry, never skip it, and the watchdog is far too long to be what recovers it.
func TestEventsSurvivesLaggingAndRateLimitedBackends(t *testing.T) {
	if testing.Short() {
		t.Skip("drives the real cometbft client against a local RPC server")
	}

	node := newRPCNode(t, 100)
	node.set(func(n *rpcNode) { n.lagEvery = 3 })

	sub := startRPCService(t, node, tuning{
		watchdog:         time.Hour,
		minRetryInterval: 20 * time.Millisecond,
		prunedCheckAfter: time.Hour,
	})
	produceBlocks(t, node, 50*time.Millisecond)

	requireInOrderThrough(t, sub, 101, 115)

	node.set(func(n *rpcNode) { n.rejectPosts = 5 })
	requireInOrderThrough(t, sub, 116, 130)

	require.Positive(t, node.get(func(n *rpcNode) int { return len(n.lagged) }), "no fetch landed on a lagging backend")
	require.Equal(t, 5, node.get(func(n *rpcNode) int { return n.rejected }), "the rate limiter never refused a fetch")
}

func testWebsocketDrop(t *testing.T, reset bool) {
	node := newRPCNode(t, 100)

	sub := startRPCService(t, node, tuning{
		watchdog:         300 * time.Millisecond,
		minRetryInterval: 50 * time.Millisecond,
		prunedCheckAfter: time.Hour,
	})
	produceBlocks(t, node, 50*time.Millisecond)

	time.Sleep(200 * time.Millisecond)
	node.drop(reset)

	require.Eventually(t, func() bool { return node.subscribeCount() >= 2 }, 5*time.Second, 10*time.Millisecond,
		"the service never re-established the block subscription after the socket dropped")

	// From here on the watchdog's catch-up cannot get past frozen, so every height
	// above it must arrive as a header on the re-established subscription.
	frozen := node.freezeStatus()

	// Every height from the start must have been published exactly once, in order.
	requireInOrderThrough(t, sub, 101, dseqAt(frozen)+10)
}
