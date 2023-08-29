# Outline Connectivity Test App

## Overview

This is a simple cross-platform app to test connectivity to Outline servers, using the Outline SDK. It is built with [Wails](https://wails.app/) and [Capacitor](https://capacitorjs.com/).

### Architecture

The overarching goal of this application is to demonstrate how the Outline SDK enables you to write each line of business logic only once across all platforms.

We achieve this by first writing a [`shared_backend`](./shared_backend) package in Go - which contains the UDP/TCP connectivity test with the Outline SDK - and a [`shared_frontend`](./shared_frontend/) GUI in TypeScript and Lit that contains an HTML for filling out all the required parameters.

Each platform used - [Wails](https://wails.app/) for desktop and [Capacitor](https://capacitorjs.com/) for mobile - then has a thin wrapper around the shared code that handles the platform-specific details.

```mermaid
graph LR
  subgraph Shared
    A["shared_backend"]
    B["shared_frontend"]
  end
  subgraph Build
    C["SharedBackend.xcframework"]
    D["SharedBackend.aar"]
  end
  subgraph app_desktop
    H["index.html"]
    G["DesktopBackend.Invoke()"]
  end
  subgraph app_mobile
    subgraph ios
      I["MobileBackendPlugin.swift"]
    end
    subgraph android
      J["MobileBackendPlugin.kt"]
    end
    K["index.html"]
    L["MobileBackend.Invoke()"]
  end

  A -.-> |gomobile| C
  A -.-> |gomobile| D
  A --> G
  B --> H
  B --> K
  C --> I
  D --> J
  I --> L
  J --> L
  L --> K
  G --> H

  style K fill:blue;
  style H fill:blue;
```

For Mobile, we use `gomobile` to build the `shared_backend` package into a `xcframework` for iOS and an `aar` for Android. You can see this for yourself by running `yarn shared_backend build`. For Desktop, Wails simply refers to the `shared_backend` package directly.

Then we implement a small piece of middleware called `Invokable` that enables the frontend to speak to the backend via the given platform.

```ts
interface Invokable {
  Invoke(parameters: { method: string; input: string; }): Promise<{ result: string; errors: string[] }>
}
```

[TO BE CONTINUED...]

## Development

### Prerequisites

- [Node.js](https://nodejs.org/)
- [Yarn](https://yarnpkg.com/)
- [Go](https://golang.org/)
- [Wails](https://wails.app/)
- [Capacitor](https://capacitorjs.com/)
- [CocoaPods](https://cocoapods.org/)
- [Xcode](https://developer.apple.com/xcode/)
- [Android SDK](https://developer.android.com/studio)
- [Android NDK](https://developer.android.com/ndk)

### Setup

1. Clone this repo
1. `cd` into the repo
1. `yarn`

If at any point you run into issues during development, try `yarn reset`.

### Development Server

`yarn watch`

### Build

> NOTE: You will need credentials to build for iOS and Android. Talk to @jigsaw.

`yarn build`

### Needed Improvements

1. **\[P1\]** android (in progress)
1. **\[P1\]** read browser language on load (and only localize once)

### Current Issues

1. <span style="color:red">**\[P0\]** add server url to an ENV var somehow... pretty dumb capacitor...</span>
1. **\[P1\]** Results dialog isn't rendering as intended (likely because of the `{ all: initial }`)
1. **\[P2\]** `cap ___ run` breaks (have workaround and [issue filed](https://github.com/ionic-team/capacitor/issues/6791))
1. <span style="color:gray">**\[P3\]** spurious lit localize TS error</span>
