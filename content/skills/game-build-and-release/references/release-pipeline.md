# Release pipeline

## Reproducible inputs

Pin or record:

- engine and export-template version;
- runtime, package manager, and lockfile;
- dependencies and addons;
- build flags and public environment;
- asset and content versions;
- target platform and architecture;
- version and build number source.

Keep credentials outside tracked configuration and logs.

## Pipeline stages

1. Validate repository and dependency state.
2. Run formatting, static checks, and tests.
3. Import or preprocess assets deterministically.
4. Build the release artifact.
5. Inspect manifest, contents, licenses, sizes, and secret exposure.
6. Generate hashes and provenance.
7. Serve, install, or unpack exactly as users will receive it.
8. Run artifact-level smoke tests.
9. Publish only with explicit authority.

## Artifact smoke route

Verify:

- application starts without console or engine errors;
- title or first scene loads;
- input works;
- one gameplay interaction completes;
- scene transition or level load works;
- audio unlocks or starts under platform policy;
- save and reload works when in scope;
- resize, fullscreen, focus loss, and pause behave;
- missing or misbased assets are absent;
- application can exit or return to menu cleanly.

## Web delivery

Check HTTP status, MIME types, compression, cache headers, service worker behavior, base URL, relative assets, CORS or cross-origin isolation when required, mobile viewport, and browser console.

## Release evidence

Record:

- source revision or dirty-state note;
- command and environment;
- tool versions;
- artifact path, size, and checksum;
- automated results;
- runtime smoke evidence;
- unsupported or unverified targets;
- manual store, device, accessibility, and performance checks.

## Rollback

Keep previous known-good artifacts and document compatibility of saves, servers, and cached web clients. Never assume a binary rollback can safely read data written by a newer version.
