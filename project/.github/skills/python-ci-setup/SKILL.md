---
name: python-ci-setup
description: "Set up a Python CI/CD pipeline with linting, formatting, type checking, tests, security scanning, and publishing. Use when: creating CI for Python, adding checks to workflow. Do not use: for local development setup, non-CI automation."
---


# Python CI Setup

Multi-step workflow to create a production-grade Python CI pipeline.

## When to Use

- New Python project needs CI from scratch
- Existing project has incomplete or fragile CI
- Adding security scanning, multi-version testing, or Trusted Publishing

## Pipeline Layers (in order)

Each layer gates the next. If a layer fails, the pipeline stops.

1. **Format** — `ruff format --check .` (or `black --check .`) must pass
2. **Lint** — `ruff check .` (or `flake8`) must pass
3. **Type check** — `mypy src` (or Pyright) must pass
4. **Test** — `pytest` with coverage threshold
5. **Security scan** — `bandit -r src` + `pip-audit`
6. **Build** — `python -m build` produces valid sdist/wheel
7. **Publish** (optional) — Trusted Publishing to PyPI via OIDC

## Steps

### 1. Choose CI Platform

Default to GitHub Actions for GitHub-hosted repositories.
Ensure matrix testing across supported Python versions (3.11, 3.12, 3.13, 3.14).

### 2. Create Quality Workflow

For GitHub Actions: `.github/workflows/ci.yml`

```yaml
name: CI
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  quality:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        python-version: ["3.11", "3.12", "3.13", "3.14"]
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-python@v5
        with:
          python-version: ${{ matrix.python-version }}

      - run: python -m pip install -U pip
      - run: python -m pip install -e ".[dev]"

      - run: ruff format --check .
      - run: ruff check .
      - run: mypy src
      - run: pytest --tb=short
      - run: coverage run -m pytest && coverage report --fail-under=85
      - run: bandit -r src
      - run: pip-audit
```

### 3. Configure Tools in pyproject.toml

All tool configuration lives in `pyproject.toml`:

```toml
[tool.ruff]
line-length = 100
target-version = "py311"

[tool.ruff.lint]
select = ["E", "F", "I", "B", "UP"]

[tool.pytest.ini_options]
testpaths = ["tests"]
addopts = "-ra"

[tool.coverage.run]
branch = true
source = ["src/my_package"]

[tool.coverage.report]
fail_under = 85

[tool.mypy]
python_version = "3.11"
strict = true
mypy_path = ["src"]

[tool.bandit]
exclude_dirs = ["tests"]
```

### 4. Add Pre-commit Hooks

`.pre-commit-config.yaml`:

```yaml
repos:
  - repo: https://github.com/astral-sh/ruff-pre-commit
    rev: v0.14.2
    hooks:
      - id: ruff-check
      - id: ruff-format

  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v5.0.0
    hooks:
      - id: end-of-file-fixer
      - id: trailing-whitespace
      - id: check-yaml
```

### 5. Add Publish Workflow (optional)

For libraries published to PyPI, use Trusted Publishing:

```yaml
name: Publish
on:
  release:
    types: [published]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-python@v5
        with:
          python-version: "3.14"
      - run: python -m pip install -U build
      - run: python -m build
      - uses: actions/upload-artifact@v4
        with:
          name: dist
          path: dist/*

  publish:
    needs: build
    runs-on: ubuntu-latest
    environment: pypi
    permissions:
      id-token: write
    steps:
      - uses: actions/download-artifact@v4
        with:
          name: dist
          path: dist
      - uses: pypa/gh-action-pypi-publish@release/v1
```

Key decisions:
- Trusted Publishing uses OIDC tokens (15-min lifetime), no stored secrets
- Separate `build` and `publish` jobs for auditability
- `environment: pypi` enables manual approval gates

### 6. Add Multi-version Testing (optional)

Use `nox` or `tox` for local matrix testing:

```python
# noxfile.py
import nox

@nox.session(python=["3.11", "3.12", "3.13", "3.14"])
def tests(session):
    session.install("-e", ".[dev]")
    session.run("pytest")
```

### 7. Verify Pipeline

- Push a branch with intentional failures to confirm gates work
- Confirm coverage threshold is enforced
- Confirm `bandit` and `pip-audit` output is visible in CI logs
- Test publish workflow against TestPyPI before production

## Checklist

- [ ] Format check (ruff format or black) enforced
- [ ] Lint (ruff check) enabled with curated rule set
- [ ] Type check (mypy or Pyright) in strict mode
- [ ] Tests run with coverage threshold
- [ ] `bandit` scans for common security issues
- [ ] `pip-audit` checks for vulnerable dependencies
- [ ] Python version matrix covers supported range
- [ ] All tools configured in `pyproject.toml`
- [ ] Pre-commit hooks installed for local feedback
- [ ] Publish uses Trusted Publishing, not manual tokens
