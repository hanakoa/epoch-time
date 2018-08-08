package epoch

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
	"time"
)

type Person struct {
	Name      string `json:"name"`
	BirthTime Time   `json:"birth_time"`
}

func TestTime(t *testing.T) {
	Convey("Given a struct", t, func() {
		p := Person{
			Name:      "Kevin Chen",
			BirthTime: Time(time.Date(1993, 04, 17, 23, 0, 0, 0, time.UTC)),
		}
		b, err := json.Marshal(p)
		So(err, ShouldBeNil)
		s := string(b)
		log.Println(s)
		So(s, ShouldEqual, `{"name":"Kevin Chen","birth_time":735087600}`)
	})
}
