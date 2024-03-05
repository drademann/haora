package command

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/test"
	"reflect"
	"testing"
	"time"
)

func TestAddCmd(t *testing.T) {
	now := datetime.MockNowAt(t, test.MockDate("26.02.2024 13:37"))

	prepareTestData := func() {
		data.State.DayList.Days = nil
	}

	testCases := []struct {
		argLine       string
		expectedStart time.Time
		expectedText  string
		expectedTags  []string
	}{
		{
			"--date 26.2.2024 add --start 12:15 --tags haora simple task",
			test.MockDate("26.02.2024 12:15"),
			"simple task",
			[]string{"haora"},
		},
		{
			"--date 26.2.2024 add -s now --tags haora nowadays",
			now,
			"nowadays",
			[]string{"haora"},
		},
		{
			"--date 26.2.2024 add -s now haora programming",
			now,
			"programming",
			[]string{"haora"},
		},
		{
			"--date 26.2.2024 add -s now --no-tags haora programming",
			now,
			"haora programming",
			nil,
		},
	}

	for _, tc := range testCases {
		prepareTestData()

		test.ExecuteCommand(t, Root, tc.argLine)

		d := data.State.DayList.Day(tc.expectedStart)
		if len(d.Tasks) != 1 {
			t.Fatalf("expected 1 task, got %d", len(d.Tasks))
		}
		task := d.Tasks[0]
		if task.Start != tc.expectedStart {
			t.Errorf("expected start time %v, got %v", tc.expectedStart, task.Start)
		}
		if task.Text != tc.expectedText {
			t.Errorf("expected text %q, got %q", tc.expectedText, task.Text)
		}
		if !reflect.DeepEqual(task.Tags, tc.expectedTags) {
			t.Errorf("expected tags %v, got %v", tc.expectedTags, task.Tags)
		}
	}
}

func TestAddShouldUpdateExistingTaskAtSameTime(t *testing.T) {
	data.State.DayList = data.DayListType{
		Days: []data.Day{
			{
				Date: test.MockDate("26.02.2024 00:00"),
				Tasks: []data.Task{
					{
						Start: test.MockDate("26.02.2024 12:15"),
						Text:  "existing task",
						Tags:  []string{"beer"},
					},
				},
			},
		},
	}

	test.ExecuteCommand(t, Root, "--date 26.02.2024 add --start 12:15 --tags haora simple task")

	d := data.State.DayList.Day(test.MockDate("26.02.2024 00:00"))
	if len(d.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(d.Tasks))
	}
	task := d.Tasks[0]
	if task.Text != "simple task" {
		t.Errorf("expected updated task's text to be %q, but got %q", "simple task", task.Text)
	}
	if !reflect.DeepEqual(task.Tags, []string{"haora"}) {
		t.Errorf("expected updated task's tags to be %v, but got %v", []string{"haora"}, task.Tags)
	}
}
