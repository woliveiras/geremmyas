# Creating packs, skills, and instructions

This guide is for contributors changing the geremmyas catalog or `project/`
tree. End users only add pack **names** to `geremmyas.yml`; they do not edit
`catalog/packs.json`.

## Prerequisites

```bash
mise trust && mise install
go test ./...
go build -o geremmyas ./cmd/geremmyas
./geremmyas doctor
```

Work from a branch off `main`. Use Conventional Commits (`feat:`, `fix:`,
`docs:`, `chore:`). Use `feat!:` or `BREAKING CHANGE:` when pack removal or
path changes break existing consumer repos.

## Add a new pack

### 1. Place files under `project/`

All synced content lives under `project/`:

| Kind | Typical path |
| --- | --- |
| Instruction | `project/.github/instructions/<name>.instructions.md` |
| Skill | `project/.github/skills/<name>/SKILL.md` (+ optional `assets/`, `references/`) |
| Agent | `project/.github/agents/<name>.agent.md` |
| Template | `project/templates/...` |

Follow naming:

- Instructions: `kebab-case.instructions.md` with YAML frontmatter and `applyTo` globs.
- Skills: folder `kebab-case` with one top-level `SKILL.md`; supporting files
  under `assets/` or `references/` must not be named `SKILL.md`.
- Agents: `kebab-case.agent.md`.

### 2. Register the pack in `catalog/packs.json`

```json
{
  "name": "my-pack",
  "tier": "stack",
  "description": "One line shown in geremmyas list.",
  "depends": ["typescript-base"],
  "files": [
    {
      "source": "project/.github/instructions/my-stack.instructions.md",
      "target": ".github/instructions/my-stack.instructions.md"
    }
  ]
}
```

Rules:

- `tier` is required and must be `core` or `stack`. `core` is reserved for the
  always-on SDD pipeline and guardrails (the `core` and `sdd` packs). Every other
  pack is `stack` (opt-in per project). `geremmyas lint` and `geremmyas doctor`
  fail when a pack is missing a tier or uses an invalid value.
- `source` paths are relative to the repo root and must match embed paths.
- `target` paths are relative to the consumer repo root.
- Use `depends` when the pack needs another pack’s files (instructions, base rules).
- Directory entries copy recursively:

```json
{
  "source": "project/.github/skills/my-skill",
  "target": ".github/skills/my-skill"
}
```

### 3. Validate

```bash
go test ./internal/cli/...
./geremmyas doctor
./geremmyas list | grep my-pack
```

In a temporary directory:

```bash
mkdir /tmp/geremmyas-pack-test && cd /tmp/geremmyas-pack-test
/path/to/geremmyas init --packs core,my-pack
/path/to/geremmyas sync
ls -la .github/instructions/
```

### 4. Document

- Add the pack to the catalog table in [README.md](../README.md) if it is a
  user-facing pack.
- Mention breaking changes in the commit message for release-please.

## Extend an existing pack

Prefer adding `files` entries to an existing pack when the content belongs to the
same stack (for example a new Fastify instruction under `fastify`).

Avoid duplicating the same `target` in multiple packs; the second install is skipped.

## Add a skill to the `sdd` pack

SDD is limited to user-invoked, general workflow capabilities. Before adding a
discoverable skill, verify that the behavior cannot be a reference in an
existing workflow or a custom agent role. To ship a new workflow skill:

1. Create `project/.github/skills/<name>/SKILL.md`.
2. Add a file entry to the `sdd` pack’s `files` array.
3. Reference the skill in `project/AGENTS.md` skill routing if agents should use it by default.
4. Run tests and `doctor`.

Skills that are optional or personal should get their own pack (like `blog` or
`premortem`), not `sdd`. Maintainer and decision helpers use the opt-in
`skill-maintenance` and `decision-support` packs.

### Skill taxonomy

- **Skill:** a capability a user can request directly.
- **Reference:** internal policy, checklist, template, example, or composition
  step loaded only by its owning skill or agent.
- **Agent:** an isolated role for expensive exploration, specification, review,
  or architecture work.
- **Instruction:** file-scoped technology guidance selected by path.

### Context budgets

`geremmyas lint` enforces discovery and always-on context limits:

| Surface | Limit |
| --- | --- |
| Skill frontmatter `description` | 240 characters |
| Top-level skill body | 250 lines |
| Nested files named `SKILL.md` | 0 |
| Public skills in the `sdd` pack | 10 |
| `project/AGENTS.md` | 700 words |

Keep the top-level skill focused on routing, required gates, and the normal
workflow. Put long examples, optional variants, detailed templates, and
background material in clearly named files under `references/` or `assets/`,
then tell the skill when to load each file. These limits keep metadata cheap to
discover and prevent the default SDD contract from growing without review.

## Add an instruction

1. Create `project/.github/instructions/<topic>.instructions.md`.
2. Set `description` and `applyTo` globs in frontmatter.
3. Wire the file in the appropriate pack’s `files` list.
4. Keep instructions short; move long examples to skill `assets/` or `references/`.

Instructions under `.github/instructions/` can also be installed globally via
`geremmyas global <pack>`. They are always copied to `~/.copilot/instructions/`.
With the `codex` target they are also mirrored to `~/.codex/instructions/` and
indexed in `~/.codex/AGENTS.md` by their `applyTo` glob, since Codex has no
`applyTo` auto-loading. Always set a meaningful `applyTo` so the Codex index can
tell the model when to read the file.

## Orphan skills

Files under `project/.github/skills/` that are **not** listed in any pack are
not shipped to consumers. Either add them to a pack or delete them to avoid drift.

## Testing checklist

- [ ] `go test ./...` passes
- [ ] `geremmyas lint` passes
- [ ] `geremmyas doctor` passes
- [ ] New pack appears in `geremmyas list`
- [ ] `init` + `sync` in a clean directory produces expected paths
- [ ] If changing customizable templates, note preserve behavior in PR description
- [ ] Rebuild binary before manual local sync tests

## Maintainer repository

Do not use `geremmyas project` in this repo. Edit `project/` directly; root
symlinks expose the same content to Copilot. See [architecture.md](architecture.md).

**Copilot instructions split:** `project/.github/copilot-instructions.md` is the
consumer template (in the `core` pack). Maintainer context lives in
`project/.github/copilot-instructions.geremmyas.md` (not in `packs.json`). Root
`.github/copilot-instructions.md` symlinks to the `.geremmyas` file.

For feature work on geremmyas itself, use `specs/` at the repo root (see
[`specs/README.md`](../specs/README.md)).
