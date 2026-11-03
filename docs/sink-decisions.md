# Sink SDK contract

The unreleased `akash.sink.v1` schema uses genesis `schema_version: 3`. This is a coordinated fresh-chain migration. Old sink genesis, numeric IDs, record keys, programs, and client history do not migrate in place. Existing deployments remain separate. The node consensus version is 2.

## Identity and proof bytes

`ID` contains a canonical nonzero 20-byte Akash owner, positive `int64` height, and positive `uint64` sequence. Its text form is `owner/height/sequence`, with canonical decimal numbers. The node allocates sequences globally within a block, starting at 1. Failed deposits roll back allocation.

`IDKey` encodes `0x14 || owner20 || BE64(height) || BE64(sequence)`. The transfer key is `0x1200 || IDKey`, and the void-order key is `0x1201 || IDKey`. Each full key is 39 bytes. `RecordDataValue` is exactly `0x02 || recipient32 || LE64(amount)`, 41 bytes. `VoidOrderValue` is the single byte `0x02`. Protobuf does not encode these proven values.

Solana's permanent transfer status PDA uses `[transfer, owner20, LE64(height), LE64(sequence)]`. Proof-buffer seeds use `[proof, submitter32, SHA256(actual_proven_key)]`, where submitter is the buffer creator. The nonce is the full hash, and mint and void proofs have different nonces. Resolved status accounts remain open for replay protection.

The shared fixture is [`proof_vectors.json`](../go/node/sink/v1/testdata/proof_vectors.json). It covers minimum, ordinary, and maximum IDs and amounts. Go's `TestProofVectors` generates it with `UPDATE_PROOF_VECTORS=1`. TypeScript and Rust independently check these bytes and nonce hashes.

## Financial status and authorization

`TransferState` contains the ID, immutable `RecordData`, complete `SettlementData`, and `void_ordered`. `SettlementID` (status, owner, height, sequence) keys the node's status index. The four financial populations are `STATUS_TRANSFER_PENDING`, `STATUS_MINTED`, `STATUS_REFUND_PENDING`, and `STATUS_REFUNDED`. Value 3 and `STATUS_VOIDED` are reserved. The four Coin/count pairs in `Totals` count exactly these populations, including held transfers.

A void order does not authorize payment. A void observation can establish refund authorization. Refund-pending includes its challenge interval. Refunded means the bank payment occurred. A conflict can overlay any status; a veto holds unpaid state. Module escrow backs pending, minted, and refund-pending amounts. Refunded totals are outside escrow.

Votes retain their own `policy_version`. Dedupe uses reporter, outcome, and policy version, ignoring changed evidence for the same tuple. Opposing observations remain available for conflict detection. Quorum tallies use the settlement's authorization version. Minted and refunded state retain their original quorum or governance evidence after reporter rotation.

Policies contain at most 32 canonical ordinary account addresses. `Params.reporters_threshold` is an SDK `LegacyDec`. With configured reporters, `3*r > 2` and `r <= 1`; the effective quorum is `ceil(r*n)`. Policy versions change only when membership or effective quorum changes. Member order alone does not change the rule.

On refresh, transfer-pending and quorum-authorized refund-pending adopt the current policy. The latter returns to pending and loses its deadline. Prior votes and holds survive. Governance refund authorization and terminal authorization retain their evidence and policy version.

`MsgRestartTransfer`, `MsgConfirmMint`, and `MsgConfirmVoid` replace `MsgResolveTransfer`. Each carries `expected_round`; a successful resolution increments the round. Confirmations carry slot/signature evidence without a second outcome field. Restart is unpaid-only. ConfirmMint cannot reverse a paid refund. ConfirmVoid requires a void order for unpaid or conflict-held minted state, but can resolve an already-paid dispute without a new order. Governance confirmation clears votes, vetoes, and holds.

Refund queue scope is zero for governance and the policy version for quorum. Only unheld refund-pending transfers belong. The node compares scope zero with the current policy's head. Queue presence does not imply maturity, available capacity, or enabled refunds. The original owner receives at most one payment.

## Amount limits and recipient validation

Amounts are SDK `Coin` values in `sdkutil.DenomUakt`. The account epoch cap may be zero to disable that limit; usage is still tracked. The minimum and other caps are positive. The minimum is at most the maximum transfer, which fits `uint64`, the aggregate sending and refund caps, and any enabled account cap. Totals use checked SDK integer arithmetic and may exceed `uint64`.

Sending and refund usage share the current number of the existing native `day` epoch, with separate consumed amounts. The epochs module owns the schedule. Its duration and start are configurable in genesis, so the identifier does not imply a UTC-day boundary. A started epoch advances only when block time is strictly after its end, at most once per block. After downtime, successive blocks can each advance one epoch and receive one fresh allowance. Unused allowances do not accumulate.

Refunds do not restore sending capacity. Queries normalize stale counters on copies at the queried height; export retains raw stamps. Missing or malformed native epoch state fails explicitly. A valid unstarted schedule reports zero, but deposits and refund payments wait until counting starts. Future stamps and positive consumption at epoch zero fail. Lowered caps retain consumed capacity, and the node scans unpaid records before lowering the refund cap. The account cap's zero value disables only that comparison and does not reset account or aggregate consumption. Account count quotas are removed.

Artur's `reporter_challenge_period` remains an independent Duration measured from authorization. Crossing a quota boundary never shortens that interval. Existing challenge deadlines do not change with parameters. Immediate Solana mints and the separate Solana amount caps are unchanged.

Defaults disable deposits and refunds. The minimum is 1,000,000 uakt. Maximum transfer, account sending cap, and refund cap are 1,000,000,000,000 uakt. The aggregate sending cap is 10,000,000,000,000 uakt. The challenge period is 72 hours and the threshold ratio is 0.75. Reporters, vaults, and guardian start empty. Release builds retain the 72-hour floor; `sinkdev` permits a one-second floor. These are development defaults.

Both builds cap the challenge period at 9,223,372,035.999999999 seconds. `MaxChallengePeriod` derives this ceiling from `cosmossdk.io/x/tx/signing/aminojson.MaxDurationSeconds`, including its final fractional second. The pinned v0.14.0 Amino encoder rejects the next nanosecond, even though Go can represent it. Parameter, message, and genesis validation reject durations above this signing limit. This is an encoder constraint, not a product policy cap.

`MsgSink` carries `dest_type` and `dest`. Only `DEST_TYPE_SOLANA` validates, and `dest` is the Solana public key that `RecordData.dest` records. `SolanaAddress` parses a canonical nonzero 32-byte base58 address. Stateless message validation checks this format. The keeper applies the curve check, with listed vaults exempt from that check only. Solana evidence validation checks positive slot and a nonzero 64-byte signature; it does not establish finality or truth. Solana keeps independent transfer and daily mint amount caps, with immediate repeatable submissions and retryable cap failures.

## Recoverable genesis

Genesis contains params, totals, complete policy history, raw send/refund usage, every complete transfer, and each owner's usage exactly once. Validation rejects old schema versions, duplicate IDs, reused height/sequence pairs across owners, duplicate owners/policies, missing authorizations, malformed evidence, invalid Coins, inconsistent deadlines, and accounting mismatches. Every terminal or refund-pending transfer retains an appropriate quorum or governance authorization. Empty zero-usage owners remain valid.

Global send consumption equals the sum of account epoch amounts stamped with its epoch. Account lifetime totals equal all their transfer amounts. Refund consumption cannot exceed total paid refunds. Lowered limits do not invalidate already-consumed capacity.

Schema marker 3 distinguishes native quota numbers from marker 2's fixed-day counters. An old day stamp of 20759 with positive consumption is numerically older than native epoch 50000. Accepting it as a native stamp would clear consumption during normalization. Marker 2 therefore fails before interpretation. This guard does not change proof format, IDs, KV domains, or the module consensus version. The install-only upgrade still rejects a version map containing Sink. There is no supported in-place conversion.

Recovery preserves the exported `day` record before Sink initialization, including its duration, start, current number, current start, and counting flag. Empty zero-stamped usage is valid when installing into an existing nonzero epoch. Neither import nor queries advance the timer. Raw consumption must not exceed the native number.

Node import additionally requires `ID.height < max(1, initial_height)`, rejects future epochs and existing sink state, rebuilds derived indexes, and checks exact bank backing. Operators retain an `initial_height` greater than the exported chain height. Import does not refund, sweep, rotate policies, or normalize raw counters.

## Queries and CLI

The only point lookup is `Transfer`, at `/akash/sink/v1/transfers/{id.owner}/{id.height}/{id.sequence}`. `Transfers` accepts optional owner and financial status filters. Unfiltered history sorts by height, sequence, and owner. `Disputes` includes terminal conflicts. `Account` returns normalized owner usage. `State` includes params, totals, current policy, independent normalized usage, and the eligible-scope refund head.

Lists use native `PageRequest`/`PageResponse`. Default limit is 100; maximum is 500. Offset and count-total scans are rejected. The node validates cursor bytes against the selected index and filters. Keep `x-cosmos-block-height` fixed for every page; a pruned-height error ends that scan. Durable discovery retains the last processed chronological identity even when `next_key` is empty. This prevents missing a newly active owner that sorts before earlier owners.

The CLI exposes `transfer`, `transfers`, `disputes`, `account`, and `state`, plus all ten transaction methods. `transfers` uses `--owner`, `--status`, `--page-key`, `--limit`, and `--reverse`. Governance commands take canonical positive `expected-round` arguments; confirmations take a slot and base58 transaction signature. `--authority` defaults to the governance module account. Proposal preparation uses `--generate-only`. `deposit` accepts exact AKT decimals or integer uakt and always sets `dest_type` to `DEST_TYPE_SOLANA`.

## Generation and verification

Legacy and native Amino signing omit empty scalar fields and both nil and empty reporter/vault lists. These lists have one canonical absent representation. The unreleased schema selects this with `omitempty` Go JSON tags. `sdkutil.MakeEncodingConfig` gives the resolved protobuf descriptor graph only to its Amino JSON sign-mode handler, so nested IDs, parameters, observations, and evidence remain in the signed bytes. Signer lookups keep the global descriptors shared by independent transaction configurations. Direct and direct-aux modes retain their existing encoding. Use `TxConfig.SignModeHandler()` for Amino signing; constructing another encoder from `SigningOptions` or calling `ProtoCodec.MarshalAminoJSON` does not use this handler.

The shared handler also changes Amino sign bytes for existing non-Sink messages with unresolved nested fields. Ship it with the node in a coordinated upgrade and update Amino-signing clients together. For wasm instantiate permissions, the handler uses numeric enum values. Existing cosmjs Amino converters for `MsgStoreCode.instantiate_permission` use names and remain incompatible; those clients should use direct signing. This does not affect Sink's measured Amino parity.

`script/sink-protocgen.sh` uses the repository's Buf configs and pinned local Go generators. It generates only Sink Go/Rust messages and mocks. The TypeScript node descriptor/index pass includes all node files so the aggregate SDK remains complete. It also regenerates API reference documentation. Generator prerequisites are `buf`, `protoc-gen-gocosmos`, `protoc-gen-grpc-gateway`, `mockery`, and `protoc-gen-doc` in `.cache/bin`, plus installed TypeScript dependencies and vendored proto imports.

Run the SDK gate from `go` with `GOTOOLCHAIN=go1.25.0 go test -mod=readonly ./node/sink/v1`, then repeat with `-tags sinkdev`. Run CLI checks from `go/cli` with `GOTOOLCHAIN=go1.25.1 go test -mod=readonly . -run TestSink`. TypeScript uses `npm run build` and `npm run test:unit`. `script/test-sink-rust.sh` compiles the generated Sink clients and verifies proof fixtures in an isolated crate because the root workspace's `rust` member is absent from this checkout.
