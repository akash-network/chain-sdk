# akash.sink.v1 integration decisions

Status as of 2026-10-19, for maintainers and agents preparing the sink API for
a release. This file is its own commit; drop it before opening an upstream
pull request if it is not wanted there. Do not act on an Open item without its
owner; record each decision here with its date.

## What this branch contains

Based on `main` at `141c80c3`:

1. `feat(sink): add akash.sink.v1 protobuf api`.
2. `feat(sink): generate go and ts bindings for sink api`, produced by
   `make proto-gen-go`, `proto-gen-ts` and `proto-gen-doc`.
3. `feat(sink): add go types, validation and codecs` (`ValidateBasic`,
   parameter validation).
4. `feat(sink): expose sink queries in the node client`.
5. `feat(cli): add sink tx and query commands` (deposit, guardian-pause,
   record, open-transfers).
6. `test(ts): cover sink sdk messages and queries`.

The node module lives in node's matching `feat/sink` branch, which needs a
tagged release of this repository.

## Compatibility constraints

A Solana program verifies x/sink state with ICS23 proofs, and deployed
clients sign these messages. Coordinate before changing:

- Protobuf package `akash.sink.v1`, message names and field numbers.
- Amino names `sink/MsgSink` and `sink/MsgGuardianPause` (wallet sign bytes).
- uint64 fields must round-trip values above 2^63 exactly in TypeScript; the
  bigint SDK does, and a spec covers it.

## Decisions needed

| ID  | Decision       | Context                                                                                                                       | Recommendation                                                 | Owner                 | Status |
| --- | -------------- | ----------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------- | --------------------- | ------ |
| A3  | Release tag    | node's sink branch pins a placeholder `go.mod` until `pkg.akt.dev/go` and `pkg.akt.dev/go/cli` are tagged with this API        | Tag both Go modules, then update node                          | chain-sdk maintainers | Open   |
| A4  | Amino names    | `sink/...` differs from bme's `akash/bme/v1/...` style; changing it changes wallet sign bytes                                  | Decide before any wallet integration ships                     | Akash maintainers     | Open   |
| A6  | Rust bindings  | Upstream generates them with remote Buf plugins, which upload the schema; local plugins don't reproduce upstream's output       | Regenerate once the schema is public                           | chain-sdk maintainers | Open   |
| A7  | TS consumers   | The existing Solana-side tooling was built against a `Long`-based SDK; this branch uses upstream's `bigint`                    | Consumers migrate when they adopt this release                 | Consumers             | Open   |

## How to test

`make test` for the Go modules (Go 1.25.1, as CI uses), the TypeScript tests,
`make lint-ts`, `make lint-proto` and the generated-code check all passed on
2026-10-19.
