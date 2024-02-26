package app

import (
	"time"
)

type DayList struct {
	Days []Day
}

func (d *DayList) Day(date time.Time) Day {
	for _, day := range d.Days {
		if isSameDay(day.Date, date) {
			return day
		}
	}
	newDay := NewDay(date)
	d.Days = append(d.Days, newDay)
	return newDay
}

func (d *DayList) UpdateDay(day Day) {
	for i, e := range d.Days {
		if isSameDay(e.Date, day.Date) {
			d.Days = append(d.Days[:i], d.Days[i+1:]...)
		}
	}
	d.Days = append(d.Days, day)
}
