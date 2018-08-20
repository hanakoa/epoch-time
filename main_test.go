package epoch

import (
	"testing"
	"os"
	"time"
	"log"
)

var Birthday = time.Date(1993, 04, 17, 23, 0, 0, 0, time.UTC)

func TestMain(m *testing.M) {
	code := m.Run()
	log.Println("Birthday is", Birthday.Unix())
	os.Exit(code)
}