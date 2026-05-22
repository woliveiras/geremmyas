# Contributing to geremmyas

Thanks for your interest in contributing! This project is a personal set of GitHub Copilot configurations, but contributions that improve quality, fix bugs, or add useful defaults are welcome.

## Setup

1. **Fork** and clone the repository
2. **Install tools** with [mise](https://mise.jdx.dev/):
   ```bash
   mise trust
   mise install
   ```
3. **Build** the binary:
   ```bash
   go build ./cmd/geremmyas
   ```
4. **Run tests**:
   ```bash
   go test ./...
   ```

## How to Contribute

1. **Create a branch** from `main` (`git checkout -b feat/my-change`)
2. **Make your changes**
3. **Run tests and linting**:
   ```bash
   go test ./...
   shellcheck install.sh uninstall.sh
   ```
4. **Commit** using [Conventional Commits](https://www.conventionalcommits.org/):
   - `feat:` new feature or config
   - `fix:` bug fix
   - `docs:` documentation only
   - `chore:` maintenance, CI, tooling
5. **Open a Pull Request** against `main`

## What Makes a Good Contribution

- **Packs**: new catalog packs with useful instructions, agents, or skills for a specific stack
- **Instructions**: language-specific patterns that are broadly useful, not project-specific
- **Agents**: well-defined role with clear triggers and structured output
- **Skills**: reusable procedures with templates/assets when needed
- **Guardrails**: rules that prevent common destructive mistakes
- **Bug fixes**: especially for `install.sh` across different OS/shell environments

## Guidelines

- Keep files concise, Copilot works better with focused instructions
- Follow existing naming conventions (`kebab-case.instructions.md`, etc.)
- Test shell scripts with `shellcheck` before submitting
- All content must be in English
- Add new packs to `catalog/packs.json` with proper dependencies

## Project Structure

```
cmd/geremmyas/main.go   # Entry point
internal/cli/           # CLI commands (init, sync, global, catalog, config)
catalog/packs.json      # Pack definitions and dependencies
project/                # Embedded project-level files (synced to repos)
user/                   # Embedded user-level files (global install)
```

## Reporting Issues

Use [GitHub Issues](https://github.com/woliveiras/geremmyas/issues) with the provided templates. Include your OS and VS Code version when reporting install problems.

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you agree to uphold it.
