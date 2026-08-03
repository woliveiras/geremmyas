# Creating packs, skills, instructions, and adapters

This guide is for contributors changing canonical content, target adapters, or
`catalog/packs.json`. End users add pack names to `geremmyas.yml`.

## Prerequisites

```bash
mise trust && mise install
go test ./...
go build -o geremmyas ./cmd/geremmyas
./geremmyas doctor
```

Use Conventional Commits. Treat removal of a pack or a consumer output path as
a breaking change.

## Choose the correct source tree

Put portable artifacts under `content/`:

| Kind | Typical source |
| --- | --- |
| Contract | `content/AGENTS.md` |
| Instruction | `content/instructions/<name>.instructions.md` |
| Skill | `content/skills/<name>/SKILL.md` |
| Template | `content/templates/...` |
| Prompt | `content/prompts/...` |
| Guardrail policy | `content/guardrails/...` |

Put an artifact under `targets/<assistant>/` only when its format or behavior
belongs to that assistant. Copilot hooks and Copilot instruction templates are
examples. Do not create a `.github/` source path merely because Copilot is one
supported destination.

## Add or extend a pack

Register semantic artifacts in `catalog/packs.json`:

```json
{
  "name": "my-pack",
  "tier": "stack",
  "description": "One line shown in geremmyas list.",
  "depends": ["typescript-base"],
  "files": [
    {
      "kind": "instruction",
      "source": "content/instructions/my-stack.instructions.md",
      "path": "my-stack.instructions.md"
    },
    {
      "kind": "skill",
      "source": "content/skills/my-skill",
      "path": "my-skill"
    }
  ]
}
```

Rules:

- `tier` is `core` or `stack`; `core` is reserved for the baseline workflow.
- `kind` controls target routing and must match the artifact semantics.
- `source` is relative to the repository root and must exist in the embed.
- `path` is target-neutral; it is not a consumer repository destination.
- Directory sources copy recursively.
- Use `depends` when another pack supplies required behavior.
- Avoid registering the same semantic artifact in multiple packs.

The target planner maps kinds to destinations. For example, one `skill` source
becomes `.github/skills/<path>` for Copilot and `.agents/skills/<path>` for
Codex and other neutral targets.

## Naming and context budgets

- Instructions: `kebab-case.instructions.md` with `description` and `applyTo`.
- Skills: `kebab-case/SKILL.md`; support files belong in `assets/` or
  `references/` and must not also be named `SKILL.md`.

`geremmyas lint` enforces:

| Surface | Limit |
| --- | --- |
| Skill frontmatter description | 240 characters |
| Top-level skill body | 250 lines |
| Nested files named `SKILL.md` | 0 |
| Public skills in the `base` pack | 7 |
| `content/AGENTS.md` | 700 words |

Keep discovery metadata and top-level workflows compact. Put detailed examples
and optional variants in named support files.

## Add a skill

1. Create `content/skills/<name>/SKILL.md`.
2. Add a `kind: "skill"` entry to the appropriate pack.
3. Reference it from `content/AGENTS.md` only when it is part of default routing.
4. Run lint, tests, and a target-matrix sync.

Phase-local coding workflows belong to `coding`; completion workflows belong to
`quality`. The technology-neutral `base` pack depends on both. Stack-specific
or personal capabilities should use their own opt-in pack.

## Add an instruction

1. Create `content/instructions/<topic>.instructions.md`.
2. Set meaningful `description` and `applyTo` frontmatter.
3. Add a `kind: "instruction"` entry to the appropriate pack.
4. Verify at least Copilot and Codex materialization.

Copilot uses the `applyTo` metadata in `.github/instructions`. Codex receives
the same source in `.codex/instructions`, with its generated instruction index
describing when each file applies.

## Add a delegation contract

Geremmyas does not bundle permanent custom agents. Prefer a compact contract in
the owning skill's `references/` directory, loaded only when delegation is due.
It should bound scope and ownership and require evidence, unknowns, findings,
and a stable output schema. The assistant runtime creates the temporary
subagent; the primary agent owns integration and Git.

The CLI retains generic `agent` adapter support for future target-specific
capabilities that provide a concrete guarantee, such as restricted tools or a
specialized runtime. Adding one is an exceptional adapter change and requires
focused cross-target generator tests.

## Add a target-specific adapter

Use `targets/<assistant>/` when a portable kind cannot express the integration.
Keep the adapter thin and continue sourcing shared prose or policy from
`content/` when possible. Add focused planner or generator tests for the new
target behavior.

## Validation

```bash
go test ./internal/cli -count=1
go test ./... -count=1
go run ./cmd/geremmyas lint
go run ./cmd/geremmyas doctor
```

Build once, then test in temporary directories:

```bash
go build -o /tmp/geremmyas ./cmd/geremmyas
mkdir /tmp/geremmyas-codex && cd /tmp/geremmyas-codex
/tmp/geremmyas init --packs core,my-pack --targets codex
/tmp/geremmyas sync
find . -maxdepth 3 -type f
```

Repeat for every affected target and for a mixed target selection. Verify that
Copilot-only `.github` content does not appear in a Codex-only project.

## Orphan content

Canonical files not referenced by any pack are not shipped. Add them to a pack
or remove them to avoid drift.

## Maintainer repository

Do not run `geremmyas project` here. Root symlinks expose canonical content and
the Copilot maintainer adapter for dogfooding. Feature work uses `specs/` at the
repository root.
