package cmd

import (
	"testing"
)

func TestFinishCmd_plainTime(t *testing.T) {
	mockNowAt(t, mockDate("22.02.2024 16:32"))

	d := newDay(mockDate("22.02.2024 00:00"))
	d.addTasks(newTask(mockDate("22.02.2024 9:00"), "a task", "Haora"))
	if !d.finished.IsZero() {
		t.Fatal("day to test should not be finished already")
	}
	ctx.data = dayList{days: []day{d}}

	out := executeCommand(t, "finish 18:00")

	d = ctx.data.day(mockDate("22.02.2024 00:00"))
	expected := mockDate("22.02.2024 18:00")
	if d.finished != expected {
		t.Errorf("expected day to be finished at %v, but got %v", expected, d.finished)
	}
	assertOutput(t, out,
		`
		Today finished at 18:00
		`)
}

func TestFinishCmd_flagTime(t *testing.T) {
	mockNowAt(t, mockDate("22.02.2024 16:32"))

	d := newDay(mockDate("22.02.2024 00:00"))
	d.addTasks(newTask(mockDate("22.02.2024 9:00"), "a task", "Haora"))
	if !d.finished.IsZero() {
		t.Fatal("day to test should not be finished already")
	}
	ctx.data = dayList{days: []day{d}}

	out := executeCommand(t, "finish -e 18:00")

	d = ctx.data.day(mockDate("22.02.2024 00:00"))
	expected := mockDate("22.02.2024 18:00")
	if d.finished != expected {
		t.Errorf("expected day to be finished at %v, but got %v", expected, d.finished)
	}
	assertOutput(t, out,
		`
		Today finished at 18:00
		`)
}

func TestFinishCmd_shouldAcceptNowAsTime(t *testing.T) {
	mockNowAt(t, mockDate("22.02.2024 16:32"))

	d := newDay(mockDate("22.02.2024 00:00"))
	d.addTasks(newTask(mockDate("22.02.2024 9:00"), "a task", "Haora"))
	if !d.finished.IsZero() {
		t.Fatal("day to test should not be finished already")
	}
	ctx.data = dayList{days: []day{d}}

	out := executeCommand(t, "finish now")

	d = ctx.data.day(mockDate("22.02.2024 00:00"))
	expected := mockDate("22.02.2024 16:32")
	if d.finished != expected {
		t.Errorf("expected day to be finished at %v, but got %v", expected, d.finished)
	}
	assertOutput(t, out,
		`
		Today finished at 16:32
		`)
}

func TestFinishCmd_shouldAcceptNowAsFlagTime(t *testing.T) {
	mockNowAt(t, mockDate("22.02.2024 16:32"))

	d := newDay(mockDate("22.02.2024 00:00"))
	d.addTasks(newTask(mockDate("22.02.2024 9:00"), "a task", "Haora"))
	if !d.finished.IsZero() {
		t.Fatal("day to test should not be finished already")
	}
	ctx.data = dayList{days: []day{d}}

	out := executeCommand(t, "finish -e now")

	d = ctx.data.day(mockDate("22.02.2024 00:00"))
	expected := mockDate("22.02.2024 16:32")
	if d.finished != expected {
		t.Errorf("expected day to be finished at %v, but got %v", expected, d.finished)
	}
	assertOutput(t, out,
		`
		Today finished at 16:32
		`)
}
