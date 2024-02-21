package app

import "time"

type Day struct {
	Date     time.Time
	Tasks    []Task
	Finished time.Time
}

func NewDay(date time.Time) *Day {
	return &Day{
		Date:     date,
		Tasks:    []Task{},
		Finished: time.Time{},
	}
}

func IsSameDay(date1, date2 time.Time) bool {
	return date1.Location() == date2.Location() && date1.Day() == date2.Day() && date1.Month() == date2.Month() && date1.Year() == date2.Year()
}
