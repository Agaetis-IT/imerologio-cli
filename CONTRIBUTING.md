# Contributing
Welcome to Imerologio CLI project ! 

## About development
### Go 
Imerologio CLI is a Go application so make sure you have a development environment with Go 1.2+.  
[See the install instructions for Go](http://golang.org/doc/install.html).

The project is structured like advised in [golang-standards](https://github.com/golang-standards/project-layout).

### Dependencies
Install the following dependencies with `go get XXX`:
- [survey](https://github.com/AlecAivazis/survey) to ensure an interactive CLI : `go get gopkg.in/AlecAivazis/survey.v1`
- [Color](https://github.com/fatih/color) in shell : `go get github.com/fatih/color`
- [pb](https://github.com/cheggaaa/pb) in shell : `go get gopkg.in/cheggaaa/pb.v1`

### Build
Launch `go build -o imerologio-cli cmd/imerologio_cli/imerologio_cli.go` to generate an executable named `imerologio-cli` at the root.

### Launch app
Simply run `./imerologio-cli` once build finished.

## About releasing
This product is built using [goreleaser](https://goreleaser.com/), take a look at the config at [.goreleaser.yml](.goreleaser.yml).  
An automated build is launched on [Circle CI](https://circleci.com/gh/Agaetis-IT/imerologio-cli) each time a tag is push on this repository.
