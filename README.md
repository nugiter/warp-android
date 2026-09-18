# WarpScout Chain — Wails v3 Android/Windows migration

This package is rebuilt from the original `warp-ready.zip` source and migrates the Wails v2 application to **Wails v3 beta.23**, with GitHub Actions workflows for Windows x64 and Android arm64.

## Build targets

- Windows x64: `WarpScoutChain.exe`
- Android arm64: `WarpScoutChain.apk`

Wails v3 currently requires Go 1.25+ and provides Android builds through its Taskfile-based mobile toolchain. Android uses the NDK to compile the Go application as `libwails.so` and Gradle to package the APK.

## GitHub Actions

After uploading this project to a GitHub repository:

1. Open **Actions**.
2. Select **Build Android APK**.
3. Click **Run workflow**.
4. After the job succeeds, download artifact **WarpScoutChain-Android-arm64**.

The Windows workflow is analogous and produces **WarpScoutChain-Windows-x64**.

The workflows intentionally use:

- Go 1.25
- Node.js 24
- Java 17
- Android API 35
- Android NDK 26.3.11579264
- Wails v3.0.0-beta.23
- current GitHub Actions Node 24-compatible major versions

The workflows obtain the official Wails v3 build/task assets from the matching Wails tag at build time. This keeps the project archive small while ensuring the Android Gradle/NDK build tree matches the selected Wails release.

## Important Android notes

The original application used Wails v2 runtime APIs. The backend has been migrated to the Wails v3 service/event model:

- `runtime.EventsEmit` → `app.Event.Emit`
- v2 `Bind` → `application.Service`
- `wails.Run` → `application.New` + `app.Run`
- v2 frontend `window.go.main.App.*` → generated Wails v3 service bindings

The frontend now uses `@wailsio/runtime` events and generated bindings. The ZIP result is returned to JavaScript as Base64 so the Android WebView does not depend on the old localhost download route.

The application still keeps its localhost subscription endpoint for desktop compatibility. A separate Android client cannot automatically consume an app-local `127.0.0.1` subscription URL; use the exported configuration files when moving the generated configuration to another Android app.

## Local toolchain check

Install Go 1.25+, then:

```bash
go install github.com/wailsapp/wails/v3/cmd/wails3@v3.0.0-beta.23
wails3 doctor
```

For Android, install Android SDK, NDK 26.3.11579264 and a JDK. Then run:

```bash
wails3 task android:package ARCH=arm64
```

## Status

This archive is a **migration/build package**, not a precompiled APK. The final APK is produced by the GitHub Actions Android runner, where the Android SDK/NDK and Wails build assets are available.
