# go-obs-websocket

[![Build Status](https://travis-ci.com/christopher-dG/go-obs-websocket.svg?branch=master)](https://travis-ci.com/christopher-dG/go-obs-websocket)

`go-obs-websocket` is a package for interacting with [`obs-websocket`](https://github.com/Palakis/obs-websocket).

## Installation

```sh
go get github.com/christopher-dG/go-obs-websocket
```

## Usage

```go
package main

import (
    "log"

    obs "github.com/christopher-dG/go-obs-websocket"
)

client := obs.Client{Host: "localhost", Port: 4444}
if err := client.Connect(); err != nil {
    log.Fatal(err)
}
defer client.Disconnect()

// TODO
```
