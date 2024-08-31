# statuscake-go ![test](https://github.com/StatusCakeDev/statuscake-go/workflows/test/badge.svg)

The [Go](https://golang.org/) implementation of the [StatusCake
API](https://www.statuscake.com/api/v1) client. Documentation for this library
can be found [here](https://www.statuscake.com/api/v1).

## Prerequisites

You will need the following things properly installed on your computer:

- [Go](https://golang.org/): any one of the **three latest major**
  [releases](https://golang.org/doc/devel/release.html)

## Installation

With [Go module](https://github.com/golang/go/wiki/Modules) support (Go 1.11+),
add the following import

```go
import "github.com/StatusCakeDev/statuscake-go"
```

to your code, and then `go [build|run|test]` will automatically fetch the
necessary dependencies.

Otherwise, to install the `statuscake-go` package, run the following command:

```bash
go get -u github.com/StatusCakeDev/statuscake-go
```

## Usage

Within any Go file instantiate an API client and execute a request:

```go
package main

import (
  "context"
  "fmt"

  "github.com/StatusCakeDev/statuscake-go"
  "github.com/StatusCakeDev/statuscake-go/credentials"
)

func main() {
  bearer := credentials.NewBearerWithStaticToken(apiToken)
  client := statuscake.NewClient(statuscake.WithRequestCredentials(bearer))

  tests, err := client.ListUptimeTests(context.Background()).Execute()
  if err != nil {
    panic(err)
  }

  fmt.Printf("%+v\n", tests.Data)
}
```

## License

This project is licensed under the [MIT License](LICENSE).
