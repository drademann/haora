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

package data

import (
	"github.com/drademann/fugo/test"
	"testing"
	"time"
)

func TestDay(t *testing.T) {
	t.Run("should return day if it exists", func(t *testing.T) {
		dayList, existingDay := dayListForTesting()

		date := test.Time("13:48")
		d := dayList.Day(date)

		if !IsSameDay(d.Date, existingDay.Date) {
			t.Errorf("got unexpected task: %+v", d)
		}
		if len(dayList.Days) != 1 {
			t.Errorf("number of days shouldn't have changed, but is now %d", len(dayList.Days))
		}
	})
	t.Run("should create a new day if it doesn't exist", func(t *testing.T) {
		dayList, _ := dayListForTesting()

		date := test.Date("30.06.2024 10:00")
		d := dayList.Day(date)

		if !IsSameDay(d.Date, date) {
			t.Errorf("got unexpected task: %+v", d)
		}
		if len(dayList.Days) != 2 {
			t.Errorf("number of days should have increased to 2, but is now %d", len(dayList.Days))
		}
		if _ = dayList.Day(date); len(dayList.Days) != 2 {
			t.Errorf("number of days shouldn't have increased calling the same date a second time")
		}
	})
}

func dayListForTesting() (*DayList, Day) {
	existingDay := Day{
		Date: test.Time("9:00"),
		Tasks: []*Task{
			{Start: test.Time("9:00"),
				Text:    "a task",
				IsPause: false,
				Tags:    []string{}},
		},
		Finished: time.Time{}}
	return &DayList{
		Days: []*Day{&existingDay},
	}, existingDay
}
