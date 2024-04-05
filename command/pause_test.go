//
// Copyright 2024-2024 The Haora Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package command

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestPause(t *testing.T) {
	now := test.Date("03.03.2024 17:27")
	datetime.AssumeForTestNowAt(t, now)

	prepareTestDay := func() *data.DayList {
		d := data.NewDay(test.Date("03.03.2024 00:00"))
		d.AddTask(data.NewTask(test.Date("03.03.2024 9:00"), "a task", "Haora"))
		return &data.DayList{Days: []*data.Day{d}}
	}

	testCases := []struct {
		argLine       string
		expectedStart time.Time
		expectedText  string
	}{
		{"pause now", test.Date("03.03.2024 17:27"), ""},
		{"pause now lunch with friends", test.Date("03.03.2024 17:27"), "lunch with friends"},
		{"pause 12:15", test.Date("03.03.2024 12:15"), ""},
		{"pause 12:15 lunch", test.Date("03.03.2024 12:15"), "lunch"},
		{"pause -s now", test.Date("03.03.2024 17:27"), ""},
		{"pause -s now nice dinner", test.Date("03.03.2024 17:27"), "nice dinner"},
		{"pause -s 12:15", test.Date("03.03.2024 12:15"), ""},
		{"pause -s 12:15 break", test.Date("03.03.2024 12:15"), "break"},
	}

	for _, tc := range testCases {
		t.Run(tc.argLine, func(t *testing.T) {
			dayList := prepareTestDay()
			defer data.MockLoadSave(dayList)()

			test.ExecuteCommand(t, Root, tc.argLine)

			if len(dayList.Days) != 1 {
				t.Fatalf("expected still 1 day in the day list, got %d", len(dayList.Days))
			}
			day := dayList.Day(now)
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
	now := test.Date("03.03.2024 17:27")
	datetime.AssumeForTestNowAt(t, now)

	d := data.NewDay(test.Date("03.03.2024 00:00"))
	d.AddTask(data.NewTask(test.Date("03.03.2024 9:00"), "a task", "Haora"))
	d.AddTask(data.NewPause(test.Date("03.03.2024 12:00"), "lunch"))
	dayList := &data.DayList{Days: []*data.Day{d}}
	defer data.MockLoadSave(dayList)()

	test.ExecuteCommand(t, Root, "pause 12:00 breakfast")

	day := dayList.Day(now)
	if len(day.Tasks) != 2 {
		t.Fatalf("expected 2 tasks in the day, got %d", len(day.Tasks)) // should update existing pause
	}
	pauseTask := day.Tasks[1]
	expected := "breakfast"
	if pauseTask.Text != expected {
		t.Errorf("expected updated pause text to be %q, got %q", expected, pauseTask.Text)
	}
}
