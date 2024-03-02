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
