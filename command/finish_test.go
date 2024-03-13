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

func TestFinish(t *testing.T) {
	now := test.Date("22.02.2024 16:32")
	datetime.AssumeForTestNowAt(t, now)

	prepareTestDay := func() {
		d := data.NewDay(test.Date("22.02.2024 00:00"))
		d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "a task", "Haora"))
		if !d.Finished.IsZero() {
			t.Fatal("day to test should not be finished already")
		}
		data.State.DayList = &data.DayListType{Days: []*data.Day{d}}
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
			prepareTestDay()

			test.ExecuteCommand(t, Root, tc.argLine)

			d := data.State.DayList.Day(now)
			if d.Finished != tc.expectedFinished {
				t.Errorf("expected finished time %v, but got %v", tc.expectedFinished, d.Finished)
			}
		})
	}
}
