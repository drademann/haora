package cmd

import (
	"time"
)

type dayList struct {
	days []day
}

// Day returns the Day struct for the specified date.
//
// The returned struct is a copy of the day.
// Changes to this day won't be applied to the data model automatically.
func (d *dayList) day(date time.Time) day {
	for _, day := range d.days {
		if isSameDay(day.date, date) {
			return day
		}
	}
	newDay := newDay(date)
	d.days = append(d.days, newDay)
	return newDay
}

func (d *dayList) update(day day) {
	for i, e := range d.days {
		if e.id == day.id {
			d.days[i] = day
		}
	}
}
