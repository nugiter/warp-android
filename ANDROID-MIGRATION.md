# Android migration notes

The source application has been migrated from Wails v2 to Wails v3 service/event APIs.

The Android workflow uses the official Wails v3 beta.23 mobile build tree at build time. This avoids committing generated Gradle/JNI assets to the application repository and keeps them synchronized with the Wails CLI version used by CI.

The current APK target is arm64 (`arm64-v8a`), which is the normal physical-device target. The Wails Android task can also build x86_64 for emulators and a fat APK when required.
