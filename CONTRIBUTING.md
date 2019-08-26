# Contributing

Welcome to Imerologio CLI project!

## About development

### Go

Imerologio CLI is a Go application so make sure you have a development environment with Go 1.12+.  
[See the install instructions for Go](http://golang.org/doc/install.html).

The project is structured like advised in [golang-standards](https://github.com/golang-standards/project-layout).

### Build

Launch `go build -o imerologio-cli cmd/imerologio_cli/main.go` to generate an executable named `imerologio-cli` at the root.

### Launch app

Simply run `./imerologio-cli` once build finished.

## About releasing

This product is built using [goreleaser](https://goreleaser.com/), take a look at the config at [.goreleaser.yml](.goreleaser.yml).  

### How to release

- push a tag with a pattern like `v0.1.4`
- an automated build is launched on [Circle CI](https://circleci.com/gh/Agaetis-IT/imerologio-cli)
- a new release (named like the tag) is pushed to [Github releases](https://github.com/Agaetis-IT/imerologio-cli/releases).
