package app

import (
	"errors"
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
// If a task at the specific timestamp already exists, it will be updated instead of added.
func (d *DayList) AddNewTask(start time.Time, text string, tags []string) error {
	day := d.Day(WorkingDate)
	task, err := day.taskAt(start)
	if err != nil {
		if errors.Is(err, NoTask) {
			task = newTask(start, text, tags)
			day.Tasks = append(day.Tasks, task)
		} else {
			return err
		}
	} else {
		task = task.with(start, text, tags)
		day.update(task)
	}
	d.update(day)
	return nil
}

func (d *DayList) update(day Day) {
	for i, e := range d.Days {
		if e.Id == day.Id {
			d.Days[i] = day
		}
	}
}
