package cmd

import (
	"reflect"
	"testing"
	"time"
)

func TestNewDay(t *testing.T) {
	date := mockDate(2024, time.February, 21, 14, 58)
	day := NewDay(date)

	if len(day.tasks) != 0 {
		t.Errorf("expected new day to have no tasks, but found %d", len(day.tasks))
	}
	if !day.finished.IsZero() {
		t.Errorf("didn't expect new day to be finished from the beginning")
	}
}

func TestHasNoTasks(t *testing.T) {
	day := NewDay(time.Now())

	result := day.IsEmpty()

	if !result {
		t.Errorf("expected day to have no tasks, but it has %d", len(day.tasks))
	}
}

func TestIsToday(t *testing.T) {
	today := time.Now()
	day := NewDay(today)

	result := day.IsToday()

	if !result {
		t.Errorf("expected day to be today, but it is not")
	}
}

func TestTaskAt(t *testing.T) {
	task1 := Task{
		start: mockTime(10, 20),
		text:  "existing text",
		tags:  []string{"haora"},
	}
	task2 := Task{
		start: mockTime(12, 30),
		text:  "lunch",
		tags:  nil,
	}
	day := Day{tasks: []Task{task1, task2}}

	found, err := day.taskAt(mockTime(10, 20))

	if err != nil {
		t.Fatal(err)
	}
	if found.id != task1.id {
		t.Errorf("expected found task to be task1, but got %v", found)
	}
}

func TestTags(t *testing.T) {
	day := Day{
		tasks: []Task{
			{tags: []string{"T1"}},
			{tags: []string{"T2", "T4"}},
			{tags: []string{"T3", "T4"}},
		},
	}

	tags := day.tags()

	expected := []string{"T1", "T2", "T3", "T4"}
	if !reflect.DeepEqual(expected, tags) {
		t.Errorf("expected tags to be %+v, but got %+v", expected, tags)
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
			mockDate(2024, time.February, 21, 10, 0),
			mockDate(2024, time.February, 21, 10, 0),
			true},
		{"dates at same day should return true",
			mockDate(2024, time.February, 21, 10, 0),
			mockDate(2024, time.February, 21, 15, 22),
			true},
		{"dates at different days should return false",
			mockDate(2024, time.February, 21, 10, 0),
			mockDate(2024, time.February, 12, 10, 0),
			false},
		{"dates at different month should return false",
			mockDate(2024, time.February, 21, 10, 0),
			mockDate(2024, time.December, 21, 10, 0),
			false},
		{"dates at different years should return false",
			mockDate(2024, time.February, 21, 10, 0),
			mockDate(2025, time.February, 21, 10, 0),
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
