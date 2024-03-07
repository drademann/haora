package datetime

import (
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t.Run("should return the current time without seconds nor nanoseconds", func(t *testing.T) {
		n := Now()

		if n.Second() != 0 {
			t.Errorf("expected seconds to be 0, got %d", n.Second())
		}
		if n.Nanosecond() != 0 {
			t.Errorf("expected nanoseconds to be 0, got %d", n.Nanosecond())
		}
	})
}

func TestFindWeekday(t *testing.T) {
	d := FindWeekday(test.Date("05.03.2024 10:42"), Previous, time.Monday)

	got := d.Format("02.01.2006")
	want := "04.03.2024"
	if got != want {
		t.Errorf("expected %s, got %s", want, got)
	}
}

func TestFindWeekday_shouldNotFindToday(t *testing.T) {
	d := FindWeekday(test.Date("05.03.2024 10:42"), Previous, time.Tuesday)

	got := d.Format("02.01.2006")
	want := "27.02.2024"
	if got != want {
		t.Errorf("expected %s, got %s", want, got)
	}
}
