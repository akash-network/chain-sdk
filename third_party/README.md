# Vendored protobuf inputs

This directory contains the minimal imported source set needed to lint and
generate the Akash protobufs without fetching dependency modules from the Buf
Schema Registry.

The files are copied without modification from these locked Apache-2.0 modules:

- `buf.build/cosmos/ics23` commit `dc427cb4519143d8996361c045a29ad7`
- `buf.build/googleapis/googleapis` commit `004180b77378443887d3b55cabc00384`
- `buf.build/protocolbuffers/wellknowntypes` commit `4e1ccfa6827947beb55974645a315b8d`

Update these inputs and the generated bindings together.
