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

package finish

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/cmd/root"
	"github.com/drademann/haora/test"
	"github.com/drademann/haora/test/assert"
	"testing"
	"time"
)

func TestFinish(t *testing.T) {
	now := test.Date("22.02.2024 16:32")
	datetime.AssumeForTestNowAt(t, now)

	prepareTestDay := func() *data.DayList {
		d := data.NewDay(test.Date("22.02.2024 00:00"))
		d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "a task", "Haora"))
		if !d.Finished.IsZero() {
			t.Fatal("day to test should not be finished already")
		}
		return &data.DayList{Days: []*data.Day{d}}
	}

	testCases := []struct {
		argLine          string
		expectedFinished time.Time
	}{
		{"finish now", now},
		{"finish 18:00", test.Date("22.02.2024 18:00")},
		{"finish -e now", now},
		{"finish -e 18:00", test.Date("22.02.2024 18:00")},
	}

	for _, tc := range testCases {
		t.Run(tc.argLine, func(t *testing.T) {
			dayList := prepareTestDay()
			data.MockLoadSave(t, dayList)

			test.ExecuteCommand(t, root.Command, tc.argLine)

			d := dayList.Day(now)
			if d.Finished != tc.expectedFinished {
				t.Errorf("expected finished time %v, but got %v", tc.expectedFinished, d.Finished)
			}
		})
	}
}

func TestFinishWithoutTasks(t *testing.T) {
	now := test.Date("22.02.2024 16:32")
	datetime.AssumeForTestNowAt(t, now)

	d := data.NewDay(test.Date("22.02.2024 00:00"))
	if !d.Finished.IsZero() {
		t.Fatal("day to test should not be finished already")
	}
	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{d}})

	out := test.ExecuteCommand(t, root.Command, "finish 18:00")

	assert.Output(t, out, "error: no tasks to finish\n")
}

func TestFinishBeforeLastTask(t *testing.T) {
	now := test.Date("22.02.2024 16:32")
	datetime.AssumeForTestNowAt(t, now)

	d := data.NewDay(test.Date("22.02.2024 00:00"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "a task", "Haora"))
	if !d.Finished.IsZero() {
		t.Fatal("day to test should not be finished already")
	}
	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{d}})

	out := test.ExecuteCommand(t, root.Command, "finish 8:00")

	assert.Output(t, out, "error: can't finish before last task's start timestamp (09:00)\n")
}

func TestUnfinished(t *testing.T) {
	now := test.Date("22.02.2024 16:32")
	datetime.AssumeForTestNowAt(t, now)

	d := data.NewDay(test.Date("22.02.2024 00:00"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "a task", "Haora"))
	d.Finished = test.Date("22.02.2024 18:00")
	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{d}})

	test.ExecuteCommand(t, root.Command, "finish --remove")

	if d.IsFinished() {
		t.Errorf("expected day to be unfinished now, but it is still finished")
	}
}
