# Vendored protobuf inputs

This directory contains the minimal imported source set needed to lint and
generate the Akash protobufs without fetching dependency modules from the Buf
Schema Registry.

The files are copied without modification from these locked modules. The
parenthetical license is the upstream repository license; individual files
retain their source headers when present.

- `buf.build/cosmos/ics23` commit `dc427cb4519143d8996361c045a29ad7`
  (Apache-2.0)
- `buf.build/googleapis/googleapis` commit
  `004180b77378443887d3b55cabc00384` (Apache-2.0)
- `buf.build/protocolbuffers/wellknowntypes` commit
  `4e1ccfa6827947beb55974645a315b8d` (`protocolbuffers/protobuf` v33.2,
  BSD-3-Clause)

Update these inputs and the generated bindings together.
