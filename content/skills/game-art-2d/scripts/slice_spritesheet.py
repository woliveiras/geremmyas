#!/usr/bin/env python3
# /// script
# requires-python = ">=3.10"
# dependencies = ["Pillow>=10"]
# ///
"""Slice a regular spritesheet into ordered PNG frames."""

from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

from PIL import Image


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Slice a regular row-major spritesheet into frames."
    )
    parser.add_argument("input", type=Path)
    parser.add_argument("out_dir", type=Path)
    parser.add_argument("--frame-width", type=int)
    parser.add_argument("--frame-height", type=int)
    parser.add_argument("--columns", type=int)
    parser.add_argument("--rows", type=int)
    parser.add_argument("--count", type=int)
    parser.add_argument("--prefix")
    parser.add_argument("--start-index", type=int, default=1)
    parser.add_argument("--manifest", type=Path)
    return parser.parse_args()


def resolve_grid(
    width: int, height: int, args: argparse.Namespace
) -> tuple[int, int, int, int]:
    frame_mode = args.frame_width is not None or args.frame_height is not None
    grid_mode = args.columns is not None or args.rows is not None
    if frame_mode == grid_mode:
        raise SystemExit(
            "provide either --frame-width/--frame-height or --columns/--rows"
        )

    if frame_mode:
        if not args.frame_width or not args.frame_height:
            raise SystemExit("both frame dimensions are required")
        if width % args.frame_width or height % args.frame_height:
            raise SystemExit("sheet dimensions are not divisible by frame dimensions")
        columns = width // args.frame_width
        rows = height // args.frame_height
        return args.frame_width, args.frame_height, columns, rows

    if not args.columns or not args.rows:
        raise SystemExit("both rows and columns are required")
    if width % args.columns or height % args.rows:
        raise SystemExit("sheet dimensions are not divisible by rows and columns")
    return width // args.columns, height // args.rows, args.columns, args.rows


def main() -> int:
    args = parse_args()
    with Image.open(args.input) as source:
        sheet = source.convert("RGBA")

    frame_width, frame_height, columns, rows = resolve_grid(
        sheet.width, sheet.height, args
    )
    capacity = columns * rows
    count = args.count if args.count is not None else capacity
    if count <= 0 or count > capacity:
        raise SystemExit(f"count must be between 1 and sheet capacity {capacity}")

    prefix = args.prefix or args.input.stem
    args.out_dir.mkdir(parents=True, exist_ok=True)
    digits = max(2, len(str(args.start_index + count - 1)))
    frames: list[dict[str, object]] = []

    for offset in range(count):
        row, column = divmod(offset, columns)
        left = column * frame_width
        top = row * frame_height
        frame = sheet.crop(
            (left, top, left + frame_width, top + frame_height)
        )
        index = args.start_index + offset
        output = args.out_dir / f"{prefix}-{index:0{digits}d}.png"
        frame.save(output, format="PNG", optimize=True)
        frames.append(
            {
                "index": index,
                "path": str(output),
                "source_rect": [left, top, frame_width, frame_height],
                "visible_bbox": (
                    list(frame.getchannel("A").getbbox())
                    if frame.getchannel("A").getbbox()
                    else None
                ),
            }
        )

    if args.manifest:
        args.manifest.parent.mkdir(parents=True, exist_ok=True)
        args.manifest.write_text(
            json.dumps(
                {
                    "source": str(args.input),
                    "frame_size": [frame_width, frame_height],
                    "columns": columns,
                    "rows": rows,
                    "frames": frames,
                },
                indent=2,
            )
            + "\n",
            encoding="utf-8",
        )

    print(
        f"wrote {count} frames to {args.out_dir} "
        f"({frame_width}x{frame_height}, row-major)"
    )
    return 0


if __name__ == "__main__":
    sys.exit(main())
