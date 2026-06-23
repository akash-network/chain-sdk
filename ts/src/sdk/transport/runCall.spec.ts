import { getEventListeners } from "node:events";
import * as v8 from "node:v8";
import * as vm from "node:vm";

import { describe, expect, it } from "vitest";

import { runStreamingCall, runUnaryCall } from "./runCall.ts";
import { TransportError } from "./TransportError.ts";

type StreamOptions = Parameters<typeof runStreamingCall>[0];
type StreamReq = StreamOptions["req"];
type UnaryOptions = Parameters<typeof runUnaryCall>[0];

// runStreamingCall only ever touches `method.input.fromPartial`, so a thin stub
// is enough to drive it without pulling in real generated message descriptors.
const method = {
  input: { fromPartial: (m: unknown) => m },
} as unknown as StreamReq["method"];

describe("runUnaryCall", () => {
  it("registers no abort listener and resolves with the response", async () => {
    let signal: AbortSignal | undefined;
    const response = { stream: false, message: {}, header: new Headers(), trailer: new Headers(), method };
    const next = (async (req: { signal: AbortSignal }) => {
      signal = req.signal;
      return response;
    }) as unknown as UnaryOptions["next"];

    const result = await runUnaryCall({
      req: { message: {}, method, service: {}, requestMethod: "POST", url: "https://example.test/unary" },
      next,
    } as unknown as UnaryOptions);

    expect(result).toBe(response);
    expect(getEventListeners(signal!, "abort")).toHaveLength(0);
  });

  it("maps a rejection to a TransportError", async () => {
    const next = (async () => {
      throw new Error("boom");
    }) as unknown as UnaryOptions["next"];

    await expect(
      runUnaryCall({
        req: { message: {}, method, service: {}, requestMethod: "POST", url: "https://example.test/unary" },
        next,
      } as unknown as UnaryOptions),
    ).rejects.toBeInstanceOf(TransportError);
  });
});

// Deterministic, GC-free guard. The composite AbortSignal.any() signal is stored
// on req.signal; after the stream settles there must be no "abort" listeners left
// on it - that listener is exactly what pins the request graph in Node's
// `gcPersistentSignals` set. The first case is the load-bearing one: a normal
// `for await` completion never triggers IteratorClose, so it is the path a partial
// fix (cleanup only in return()/throw()) misses.
describe("runStreamingCall", () => {
  const gc = getGc();

  it("removes the abort listener after the stream completes normally", async () => {
    expect(getEventListeners(await drive({ responses: 3 }), "abort")).toHaveLength(0);
  });

  it("removes the abort listener when the consumer abandons the stream early", async () => {
    expect(getEventListeners(await drive({ responses: 5, breakAfter: 1 }), "abort")).toHaveLength(0);
  });

  // Stronger guard: prove the request graph is actually collectable after a normal
  // completion, with a long-lived caller signal held for the whole test. This pins
  // the cause on Node's `gcPersistentSignals` set (released only when the listener
  // is removed) rather than on the caller signal.
  it.skipIf(gc === undefined)("lets the request graph be garbage collected after completion", async () => {
    const callerSignal = new AbortController().signal;

    // Build, run and fully drain the call in a nested scope so the response wrapper
    // and the request message both fall out of scope on return. The only thing that
    // could keep the request message alive afterwards is a lingering abort listener.
    const collectRef = async (): Promise<WeakRef<object>> => {
      const requestMessage = {
        async *[Symbol.asyncIterator]() {
          yield {};
        },
      };
      const ref = new WeakRef(requestMessage);
      const res = await runStreamingCall({
        interceptors: undefined,
        signal: callerSignal,
        req: {
          stream: true,
          message: requestMessage,
          method: { input: { fromPartial: (m: unknown) => m } },
        },
        next: async () => ({
          message: {
            async *[Symbol.asyncIterator]() {
              yield { i: 0 };
            },
          },
        }),
      } as unknown as Parameters<typeof runStreamingCall>[0]);

      for await (const _ of res.message as AsyncIterable<unknown>) {
        // drain to completion
      }
      return ref;
    };

    const ref = await collectRef();
    await forceCollection(gc!);

    expect(ref.deref()).toBeUndefined();
    // Keep the caller signal reachable past the GC to make the point explicit.
    expect(callerSignal.aborted).toBe(false);
  });

  async function drive(opts: { responses: number; breakAfter?: number }): Promise<AbortSignal> {
    let composite!: AbortSignal;
    const res = await runStreamingCall({
      interceptors: undefined,
      signal: new AbortController().signal, // a caller signal, kept alive for the test
      req: {
        stream: true,
        message: { async *[Symbol.asyncIterator]() { yield {}; } },
        method: { input: { fromPartial: (m: unknown) => m } },
      },
      next: async (req: { signal: AbortSignal }) => {
        composite = req.signal; // the AbortSignal.any() composite
        return {
          message: {
            async *[Symbol.asyncIterator]() {
              for (let i = 0; i < opts.responses; i++) yield { i };
            },
          },
        };
      },
    } as unknown as Parameters<typeof runStreamingCall>[0]);

    let n = 0;
    for await (const _ of res.message as AsyncIterable<unknown>) {
      if (opts.breakAfter !== undefined && ++n >= opts.breakAfter) break;
    }
    return composite;
  }
});

/**
 * Returns a callable `gc()` if one is reachable. Prefers the global exposed by
 * `--expose-gc`; otherwise enables it at runtime via the v8 module so the
 * WeakRef regression test can run without special node flags. Returns undefined
 * if neither is available, in which case the test is skipped.
 */
function getGc(): (() => void) | undefined {
  const globalGc = (globalThis as { gc?: () => void }).gc;
  if (typeof globalGc === "function") return globalGc;
  try {
    v8.setFlagsFromString("--expose-gc");
    const gc = vm.runInNewContext("gc") as unknown;
    v8.setFlagsFromString("--no-expose-gc");
    if (typeof gc === "function") return gc as () => void;
  } catch {
    // Unable to obtain a gc handle; the regression test will be skipped.
  }
  return undefined;
}

async function forceCollection(gc: () => void): Promise<void> {
  for (let i = 0; i < 10; i++) {
    gc();
    await new Promise((resolve) => setImmediate(resolve));
  }
}
