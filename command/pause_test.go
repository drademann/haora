package command

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestPause(t *testing.T) {
	now := test.MockDate("03.03.2024 17:27")
	test.MockNowAt(t, now)

	prepareTestDay := func() {
		d := data.NewDay(test.MockDate("03.03.2024 00:00"))
		d.AddTasks(data.NewTask(test.MockDate("03.03.2024 9:00"), "a task", "Haora"))
		data.State.DayList = data.DayListType{Days: []data.Day{d}}
	}

	testCases := []struct {
		argLine       string
		expectedStart time.Time
		expectedText  string
	}{
		{"pause now", test.MockDate("03.03.2024 17:27"), ""},
		{"pause now lunch with friends", test.MockDate("03.03.2024 17:27"), "lunch with friends"},
		{"pause 12:15", test.MockDate("03.03.2024 12:15"), ""},
		{"pause 12:15 lunch", test.MockDate("03.03.2024 12:15"), "lunch"},
		{"pause -s now", test.MockDate("03.03.2024 17:27"), ""},
		{"pause -s now nice dinner", test.MockDate("03.03.2024 17:27"), "nice dinner"},
		{"pause -s 12:15", test.MockDate("03.03.2024 12:15"), ""},
		{"pause -s 12:15 break", test.MockDate("03.03.2024 12:15"), "break"},
	}

	for _, tc := range testCases {
		t.Run(tc.argLine, func(t *testing.T) {
			prepareTestDay()

			test.ExecuteCommand(t, Root, tc.argLine)

			if len(data.State.DayList.Days) != 1 {
				t.Fatalf("expected still 1 day in the day list, got %d", len(data.State.DayList.Days))
			}
			day := data.State.DayList.Day(now)
			if len(day.Tasks) != 2 {
				t.Fatalf("expected 2 tasks in the day, got %d", len(day.Tasks))
			}
			pauseTask := day.Tasks[1]
			if !pauseTask.IsPause {
				t.Errorf("expected the second task to be a pause task")
			}
			if pauseTask.Start != tc.expectedStart {
				t.Errorf("expected pause start time to be %v, got %v", tc.expectedStart, pauseTask.Start)
			}
			if pauseTask.Text != tc.expectedText {
				t.Errorf("expected pause task text to be %q, got %q", tc.expectedText, pauseTask.Text)
			}
			if len(pauseTask.Tags) != 0 {
				t.Errorf("expected pause task tags to be empty, got %d tags", len(pauseTask.Tags))
			}
		})
	}
}

func TestPauseUpdate(t *testing.T) {
	now := test.MockDate("03.03.2024 17:27")
	test.MockNowAt(t, now)

	d := data.NewDay(test.MockDate("03.03.2024 00:00"))
	d.AddTasks(
		data.NewTask(test.MockDate("03.03.2024 9:00"), "a task", "Haora"),
		data.NewPause(test.MockDate("03.03.2024 12:00"), "lunch"),
	)
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	test.ExecuteCommand(t, Root, "pause 12:00 breakfast")

	day := data.State.DayList.Day(now)
	if len(day.Tasks) != 2 {
		t.Fatalf("expected 2 tasks in the day, got %d", len(day.Tasks)) // should update existing pause
	}
	pauseTask := day.Tasks[1]
	expected := "breakfast"
	if pauseTask.Text != expected {
		t.Errorf("expected updated pause text to be %q, got %q", expected, pauseTask.Text)
	}
}
