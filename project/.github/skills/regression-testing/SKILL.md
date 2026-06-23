---
name: regression-testing
description: "Mandatory regression test creation for every bug fix. Ensures fixed bugs don't reappear. Use after bug identification, before applying the fix."
---

# Regression Testing

## When to Use

- After identifying bug root cause (never skip this)
- Before applying the fix
- After fix verification to confirm test catches the bug
- Before considering bug truly closed

## When NOT to Use

- During initial bug investigation (investigate first)
- For speculative issues (only after reproducing)
- For features (feature tests ≠ regression tests)
- For refactoring with no behavioral change

## Regression Test Purpose

**A regression test** is a test that would FAIL before the fix and PASS after.

It proves:
1. Bug was real (test catches the broken state)
2. Fix actually works (test passes after fix)
3. Bug won't sneak back (test is part of suite)

**Language agnostic**: Applies to Go, Python, JavaScript/TypeScript, Java, Rust, etc. Concepts are universal; syntax varies.

## Workflow

**1. Reproduce the Bug**
```
App state: A
User action: B
Expected: C
Actual: D (broken)
```

**2. Write Test for Expected Behavior (RED)**
```go
func TestLoginWithExpiredSessionReturnsUnauthorized(t *testing.T) {
    session := createExpiredSession()
    resp := httpClient.Post("/login", session)

    assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
    // This test FAILS before fix, PASSES after
}
```

**3. Verify Test Fails**
```
$ go test -run TestLoginWithExpiredSessionReturnsUnauthorized
--- FAIL: TestLoginWithExpiredSessionReturnsUnauthorized
Expected: 401
Actual: 500
```

**4. Apply Fix**
Fix the bug in production code.

**5. Verify Test Passes**
```
$ go test -run TestLoginWithExpiredSessionReturnsUnauthorized
--- PASS: TestLoginWithExpiredSessionReturnsUnauthorized
```

**6. Run Full Suite**
```
$ go test ./...
ok  example.com/pkg  0.324s
```

**7. Verify No Regressions**
- All existing tests still pass
- No new failures introduced
- Edge cases still handled

---

## Regression Test Template

**Location**: Same test file as feature being tested

**Naming Pattern** (language-independent):
- `[Feature][ConditionThatCausedBug][ExpectedBehavior]`
- Example: `CartCheckoutWithHighInventoryDoesNotDoubleCharge`
- Goal: Name should explain what was broken and what we're checking

### Go Pattern

```go
// Bad test (insufficient):
func TestLogin(t *testing.T) {
    resp := login()
    assert.NotNil(t, resp)
}

// Good regression test (specific, reproducible):
func TestLoginWithExpiredSessionReturnsUnauthorizedNotInternalError(t *testing.T) {
    // Setup: state that triggers the bug
    session := createSessionWithExpiry(time.Now().Add(-1 * time.Hour))

    // Action: trigger the bug path
    resp := httpClient.Post("/api/login", loginRequest{
        sessionID: session.ID,
    })

    // Verify: expected behavior (not the bug)
    assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
    assert.NotContains(t, resp.Body, "Internal Server Error") // was the bug
    assert.Empty(t, resp.Cookie("auth_token"))
}
```

### JavaScript/TypeScript Pattern

```typescript
// Bad test (insufficient):
test("login works", () => {
  const resp = login();
  expect(resp).toBeDefined();
});

// Good regression test (specific, reproducible):
test("LoginWithExpiredSessionReturnsUnauthorizedNotInternalError", () => {
  // Setup: state that triggers the bug
  const session = createSessionWithExpiry(Date.now() - 3600000);

  // Action: trigger the bug path
  const resp = httpClient.post("/api/login", {
    sessionID: session.id,
  });

  // Verify: expected behavior (not the bug)
  expect(resp.status).toBe(401);
  expect(resp.body).not.toContain("Internal Server Error"); // was the bug
  expect(resp.cookies["auth_token"]).toBeUndefined();
});
```

### Python Pattern

```python
# Bad test (insufficient):
def test_login():
    resp = login()
    assert resp is not None

# Good regression test (specific, reproducible):
def test_login_with_expired_session_returns_unauthorized_not_internal_error():
    # Setup: state that triggers the bug
    session = create_session_with_expiry(datetime.now() - timedelta(hours=1))

    # Action: trigger the bug path
    resp = http_client.post(
        "/api/login",
        json={"session_id": session.id}
    )

    # Verify: expected behavior (not the bug)
    assert resp.status_code == 401, f"Expected 401, got {resp.status_code}"
    assert "Internal Server Error" not in resp.text, "Bug still present: got 500"
    assert "auth_token" not in resp.cookies
```

---

## Regression Categories

### Category 1: Boundary Condition

**Bug**: Off-by-one error in pagination

**Go**:
```go
func TestPaginationWithExactPageBoundary(t *testing.T) {
    items := createItems(25)  // Exactly 1 page
    page := fetchPage(1)

    assert.Len(t, page.Items, 25)
    assert.Equal(t, page.Total, 25)
    assert.False(t, page.HasNext)
}
```

**Python**:
```python
def test_pagination_with_exact_page_boundary():
    items = create_items(25)  # Exactly 1 page
    page = fetch_page(1)

    assert len(page["items"]) == 25
    assert page["total"] == 25
    assert page["has_next"] is False
```

**JavaScript/TypeScript**:
```typescript
test("PaginationWithExactPageBoundary", () => {
  const items = createItems(25);  // Exactly 1 page
  const page = fetchPage(1);

  expect(page.items).toHaveLength(25);
  expect(page.total).toBe(25);
  expect(page.hasNext).toBe(false);
});
```

### Category 2: State Transition

**Bug**: State machine skips validation on edge case

**Go**:
```go
func TestOrderTransitionFromPendingToRefundedSkipsInProgressValidation(t *testing.T) {
    order := createOrder(statusPending)

    err := order.TransitionTo(statusRefunded)  // Should fail

    assert.NotNil(t, err)
    assert.Equal(t, order.Status, statusPending)  // State unchanged
}
```

**Python**:
```python
def test_order_transition_from_pending_to_refunded_skips_in_progress_validation():
    order = create_order(Status.PENDING)

    with pytest.raises(ValueError):
        order.transition_to(Status.REFUNDED)  # Should fail

    assert order.status == Status.PENDING  # State unchanged
```

**JavaScript/TypeScript**:
```typescript
test("OrderTransitionFromPendingToRefundedSkipsInProgressValidation", () => {
  const order = createOrder(Status.PENDING);

  expect(() => order.transitionTo(Status.REFUNDED)).toThrow();
  expect(order.status).toBe(Status.PENDING);  // State unchanged
});
```

### Category 3: Concurrency

**Bug**: Race condition in shared state

**Go**:
```go
func TestConcurrentUpdateToCounterDoesNotLoseUpdates(t *testing.T) {
    counter := NewAtomicCounter(0)

    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            counter.Increment()
            wg.Done()
        }()
    }
    wg.Wait()

    assert.Equal(t, counter.Value(), 100)  // Would be <100 before fix
}
```

**Python**:
```python
def test_concurrent_update_to_counter_does_not_lose_updates():
    counter = AtomicCounter(0)
    threads = []

    for _ in range(100):
        t = threading.Thread(target=counter.increment)
        threads.append(t)
        t.start()

    for t in threads:
        t.join()

    assert counter.value == 100  # Would be <100 before fix
```

**JavaScript/TypeScript** (using Promise.all):
```typescript
test("ConcurrentUpdateToCounterDoesNotLoseUpdates", async () => {
  const counter = new AtomicCounter(0);

  await Promise.all(
    Array.from({ length: 100 }, () => counter.increment())
  );

  expect(counter.value).toBe(100);  // Would be <100 before fix
});
```

### Category 4: NULL/Empty Handling

**Bug**: Nil pointer dereference on empty input

**Go**:
```go
func TestProcessPaymentWithEmptyInvoiceDoesNotPanic(t *testing.T) {
    var invoice *Invoice  // nil

    result, err := processPayment(invoice)

    assert.Nil(t, result)
    assert.NotNil(t, err)
    assert.Equal(t, err.Error(), "invoice required")
}
```

**Python**:
```python
def test_process_payment_with_empty_invoice_does_not_panic():
    invoice = None

    with pytest.raises(ValueError, match="invoice required"):
        process_payment(invoice)
```

**JavaScript/TypeScript**:
```typescript
test("ProcessPaymentWithEmptyInvoiceDoesNotPanic", () => {
  const invoice = null;

  expect(() => processPayment(invoice)).toThrow("invoice required");
});
```

### Category 5: External Service Failure

**Bug**: Unhandled timeout on external API call

**Go**:
```go
func TestCheckoutWithPaymentServiceTimeoutReturnsGracefulError(t *testing.T) {
    mockPaymentService.SetTimeout(100 * time.Millisecond)

    err := checkout(testOrder)

    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "payment service")  // Clear error, not generic
    assert.NotContains(t, err.Error(), "connection refused")  // Not infrastructure detail
}
```

**Python**:
```python
def test_checkout_with_payment_service_timeout_returns_graceful_error():
    mock_payment_service.set_timeout(0.1)  # 100ms

    with pytest.raises(PaymentError) as exc_info:
        checkout(test_order)

    assert "payment service" in str(exc_info.value)
    assert "connection refused" not in str(exc_info.value)
```

**JavaScript/TypeScript**:
```typescript
test("CheckoutWithPaymentServiceTimeoutReturnsGracefulError", async () => {
  mockPaymentService.setTimeout(100);  // 100ms

  await expect(checkout(testOrder)).rejects.toThrow(/payment service/);
  await expect(checkout(testOrder)).rejects.not.toThrow(/connection refused/);
});
```

---

## Red Flags — Weak Regression Test

| Flag | Why | Fix |
|------|-----|-----|
| Test passes even with bug in place | Not testing the bug | Use pre-fix run to verify test fails |
| "Test" is just a code review | No assertion, just reads code | Add concrete assertion |
| Flaky on repeated runs | Environmental dependency | Mock external deps, use fixed seeds |
| Test name doesn't match bug | Future you won't know what it prevents | Rename to "BugXYZ_[condition]_[behavior]" |
| Only tests happy path | Bug was in edge case | Test the edge case |

---

## Verification After Fix

**Before committing the fix:**

### Universal Steps (all languages)

1. **Show test FAILS with bug present** (do NOT fix yet)
   - Keep the test file (you just created it)
   - Run the test against broken code
   - Verify it FAILS with the expected error

2. **Apply the fix** in production code

3. **Show test PASSES after fix**
   - Run the same test again
   - Verify it now PASSES

4. **Run full suite**
   - Ensure no regressions introduced
   - All existing tests still pass

5. **Commit with reference**

### Language-Specific Commands

**Go**:
```bash
# Before fix
go test -run TestLoginWithExpiredSession  # FAILS

# Apply fix, then
go test -run TestLoginWithExpiredSession  # PASSES
go test ./...  # All pass
```

**Python**:
```bash
# Before fix
pytest tests/test_auth.py::test_login_with_expired_session -v  # FAILS

# Apply fix, then
pytest tests/test_auth.py::test_login_with_expired_session -v  # PASSES
pytest tests/  # Full suite passes
```

**JavaScript/TypeScript**:
```bash
# Before fix
npm test -- LoginWithExpiredSession  # FAILS

# Apply fix, then
npm test -- LoginWithExpiredSession  # PASSES
npm test  # Full suite passes
```

### Commit Pattern

```bash
git commit -m "fix(auth): handle expired sessions correctly

Before: login() returned 500 for expired session
After: login() returns 401, clears auth token

Regression test: LoginWithExpiredSessionReturnsUnauthorized
Fixes: #BUG-123"
```

---

**Key Principle**: Regression tests are not optional. They are insurance against the same bug appearing twice. Every bug fix gets a test that catches it. No exceptions.
