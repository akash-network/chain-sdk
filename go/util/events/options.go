package events

import (
	"errors"
	"fmt"
	"time"

	"cosmossdk.io/log"
)

const (
	// DefaultWatchdogTimeout is how long the service waits for a new block before it
	// treats the block subscription as dead, re-subscribes and fetches the missed
	// blocks itself.
	DefaultWatchdogTimeout = 8 * time.Second

	defaultMinRetryInterval = 2 * time.Second
	// defaultPrunedCheckAfter is how long a height stays stalled before we ask the node
	// whether it has been pruned. The answer only explains the stall in the log: the
	// cursor never skips a height.
	defaultPrunedCheckAfter = 30 * time.Second
)

// options holds the settings the service runs on.
type options struct {
	watchdog         time.Duration
	minRetryInterval time.Duration
	prunedCheckAfter time.Duration
	log              log.Logger
	startHeight      int64
	startHeightSet   bool
}

// Option configures the events service. It follows the functional options pattern,
// so settings can be added without breaking NewEvents callers. An option that
// returns an error makes NewEvents fail with that error.
type Option func(*options) error

// WithWatchdogTimeout sets how long the service waits for a new block before it
// treats the block subscription as dead, re-subscribes and fetches the missed blocks
// itself. It must be positive. The default is DefaultWatchdogTimeout.
func WithWatchdogTimeout(d time.Duration) Option {
	return func(o *options) error {
		if d <= 0 {
			return fmt.Errorf("events: watchdog timeout must be positive, got %s", d)
		}
		o.watchdog = d
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
// node no longer retains parks the service until a node that retains it answers.
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

// withMinRetryInterval sets how often a height that cannot be fetched is retried.
// It is unexported, so tests can run fast without widening the public API.
func withMinRetryInterval(d time.Duration) Option {
	return func(o *options) error {
		if d <= 0 {
			return fmt.Errorf("events: retry interval must be positive, got %s", d)
		}
		o.minRetryInterval = d
		return nil
	}
}

// withPrunedCheckAfter sets how long a height stays stalled before the node is asked
// whether it has been pruned, to explain the stall in the log. It is unexported for the
// same reason.
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
		minRetryInterval: defaultMinRetryInterval,
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
