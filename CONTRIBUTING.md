# Contributing to geremmyas

Thanks for your interest in contributing! This project is a personal set of GitHub Copilot configurations, but contributions that improve quality, fix bugs, or add useful defaults are welcome.

## How to Contribute

1. **Fork** the repository
2. **Create a branch** from `main` (`git checkout -b feat/my-change`)
3. **Make your changes**
4. **Test the install script** locally:
   ```bash
   bash install.sh
   bash install.sh --project
   bash uninstall.sh
   ```
5. **Commit** using [Conventional Commits](https://www.conventionalcommits.org/):
   - `feat:` new feature or config
   - `fix:` bug fix
   - `docs:` documentation only
   - `chore:` maintenance, CI, tooling
6. **Open a Pull Request** against `main`

## What Makes a Good Contribution

- **Instructions**: language-specific patterns that are broadly useful, not project-specific
- **Agents**: well-defined role with clear triggers and structured output
- **Skills**: reusable procedures with templates/assets when needed
- **Guardrails**: rules that prevent common destructive mistakes
- **Bug fixes**: especially for `install.sh` across different OS/shell environments

## Guidelines

- Keep files concise — Copilot works better with focused instructions
- Follow existing naming conventions (`kebab-case.instructions.md`, etc.)
- Test shell scripts with `shellcheck` before submitting
- All content must be in English

## Reporting Issues

Use [GitHub Issues](https://github.com/woliveiras/geremmyas/issues) with the provided templates. Include your OS and VS Code version when reporting install problems.

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you agree to uphold it.
