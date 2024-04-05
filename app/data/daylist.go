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
	"time"
)

type DayList struct {
	Days []*Day
}

// Day returns the Day struct for the specified date.
//
// The returned struct is a copy of the day.
// Changes to this day won't be applied to the data model automatically.
func (dl *DayList) Day(date time.Time) *Day {
	for _, day := range dl.Days {
		if isSameDay(day.Date, date) {
			return day
		}
	}
	day := NewDay(date)
	dl.Days = append(dl.Days, day)
	return day
}

func (dl *DayList) Week(start time.Time) Week {
	var week Week
	var date = start
	for i := 0; i < 7; i++ {
		week.Days[i] = *dl.Day(date)
		date = date.Add(24 * time.Hour)
	}
	return week
}

func (dl *DayList) SanitizedDays() []*Day {
	var r = make([]*Day, 0)
	for _, d := range dl.Days {
		if !d.IsEmpty() { // ignore days without any task
			r = append(r, d.sanitize())
		}
	}
	return r
}
