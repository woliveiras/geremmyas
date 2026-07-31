# Phaser build and release

Inspect `package.json`, lockfile, Vite or other bundler config, TypeScript config, public asset layout, environment variables, hosting base path, and supported browsers.

## Build

- Use the repository’s pinned package manager and production script.
- Expose only explicitly public environment variables.
- Verify dynamic imports, hashed chunks, asset URLs, source-map policy, and bundle warnings.
- Keep runtime asset keys stable even when filenames are hashed.
- Account for deployment under a subpath.

## Verify

Serve the built output over HTTP rather than opening files directly. Inspect network failures, console errors, WebGL or Canvas selection, audio unlock, font loading, resize, DPR, fullscreen, focus loss, and storage.

## Delivery

Configure compression and cache immutable hashed assets aggressively while ensuring the HTML entry point can discover new versions. Review service-worker update behavior and offline caches when used.

## Official references

- [Phaser game configuration](https://docs.phaser.io/phaser/concepts/game)
- [Loader](https://docs.phaser.io/phaser/concepts/loader)
- [Scale Manager](https://docs.phaser.io/phaser/concepts/scale-manager)
- [Vite production build](https://vite.dev/guide/build.html)
