# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

This changelog is automatically maintained by [release-please](https://github.com/googleapis/release-please).

## [1.5.0](https://github.com/woliveiras/geremmyas/compare/v1.4.0...v1.5.0) (2026-05-21)


### Features

* add LangChain, LangGraph, and LLM service instruction files; update FastAPI conventions and detection logic ([826dcf5](https://github.com/woliveiras/geremmyas/commit/826dcf502639faf6838da8a7f8cded7eacc8eb8f))
* add NestJS and Fastify instruction files; update detection logic for relevant instructions ([37c6234](https://github.com/woliveiras/geremmyas/commit/37c6234c35ec18670364fa83ddb04273fa97b1c3))
* add new instruction files for Echo, Go Embed, and Air; update README and install script for detection logic ([b185285](https://github.com/woliveiras/geremmyas/commit/b1852857123d493c35eebda81d1bbad018e4de64))
* add new skills for Terraform, GCloud, CI/CD workflows, LangGraph agent design, LLM integration review, PostgreSQL query review, Supabase workflow, and ChromaDB RAG workflow ([24cb506](https://github.com/woliveiras/geremmyas/commit/24cb506e653fa70a212babcb35e4d2c97109d4f4))
* add new workflow skills for Terraform, GCloud, CI/CD, LLM integration, LangGraph design, Supabase, PostgreSQL, and ChromaDB ([e3e718a](https://github.com/woliveiras/geremmyas/commit/e3e718a166847b6acb008a814ce56fcbe1730de9))
* add PostgreSQL, ChromaDB, and Supabase instruction files; update detection logic for relevant instructions ([271009d](https://github.com/woliveiras/geremmyas/commit/271009de43fd96291d214bd7404410b558f7b4d6))
* add Pydantic instruction file; update detection logic to include Pydantic instructions ([daf3d0b](https://github.com/woliveiras/geremmyas/commit/daf3d0b1c095cc967bbf2c66bcfdd8ab2389571e))
* add SQLite instruction files for Node.js, Go, Python, and Android; update existing SQLite conventions ([7a6a70f](https://github.com/woliveiras/geremmyas/commit/7a6a70fc6854c7589dc03f383ccb8a1833ee2478))
* enhance instruction files for Docker, GitHub Actions, GCP, and Terraform; add new conventions and update detection logic ([33b062c](https://github.com/woliveiras/geremmyas/commit/33b062c487356f290c8c1acac74d0d422676e10d))
* remove tasks-tech-stack.md to streamline project structure ([0b13c90](https://github.com/woliveiras/geremmyas/commit/0b13c90c2d0a6171731014b86389debfd805b4d2))

## [1.4.0](https://github.com/woliveiras/geremmyas/compare/v1.3.0...v1.4.0) (2026-05-21)


### Features

* add AGENTS.md for project-level agent workflows and instructions; update global Copilot instructions for clarity ([d726e99](https://github.com/woliveiras/geremmyas/commit/d726e994d64507b22dc464e81548fc8b299bee3f))
* add bugfix loop skill documentation with process and template ([6f3f9eb](https://github.com/woliveiras/geremmyas/commit/6f3f9ebe47a4377921a6d4fe0c145b13c3c8b375))
* add generate-adr skill for creating Architectural Decision Records in MADR format ([16e69d3](https://github.com/woliveiras/geremmyas/commit/16e69d364435941ca75cdd978bf0ab6fbb4afd67))
* add generate-glossary skill and template for formalizing domain terminology ([b094309](https://github.com/woliveiras/geremmyas/commit/b0943097c382f60d87d2d27c253ed1e57e9bf489))
* add generate-spec skill and template for structured feature specifications ([29d8ead](https://github.com/woliveiras/geremmyas/commit/29d8ead43ce2b6016c05d72920bc4ac5ad904c14))
* add generate-tests-from-spec skill for generating unit tests from specifications ([701f4ee](https://github.com/woliveiras/geremmyas/commit/701f4eeb5e63d1ea698a3fdd6b6e63e39e7cd246))
* add SKILL.md for conducting requirements interviews to clarify product and technical ambiguity ([372bfd8](https://github.com/woliveiras/geremmyas/commit/372bfd88ffafa2bba0611a469ef6faef66e31528))
* add SKILL.md for generating git commit messages in Conventional Commits format ([d92a8f6](https://github.com/woliveiras/geremmyas/commit/d92a8f6c9d591e52ad50b843007e741dbd32408e))
* add SKILL.md for managing state with Zustand ([9130b99](https://github.com/woliveiras/geremmyas/commit/9130b996e6a4530a46d9350b533707517cdb08ab))
* add SKILL.md for migrating from React Router v6 to v7 ([9130b99](https://github.com/woliveiras/geremmyas/commit/9130b996e6a4530a46d9350b533707517cdb08ab))
* add SKILL.md for modeling state with XState ([9130b99](https://github.com/woliveiras/geremmyas/commit/9130b996e6a4530a46d9350b533707517cdb08ab))
* add SKILL.md for safe Git commit process with Conventional Commits format ([55e1d5e](https://github.com/woliveiras/geremmyas/commit/55e1d5e1fab5a6f0e23b7f86bc2c2a13fa3027c3))
* add SKILL.md for session handoff process to streamline transitions between Copilot sessions ([4ef51ce](https://github.com/woliveiras/geremmyas/commit/4ef51ce78957b36c6dd9c1ef5c19a1001b8a0a6e))
* add SKILL.md for skill authoring process to guide creation and revision of Copilot skills ([5d424c5](https://github.com/woliveiras/geremmyas/commit/5d424c5860534e4ab5bac721be7ab4fc46179e7b))
* add SKILL.md for task breakdown process to guide creation of tasks from PRD, spec, or plan ([92522f7](https://github.com/woliveiras/geremmyas/commit/92522f78385b470638c3ad59248ba18261ceddf7))
* add SKILL.md for updating project documentation after feature implementation ([7bf3ad0](https://github.com/woliveiras/geremmyas/commit/7bf3ad04535a5847830c65a20c5f678ab06790c9))
* add SKILL.md for validating with Zod ([9130b99](https://github.com/woliveiras/geremmyas/commit/9130b996e6a4530a46d9350b533707517cdb08ab))
* add SKILL.md for vertical TDD process to guide test-driven development ([e63f7bb](https://github.com/woliveiras/geremmyas/commit/e63f7bb1b32f4939032270a9f5cd5557efced7cb))
* enhance installation script to support exclusion of specific instruction files and improve relevant instruction detection ([1bd553b](https://github.com/woliveiras/geremmyas/commit/1bd553b6324f690913da7924f503fe19f2b1e9d8))
* enhance installer to support skill dependencies and improve instruction installation logic ([91538bf](https://github.com/woliveiras/geremmyas/commit/91538bf2fd4217afe02d227869c95f8b8aef9582))
* enhance reviewer and review prompts for clarity and spec alignment ([f283cb8](https://github.com/woliveiras/geremmyas/commit/f283cb8ea252a9092daebe395bab9b9b3ff73bf5))
* enhance security and testing documentation; add specific guidelines for Android, API, E2E, integration, and web security; update Python instructions and remove obsolete security file ([d6eeba4](https://github.com/woliveiras/geremmyas/commit/d6eeba45a879edbc255df854de3748cb4a04536d))
* refine spec-writing agent description and process for clarity and alignment with workflows ([99e814b](https://github.com/woliveiras/geremmyas/commit/99e814bcde1a0b48eaefa99b18ae3543a8163891))
* remove deprecated skills and templates from project ([3c91d18](https://github.com/woliveiras/geremmyas/commit/3c91d186b93163309c2e5c577de7135a428e05b1))
* remove mise.toml handling from project installation process and update README instructions ([4817aea](https://github.com/woliveiras/geremmyas/commit/4817aea418e5f5fdd11837b48cc42cf4976ea3eb))
* update agent documentation to reflect new skill formats and improve clarity ([c721fa5](https://github.com/woliveiras/geremmyas/commit/c721fa5a738f8abd6e8c644516ffa1e134b0559a))
* update skill references in documentation for React Router, XState, Zod, and Zustand ([41afcea](https://github.com/woliveiras/geremmyas/commit/41afcea586b0ccc19e57444d5cc55e5baaa30112))
* update SKILL.md files to clarify skill usage and improve documentation ([0807703](https://github.com/woliveiras/geremmyas/commit/08077039d6c5544633d75c39b51966e10226e682))


### Bug Fixes

* correct skill names in review and SDD prompts for consistency ([52ce5f8](https://github.com/woliveiras/geremmyas/commit/52ce5f8b86f247f414067a61dac1449631b2e363))
* update descriptions in architect agent and generate-adr skill for clarity and consistency ([8b2e872](https://github.com/woliveiras/geremmyas/commit/8b2e872f76b3db347fd3a76a10ffcb405400536c))

## [1.3.0](https://github.com/woliveiras/geremmyas/compare/v1.2.0...v1.3.0) (2026-04-24)


### Features

* add ADR template skill for documenting architectural decisions ([14a12e2](https://github.com/woliveiras/geremmyas/commit/14a12e27c75d004dd2eb8d8ac4234b5110f1a4b5))
* add comprehensive guidelines for React Router v7 in framework mode ([2ce2b20](https://github.com/woliveiras/geremmyas/commit/2ce2b205fba7bf4aaf4964ac5d948ca1574bcd1c))
* add detailed instructions for XState and Zustand patterns, including React integration and testing guidelines ([da8bb06](https://github.com/woliveiras/geremmyas/commit/da8bb06fe7c710f7773802aa9628c34891f675d8))
* add guidelines for Conventional Commits format and usage ([967f849](https://github.com/woliveiras/geremmyas/commit/967f849ac74a50566ee1e1b59bc5f57faed3a5a9))
* add migration guide for React Router v6 to v7 framework mode ([b44cf07](https://github.com/woliveiras/geremmyas/commit/b44cf0780b8f5586b41e0af093f715d885ce67e2))
* add RFC template for structured documentation ([668f036](https://github.com/woliveiras/geremmyas/commit/668f036a3229fa8054372816c1552d4ad22a8528))
* add Spec Driven Development (SDD) workflow section to README with detailed steps and prompt usage ([d6f2616](https://github.com/woliveiras/geremmyas/commit/d6f261616b30d4d3c93026c54ac05bece341dee6))
* add TanStack Query v5 guidelines covering hooks, queries, and mutations ([a3cf3d5](https://github.com/woliveiras/geremmyas/commit/a3cf3d5fd4b084ce9fc693f33b375f09317826c0))
* add Zod patterns skill documentation for validation use cases ([7cd873f](https://github.com/woliveiras/geremmyas/commit/7cd873f859e2cc67f1a462945ae1c128b5cdc44b))
* enhance FastAPI guidelines with async patterns and dependency injection practices ([b247f16](https://github.com/woliveiras/geremmyas/commit/b247f168a9697efd7aeb3c924fa87ee87ae7d26a))
* enhance installation script with interactive configuration options for copilot-instructions.md ([c558513](https://github.com/woliveiras/geremmyas/commit/c55851384c8e60b1dfebb775d5272cf69c7b1afd))
* enhance testability section in reviewer agent with additional checks and references ([424e733](https://github.com/woliveiras/geremmyas/commit/424e7339a3da4ea3340a0e3cee7316f69241089e))
* update installation and uninstallation instructions in README for clarity and detail ([fb2cd3f](https://github.com/woliveiras/geremmyas/commit/fb2cd3f5fcd04d6b79cd8304cb33aff7b26583c7))
* update installation script to preserve customizable files during updates and clarify force overwrite behavior in README ([9ad625f](https://github.com/woliveiras/geremmyas/commit/9ad625f23c0a3ead626ddd5e4c8d0ecbf97e31cc))
* update RFC writing instructions to include rfc-template skill format for quick generation ([5597666](https://github.com/woliveiras/geremmyas/commit/559766692ad80e45eee65ddeee9ad1077bd02b07))
* update skills and prompts count in README for accuracy ([9597c6f](https://github.com/woliveiras/geremmyas/commit/9597c6f0c947bb4effa61b02624e6ceb6b3a4b22))
* update skills section to reflect additional templates and new RFC and Zod patterns entries ([b79ccc6](https://github.com/woliveiras/geremmyas/commit/b79ccc691601aa1998203156c38f5730ed7c9d20))


### Bug Fixes

* update guideline for unexported package-level variable naming ([f1d8dfd](https://github.com/woliveiras/geremmyas/commit/f1d8dfd093ff9776c18b227b01df1591d8a1897d))

## [1.2.0](https://github.com/woliveiras/geremmyas/compare/v1.1.0...v1.2.0) (2026-04-17)


### Features

* add comprehensive Docker instructions for best practices and security ([da5949f](https://github.com/woliveiras/geremmyas/commit/da5949f1413a2d004968c6a1480b7f3f0e925644))
* add comprehensive SQLite best practices and guidelines for schema design, migrations, and performance optimization ([cd820da](https://github.com/woliveiras/geremmyas/commit/cd820da2f2bc2510d070e8f49aa0edc7a3781dbe))
* add Tailwind CSS v4 conventions and patterns documentation ([2845238](https://github.com/woliveiras/geremmyas/commit/2845238d90c89472c88bdaed7cd8153dc3ee2957))
* add XState v5 conventions and guidelines for state machines, testing, and React integration ([06debd8](https://github.com/woliveiras/geremmyas/commit/06debd842f4e830b1f3f58b78fa35c7da9ca7265))
* add Zod v4 conventions and guidelines for schema definition, validation patterns, and runtime validation ([a4cd2f8](https://github.com/woliveiras/geremmyas/commit/a4cd2f898ac0c4fd033131774ea582b4ef2924ce))
* add Zustand v5 conventions and guidelines for store structure, middleware, mutations, persistence, and selectors ([7652540](https://github.com/woliveiras/geremmyas/commit/7652540bfbc5cf8d52ad694d0ac567193e47b410))
* enhance Go code quality guidelines with detailed sections on concurrency, structs, functions, performance, testing, linting, and anti-patterns ([9cd86ad](https://github.com/woliveiras/geremmyas/commit/9cd86ad92420975b1a6d7713af2d2240dabc4ec5))

## [1.1.0](https://github.com/woliveiras/geremmyas/compare/v1.0.0...v1.1.0) (2026-04-10)


### Features

* update CI and release workflows to ignore specific paths on push ([79dff1e](https://github.com/woliveiras/geremmyas/commit/79dff1e0a5a697cd708ab3beb9299e7c7f9fae32))

## 1.0.0 (2026-04-10)


### Features

* add .gitignore file to exclude OS and editor-specific files ([60141ca](https://github.com/woliveiras/geremmyas/commit/60141ca3fe48485986b3cc17d5c901fc429025c3))
* add architecture improvement, code explorer, and code reviewer agents ([bc1cbf8](https://github.com/woliveiras/geremmyas/commit/bc1cbf888dbf56a9daac88cbfaab3c075f4406c3))
* add bug report template to facilitate issue reporting and improve troubleshooting ([f9b8df3](https://github.com/woliveiras/geremmyas/commit/f9b8df3d94da9091e47d8250bf5588573ede57ad))
* add CI workflow for shell script linting and installation testing ([5ad0e92](https://github.com/woliveiras/geremmyas/commit/5ad0e928598309c49012135dd4aa653ab6fc6404))
* add CONTRIBUTING.md to guide contributions and improve project quality ([b03c415](https://github.com/woliveiras/geremmyas/commit/b03c4154f8e60c8b77ec573569f79118dfd5257f))
* add Contributor Covenant Code of Conduct to promote a respectful community ([d77bb51](https://github.com/woliveiras/geremmyas/commit/d77bb513f32621e448ee4ffde4b4a09feee25caf))
* add documentation and templates for glossary, spec generation, and unit testing ([2a362c6](https://github.com/woliveiras/geremmyas/commit/2a362c62939c1dc031068c13ef27aa3ed9d3afa7))
* add feature request template to streamline suggestions and enhance clarity ([18eb98c](https://github.com/woliveiras/geremmyas/commit/18eb98cd7c2edb7892b0405e21a3b3bce007027e))
* add guardrails for command safety and implement blocking script ([e5cfd50](https://github.com/woliveiras/geremmyas/commit/e5cfd509e5ed1aba8dc0f5b7a9fae0c03cfcf150))
* add guidelines for various programming languages including Go, Python, TypeScript, Kotlin, React, and security best practices ([0b1ce70](https://github.com/woliveiras/geremmyas/commit/0b1ce701ab980a37a9210f0679fa6ad174b5ce75))
* add initial release configuration and changelog for project ([e0bef0e](https://github.com/woliveiras/geremmyas/commit/e0bef0e17d5fe2be64499c1f06243473913b508c))
* add installer and uninstaller scripts for geremmyas ([c8ff593](https://github.com/woliveiras/geremmyas/commit/c8ff593968ff8472f1eba9b34e273eb57629f7bd))
* add mise.toml configuration file for tool version management ([fd1fdeb](https://github.com/woliveiras/geremmyas/commit/fd1fdebb9be782a78a1fe7d40943bd39b199e713))
* add MIT License to the project for legal clarity and usage rights ([925513e](https://github.com/woliveiras/geremmyas/commit/925513e8a4a9b403f505b7c4fab07f728e44f251))
* add pull request template to standardize contributions and improve clarity ([a37057d](https://github.com/woliveiras/geremmyas/commit/a37057dc461fc2857bb00744a009a66f9559f74e))
* add reference documents for complexity signals, deep modules, dependency categories, interface design, pragmatic heuristics, and seam finding ([c32dbec](https://github.com/woliveiras/geremmyas/commit/c32dbec9507566d2ae4bebc8eeb018ac28e2b402))
* add spec writer agent for structured feature specifications ([664664c](https://github.com/woliveiras/geremmyas/commit/664664cc24da25321181dd82b2ba2f2bcbec5faf))
* add structured prompts for code refactoring, review, and testing ([3d28698](https://github.com/woliveiras/geremmyas/commit/3d2869874a5aeef558327d60678d3adf60f7efba))
* enhance CI workflow with simulated global install and improve command parsing in hook script ([50989fa](https://github.com/woliveiras/geremmyas/commit/50989fae492fe3d09b836be6ce4c3ae38e9847b2))

## [0.1.0](https://github.com/woliveiras/geremmyas/releases/tag/v0.1.0) — 2026-04-10

### Features

- Install script (`install.sh`) with OS detection (macOS/Linux), `--project` and `--force` flags
- Uninstall script (`uninstall.sh`) for global config removal
- 3 global user prompts: review, refactor, test
- Project template with `copilot-instructions.md` (Spec Driven Development workflow)
- 9 instruction files: TypeScript, Go, Python, Kotlin, React, Astro/MDX, testing, security, conventional commits
- 4 agents: spec-writer, explorer, reviewer, architect
- 6 agent references: deep modules, interface design, complexity signals, dependency categories, pragmatic heuristics, seam finding
- 4 skills: spec-template, test-from-spec, doc-updater, glossary
- Command guardrails (hooks) with configurable BLOCK/ASK rules
- mise.toml template for tool version management
- GitHub Actions CI: ShellCheck, structure validation, install tests (Ubuntu + macOS)
- Contributing guide, Code of Conduct, issue and PR templates
