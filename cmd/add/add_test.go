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

package add

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/cmd/root"
	"github.com/drademann/haora/test"
	"reflect"
	"testing"
	"time"
)

func TestAddCmd(t *testing.T) {
	now := datetime.AssumeForTestNowAt(t, test.Date("26.02.2024 13:37"))

	testCases := []struct {
		argLine       string
		expectedStart time.Time
		expectedText  string
		expectedTags  []string
	}{
		{
			"--date 26.2.2024 add --start 12:15 --tags haora simple task",
			test.Date("26.02.2024 12:15"),
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
		t.Run(tc.argLine, func(t *testing.T) {
			dayList := data.DayList{}
			data.MockLoadSave(t, &dayList)

			test.ExecuteCommand(t, root.Command, tc.argLine)

			d := dayList.Day(tc.expectedStart)
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
		})
	}
}

func TestAddShouldUpdateExistingTaskAtSameTime(t *testing.T) {
	dayList := data.DayList{
		Days: []*data.Day{
			{
				Date: test.Date("26.02.2024 00:00"),
				Tasks: []*data.Task{
					{
						Start: test.Date("26.02.2024 12:15"),
						Text:  "existing task",
						Tags:  []string{"beer"},
					},
				},
			},
		},
	}
	data.MockLoadSave(t, &dayList)

	test.ExecuteCommand(t, root.Command, "--date 26.02.2024 add --start 12:15 --tags haora simple task")

	d := dayList.Day(test.Date("26.02.2024 00:00"))
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
