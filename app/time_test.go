package app

import "testing"

func TestNow(t *testing.T) {
	t.Run("should return the current time wihout seconds nor nanoseconds", func(t *testing.T) {
		now := Now()

		if now.Second() != 0 {
			t.Errorf("expected seconds to be 0, but got %d", now.Second())
		}
		if now.Nanosecond() != 0 {
			t.Errorf("expected nanoseconds to be 0, but got %d", now.Nanosecond())
		}
	})
}
