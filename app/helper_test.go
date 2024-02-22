package app_test

import "time"

func testDate(year int, month time.Month, day, hour, minute int) time.Time {
	return time.Date(year, month, day, hour, minute, 0, 0, time.Local)
}

func testTime(hour, minute int) time.Time {
	return testDate(2024, time.June, 21, hour, minute)
}
