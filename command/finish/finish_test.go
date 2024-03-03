package finish

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command"
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestFinish(t *testing.T) {
	now := test.MockDate("22.02.2024 16:32")
	test.MockNowAt(t, now)

	prepareTestDay := func() {
		d := data.NewDay(test.MockDate("22.02.2024 00:00"))
		d.AddTasks(data.NewTask(test.MockDate("22.02.2024 9:00"), "a task", "Haora"))
		if !d.Finished.IsZero() {
			t.Fatal("day to test should not be finished already")
		}
		data.State.DayList = data.DayListType{Days: []data.Day{d}}
	}

	testCases := []struct {
		argLine          string
		expectedFinished time.Time
	}{
		{"finish now", now},
		{"finish 18:00", test.MockDate("22.02.2024 18:00")},
		{"finish -e now", now},
		{"finish -e 18:00", test.MockDate("22.02.2024 18:00")},
	}

	for _, tc := range testCases {
		t.Run(tc.argLine, func(t *testing.T) {
			prepareTestDay()

			test.ExecuteCommand(t, command.Root, tc.argLine)

			d := data.State.DayList.Day(now)
			if d.Finished != tc.expectedFinished {
				t.Errorf("expected finished time %v, but got %v", tc.expectedFinished, d.Finished)
			}
		})
	}
}
