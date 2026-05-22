#!/usr/bin/env python3
"""
validate-bib.py - Bibliography validator for LaTeX papers.

Checks:
  1. Orphaned \\cite{} keys: used in .tex but missing from .bib
  2. Unused .bib entries: defined in .bib but never cited in .tex
  3. Malformed .bib entries: missing required fields for their entry type

Usage:
  python validate-bib.py --tex paper.tex --bib references.bib
  python validate-bib.py --tex paper.tex --bib references.bib --strict
  python validate-bib.py --tex paper.tex --bib references.bib --no-unused

Run after any round of work where new citations were added.
Skip if no new \\cite{} keys were introduced since last run.
"""

import argparse
import re
import sys
from pathlib import Path


# BibTeX required fields by entry type.
# Source: BibTeX documentation + common SE venue expectations.
REQUIRED_FIELDS = {
    "article":       {"author", "title", "journal", "year"},
    "inproceedings": {"author", "title", "booktitle", "year"},
    "proceedings":   {"title", "year"},
    "book":          {"author", "title", "publisher", "year"},
    "incollection":  {"author", "title", "booktitle", "publisher", "year"},
    "phdthesis":     {"author", "title", "school", "year"},
    "mastersthesis": {"author", "title", "school", "year"},
    "techreport":    {"author", "title", "institution", "year"},
    "misc":          {"author", "title", "year"},
    "unpublished":   {"author", "title", "note"},
    "inbook":        {"author", "title", "chapter", "publisher", "year"},
    "manual":        {"title"},
    "conference":    {"author", "title", "booktitle", "year"},  # alias for inproceedings
}

# Fields that count as "author" for entry types that allow "editor" instead
AUTHOR_OR_EDITOR_TYPES = {"proceedings", "book", "incollection", "inbook"}


def extract_cite_keys(tex_content: str) -> set[str]:
    """
    Extract all citation keys from LaTeX content.
    Handles: \\cite{key}, \\cite{k1,k2}, \\citet{}, \\citep{}, \\citealt{},
    \\citealp{}, \\citeauthor{}, \\citeyear{}, \\citenum{}.
    Skips commented lines (lines starting with %).
    """
    keys: set[str] = set()
    # Remove comments (lines starting with optional whitespace + %)
    lines = tex_content.splitlines()
    active_lines = []
    for line in lines:
        stripped = line.lstrip()
        if stripped.startswith("%"):
            continue
        # Inline comment: strip everything after unescaped %
        line = re.sub(r"(?<!\\)%.*$", "", line)
        active_lines.append(line)
    clean_content = "\n".join(active_lines)

    cite_pattern = re.compile(
        r"\\cite[a-zA-Z*]*\{([^}]+)\}"
    )
    for match in cite_pattern.finditer(clean_content):
        raw_keys = match.group(1)
        for key in raw_keys.split(","):
            key = key.strip()
            if key:
                keys.add(key)
    return keys


def parse_bib_file(bib_content: str) -> dict[str, dict]:
    """
    Parse a .bib file and return a dict mapping key -> {type, fields}.
    Handles multi-line field values and braces/quotes.
    Does NOT handle @string or @preamble (skipped).
    """
    entries: dict[str, dict] = {}

    # Match @type{key, ...}
    entry_pattern = re.compile(
        r"@(\w+)\s*\{\s*([^,\s]+)\s*,\s*(.*?)\n\}",
        re.DOTALL | re.IGNORECASE,
    )

    for match in entry_pattern.finditer(bib_content):
        entry_type = match.group(1).lower()
        entry_key = match.group(2).strip()
        fields_raw = match.group(3)

        if entry_type in ("string", "preamble", "comment"):
            continue

        fields = _parse_fields(fields_raw)
        entries[entry_key] = {
            "type": entry_type,
            "fields": {k.lower() for k in fields},
        }

    return entries


def _parse_fields(fields_raw: str) -> dict[str, str]:
    """
    Extract field names from a BibTeX entry body.
    Returns dict of field_name -> field_value (value not used, only name).
    """
    fields = {}
    # Match field = {value} or field = "value" or field = number
    field_pattern = re.compile(
        r"^\s*(\w+)\s*=\s*(?:\{.*?\}|\".*?\"|[\w\d]+)",
        re.MULTILINE | re.DOTALL,
    )
    for match in field_pattern.finditer(fields_raw):
        field_name = match.group(1).lower()
        fields[field_name] = ""
    return fields


def check_required_fields(entry_key: str, entry: dict) -> list[str]:
    """
    Check that a BibTeX entry has all required fields.
    Returns list of missing field names (empty if OK).
    """
    entry_type = entry["type"]
    present_fields = entry["fields"]
    required = REQUIRED_FIELDS.get(entry_type)

    if required is None:
        # Unknown entry type - skip field check
        return []

    # For types that accept editor instead of author
    effective_required = set(required)
    if entry_type in AUTHOR_OR_EDITOR_TYPES:
        if "editor" in present_fields and "author" not in present_fields:
            effective_required.discard("author")

    missing = sorted(effective_required - present_fields)
    return missing


def load_tex_recursive(tex_path: Path) -> str:
    """
    Load a .tex file and recursively expand \\input{} and \\include{} directives.
    This handles multi-file papers (e.g., chapters in separate files).
    Limits recursion depth to 10 to avoid infinite loops.
    """
    def _load(path: Path, depth: int = 0) -> str:
        if depth > 10:
            return ""
        try:
            content = path.read_text(encoding="utf-8", errors="replace")
        except FileNotFoundError:
            return ""

        # Expand \input{file} and \include{file}
        def expand_include(match):
            included_name = match.group(1).strip()
            if not included_name.endswith(".tex"):
                included_name += ".tex"
            included_path = path.parent / included_name
            return _load(included_path, depth + 1)

        content = re.sub(r"\\(?:input|include)\{([^}]+)\}", expand_include, content)
        return content

    return _load(tex_path)


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Validate LaTeX bibliography: orphaned cites, unused entries, missing fields."
    )
    parser.add_argument("--tex", required=True, help="Path to main .tex file")
    parser.add_argument("--bib", required=True, help="Path to .bib file")
    parser.add_argument(
        "--no-unused",
        action="store_true",
        help="Do not report unused .bib entries (useful for shared bibliography files)",
    )
    parser.add_argument(
        "--strict",
        action="store_true",
        help="Exit with code 1 if any orphaned cites or missing required fields found",
    )
    args = parser.parse_args()

    tex_path = Path(args.tex)
    bib_path = Path(args.bib)

    if not tex_path.exists():
        print(f"ERROR: .tex file not found: {tex_path}", file=sys.stderr)
        return 2
    if not bib_path.exists():
        print(f"ERROR: .bib file not found: {bib_path}", file=sys.stderr)
        return 2

    tex_content = load_tex_recursive(tex_path)
    bib_content = bib_path.read_text(encoding="utf-8", errors="replace")

    cite_keys = extract_cite_keys(tex_content)
    bib_entries = parse_bib_file(bib_content)

    bib_keys = set(bib_entries.keys())

    # ----------------------------------------------------------------
    # Check 1: Orphaned \cite{} - used in .tex but absent from .bib
    # ----------------------------------------------------------------
    orphaned = sorted(cite_keys - bib_keys)

    # ----------------------------------------------------------------
    # Check 2: Unused .bib entries - defined but never cited
    # ----------------------------------------------------------------
    unused = sorted(bib_keys - cite_keys)

    # ----------------------------------------------------------------
    # Check 3: Missing required fields
    # ----------------------------------------------------------------
    field_errors: list[tuple[str, str, list[str]]] = []
    for key, entry in bib_entries.items():
        missing = check_required_fields(key, entry)
        if missing:
            field_errors.append((key, entry["type"], missing))
    field_errors.sort(key=lambda x: x[0])

    # ----------------------------------------------------------------
    # Report
    # ----------------------------------------------------------------
    has_errors = False

    print(f"\n{'='*60}")
    print(f"  Bibliography Validation Report")
    print(f"  .tex : {tex_path}")
    print(f"  .bib : {bib_path}")
    print(f"{'='*60}")
    print(f"  Cite keys in .tex : {len(cite_keys)}")
    print(f"  Entries in .bib   : {len(bib_keys)}")
    print()

    # Orphaned cites (critical)
    if orphaned:
        has_errors = True
        print(f"[CRITICAL] Orphaned \\cite{{}} keys ({len(orphaned)}) - in .tex but missing from .bib:")
        for key in orphaned:
            print(f"    \\cite{{{key}}}")
        print()
    else:
        print("[OK] No orphaned \\cite{} keys.")
        print()

    # Unused entries (informational unless --no-unused suppressed)
    if not args.no_unused:
        if unused:
            print(f"[INFO] Unused .bib entries ({len(unused)}) - defined in .bib but never cited:")
            for key in unused:
                entry_type = bib_entries[key]["type"]
                print(f"    {key}  (@{entry_type})")
            print()
        else:
            print("[OK] No unused .bib entries.")
            print()

    # Missing required fields (warning)
    if field_errors:
        has_errors = True
        print(f"[WARNING] Missing required fields ({len(field_errors)} entries):")
        for key, entry_type, missing_fields in field_errors:
            print(f"    {key}  (@{entry_type})  - missing: {', '.join(missing_fields)}")
        print()
    else:
        print("[OK] All .bib entries have required fields.")
        print()

    print(f"{'='*60}")
    if not orphaned and not field_errors:
        print("  Result: PASS - bibliography is consistent.")
    else:
        print("  Result: ISSUES FOUND - see above.")
    print(f"{'='*60}\n")

    if args.strict and has_errors:
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main())
