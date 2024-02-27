package app

import (
	"time"
)

type DayList struct {
	Days []Day
}

// Day returns the Day struct for the specified date.
//
// The returned struct is a copy of the day.
// Changes to this day won't be applied to the data model automatically.
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

// AddNewTask creates a new task for the current working date day.
//
// The new task starts at the start timestamp with given text and tags.
// The date part of the start timestamp is not used, instead the working date's date is applied.
func (d *DayList) AddNewTask(start time.Time, text string, tags []string) {
	task := NewTask(start, text, tags)
	day := d.Day(WorkingDate)
	day.Tasks = append(day.Tasks, task)
	d.update(day)
}

func (d *DayList) update(day Day) {
	for i, e := range d.Days {
		if isSameDay(e.Date, day.Date) {
			d.Days = append(d.Days[:i], d.Days[i+1:]...)
		}
	}
	d.Days = append(d.Days, day)
}
