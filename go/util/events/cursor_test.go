package events

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	errFetch = errors.New("fetch failed")
	t0       = time.Unix(0, 0)
)

// evs stands in for a height's decoded events: one marker naming the height.
func evs(h int64) []interface{} { return []interface{}{h} }

// The cursor hands out up to window heights at once, never beyond the highest height
// known to exist, and never a height that is already in flight.
func TestCursorFetchesUpToWindowAndKnown(t *testing.T) {
	require.Empty(t, newCursor(100, 4, time.Second, time.Hour).toFetch(t0), "nothing is known above the start height")

	c := newCursor(100, 4, time.Second, time.Hour)
	c.observe(110)
	require.Equal(t, []int64{101, 102, 103, 104}, c.toFetch(t0))
	require.Empty(t, c.toFetch(t0), "heights already in flight are not fetched twice")

	c = newCursor(100, 4, time.Second, time.Hour)
	c.observe(102)
	c.observe(99)
	require.Equal(t, []int64{101, 102}, c.toFetch(t0), "nothing beyond the highest known height, which never goes down")
}

// Fetches complete in any order; heights come out strictly in order, each once.
func TestCursorPublishesInOrder(t *testing.T) {
	c := newCursor(100, 4, time.Second, time.Hour)
	c.observe(103)
	require.Equal(t, []int64{101, 102, 103}, c.toFetch(t0))

	_, parked := c.done(103, evs(103), nil, t0)
	require.False(t, parked)
	_, ok := c.pop(t0)
	require.False(t, ok, "103 waits for 101 and 102")

	c.done(101, evs(101), nil, t0)
	c.done(102, evs(102), nil, t0)
	for _, want := range []int64{101, 102, 103} {
		r, ok := c.pop(t0)
		require.True(t, ok)
		require.Equal(t, readyBlock{height: want, events: evs(want)}, r)
	}
	_, ok = c.pop(t0)
	require.False(t, ok)
	require.Equal(t, int64(103), c.last())
}

// The cursor height failing parks it until its retry is due, and the wake-up time is
// reported only while no retry is in flight.
func TestCursorParksAndRetriesWhenDue(t *testing.T) {
	c := newCursor(100, 4, time.Second, time.Hour)
	c.observe(101)
	require.Equal(t, []int64{101}, c.toFetch(t0))

	pe, parked := c.done(101, nil, errFetch, t0)
	require.True(t, parked)
	require.Equal(t, parkEvent{height: 101, err: errFetch, first: true}, pe)
	require.True(t, c.parkedOn(101))

	require.Empty(t, c.toFetch(t0.Add(999*time.Millisecond)), "not due before the retry interval")
	at, ok := c.wakeAt()
	require.True(t, ok)
	require.Equal(t, t0.Add(time.Second), at)

	require.Equal(t, []int64{101}, c.toFetch(t0.Add(time.Second)))
	_, ok = c.wakeAt()
	require.False(t, ok, "no wake-up while the retry is in flight")

	pe, parked = c.done(101, nil, errFetch, t0.Add(time.Second))
	require.True(t, parked)
	require.False(t, pe.first)
	require.Equal(t, time.Second, pe.stalled)
	at, ok = c.wakeAt()
	require.True(t, ok)
	require.Equal(t, t0.Add(2*time.Second), at)
}

// A long stall is escalated, and the pruned check started, exactly once.
func TestCursorEscalatesAndPruneChecksOnce(t *testing.T) {
	c := newCursor(100, 1, time.Second, 30*time.Second)
	c.observe(101)

	for _, tc := range []struct {
		after time.Duration
		once  bool
	}{
		{after: 0},
		{after: 10 * time.Second},
		{after: 30 * time.Second, once: true},
		{after: 40 * time.Second},
	} {
		now := t0.Add(tc.after)
		require.Equal(t, []int64{101}, c.toFetch(now), "after %s", tc.after)
		pe, parked := c.done(101, nil, errFetch, now)
		require.True(t, parked)
		require.Equal(t, tc.once, pe.escalate, "escalate after %s", tc.after)
		require.Equal(t, tc.once, pe.pruneCheck, "pruned check after %s", tc.after)
		require.Equal(t, tc.after, pe.stalled)
	}
}

// A height above the cursor that fails stops prefetching at it. It is fetched again
// only once it is the cursor height, where normal parking applies.
func TestCursorStopsPrefetchAtFailedHeight(t *testing.T) {
	c := newCursor(100, 4, time.Second, time.Hour)
	c.observe(110)
	require.Equal(t, []int64{101, 102, 103, 104}, c.toFetch(t0))

	_, parked := c.done(102, nil, errFetch, t0)
	require.False(t, parked, "a failure above the cursor does not park it")
	require.False(t, c.parkedOn(102))
	require.Empty(t, c.toFetch(t0.Add(time.Hour)), "nothing at or above the failed height is fetched before the cursor reaches it")

	c.done(101, evs(101), nil, t0)
	r, ok := c.pop(t0)
	require.True(t, ok)
	require.Equal(t, int64(101), r.height)
	require.Equal(t, []int64{102, 105}, c.toFetch(t0), "the failed height is fetched as the cursor height; 103 and 104 are still in flight")

	pe, parked := c.done(102, nil, errFetch, t0)
	require.True(t, parked)
	require.True(t, pe.first)
}

// A parked height that is finally served reports how long it was parked, and the
// cursor is no longer parked.
func TestCursorRecoveryReportsStall(t *testing.T) {
	c := newCursor(100, 1, time.Second, time.Hour)
	c.observe(101)
	c.toFetch(t0)
	c.done(101, nil, errFetch, t0)

	later := t0.Add(3 * time.Second)
	require.Equal(t, []int64{101}, c.toFetch(later))
	_, parked := c.done(101, evs(101), nil, later)
	require.False(t, parked)

	r, ok := c.pop(later)
	require.True(t, ok)
	require.Equal(t, readyBlock{height: 101, events: evs(101), recovered: true, stalled: 3 * time.Second}, r)
	require.False(t, c.parkedOn(101))
	_, ok = c.wakeAt()
	require.False(t, ok)
}

// Whatever order fetches complete or fail in, the cursor never has more than window
// heights outstanding, never hands out a height twice, never refetches a failed height
// before it is due, and publishes every height exactly once, in order.
func TestCursorNeverExceedsWindow(t *testing.T) {
	const window = 4
	rng := rand.New(rand.NewSource(1)) //nolint:gosec // seeded on purpose: a reproducible simulation, not security
	c := newCursor(100, window, time.Second, time.Hour)

	var (
		now       = t0
		known     = int64(100)
		next      = int64(101) // next height the cursor must publish
		inflight  []int64
		failedAt  = map[int64]time.Time{} // failures of the cursor height: not due before +1s
		failedAbv = map[int64]bool{}      // failures above the cursor: wait for the cursor
	)

	for step := 0; step < 2000; step++ {
		now = now.Add(time.Duration(rng.Intn(300)) * time.Millisecond)
		known += rng.Int63n(2)
		c.observe(known)

		for _, h := range c.toFetch(now) {
			require.NotContains(t, inflight, h, "step %d: height %d handed out twice", step, h)
			require.Less(t, h, next+window, "step %d: height %d beyond the window", step, h)
			require.LessOrEqual(t, h, known, "step %d: height %d not known to exist", step, h)
			if failedAbv[h] {
				require.Equal(t, next, h, "step %d: failed height %d refetched before the cursor reached it", step, h)
				delete(failedAbv, h)
			}
			if at, ok := failedAt[h]; ok {
				require.False(t, now.Before(at.Add(time.Second)), "step %d: parked height %d refetched early", step, h)
			}
			inflight = append(inflight, h)
		}

		if len(inflight) > 0 && rng.Intn(2) == 0 {
			i := rng.Intn(len(inflight))
			h := inflight[i]
			inflight = append(inflight[:i], inflight[i+1:]...)
			if rng.Intn(4) == 0 {
				if _, parked := c.done(h, nil, errFetch, now); parked {
					require.Equal(t, next, h)
					failedAt[h] = now
				} else {
					require.Greater(t, h, next)
					failedAbv[h] = true
				}
			} else {
				c.done(h, evs(h), nil, now)
			}
		}

		for {
			r, ok := c.pop(now)
			if !ok {
				break
			}
			require.Equal(t, next, r.height, "step %d: out of order", step)
			require.Equal(t, evs(next), r.events)
			delete(failedAt, next)
			next++
		}
	}

	t.Logf("published through height %d of %d known", next-1, known)
	require.Greater(t, next, int64(300), "the simulation made progress")
	require.Equal(t, next-1, c.last())
}
