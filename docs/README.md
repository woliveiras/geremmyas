# geremmyas documentation

Human-oriented docs for the CLI and the maintainer repository. Agent workflow
rules live in [`AGENTS.md`](../AGENTS.md) (symlink to [`project/AGENTS.md`](../project/AGENTS.md)).

| Document | Audience | Contents |
| --- | --- | --- |
| [architecture.md](architecture.md) | Contributors | Embed FS, packs, sync, **multi-IDE targets**, global install, repo symlinks |
| [creating-packs.md](creating-packs.md) | Contributors | Add packs, skills, and instructions to the catalog |

Consumer projects use the SDD template installed by the `sdd` pack:
`specs/README.md` at the project root (see [`project/templates/specs/README.md`](../project/templates/specs/README.md)).

**Copilot instructions:** [`project/.github/copilot-instructions.md`](../project/.github/copilot-instructions.md)
is the synced template. Maintainer-only context:
[`copilot-instructions.geremmyas.md`](../project/.github/copilot-instructions.geremmyas.md).
