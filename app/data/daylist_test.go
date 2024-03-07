package data

import (
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestDay(t *testing.T) {
	existingDay := Day{
		Date: test.Time("9:00"),
		Tasks: []Task{
			{Start: test.Time("9:00"),
				Text:    "a task",
				IsPause: false,
				Tags:    []string{}},
		},
		Finished: time.Time{}}
	State.DayList = DayListType{
		Days: []Day{existingDay},
	}

	t.Run("should return day if it exists", func(t *testing.T) {
		date := test.Time("13:48")

		d := State.DayList.Day(date)

		if !isSameDay(d.Date, existingDay.Date) {
			t.Errorf("got unexpected task: %+v", d)
		}
		if len(State.DayList.Days) != 1 {
			t.Errorf("number of days should'nt have changed, but is now %d", len(State.DayList.Days))
		}
	})
	t.Run("should create a new day if it doesn't exist", func(t *testing.T) {
		date := test.Date("30.06.2024 10:00")

		d := State.DayList.Day(date)

		if !isSameDay(d.Date, date) {
			t.Errorf("got unexpected task: %+v", d)
		}
		if len(State.DayList.Days) != 2 {
			t.Errorf("number of days should have increased to 2, but is now %d", len(State.DayList.Days))
		}
		if _ = State.DayList.Day(date); len(State.DayList.Days) != 2 {
			t.Errorf("number of days shouldn't have increased calling the same date a second time")
		}
	})
}
