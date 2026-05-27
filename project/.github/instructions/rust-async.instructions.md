---
description: "Use when writing or reviewing async Rust code with Tokio, concurrency primitives, channels, or shared state."
applyTo: "**/*.rs"
---

# Rust Async & Concurrency

## Runtime

- Tokio is the default async runtime unless the project already uses another
- Every binary with async code needs `#[tokio::main]` or explicit runtime build
- Futures are lazy; they only progress when polled (awaited)

## Task Design

- Use `tokio::spawn` for independent concurrent work
- Every spawned task must have a cancellation mechanism (`CancellationToken`, `select!`, or channel close)
- Never block the async runtime: use `tokio::task::spawn_blocking` for CPU-bound or blocking I/O
- Long computation without `.await` points starves other tasks; break it up or offload

## Concurrency Primitives (preference order)

1. Message passing (`tokio::sync::mpsc`, `oneshot`, `broadcast`) when ownership can move between tasks
2. `Arc<Mutex<T>>` or `Arc<RwLock<T>>` when shared state is unavoidable
3. Atomics only when you can prove correctness and the operation is simple

## Locks

- Minimize lock scope; never hold a lock across an `.await` point
- Prefer `tokio::sync::Mutex` over `std::sync::Mutex` when the lock must be held across awaits
- Use `std::sync::Mutex` for synchronous-only critical sections (cheaper)

## Send & Sync

- Futures passed to `tokio::spawn` must be `Send`
- Do not hold non-Send types (like `Rc`, `RefCell`) across `.await` boundaries
- If you need single-thread context, use `tokio::task::LocalSet`

## Channels

- `mpsc` for many-to-one; bounded channels preferred (backpressure)
- `oneshot` for single response patterns (request/reply)
- `broadcast` for fan-out to multiple consumers
- `watch` for latest-value sharing (config reload, state)

## Error Handling in Async

- `JoinHandle` returns `Result<T, JoinError>`; always handle the outer error
- Use `tokio::select!` with cancellation safety awareness
- Timeouts: `tokio::time::timeout` wrapping futures, not manual timer logic

## Structured Concurrency

- Prefer `JoinSet` or `FuturesUnordered` over unbounded `tokio::spawn` scattering
- Group related tasks and await all completions before proceeding
- Graceful shutdown: signal all tasks, then await with deadline

## Observability

- Use `tracing` with `#[instrument]` for async-aware structured spans
- `tokio-console` for runtime debugging (task states, poll times, waker counts)
- Never use `println!` for production diagnostics; use structured fields
