package data

import (
	"github.com/drademann/haora/app"
	"time"
)

type Week struct {
	Days [7]Day
}

func CollectWeek(start time.Time) Week {
	var week Week
	var date = start
	for i := 0; i < 7; i++ {
		week.Days[i] = State.DayList.Day(date)
		date = date.Add(24 * time.Hour)
	}
	return week
}

func (w Week) TotalWorkDuration() time.Duration {
	total := 0 * time.Nanosecond
	for _, day := range w.Days {
		total += day.TotalWorkDuration()
	}
	return total
}

func (w Week) TotalOvertimeDuration() (time.Duration, error) {
	durationPerWeek, err := app.DurationPerWeek()
	if err != nil {
		return 0, err
	}
	return w.TotalWorkDuration() - durationPerWeek, nil
}
