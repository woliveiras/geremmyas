---
description: "Use when writing or reviewing unsafe Rust code, FFI boundaries, raw pointers, or repr annotations."
applyTo: "**/*.rs"
---

# Rust Unsafe & FFI

## Unsafe Blocks

- Isolate `unsafe` to the smallest possible scope
- Wrap every `unsafe` block in a safe public API
- Add a `// SAFETY:` comment immediately above explaining why the usage is sound
- Never implement `Send` or `Sync` manually unless you can prove the invariants

## Safety Comments

Format:
```rust
// SAFETY: <invariant that makes this sound in this specific context>
unsafe { ... }
```

The comment must explain the specific local reasoning, not just restate the docs.

## FFI Boundaries

- Types crossing FFI must have stable ABI: use `#[repr(C)]`, `#[repr(transparent)]`, or primitive types
- `repr(Rust)` has no layout guarantees; never pass it across FFI
- Use `CString`/`CStr` for NUL-terminated strings; validate before conversion
- Ownership transfer across languages must be explicitly documented (who allocates, who frees)
- Wrap raw C APIs in a safe Rust interface; keep the `unsafe extern` block private

## Raw Pointers

- Validate non-null before dereferencing
- Document lifetime expectations (how long the pointer remains valid)
- Prefer references and slices over raw pointers whenever possible
- Use `NonNull<T>` to encode non-null invariant in the type

## Repr & Layout

- `#[repr(C)]` for FFI structs and stable layout
- `#[repr(transparent)]` for newtypes that must be ABI-compatible with inner type
- `#[repr(u8)]` / `#[repr(i32)]` etc. for enums crossing FFI
- Never assume field order or padding without explicit repr

## Validation Tools

- Miri: run `cargo +nightly miri test` to detect UB in tests (nightly only)
- Sanitizers: AddressSanitizer, ThreadSanitizer for memory and concurrency errors
- Fuzzing: `cargo-fuzz` or `libfuzzer` for parsers, codecs, serializers, and FFI wrappers
- These tools complement each other; no single tool catches all classes of bugs

## Concurrency Safety

- `Send`: type can be moved to another thread
- `Sync`: type can be shared (via `&T`) across threads
- `Rc<T>` and `RefCell<T>` are single-thread only; use `Arc<T>` and `Mutex<T>` for multi-thread
- Never implement `Send`/`Sync` without auditing all interior mutability and pointer usage
- Data races are UB; race conditions (logic bugs) are not prevented by the compiler
