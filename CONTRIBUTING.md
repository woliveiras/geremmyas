# Contributing to geremmyas

Thanks for your interest in contributing! This project ships opinionated Copilot
configurations and a pack-based CLI. Contributions that improve quality, fix
bugs, or add broadly useful defaults are welcome.

## Documentation

| Doc | When to read |
| --- | --- |
| [docs/architecture.md](docs/architecture.md) | How embed, sync, global install, and symlinks work |
| [docs/creating-packs.md](docs/creating-packs.md) | Add or change packs, skills, and instructions |
| [README.md](README.md) | User-facing install, CLI, and pack catalog |
| [AGENTS.md](AGENTS.md) | Agent operating contract (symlink to `project/AGENTS.md`) |

## Setup

1. **Fork** and clone the repository
2. **Install tools** with [mise](https://mise.jdx.dev/):
   ```bash
   mise trust
   mise install
   ```
3. **Build** the binary:
   ```bash
   go build -o geremmyas ./cmd/geremmyas
   ```
4. **Run tests**:
   ```bash
   go test ./...
   ./geremmyas doctor
   ```

## How to Contribute

1. **Create a branch** from `main` (`git checkout -b feat/my-change`)
2. **Make your changes** under `project/` and/or `catalog/packs.json`
3. **Run tests and linting**:
   ```bash
   go test ./...
   shellcheck install.sh uninstall.sh
   ./geremmyas doctor
   ```
4. **Commit** using [Conventional Commits](https://www.conventionalcommits.org/):
   - `feat:` new feature or config
   - `fix:` bug fix
   - `docs:` documentation only
   - `chore:` maintenance, CI, tooling
   - `feat!:` or footer `BREAKING CHANGE:` for major releases (pack removals, path changes)
5. **Open a Pull Request** against `main`

## What Makes a Good Contribution

- **Packs**: new catalog entries with clear `depends` and focused file lists
- **Instructions**: language or stack patterns with `applyTo` globs, not project-specific trivia
- **Agents**: narrow roles with clear triggers (see `project/.github/agents/`)
- **Skills**: reusable procedures; use `skill-authoring` skill in repo as reference
- **Guardrails**: safe defaults in `project/.github/hooks/guardrails-rules.txt`
- **CLI**: changes in `internal/cli/` with tests in `*_test.go`
- **Docs**: updates to `docs/` and README when behavior or packs change

## Adding a pack, skill, or instruction

See **[docs/creating-packs.md](docs/creating-packs.md)** for the full checklist.

Short version:

1. Add files under `project/`
2. Register paths in `catalog/packs.json`
3. Run `go test ./...` and `geremmyas doctor`
4. Test `init` + `sync` in a temporary directory

**Do not** run `geremmyas project` in this repository; edit `project/` directly.
Root symlinks keep Copilot aligned with the canonical tree. Copilot project
context for **this repo** is in `project/.github/copilot-instructions.geremmyas.md`
(not synced). The consumer template is `project/.github/copilot-instructions.md`.

## Guidelines

- Keep instruction and skill files concise
- Follow naming: `kebab-case.instructions.md`, `skills/<name>/SKILL.md`
- Every skill in `project/.github/skills/` should belong to a pack or be removed
- All user-facing prose in English
- Shell scripts: run `shellcheck` before submitting

## Project structure

```text
cmd/geremmyas/          Entry point
internal/cli/           CLI implementation
catalog/packs.json      Pack definitions
project/                Canonical content synced to consumer repos
user/                   Embedded prompts (optional, not synced by default)
docs/                   Contributor and architecture documentation
specs/                  Specs for geremmyas features (maintainer SDD)
```

## Reporting issues

Use [GitHub Issues](https://github.com/woliveiras/geremmyas/issues) with the
provided templates. Include OS and editor version for install problems.

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md).
