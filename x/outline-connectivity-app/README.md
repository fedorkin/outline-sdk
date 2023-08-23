# Outline Connectivity Test App

## Overview

This is a simple cross-platform app to test connectivity to Outline servers, using the Outline SDK. It is built with [Wails](https://wails.app/) and [Capacitor](https://capacitorjs.com/).

### Architecture

TODO

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

1. **\[P1\]** Results dialog isn't rendering as intended (likely because of the `{ all: initial }`)
1. **\[P2\]** `cap ___ run` breaks (have workaround and [issue filed](https://github.com/ionic-team/capacitor/issues/6791))
1. **\[P2\]** android live reload (simulator doesn't see localhost)
1. <span style="color:gray">**\[P3\]** spurious lit localize TS error</span>
