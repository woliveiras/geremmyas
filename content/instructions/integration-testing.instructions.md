---
description: "Use when writing or reviewing integration tests. Focuses on real boundaries between modules and controlled external dependencies."
applyTo: "**/integration/**, **/*.integration.*, **/tests/integration/**"
---

# Integration Testing Conventions

- Test real collaboration between modules at meaningful boundaries.
- Use real infrastructure adapters when practical, with controlled test
  databases, queues, filesystems, or HTTP servers.
- Keep external third-party services behind fakes, contract fixtures, or local
  test doubles.
- Seed data explicitly in the test; do not depend on developer machine state.
- Verify durable side effects through the public boundary under test.
- Clean up resources created by the test.
- Keep integration tests slower and fewer than unit tests, but strong enough to
  catch wiring, serialization, transaction, and configuration bugs.
