package datetime

import "testing"

func TestNow(t *testing.T) {
	t.Run("should return the current time without seconds nor nanoseconds", func(t *testing.T) {
		n := Now()

		if n.Second() != 0 {
			t.Errorf("expected seconds to be 0, but got %d", n.Second())
		}
		if n.Nanosecond() != 0 {
			t.Errorf("expected nanoseconds to be 0, but got %d", n.Nanosecond())
		}
	})
}
