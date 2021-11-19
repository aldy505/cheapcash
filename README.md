# Cheapcash

[![Go Reference](https://pkg.go.dev/badge/github.com/aldy505/cheapcash.svg)](https://pkg.go.dev/github.com/aldy505/cheapcash) [![Go Report Card](https://goreportcard.com/badge/github.com/aldy505/cheapcash)](https://goreportcard.com/report/github.com/aldy505/cheapcash) ![GitHub](https://img.shields.io/github/license/aldy505/cheapcash) [![CodeFactor](https://www.codefactor.io/repository/github/aldy505/cheapcash/badge)](https://www.codefactor.io/repository/github/aldy505/cheapcash) [![codecov](https://codecov.io/gh/aldy505/cheapcash/branch/master/graph/badge.svg?token=Noeexg5xEJ)](https://codecov.io/gh/aldy505/cheapcash) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/9b78970127c74c1a923533e05f65848d)](https://www.codacy.com/gh/aldy505/cheapcash/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=aldy505/cheapcash&amp;utm_campaign=Badge_Grade) [![Test and coverage](https://github.com/aldy505/cheapcash/actions/workflows/ci.yml/badge.svg)](https://github.com/aldy505/cheapcash/actions/workflows/ci.yml)

SSD is cheap. Why don't we use it for caching?

A simple library implementing filesystem I/O as a cache. Should be must useful when used again a Solid State Drive
for maximum speed and to handle good amount of concurrent read/write.

The API itself is also pretty simple considering I don't want this to be a full-blown caching library like Redis,
I just want it to be simple like Bigcache or similar caching library.

## Install

```go
import "github.com/aldy505/cheapcash"
```

## Usage

It has simple API for reading & storing cache.

```go
package main

import (
  "log"

  "github.com/aldy505/cheapcash"
)

func main() {
  // Create a Cheapcash instance.
  // Of course you can make multiple instance for multiple
  // root directories.
  cache := cheapcash.New("/tmp/cheapcash")
  // or if you are feeling lazy
  cache = cheapcash.Default()
  // path defaults to /tmp/cheapcash

  err := cache.Write("users:list", usersList)
  if err != nil {
    log.Fatal(err)
  }

  val, err := cache.Read("users:list")
  if err != nil {
    log.Fatal(err)
  }

  log.Println(string(val))

  err = cache.Append("users:list", []byte("\nMarcel"))
  if err != nil {
    log.Fatal(err)
  }

  err = cache.Delete("users:list")
  if err != nil {
    log.Fatal(err)
  }
}
```

See Godoc documentation (link above, beneath the title) for more complete documentation of the package.

## License

[MIT](./LICENSE)
