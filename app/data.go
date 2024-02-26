package app

import "time"

// Data represents the so far added Days.
//
// The app will load the list before executing a command,
// and will save the (changed) list after the command finishes without an error.
var Data DayList

// WorkingDate represents the global set date to apply commands on.
var WorkingDate time.Time
