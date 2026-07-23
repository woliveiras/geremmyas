---
name: update-docs
description: "Update project documentation after implementing a feature. Use when: updating docs, documentation, update docs after implementation. Do not use: for doc-first design, specification writing."
---


# Update Docs

Update documentation in `docs/` after a feature is implemented.

## When to Use

- After finishing a feature implementation
- When specs, code, or architecture changed and docs are stale
- When onboarding docs need to reflect new patterns

## Procedure

1. Read the spec for the implemented feature (from `specs/` or the conversation)
2. Read the implemented code to understand what actually shipped
3. Search `docs/` for existing documentation that may need updates
4. For each affected doc:
   - If it exists: update the relevant sections, preserving the existing format
   - If it doesn't exist: create it following the format of neighboring docs
5. Check for broken references (links to moved/renamed files)

## What to Document

- **API changes**: new endpoints, changed parameters, updated responses
- **Architecture changes**: new modules, changed boundaries, new patterns
- **Setup changes**: new dependencies, changed build steps, new env vars
- **Configuration changes**: new settings, changed defaults

## Rules

- Match the existing documentation style and format
- Do NOT rewrite docs that don't need changes
- Keep docs concise — one paragraph per concept, not walls of text
- Update the table of contents if the doc has one
- If no `docs/` directory exists, create it with a README

## Write like a human

Write prose a practitioner would actually write, not generic assistant output. Match the existing documentation style, but keep these AI writing tells out of the prose:

- No em dashes and no curly quotes: use commas, parentheses, or new sentences, and straight quotes (" ').
- Cut filler vocabulary: "delve", "leverage", "utilize" (use "use"), "robust", "seamless", "crucial", "pivotal", "testament", "underscore", "showcase", "foster", "landscape" (as an abstract noun).
- Drop significance padding ("stands as a testament to", "plays a pivotal role", "reflects a broader shift") and trailing "-ing" analysis clauses ("...highlighting its importance").
- Prefer plain "is"/"has" over "serves as"/"boasts". Attribute every claim to a named source or delete it.
- Skip forced tricolons and decorative emoji. State the point first; cut sentences that only announce a topic or restate the previous one.
