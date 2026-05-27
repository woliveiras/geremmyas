---
name: android-ci-setup
description: "Set up an Android CI/CD pipeline with lint, static analysis, unit tests, dependency review, build verification, and Play Store deployment."
---

# Android CI Setup

Multi-step workflow to create a production-grade Android CI pipeline with
GitHub Actions, Gradle, and Fastlane.

## When to Use

- New Android project needs CI from scratch
- Existing project has incomplete or fragile CI
- Adding static analysis, dependency review, or automated deployment

## Pipeline Layers (in order)

Each layer gates the next. If a layer fails, the pipeline stops.

1. **Setup** — JDK, Gradle cache, dependency resolution
2. **Lint** — Android Lint (`lintDebug`)
3. **Static Analysis** — detekt + ktlint
4. **Unit Tests** — `testDebugUnitTest`
5. **Build** — `assembleRelease` or `bundleRelease`
6. **Dependency Review** — Block PRs introducing vulnerable deps
7. **Deploy** (release branches) — Fastlane to Play Console tracks

## Steps

### 1. Choose Build Configuration

| Tool | Recommendation |
|------|----------------|
| Gradle DSL | **Kotlin DSL** (`build.gradle.kts`) |
| Dependencies | **Version Catalogs** (`gradle/libs.versions.toml`) |
| Annotation Processing | **KSP** (not kapt, when supported) |
| JDK | 17 (minimum for AGP 8+/9) |

### 2. Create Workflow File

`.github/workflows/android-ci.yml`:

```yaml
name: Android CI

on:
  pull_request:
  push:
    branches: [main]

jobs:
  quality:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-java@v4
        with:
          distribution: temurin
          java-version: "17"
          cache: gradle

      - name: Lint
        run: ./gradlew lintDebug

      - name: Static analysis
        run: ./gradlew detekt ktlintCheck

      - name: Unit tests
        run: ./gradlew testDebugUnitTest

      - name: Build release
        run: ./gradlew assembleRelease

  dependency-review:
    if: github.event_name == 'pull_request'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/dependency-review-action@v4
```

### 3. Configure Android Lint

In `app/build.gradle.kts`:

```kotlin
android {
    lint {
        warningsAsErrors = true
        abortOnError = true
        checkDependencies = true
        baseline = file("lint-baseline.xml")
    }
}
```

Generate baseline for existing issues: `./gradlew lintDebug` (first run
creates the baseline file).

### 4. Configure detekt

`config/detekt/detekt.yml` with project rules. In root `build.gradle.kts`:

```kotlin
plugins {
    alias(libs.plugins.detekt)
}

detekt {
    buildUponDefaultConfig = true
    allRules = false
    config.setFrom(files("$rootDir/config/detekt/detekt.yml"))
    parallel = true
}

dependencies {
    detektPlugins("io.gitlab.arturbosch.detekt:detekt-formatting:${libs.versions.detekt.get()}")
}
```

### 5. Configure ktlint

`.editorconfig`:

```ini
root = true

[*.{kt,kts}]
charset = utf-8
end_of_line = lf
insert_final_newline = true
indent_style = space
indent_size = 4
ij_kotlin_allow_trailing_comma = true
ktlint_code_style = ktlint_official
```

Use the Gradle plugin (`org.jlleitschuh.gradle.ktlint`) or run via detekt
formatting ruleset.

### 6. Configure R8 for Release

```kotlin
android {
    buildTypes {
        release {
            isMinifyEnabled = true
            isShrinkResources = true
            proguardFiles(
                getDefaultProguardFile("proguard-android-optimize.txt"),
                "proguard-rules.pro",
            )
        }
    }
}
```

Keep rules for reflection (Retrofit models, Room entities, serialization):

```pro
-keep class com.yourpackage.api.model.** { *; }
-keepattributes *Annotation*
```

### 7. Add Version Catalog

`gradle/libs.versions.toml`:

```toml
[versions]
agp = "9.0.1"
kotlin = "2.3.4"
hilt = "2.57.1"
detekt = "1.23.8"
room = "2.7.1"

[plugins]
android-application = { id = "com.android.application", version.ref = "agp" }
ksp = { id = "com.google.devtools.ksp", version.ref = "kotlin" }
hilt = { id = "com.google.dagger.hilt.android", version.ref = "hilt" }
detekt = { id = "dev.detekt", version.ref = "detekt" }

[libraries]
hilt-android = { module = "com.google.dagger:hilt-android", version.ref = "hilt" }
hilt-compiler = { module = "com.google.dagger:hilt-compiler", version.ref = "hilt" }
room-runtime = { module = "androidx.room:room-runtime", version.ref = "room" }
room-compiler = { module = "androidx.room:room-compiler", version.ref = "room" }
room-ktx = { module = "androidx.room:room-ktx", version.ref = "room" }
```

### 8. Add Dependabot

`.github/dependabot.yml`:

```yaml
version: 2
updates:
  - package-ecosystem: gradle
    directory: /
    schedule:
      interval: weekly
    open-pull-requests-limit: 10
```

### 9. Add Fastlane for Deployment (optional)

`fastlane/Fastfile`:

```ruby
platform :android do
  desc "Deploy to internal track"
  lane :internal do
    gradle(task: "bundle", build_type: "Release")
    upload_to_play_store(
      track: "internal",
      aab: "app/build/outputs/bundle/release/app-release.aab",
      skip_upload_images: true,
      skip_upload_screenshots: true,
      skip_upload_metadata: true
    )
  end

  desc "Promote internal to production"
  lane :promote do
    upload_to_play_store(
      track: "internal",
      track_promote_to: "production",
      skip_upload_changelogs: false
    )
  end
end
```

Use Play App Signing with a separate upload key. Never store signing keys
in the repository.

### 10. Add CodeQL (optional)

`.github/workflows/codeql.yml`:

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
      - uses: actions/setup-java@v4
        with:
          distribution: temurin
          java-version: "17"
      - uses: github/codeql-action/init@v3
        with:
          languages: java-kotlin
      - uses: github/codeql-action/autobuild@v3
      - uses: github/codeql-action/analyze@v3
```

### 11. Verify Pipeline

- Push a branch with a lint violation to confirm `lintDebug` blocks merge
- Push a branch with a detekt finding to confirm static analysis blocks
- Verify unit test failure blocks the pipeline
- Confirm dependency-review catches a vulnerable dependency in a PR
- Verify release build succeeds with R8 enabled
- Test Fastlane deployment to internal track (manual trigger first)

## Checklist

- [ ] Gradle cache configured in CI (`cache: gradle`)
- [ ] `./gradlew lintDebug` runs with `warningsAsErrors = true`
- [ ] detekt + ktlint run as separate analysis step
- [ ] Unit tests run (`testDebugUnitTest`)
- [ ] Release build with R8 enabled succeeds in CI
- [ ] Dependency Review blocks PRs with vulnerable deps
- [ ] JDK version from setup (not hardcoded in Gradle wrapper)
- [ ] Version catalogs centralize all dependency versions
- [ ] KSP used instead of kapt where supported
- [ ] Dependabot configured for Gradle ecosystem
- [ ] Mapping file preserved for crash deobfuscation
- [ ] Play App Signing with separate upload key (never in repo)
- [ ] CodeQL or Dependency-Check for supply chain scanning
