package finish

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command"
	"github.com/drademann/haora/test"
	"testing"
)

func TestFinishCmd_plainTime(t *testing.T) {
	test.MockNowAt(t, test.MockDate("22.02.2024 16:32"))

	d := data.NewDay(test.MockDate("22.02.2024 00:00"))
	d.AddTasks(data.NewTask(test.MockDate("22.02.2024 9:00"), "a task", "Haora"))
	if !d.Finished.IsZero() {
		t.Fatal("day to test should not be finished already")
	}
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	out := test.ExecuteCommand(t, command.Root, "finish 18:00")

	d = data.State.DayList.Day(test.MockDate("22.02.2024 00:00"))
	expected := test.MockDate("22.02.2024 18:00")
	if d.Finished != expected {
		t.Errorf("expected day to be finished at %v, but got %v", expected, d.Finished)
	}
	test.AssertOutput(t, out,
		`
		Today finished at 18:00
		`)
}

func TestFinishCmd_flagTime(t *testing.T) {
	test.MockNowAt(t, test.MockDate("22.02.2024 16:32"))

	d := data.NewDay(test.MockDate("22.02.2024 00:00"))
	d.AddTasks(data.NewTask(test.MockDate("22.02.2024 9:00"), "a task", "Haora"))
	if !d.Finished.IsZero() {
		t.Fatal("day to test should not be finished already")
	}
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	out := test.ExecuteCommand(t, command.Root, "finish -e 18:00")

	d = data.State.DayList.Day(test.MockDate("22.02.2024 00:00"))
	expected := test.MockDate("22.02.2024 18:00")
	if d.Finished != expected {
		t.Errorf("expected day to be finished at %v, but got %v", expected, d.Finished)
	}
	test.AssertOutput(t, out,
		`
		Today finished at 18:00
		`)
}

func TestFinishCmd_shouldAcceptNowAsTime(t *testing.T) {
	test.MockNowAt(t, test.MockDate("22.02.2024 16:32"))

	d := data.NewDay(test.MockDate("22.02.2024 00:00"))
	d.AddTasks(data.NewTask(test.MockDate("22.02.2024 9:00"), "a task", "Haora"))
	if !d.Finished.IsZero() {
		t.Fatal("day to test should not be finished already")
	}
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	out := test.ExecuteCommand(t, command.Root, "finish now")

	d = data.State.DayList.Day(test.MockDate("22.02.2024 00:00"))
	expected := test.MockDate("22.02.2024 16:32")
	if d.Finished != expected {
		t.Errorf("expected day to be finished at %v, but got %v", expected, d.Finished)
	}
	test.AssertOutput(t, out,
		`
		Today finished at 16:32
		`)
}

func TestFinishCmd_shouldAcceptNowAsFlagTime(t *testing.T) {
	test.MockNowAt(t, test.MockDate("22.02.2024 16:32"))

	d := data.NewDay(test.MockDate("22.02.2024 00:00"))
	d.AddTasks(data.NewTask(test.MockDate("22.02.2024 9:00"), "a task", "Haora"))
	if !d.Finished.IsZero() {
		t.Fatal("day to test should not be finished already")
	}
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	out := test.ExecuteCommand(t, command.Root, "finish -e now")

	d = data.State.DayList.Day(test.MockDate("22.02.2024 00:00"))
	expected := test.MockDate("22.02.2024 16:32")
	if d.Finished != expected {
		t.Errorf("expected day to be finished at %v, but got %v", expected, d.Finished)
	}
	test.AssertOutput(t, out,
		`
		Today finished at 16:32
		`)
}
