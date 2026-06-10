# kamu

Unified CLI for the Kamu platform — drive **kamudb** (databases), **kamubee** (apps), **kamudns** (DNS), and **kamustatus** (uptime monitoring) from one binary with one login against **kamuid**.

```
kamu auth login
kamu db list
kamu bee apps
kamu dns zones
kamu status projects list
```

`kamu status` talks to [kamustatus](https://github.com/kontakto-fi/kamustatus). Until [kamustatus#5](https://github.com/kontakto-fi/kamustatus/issues/5) lands it needs a project-scoped key:

```sh
export KAMU_KAMUSTATUS_API_KEY=km_...
```

## Install

### Homebrew

```sh
brew install kotisivukamu/tap/kamu
```

### From source

Requires Go 1.25+.

```sh
go install github.com/kotisivukamu/kamu-cli/cmd/kamu@latest
```

### Pre-built binaries

Download from [Releases](https://github.com/kotisivukamu/kamu-cli/releases). Archives for `darwin_amd64`, `darwin_arm64`, `linux_amd64`, `linux_arm64`.

## Status

Early — see [#1](https://github.com/kotisivukamu/kamu-cli/issues/1) for the milestone plan. M0 (scaffold) and the release pipeline are in; auth and the per-service subcommands are stubs until the platform-side conventions land ([kamuid#1](https://github.com/kotisivukamu/kamuid/issues/1), [kamuid#2](https://github.com/kotisivukamu/kamuid/issues/2)).

## Development

```sh
go build -o kamu ./cmd/kamu
./kamu --help
./kamu version
```

Layout follows [flyctl](https://github.com/superfly/flyctl): one package per noun under `internal/command/`, one file per verb.

## Release

Push a `vX.Y.Z` tag; GitHub Actions runs GoReleaser, publishes the GitHub release, and pushes the Homebrew formula to [kotisivukamu/homebrew-tap](https://github.com/kotisivukamu/homebrew-tap).
