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

// AddNewTask creates a new task for the current working date day,
// starting at start timestamp with given text and tags.
// The date part of the start timestamp is not used, instead the working date
// is applied.
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
