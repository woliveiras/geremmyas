# Android CI Configuration Examples

Load only the section needed for the current pipeline change. Replace example
versions and package names with values verified for the target project.

## Version Catalog

```toml
[versions]
agp = "9.0.1"
kotlin = "2.3.4"
detekt = "1.23.8"

[plugins]
android-application = { id = "com.android.application", version.ref = "agp" }
ksp = { id = "com.google.devtools.ksp", version.ref = "kotlin" }
detekt = { id = "dev.detekt", version.ref = "detekt" }
```

## Dependabot

```yaml
version: 2
updates:
  - package-ecosystem: gradle
    directory: /
    schedule:
      interval: weekly
    open-pull-requests-limit: 10
```

## Fastlane

```ruby
platform :android do
  lane :internal do
    gradle(task: "bundle", build_type: "Release")
    upload_to_play_store(
      track: "internal",
      aab: "app/build/outputs/bundle/release/app-release.aab"
    )
  end

  lane :promote do
    upload_to_play_store(track: "internal", track_promote_to: "production")
  end
end
```

Use Play App Signing with a separate upload key. Never store signing material
in the repository.

## CodeQL

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
