package app_test

import (
	"haora/app"
	"testing"
	"time"
)

func TestDay(t *testing.T) {
	existingDay := app.Day{
		Date: time.Date(2024, time.February, 22, 6, 24, 13, 0, time.Local),
		Tasks: []app.Task{
			{Start: time.Date(2024, time.February, 9, 0, 0, 0, 0, time.Local),
				Text:    "a task",
				IsPause: false,
				Tags:    []string{}},
		},
		Finished: time.Time{}}
	app.Data = app.DayList{
		existingDay,
	}

	t.Run("should return day if it exists", func(t *testing.T) {
		date := time.Date(2024, time.February, 22, 10, 0, 0, 0, time.Local)

		day := app.Data.Day(date)

		if !app.IsSameDay(day.Date, existingDay.Date) {
			t.Errorf("got unexpected task: %+v", day)
		}
		if len(app.Data) != 1 {
			t.Errorf("number of days should'nt have changed, but is now %d", len(app.Data))
		}
	})
	t.Run("should create a new day if it doesn't exist", func(t *testing.T) {
		date := time.Date(2024, time.June, 30, 10, 0, 0, 0, time.Local)

		day := app.Data.Day(date)

		if !app.IsSameDay(day.Date, date) {
			t.Errorf("got unexpected task: %+v", day)
		}
		if len(app.Data) != 2 {
			t.Errorf("number of days should have increased to 2, but is now %d", len(app.Data))
		}
	})
}
