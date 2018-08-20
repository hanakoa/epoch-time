# epoch-time
[![GoDoc](https://godoc.org/github.com/hanakoa/epoch-time?status.svg)](https://godoc.org/github.com/hanakoa/epoch-time)
[![Go report](http://goreportcard.com/badge/hanakoa/epoch-time)](http://goreportcard.com/report/hanakoa/epoch-time)
[![CircleCI](https://circleci.com/gh/hanakoa/epoch-time.svg?style=shield)](https://circleci.com/gh/hanakoa/epoch-time)
[![Coverage Status](https://coveralls.io/repos/github/hanakoa/epoch-time/badge.svg?branch=master)](https://coveralls.io/github/hanakoa/epoch-time?branch=master)
[![GitHub Release](https://img.shields.io/github/release/hanakoa/epoch-time.svg)](https://github.com/hanakoa/epoch-time/releases)

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

type Invitation struct {
	CreatedTime   epoch.Time     `json:"created_time"`
	ApprovedTime  epoch.NullTime `json:"approved_time"`
}

func main() {
	e := Event{
		CreatedTime: epoch.Time(time.Date(1993, 04, 17, 23, 0, 0, 0, time.UTC)),
		ApprovedTime:   epoch.NullTimeFromPtr(nil),
	}

	b, err := json.Marshal(e)

	// prints `{"created_time":735087600,"approved_time":null}`
	log.Println(string(b))

	e.ApprovedTime = epoch.Time(time.Date(1993, 04, 17, 23, 0, 0, 0, time.UTC))

	b, err := json.Marshal(e)

	// prints `{"created_time":735087600,"approved_time":735087600}`
	log.Println(string(b))
}
```
