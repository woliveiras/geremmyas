# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

This changelog is automatically maintained by [release-please](https://github.com/googleapis/release-please).

## [3.1.0](https://github.com/woliveiras/geremmyas/compare/v3.0.0...v3.1.0) (2026-06-23)


### Features

* add instructions to update tasks and reconcile plans before completion ([ccdb657](https://github.com/woliveiras/geremmyas/commit/ccdb657637de314bfc38676e2928a7b85341666a))
* add lint command over skill tree ([1c4b27f](https://github.com/woliveiras/geremmyas/commit/1c4b27f119940cf8037335ac8668bc9034fe36d3))
* add skill description lint engine ([9f96fc4](https://github.com/woliveiras/geremmyas/commit/9f96fc461f5d4c5d87e56bfb4cfab81d3e1a2fe5))
* add skill name and body lint rules ([c9cf5c9](https://github.com/woliveiras/geremmyas/commit/c9cf5c941d56127725caeffb594023ce2d6180d2))
* **agents:** add guardrails framework for error prevention and decision-making ([24c83de](https://github.com/woliveiras/geremmyas/commit/24c83dedd8a59397c42bb543280f4d7df84aaeff))
* **cli:** surface codex in help, usage, and destination summary ([deee9f6](https://github.com/woliveiras/geremmyas/commit/deee9f689a2b75b473dc4a391adacfedd3a2e601))
* **codex:** generate codex document at global scope ([e1a1ebe](https://github.com/woliveiras/geremmyas/commit/e1a1ebec864029bb86cbf41c1323fbe0922e2778))
* **codex:** generate codex document at project scope ([7dfa173](https://github.com/woliveiras/geremmyas/commit/7dfa1733a9d6a7efd20cf9db464ef8e95dd387fc))
* enhance reconciliation process for task completion in vertical TDD ([dc0249e](https://github.com/woliveiras/geremmyas/commit/dc0249e8e047cebbe91e002a7af45447e35bffd3))
* **guardrails:** implement comprehensive guardrails framework for AI coding agents ([09170e4](https://github.com/woliveiras/geremmyas/commit/09170e4858072d214b43a269f84475bd98849a18))
* **readme:** update documentation to include guardrails framework and its error-prevention features ([0501174](https://github.com/woliveiras/geremmyas/commit/05011744a522723c63a337783fc7d83a13afbe27))
* remove obsolete specs and tasks for dashboard features ([f384c2a](https://github.com/woliveiras/geremmyas/commit/f384c2ab21f62e34475fa81f92b58c5bb22eaff3))
* report missing skill files in lint ([ab51eb9](https://github.com/woliveiras/geremmyas/commit/ab51eb9c0e3801794187d5a01a0eb8ba14ca9b8d))
* **skills:** add abort criteria guidelines for task management ([c8c9a87](https://github.com/woliveiras/geremmyas/commit/c8c9a87de00a1d34e723570718b90d62e3cfd5c5))
* **skills:** add agent rationalization blocking guidelines ([a194a4c](https://github.com/woliveiras/geremmyas/commit/a194a4c1860d3750300dee81dc6861232f81d20b))
* **skills:** add approval gates before implementation guidelines ([fe81e79](https://github.com/woliveiras/geremmyas/commit/fe81e79b8b680889fb077036abc4c6b56978290f))
* **skills:** add comprehensive regression testing guidelines ([741b0f1](https://github.com/woliveiras/geremmyas/commit/741b0f147595674dc84babd3ca21b7c2ed6fbf25))
* **skills:** add decision framework for structured decision-making ([8694d58](https://github.com/woliveiras/geremmyas/commit/8694d58220c0b5dc1267119c585320e377be1a66))
* **skills:** add structured workflow for code review requesting ([133856a](https://github.com/woliveiras/geremmyas/commit/133856a405f83e90fb8a28f6b3520c147a0c4d62))
* **skills:** add subagent selection decision framework and guidelines ([0641b4a](https://github.com/woliveiras/geremmyas/commit/0641b4a8bde95e592fab33ca0a6717d26262649b))
* **skills:** add verification checklists to enforce evidence before task completion ([e4762a1](https://github.com/woliveiras/geremmyas/commit/e4762a15ed7821e43fec233334a66b8f6b74f017))
* **skills:** refine descriptions for various skills to enhance clarity and usage guidelines ([f669575](https://github.com/woliveiras/geremmyas/commit/f669575c9b6fda7985377755ef99ee28248d91c1))
* **targets:** add codex as valid target ([ffb2a4c](https://github.com/woliveiras/geremmyas/commit/ffb2a4c0178156b8dd672955f03a668c1f8ee8ca))
* **targets:** test mixed-target and force-overwrite behavior ([d204fa3](https://github.com/woliveiras/geremmyas/commit/d204fa3f9421bfae94e8757a314046faa9f729ef))


### Bug Fixes

* correct formatting of missing skill file violation constant ([4ebe744](https://github.com/woliveiras/geremmyas/commit/4ebe744db6df2b640696b468139501b4d2236ac7))


### Documentation

* **specs:** mark spec 0001 as Completed ([33a7868](https://github.com/woliveiras/geremmyas/commit/33a786886988d96a33cadbad73c018ac81adf4ae))


### Miscellaneous

* **docs:** update readme.md ([5aaf6f0](https://github.com/woliveiras/geremmyas/commit/5aaf6f09157edce539adeaf3071dfab1cb0c2c58))
* **skills:** refine skill descriptions across multiple files to clarify usage contexts and avoid misuse ([26ac4f9](https://github.com/woliveiras/geremmyas/commit/26ac4f9a6aef8d0f5e5e214995403707a43c9f3e))

## [3.0.0](https://github.com/woliveiras/geremmyas/compare/v2.8.0...v3.0.0) (2026-06-04)


### ⚠ BREAKING CHANGES

* update geremmyas workflow

### Features

* add PRD for Dashboard to improve spec management and visibility ([9103136](https://github.com/woliveiras/geremmyas/commit/9103136d5ebdeb6fc22f4bb90316aef427dac6d4))
* add Rust skills ([ad8e8aa](https://github.com/woliveiras/geremmyas/commit/ad8e8aa3830658ecc79b104aff82eecc06ac4a86))
* **dashboard:** add spec parser and data model (spec 0001) ([cf8063c](https://github.com/woliveiras/geremmyas/commit/cf8063c531157d148c67cf0a12654849be65ace9))
* **dashboard:** enhance dashboard templates and styles for improved UI ([17ce4eb](https://github.com/woliveiras/geremmyas/commit/17ce4ebf929379fda5e12bb17b7dc1d0540d8245))
* **dashboard:** enhance spec scanning and README generation ([6fba246](https://github.com/woliveiras/geremmyas/commit/6fba24649c64531ea4e2b4a70e538df181317d7f))
* **dashboard:** enhance WatchDirs to return stop function and add tests ([a5380b7](https://github.com/woliveiras/geremmyas/commit/a5380b72b7c7bfbf1ae4adc950480c054cfccb4a))
* **dashboard:** git dates, cache, and metrics charts (spec 0004) ([3bd6c76](https://github.com/woliveiras/geremmyas/commit/3bd6c762cb7cddf40cfe91ab3e6efe45f7b030f0))
* **dashboard:** overview HTML and CLI command (spec 0002) ([54652b9](https://github.com/woliveiras/geremmyas/commit/54652b9373a9dff888af45362c478a2168fdcf75))
* **dashboard:** per-family kanban board view (spec 0003) ([1068f39](https://github.com/woliveiras/geremmyas/commit/1068f39b92a2dcd02c984c946e2ed050e0706629))
* **dashboard:** PRD/bugfix linking and dependency navigation (spec 0006) ([aa0e6a0](https://github.com/woliveiras/geremmyas/commit/aa0e6a0b8854549b6ed72da3eabb60f687333fe7))
* **dashboard:** serve and watch mode with SSE reload (spec 0005) ([ffab6a2](https://github.com/woliveiras/geremmyas/commit/ffab6a2b2902456c5aded0879acac34b5dcdb5cb))
* enhance Android instructions with security guidelines and add CI/CD setup skill ([4bd8996](https://github.com/woliveiras/geremmyas/commit/4bd8996d516a0bc75da3fadf1d4a73bb06e533fd))
* enhance dashboard serve command to support full regeneration on file changes and update dependencies ([eac2f8f](https://github.com/woliveiras/geremmyas/commit/eac2f8fb12d010bc3b3c5f676fdadf5bb3e583fd))
* enhance global install command with multi-target support and update documentation ([14a20d3](https://github.com/woliveiras/geremmyas/commit/14a20d311692b3010e8edbd2bbf97e8aae25752c))
* enhance Python instructions with security guidelines and add CI setup skill ([aeaae7e](https://github.com/woliveiras/geremmyas/commit/aeaae7ed4b35ca5b9d16b3941e48a138fc64e557))
* enhance TypeScript instructions with security guidelines and add CI setup skill ([bc783cb](https://github.com/woliveiras/geremmyas/commit/bc783cb2928a10dd16de0611b70d5b00c3c6adce))
* implement dashboard parser and overview generation with dependency tracking ([88bdd91](https://github.com/woliveiras/geremmyas/commit/88bdd91285d991d6bc389c80b74be73d6dfeba5f))
* introduce multi-IDE target support and enhance global install functionality ([93f09d4](https://github.com/woliveiras/geremmyas/commit/93f09d426e11934d09d68637890d94be8494ae0e))
* update geremmyas workflow ([5beba39](https://github.com/woliveiras/geremmyas/commit/5beba3995c6b73a41f91242ecb6e814efab1a6c5))
* update Go instructions and add Go CI setup skill; remove Rust review skill ([0e76cc8](https://github.com/woliveiras/geremmyas/commit/0e76cc8c6cf2ac7a96ae53a8ab09c87d9813fb26))


### Bug Fixes

* **dashboard:** update SSE reload logic to trigger on specific event ([eee809e](https://github.com/woliveiras/geremmyas/commit/eee809ed657d548f1b58859f0a30eb370fa16a24))


### Documentation

* create docs ([bc9e60d](https://github.com/woliveiras/geremmyas/commit/bc9e60d7d421d8f2cb127804299196577ed726a3))
* document dashboard command and mark specs implemented ([bd0d9b3](https://github.com/woliveiras/geremmyas/commit/bd0d9b392be0698d4d527bc0a9e60679bec8eb02))
* update dashboard specs status to Implemented ([3961cd2](https://github.com/woliveiras/geremmyas/commit/3961cd25f2c7fe8ff2aa230c584333cda6a6f38f))


### Miscellaneous

* add caveman ([c2e0610](https://github.com/woliveiras/geremmyas/commit/c2e0610e66dc75d78f79c6370e26b84f0a2d0cf0))
* update .gitignore to include additional dashboard cache and asset paths ([54d9fb9](https://github.com/woliveiras/geremmyas/commit/54d9fb9647736e5ae55ff1453e944402e7f0ed56))
* update .gitignore to include geremmyas build directory ([68156b0](https://github.com/woliveiras/geremmyas/commit/68156b08d4239512b2e23c4feaa4f71689497063))
* update dashboard specs status from Draft to Approved for all phases ([2c9e50f](https://github.com/woliveiras/geremmyas/commit/2c9e50f6d92aede98224c809e726bfddaaef24b6))


### Refactoring

* streamline target handling by introducing applyTargetsFlag function and add unit tests for validation ([5858e04](https://github.com/woliveiras/geremmyas/commit/5858e04d4d2b7881e089bc3cbfada23954c3d9ae))

## [2.8.0](https://github.com/woliveiras/geremmyas/compare/v2.7.0...v2.8.0) (2026-05-25)


### Features

* add premortem skill for stress-testing plans and decisions ([0b9aa73](https://github.com/woliveiras/geremmyas/commit/0b9aa7322887ec67ab5cff82bd62e65d288f259f))
* restructure GitHub directory to link project resources ([c9efd47](https://github.com/woliveiras/geremmyas/commit/c9efd47258b7db7969b2c2940b15b96d77bbc512))

## [2.7.0](https://github.com/woliveiras/geremmyas/compare/v2.6.0...v2.7.0) (2026-05-25)


### Features

* add daily workflow for scrum standup and meeting notes ([2340295](https://github.com/woliveiras/geremmyas/commit/23402959dfe44a683aa6d5bc48f750d6233f4f36))

## [2.6.0](https://github.com/woliveiras/geremmyas/compare/v2.5.0...v2.6.0) (2026-05-25)


### Features

* add daily skill and template for recording work progress and standup notes ([3422c1a](https://github.com/woliveiras/geremmyas/commit/3422c1ab966bf7c0a37ec9945440678e0f63738a))

## [2.5.0](https://github.com/woliveiras/geremmyas/compare/v2.4.0...v2.5.0) (2026-05-25)


### Features

* add 'brag-me' pack for generating demo presentations with Reveal.js ([3c1b33a](https://github.com/woliveiras/geremmyas/commit/3c1b33a6ba5d4ea177e31c69e0c1a49a38beb828))
* add logo to README for enhanced branding ([4f15cc3](https://github.com/woliveiras/geremmyas/commit/4f15cc3083476db1f4f15078e114f878e694cdad))

## [2.4.0](https://github.com/woliveiras/geremmyas/compare/v2.3.0...v2.4.0) (2026-05-22)


### Features

* add 'project' command to install packs and update configuration, with interactive selection and file preservation options ([e35f136](https://github.com/woliveiras/geremmyas/commit/e35f1367a58a7ff6eeaf728ebc2d56a8f7a90cd5))
* add systematic review protocol, threats to validity framework, venue submission checklists, and bibliography validation script ([0e3b0cc](https://github.com/woliveiras/geremmyas/commit/0e3b0cc25284652ba446dacd141d42558464f74a))

## [2.3.0](https://github.com/woliveiras/geremmyas/compare/v2.2.0...v2.3.0) (2026-05-22)


### Features

* update global installation paths for skills and instructions ([06dcf61](https://github.com/woliveiras/geremmyas/commit/06dcf6174aa8a6411892264d79eaea7d721961d1))


### Miscellaneous

* add symlink for project AGENTS.md and .github directory ([d0539fd](https://github.com/woliveiras/geremmyas/commit/d0539fd4ac2aca7c63f8e4e037213b26c5390deb))

## [2.2.0](https://github.com/woliveiras/geremmyas/compare/v2.1.0...v2.2.0) (2026-05-22)


### Features

* enhance runRemove with interactive mode and add isGlobalTarget function ([98745df](https://github.com/woliveiras/geremmyas/commit/98745df311d158b7fe37f8d9dfdba31b62ac15fb))


### Documentation

* add mise.toml and update contributing guide ([639c28a](https://github.com/woliveiras/geremmyas/commit/639c28a81d9e8cb9d78e13128be0611fb0f9f810))
* document global install and mise environment setup ([06a6a43](https://github.com/woliveiras/geremmyas/commit/06a6a438f4785b1eccdfa4d9ccebbd407c9fe4b6))

## [2.1.0](https://github.com/woliveiras/geremmyas/compare/v2.0.2...v2.1.0) (2026-05-22)


### Features

* add interactive init and global pack install ([7473b78](https://github.com/woliveiras/geremmyas/commit/7473b7805a464ad0ae6e7797065d8322ce9682f2))

## [2.0.2](https://github.com/woliveiras/geremmyas/compare/v2.0.1...v2.0.2) (2026-05-22)


### Bug Fixes

* build and upload binaries in release workflow ([c6e0142](https://github.com/woliveiras/geremmyas/commit/c6e014205ffcedaf10194ffd82cabf93ae62cceb))
* correct install examples in README ([49928ff](https://github.com/woliveiras/geremmyas/commit/49928ff7325c5db83b3e4491e455ed466338610a))
* show mise trust hint after sync installs mise.toml ([a3e1dbd](https://github.com/woliveiras/geremmyas/commit/a3e1dbd1fb67e40d38688c0f79e9d303154528e3))

## [2.0.1](https://github.com/woliveiras/geremmyas/compare/v2.0.0...v2.0.1) (2026-05-22)


### Bug Fixes

* allow release workflow to publish release-please releases ([4f34c77](https://github.com/woliveiras/geremmyas/commit/4f34c77bfa9766505e4a977f817b243678530d89))
* update permissions and token configuration in release workflow ([ff8fa4c](https://github.com/woliveiras/geremmyas/commit/ff8fa4c17bb23a4d28c3a3c1ecb943384c4e7bf2))
* use manifest mode for release-please action ([225ae5f](https://github.com/woliveiras/geremmyas/commit/225ae5f15decee725d8d751b0707a9c6d259a196))

## [2.0.0](https://github.com/woliveiras/geremmyas/compare/v1.5.0...v2.0.0) (2026-05-22)


### ⚠ BREAKING CHANGES

* geremmyas is now the primary CLI and the shell installer only manages the geremmyas binary.

### Features

* add AFK task triage and agent brief skills, update task breakdown and agents documentation ([6e96b1a](https://github.com/woliveiras/geremmyas/commit/6e96b1ae035f88b27d516d8ba46a75721111e4e4))
* enhance AFK task guidelines and clarify agent brief requirements ([57c7dca](https://github.com/woliveiras/geremmyas/commit/57c7dca2059363aee9811f5b98cbaacac9fe8307))
* release geremmyas 2.0 ([8afed42](https://github.com/woliveiras/geremmyas/commit/8afed42bfaf0457df6eb9bf61a6979c6480b1009))
* update documentation for React Router, Tailwind CSS, XState, Zod, and Zustand conventions ([9aab9f9](https://github.com/woliveiras/geremmyas/commit/9aab9f9522905460eed351f5eb7cc715c22873b3))


### Bug Fixes

* remove unnecessary prefix from project description in README ([ecc18b6](https://github.com/woliveiras/geremmyas/commit/ecc18b6716294215eea412eb6e130ca99748de92))

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
