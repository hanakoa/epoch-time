package epoch

import (
	"testing"
	"time"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
)

type Event struct {
	StartTime Time     `json:"start_time"`
	EndTime   NullTime `json:"end_time"`
}

func TestNullTimeMarshal(t *testing.T) {
	Convey("Given an event with no end time", t, func() {
		p := Event{
			StartTime: Time(time.Date(1993, 04, 17, 23, 0, 0, 0, time.UTC)),
			EndTime:   NullTimeFromPtr(nil),
		}
		b, err := json.Marshal(p)
		So(err, ShouldBeNil)
		s := string(b)
		log.Println(s)
		So(s, ShouldEqual, `{"start_time":735087600,"end_time":null}`)
		So(p.EndTime.Valid, ShouldBeFalse)
	})
	Convey("Given an event with an end time", t, func() {
		at := time.Date(1993, 04, 17, 23, 0, 0, 0, time.UTC)
		p := Event{
			StartTime: Time(time.Date(1993, 04, 17, 23, 0, 0, 0, time.UTC)),
			EndTime:   NullTimeFromPtr(&at),
		}
		b, err := json.Marshal(p)
		So(err, ShouldBeNil)
		s := string(b)
		log.Println(s)
		So(s, ShouldEqual, `{"start_time":735087600,"end_time":735087600}`)
	})
}

func TestNullTimeUnmarshal(t *testing.T) {
	Convey("Given a JSON string without an end time", t, func() {
		p := Event{}
		err := json.Unmarshal([]byte(`{"start_time":735087600,"end_time":null}`), &p)
		So(err, ShouldBeNil)
		tt := time.Time(p.StartTime)
		So(tt.Year(), ShouldEqual, 1993)
		So(tt.Month(), ShouldEqual, time.April)
		So(tt.Day(), ShouldEqual, 17)
		So(tt.Hour(), ShouldEqual, 23)
		So(p.EndTime.Valid, ShouldBeFalse)
	})
	Convey("Given a JSON string with an end time", t, func() {
		p := Event{}
		err := json.Unmarshal([]byte(`{"start_time":735087600,"end_time":735087600}`), &p)
		So(err, ShouldBeNil)
		So(p.EndTime.Valid, ShouldBeTrue)
		tt := time.Time(p.EndTime.Time)
		So(tt.Year(), ShouldEqual, 1993)
		So(tt.Month(), ShouldEqual, time.April)
		So(tt.Day(), ShouldEqual, 17)
		So(tt.Hour(), ShouldEqual, 23)
	})
}