package test

import (
	"time"
)

func Date(dateStr string) time.Time {
	tm, err := time.Parse("02.01.2006 15:04", dateStr)
	if err != nil {
		panic(err)
	}
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), 0, 0, time.Local)
}

func Time(timeStr string) time.Time {
	return Date("21.06.2024 " + timeStr)
}
