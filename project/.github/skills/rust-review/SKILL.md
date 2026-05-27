---
name: rust-review
description: "Systematic code review checklist for Rust. Use when: reviewing Rust code, reviewing a Rust PR, checking Rust code quality, or validating Copilot-generated Rust suggestions."
---

# Rust Code Review

Systematic review checklist for Rust code, designed to catch the most common
issues in both human-written and AI-generated code.

## When to Use

- Reviewing a pull request with Rust changes
- Validating Copilot-generated Rust code before accepting
- Self-reviewing before submitting a PR
- Auditing unsafe code or FFI boundaries

## Process

1. Read the spec or task description to understand intent.
2. Run the checklist below against the diff.
3. For each finding, note the file, location, problem, and suggested fix.
4. Verify that tests exist and cover the changed behavior.
5. Run `cargo clippy` and `cargo test` locally if possible.

## Checklist

### Ownership & Allocation

- [ ] No unnecessary `.clone()` — data should be borrowed when only read
- [ ] Input parameters use `&str`, `&[T]`, `&Path` for read-only access
- [ ] Owned types returned only when transformation occurs
- [ ] No `Box<dyn Trait>` where a generic would suffice
- [ ] No `Arc`/`Rc` without shared ownership need

### Error Handling

- [ ] No `unwrap()` or `expect()` in production paths
- [ ] Errors propagated with `?` and enriched with `.context()`
- [ ] Library errors use typed enums (not `String` or `Box<dyn Error>`)
- [ ] No hidden panics (indexing without bounds check, unchecked arithmetic)
- [ ] `Option` used for absence, `Result` for failures

### Concurrency

- [ ] No lock held across `.await` (use `tokio::sync::Mutex` if unavoidable)
- [ ] Lock scope is minimal (guard dropped before next operation)
- [ ] Spawned tasks have cancellation/shutdown mechanism
- [ ] No blocking I/O or CPU work on async runtime without `spawn_blocking`
- [ ] Channels are bounded (backpressure considered)

### Unsafe & FFI

- [ ] Every `unsafe` block has a `// SAFETY:` comment with local reasoning
- [ ] `unsafe` scope is minimal (no extra safe code inside the block)
- [ ] FFI types use `#[repr(C)]` or stable layout
- [ ] Raw pointers validated before dereference
- [ ] Ownership across FFI is documented (who allocates, who frees)

### API Design

- [ ] Public API follows naming conventions (`as_`, `to_`, `into_`, `iter`)
- [ ] No leaking internal types in public signatures
- [ ] Generic bounds are minimal (only what the implementation uses)
- [ ] Builder pattern for types with many optional fields
- [ ] Newtypes for semantic distinctions (not bare primitives)

### Testing

- [ ] New behavior has corresponding test(s)
- [ ] Tests verify observable behavior, not implementation details
- [ ] Edge cases covered (empty input, boundary values, error paths)
- [ ] No test-only public visibility (`pub` added just for testing)
- [ ] Async tests use `#[tokio::test]`

### Performance & Observability

- [ ] No allocation in hot loops without justification
- [ ] `tracing` spans and structured fields for async operations
- [ ] No `println!`/`eprintln!` in production paths
- [ ] Iterators preferred over manual index manipulation

### Style

- [ ] Code passes `cargo fmt --check`
- [ ] Code passes `cargo clippy` without warnings
- [ ] `#[allow(...)]` has a justification comment
- [ ] Modules are not oversized; split at ~300-400 lines

## Output Format

For each finding:

```
<file>:<line> — <category>
Problem: <what is wrong>
Fix: <specific suggestion>
```

Summarize with: total findings, severity breakdown, and whether the change is
safe to merge after fixes.
