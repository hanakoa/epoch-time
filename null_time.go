package epoch

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// NullTime is a nullable time.Time. It supports SQL and JSON serialization.
// It will marshal to null if null.
type NullTime struct {
	Time  Time
	Valid bool
}

// Scan implements the Scanner interface.
func (t *NullTime) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case time.Time:
		t.Time = Time(x)
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", value, value)
	}
	t.Valid = err == nil
	return err
}

// Value implements the driver Valuer interface.
func (t NullTime) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return time.Time(t.Time), nil
}

// NewTime creates a new NullTime.
func NewNullTime(t time.Time, valid bool) NullTime {
	return NullTime{
		Time:  Time(t),
		Valid: valid,
	}
}

// NullTimeFrom creates a new NullTime that will always be valid.
func NullTimeFrom(t time.Time) NullTime {
	return NewNullTime(t, true)
}

// NullTimeFromPtr creates a new NullTime that will be null if t is nil.
func NullTimeFromPtr(t *time.Time) NullTime {
	if t == nil {
		return NewNullTime(time.Time{}, false)
	}
	return NewNullTime(*t, true)
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (t NullTime) ValueOrZero() time.Time {
	if !t.Valid {
		return time.Time{}
	}
	return time.Time(t.Time)
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t NullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return Time(t.Time).MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string, object (e.g. pq.NullTime and friends)
// and null input.
func (t *NullTime) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case float64:
		et := Time(t.Time)
		err = (&et).UnmarshalJSON(data)
		t.Time = et
		t.Valid = true
	case string:
		et := Time(t.Time)
		err = (&et).UnmarshalJSON(data)
		t.Time = et
		t.Valid = true
	case map[string]interface{}:
		ti, tiOK := x["Time"].(string)
		valid, validOK := x["Valid"].(bool)
		if !tiOK || !validOK {
			return fmt.Errorf(`json: unmarshalling object into Go value of type null.Time requires key "Time" to be of type string and key "Valid" to be of type bool; found %T and %T, respectively`, x["Time"], x["Valid"])
		}
		tt := time.Time(t.Time)
		err = (&tt).UnmarshalText([]byte(ti))
		t.Valid = valid
		return err
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type null.Time", reflect.TypeOf(v).Name())
	}
	t.Valid = err == nil
	return err
}

func (t NullTime) MarshalText() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return time.Time(t.Time).MarshalText()
}

func (t *NullTime) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		t.Valid = false
		return nil
	}
	tt := time.Time(t.Time)
	if err := (&tt).UnmarshalText(text); err != nil {
		return err
	}
	t.Valid = true
	return nil
}

// SetValid changes this NullTime's value and sets it to be non-null.
func (t *NullTime) SetValid(v time.Time) {
	t.Time = Time(v)
	t.Valid = true
}

// Ptr returns a pointer to this Time's value, or a nil pointer if this Time is null.
func (t NullTime) Ptr() *time.Time {
	if !t.Valid {
		return nil
	}
	tt := time.Time(t.Time)
	return &tt
}
