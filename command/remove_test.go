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
	"github.com/drademann/haora/test/assert"
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("21.06.2024 16:23"))

	dayList := data.DayList{
		Days: []*data.Day{
			{
				Date: test.Date("21.06.2024 00:00"),
				Tasks: []*data.Task{
					{Start: test.Time("09:00"), Text: "A"},
					{Start: test.Time("10:00"), Text: "B"},
					{Start: test.Time("11:00"), Text: "C"},
				},
			},
		},
	}
	data.MockLoadSave(t, &dayList)

	test.ExecuteCommand(t, Root, "remove 10:00")

	if len(dayList.Days) != 1 {
		t.Errorf("expected still 1 day, got %d", len(dayList.Days))
	}
	if len(dayList.Days[0].Tasks) != 2 {
		t.Errorf("expected 2 task left, got %d", len(dayList.Days[0].Tasks))
	}

	taskTexts := make([]string, len(dayList.Days[0].Tasks))
	for i, task := range dayList.Days[0].Tasks {
		taskTexts[i] = task.Text
	}
	if !reflect.DeepEqual(taskTexts, []string{"A", "C"}) {
		t.Errorf("expected leftover tasks text's to be A and C, got %v", taskTexts)
	}
}

func TestRemoveOneTask(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("21.06.2024 16:23"))

	dayList := data.DayList{
		Days: []*data.Day{
			{
				Date: test.Date("21.06.2024 00:00"),
				Tasks: []*data.Task{
					{Start: test.Time("10:00"), Text: "B"},
				},
			},
		},
	}
	data.MockLoadSave(t, &dayList)

	test.ExecuteCommand(t, Root, "remove 10:00")

	if len(dayList.Days) != 1 {
		t.Errorf("expected still 1 day, got %d", len(dayList.Days))
	}
	if len(dayList.Days[0].Tasks) != 0 {
		t.Errorf("expected no task left, got %d", len(dayList.Days[0].Tasks))
	}
}

func TestRemoveNoTask_shouldPrintErrorMessage(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("21.06.2024 16:23"))

	dayList := data.DayList{
		Days: []*data.Day{
			{
				Date: test.Date("21.06.2024 00:00"),
				Tasks: []*data.Task{
					{Start: test.Time("10:00"), Text: "B"},
				},
			},
		},
	}
	data.MockLoadSave(t, &dayList)

	out := test.ExecuteCommand(t, Root, "remove 12:00")

	assert.Output(t, out, "error: no task found at 12:00\n")
}