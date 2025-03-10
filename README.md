# wrike.go [![Go](https://github.com/golyakov/wrike.go/workflows/Go/badge.svg?branch=master)](https://github.com/golyakov/wrike.go/actions) [![PkgGoDev](https://pkg.go.dev/badge/github.com/golyakov/wrike.go)](https://pkg.go.dev/github.com/golyakov/wrike.go) [![Go Report Card](https://goreportcard.com/badge/github.com/golyakov/wrike.go)](https://goreportcard.com/report/github.com/golyakov/wrike.go)

Go implementation of [Wrike API](https://developers.wrike.com/documentation/api/overview) client

## Basic usage

```go
package main

import (
    "fmt"

    "github.com/golyakov/wrike.go"
)

func main() {
    conf := wrike.NewConfig("wrike-access-token", "") // Default host name is "app-eu.wrike.com"
    // To set a different host name:
    // conf := wrike.NewConfig("wrike-access-token", "www.wrike.com")
    api := wrike.NewAPI(conf)

    user, err := api.QueryUser("KUAAAA3E")
    if err != nil {
      panic(err)
    }

    fmt.Println(user.Kind)                      // => "users"
    fmt.Println(user.Data[0].ID)                // => "KUAAAA3E"
    fmt.Println(user.Data[0].Profiles[0].Email) // => "kqri7kgjlb@y21z0uysjx.com"
}
```
