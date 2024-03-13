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

var State *StateType

func init() {
	State = &StateType{
		DayList: &DayListType{},
	}
}

type StateType struct {

	// DayList with all days recorded so far.
	DayList *DayListType

	// WorkingDate represents the global set date to apply commands on.
	WorkingDate time.Time
}

func (s *StateType) WorkingDay() *Day {
	return s.DayList.Day(s.WorkingDate)
}

func (s *StateType) SanitizedDays() []*Day {
	var r = make([]*Day, 0)
	for _, d := range s.DayList.Days {
		if !d.IsEmpty() { // ignore days without any task
			r = append(r, d.sanitize())
		}
	}
	return r
}
