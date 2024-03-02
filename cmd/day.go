package cmd

import (
	"errors"
	"github.com/google/uuid"
	"slices"
	"time"
)

type day struct {
	id       uuid.UUID
	date     time.Time
	tasks    []Task
	finished time.Time
}

func newDay(date time.Time) day {
	return day{
		id:       uuid.New(),
		date:     date,
		tasks:    []Task{},
		finished: time.Time{},
	}
}

func (d *day) isEmpty() bool {
	return len(d.tasks) == 0
}

func (d *day) isToday() bool {
	return isSameDay(d.date, now())
}

// addNewTask creates a new task.
//
// The new task starts at the start timestamp with given text and tags.
// The date part of the start timestamp is not used, instead the day's date is applied.
// If a task at the specific timestamp already exists, it will be updated instead of added.
func (d *day) addNewTask(s time.Time, tx string, tgs []string) error {
	s = combineDateTime(d.date, s)
	t, err := d.taskAt(s)
	if err != nil {
		if errors.Is(err, NoTask) {
			t = newTask(s, tx, tgs...)
			d.addTasks(t)
		} else {
			return err
		}
	} else {
		t = t.with(s, tx, tgs...)
		d.updateTask(t)
	}
	ctx.data.update(*d)
	return nil
}

func (d *day) addTasks(tasks ...Task) {
	d.tasks = append(d.tasks, tasks...)
	slices.SortFunc(d.tasks, tasksByStart)
}

func (d *day) finish(f time.Time) {
	f = combineDateTime(d.date, f)
	d.finished = f
	ctx.data.update(*d)
}

func (d *day) totalDuration() time.Duration {
	if len(d.tasks) == 0 {
		return 0
	}
	start := d.tasks[0].start
	end := now()
	if !d.finished.IsZero() {
		end = d.finished
	}
	return end.Sub(start)
}

func (d *day) totalBreakDuration() time.Duration {
	var sum time.Duration = 0
	for _, t := range d.tasks {
		if t.isPause {
			sum += d.taskDuration(t)
		}
	}
	return sum
}

func (d *day) totalWorkDuration() time.Duration {
	var sum time.Duration = 0
	for _, t := range d.tasks {
		if !t.isPause {
			sum += d.taskDuration(t)
		}
	}
	return sum
}

func (d *day) totalTagDuration(tag string) time.Duration {
	var sum time.Duration = 0
	for _, t := range d.tasks {
		if slices.Contains(t.tags, tag) {
			sum += d.taskDuration(t)
		}
	}
	return sum
}

func (d *day) taskDuration(task Task) time.Duration {
	s, err := d.succ(task)
	if errors.Is(err, NoTaskSucc) {
		if d.finished.IsZero() {
			return now().Sub(task.start)
		}
		return d.finished.Sub(task.start)
	}
	return s.start.Sub(task.start)
}

var (
	NoTask     = errors.New("no task")
	NoTaskSucc = errors.New("no succeeding task")
	NoTaskPred = errors.New("no preceding task")
)

func (d *day) succ(task Task) (Task, error) {
	for i, t := range d.tasks {
		if t.id == task.id {
			j := i + 1
			if j < len(d.tasks) {
				return d.tasks[j], nil
			}
		}
	}
	return task, NoTaskSucc
}

func (d *day) pred(task Task) (Task, error) {
	for i, t := range d.tasks {
		if t.id == task.id {
			j := i - 1
			if j >= 0 {
				return d.tasks[j], nil
			}
		}
	}
	return task, NoTaskPred
}

func (d *day) taskAt(start time.Time) (Task, error) {
	for _, t := range d.tasks {
		if t.start.Hour() == start.Hour() && t.start.Minute() == start.Minute() {
			return t, nil
		}
	}
	return Task{}, NoTask
}

func (d *day) updateTask(task Task) {
	for i, t := range d.tasks {
		if t.id == task.id {
			d.tasks[i] = task
		}
	}
}

func (d *day) tags() []string {
	set := make(map[string]struct{})
	for _, t := range d.tasks {
		for _, tag := range t.tags {
			set[tag] = struct{}{}
		}
	}
	var res []string
	for tag := range set {
		res = append(res, tag)
	}
	slices.Sort(res)
	return res
}

func isSameDay(date1, date2 time.Time) bool {
	return date1.Location() == date2.Location() && date1.Day() == date2.Day() && date1.Month() == date2.Month() && date1.Year() == date2.Year()
}
