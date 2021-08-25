# statuscake-go ![test](https://github.com/StatusCakeDev/statuscake-go/workflows/test/badge.svg)

The [Go](https://golang.org/) implementation of the [StatusCake
API](https://www.statuscake.com/api/v1) client. Documentation for this library
can be found [here](https://www.statuscake.com/api/v1).

## Prerequisites

You will need the following things properly installed on your computer.

* [Go](https://golang.org/): any one of the **three latest major**
  [releases](https://golang.org/doc/devel/release.html)

## Installation

With [Go module](https://github.com/golang/go/wiki/Modules) support (Go 1.11+),
simply add the following import

```go
import "github.com/StatusCakeDev/statuscake-go"
```

to your code, and then `go [build|run|test]` will automatically fetch the
necessary dependencies.

Otherwise, to install the `statuscake-go` package, run the following command:

```bash
$ go get -u github.com/StatusCakeDev/statuscake-go
```

## License

This project is licensed under the [MIT License](LICENSE.md).
