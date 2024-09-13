# Go Repository Template

[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/humweb/app-boilerplate)](https://github.com/humweb/app-boilerplate/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/humweb/app-boilerplate.svg)](https://pkg.go.dev/github.com/humweb/app-boilerplate)
[![go.mod](https://img.shields.io/github/go-mod/go-version/humweb/app-boilerplate)](go.mod)
[![LICENSE](https://img.shields.io/github/license/humweb/app-boilerplate)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/humweb/app-boilerplate/build.yml?branch=main)](https://github.com/humweb/app-boilerplate/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/humweb/app-boilerplate)](https://goreportcard.com/report/github.com/humweb/app-boilerplate)
[![codecov](https://codecov.io/gh/humweb/app-boilerplate/graph/badge.svg?token=TJZL21LFKC)](https://codecov.io/gh/humweb/app-boilerplate)

⭐ `Star` this repository if you find it valuable and worth maintaining.

👁 `Watch` this repository to get notified about new releases, issues, etc.

## Description

This is a GitHub repository template for a Go application.
You can use it:

- to create a new repository with automation and environment setup,
- as reference when improving automation for an existing repository.

It includes:

- continuous integration via [GitHub Actions](https://github.com/features/actions),
- build automation via [Make](https://www.gnu.org/software/make),
- dependency management using [Go Modules](https://github.com/golang/go/wiki/Modules),
- code formatting using [gofumpt](https://github.com/mvdan/gofumpt),
- linting with [golangci-lint](https://github.com/golangci/golangci-lint)
  and [misspell](https://github.com/client9/misspell),
- unit testing with
  [race detector](https://blog.golang.org/race-detector),
  code coverage [HTML report](https://blog.golang.org/cover)
  and [Codecov report](https://codecov.io/),
- releasing using [GoReleaser](https://github.com/goreleaser/goreleaser),
- dependencies scanning and updating thanks to [Dependabot](https://dependabot.com),
- security code analysis using [CodeQL Action](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning),
  and [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck),
- [Visual Studio Code](https://code.visualstudio.com) configuration with [Go](https://code.visualstudio.com/docs/languages/go) support.

## Usage

1. Sign up on [Codecov](https://codecov.io/) and configure
   [Codecov GitHub Application](https://github.com/apps/codecov).
1. Click the `Use this template` button (alt. clone or download this repository).
1. Replace all occurrences of `humweb/app-boilerplate` to `your_org/repo_name` in all files.
1. Replace all occurrences of `app-boilerplate` to `repo_name` in [Dockerfile](Dockerfile).
1. Follow [these](https://docs.codecov.com/docs/adding-the-codecov-token#github-actions)
   instructions to add the `CODECOV_TOKEN` GitHub Actions and Dependabot secret.
1. Update the following files:
   - [CHANGELOG.md](CHANGELOG.md)
   - [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)
   - [LICENSE](LICENSE)
   - [README.md](README.md)

## Setup

Below you can find sample instructions on how to set up the development environment.
Of course, you can use other tools like [GoLand](https://www.jetbrains.com/go/),
[Vim](https://github.com/fatih/vim-go), [Emacs](https://github.com/dominikh/go-mode.el).
However, take notice that the Visual Studio Go extension is
[officially supported](https://blog.golang.org/vscode-go) by the Go team.

1. Install [Go](https://golang.org/doc/install).
1. Install [Visual Studio Code](https://code.visualstudio.com/).
1. Install [Go extension](https://code.visualstudio.com/docs/languages/go).
1. Clone and open this repository.
1. `F1` -> `Go: Install/Update Tools` -> (select all) -> OK.

## Build

### Terminal

- `make` - execute the build pipeline.
- `make help` - print help for the [Make targets](Makefile).

### Visual Studio Code

`F1` → `Tasks: Run Build Task (Ctrl+Shift+B or ⇧⌘B)` to execute the build pipeline.

## Release

The release workflow is triggered each time a tag with `v` prefix is pushed.

_CAUTION_: Make sure to understand the consequences before you bump the major version.
More info: [Go Wiki](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher),
[Go Blog](https://blog.golang.org/v2-go-modules).

## Maintenance

Notable files:

- [.github/workflows](.github/workflows) - GitHub Actions workflows,
- [.github/dependabot.yml](.github/dependabot.yml) - Dependabot configuration,
- [.vscode](.vscode) - Visual Studio Code configuration files,
- [.golangci.yml](.golangci.yml) - golangci-lint configuration,
- [.goreleaser.yml](.goreleaser.yml) - GoReleaser configuration,
- [Dockerfile](Dockerfile) - Dockerfile used by GoReleaser to create a container image,
- [Makefile](Makefile) - Make targets used for development, [CI build](.github/workflows) and [.vscode/tasks.json](.vscode/tasks.json),
- [go.mod](go.mod) - [Go module definition](https://github.com/golang/go/wiki/Modules#gomod),
- [tools.go](tools.go) - [build tools](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module).

## FAQ

### Why Visual Studio Code editor configuration

Developers that use Visual Studio Code can take advantage of the editor configuration.
While others do not have to care about it.
Setting configs for each repo is unnecessary time consuming.
VS Code is the most popular Go editor ([survey](https://blog.golang.org/survey2019-results))
and it is officially [supported by the Go team](https://blog.golang.org/vscode-go).

You can always remove the [.vscode](.vscode) directory if it really does not help you.

### Why GitHub Actions, not any other CI server

GitHub Actions is out-of-the-box if you are already using GitHub.
[Here](https://github.com/mvdan/github-actions-golang) you can learn how to use it for Go.

However, changing to any other CI server should be very simple,
because this repository has build logic and tooling installation in [Makefile](Makefile).

### How can I customize the release

Take a look at GoReleaser [docs](https://goreleaser.com/customization/)
as well as [its repo](https://github.com/goreleaser/goreleaser/)
how it is dogfooding its functionality.
You can use it to add deb/rpm/snap packages, Homebrew Tap, Scoop App Manifest etc.

If you are developing a library and you like handcrafted changelog and release notes,
you are free to remove any usage of GoReleaser.

## Contributing

Feel free to create an issue or propose a pull request.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md).
