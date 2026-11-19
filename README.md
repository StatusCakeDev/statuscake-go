# statuscake-go [![test](https://github.com/StatusCakeDev/statuscake-go/actions/workflows/test.yaml/badge.svg?event=push)](https://github.com/StatusCakeDev/statuscake-go/actions/workflows/test.yaml)

The [Go](https://golang.org/) implementation of the [StatusCake API][api-docs]
client. Documentation for this library can be found on the [StatusCake Developer
Portal][guides].

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
  "os"

  "github.com/StatusCakeDev/statuscake-go"
  "github.com/StatusCakeDev/statuscake-go/credentials"
)

func main() {
  bearer := credentials.NewBearerWithStaticToken(os.Getenv("STATUSCAKE_API_TOKEN"))
  client := statuscake.NewClient(statuscake.WithRequestCredentials(bearer))

  tests, err := client.ListUptimeTests(context.Background()).Execute()
  if err != nil {
    fmt.Println(err)
    fmt.Printf("%+v\n", statuscake.Errors(err))
    os.Exit(1)
  }

  fmt.Printf("%+v\n", tests.Data)
}
```

## Examples

The [`examples`](examples) directory has a program for each kind of check, which
creates, reads, updates, lists and deletes it. They create resources in your
account, and delete them again before exiting:

```bash
export STATUSCAKE_API_TOKEN=...
go run ./examples/uptime
```

## License

This project is licensed under the [MIT License](LICENSE).

[api-docs]: https://developers.statuscake.com/api/
[guides]: https://developers.statuscake.com/guides/sdks/introduction
