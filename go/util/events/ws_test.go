package events

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"

	cmtlog "github.com/cometbft/cometbft/libs/log"
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
// clients under test speak to the real thing.
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
	// subscribeErr, when set, is what subscribe answers, the way a node at its
	// subscription limits does. The call is still counted.
	subscribeErr error

	// handler serves the node's RPC, so a test can put it behind another listener.
	handler http.Handler
	// conns holds the HTTP connections the node serves that are still open. A
	// websocket leaves it once upgraded.
	conns map[net.Conn]http.ConnState
}

// trackConn records the state of every HTTP connection the node serves.
func (n *rpcNode) trackConn(c net.Conn, s http.ConnState) {
	n.mu.Lock()
	defer n.mu.Unlock()
	switch s {
	case http.StateClosed, http.StateHijacked:
		delete(n.conns, c)
	default:
		n.conns[c] = s
	}
}

// openHTTPConns is how many HTTP connections clients still hold open to the node.
func (n *rpcNode) openHTTPConns() int {
	n.mu.Lock()
	defer n.mu.Unlock()
	return len(n.conns)
}

// newRPCNode starts a node whose chain is at height 100.
func newRPCNode(t *testing.T) *rpcNode {
	t.Helper()

	n := &rpcNode{t: t, height: 100, lagged: map[int64]bool{}, conns: map[net.Conn]http.ConnState{}}

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
	n.handler = n.proxy(mux)

	srv := &http.Server{Handler: n.handler, ReadHeaderTimeout: time.Second, ConnState: n.trackConn}
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
	if n.subscribeErr != nil {
		return nil, n.subscribeErr
	}
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

// push sends a header for height h to every live subscriber without moving the chain,
// the way a lagging backend repeats one already delivered.
func (n *rpcNode) push(h int64) {
	n.mu.Lock()
	defer n.mu.Unlock()
	for _, s := range n.subs {
		_ = s.conn.TryWriteRPCResponse(rpctypes.NewRPCSuccessResponse(s.req.ID, coretypes.ResultEvent{
			Query: s.query,
			Data:  cmtypes.EventDataNewBlockHeader{Header: cmtypes.Header{Height: h}},
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

// The websocket can drop at any moment and in more than one way: a FIN or a reset.
// Either way the service must notice, redial, get the subscription re-established, and
// publish every block in between, in order.
func TestEventsRecoversWebsocketDrop(t *testing.T) {
	if testing.Short() {
		t.Skip("drives a real websocket against a local RPC server")
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

// startRPCService runs the service against node, over the connection it dials itself,
// and waits for its block subscription to reach the node.
func startRPCService(t *testing.T, node *rpcNode, tune tuning, extra ...Option) pubsub.Subscriber {
	t.Helper()
	return startRPCServiceAt(t, node, node.addr(), tune, extra...)
}

// startRPCServiceAt is startRPCService reaching node through addr, such as a proxy.
func startRPCServiceAt(t *testing.T, node *rpcNode, addr string, tune tuning, extra ...Option) pubsub.Subscriber {
	t.Helper()

	bus := pubsub.NewBus()
	sub, err := bus.Subscribe()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	svc, err := NewEvents(ctx, addr, "test", bus, append(tune.options(), extra...)...)
	require.NoError(t, err)
	t.Cleanup(func() {
		svc.Shutdown()
		bus.Close()
	})

	require.Eventually(t, func() bool { return node.subscribeCount() >= 1 }, 2*time.Second, 10*time.Millisecond)

	return sub
}

// blockInterval is how often produceBlocks commits a block.
const blockInterval = 50 * time.Millisecond

// produceBlocks commits a block every blockInterval until the test ends.
func produceBlocks(t *testing.T, node *rpcNode) {
	t.Helper()

	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		tick := time.NewTicker(blockInterval)
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
		t.Skip("drives a real websocket against a local RPC server")
	}

	node := newRPCNode(t)
	node.set(func(n *rpcNode) { n.lagEvery = 3 })

	sub := startRPCService(t, node, tuning{
		watchdog:         time.Hour,
		retryInterval:    20 * time.Millisecond,
		prunedCheckAfter: time.Hour,
	})
	produceBlocks(t, node)

	requireInOrderThrough(t, sub, 101, 115)

	node.set(func(n *rpcNode) { n.rejectPosts = 5 })
	requireInOrderThrough(t, sub, 116, 130)

	require.Positive(t, node.get(func(n *rpcNode) int { return len(n.lagged) }), "no fetch landed on a lagging backend")
	require.Equal(t, 5, node.get(func(n *rpcNode) int { return n.rejected }), "the rate limiter never refused a fetch")
}

// A websocket whose path goes dead without either end closing it, as behind a NAT that
// dropped its mapping or a proxy that keeps the client side open, never fails a read or
// a write on its own. The service must notice through the missing pongs, redial, and
// carry on. Once status is frozen, only headers on a redialed subscription can carry
// the service past the frozen height.
func TestEventsRecoversHalfOpenWebsocket(t *testing.T) {
	if testing.Short() {
		t.Skip("drives a real websocket against a local RPC server")
	}

	node := newRPCNode(t)
	proxy := newBlackholeProxy(t, node.ln.Addr().String())

	sub := startRPCServiceAt(t, node, proxy.addr(), tuning{
		watchdog:         300 * time.Millisecond,
		retryInterval:    50 * time.Millisecond,
		prunedCheckAfter: time.Hour,
	}, withPing(20*time.Millisecond, 100*time.Millisecond))
	produceBlocks(t, node)

	requireInOrderThrough(t, sub, 101, 105)

	proxy.blackholeWebsocket()
	frozen := node.freezeStatus()

	requireInOrderThrough(t, sub, 106, dseqAt(frozen)+10)
	require.GreaterOrEqual(t, node.subscribeCount(), 2)
}

// A node that rejects the subscription, as one at its subscription limits does, is
// asked again once per retry interval on a fresh connection, never in a storm, while
// events keep flowing through the watchdog's polling.
func TestEventsPacesRedialsWhenSubscribeRejected(t *testing.T) {
	if testing.Short() {
		t.Skip("drives a real websocket against a local RPC server")
	}

	node := newRPCNode(t)
	node.set(func(n *rpcNode) { n.subscribeErr = errors.New("max_subscription_clients 100 reached") })

	sub := startRPCService(t, node, tuning{
		watchdog:         100 * time.Millisecond,
		retryInterval:    100 * time.Millisecond,
		prunedCheckAfter: time.Hour,
	})
	produceBlocks(t, node)

	before := node.subscribeCount()
	time.Sleep(time.Second)
	subscribes := node.subscribeCount() - before
	require.GreaterOrEqual(t, subscribes, 5, "the subscription keeps being asked for")
	require.LessOrEqual(t, subscribes, 12, "about once per retry interval, not a storm")

	requireInOrderThrough(t, sub, 101, 110)
}

func testWebsocketDrop(t *testing.T, reset bool) {
	node := newRPCNode(t)

	sub := startRPCService(t, node, tuning{
		watchdog:         300 * time.Millisecond,
		retryInterval:    50 * time.Millisecond,
		prunedCheckAfter: time.Hour,
	})
	produceBlocks(t, node)

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

// proxiedConn is one connection through a blackholeProxy.
type proxiedConn struct {
	client, upstream net.Conn
	ws, frozen       atomic.Bool
}

// blackholeProxy forwards TCP connections to a target until told to swallow the
// websocket ones: then both directions are read and dropped and neither side is
// closed, the way a dead path behind a NAT, or a proxy that keeps the client side
// open, looks to the client.
type blackholeProxy struct {
	ln    net.Listener
	mu    sync.Mutex
	conns []*proxiedConn
}

func newBlackholeProxy(t *testing.T, target string) *blackholeProxy {
	t.Helper()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	p := &blackholeProxy{ln: ln}

	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			u, err := net.Dial("tcp", target)
			if err != nil {
				_ = c.Close()
				continue
			}
			pc := &proxiedConn{client: c, upstream: u}
			p.mu.Lock()
			p.conns = append(p.conns, pc)
			p.mu.Unlock()
			go pc.pipe(c, u, true)
			go pc.pipe(u, c, false)
		}
	}()

	t.Cleanup(func() {
		_ = ln.Close()
		p.mu.Lock()
		defer p.mu.Unlock()
		for _, pc := range p.conns {
			_ = pc.client.Close()
			_ = pc.upstream.Close()
		}
	})

	return p
}

func (p *blackholeProxy) addr() string { return "tcp://" + p.ln.Addr().String() }

func (pc *proxiedConn) pipe(src, dst net.Conn, fromClient bool) {
	buf := make([]byte, 32<<10)
	first := true
	for {
		n, err := src.Read(buf)
		if n > 0 {
			if first && fromClient && bytes.HasPrefix(buf[:n], []byte("GET /websocket")) {
				pc.ws.Store(true)
			}
			first = false
			if !pc.frozen.Load() {
				if _, werr := dst.Write(buf[:n]); werr != nil {
					return
				}
			}
		}
		if err != nil {
			if !pc.frozen.Load() {
				_ = dst.Close()
			}
			return
		}
	}
}

// blackholeWebsocket swallows every websocket connection open so far.
func (p *blackholeProxy) blackholeWebsocket() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, pc := range p.conns {
		if pc.ws.Load() {
			pc.frozen.Store(true)
		}
	}
}

// streamConfig aims a header stream at nodeURI with the given liveness timings.
func streamConfig(t *testing.T, nodeURI string, ping, pong time.Duration) wsConfig {
	t.Helper()

	u, err := parseNodeURI(nodeURI)
	require.NoError(t, err)

	return wsConfig{
		target:     newWSTarget(u),
		query:      blkHeaderQuery().String(),
		timeout:    2 * time.Second,
		pingPeriod: ping,
		pongWait:   pong,
	}
}

// dialStream opens a header stream and closes it with the test.
func dialStream(t *testing.T, cfg wsConfig) headerStream {
	t.Helper()

	s, err := dialHeaders(context.Background(), cfg)
	require.NoError(t, err)
	t.Cleanup(s.Close)

	return s
}

// requireOpen fails if s has ended.
func requireOpen(t *testing.T, s headerStream) {
	t.Helper()

	select {
	case <-s.Done():
		t.Fatalf("stream ended: %v", s.Err())
	default:
	}
}

// requireEnded waits for s to end and returns why.
func requireEnded(t *testing.T, s headerStream, within time.Duration) error {
	t.Helper()

	select {
	case <-s.Done():
		return s.Err()
	case <-time.After(within):
		t.Fatalf("stream still open after %s", within)
		return nil
	}
}

// waitLatest waits until the highest header s has seen is h.
func waitLatest(t *testing.T, s headerStream, h int64) {
	t.Helper()

	deadline := time.After(2 * time.Second)
	for s.Latest() < h {
		select {
		case <-s.Notify():
		case <-deadline:
			t.Fatalf("stream latest %d, want %d", s.Latest(), h)
		}
	}
	require.Equal(t, h, s.Latest())
}

// unixSocketURI is a node reached over a unix socket.
const unixSocketURI = "unix:///tmp/cmt.sock"

func TestWSTargetMapsNodeURIs(t *testing.T) {
	const (
		exampleAddr = "rpc.example.com:443"
		exampleURL  = "wss://" + exampleAddr + "/websocket"
	)

	for _, tc := range []struct {
		uri, url, network, address, auth string
	}{
		{uri: "tcp://127.0.0.1:26657", url: "ws://127.0.0.1:26657/websocket", network: networkTCP, address: "127.0.0.1:26657"},
		{uri: "http://node:26657", url: "ws://node:26657/websocket", network: networkTCP, address: "node:26657"},
		{uri: "//node:26657", url: "ws://node:26657/websocket", network: networkTCP, address: "node:26657"},
		{uri: "//node", url: "ws://node:80/websocket", network: networkTCP, address: "node:80"},
		{uri: "https://TOKEN@rpc.example.com", url: exampleURL, network: networkTCP, address: exampleAddr, auth: "Basic VE9LRU46"},
		{uri: "https://rpc.example.com", url: exampleURL, network: networkTCP, address: exampleAddr},
		{uri: "wss://rpc.example.com:443/", url: exampleURL, network: networkTCP, address: exampleAddr},
		{ //nolint:gosec // a test fixture's credentials, not a secret
			uri: "https://user:pass@rpc.example.com/rpc", url: "wss://rpc.example.com:443/rpc/websocket",
			network: networkTCP, address: exampleAddr, auth: "Basic dXNlcjpwYXNz",
		},
		{uri: unixSocketURI, url: "ws://.tmp.cmt.sock/websocket", network: networkUnix, address: "/tmp/cmt.sock"},
	} {
		t.Run(tc.uri, func(t *testing.T) {
			u, err := parseNodeURI(tc.uri)
			require.NoError(t, err)

			target := newWSTarget(u)
			require.Equal(t, tc.url, target.url)
			require.Equal(t, tc.network, target.network)
			require.Equal(t, tc.address, target.address)
			require.Equal(t, tc.auth, target.header.Get("Authorization"))
		})
	}

	for _, uri := range []string{"", "localhost:26657", "ftp://node:21", "tcp://", "unix://"} {
		_, err := parseNodeURI(uri)
		require.Error(t, err, "%q", uri)
	}
}

// An error about a node URI never repeats the credentials in it, whether they sit in the
// password or, as a token, in the user part.
func TestParseNodeURIHidesCredentials(t *testing.T) {
	for _, uri := range []string{
		"https://user:s3cret@rpc.example.com:bad", // unparseable port
		"ftp://user:s3cret@rpc.example.com:21",    // unsupported scheme
		"ftp://s3cret@rpc.example.com:21",         // a token as the user
		"https://user:s3cret@",                    // no host
	} {
		_, err := parseNodeURI(uri)
		require.Error(t, err, uri)
		require.NotContains(t, err.Error(), "s3cret", uri)
	}
}

// host:port is not a URI; the error says which schemes are.
func TestParseNodeURINamesTheSchemes(t *testing.T) {
	_, err := parseNodeURI("localhost:26657")
	require.ErrorContains(t, err, "tcp://")
}

// The HTTP client is given the HTTP side of a websocket URI: it cannot dial a ws or wss
// network, nor post to such a URL.
func TestQueryURLMapsWebsocketSchemes(t *testing.T) {
	for uri, want := range map[string]string{
		"ws://node:26657":               "http://node:26657",
		"wss://rpc.example.com:443/rpc": "https://rpc.example.com:443/rpc",
		"tcp://node:26657":              "tcp://node:26657",
		"//node":                        "tcp://node:80",
		"https://rpc.example.com":       "https://rpc.example.com:443",
		unixSocketURI:                   unixSocketURI,
	} {
		u, err := parseNodeURI(uri)
		require.NoError(t, err)
		require.Equal(t, want, queryURL(u).String(), uri)
	}
}

// A node on a unix socket is reached there by its websocket as well as its queries.
func TestEventsOverUnixSocket(t *testing.T) {
	if testing.Short() {
		t.Skip("drives a real websocket against a local RPC server")
	}

	// A short directory: a socket path is limited to about a hundred bytes.
	dir, err := os.MkdirTemp("", "ev")
	require.NoError(t, err)
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	sock := filepath.Join(dir, "node.sock")

	node := newRPCNode(t)
	ln, err := net.Listen(networkUnix, sock)
	require.NoError(t, err)
	srv := &http.Server{Handler: node.handler, ReadHeaderTimeout: time.Second}
	go func() { _ = srv.Serve(ln) }()
	t.Cleanup(func() { _ = srv.Close() })

	// An hour's watchdog: only the websocket's headers can move the service on.
	sub := startRPCServiceAt(t, node, "unix://"+sock, tuning{
		watchdog:         time.Hour,
		retryInterval:    20 * time.Millisecond,
		prunedCheckAfter: time.Hour,
	})
	produceBlocks(t, node)

	requireInOrderThrough(t, sub, 101, 105)
}

// A node given by its websocket URI is queried over HTTP on the same address.
func TestEventsOverWebsocketSchemeURI(t *testing.T) {
	if testing.Short() {
		t.Skip("drives a real websocket against a local RPC server")
	}

	node := newRPCNode(t)
	sub := startRPCServiceAt(t, node, "ws://"+node.ln.Addr().String(), tuning{
		watchdog:         time.Hour,
		retryInterval:    20 * time.Millisecond,
		prunedCheckAfter: time.Hour,
	})
	produceBlocks(t, node)

	requireInOrderThrough(t, sub, 101, 105)
}

// The stream reports the highest header the node sent. Repeated and older headers are
// not news, and do not signal.
func TestHeaderStreamDeliversHighestHeader(t *testing.T) {
	node := newRPCNode(t)
	s := dialStream(t, streamConfig(t, node.addr(), time.Second, 5*time.Second))
	require.Eventually(t, func() bool { return node.subscribeCount() == 1 }, 2*time.Second, time.Millisecond)

	node.commit()
	node.commit()
	node.commit()
	waitLatest(t, s, 103)

	time.Sleep(50 * time.Millisecond) // let the signal for 103 land before draining it
	select {
	case <-s.Notify():
	default:
	}
	node.push(102)
	node.push(103)
	select {
	case <-s.Notify():
		t.Fatalf("a repeated or older header signalled a new one (latest %d)", s.Latest())
	case <-time.After(100 * time.Millisecond):
	}

	node.commit()
	waitLatest(t, s, 104)
	requireOpen(t, s)
}

// A node that rejects the subscription, as one at its subscription limits does, ends
// the stream with its answer; nothing on the client side retries it.
func TestHeaderStreamEndsOnSubscriptionError(t *testing.T) {
	node := newRPCNode(t)
	node.set(func(n *rpcNode) { n.subscribeErr = errors.New("max_subscription_clients 100 reached") })

	s := dialStream(t, streamConfig(t, node.addr(), time.Second, 5*time.Second))

	require.ErrorContains(t, requireEnded(t, s, 2*time.Second), "max_subscription_clients")
	require.Equal(t, 1, node.subscribeCount())
}

// A path that answers nothing, while neither end closes it, ends the stream once a pong
// is overdue.
func TestHeaderStreamDetectsHalfOpenConnection(t *testing.T) {
	node := newRPCNode(t)
	proxy := newBlackholeProxy(t, node.ln.Addr().String())
	s := dialStream(t, streamConfig(t, proxy.addr(), 20*time.Millisecond, 100*time.Millisecond))
	require.Eventually(t, func() bool { return node.subscribeCount() == 1 }, 2*time.Second, time.Millisecond)

	node.commit()
	waitLatest(t, s, 101)

	proxy.blackholeWebsocket()

	require.Error(t, requireEnded(t, s, time.Second))
}

// A chain that produces no block for longer than the pong wait is not a dead
// connection: pongs keep the stream open.
func TestHeaderStreamSurvivesIdleChain(t *testing.T) {
	node := newRPCNode(t)
	s := dialStream(t, streamConfig(t, node.addr(), 20*time.Millisecond, 100*time.Millisecond))
	require.Eventually(t, func() bool { return node.subscribeCount() == 1 }, 2*time.Second, time.Millisecond)

	time.Sleep(500 * time.Millisecond) // five pong waits without a block
	requireOpen(t, s)

	node.commit()
	waitLatest(t, s, 101)
}

// Close ends the stream at once, and says so.
func TestHeaderStreamCloseIsPrompt(t *testing.T) {
	node := newRPCNode(t)
	s, err := dialHeaders(context.Background(), streamConfig(t, node.addr(), time.Second, 5*time.Second))
	require.NoError(t, err)
	require.Eventually(t, func() bool { return node.subscribeCount() == 1 }, 2*time.Second, time.Millisecond)

	start := time.Now()
	s.Close()
	require.Less(t, time.Since(start), 100*time.Millisecond)
	require.ErrorIs(t, requireEnded(t, s, time.Second), errStreamClosed)

	s.Close() // a second Close is harmless
}

// newSilentListener accepts connections and holds them open without ever answering,
// the way a node or load balancer stuck behind a full backend does. accepted is
// signalled for every connection.
func newSilentListener(t *testing.T) (addr string, accepted <-chan struct{}) {
	t.Helper()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	signal := make(chan struct{}, 16)
	var (
		mu   sync.Mutex
		held []net.Conn
	)
	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			mu.Lock()
			held = append(held, c)
			mu.Unlock()
			select {
			case signal <- struct{}{}:
			default:
			}
		}
	}()
	t.Cleanup(func() {
		_ = ln.Close()
		mu.Lock()
		defer mu.Unlock()
		for _, c := range held {
			_ = c.Close()
		}
	})

	return ln.Addr().String(), signal
}

// A node that accepts the connection but never answers the handshake holds the dial
// only until its context ends.
func TestDialHeadersHonoursContext(t *testing.T) {
	addr, _ := newSilentListener(t)
	cfg := streamConfig(t, "tcp://"+addr, time.Second, 5*time.Second)
	cfg.timeout = time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := dialHeaders(ctx, cfg)
	require.Error(t, err)
	require.Less(t, time.Since(start), 500*time.Millisecond)
}

// Cancelling the dial's context, rather than letting it time out, ends a handshake the
// node never answers at once. Otherwise the service's shutdown would wait out the whole
// dial timeout.
func TestDialHeadersCancelledDuringHandshake(t *testing.T) {
	addr, accepted := newSilentListener(t)
	cfg := streamConfig(t, "tcp://"+addr, time.Second, 5*time.Second)
	cfg.timeout = time.Minute

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1)
	go func() {
		_, err := dialHeaders(ctx, cfg)
		done <- err
	}()

	requireSignal(t, accepted, "the dial to reach the node")
	start := time.Now()
	cancel()

	select {
	case err := <-done:
		require.Error(t, err)
		require.Less(t, time.Since(start), 500*time.Millisecond)
	case <-time.After(2 * time.Second):
		t.Fatal("the dial kept waiting on the handshake after its context was cancelled")
	}
}

// The service's HTTP connections to the node are its own, so they close when it stops
// rather than idling on until the transport times them out.
func TestEventsShutdownClosesNodeConnections(t *testing.T) {
	if testing.Short() {
		t.Skip("drives a real websocket against a local RPC server")
	}

	node := newRPCNode(t)
	bus := pubsub.NewBus()
	t.Cleanup(bus.Close)
	sub, err := bus.Subscribe()
	require.NoError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	svc, err := NewEvents(ctx, node.addr(), "test", bus,
		tuning{watchdog: time.Hour, retryInterval: 20 * time.Millisecond, prunedCheckAfter: time.Hour}.options()...)
	require.NoError(t, err)
	require.Eventually(t, func() bool { return node.subscribeCount() >= 1 }, 2*time.Second, 10*time.Millisecond)
	produceBlocks(t, node)

	requireInOrderThrough(t, sub, 101, 103)
	require.Positive(t, node.openHTTPConns(), "the service queries the node over HTTP")

	svc.Shutdown()
	require.Eventually(t, func() bool { return node.openHTTPConns() == 0 }, 2*time.Second, 10*time.Millisecond,
		"the service's HTTP connections outlived it")
}

// Shutdown does not wait on a websocket dial stuck in the handshake.
func TestEventsShutdownDuringStalledHandshake(t *testing.T) {
	addr, accepted := newSilentListener(t)

	bus := pubsub.NewBus()
	t.Cleanup(bus.Close)
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	svc, err := NewEvents(ctx, "tcp://"+addr, "test", bus,
		WithStartHeight(100), WithWatchdogTimeout(time.Hour), WithFetchTimeout(time.Minute))
	require.NoError(t, err)

	requireSignal(t, accepted, "the service to dial the node")
	start := time.Now()
	stopped := make(chan struct{})
	go func() {
		svc.Shutdown()
		close(stopped)
	}()

	select {
	case <-stopped:
		require.Less(t, time.Since(start), 500*time.Millisecond)
	case <-time.After(3 * time.Second):
		t.Fatal("Shutdown waited on the stalled handshake")
	}
}

// A node behind TLS is reached over wss.
func TestDialHeadersOverTLS(t *testing.T) {
	node := newRPCNode(t)
	srv := httptest.NewUnstartedServer(node.handler)
	srv.StartTLS()
	t.Cleanup(srv.Close)

	pool := x509.NewCertPool()
	pool.AddCert(srv.Certificate())
	cfg := streamConfig(t, "https://"+srv.Listener.Addr().String(), time.Second, 5*time.Second)
	cfg.tlsConfig = &tls.Config{RootCAs: pool, MinVersion: tls.VersionTLS12}

	s := dialStream(t, cfg)
	require.Eventually(t, func() bool { return node.subscribeCount() == 1 }, 2*time.Second, time.Millisecond)

	node.commit()
	waitLatest(t, s, 101)
}

// A frame that is not JSON-RPC is skipped; the stream carries on.
func TestHeaderStreamIgnoresMalformedFrames(t *testing.T) {
	query := blkHeaderQuery().String()
	release := make(chan struct{})
	upgrader := websocket.Upgrader{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		var req rpctypes.RPCRequest
		if err := conn.ReadJSON(&req); err != nil {
			return
		}
		_ = conn.WriteMessage(websocket.TextMessage, []byte("{not json"))
		_ = conn.WriteJSON(rpctypes.NewRPCSuccessResponse(req.ID, coretypes.ResultEvent{
			Query: query,
			Data:  cmtypes.EventDataNewBlockHeader{Header: cmtypes.Header{Height: 7}},
		}))
		<-release
	}))
	t.Cleanup(func() {
		close(release)
		srv.Close()
	})

	s := dialStream(t, streamConfig(t, "tcp://"+srv.Listener.Addr().String(), time.Second, 5*time.Second))

	waitLatest(t, s, 7)
	requireOpen(t, s)
}
