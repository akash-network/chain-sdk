# AEP-86 Auditor

`aep86-auditor` is the local starting point for the AEP-86 L-5 reference auditor CLI.

This tool intentionally lives outside consensus code. The current slice implements local collection, offline evidence checks, sustained baseline comparison, and transaction command preparation:

- generates a cryptographically random 32-byte nonce
- calls provider `akash.inventory.v1.InventoryService/GetInventorySnapshot`
- decodes `akash.inventory.v1.SnapshotPayload`
- verifies nonce/provider/chain binding
- queries the chain gRPC auth account for the provider public key
- verifies the provider signature over the raw `snapshot_payload` bytes
- writes raw artifacts plus a draft `akash.audit.evidence.v1` JSON document
- verifies that `evidence.draft.json` is schema-valid canonical JSON and matches `evidence.draft.sha256`
- compares a current evidence artifact against a baseline artifact
- prepares submit and revoke transaction commands from verified evidence artifacts

## Run

```sh
go run ./go/tools/aep86-auditor collect \
  --provider-grpc provider.example.com:8443 \
  --chain-grpc rpc.example.com:9090 \
  --auditor akash1... \
  --audit-escrow-id 0 \
  --target-tier L1 \
  --software-binary-hash sha256:<64-hex> \
  --output-dir ./aep86-audit
```

The provider endpoint is the existing provider daemon gRPC endpoint with the public AEP-86 inventory service registered.
The chain endpoint is an Akash node gRPC endpoint used to query the provider account public key and best-effort
verification facts.
`--software-binary-hash` is required and must use `sha256:<64-hex>` form so the draft evidence satisfies the
strict evidence schema.

For local devnets with self-signed provider certificates, add `--provider-skip-tls-verify`. For plaintext test servers,
add `--provider-insecure`.

Validate the local artifact directory before submitting evidence elsewhere:

```sh
go run ./go/tools/aep86-auditor verify ./aep86-audit
```

Prepare a submission command from a verified artifact directory:

```sh
go run ./go/tools/aep86-auditor submit \
  --fee 100uakt \
  --deposit 200uakt \
  ./aep86-audit
```

Compare sustained-validation evidence against the original baseline:

```sh
go run ./go/tools/aep86-auditor sustain ./aep86-baseline ./aep86-current
```

Write a sustained-validation evidence artifact that can drive revocation command
preparation:

```sh
go run ./go/tools/aep86-auditor sustain \
  --output-dir ./aep86-sustained \
  ./aep86-baseline \
  ./aep86-current
```

Prepare a revocation command from verified revocation evidence:

```sh
go run ./go/tools/aep86-auditor revoke \
  --reason software_identity_changed \
  ./aep86-sustained
```

`submit` and `revoke` are dry-run helpers. They validate canonical `evidence.draft.json`,
verify `evidence.draft.sha256`, check optional provider/auditor/audit-escrow/tier/capability/chain
flags against the evidence, and print the exact `akash tx verification ...` command.
