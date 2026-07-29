#!/usr/bin/env python3
# /// script
# requires-python = ">=3.10"
# dependencies = ["Pillow>=10"]
# ///
"""Inspect raster game assets and enforce simple technical contracts."""

from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path
from typing import Any

from PIL import Image


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Inspect image dimensions, alpha, bounds, grid, and palette."
    )
    parser.add_argument("images", nargs="+", type=Path)
    parser.add_argument("--expect-width", type=int)
    parser.add_argument("--expect-height", type=int)
    parser.add_argument("--cell-width", type=int)
    parser.add_argument("--cell-height", type=int)
    parser.add_argument(
        "--require-alpha",
        action="store_true",
        help="Require an alpha channel and at least one transparent pixel.",
    )
    parser.add_argument("--max-colors", type=int)
    parser.add_argument(
        "--color-scan-limit",
        type=int,
        default=1_000_000,
        help="Stop exact color counting above this many colors.",
    )
    parser.add_argument("--json", action="store_true", dest="as_json")
    return parser.parse_args()


def has_alpha_channel(image: Image.Image) -> bool:
    return image.mode in {"RGBA", "LA", "PA"} or "transparency" in image.info


def inspect_image(path: Path, args: argparse.Namespace) -> dict[str, Any]:
    issues: list[str] = []
    with Image.open(path) as source:
        source.load()
        rgba = source.convert("RGBA")
        width, height = rgba.size
        alpha = rgba.getchannel("A")
        alpha_min, alpha_max = alpha.getextrema()
        visible_bbox = alpha.getbbox()
        colors = rgba.getcolors(maxcolors=args.color_scan_limit)
        color_count: int | str = (
            sum(1 for _, color in colors if color[3] > 0)
            if colors is not None
            else f">{args.color_scan_limit}"
        )

        if args.expect_width is not None and width != args.expect_width:
            issues.append(f"width {width} != expected {args.expect_width}")
        if args.expect_height is not None and height != args.expect_height:
            issues.append(f"height {height} != expected {args.expect_height}")
        if args.cell_width is not None and width % args.cell_width != 0:
            issues.append(
                f"width {width} is not divisible by cell width {args.cell_width}"
            )
        if args.cell_height is not None and height % args.cell_height != 0:
            issues.append(
                f"height {height} is not divisible by cell height {args.cell_height}"
            )
        if args.require_alpha:
            if not has_alpha_channel(source):
                issues.append("image has no alpha channel")
            elif alpha_min == 255:
                issues.append("alpha channel contains no transparent pixels")
        if args.max_colors is not None:
            if isinstance(color_count, str):
                issues.append(
                    f"color count exceeds scan limit and max {args.max_colors}"
                )
            elif color_count > args.max_colors:
                issues.append(
                    f"color count {color_count} exceeds max {args.max_colors}"
                )

        pixel_count = width * height
        transparent_pixels = sum(count for count, value in alpha.getcolors() if value == 0)
        return {
            "path": str(path),
            "format": source.format,
            "mode": source.mode,
            "width": width,
            "height": height,
            "has_alpha_channel": has_alpha_channel(source),
            "alpha_range": [alpha_min, alpha_max],
            "transparent_pixels": transparent_pixels,
            "transparent_ratio": (
                round(transparent_pixels / pixel_count, 6) if pixel_count else 0
            ),
            "visible_bbox": list(visible_bbox) if visible_bbox else None,
            "color_count": color_count,
            "issues": issues,
        }


def main() -> int:
    args = parse_args()
    results: list[dict[str, Any]] = []
    failed = False

    for path in args.images:
        try:
            result = inspect_image(path, args)
        except (OSError, ValueError) as error:
            result = {"path": str(path), "issues": [f"cannot inspect: {error}"]}
        results.append(result)
        failed = failed or bool(result["issues"])

    if args.as_json:
        print(json.dumps(results, indent=2))
    else:
        for result in results:
            print(
                f"{result['path']}: "
                f"{result.get('width', '?')}x{result.get('height', '?')} "
                f"{result.get('format', '?')} {result.get('mode', '?')} "
                f"colors={result.get('color_count', '?')} "
                f"bbox={result.get('visible_bbox', '?')}"
            )
            for issue in result["issues"]:
                print(f"  ERROR: {issue}")

    return 1 if failed else 0


if __name__ == "__main__":
    sys.exit(main())
