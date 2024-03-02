package data

import (
	"errors"
	"github.com/drademann/haora/app/datetime"
	"github.com/google/uuid"
	"slices"
	"time"
)

type Day struct {
	Id       uuid.UUID
	Date     time.Time
	Tasks    []Task
	Finished time.Time
}

func NewDay(date time.Time) Day {
	return Day{
		Id:       uuid.New(),
		Date:     date,
		Tasks:    []Task{},
		Finished: time.Time{},
	}
}

func (d *Day) IsEmpty() bool {
	return len(d.Tasks) == 0
}

func (d *Day) IsToday() bool {
	return isSameDay(d.Date, datetime.Now())
}

// AddNewTask creates a new task.
//
// The new task starts at the start timestamp with given text and tags.
// The date part of the start timestamp is not used, instead the day's date is applied.
// If a task at the specific timestamp already exists, it will be updated instead of added.
func (d *Day) AddNewTask(s time.Time, tx string, tgs []string) error {
	s = datetime.Combine(d.Date, s)
	t, err := d.taskAt(s)
	if err != nil {
		if errors.Is(err, NoTask) {
			t = NewTask(s, tx, tgs...)
			d.AddTasks(t)
		} else {
			return err
		}
	} else {
		t = t.with(s, tx, tgs...)
		d.updateTask(t)
	}
	State.DayList.update(*d)
	return nil
}

func (d *Day) AddTasks(tasks ...Task) {
	d.Tasks = append(d.Tasks, tasks...)
	slices.SortFunc(d.Tasks, tasksByStart)
}

func (d *Day) Finish(f time.Time) {
	f = datetime.Combine(d.Date, f)
	d.Finished = f
	State.DayList.update(*d)
}

func (d *Day) TotalDuration() time.Duration {
	if len(d.Tasks) == 0 {
		return 0
	}
	start := d.Tasks[0].Start
	end := datetime.Now()
	if !d.Finished.IsZero() {
		end = d.Finished
	}
	return end.Sub(start)
}

func (d *Day) TotalBreakDuration() time.Duration {
	var sum time.Duration = 0
	for _, t := range d.Tasks {
		if t.IsBreak {
			sum += d.TaskDuration(t)
		}
	}
	return sum
}

func (d *Day) TotalWorkDuration() time.Duration {
	var sum time.Duration = 0
	for _, t := range d.Tasks {
		if !t.IsBreak {
			sum += d.TaskDuration(t)
		}
	}
	return sum
}

func (d *Day) TotalTagDuration(tag string) time.Duration {
	var sum time.Duration = 0
	for _, t := range d.Tasks {
		if slices.Contains(t.Tags, tag) {
			sum += d.TaskDuration(t)
		}
	}
	return sum
}

func (d *Day) TaskDuration(task Task) time.Duration {
	s, err := d.Succ(task)
	if errors.Is(err, NoTaskSucc) {
		if d.Finished.IsZero() {
			return datetime.Now().Sub(task.Start)
		}
		return d.Finished.Sub(task.Start)
	}
	return s.Start.Sub(task.Start)
}

var (
	NoTask     = errors.New("no task")
	NoTaskSucc = errors.New("no succeeding task")
	NoTaskPred = errors.New("no preceding task")
)

func (d *Day) Succ(task Task) (Task, error) {
	for i, t := range d.Tasks {
		if t.Id == task.Id {
			j := i + 1
			if j < len(d.Tasks) {
				return d.Tasks[j], nil
			}
		}
	}
	return task, NoTaskSucc
}

func (d *Day) pred(task Task) (Task, error) {
	for i, t := range d.Tasks {
		if t.Id == task.Id {
			j := i - 1
			if j >= 0 {
				return d.Tasks[j], nil
			}
		}
	}
	return task, NoTaskPred
}

func (d *Day) taskAt(start time.Time) (Task, error) {
	for _, t := range d.Tasks {
		if t.Start.Hour() == start.Hour() && t.Start.Minute() == start.Minute() {
			return t, nil
		}
	}
	return Task{}, NoTask
}

func (d *Day) updateTask(task Task) {
	for i, t := range d.Tasks {
		if t.Id == task.Id {
			d.Tasks[i] = task
		}
	}
}

func (d *Day) Tags() []string {
	set := make(map[string]struct{})
	for _, t := range d.Tasks {
		for _, tag := range t.Tags {
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
