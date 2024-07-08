# Overview

__N.b.__ Latest version should always be on top

| Heading        | Description                            |
| :------------- | :------------------------------------- |
| __Added__      | for new features.                      |
| __Changed__    | for changes in existing functionality. |
| __Deprecated__ | for soon-to-be removed features.       |
| __Removed__    | for now removed features.              |
| __Fixed__      | for any bug fixes.                     |
| __Security__   | in case of vulnerabilities.            |

__Reference__:
* [keepachangelog](https://keepachangelog.com/en/1.0.0/)
* [Semantic Versioning](https://semver.org/)

# CHANGELOG

## [0.6.0]

### Changed

* Improved dev dependency management with tools.go and a Makefile

## [0.5.0]

### Added

* golangci config and updated pre-commit config to utilise linting

## [0.4.0]

### Chnaged

* cmd package to be called by the main.go in the root directory rather than
  directly.

## [0.3.0]

### Changed

* Simplified the conversion logic

## [0.2.0]

### Changed

* Moved validation logic into its own module in internal/
* Abstracted code in prompt and validation scripts

## [0.1.0]

### Added

* Initial release
