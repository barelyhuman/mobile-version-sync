# mobile-version-sync

simple too to change versionName on Android and CFBundleShortVersionString on iOS
additionally allows to increment versionCode and CFBundleVersion

# UNDER ACTIVE DEVELOPMENT

#### Files to edit

**Android**
app/src/main/AndroidManifest.xml
/app/build.gradle

**iOS**

- [x] <app_name>/Info.plist

## What Works

As of now the iOS platform works, android is still being worked on.

## Usage

```sh
Usage of mobile-version-sync
  -app string
        app name
  -b    alias to bump
  -bump
        bump the build/versionCode as well (default: false)
  -d string
        alias to dir (default "./ios/")
  -dir string
        directory that contains the android or iOS structure (default "./ios/")
  -p string
        alias to platform (default "ios")
  -platform string
        platform to run the version sync on [ios|android] (default "ios")
  -v string
        alias to version
  -version string
        version string to sync with
```
