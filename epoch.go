package epoch

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// Time is a wrapper over time.Time that gets serialized
// as epoch seconds rather than RFC3339.
// It supports SQL and JSON serialization.
type Time time.Time

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in epoch seconds format.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(time.Time(t).Unix()))), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in epoch seconds format.
func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	var err error
	i, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	tt := time.Unix(int64(i), 0).UTC()
	*t = Time(tt)
	return nil
}

// Scan implements the Scanner interface.
func (t *Time) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case time.Time:
		*t = Time(x)
	case nil:
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into Time: %v", value, value)
	}
	return err
}

// Value implements the driver Valuer interface.
func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t Time) String() string {
	return time.Time(t).String()
}
