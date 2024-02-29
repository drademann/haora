package cmd

import (
	"testing"
	"time"
)

func TestDay(t *testing.T) {
	existingDay := Day{
		Date: mockTime(9, 0),
		Tasks: []Task{
			{Start: mockTime(9, 0),
				Text:    "a task",
				IsPause: false,
				Tags:    []string{}},
		},
		Finished: time.Time{}}
	ctx.data = dayList{
		days: []Day{existingDay},
	}

	t.Run("should return day if it exists", func(t *testing.T) {
		date := mockTime(13, 48)

		day := ctx.data.day(date)

		if !isSameDay(day.Date, existingDay.Date) {
			t.Errorf("got unexpected task: %+v", day)
		}
		if len(ctx.data.days) != 1 {
			t.Errorf("number of days should'nt have changed, but is now %d", len(ctx.data.days))
		}
	})
	t.Run("should create a new day if it doesn't exist", func(t *testing.T) {
		date := mockDate(2024, time.June, 30, 10, 0)

		day := ctx.data.day(date)

		if !isSameDay(day.Date, date) {
			t.Errorf("got unexpected task: %+v", day)
		}
		if len(ctx.data.days) != 2 {
			t.Errorf("number of days should have increased to 2, but is now %d", len(ctx.data.days))
		}
		if _ = ctx.data.day(date); len(ctx.data.days) != 2 {
			t.Errorf("number of days shouldn't have increased calling the same date a second time")
		}
	})
}
