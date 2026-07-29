#!/usr/bin/env python3
# /// script
# requires-python = ">=3.10"
# dependencies = ["Pillow>=10"]
# ///
"""Reduce a source image to logical-resolution, palette-limited pixel art."""

from __future__ import annotations

import argparse
import json
import sys
from collections import Counter
from pathlib import Path

from PIL import Image


def image_pixels(image: Image.Image):
    getter = getattr(image, "get_flattened_data", None)
    return getter() if getter is not None else image.getdata()


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Create a hard-edged logical-resolution pixel-art PNG."
    )
    parser.add_argument("input", type=Path)
    parser.add_argument("output", type=Path)
    parser.add_argument("--width", type=int, required=True)
    parser.add_argument("--height", type=int, required=True)
    parser.add_argument("--colors", type=int, default=16)
    parser.add_argument("--padding", type=int, default=0)
    parser.add_argument(
        "--anchor",
        choices=("bottom-center", "center", "top-left"),
        default="bottom-center",
    )
    parser.add_argument("--alpha-threshold", type=int, default=96)
    parser.add_argument("--no-trim", action="store_true")
    parser.add_argument("--preview", type=Path)
    parser.add_argument("--preview-scale", type=int, default=8)
    parser.add_argument("--palette-output", type=Path)
    return parser.parse_args()


def anchor_position(
    anchor: str,
    canvas_size: tuple[int, int],
    subject_size: tuple[int, int],
    padding: int,
) -> tuple[int, int]:
    canvas_width, canvas_height = canvas_size
    subject_width, subject_height = subject_size
    if anchor == "bottom-center":
        return (
            (canvas_width - subject_width) // 2,
            canvas_height - padding - subject_height,
        )
    if anchor == "center":
        return (
            (canvas_width - subject_width) // 2,
            (canvas_height - subject_height) // 2,
        )
    return padding, padding


def visible_bbox(image: Image.Image, threshold: int) -> tuple[int, int, int, int] | None:
    alpha = image.getchannel("A")
    hard_mask = alpha.point(lambda value: 255 if value >= threshold else 0)
    return hard_mask.getbbox()


def first_visible_rgb(image: Image.Image, alpha: Image.Image) -> tuple[int, int, int]:
    rgb = image.convert("RGB")
    for color, opacity in zip(image_pixels(rgb), image_pixels(alpha)):
        if opacity:
            return color
    return (0, 0, 0)


def main() -> int:
    args = parse_args()
    if args.width <= 0 or args.height <= 0:
        raise SystemExit("width and height must be positive")
    if not 2 <= args.colors <= 256:
        raise SystemExit("colors must be between 2 and 256")
    if not 0 <= args.alpha_threshold <= 255:
        raise SystemExit("alpha threshold must be between 0 and 255")
    if args.padding < 0 or args.padding * 2 >= min(args.width, args.height):
        raise SystemExit("padding leaves no drawable area")

    with Image.open(args.input) as source:
        source_rgba = source.convert("RGBA")

    bbox = (
        (0, 0, *source_rgba.size)
        if args.no_trim
        else visible_bbox(source_rgba, args.alpha_threshold)
    )
    if bbox is None:
        raise SystemExit("input contains no visible pixels")

    subject = source_rgba.crop(bbox)
    available_width = args.width - args.padding * 2
    available_height = args.height - args.padding * 2
    scale = min(
        available_width / subject.width,
        available_height / subject.height,
    )
    resized_size = (
        max(1, round(subject.width * scale)),
        max(1, round(subject.height * scale)),
    )
    resampling = (
        Image.Resampling.LANCZOS if scale < 1 else Image.Resampling.NEAREST
    )
    subject = subject.resize(resized_size, resampling)

    canvas = Image.new("RGBA", (args.width, args.height), (0, 0, 0, 0))
    position = anchor_position(
        args.anchor, canvas.size, subject.size, args.padding
    )
    canvas.alpha_composite(subject, position)

    alpha = canvas.getchannel("A").point(
        lambda value: 255 if value >= args.alpha_threshold else 0
    )
    filler = first_visible_rgb(canvas, alpha)
    flattened = Image.new("RGB", canvas.size, filler)
    flattened.paste(canvas.convert("RGB"), mask=alpha)
    quantized = flattened.quantize(
        colors=args.colors,
        method=Image.Quantize.MEDIANCUT,
        dither=Image.Dither.NONE,
    ).convert("RGBA")
    quantized.putalpha(alpha)

    args.output.parent.mkdir(parents=True, exist_ok=True)
    quantized.save(args.output, format="PNG", optimize=True)

    if args.preview:
        if args.preview_scale <= 0:
            raise SystemExit("preview scale must be positive")
        preview = quantized.resize(
            (
                args.width * args.preview_scale,
                args.height * args.preview_scale,
            ),
            Image.Resampling.NEAREST,
        )
        args.preview.parent.mkdir(parents=True, exist_ok=True)
        preview.save(args.preview, format="PNG", optimize=True)

    colors = Counter(
        pixel[:3] for pixel in image_pixels(quantized) if pixel[3] > 0
    ).most_common()
    if args.palette_output:
        args.palette_output.parent.mkdir(parents=True, exist_ok=True)
        args.palette_output.write_text(
            json.dumps(
                [
                    {
                        "hex": f"#{red:02x}{green:02x}{blue:02x}",
                        "pixels": count,
                    }
                    for (red, green, blue), count in colors
                ],
                indent=2,
            )
            + "\n",
            encoding="utf-8",
        )

    print(
        f"wrote {args.output} "
        f"({args.width}x{args.height}, {len(colors)} visible colors, "
        f"source scale {scale:.4f})"
    )
    if args.preview:
        print(f"wrote {args.preview} ({args.preview_scale}x nearest preview)")
    return 0


if __name__ == "__main__":
    sys.exit(main())
