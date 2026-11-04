#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
mkdir -p .cache/sink-rust-check .cache/sink-rust-deps
sed '/^inputs:/,$d; s@out: rs/src/gen/node@out: .cache/sink-rust-deps@' buf.gen.rs.yaml > .cache/buf.sink.rs-deps.yaml
.cache/bin/buf generate --template .cache/buf.sink.rs-deps.yaml \
  --path go/vendor/github.com/cosmos/cosmos-sdk/proto/cosmos/base/v1beta1/coin.proto \
  --path go/vendor/github.com/cosmos/cosmos-sdk/proto/cosmos/base/query/v1beta1/pagination.proto
cat > .cache/sink-rust-check/Cargo.toml <<'TOML'
[package]
name = "sink-sdk-contract-check"
version = "0.0.0"
edition = "2021"
[workspace]
[lib]
path = "../../rs/tests/sink_vectors.rs"
[dependencies]
prost = "=0.13.5"
prost-types = "=0.13.5"
tonic = "=0.12.3"
serde = { version = "1", features = ["derive"] }
serde_json = "1"
pbjson = "=0.7.0"
pbjson-types = "=0.7.0"
sha2 = "0.10"
hex = "0.4"
TOML
cargo test --manifest-path .cache/sink-rust-check/Cargo.toml
