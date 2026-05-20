# AEP-86 Auditor

`aep86-auditor` is the local starting point for the AEP-86 L-5 reference auditor CLI.

This tool intentionally lives outside consensus code. The first slice implements `collect` only:

- generates a cryptographically random 32-byte nonce
- calls provider `akash.inventory.v1.InventoryService/GetInventorySnapshot`
- decodes `akash.inventory.v1.SnapshotPayload`
- verifies nonce/provider/chain binding
- queries the chain gRPC auth account for the provider public key
- verifies the provider signature over the raw `snapshot_payload` bytes
- writes raw artifacts plus a draft `akash.audit.evidence.v1` JSON document

## Run

```sh
go run ./tools/aep86-auditor collect \
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
