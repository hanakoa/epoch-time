# epoch-time
This library offers a wrapper over `time.Time` that gets serialized as epoch seconds
rather than RFC3339.

## Usage
```golang
package main

import (
	"time"
	epoch "github.com/hanakoa/epoch-time"
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
