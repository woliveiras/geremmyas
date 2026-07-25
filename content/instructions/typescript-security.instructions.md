---
description: "Use when writing or reviewing TypeScript code at trust boundaries, handling external data, dependencies, or security-sensitive operations."
applyTo: "**/*.ts,**/*.tsx"
---

# TypeScript Security

## Types Are Not Runtime Protection

- TypeScript types are erased at runtime; they prevent zero attacks by themselves
- Vite, SWC, Babel, esbuild, and Node type-stripping do NOT type-check
- `tsc --noEmit` must run in CI as an independent verification step
- Every external boundary (HTTP, env, files, queues, WebSocket, localStorage)
  needs runtime validation

## Runtime Validation at Boundaries

- Validate all external input with schema libraries (Zod, Valibot, io-ts)
- Infer types from schemas, not the other way around
- Never trust `as` casts on data from network, storage, or third parties
- Validate environment variables at startup; fail fast on missing values
- Validate URL parameters, query strings, and form data before use

## Strict Compiler as First Layer

- `strict: true` catches null derefs, implicit any, and unsafe operations
- `noUncheckedIndexedAccess` prevents unsafe array/object access
- `exactOptionalPropertyTypes` distinguishes missing from explicit undefined
- `useUnknownInCatchVariables` forces narrowing in catch blocks
- These flags prevent entire classes of bugs at zero runtime cost

## Dependency & Supply Chain Security

- Run `npm audit` (or equivalent) in CI; fail on high/critical
- Enable Dependabot or Renovate for automated dependency updates
- Use lockfile-based installs in CI: `npm ci`, `pnpm install --frozen-lockfile`,
  `yarn install --immutable`
- Audit new dependencies before adding; check maintainership, size, and scope
- For published packages: use trusted publishing with OIDC provenance
- Pin exact versions in lockfile; review lockfile diffs in PRs

## Error & Exception Safety

- Only throw `Error` instances; never throw strings or plain objects
- Catch `unknown`; narrow with `instanceof Error` before accessing properties
- Use `Error.cause` for chaining; never discard context
- Never `catch (e) { }` without handling, logging, or rethrowing
- Floating promises are security risks (unhandled rejections, race conditions)

## Secrets & Sensitive Data

- Never hardcode secrets, tokens, or keys in source
- Load from environment or secret manager at runtime
- Do not log secrets, tokens, PII, or full request bodies
- Use content exclusion for paths containing sensitive configuration
- `.env` files must be in `.gitignore`; never commit them

## XSS & Injection (Browser Context)

- Use framework escaping (React JSX, Angular templates) for user content
- Never use `innerHTML`, `dangerouslySetInnerHTML`, or `eval` with user input
- Sanitize HTML only with established libraries (DOMPurify)
- Validate and encode URLs before rendering in `href` or `src`
- Use CSP headers to limit script sources

## Node.js Specific

- Set timeouts on all HTTP clients, database connections, and external calls
- Validate and limit request body size
- Use parameterized queries for databases; never string-interpolate SQL
- Avoid `child_process.exec` with user input; prefer `execFile` with args array
- Do not use `Function()` constructor or `vm.runInNewContext` with untrusted code

## Checklist

- [ ] `tsc --noEmit` runs in CI independently of bundler
- [ ] External data validated with schema library at boundaries
- [ ] No `as` casts on untrusted data
- [ ] `npm audit` (or equivalent) in CI pipeline
- [ ] Lockfile-based install in CI (no floating resolutions)
- [ ] Dependabot or Renovate configured
- [ ] Secrets loaded from environment, never committed
- [ ] Error handling: only `Error` instances, always with cause
- [ ] No floating promises (lint rule enforced)
- [ ] Source maps published for production error tracking
