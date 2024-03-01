package cmd

import (
	"testing"
	"time"
)

func TestDay(t *testing.T) {
	existingDay := day{
		date: mockTime(9, 0),
		tasks: []Task{
			{start: mockTime(9, 0),
				text:    "a task",
				isPause: false,
				tags:    []string{}},
		},
		finished: time.Time{}}
	ctx.data = dayList{
		days: []day{existingDay},
	}

	t.Run("should return day if it exists", func(t *testing.T) {
		date := mockTime(13, 48)

		day := ctx.data.day(date)

		if !isSameDay(day.date, existingDay.date) {
			t.Errorf("got unexpected task: %+v", day)
		}
		if len(ctx.data.days) != 1 {
			t.Errorf("number of days should'nt have changed, but is now %d", len(ctx.data.days))
		}
	})
	t.Run("should create a new day if it doesn't exist", func(t *testing.T) {
		date := mockDate(2024, time.June, 30, 10, 0)

		day := ctx.data.day(date)

		if !isSameDay(day.date, date) {
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
