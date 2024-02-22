package app

type dayList []Day

// DayList represents the so far added Days.
//
// The app will load the list before executing a command,
// and will save the (changed) list after the command finishes without an error.
var DayList dayList
