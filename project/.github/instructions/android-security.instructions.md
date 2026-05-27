---
description: "Use when writing or reviewing Android Kotlin code, mobile storage, permissions, networking, or authentication flows."
applyTo: "**/*.kt, **/AndroidManifest.xml"
---

# Android Security Conventions

Security in Android is a design discipline. Apply least privilege, secure
communication, defensive storage, and reduced operational surface across
components, logs, and dependencies.

## Input Validation

- Validate all external input as early as possible (intents, deep links,
  content providers, files from external storage)
- Use allowlists for structured formats (regex, enum matching)
- Validate at syntactic and semantic levels before data reaches persistence
  or network
- Use `require` / `check` for preconditions; `takeIf` for safe narrowing
- Treat data from external storage, clipboard, and other apps as untrusted

## Secure Storage

- Store cryptographic keys in **Android Keystore** (non-exportable, with usage
  policies)
- Use **DataStore** for preferences and configuration (async, transactional,
  coroutine-based); never SharedPreferences in new code
- `EncryptedSharedPreferences` is deprecated; avoid in new projects
- For encrypted preferences, evaluate **DataStore + Tink** (`datastore-tink`)
  with maturity assessment (currently alpha)
- Room for structured local data; encrypt sensitive databases with SQLCipher
  when required
- Never store secrets, tokens, credentials, or PII in plain files, resources,
  or source code

## Network Security

- HTTPS/TLS always; never disable certificate validation
- Use **Network Security Configuration** for declarative rules per domain
- Set `cleartextTrafficPermitted="false"` in base-config
- Certificate pinning only with backup pin and documented rotation process
- Debug overrides only in `<debug-overrides>` block
- Update security provider via `ProviderInstaller` when needed
- Set timeouts on all HTTP clients

## Permissions

- Request minimum permissions; ask only when the user triggers the action
- Degrade gracefully on permission denial
- On Android 13+, revoke unused permissions with
  `revokeSelfPermissionOnKill()`
- Document every dangerous permission with user-facing rationale
- Never request permissions at app startup without context

## Components & Intents

- Components are **not exported** by default (`android:exported="false"`)
- Use implicit intents with chooser for cross-app sharing
- Use `FileProvider` / `content://` URIs; never `file://`
- Grant URI read/write permissions with flags, scoped and temporary
- Validate all intent extras and deep link parameters before use

## R8 / ProGuard

- Enable R8 in release builds (`isMinifyEnabled = true`,
  `isShrinkResources = true`)
- Maintain keep rules for reflection, serialization, and Retrofit models
- Use `proguard-android-optimize.txt` as base
- Keep mapping file for crash deobfuscation (R8 Retrace)
- Retrofit includes R8 rules automatically; verify custom converters

## App Integrity

- Use **Play Integrity API** on the backend to verify genuine binary, Play
  install, and trusted environment
- Treat integrity checks as part of anti-abuse strategy, not sole protection
- Use **Play App Signing** with separate upload key
- Never bundle signing keys in source or CI artifacts

## Logging & Observability

- Never log tokens, credentials, PII, or full server responses
- Use Timber with a release tree that strips verbose/debug logs
- Crashlytics for production crashes; validate with test crash before release
- Add breadcrumbs and custom keys (without PII) for debugging context
- Source maps / mapping files must be uploaded for deobfuscation

## Dependency & Supply Chain Security

- Review dependencies via **Play SDK Index** for known vulnerabilities
- Use **GitHub Dependency Review** to block vulnerable deps in PRs
- Run **CodeQL** or **OWASP Dependency-Check** in CI
- Enable Gradle **dependency verification** for third-party plugins
- Audit new SDKs for permissions, data collection, and maintenance status
- Pin dependency versions in lockfile; review lockfile diffs

## Checklist

- [ ] All external input validated (intents, deep links, files, clipboard)
- [ ] `cleartextTrafficPermitted` is **false** by default
- [ ] Keys stored in **Android Keystore**
- [ ] Preferences use **DataStore**, not SharedPreferences
- [ ] No `EncryptedSharedPreferences` in new code
- [ ] R8 enabled with `shrinkResources` in release
- [ ] Components not exported unless explicitly required
- [ ] Permissions requested only on user action, with graceful fallback
- [ ] No PII in logs; Timber release tree strips debug output
- [ ] Dependency review in CI (SDK Index / Dependency Review / CodeQL)
- [ ] Play App Signing with separate upload key
- [ ] Mapping file preserved for production crash deobfuscation
