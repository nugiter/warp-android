# Android migration notes

## Existing desktop-only bridge

The current application uses:

- `github.com/wailsapp/wails/v2`
- `wails.Run(...)` in `main.go`
- `github.com/wailsapp/wails/v2/pkg/runtime` in `app.go`
- generated Wails JavaScript bindings: `window.go.main.App.GenerateConfigs(...)`
- a loopback HTTP server on `127.0.0.1:8888`

These pieces are not Android-ready as-is.

## Required migration

1. Move business logic out of the Wails v2 `App` bridge into a platform-neutral Go package.
2. Replace Wails v2 runtime event emission with the Wails v3 mobile event API (or a dedicated callback bridge).
3. Replace the generated v2 `window.go.main.App.*` calls in the frontend with the v3 runtime bindings.
4. Replace desktop file-dialog assumptions with Android Storage Access Framework / Android download handling.
5. Bind the local subscription service to an Android-safe interface; do not assume that a desktop loopback URL is reachable from every Android WebView configuration.
6. Build and test at least `arm64-v8a` on Android 13/14+.

## Security

Do not ship signing keys in the repository. For release builds, use GitHub Actions secrets and sign the APK/AAB in a separate release workflow.
