package epoch_time

import (
	"strconv"
	"time"
)

type EpochTime time.Time

func (t EpochTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	unix := time.Time(t).Unix()
	return []byte(strconv.FormatInt(unix, 10)), nil
}
