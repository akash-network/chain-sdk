package events

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"

	cmtjson "github.com/cometbft/cometbft/libs/json"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	cmtypes "github.com/cometbft/cometbft/types"
)

// errStreamClosed is why a header stream ended when its owner closed it.
var errStreamClosed = errors.New("events: header stream closed")

// headerStream is one live subscription to the node's block headers. It ends exactly
// once, on its own when the connection or the subscription under it fails, or when it
// is closed.
type headerStream interface {
	// Notify is signalled when a header above Latest arrives. Signals coalesce: one
	// pending signal stands for every header since the last one was received.
	Notify() <-chan struct{}
	// Latest is the highest header height the stream has seen.
	Latest() int64
	// Done is closed once the stream has ended.
	Done() <-chan struct{}
	// Err is why the stream ended. It is nil until Done is closed.
	Err() error
	// Close ends the stream and waits for its goroutines.
	Close()
}

// dialFunc opens a header stream. ctx bounds the dial only, not the stream.
type dialFunc func(ctx context.Context) (headerStream, error)

// wsConfig is how dialHeaders reaches the node and decides the connection is dead.
type wsConfig struct {
	target wsTarget
	// query is the subscription query, blkHeaderQuery().String().
	query string
	// timeout bounds the dial and handshake, and every write.
	timeout time.Duration
	// pingPeriod is how often the stream pings the node; pongWait is how long it goes
	// without a pong or a message before it treats the connection as dead.
	pingPeriod time.Duration
	pongWait   time.Duration
	// tlsConfig is nil in production, which means the system's roots.
	tlsConfig *tls.Config
}

// wsStream is a headerStream over a websocket the stream owns.
type wsStream struct {
	conn   *websocket.Conn
	cfg    wsConfig
	notify chan struct{}
	latest atomic.Int64
	done   chan struct{}
	once   sync.Once
	err    error
	wg     sync.WaitGroup
}

// dialHeaders opens a websocket to the node, subscribes to block headers and returns
// the stream. It does not wait for the node to accept the subscription: a rejection
// arrives as an error answer and ends the stream, so the caller decides when to try
// again rather than the client retrying on its own.
//
// The connection is watched by ping and pong: the stream pings every pingPeriod, and
// any pong or message within pongWait proves the path alive. A path that stops
// answering, whether reset, closed or silently dropped along the way, ends the stream
// within pongWait; a chain that merely produces no block does not.
//
// Cancelling ctx ends the dial at once, in any phase. The websocket library bounds the
// handshake by deadlines only, so the connection is also poked when ctx ends, until the
// stream is handed over.
func dialHeaders(ctx context.Context, cfg wsConfig) (headerStream, error) {
	nd := &net.Dialer{Timeout: cfg.timeout}
	dial := nd.DialContext
	proxy := http.ProxyFromEnvironment
	if cfg.target.network == networkUnix {
		// The URL's host only stands in for the socket, so there is nothing to proxy.
		proxy = nil
		dial = func(ctx context.Context, _, _ string) (net.Conn, error) {
			return nd.DialContext(ctx, networkUnix, cfg.target.address)
		}
	}

	// poke makes every pending and future read and write on c fail at once.
	poke := func(c net.Conn) { _ = c.SetDeadline(time.Unix(1, 0)) }

	var (
		rawMu sync.Mutex
		raw   net.Conn
	)
	d := websocket.Dialer{
		NetDialContext: func(dialCtx context.Context, network, addr string) (net.Conn, error) {
			c, err := dial(dialCtx, network, addr)
			if err != nil {
				return nil, err
			}
			rawMu.Lock()
			raw = c
			rawMu.Unlock()
			if ctx.Err() != nil {
				poke(c) // ctx ended before the hook below could see this connection
			}
			return c, nil
		},
		Proxy:            proxy,
		HandshakeTimeout: cfg.timeout,
		TLSClientConfig:  cfg.tlsConfig,
	}

	stop := context.AfterFunc(ctx, func() {
		rawMu.Lock()
		defer rawMu.Unlock()
		if raw != nil {
			poke(raw)
		}
	})
	defer stop()

	conn, resp, err := d.DialContext(ctx, cfg.target.url, cfg.target.header)
	if resp != nil && resp.Body != nil {
		_ = resp.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("events: dial %s: %w", cfg.target.url, err)
	}

	req, err := rpctypes.MapToRequest(rpctypes.JSONRPCIntID(1), "subscribe", map[string]interface{}{"query": cfg.query})
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("events: subscribe request: %w", err)
	}
	if err := conn.SetWriteDeadline(time.Now().Add(cfg.timeout)); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("events: subscribe: %w", err)
	}
	if err := conn.WriteJSON(req); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("events: subscribe: %w", err)
	}

	// Detach the cancellation hook before handing the stream over. If ctx ended first,
	// the hook may already have poisoned the connection's deadlines.
	if !stop() {
		_ = conn.Close()
		return nil, fmt.Errorf("events: dial %s: %w", cfg.target.url, ctx.Err())
	}

	s := &wsStream{
		conn:   conn,
		cfg:    cfg,
		notify: make(chan struct{}, 1),
		done:   make(chan struct{}),
	}
	s.wg.Add(2)
	go s.read()
	go s.ping()

	return s, nil
}

func (s *wsStream) Notify() <-chan struct{} { return s.notify }

func (s *wsStream) Latest() int64 { return s.latest.Load() }

func (s *wsStream) Done() <-chan struct{} { return s.done }

func (s *wsStream) Err() error {
	select {
	case <-s.done:
		return s.err
	default:
		return nil
	}
}

func (s *wsStream) Close() {
	s.end(errStreamClosed)
	s.wg.Wait()
}

// end ends the stream with err, once: closing the socket stops both goroutines, and
// drops the subscription on the node's side too.
func (s *wsStream) end(err error) {
	s.once.Do(func() {
		s.err = err
		_ = s.conn.Close()
		close(s.done)
	})
}

// read receives the node's messages until the connection fails. Every pong and every
// message pushes the read deadline back by pongWait.
func (s *wsStream) read() {
	defer s.wg.Done()

	alive := func() error { return s.conn.SetReadDeadline(time.Now().Add(s.cfg.pongWait)) }
	if err := alive(); err != nil {
		s.end(fmt.Errorf("events: read: %w", err))
		return
	}
	s.conn.SetPongHandler(func(string) error { return alive() })

	for {
		_, data, err := s.conn.ReadMessage()
		if err != nil {
			s.end(fmt.Errorf("events: read: %w", err))
			return
		}
		if err := alive(); err != nil {
			s.end(fmt.Errorf("events: read: %w", err))
			return
		}

		var resp rpctypes.RPCResponse
		if err := json.Unmarshal(data, &resp); err != nil {
			continue
		}
		if resp.Error != nil {
			s.end(fmt.Errorf("events: node ended the subscription: %w", resp.Error))
			return
		}

		var ev ctypes.ResultEvent
		if err := cmtjson.Unmarshal(resp.Result, &ev); err != nil {
			continue
		}
		hdr, ok := ev.Data.(cmtypes.EventDataNewBlockHeader)
		if !ok || hdr.Header.Height <= s.latest.Load() {
			continue
		}

		s.latest.Store(hdr.Header.Height)
		select {
		case s.notify <- struct{}{}:
		default:
		}
	}
}

// ping pings the node every pingPeriod until the stream ends. A ping that cannot be
// written ends it.
func (s *wsStream) ping() {
	defer s.wg.Done()

	t := time.NewTicker(s.cfg.pingPeriod)
	defer t.Stop()

	for {
		select {
		case <-s.done:
			return
		case <-t.C:
			if err := s.conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(s.cfg.timeout)); err != nil {
				s.end(fmt.Errorf("events: ping: %w", err))
				return
			}
		}
	}
}
