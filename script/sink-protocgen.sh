#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
export PATH="$PWD/.cache/bin:$PWD/ts/script:$PATH"
mkdir -p .cache/sink-gen-go .cache/sink-gen-rust
sed 's@out: .@out: .cache/sink-gen-go@' buf.gen.gogo.yaml > .cache/buf.sink.go.yaml
buf generate --template .cache/buf.sink.go.yaml --path proto/node/akash/sink/v1
cp .cache/sink-gen-go/pkg.akt.dev/go/node/sink/v1/*.go go/node/sink/v1/
PROTO_SOURCE=node buf generate --template buf.gen.ts.yaml ./proto/node
node --experimental-strip-types --no-warnings ts/script/fix-ts-proto-generated-types.ts
sed '/^inputs:/,$d; s@out: rs/src/gen/node@out: .cache/sink-gen-rust@' buf.gen.rs.yaml > .cache/buf.sink.rs.yaml
buf generate --template .cache/buf.sink.rs.yaml --path proto/node/akash/sink/v1
cp .cache/sink-gen-rust/akash.sink.v1*.rs rs/src/gen/node/
python3 rs/script/fix-prost-serde.py rs/src/gen/node
rustfmt --edition 2021 --config skip_children=true rs/src/gen/node/akash.sink.v1.rs rs/src/gen/node/akash.sink.v1.serde.rs rs/src/gen/node/akash.sink.v1.tonic.rs
cat > .cache/mockery.sink.yaml <<'YAML'
mockname: "{{.InterfaceName}}"
filename: "{{.MockName}}_mock.go"
with-expecter: True
resolve-type-alias: False
issue-845-fix: True
packages:
  pkg.akt.dev/go/node/sink/v1:
    config:
      dir: mocks/node/client/sink
    interfaces:
      QueryClient: {}
YAML
(cd go && GOFLAGS=-mod=readonly GOTOOLCHAIN=go1.25.0 ../.cache/bin/mockery --config ../.cache/mockery.sink.yaml)
buf generate --template proto/node/buf.gen.doc.yaml ./proto/node
python3 - <<'PY'
from pathlib import Path
path = Path('docs/proto/node.md')
lines = path.read_text().splitlines(keepends=True)
in_sink = False
for index, line in enumerate(lines):
    if '<a name="akash/' in line and '.proto"></a>' in line:
        in_sink = '<a name="akash/sink/v1/' in line
    if in_sink:
        lines[index] = line.rstrip() + '\n'
path.write_text(''.join(lines))
PY
