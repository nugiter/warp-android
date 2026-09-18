# Android build status

This repository currently contains a Wails v2 desktop application plus an Android migration scaffold.

The Android workflow intentionally validates the Android/Java/Wails v3 toolchain and uploads the migration notes; it does **not** claim to produce a final APK yet. The existing application imports `github.com/wailsapp/wails/v2` and uses the Wails v2 desktop runtime API, so a real Android APK requires migration to the Wails v3 mobile API first.

The GitHub Actions warnings have been addressed by using Node 24-compatible action versions:
- `actions/checkout@v5`
- `actions/setup-go@v6`
- `actions/setup-java@v6`
- `android-actions/setup-android@v4`

Go module caching is disabled in the Android workflow because the repository currently has no `go.sum`; this removes the `go.sum` cache warning. Once dependencies are resolved and a committed `go.sum` is added, caching can be enabled again.
