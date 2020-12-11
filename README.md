# Go-Scratchpad

Repository that contains experiments done in Golang, as a means to learn and test the language.

## Prerequisites

Either use Go on a local machine or copy the code to one of the many online playgrounds.
Code that relies on external libraries might not run as expected in a playground due to dependencies not being accessible.
Tested with Go version 1.5.

## Getting Started

To run this code on a local machine, pull the repository and place the source files in the GOPATH. It is not necessary to place the source files in the GOPATH, but it makes execution simpler.

Open a terminal in the GOPATH directory and run the code you want to test with:

```
go test ./src/github.com/HDauven/go-scratchpad/[package name]
```

To run documentation locally run the following command:

```
go doc -http :8080
```

The documentation for this repository can then be found at this URL: "http://localhost:8080/pkg/github.com/HDauven/go-scratchpad"

The top-level main package utilizes all the other packages that can be found in this repository. These packages contain implementations and tests. The main package shows how the functions in other packages can be utilized.

Run the main package with following command if you're in the GOPATH:

```
go run ./src/github.com/HDauven/go-scratchpad/main.go
```

## Contributing

Do you have any tips, tricks or general improvements to the repo?
In that case, feel free to contact me or to create a pull request with your improvements.
