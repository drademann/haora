package app

import (
	"haora/test"
	"testing"
	"time"
)

func TestNewDay(t *testing.T) {
	date := test.MockDate(2024, time.February, 21, 14, 58)
	day := NewDay(date)

	if len(day.Tasks) != 0 {
		t.Errorf("expected new day to have no tasks, but found %d", len(day.Tasks))
	}
	if !day.Finished.IsZero() {
		t.Errorf("didn't expect new day to be finished from the beginning")
	}
}

func TestHasNoTasks(t *testing.T) {
	day := NewDay(time.Now())

	result := day.IsEmpty()

	if !result {
		t.Errorf("expected day to have no tasks, but it has %d", len(day.Tasks))
	}
}

func TestTaskAt(t *testing.T) {
	task1 := Task{
		Start: test.MockTime(10, 20),
		Text:  "existing text",
		Tags:  []string{"haora"},
	}
	task2 := Task{
		Start: test.MockTime(12, 30),
		Text:  "lunch",
		Tags:  nil,
	}
	day := Day{Tasks: []Task{task1, task2}}

	found, err := day.taskAt(test.MockTime(10, 20))

	if err != nil {
		t.Fatal(err)
	}
	if found.Id != task1.Id {
		t.Errorf("expected found task to be task1, but got %v", found)
	}
}

func TestSameDay(t *testing.T) {
	testCases := []struct {
		name     string
		date1    time.Time
		date2    time.Time
		expected bool
	}{
		{"dates at exact same time should return true",
			test.MockDate(2024, time.February, 21, 10, 0),
			test.MockDate(2024, time.February, 21, 10, 0),
			true},
		{"dates at same day should return true",
			test.MockDate(2024, time.February, 21, 10, 0),
			test.MockDate(2024, time.February, 21, 15, 22),
			true},
		{"dates at different days should return false",
			test.MockDate(2024, time.February, 21, 10, 0),
			test.MockDate(2024, time.February, 12, 10, 0),
			false},
		{"dates at different month should return false",
			test.MockDate(2024, time.February, 21, 10, 0),
			test.MockDate(2024, time.December, 21, 10, 0),
			false},
		{"dates at different years should return false",
			test.MockDate(2024, time.February, 21, 10, 0),
			test.MockDate(2025, time.February, 21, 10, 0),
			false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isSameDay(tc.date1, tc.date2)
			if result != tc.expected {
				t.Errorf("expected %t, but got %t", tc.expected, result)
			}
		})
	}
}
