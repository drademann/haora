package app

import (
	"time"
)

type DayList []Day

func (dl *DayList) Day(date time.Time) Day {
	for _, day := range *dl {
		if isSameDay(day.Date, date) {
			return day
		}
	}
	newDay := *NewDay(date)
	*dl = append(*dl, newDay)
	return newDay
}
