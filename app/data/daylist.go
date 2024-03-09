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

type DayListType struct {
	Days []Day
}

// Day returns the Day struct for the specified date.
//
// The returned struct is a copy of the day.
// Changes to this day won't be applied to the data model automatically.
func (d *DayListType) Day(date time.Time) Day {
	for _, day := range d.Days {
		if isSameDay(day.Date, date) {
			return day
		}
	}
	day := NewDay(date)
	d.Days = append(d.Days, day)
	return day
}

func (d *DayListType) update(day Day) {
	for i, e := range d.Days {
		if e.Id == day.Id {
			d.Days[i] = day
		}
	}
}
