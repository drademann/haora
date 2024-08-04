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

package edit

import (
	"fmt"
	"github.com/drademann/fugo/test"
	"github.com/drademann/fugo/test/assert"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/cmd"
	"github.com/drademann/haora/cmd/root"
	"reflect"
	"testing"
)

func TestNoTaskFound(t *testing.T) {
	dayList := data.DayList{}
	data.MockLoadSave(t, &dayList)

	testCases := []struct {
		argLine       string
		expectedError string
	}{
		{"--date 4.8.2024 edit -u 10:00", "no task found starting at 10:00"},
		{"--date 4.8.2024 edit --update=10:00", "no task found starting at 10:00"},
	}

	for _, tc := range testCases {
		out := cmd.TestExecute(t, root.Command, tc.argLine)

		assert.Output(t, out, fmt.Sprintf(
			`
			error: %s
			`,
			tc.expectedError))
	}
}

func TestEditTask(t *testing.T) {
	dayList := data.DayList{
		Days: []*data.Day{
			{
				Date: test.Date("04.08.2024 00:00"),
				Tasks: []*data.Task{
					{
						Start: test.Date("04.08.2024 09:00"),
						Text:  "existing task",
						Tags:  []string{"beer"},
					},
				},
			},
		},
	}
	data.MockLoadSave(t, &dayList)

	cmd.TestExecute(t, root.Command, "--date 04.08.2024 edit -u 900 -s 1000 --tags haora --text simple task")

	d := dayList.Day(test.Date("04.08.2024 00:00"))
	if len(d.Tasks) != 1 {
		t.Fatalf("expected still one task, got %d", len(d.Tasks))
	}
	task := d.Tasks[0]
	expectedNewStart := test.Date("04.08.2024 10:00")
	if task.Start != expectedNewStart {
		t.Errorf("expected task start to be %s, got %s", expectedNewStart, task.Start)
	}
	if task.Text != "simple task" {
		t.Errorf("expected the updated task's text to be %q but got %q", "simple task", task.Text)
	}
	if !reflect.DeepEqual(task.Tags, []string{"haora"}) {
		t.Errorf("expected the updated task's tags to be %v but got %v", []string{"haora"}, task.Tags)
	}
}

func TestEditTask_onlyTime(t *testing.T) {
	dayList := data.DayList{
		Days: []*data.Day{
			{
				Date: test.Date("04.08.2024 00:00"),
				Tasks: []*data.Task{
					{
						Start: test.Date("04.08.2024 09:00"),
						Text:  "existing task",
						Tags:  []string{"beer"},
					},
				},
			},
		},
	}
	data.MockLoadSave(t, &dayList)

	cmd.TestExecute(t, root.Command, "--date 04.08.2024 edit -u 900 -s 12:00")

	d := dayList.Day(test.Date("04.08.2024 00:00"))
	if len(d.Tasks) != 1 {
		t.Fatalf("expected still one task, got %d", len(d.Tasks))
	}
	task := d.Tasks[0]
	expectedNewStart := test.Date("04.08.2024 12:00")
	if task.Start != expectedNewStart {
		t.Errorf("expected task start to be %s, got %s", expectedNewStart, task.Start)
	}
	expectedExistingText := "existing task"
	if task.Text != expectedExistingText {
		t.Errorf("expected task's text still to be %q but got %q", expectedExistingText, task.Text)
	}
	expectedExistingTags := []string{"beer"}
	if !reflect.DeepEqual(task.Tags, expectedExistingTags) {
		t.Errorf("expected task's tags still to be %v but got %v", expectedExistingTags, task.Tags)
	}
}

func TestEditTask_onlyText(t *testing.T) {
	dayList := data.DayList{
		Days: []*data.Day{
			{
				Date: test.Date("04.08.2024 00:00"),
				Tasks: []*data.Task{
					{
						Start: test.Date("04.08.2024 09:00"),
						Text:  "existing task",
						Tags:  []string{"beer"},
					},
				},
			},
		},
	}
	data.MockLoadSave(t, &dayList)

	cmd.TestExecute(t, root.Command, "--date 04.08.2024 edit -u 900 --text hello world")

	d := dayList.Day(test.Date("04.08.2024 00:00"))
	if len(d.Tasks) != 1 {
		t.Fatalf("expected still one task, got %d", len(d.Tasks))
	}
	task := d.Tasks[0]
	expectedExistingStart := test.Date("04.08.2024 09:00")
	if task.Start != expectedExistingStart {
		t.Errorf("expected task's start still to be %s, got %s", expectedExistingStart, task.Start)
	}
	expectedNewText := "hello world"
	if task.Text != expectedNewText {
		t.Errorf("expected task's new text to be %q but got %q", expectedNewText, task.Text)
	}
	expectedExistingTags := []string{"beer"}
	if !reflect.DeepEqual(task.Tags, expectedExistingTags) {
		t.Errorf("expected task's tags still to be %v but got %v", expectedExistingTags, task.Tags)
	}
}

func TestEditTask_removeTags(t *testing.T) {
	dayList := data.DayList{
		Days: []*data.Day{
			{
				Date: test.Date("04.08.2024 00:00"),
				Tasks: []*data.Task{
					{
						Start: test.Date("04.08.2024 09:00"),
						Text:  "existing task",
						Tags:  []string{"beer", "gin"},
					},
				},
			},
		},
	}
	data.MockLoadSave(t, &dayList)

	cmd.TestExecute(t, root.Command, "--date 04.08.2024 edit -u 900 --no-tags")

	d := dayList.Day(test.Date("04.08.2024 00:00"))
	if len(d.Tasks) != 1 {
		t.Fatalf("expected still one task, got %d", len(d.Tasks))
	}
	task := d.Tasks[0]
	expectedExistingStart := test.Date("04.08.2024 09:00")
	if task.Start != expectedExistingStart {
		t.Errorf("expected task's start still to be %s, got %s", expectedExistingStart, task.Start)
	}
	expectedExistingText := "existing task"
	if task.Text != expectedExistingText {
		t.Errorf("expected task's text still to be %q but got %q", expectedExistingText, task.Text)
	}
	if len(task.Tags) != 0 {
		t.Errorf("expected task's tags to be empty but got %v", task.Tags)
	}
}
