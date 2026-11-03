#!/usr/bin/env python3
import argparse
from pathlib import Path
import re


MARKER = "\n// Generated from rs/serde/sink_legacy_dec.rs by rs/script/fix-prost-serde.py.\n"


def patch(source: str, adapter: str) -> str:
    source = source.split(MARKER)[0]
    for field, json_name in [
        ("reporters_threshold", "reportersThreshold"),
        ("threshold", "threshold"),
    ]:
        replacements = [
            (
                rf'serialize_field\(\s*"{json_name}",\s*&self\.{field}\s*\)',
                f'serialize_field("{json_name}", '
                f'&sink_legacy_dec::to_json(&self.{field})'
                '.map_err(serde::ser::Error::custom)?)',
                rf"sink_legacy_dec::to_json\(&self\.{field}\)",
            ),
            (
                rf"\b{field}__ = Some\(map_\.next_value\(\)\?\);",
                f"{field}__ = Some("
                "sink_legacy_dec::from_json(&map_.next_value::<String>()?)"
                ".map_err(serde::de::Error::custom)?);",
                rf"\b{field}__\s*=\s*Some\(\s*sink_legacy_dec::from_json",
            ),
        ]
        for pattern, replacement, patched_pattern in replacements:
            source, count = re.subn(pattern, replacement, source)
            if count > 1 or len(re.findall(patched_pattern, source)) != 1:
                raise ValueError(f"unexpected Sink serde output for {field}")
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
