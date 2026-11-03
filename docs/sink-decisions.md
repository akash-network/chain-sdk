# Sink SDK contract

The unreleased `akash.sink.v1` schema uses genesis `schema_version: 4`. This is a coordinated fresh-chain migration. Old sink genesis, numeric IDs, record keys, programs, and client history do not migrate in place. Existing deployments remain separate. The node consensus version is 2.

## Identity and proof bytes

`ID` contains a canonical nonzero 20-byte Akash owner, positive `int64` height, and positive `uint64` sequence. Its text form is `owner/height/sequence`, with canonical decimal numbers. The node allocates sequences globally within a block, starting at 1. Failed deposits roll back allocation.

`IDKey` encodes `0x14 || owner20 || BE64(height) || BE64(sequence)`. The native transfer key is `0x0200 || IDKey` (39 bytes). The canonical settlement key is `0x0301 || status || IDKey` (40 bytes). Both values use `codec.CollValue`: native protobuf `RecordData` with destination and SDK Coin, and native protobuf `SettlementData`. There is no separate void map. Settlement tag 9 stores the permanent void authorization; status appears only in the canonical key.

Solana's permanent transfer status PDA uses `[transfer, owner20, LE64(height), LE64(sequence)]`. Proof-buffer seeds use `[proof, submitter32, SHA256(actual_proven_key)]`, where submitter is the buffer creator. The nonce is the full hash, and mint and void proofs have different nonces. Resolved status accounts remain open for replay protection.

The native fixtures are [`proof_vectors.json`](../go/node/sink/v1/testdata/proof_vectors.json) and [`settlement_vectors.json`](../go/node/sink/v1/testdata/settlement_vectors.json). They cover boundary IDs, base58 recipient lengths, amount-length boundaries, all five statuses, and a 6,500-byte legitimate maximum settlement certificate plus opposite witness. Node generates them through its production key codecs and native Cosmos value codecs: run `UPDATE_PROOF_VECTORS=1 go test ./x/sink/keeper -run '^(TestProofVectors|TestSettlementProofVectors)$' -count=1` from the node checkout, then copy `contracts/solana/programs/akt-solana/tests/fixtures/{proof_vectors,settlement_vectors}.json` into this SDK fixture directory. TypeScript and Rust independently encode the generated protobuf types and verify full-key nonce hashes.

## Financial status and authorization

`TransferState` contains ID, immutable data, settlement, and top-level status. `SettlementID` is the canonical settlement key. The five financial populations are pending 1, minted 2, voided 3, refund-pending 4, and refunded 5. `VOIDED` means a governance-ordered attempt awaiting Solana confirmation. The derived HTTP receipt label is `void_requested`. It does not authorize a refund. Holds are independent of status, and terminal conflicts remain discoverable. Escrow backs the first four populations; refunded amounts are outside escrow.

`Totals` and flat `EventTotalsUpdated` pair pending amount/count at tags 1/2, minted 3/4, voided 5/6, refund-pending 7/8, and refunded 9/10. Amounts are SDK Coins and counts are uint64. Only the event adds normalized `refunded_in_epoch` at tag 11. Stored totals contain no allocator, policy, or epoch facts.

`SettlementData` restores hold 1, round 2, embedded policy 3, votes 4, integer challenge deadline 5, flat governance observation 6, veto signer 7 and reason 8. Its permanent void authorization is tag 9; veto membership provenance is tag 10. Observation fields are outcome, slot and signature. Votes retain their membership version. Active attempts accept one vote per reporter for the first outcome, freeze the certificate at quorum, and retain at most one opposite witness. Conflict freezes further reports. Terminal and held evidence survives rotation.

Policies contain at most 32 canonical ordinary account addresses. `Params.reporters_threshold` is an SDK `LegacyDec` between zero and one. The effective quorum is `ceil(r*n)` and must exceed two thirds of the members. For example, a ratio of `0.6` with four reporters requires three signatures. Policy versions change only when membership or effective quorum changes. Member order alone does not change the rule.

A current report refreshes an unheld unpaid reporter attempt after rotation: replace its embedded policy, clear stale votes and the deadline, and return to pending or voided according to its permanent receipt. Old-policy reports fail before clearing. The round does not change. Governance-authorized refunds, terminal certificates and supported holds retain their evidence and policy version.

`MsgRequestVoid.expected_round` occupies tag 3. New orders require the current round; an authenticated retry is a no-op. The permanent receipt retains its original positive height and round through mint, conflict, restart, rotation, refund and recovery. A reporter veto retains a positive membership version; zero means governance and must match the configured authority at the node boundary.

`MsgRestartTransfer`, `MsgConfirmMint`, and `MsgConfirmVoid` replace `MsgResolveTransfer`. Each carries `expected_round`; a successful resolution increments the round. Confirmations carry slot/signature evidence without a second outcome field. Restart is unpaid-only. ConfirmMint cannot reverse a paid refund. ConfirmVoid requires a void order for unpaid or conflict-held minted state, but can resolve an already-paid dispute without a new order. Governance confirmation clears votes, vetoes, and holds.

Refund queue scope is zero for governance and the policy version for quorum. Only unheld refund-pending transfers belong. The node compares scope zero with the current policy's head. Queue presence does not imply maturity, available capacity, or enabled refunds. The original owner receives at most one payment.

## Amount limits and recipient validation

Amounts are SDK `Coin` values in `sdkutil.DenomUakt`. The account epoch cap may be zero to disable that limit; usage is still tracked. The minimum must be positive while deposits are enabled and may be zero while disabled. The other caps are positive. The minimum fits the transfer maximum, the aggregate sending cap, and any enabled account cap. The transfer maximum fits `uint64` and the refund cap. Account and aggregate sending caps may be below that maximum; every deposit must fit all applicable limits. Totals use checked SDK integer arithmetic and may exceed `uint64`.

Sending and refund usage share the current number of the existing native `day` epoch, with separate consumed amounts. The epochs module owns the schedule. Its duration and start are configurable in genesis, so the identifier does not imply a UTC-day boundary. A started epoch advances only when block time is strictly after its end, at most once per block. After downtime, successive blocks can each advance one epoch and receive one fresh allowance. Unused allowances do not accumulate.

Refunds do not restore sending capacity. Queries normalize stale counters on copies at the queried height; export retains raw stamps. Missing or malformed native epoch state fails explicitly. A valid unstarted schedule reports zero, but deposits and refund payments wait until counting starts. Future stamps and positive consumption at epoch zero fail. Lowered caps retain consumed capacity, and the node scans unpaid records before lowering the refund cap. The account cap's zero value disables only that comparison and does not reset account or aggregate consumption. Account count quotas are removed.

Artur's `reporter_challenge_period` remains an independent Duration measured from authorization. `challenge_end` is the checked ceiling of block time plus Duration in Unix seconds, including fractional carry; it is positive only for refund-pending. Crossing a quota boundary never shortens that interval. Existing challenge deadlines do not change with parameters. Immediate Solana mints and the separate Solana amount caps are unchanged.

Defaults disable deposits and refunds. The minimum is 1,000,000 uakt. Maximum transfer, account sending cap, and refund cap are 1,000,000,000,000 uakt. The aggregate sending cap is 10,000,000,000,000 uakt. The challenge period is 72 hours and the threshold ratio is 0.75. Reporters, vaults, and guardian start empty. Release builds retain the 72-hour floor; `sinkdev` permits a one-second floor. These are development defaults.

Both builds cap the challenge period at 9,223,372,035.999999999 seconds. `MaxChallengePeriod` derives this ceiling from `cosmossdk.io/x/tx/signing/aminojson.MaxDurationSeconds`, including its final fractional second. The pinned v0.14.0 Amino encoder rejects the next nanosecond, even though Go can represent it. Parameter, message, and genesis validation reject durations above this signing limit. This is an encoder constraint, not a product policy cap.

`MsgSink` carries `dest_type` and `dest`. Only `DEST_TYPE_SOLANA` validates, and `dest` is the Solana public key that `RecordData.dest` records. `SolanaAddress` parses a canonical nonzero 32-byte base58 address. Stateless message validation checks this format. The keeper applies the curve check, with listed vaults exempt from that check only. Solana evidence validation checks positive slot and a nonzero 64-byte signature; it does not establish finality or truth. Solana keeps independent transfer and daily mint amount caps, with immediate repeatable submissions and retryable cap failures.

## Recoverable genesis

Genesis contains params, totals, raw send/refund usage, `GenesisRecord` entries, one account usage entry per owner, and explicit `current_policy_version`. Current and all referenced policies must be present and consistent; sparse versions are valid. Extra retained history is allowed. Validation rejects old schema versions, duplicate IDs, reused height/sequence pairs across owners, duplicate owners/policies, missing authorizations, malformed evidence, invalid Coins, inconsistent deadlines, and accounting mismatches. Every terminal or refund-pending transfer retains an appropriate quorum or governance authorization. Empty zero-usage owners remain valid.

Global send consumption equals the sum of account epoch amounts stamped with its epoch. Account lifetime totals equal all their transfer amounts. Refund consumption cannot exceed total paid refunds. Lowered limits do not invalidate already-consumed capacity.

Schema marker 4 rejects schema 3 and older snapshots before interpretation. It covers the native canonical key/value shape, restored genesis records, permanent authorization, policy provenance and renumbered totals. Old fixed-day counters also remain incompatible with native epoch numbers. The install-only upgrade rejects a version map containing Sink. There is no supported in-place conversion or vote truncation.

Recovery preserves the exported `day` record before Sink initialization, including its duration, start, current number, current start, and counting flag. Empty zero-stamped usage is valid when installing into an existing nonzero epoch. Neither import nor queries advance the timer. Raw consumption must not exceed the native number.

Node import additionally requires both ID and void-request heights below `max(1, initial_height)`, rejects future epochs and existing sink state, rebuilds derived indexes, and checks exact bank backing. Operators retain an `initial_height` greater than the exported chain height. Import does not refund, sweep, rotate policies, or normalize raw counters.

## Queries and CLI

The only point lookup is `Transfer`, at `/akash/sink/v1/transfers/{id.owner}/{id.height}/{id.sequence}`. `Transfers` accepts optional owner and financial status filters. Unfiltered history sorts by height, sequence, and owner. `Disputes` includes terminal conflicts. `Account` returns normalized owner usage. `State` includes params, totals, current policy, independent normalized usage, and the eligible-scope refund head.

Lists use native `PageRequest`/`PageResponse`. Default limit is 100; maximum is 500. Offset and count-total scans are rejected. Status queries range canonical settlement buckets. Owner-only queries merge all five buckets in chronological order. The node rejects malformed, absent, and out-of-filter cursors; `next_key` is inclusive and contains the raw selected key without the collection prefix. It carries no server-enforced height binding. Keep `x-cosmos-block-height` fixed for every page; a pruned-height error ends that scan. Durable discovery retains the last processed chronological identity even when `next_key` is empty. This prevents missing a newly active owner that sorts before earlier owners.

The CLI exposes `transfer`, `transfers`, `disputes`, `account`, and `state`, plus all ten transaction methods. `transfers` uses `--owner`, `--status`, `--page-key`, `--limit`, and `--reverse`. Governance commands take canonical positive `expected-round` arguments; confirmations take a slot and base58 transaction signature. `--authority` defaults to the governance module account. Proposal preparation uses `--generate-only`. `deposit` accepts exact AKT decimals or integer uakt and always sets `dest_type` to `DEST_TYPE_SOLANA`.

## Generation and verification

Legacy and native Amino signing omit empty scalar fields and both nil and empty reporter/vault lists. These lists have one canonical absent representation. The unreleased schema selects this with `omitempty` Go JSON tags. `sdkutil.MakeEncodingConfig` gives the resolved protobuf descriptor graph only to its Amino JSON sign-mode handler, so nested IDs, parameters, observations, and evidence remain in the signed bytes. Signer lookups keep the global descriptors shared by independent transaction configurations. Direct and direct-aux modes retain their existing encoding. Use `TxConfig.SignModeHandler()` for Amino signing; constructing another encoder from `SigningOptions` or calling `ProtoCodec.MarshalAminoJSON` does not use this handler.

The shared handler also changes Amino sign bytes for existing non-Sink messages with unresolved nested fields. Ship it with the node in a coordinated upgrade and update Amino-signing clients together. For wasm instantiate permissions, the handler uses numeric enum values. Existing cosmjs Amino converters for `MsgStoreCode.instantiate_permission` use names and remain incompatible; those clients should use direct signing. This does not affect Sink's measured Amino parity.

`script/sink-protocgen.sh` uses the repository's Buf configs and pinned local Go generators. It generates only Sink Go/Rust messages and mocks. The TypeScript node descriptor/index pass includes all node files so the aggregate SDK remains complete. It also regenerates API reference documentation. Generator prerequisites are `buf`, `protoc-gen-gocosmos`, `protoc-gen-grpc-gateway`, `mockery`, and `protoc-gen-doc` in `.cache/bin`, plus installed TypeScript dependencies and vendored proto imports.

Run the SDK gate from `go` with `GOTOOLCHAIN=go1.25.0 go test -mod=readonly ./node/sink/v1`, then repeat with `-tags sinkdev`. Run CLI checks from `go/cli` with `GOTOOLCHAIN=go1.25.1 go test -mod=readonly . -run TestSink`. TypeScript uses `npm run build` and `npm run test:unit`. `script/test-sink-rust.sh` compiles the generated Sink clients and verifies proof fixtures in an isolated crate because the root workspace's `rust` member is absent from this checkout.
