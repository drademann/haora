package test

import (
	"time"
)

func MockDate(dateStr string) time.Time {
	tm, err := time.Parse("02.01.2006 15:04", dateStr)
	if err != nil {
		panic(err)
	}
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), 0, 0, time.Local)
}

func MockTime(timeStr string) time.Time {
	return MockDate("21.06.2024 " + timeStr)
}
