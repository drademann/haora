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
	"github.com/drademann/haora/app"
	"time"
)

type Week struct {
	Days [7]Day
}

func CollectWeek(start time.Time) Week {
	var week Week
	var date = start
	for i := 0; i < 7; i++ {
		week.Days[i] = State.DayList.Day(date)
		date = date.Add(24 * time.Hour)
	}
	return week
}

func (w Week) TotalWorkDuration() time.Duration {
	total := 0 * time.Nanosecond
	for _, day := range w.Days {
		total += day.TotalWorkDuration()
	}
	return total
}

func (w Week) TotalOvertimeDuration() (time.Duration, error) {
	durationPerWeek, err := app.DurationPerWeek()
	if err != nil {
		return 0, err
	}
	return w.TotalWorkDuration() - durationPerWeek, nil
}
