# events

Package `events` (`pkg.akt.dev/go/util/events`) publishes the Akash audit, deployment
and market events a node emits onto a `pubsub.Bus`: in the order the chain executed
them, each once, with no height skipped. The recognized set is in `decode.go`.

## Usage

```go
bus := pubsub.NewBus()
defer bus.Close()

sub, err := bus.Subscribe()
if err != nil {
	return err
}

svc, err := events.NewEvents(ctx, "tcp://localhost:26657", "my-app", bus,
	events.WithStartHeight(bootstrapHeight)) // optional: the height your state was read at
if err != nil {
	return err
}
defer svc.Shutdown()

for {
	select {
	case ev := <-sub.Events():
		if closed, ok := ev.(*mtypes.EventLeaseClosed); ok {
			_ = closed // handle it
		}
	case <-svc.Done():
		return svc.Err() // non-nil only if the bus refused an event
	}
}
```

`NewEvents` fails on an invalid node URI or option and, without `WithStartHeight`, when
the node does not answer its first status query. Once started, the service stops only on
`Shutdown`, when its context ends, or when the bus refuses an event.

## Guarantees

- Events go out in execution order: PreBlock and BeginBlock events, then transactions,
  then EndBlock events.
- Each height is published once and none is skipped. A height that cannot be fetched is
  retried forever, and every later event waits behind it. A bus that refuses an event
  stops the service, possibly partway through a height.
- A recognized event that cannot be decoded, such as one with an enum value from a newer
  chain, is skipped and logged at error level, since a retry cannot change the outcome.
- Without `WithStartHeight` the service publishes from the block after the node's latest
  height, so it replays nothing. Its position is not persisted and events carry no height:
  to start without a gap after reading chain state, pass the height it was read at.

## Design

**Own connection.** The service opens its own connection from the node URI: an HTTP
JSON-RPC client with its own pool for block results and status queries, and a websocket
subscribed to new block headers. Its lifecycle is separate from any client the caller
holds. Accepted URIs are `tcp`, `http`, `https`, `ws`, `wss` and `unix`; `//host[:port]`
means `tcp` and a bare `host:port` is rejected. A missing port defaults to 80, or 443 for
`https` and `wss`. A `ws` or `wss` URI is queried over `http` or `https`. Credentials in
the URI are sent as Basic auth and masked in URI errors.

**Liveness.** The websocket is pinged every 5 s and counts as dead after 15 s without a
pong or a message. A dead or rejected connection is redialed at once, then once per retry
interval, for as long as it takes; the service never subscribes twice on one connection.
While no new header arrives, a watchdog asks the node for its latest height once per
window, so events keep flowing even while the websocket cannot connect. On a live
connection, a quiet window means a slow chain, not an outage.

**Run loop.** One goroutine owns all state and only selects on results and timers. Every
call to the node (dial, block fetch, status) runs in a goroutine of its own, bounded by
the fetch timeout and cancelled on shutdown, so the loop never blocks on the network and
`Shutdown` returns promptly.

**Cursor.** `cursor.go` decides what to fetch and what to publish. It starts no
goroutines and reads no clock.

- At most `WithPrefetch` heights, from the next one to publish on, are in flight or
  waiting at a time. They are published strictly in order, however the fetches complete.
- When the next height fails, the cursor parks on it and retries it once per retry
  interval. After 30 s stalled it logs an error once and asks the node whether the height
  was pruned; the answer only explains the stall, the height is never skipped.
- When a height above the cursor fails, no new fetch from it on starts until the cursor
  reaches it. It is then retried as the cursor height, and prefetching resumes.
- A result for another height than the one requested counts as a failed fetch.

## Options

| Option                | Default         | Effect                                                                   |
|-----------------------|-----------------|--------------------------------------------------------------------------|
| `WithWatchdogTimeout` | 8 s             | Time without a new header before the node is polled for its height.      |
| `WithFetchTimeout`    | 10 s            | Bound on every call to the node, the websocket dial and writes included. |
| `WithRetryInterval`   | 5 s             | Pace of retries for a failed height, and of redials.                     |
| `WithPrefetch`        | 4               | Heights fetched ahead at once; 1 fetches one at a time.                  |
| `WithStartHeight`     | node's latest   | Publish from block h+1.                                                  |
| `WithLogger`          | logger in `ctx` | Where the service logs.                                                  |

## Files

| File         | Contents                                                |
|--------------|---------------------------------------------------------|
| `publish.go` | `Service`, `NewEvents`, lifecycle and run loop          |
| `cursor.go`  | In-order cursor                                         |
| `ws.go`      | Websocket header stream: dial, subscribe, ping and pong |
| `node.go`    | Node URI parsing; HTTP and websocket targets            |
| `decode.go`  | Block results to typed events                           |
| `options.go` | Options                                                 |
| `query.go`   | Header subscription query                               |
