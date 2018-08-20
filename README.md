# epoch-time
[![GoDoc](https://godoc.org/github.com/hanakoa/epoch-time?status.svg)](https://godoc.org/github.com/hanakoa/epoch-time)
[![Go report](http://goreportcard.com/badge/hanakoa/epoch-time)](http://goreportcard.com/report/hanakoa/epoch-time)

This library offers a timestamp struct that is SQL-compatible and marshals/unmarshals as epoch seconds.
It also offers a nullable version.

## Why?
Traditionally, `time.Time` gets serialized and deserialized using RFC3339 rather than epoch seconds.
There is a [workaround](https://stackoverflow.com/questions/23695479/format-timestamp-in-outgoing-json-in-golang)
you can use to create your own wrapper over `time.Time`.

However, there are two additional requirements many apps need:
- timestamps need SQL support
- timestamps need to be nullable sometimes

For SQL support, we implement the [Value](https://golang.org/pkg/database/sql/driver/#Value) interface.
For nullability, we borrow heavily from [guregu/null](https://github.com/guregu/null).

## Usage
```golang
package main

import (
	"time"
	"github.com/hanakoa/epoch-time"
)

func main() {
	// Make a time.Time
	t := time.Date(1993, 04, 17, 23, 0, 0, 0, time.UTC)

	// Make a Time
	e := epoch.Time(t)

	b, err := json.Marshal(e)

	// prints 735087600
	log.Println(string(b))
}
```
