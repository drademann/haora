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

package datetime

import (
	"testing"
	"time"
)

// Now timestamp as variable to allow tests to override it.
//
// Seconds and nanoseconds are truncated and set to zero because all calculations in Haora are based on minutes.
var Now = NowFunc

func NowFunc() time.Time {
	return time.Now().Truncate(time.Minute)
}

// AssumeForTestNowAt allows pretending another timestamp for today (now).
// The returned function should be deferred to reestablish the original now() function.
// Should not be called from production code!
func AssumeForTestNowAt(t *testing.T, tm time.Time) time.Time {
	t.Helper()
	Now = func() time.Time {
		return tm
	}
	t.Cleanup(func() { Now = NowFunc })
	return tm
}

func Combine(d, t time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), t.Hour(), t.Minute(), 0, 0, d.Location())
}

type Direction int

const (
	Previous Direction = -1
	// Next Direction = 1
)

func FindWeekday(date time.Time, dir Direction, weekday time.Weekday) time.Time {
	if date.Weekday() == weekday {
		return date
	}
	step := time.Duration(dir * 24)
	date = date.Add(step * time.Hour)
	return FindWeekday(date, dir, weekday)
}
