---
description: "Rust code quality based on Effective Rust, API Guidelines, and the Rust Book."
applyTo: "**/*.rs"
---

# Rust Code Quality

Follow the [Rust API Guidelines](https://rust-lang.github.io/api-guidelines/),
[Effective Rust](https://www.lurklurk.org/effective-rust/), and the
[Rust Book](https://doc.rust-lang.org/book/).

## Ownership & Borrowing

- Accept borrowed inputs for read-only APIs: `&str` over `String`, `&[T]` over `Vec<T>`, `&Path` over `PathBuf`
- Return owned types only when there is a real transformation or the callee must own the data
- Explicit lifetime annotations only when output borrows from input; do not add unnecessary annotations
- Prefer `Clone` over `Rc`/`Arc` unless profiling shows cloning is expensive

## Error Handling

- `Option<T>` for absence without failure detail
- `Result<T, E>` for recoverable failures
- `panic!` only for broken invariants, prototypes, examples, and tests
- Library errors: typed enum with `thiserror`
- Application boundaries: `anyhow::Result<T>` with `.context()`
- No `unwrap()` or `expect()` in production paths unless the invariant is proven and commented
- Wrap errors with context: propagate with `?` and add `.context("what failed")`

## Naming (API Guidelines)

- Conversion methods: `as_` (cheap borrow), `to_` (expensive or lossy), `into_` (ownership transfer)
- Iterator methods: `iter()`, `iter_mut()`, `into_iter()`
- Constructors: `new()` for the default, `with_*()` for variants, builder pattern for complex construction
- Error types: suffix `Error` (`ParseError`, `ConfigError`)
- Newtypes for semantic distinctions (`UserId(u64)` not bare `u64`)

## Generics & Traits

- Prefer generics + trait bounds when all types are known at compile time
- Use `dyn Trait` for heterogeneous collections, plugins, or open extension points
- Keep trait methods object-safe unless there is a strong reason not to
- Verify trait compliance: `const _: () = { fn assert_send<T: Send>() {} assert_send::<MyType>(); };`

## Iterators & Closures

- Prefer `map`/`filter`/`collect` pipelines over imperative loops when readability improves
- Iterators are zero-cost abstractions; do not avoid them for performance reasons without profiling
- Use `Iterator` trait for lazy evaluation; collect only when needed

## Module Organization

- Packages group crates; crates are module trees; workspaces coordinate multiple packages
- Domain crate does not depend on infrastructure
- Binaries compose library crates; keep `main.rs` thin
- Use `pub(crate)` to limit visibility; default to private

## Testing

- Unit tests in `#[cfg(test)] mod tests` inside the module
- Integration tests in `tests/` directory, testing public API only
- Table-driven tests with clear arrange/act/assert structure
- `#[should_panic]` for invariant violation tests
- Property testing with `proptest` for parsers and input handling
- Snapshot testing with `insta` for structured output

## Linting

- `#![warn(clippy::all, clippy::pedantic)]` in `lib.rs`/`main.rs`
- Allow specific lints only with local justification comment
- `cargo fmt --all -- --check` in CI
- `cargo clippy --workspace --all-targets --all-features -- -D warnings` in CI
