# WarpScout Chain — Android migration package

This directory is the Android migration scaffold for the existing Wails v2 desktop application.

## Important

The source project is Wails v2.9.2. Wails v2's stable CLI supports desktop targets; native Android/iOS support is part of the newer Wails v3 mobile work and is currently experimental. Therefore this package deliberately does **not** claim that the existing v2 source can be turned into a working APK by merely adding an Android workflow.

The GitHub Actions workflow installs the Android SDK/NDK/JDK and Wails v3 CLI and validates the migration scaffold. The next implementation step is migrating the Wails v2 bridge (`wails.Run`, `runtime.EventsEmit`, generated `window.go.main.App.*`) to the Wails v3 Android bridge or to a dedicated JNI/gomobile bridge.

The existing Windows build remains unchanged.
