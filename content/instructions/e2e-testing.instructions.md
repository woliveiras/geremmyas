---
description: "Use when writing or reviewing end-to-end tests. Focuses on user journeys, stable selectors, real app behavior, and reliable verification."
applyTo: "**/e2e/**, **/*.e2e.*, **/playwright.config.*, **/cypress.config.*"
---

# E2E Testing Conventions

- Test user-visible journeys, not implementation details.
- Prefer accessible selectors (`getByRole`, labels, visible text) over CSS
  selectors or test-only IDs.
- Keep each test independent; do not rely on ordering between tests.
- Control external systems with seeded data, fixtures, or test accounts.
- Wait for user-visible state, not arbitrary timeouts.
- Capture useful debugging evidence on failure: console errors, network errors,
  screenshots, traces, or videos when the framework supports them.
- Keep E2E coverage focused on critical flows; push detailed edge cases to
  lower-level tests when possible.
