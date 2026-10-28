package events

import (
	"errors"
	"fmt"
	"time"

	"cosmossdk.io/log"
)

const (
	// DefaultWatchdogTimeout is how long the service waits for a new block header before
	// it asks the node how far the chain has got and fetches the blocks it has missed
	// itself (see WithWatchdogTimeout).
	DefaultWatchdogTimeout = 8 * time.Second

	// DefaultFetchTimeout is how long the service waits for the node to answer a block
	// fetch or a status query. A call that has not returned by then counts as a failed
	// attempt and is retried. It also bounds the websocket dial, handshake included, and
	// every websocket write.
	DefaultFetchTimeout = 10 * time.Second

	// DefaultRetryInterval is how often a block that cannot be fetched is retried (see
	// WithRetryInterval).
	DefaultRetryInterval = 5 * time.Second

	// DefaultPrefetch is how many blocks the service fetches at once (see WithPrefetch).
	DefaultPrefetch = 4

	// defaultPingPeriod is how often the service pings the node over its websocket, and
	// defaultPongWait how long the websocket may go without a pong or a message before
	// the service treats the connection as dead and redials.
	defaultPingPeriod = 5 * time.Second
	defaultPongWait   = 15 * time.Second

	// defaultPrunedCheckAfter is how long a height stays stalled before we ask the node
	// whether it has been pruned. The answer only explains the stall in the log: the
	// cursor never skips a height.
	defaultPrunedCheckAfter = 30 * time.Second
)

// options holds the settings the service runs on.
type options struct {
	watchdog         time.Duration
	fetchTimeout     time.Duration
	retryInterval    time.Duration
	prefetch         int
	pingPeriod       time.Duration
	pongWait         time.Duration
	prunedCheckAfter time.Duration
	log              log.Logger
	startHeight      int64
	startHeightSet   bool
}

// Option configures the events service. It follows the functional options pattern,
// so settings can be added without breaking NewEvents callers. An option that
// returns an error makes NewEvents fail with that error.
type Option func(*options) error

// WithWatchdogTimeout sets how long the service waits for a new block header before it
// asks the node how far the chain has got and fetches the blocks it has missed itself.
// It keeps asking once per window until a new header arrives. It does not judge the
// websocket: a dead connection is found by ping and pong and redialed. It must be
// positive. The default is DefaultWatchdogTimeout.
func WithWatchdogTimeout(d time.Duration) Option {
	return func(o *options) error {
		if d <= 0 {
			return fmt.Errorf("events: watchdog timeout must be positive, got %s", d)
		}
		o.watchdog = d
		return nil
	}
}

// WithFetchTimeout sets how long the service waits for the node to answer a block fetch
// or a status query before it counts the call as failed and retries it. It also bounds
// the websocket dial, handshake included, and every websocket write, so it must leave
// room for a slow node to complete a handshake. It must be positive. The default is
// DefaultFetchTimeout.
func WithFetchTimeout(d time.Duration) Option {
	return func(o *options) error {
		if d <= 0 {
			return fmt.Errorf("events: fetch timeout must be positive, got %s", d)
		}
		o.fetchTimeout = d
		return nil
	}
}

// WithRetryInterval sets how often a block that cannot be fetched is retried. The
// service never gives up on a block, because a skipped block is a lost event, so a
// block the node cannot serve is retried at this pace for as long as the service runs.
// It must be positive. The default is DefaultRetryInterval.
func WithRetryInterval(d time.Duration) Option {
	return func(o *options) error {
		if d <= 0 {
			return fmt.Errorf("events: retry interval must be positive, got %s", d)
		}
		o.retryInterval = d
		return nil
	}
}

// WithPrefetch sets how many blocks the service may fetch at once while it catches up.
// Blocks are still published strictly in order; prefetching only overlaps the waits on
// the node. 1 fetches one block at a time. It must be at least 1. The default is
// DefaultPrefetch.
func WithPrefetch(n int) Option {
	return func(o *options) error {
		if n < 1 {
			return fmt.Errorf("events: prefetch must be at least 1, got %d", n)
		}
		o.prefetch = n
		return nil
	}
}

// WithLogger sets the logger the service reports to, including every loss and
// re-establishment of the block subscription. It must not be nil. By default the
// service uses the logger carried by the context passed to NewEvents (see
// util/ctxlog), which discards everything when the context carries none.
func WithLogger(l log.Logger) Option {
	return func(o *options) error {
		if l == nil {
			return errors.New("events: logger must not be nil")
		}
		o.log = l
		return nil
	}
}

// WithStartHeight makes the service publish events from block h+1 on, where h is the
// last block the caller has already accounted for, for example the height its
// bootstrap queries were pinned to. Without it the service starts from the node's
// latest height when NewEvents runs, and anything the caller has not seen at or below
// that height is neither queried nor published. h must not be negative. A height the
// node no longer retains parks the cursor until a node that retains it answers.
func WithStartHeight(h int64) Option {
	return func(o *options) error {
		if h < 0 {
			return fmt.Errorf("events: start height must not be negative, got %d", h)
		}
		o.startHeight = h
		o.startHeightSet = true
		return nil
	}
}

// withPing sets how often the websocket is pinged and how long it may go without a pong
// or a message before it is treated as dead. It is unexported, so tests can run fast
// without widening the public API. period must be positive and below wait.
func withPing(period, wait time.Duration) Option {
	return func(o *options) error {
		if period <= 0 || wait <= period {
			return fmt.Errorf("events: ping period must be positive and below the pong wait, got %s and %s", period, wait)
		}
		o.pingPeriod, o.pongWait = period, wait
		return nil
	}
}

// withPrunedCheckAfter sets how long a height stays stalled before the node is asked
// whether it has been pruned, to explain the stall in the log. It is unexported, so
// tests can run fast without widening the public API.
func withPrunedCheckAfter(d time.Duration) Option {
	return func(o *options) error {
		if d <= 0 {
			return fmt.Errorf("events: pruned check delay must be positive, got %s", d)
		}
		o.prunedCheckAfter = d
		return nil
	}
}

func newOptions(opts ...Option) (options, error) {
	o := options{
		watchdog:         DefaultWatchdogTimeout,
		fetchTimeout:     DefaultFetchTimeout,
		retryInterval:    DefaultRetryInterval,
		prefetch:         DefaultPrefetch,
		pingPeriod:       defaultPingPeriod,
		pongWait:         defaultPongWait,
		prunedCheckAfter: defaultPrunedCheckAfter,
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(&o); err != nil {
			return options{}, err
		}
	}

	return o, nil
}
