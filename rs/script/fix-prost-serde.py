#!/usr/bin/env python3
import argparse
from pathlib import Path
import re


MARKER = "\n// Generated from rs/serde/sink_legacy_dec.rs by rs/script/fix-prost-serde.py.\n"


def patch(source: str, adapter: str) -> str:
    source = source.split(MARKER)[0]
    replacements = [
        (
            r'serialize_field\(\s*"reportersThreshold",\s*&self\.reporters_threshold\s*\)',
            'serialize_field("reportersThreshold", '
            '&sink_legacy_dec::to_json(&self.reporters_threshold)'
            '.map_err(serde::ser::Error::custom)?)',
            "sink_legacy_dec::to_json",
        ),
        (
            r"reporters_threshold__ = Some\(map_\.next_value\(\)\?\);",
            "reporters_threshold__ = Some("
            "sink_legacy_dec::from_json(&map_.next_value::<String>()?)"
            ".map_err(serde::de::Error::custom)?);",
            "sink_legacy_dec::from_json",
        ),
    ]
    for pattern, replacement, call in replacements:
        source, count = re.subn(pattern, replacement, source)
        if count > 1 or source.count(call) != 1:
            raise ValueError(f"unexpected Sink serde output for {call}")
    return source.rstrip() + "\n" + MARKER + adapter


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("output_dir", type=Path)
    args = parser.parse_args()
    path = args.output_dir / "akash.sink.v1.serde.rs"
    adapter = (Path(__file__).resolve().parents[1] / "serde/sink_legacy_dec.rs").read_text()
    path.write_text(patch(path.read_text(), adapter))


if __name__ == "__main__":
    main()
