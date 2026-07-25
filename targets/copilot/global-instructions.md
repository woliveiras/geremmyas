# Global Copilot Instructions

This file is a bootstrap, not the project contract.

When working inside a repository:

1. Look for `AGENTS.md` at the workspace root.
2. If `AGENTS.md` exists, follow it as the project operating contract.
3. If `AGENTS.md` does not exist, look for `.github/copilot-instructions.md`.
4. If neither exists and the user is starting meaningful project work, ask
   whether to install or create project-level instructions.

Project-level files always override this global default.

Use global skills as reusable capabilities, but rely on the local `AGENTS.md`
for repository-specific paths, workflows, commands, and artifact locations.
