#Contributing
Welcome to Imerologio CLI project ! 

## About development
### Go 
Imerologio CLI is a Go application so make sure your have a development environment with Go 1.2+.  
[See the install instructions for Go](http://golang.org/doc/install.html).

### Dependencies
Install the following dependencies with `go get XXX`:
- [ishell](https://github.com/abiosoft/ishell) to ensure an interactive CLI : `go get github.com/fatih/color`
- [Color](https://github.com/fatih/color) in shell : `go get github.com/fatih/color`

## About releasing
This product is built using [goreleaser](https://goreleaser.com/), take a look at the config at [.goreleaser.yml](.goreleaser.yml).  
An automated build is launched on [Circle CI](https://circleci.com/gh/Agaetis-IT/imerologio-cli) each time a tag is push on this repository.