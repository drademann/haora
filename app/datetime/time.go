package datetime

import "time"

// Now timestamp as variable to allow tests to override it.
//
// Seconds and nanoseconds are truncated and set to zero because all calculations in Haora are based on minutes.
var Now = NowFunc

func NowFunc() time.Time {
	return time.Now().Truncate(time.Minute)
}

func Combine(d, t time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), t.Hour(), t.Minute(), 0, 0, d.Location())
}
