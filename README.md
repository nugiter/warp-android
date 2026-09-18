# WarpScout Chain

Windows build: Wails v2.9.2.

Android migration scaffold: see `android/`.

## Windows

The existing `.github/workflows/build-windows.yml` builds the Windows x64 executable.

## Android

Run `.github/workflows/build-android.yml` manually from GitHub Actions. It prepares the Android/JDK/NDK toolchain and uploads the Android migration scaffold as an artifact.

**This workflow intentionally does not produce a fake APK.** The current application is Wails v2.9.2 desktop code. Native Android packaging requires migrating the Wails bridge to the newer mobile-capable architecture first.
