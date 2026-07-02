# Bugfix: Generated Codex doc instructs using Copilot instruction files

- Date: 2026-07-02
- Area: `internal/cli` IDE doc generation (spec 0005 follow-up)
- Severity: Medium (wrong guidance, not a crash)
- Status: Implemented

## Symptom

After `geremmyas global --targets codex --force core sdd`, the generated
`~/.codex/AGENTS.md` tells the assistant to use Copilot instruction files, even
though the target is Codex and the doc already ships a correct on-demand
Instructions index pointing to `~/.codex/instructions/`.

## Reproduction

```sh
./geremmyas global --targets codex --force core sdd
sed -n '1,40p' ~/.codex/AGENTS.md
grep -n '\.github/instructions/\*\.instructions\.md\|copilot-instructions\.md' ~/.codex/AGENTS.md
```

Observed: the embedded operating contract still contains:

- `Use `.github/copilot-instructions.md` as supporting context ...` (Instruction Order #2)
- `Load `.github/instructions/*.instructions.md` only when their `applyTo` patterns match ...` (Instruction Order #3)
- `Use matching `.github/instructions/*.instructions.md` for edits in a single technology ...` (Skill Routing)

## Root cause

`buildIDEAgentsDoc` (`internal/cli/generate_claude.go`) inlines the whole
`AGENTS.md` body verbatim for every IDE target. The canonical
`project/AGENTS.md` describes instruction loading in Copilot terms
(`.github/instructions/*.instructions.md` + `applyTo`,
`.github/copilot-instructions.md`). For Codex (and other non-Copilot targets)
this contradicts the generated `## Instructions (on demand)` index, which is the
correct, assistant-appropriate routing.

Copilot's `applyTo` auto-loading is driven by instruction-file frontmatter, not
by the AGENTS.md text, so the AGENTS.md wording is advisory and can be made
assistant-neutral without breaking Copilot.

## Hypotheses considered

1. The generated Instructions index is missing. Rejected: it is present and
   correct; the contradiction comes from the embedded body.
2. The Codex path is wrong. Rejected: fixed in spec 0005; the file is at
   `~/.codex/AGENTS.md` and reads correctly.
3. The embedded `AGENTS.md` body carries Copilot-specific instruction guidance.
   Confirmed: lines 9, 11, and 166 of `project/AGENTS.md`.

## Proposed fix

Two options:

- **Option C (recommended, single source):** Reword the three references in
  `project/AGENTS.md` to be assistant-neutral so no generated doc instructs
  using Copilot instruction files:
  - #2: keep the project-overview intent without hardcoding the Copilot file
    name as the mechanism.
  - #3: "Load the technology instruction that matches the files being edited,
    using the instruction index provided for your assistant."
  - Skill Routing: "Use the matching technology instruction for edits in a
    single technology ...".
  Copilot keeps working (frontmatter `applyTo` drives loading). Codex and the
  other IDE docs then contain only their own on-demand index.

- **Option A (generator-scoped):** Keep `project/AGENTS.md` Copilot-specific and
  have `buildIDEAgentsDoc` rewrite/neutralize those references only for targets
  that ship an instruction index. More code and fragile string surgery on
  user-authored content; rejected in favor of Option C.

## Regression test

Add a generator test asserting the generated global Codex `AGENTS.md`:

- does not contain the substring `.github/instructions/*.instructions.md`;
- does not instruct using `.github/copilot-instructions.md` as the instruction
  source;
- still contains the `## Instructions (on demand)` index pointing to
  `~/.codex/instructions/`.

## Verification plan

- `go test ./internal/cli`
- Regenerate and re-inspect `~/.codex/AGENTS.md`.

## Verification results

- `go test ./internal/cli` passes.
- Full suite `go test ./...` passes.
- After rebuilding the binary and running
  `./geremmyas global --targets codex --force core sdd`,
  `~/.codex/AGENTS.md` no longer contains:
  - `.github/copilot-instructions.md`
  - `.github/instructions/*.instructions.md`
  - `~/.copilot/instructions`
- The Codex instruction index remains present and points to
  `~/.codex/instructions/<file>`.
