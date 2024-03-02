package data

import (
	"time"
)

type DayListType struct {
	Days []Day
}

// Day returns the Day struct for the specified date.
//
// The returned struct is a copy of the day.
// Changes to this day won't be applied to the data model automatically.
func (d *DayListType) Day(date time.Time) Day {
	for _, day := range d.Days {
		if isSameDay(day.Date, date) {
			return day
		}
	}
	day := NewDay(date)
	d.Days = append(d.Days, day)
	return day
}

func (d *DayListType) update(day Day) {
	for i, e := range d.Days {
		if e.Id == day.Id {
			d.Days[i] = day
		}
	}
}
