# geremmyas documentation

Human-oriented docs for the CLI and the maintainer repository. Agent workflow
rules live in [`AGENTS.md`](../AGENTS.md) (symlink to
[`content/AGENTS.md`](../content/AGENTS.md)).

| Document | Audience | Contents |
| --- | --- | --- |
| [architecture.md](architecture.md) | Contributors | Embed FS, packs, sync, **multi-IDE targets**, global install, repo symlinks |
| [creating-packs.md](creating-packs.md) | Contributors | Add packs, skills, and instructions to the catalog |

Consumer projects use the specification template installed by the `coding` pack:
`specs/README.md` at the project root (see
[`content/templates/specs/README.md`](../content/templates/specs/README.md)).

Assistant-neutral sources live in [`content/`](../content/); assistant-specific
adapters and templates live in [`targets/`](../targets/).
