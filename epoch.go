package epoch

import (
	"strconv"
	"time"
)

// Time is a wrapper over time.Time that gets serialized
// as epoch seconds rather than RFC3339.
type Time time.Time

// MarshalJSON is called by json.Marshal to produce JSON.
// In this function we serialize Time as epoch seconds (an int).
func (t Time) MarshalJSON() ([]byte, error) {
	//do your serializing here
	unix := time.Time(t).Unix()
	return []byte(strconv.FormatInt(unix, 10)), nil
}

func (t Time) String() string {
	return time.Time(t).String()
}
