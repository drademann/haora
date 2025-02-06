//
// Copyright 2024-2025 The Haora Authors
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
	"github.com/drademann/haora/cmd/config"
	"time"
)

type Week struct {
	Days [7]Day
}

func (w Week) TotalWorkDuration() time.Duration {
	total := 0 * time.Nanosecond
	for _, day := range w.Days {
		total += day.TotalWorkDuration()
	}
	return total
}

func (w Week) TotalOvertimeDuration() (time.Duration, bool) {
	durationPerWeek, exist := w.TotalWeekDurationWithoutVacation()
	if !exist {
		return 0, false
	}
	return w.TotalWorkDuration() - durationPerWeek, true
}

func (w Week) TotalWeekDurationWithoutVacation() (time.Duration, bool) {
	total, ok := config.DurationPerWeek()
	if !ok {
		return 0, false
	}
	perDay, ok := config.DurationPerDay()
	if !ok {
		return 0, false
	}
	for _, day := range w.Days {
		if day.IsVacation {
			total -= perDay
		}
	}
	return total, true
}
