package app

import (
	"haora/test"
	"testing"
	"time"
)

func TestDay(t *testing.T) {
	existingDay := Day{
		Date: test.MockTime(9, 0),
		Tasks: []Task{
			{Start: test.MockTime(9, 0),
				Text:    "a task",
				IsPause: false,
				Tags:    []string{}},
		},
		Finished: time.Time{}}
	Data = DayList{
		existingDay,
	}

	t.Run("should return day if it exists", func(t *testing.T) {
		date := test.MockTime(13, 48)

		day := Data.Day(date)

		if !isSameDay(day.Date, existingDay.Date) {
			t.Errorf("got unexpected task: %+v", day)
		}
		if len(Data) != 1 {
			t.Errorf("number of days should'nt have changed, but is now %d", len(Data))
		}
	})
	t.Run("should create a new day if it doesn't exist", func(t *testing.T) {
		date := test.MockDate(2024, time.June, 30, 10, 0)

		day := Data.Day(date)

		if !isSameDay(day.Date, date) {
			t.Errorf("got unexpected task: %+v", day)
		}
		if len(Data) != 2 {
			t.Errorf("number of days should have increased to 2, but is now %d", len(Data))
		}
	})
}
