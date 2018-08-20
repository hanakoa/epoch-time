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

func TestTimeMarshal(t *testing.T) {
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

func TestTimeUnmarshal(t *testing.T) {
	Convey("Given a JSON string", t, func() {
		p := Person{}
		err := json.Unmarshal([]byte(`{"name":"Kevin Chen","birth_time":735087600}`), &p)
		So(err, ShouldBeNil)
		So(p.Name, ShouldEqual, "Kevin Chen")
		So(p.BirthTime, ShouldNotBeNil)
		t := time.Time(p.BirthTime)
		So(t.Year(), ShouldEqual, 1993)
		So(t.Month(), ShouldEqual, time.April)
		So(t.Day(), ShouldEqual, 17)
		So(t.Hour(), ShouldEqual, 23)
	})
}