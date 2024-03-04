package data

import (
	"time"
)

var State StateType

type StateType struct {

	// DayList with all days recorded so far.
	DayList DayListType

	// WorkingDate represents the global set date to apply commands on.
	WorkingDate time.Time
}

func InitState(workingDate time.Time) {
	State.WorkingDate = workingDate
}

func (s *StateType) WorkingDay() Day {
	return s.DayList.Day(s.WorkingDate)
}

func (s *StateType) SanitizedDays() []Day {
	var r = make([]Day, 0)
	for _, d := range s.DayList.Days {
		if !d.IsEmpty() { // ignore days without any task
			r = append(r, d.sanitize())
		}
	}
	return r
}
