---
name: typescript-ci-setup
description: "Set up a TypeScript CI/CD pipeline with type checking, linting, formatting, tests, security scanning, and build verification."
---

# TypeScript CI Setup

Multi-step workflow to create a production-grade TypeScript CI pipeline.

## When to Use

- New TypeScript project needs CI from scratch
- Existing project has incomplete or fragile CI
- Adding type-check gates, security scanning, or lockfile enforcement

## Pipeline Layers (in order)

Each layer gates the next. If a layer fails, the pipeline stops.

1. **Install** — Lockfile-based (`npm ci` / `pnpm install --frozen-lockfile`)
2. **Type check** — `tsc --noEmit` (independent of bundler)
3. **Lint** — ESLint with typescript-eslint type-checked rules
4. **Format check** — Prettier or Biome (`--check` mode)
5. **Test** — Vitest/Jest with coverage threshold
6. **Build** — Production build (Vite, tsup, esbuild, webpack)
7. **Audit** — `npm audit --audit-level=high` or equivalent
8. **Code scanning** (optional) — CodeQL for JavaScript/TypeScript

## Steps

### 1. Choose Package Manager & Lockfile Strategy

| Manager | CI Install Command |
|---------|-------------------|
| npm     | `npm ci` |
| pnpm    | `pnpm install --frozen-lockfile` |
| yarn    | `yarn install --immutable` |
| bun     | `bun install --frozen-lockfile` |

The lockfile must be committed. CI must fail if the lockfile is out of sync.

### 2. Create Workflow File

For GitHub Actions: `.github/workflows/ci.yml`

```yaml
name: CI
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  ci:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version-file: .node-version
          cache: npm

      - run: npm ci
      - run: npx tsc --noEmit
      - run: npx eslint .
      - run: npx prettier . --check
      - run: npx vitest run --coverage
      - run: npm run build
      - run: npm audit --audit-level=high
```

### 3. Configure tsconfig for Strictness

Minimum recommended flags:

```jsonc
{
  "compilerOptions": {
    "strict": true,
    "noUncheckedIndexedAccess": true,
    "exactOptionalPropertyTypes": true,
    "useUnknownInCatchVariables": true,
    "noImplicitOverride": true,
    "noPropertyAccessFromIndexSignature": true,
    "isolatedModules": true,
    "verbatimModuleSyntax": true,
    "skipLibCheck": false
  }
}
```

### 4. Configure ESLint

Use flat config with typescript-eslint type-checked rules:

```js
// eslint.config.mjs
import eslint from "@eslint/js";
import tseslint from "typescript-eslint";
import eslintConfigPrettier from "eslint-config-prettier";

export default tseslint.config(
  eslint.configs.recommended,
  ...tseslint.configs.recommendedTypeChecked,
  {
    languageOptions: {
      parserOptions: { projectService: true },
    },
    rules: {
      "@typescript-eslint/no-explicit-any": "error",
      "@typescript-eslint/no-unsafe-assignment": "error",
      "@typescript-eslint/no-floating-promises": "error",
      "@typescript-eslint/consistent-type-imports": "error",
      "@typescript-eslint/only-throw-error": "error",
    },
  },
  eslintConfigPrettier,
);
```

### 5. Add Prettier Config

```json
{
  "semi": true,
  "singleQuote": false,
  "trailingComma": "all",
  "printWidth": 100
}
```

Or use Biome as unified linter+formatter alternative.

### 6. Configure Test Runner

Vitest (recommended for Vite stacks):

```ts
// vitest.config.ts
import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    coverage: {
      provider: "v8",
      thresholds: { lines: 80, branches: 75 },
    },
  },
});
```

### 7. Add Dependabot

`.github/dependabot.yml`:

```yaml
version: 2
updates:
  - package-ecosystem: npm
    directory: /
    schedule:
      interval: weekly
    open-pull-requests-limit: 10
```

### 8. Add Code Scanning (optional)

For CodeQL: `.github/workflows/codeql.yml`

```yaml
name: CodeQL
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
  schedule:
    - cron: "0 6 * * 1"

jobs:
  analyze:
    runs-on: ubuntu-latest
    permissions:
      security-events: write
    steps:
      - uses: actions/checkout@v4
      - uses: github/codeql-action/init@v3
        with:
          languages: javascript-typescript
      - uses: github/codeql-action/analyze@v3
```

### 9. Verify Pipeline

- Push a branch with a type error to confirm `tsc --noEmit` blocks merge
- Push a branch with a lint violation to confirm ESLint blocks
- Verify coverage threshold fails when dropping below minimum
- Confirm `npm audit` output appears in CI logs
- Verify lockfile enforcement catches `package.json` / lockfile drift

## Checklist

- [ ] Lockfile-based install in CI (no floating resolutions)
- [ ] `tsc --noEmit` runs independently of bundler
- [ ] ESLint with typescript-eslint type-checked rules
- [ ] Prettier (or Biome) format check
- [ ] Tests run with coverage threshold enforced
- [ ] Production build step succeeds
- [ ] `npm audit` (or equivalent) fails on high/critical vulns
- [ ] Node version from `.node-version` or `package.json` (not hardcoded)
- [ ] Dependabot or Renovate configured for dependency updates
- [ ] CodeQL or equivalent code scanning (for public/sensitive repos)
- [ ] Source maps generated for production error tracking
