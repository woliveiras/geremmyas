---
description: "TypeScript code quality based on the official Handbook, Google TS Style Guide, and strict compiler conventions."
applyTo: "**/*.ts,**/*.tsx"
---

# TypeScript Conventions

Derive types from values and invariants; do not maintain parallel type
definitions. Prefer inference, composition, and narrowing over casts.

## Compiler Strictness

- Enable `strict` in all projects; it is the baseline, not aspirational
- Enable `noUncheckedIndexedAccess` for safe array/map lookups
- Enable `exactOptionalPropertyTypes` to distinguish `undefined` from missing
- Enable `useUnknownInCatchVariables` (catch variables are `unknown`)
- Enable `noImplicitOverride` for class hierarchies
- Enable `noPropertyAccessFromIndexSignature` to force bracket notation on
  dynamic keys
- Enable `isolatedModules` and `verbatimModuleSyntax` for transpiler
  compatibility
- Keep `skipLibCheck: false` as target; use `true` only during migration
- Run `tsc --noEmit` in CI even when bundler does the transpilation

## Type Modeling

- Use `interface` for object shapes and contracts that benefit from `extends`
- Use `type` for unions, intersections, conditionals, mapped types, and helpers
- Prefer discriminated unions over loose optional fields for mutually exclusive
  states
- Use `as const` to freeze literal types; `satisfies` to check conformance
  without losing inference
- Derive types with utility types (`Omit`, `Pick`, `Partial`, `Readonly`,
  `ReturnType`, `Awaited`)
- Avoid overloads when a union parameter solves the same problem with less
  friction
- Keep generics minimal; they should be inferred from arguments, not forced on
  callers

## Narrowing & Safety

- Use type guards (`is`) and assertion functions (`asserts`) for runtime checks
- Prefer `unknown` at trust boundaries; narrow with guards before using
- Use `assertNever` in switch default for exhaustive union handling
- Never use `as` casts without prior narrowing or proof of correctness
- Never use `any` without explicit justification; prefer `unknown`
- Non-null assertion (`!`) is a code smell; handle nullability properly

## Modules & Exports

- Use ESM with named exports; avoid default exports
- Use `import type` for type-only imports (`verbatimModuleSyntax` enforces this)
- Keep exported API surface minimal; internal modules stay unexported
- Do not export mutable bindings
- Organize by feature, not by technical layer, in applications

## Naming

- `PascalCase` for types, interfaces, classes, enums, and components
- `camelCase` for functions, variables, properties, and methods
- `UPPER_SNAKE_CASE` only for truly global immutable constants
- No `I` prefix on interfaces; no `T` prefix on type parameters unless
  conventional in the codebase
- File names: `kebab-case.ts` or `camelCase.ts` matching project convention

## Error Handling

- Only throw `Error` instances (or subclasses); never throw primitives
- Use `Error.cause` to chain context without losing the original error
- Catch `unknown`; narrow before accessing properties
- Never swallow errors silently; log or rethrow with context
- Lint: enable `no-floating-promises`, `only-throw-error`,
  `prefer-promise-reject-errors`

## Async

- Await all promises or explicitly mark fire-and-forget with `void`
- Never use `forEach` with async callbacks (it drops the promise)
- Use `Promise.all` / `Promise.allSettled` for concurrent operations
- Set explicit return types on public async functions
- Handle rejections at every call site or propagate explicitly

## Runtime Validation

- Types do not exist at runtime; validate external data at boundaries
- Use Zod, Valibot, or io-ts for schema validation of HTTP, env, queues, files
- Infer TypeScript types from validation schemas (`z.infer<typeof Schema>`)
- Never trust `as` casts on data from external sources

## Functions & Design

- Prefer small, focused functions with explicit return types on public APIs
- Use `readonly` arrays and properties when mutation is not intended
- Prefer immutable patterns; mutate only when performance requires it
- Avoid classes as namespaces; use module-level functions and types

## Testing

- Use Vitest for modern stacks; Jest for legacy compatibility
- Testing Library for UI behavior tests; Playwright for E2E
- Test observable behavior, not implementation details
- Type-check test files with `tsc`; do not exclude them from strict mode

## Linting

- ESLint + typescript-eslint with `recommendedTypeChecked` config
- Prettier (or Biome) for formatting; disable conflicting ESLint style rules
- Key rules: `no-explicit-any`, `no-unsafe-assignment`,
  `no-floating-promises`, `consistent-type-imports`, `only-throw-error`
- Run lint and type-check as separate CI steps

## Anti-patterns

- `any` spread across the codebase (use `unknown` + narrowing)
- Blind `as Foo` casts without proof
- Parallel type definitions maintained manually (derive instead)
- Overloads where a union suffices
- `skipLibCheck: true` as permanent setting
- Default exports (harder to refactor, rename, and trace)
- Mutable exported state
- `catch (e) { }` without handling or rethrowing
- Ignoring `tsc --noEmit` in CI because "the bundler compiles it"
