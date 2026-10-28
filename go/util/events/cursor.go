package events

import "time"

// cursor is the service's position in the chain: the next height to publish, the
// heights being fetched ahead of it, and the fetched ones waiting for their turn. It
// decides what to fetch and what may be published, and nothing else. It starts no
// goroutines and reads no clock, so the run loop drives it and tests can step it.
//
// Heights are published strictly in order, each once. At most window heights, starting
// at the next one to publish, are in flight or held at a time, so a back-fill overlaps
// its waits on the node without memory growing with the gap.
//
// The next height failing parks the cursor: it is retried once per retry interval, for
// as long as it takes, because a skipped height is a lost event. A height above it
// failing stops prefetching there until the cursor reaches it, so a node that lags or
// refuses requests is asked for one height at a time instead of being hammered.
type cursor struct {
	next             int64
	known            int64
	window           int
	retry            time.Duration
	prunedCheckAfter time.Duration
	inflight         map[int64]struct{}
	ready            map[int64][]interface{}
	// failed holds the heights above next whose fetch failed. Nothing from the lowest
	// of them on is fetched until next reaches it; each is dropped once it is next, and
	// then fetched as the cursor height. Being within the window, it holds at most
	// window heights.
	failed map[int64]struct{}
	park   parked
}

// parked is the state of the next height while it keeps failing. The zero value means
// the cursor is not parked.
type parked struct {
	height  int64
	since   time.Time
	nextTry time.Time
	// escalated and pruneChecked make a long stall logged at error level, and the node
	// asked whether the height is pruned, once rather than on every retry.
	escalated    bool
	pruneChecked bool
}

// parkEvent is what the run loop logs, and acts on, when the cursor height fails.
type parkEvent struct {
	height     int64
	err        error
	first      bool          // first failure of this height
	escalate   bool          // failing for prunedCheckAfter, reported once
	pruneCheck bool          // time to ask the node whether the height is pruned, reported once
	stalled    time.Duration // time since the first failure
}

// readyBlock is the next height to publish, in order.
type readyBlock struct {
	height    int64
	events    []interface{}
	recovered bool          // the height had been parked
	stalled   time.Duration // how long it was parked, when recovered
}

// newCursor returns a cursor that has published everything up to last.
func newCursor(last int64, window int, retry, prunedCheckAfter time.Duration) *cursor {
	return &cursor{
		next:             last + 1,
		known:            last,
		window:           window,
		retry:            retry,
		prunedCheckAfter: prunedCheckAfter,
		inflight:         make(map[int64]struct{}),
		ready:            make(map[int64][]interface{}),
		failed:           make(map[int64]struct{}),
	}
}

// observe records that the chain has reached height. Lower heights change nothing.
func (c *cursor) observe(height int64) {
	if height > c.known {
		c.known = height
	}
}

// toFetch returns the heights to start fetching at now and marks them in flight.
func (c *cursor) toFetch(now time.Time) []int64 {
	last := c.next + int64(c.window) - 1
	if c.known < last {
		last = c.known
	}
	for h := range c.failed {
		if h <= last {
			last = h - 1
		}
	}

	var heights []int64
	for h := c.next; h <= last; h++ {
		if _, ok := c.inflight[h]; ok {
			continue
		}
		if _, ok := c.ready[h]; ok {
			continue
		}
		if h == c.park.height && now.Before(c.park.nextTry) {
			continue
		}
		c.inflight[h] = struct{}{}
		heights = append(heights, h)
	}

	return heights
}

// done records the outcome of fetching height. It reports true when the cursor height
// failed and parked, with what the run loop should log and do about it.
func (c *cursor) done(height int64, evs []interface{}, err error, now time.Time) (parkEvent, bool) {
	delete(c.inflight, height)

	if err == nil {
		c.ready[height] = evs
		return parkEvent{}, false
	}

	if height != c.next {
		c.failed[height] = struct{}{}
		return parkEvent{}, false
	}

	first := c.park.height != height
	if first {
		c.park = parked{height: height, since: now}
	}
	c.park.nextTry = now.Add(c.retry)

	ev := parkEvent{height: height, err: err, first: first, stalled: now.Sub(c.park.since)}
	if ev.stalled >= c.prunedCheckAfter {
		if !c.park.escalated {
			c.park.escalated = true
			ev.escalate = true
		}
		if !c.park.pruneChecked {
			c.park.pruneChecked = true
			ev.pruneCheck = true
		}
	}

	return ev, true
}

// pop returns the next height to publish, if it has been fetched, and moves the cursor
// past it. The caller must publish it before calling pop again.
func (c *cursor) pop(now time.Time) (readyBlock, bool) {
	evs, ok := c.ready[c.next]
	if !ok {
		return readyBlock{}, false
	}
	delete(c.ready, c.next)

	r := readyBlock{height: c.next, events: evs}
	if c.park.height == c.next {
		r.recovered = true
		r.stalled = now.Sub(c.park.since)
		c.park = parked{}
	}

	c.next++
	delete(c.failed, c.next)

	return r, true
}

// wakeAt is when the parked height is next due. It reports false when the cursor is not
// parked, or when the parked height is already being fetched again.
func (c *cursor) wakeAt() (time.Time, bool) {
	if c.park.height == 0 {
		return time.Time{}, false
	}
	if _, ok := c.inflight[c.park.height]; ok {
		return time.Time{}, false
	}
	return c.park.nextTry, true
}

// parkedOn reports whether the cursor is parked on height.
func (c *cursor) parkedOn(height int64) bool {
	return height != 0 && c.park.height == height
}

// last is the highest height published so far.
func (c *cursor) last() int64 {
	return c.next - 1
}
