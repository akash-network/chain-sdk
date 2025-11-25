#!/bin/bash
# Wrapper script to set up localStorage polyfill before running the TypeScript plugin

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
export NODE_OPTIONS="--require $SCRIPT_DIR/localStorage-polyfill.js"

exec node --experimental-strip-types --no-warnings "$SCRIPT_DIR/protoc-gen-sdk-object.ts" "$@"

