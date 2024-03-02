package cmd

import (
	"reflect"
	"testing"
	"time"
)

func TestNewDay(t *testing.T) {
	date := mockDate("21.02.2024 14:58")
	d := newDay(date)

	if len(d.tasks) != 0 {
		t.Errorf("expected new day to have no tasks, but found %d", len(d.tasks))
	}
	if !d.finished.IsZero() {
		t.Errorf("didn't expect new day to be finished from the beginning")
	}
}

func TestHasNoTasks(t *testing.T) {
	d := newDay(time.Now())

	result := d.isEmpty()

	if !result {
		t.Errorf("expected day to have no tasks, but it has %d", len(d.tasks))
	}
}

func TestIsToday(t *testing.T) {
	today := time.Now()
	d := newDay(today)

	result := d.isToday()

	if !result {
		t.Errorf("expected day to be today, but it is not")
	}
}

func TestTaskAt(t *testing.T) {
	task1 := task{
		start: mockTime("10:20"),
		text:  "existing text",
		tags:  []string{"haora"},
	}
	task2 := task{
		start: mockTime("12:30"),
		text:  "lunch",
		tags:  nil,
	}
	d := day{tasks: []task{task1, task2}}

	found, err := d.taskAt(mockTime("10:20"))

	if err != nil {
		t.Fatal(err)
	}
	if found.id != task1.id {
		t.Errorf("expected found task to be task1, but got %v", found)
	}
}

func TestTags(t *testing.T) {
	d := day{
		tasks: []task{
			{tags: []string{"T1"}},
			{tags: []string{"T2", "T4"}},
			{tags: []string{"T3", "T4"}},
		},
	}

	tags := d.tags()

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
			mockDate("21.02.2024 10:00"),
			mockDate("21.02.2024 10:00"),
			true},
		{"dates at same day should return true",
			mockDate("21.02.2024 10:00"),
			mockDate("21.02.2024 15:22"),
			true},
		{"dates at different days should return false",
			mockDate("21.02.2024 10:00"),
			mockDate("12.11.2024 10:00"),
			false},
		{"dates at different month should return false",
			mockDate("21.02.2024 10:00"),
			mockDate("21.03.2024 10:00"),
			false},
		{"dates at different years should return false",
			mockDate("21.02.2023 10:00"),
			mockDate("21.02.2024 10:00"),
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
