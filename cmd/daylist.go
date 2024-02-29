package cmd

import (
	"errors"
	"time"
)

type dayList struct {
	days []Day
}

// AddNewTask creates a new task for the current working date day.
//
// The new task starts at the start timestamp with given text and tags.
// The date part of the start timestamp is not used, instead the working date's date is applied.
// If a task at the specific timestamp already exists, it will be updated instead of added.
func AddNewTask(start time.Time, text string, tags []string) error {
	day := ctx.data.day(ctx.workingDate)
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
	ctx.data.update(day)
	return nil
}

// Day returns the Day struct for the specified date.
//
// The returned struct is a copy of the day.
// Changes to this day won't be applied to the data model automatically.
func (d *dayList) day(date time.Time) Day {
	for _, day := range d.days {
		if isSameDay(day.Date, date) {
			return day
		}
	}
	newDay := NewDay(date)
	d.days = append(d.days, newDay)
	return newDay
}

func (d *dayList) update(day Day) {
	for i, e := range d.days {
		if e.Id == day.Id {
			d.days[i] = day
		}
	}
}
