package app

import "time"

func MockDate(year int, month time.Month, day, hour, minute int) time.Time {
	return time.Date(year, month, day, hour, minute, 0, 0, time.Local)
}

func MockTime(hour, minute int) time.Time {
	return MockDate(2024, time.June, 21, hour, minute)
}
