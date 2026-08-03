# Plan: Modular game-development routing

Spec: [spec.md](./spec.md)

Status: In Progress

## Approach

Lock the desired catalog topology and routing signals in failing tests first.
Then split the catalog, narrow the eleven skill descriptions, update public
documentation, and run full Geremmyas verification. Finally build the local CLI,
reconcile Rusted to its focused packs, and record a real Codex routing smoke.

## Sequencing

1. Add catalog tests for seven focused packs, the complete metapack, and the
   backward-compatible art pack name.
2. Add a versioned routing corpus and lexical conformance test covering every
   game skill.
3. Implement the pack topology and discriminating frontmatter descriptions.
4. Update README and architecture guidance for focused project selection.
5. Verify Geremmyas, then materialize the smaller selection in Rusted.
6. Run one concrete Codex routing smoke and preserve its raw evidence.

## Risks

- Pack aliases can create dependency cycles or duplicate artifacts; closure and
  materialization tests must prove the result.
- Overly narrow descriptions can hide valid tasks; each description must keep
  representative positive signals while explicitly excluding adjacent domains.
- A lexical test can pass while a runtime ignores a skill. Its name and docs
  must make that limitation visible, and the Rusted smoke remains separate.
- Rusted has preexisting uncommitted benchmark and harness artifacts. Sync must
  preserve them and remove only unchanged manifest-owned obsolete files.

## Verification

- `go test ./internal/cli -run 'Game|CatalogComposite' -count=1`
- `go test ./... -count=1`
- `go run ./cmd/geremmyas lint`
- `go run ./cmd/geremmyas doctor`
- `git diff --check`
- Rusted project sync with manifest inspection and `geremmyas context`
- One read-only, ephemeral Codex routing smoke with raw JSONL evidence

## Completion

Mark `Verified` only after structural checks pass, Rusted contains the focused
skill set, obsolete managed game skills reconcile safely, and the real runtime
smoke result is recorded honestly.
