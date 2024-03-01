package cmd

import (
	"errors"
	"time"
)

type dayList struct {
	days []day
}

// AddNewTask creates a new task for the current working date day.
//
// The new task starts at the start timestamp with given text and tags.
// The date part of the start timestamp is not used, instead the working date's date is applied.
// If a task at the specific timestamp already exists, it will be updated instead of added.
func addTask(start time.Time, text string, tags []string) error {
	day := ctx.data.day(ctx.workingDate)
	task, err := day.taskAt(start)
	if err != nil {
		if errors.Is(err, NoTask) {
			task = newTask(start, text, tags)
			day.addTasks(task)
		} else {
			return err
		}
	} else {
		task = task.with(start, text, tags)
		day.updateTask(task)
	}
	ctx.data.update(day)
	return nil
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
