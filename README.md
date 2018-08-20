# epoch-time
[![GoDoc](https://godoc.org/github.com/hanakoa/epoch-time?status.svg)](https://godoc.org/github.com/hanakoa/epoch-time)
[![Go report](http://goreportcard.com/badge/hanakoa/epoch-time)](http://goreportcard.com/report/hanakoa/epoch-time)

This library offers a wrapper over `time.Time` that gets serialized as epoch seconds
rather than RFC3339.

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

## Sources
This library draws heavily from [guregu/null](https://github.com/guregu/null).