# epoch-time
This library offers a wrapper over `time.Time` that gets serialized as epoch seconds
rather than RFC3339.

## Usage
```golang
// Make a time.Time
t := time.Date(1993, 04, 17, 23, 0, 0, 0, time.UTC)

// Make an Time
e := Time(t)

b, err := json.Marshal(e)

// prints 735087600
log.Println(string(b))
```
