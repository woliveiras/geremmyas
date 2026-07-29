#!/usr/bin/env python3
# /// script
# requires-python = ">=3.10"
# dependencies = ["Pillow>=10"]
# ///
"""Normalize ordered sprite frames with one shared scale and anchor."""

from __future__ import annotations

import argparse
import json
import math
import sys
from pathlib import Path

from PIL import Image


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Pack ordered frames into equal cells with shared scaling."
    )
    parser.add_argument("frames", nargs="+", type=Path)
    parser.add_argument("--output", type=Path, required=True)
    parser.add_argument("--cell-width", type=int, required=True)
    parser.add_argument("--cell-height", type=int, required=True)
    parser.add_argument("--columns", type=int)
    parser.add_argument("--padding", type=int, default=0)
    parser.add_argument(
        "--anchor",
        choices=("bottom-center", "center", "top-left"),
        default="bottom-center",
    )
    parser.add_argument(
        "--resample",
        choices=("nearest", "lanczos"),
        default="nearest",
    )
    parser.add_argument("--allow-upscale", action="store_true")
    parser.add_argument("--manifest", type=Path)
    return parser.parse_args()


def anchor_position(
    anchor: str,
    cell_size: tuple[int, int],
    subject_size: tuple[int, int],
    padding: int,
) -> tuple[int, int]:
    cell_width, cell_height = cell_size
    subject_width, subject_height = subject_size
    if anchor == "bottom-center":
        return (
            (cell_width - subject_width) // 2,
            cell_height - padding - subject_height,
        )
    if anchor == "center":
        return (
            (cell_width - subject_width) // 2,
            (cell_height - subject_height) // 2,
        )
    return padding, padding


def main() -> int:
    args = parse_args()
    if args.cell_width <= 0 or args.cell_height <= 0:
        raise SystemExit("cell dimensions must be positive")
    if args.padding < 0 or args.padding * 2 >= min(
        args.cell_width, args.cell_height
    ):
        raise SystemExit("padding leaves no drawable cell area")

    loaded: list[dict[str, object]] = []
    for path in args.frames:
        with Image.open(path) as source:
            rgba = source.convert("RGBA")
        bbox = rgba.getchannel("A").getbbox()
        if bbox is None:
            raise SystemExit(f"{path} contains no visible pixels")
        loaded.append({"path": path, "image": rgba, "bbox": bbox})

    max_width = max(item["bbox"][2] - item["bbox"][0] for item in loaded)
    max_height = max(item["bbox"][3] - item["bbox"][1] for item in loaded)
    available_width = args.cell_width - args.padding * 2
    available_height = args.cell_height - args.padding * 2
    shared_scale = min(
        available_width / max_width,
        available_height / max_height,
    )
    if not args.allow_upscale:
        shared_scale = min(1.0, shared_scale)

    columns = args.columns or len(loaded)
    if columns <= 0:
        raise SystemExit("columns must be positive")
    rows = math.ceil(len(loaded) / columns)
    sheet = Image.new(
        "RGBA",
        (columns * args.cell_width, rows * args.cell_height),
        (0, 0, 0, 0),
    )
    resampling = (
        Image.Resampling.NEAREST
        if args.resample == "nearest"
        else Image.Resampling.LANCZOS
    )
    manifest_frames: list[dict[str, object]] = []

    for index, item in enumerate(loaded):
        bbox = item["bbox"]
        cropped = item["image"].crop(bbox)
        resized_size = (
            max(1, round(cropped.width * shared_scale)),
            max(1, round(cropped.height * shared_scale)),
        )
        resized = cropped.resize(resized_size, resampling)
        local_x, local_y = anchor_position(
            args.anchor,
            (args.cell_width, args.cell_height),
            resized_size,
            args.padding,
        )
        row, column = divmod(index, columns)
        sheet_x = column * args.cell_width + local_x
        sheet_y = row * args.cell_height + local_y
        sheet.alpha_composite(resized, (sheet_x, sheet_y))
        manifest_frames.append(
            {
                "index": index,
                "source": str(item["path"]),
                "source_visible_bbox": list(bbox),
                "sheet_rect": [
                    column * args.cell_width,
                    row * args.cell_height,
                    args.cell_width,
                    args.cell_height,
                ],
                "content_rect": [sheet_x, sheet_y, *resized_size],
            }
        )

    args.output.parent.mkdir(parents=True, exist_ok=True)
    sheet.save(args.output, format="PNG", optimize=True)

    if args.manifest:
        args.manifest.parent.mkdir(parents=True, exist_ok=True)
        args.manifest.write_text(
            json.dumps(
                {
                    "image": str(args.output),
                    "cell_size": [args.cell_width, args.cell_height],
                    "columns": columns,
                    "rows": rows,
                    "anchor": args.anchor,
                    "shared_scale": shared_scale,
                    "frames": manifest_frames,
                },
                indent=2,
            )
            + "\n",
            encoding="utf-8",
        )

    print(
        f"wrote {args.output} ({columns}x{rows} cells, "
        f"{args.cell_width}x{args.cell_height}, shared scale {shared_scale:.4f})"
    )
    return 0


if __name__ == "__main__":
    sys.exit(main())
