#!/bin/bash
#
# Enforces: if/then/else keywords must only appear inside top-level definitions,
# never inline under properties.
#
# Usage: ./lint-schema.sh [schema-file]
# Default schema file: sdl-input.schema.yaml in the same directory as this script.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCHEMA="${1:-${SCRIPT_DIR}/sdl-input.schema.yaml}"

if [[ ! -f "$SCHEMA" ]]; then
    echo "error: schema file not found: $SCHEMA" >&2
    exit 1
fi

echo "Checking SDL schema: if/then/else must only appear inside top-level definitions"
echo "  file: ${SCHEMA}"

VIOLATIONS=()
in_definitions=false
line_num=0

while IFS= read -r line; do
    line_num=$((line_num + 1))

    # Track whether we're inside top-level definitions block or properties block.
    # Top-level keys have zero indentation.
    if [[ "$line" =~ ^definitions: ]]; then
        in_definitions=true
        continue
    fi

    if [[ "$line" =~ ^[a-zA-Z] && ! "$line" =~ ^definitions: ]]; then
        in_definitions=false
    fi

    if $in_definitions; then
        continue
    fi

    # Check for if/then/else as YAML keys (with proper indentation).
    # Match lines like "    if:", "      then:", "      else:" but not
    # things like "signedBy:" or property values containing these words.
    if [[ "$line" =~ ^[[:space:]]+(if|then|else): ]]; then
        trimmed="${line#"${line%%[![:space:]]*}"}"
        VIOLATIONS+=("${line_num}: ${trimmed}")
    fi
done < "$SCHEMA"

if [[ ${#VIOLATIONS[@]} -gt 0 ]]; then
    echo "FAIL"
    echo "" >&2
    echo "error: if/then/else keywords found outside definitions block" >&2
    echo "Convention: extract cross-field validations into named definitions and use allOf + \$ref" >&2
    echo "" >&2
    for v in "${VIOLATIONS[@]}"; do
        echo "  line ${v}" >&2
    done
    exit 1
fi

echo "OK"
