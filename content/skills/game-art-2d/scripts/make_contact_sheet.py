#!/usr/bin/env python3
# /// script
# requires-python = ">=3.10"
# dependencies = ["Pillow>=10"]
# ///
"""Render ordered images on a labeled checkerboard contact sheet."""

from __future__ import annotations

import argparse
import math
import sys
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Build a labeled contact sheet for visual review."
    )
    parser.add_argument("images", nargs="+", type=Path)
    parser.add_argument("--output", type=Path, required=True)
    parser.add_argument("--columns", type=int, default=4)
    parser.add_argument("--cell-width", type=int)
    parser.add_argument("--cell-height", type=int)
    parser.add_argument("--padding", type=int, default=12)
    parser.add_argument("--checker-size", type=int, default=12)
    parser.add_argument(
        "--resample",
        choices=("nearest", "lanczos"),
        default="nearest",
    )
    parser.add_argument("--no-labels", action="store_true")
    return parser.parse_args()


def draw_checkerboard(
    draw: ImageDraw.ImageDraw,
    box: tuple[int, int, int, int],
    size: int,
) -> None:
    left, top, right, bottom = box
    light = (214, 214, 214, 255)
    dark = (174, 174, 174, 255)
    for y in range(top, bottom, size):
        for x in range(left, right, size):
            color = light if ((x - left) // size + (y - top) // size) % 2 == 0 else dark
            draw.rectangle(
                (x, y, min(x + size, right), min(y + size, bottom)),
                fill=color,
            )


def main() -> int:
    args = parse_args()
    if args.columns <= 0 or args.padding < 0 or args.checker_size <= 0:
        raise SystemExit("columns and checker size must be positive; padding non-negative")

    images: list[tuple[Path, Image.Image]] = []
    for path in args.images:
        with Image.open(path) as source:
            images.append((path, source.convert("RGBA")))

    label_height = 20 if not args.no_labels else 0
    content_width = args.cell_width or max(image.width for _, image in images)
    content_height = args.cell_height or max(image.height for _, image in images)
    cell_width = content_width + args.padding * 2
    cell_height = content_height + args.padding * 2 + label_height
    rows = math.ceil(len(images) / args.columns)
    sheet = Image.new(
        "RGBA",
        (args.columns * cell_width, rows * cell_height),
        (32, 32, 36, 255),
    )
    draw = ImageDraw.Draw(sheet)
    font = ImageFont.load_default()
    resampling = (
        Image.Resampling.NEAREST
        if args.resample == "nearest"
        else Image.Resampling.LANCZOS
    )

    for index, (path, image) in enumerate(images):
        row, column = divmod(index, args.columns)
        cell_x = column * cell_width
        cell_y = row * cell_height
        content_box = (
            cell_x + args.padding,
            cell_y + args.padding,
            cell_x + args.padding + content_width,
            cell_y + args.padding + content_height,
        )
        draw_checkerboard(draw, content_box, args.checker_size)

        scale = min(
            content_width / image.width,
            content_height / image.height,
            1.0,
        )
        rendered_size = (
            max(1, round(image.width * scale)),
            max(1, round(image.height * scale)),
        )
        rendered = image.resize(rendered_size, resampling)
        x = content_box[0] + (content_width - rendered.width) // 2
        y = content_box[1] + (content_height - rendered.height) // 2
        sheet.alpha_composite(rendered, (x, y))

        if not args.no_labels:
            label = f"{index + 1:02d}  {path.name}"
            label_y = cell_y + cell_height - label_height
            draw.rectangle(
                (cell_x, label_y, cell_x + cell_width, cell_y + cell_height),
                fill=(18, 18, 22, 255),
            )
            draw.text(
                (cell_x + args.padding, label_y + 4),
                label,
                font=font,
                fill=(245, 245, 245, 255),
            )

    args.output.parent.mkdir(parents=True, exist_ok=True)
    sheet.save(args.output, format="PNG", optimize=True)
    print(
        f"wrote {args.output} "
        f"({args.columns} columns, {rows} rows, {len(images)} images)"
    )
    return 0


if __name__ == "__main__":
    sys.exit(main())
