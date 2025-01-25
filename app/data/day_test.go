//
// Copyright 2024-2025 The Haora Authors
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
	"github.com/drademann/fugo/test"
	"reflect"
	"testing"
	"time"
)

func TestNewDay(t *testing.T) {
	date := test.Date("21.02.2024 14:58")
	d := NewDay(date)

	if len(d.Tasks) != 0 {
		t.Errorf("expected new day to have no tasks, but found %d", len(d.Tasks))
	}
	if !d.Finished.IsZero() {
		t.Errorf("didn't expect new day to be finished from the beginning")
	}
}

func TestHasNoTasks(t *testing.T) {
	d := NewDay(time.Now())

	result := d.IsEmpty()

	if !result {
		t.Errorf("expected day to have no tasks, but it has %d", len(d.Tasks))
	}
}

func TestIsToday(t *testing.T) {
	today := time.Now()
	d := NewDay(today)

	result := d.IsToday()

	if !result {
		t.Errorf("expected day to be today, but it is not")
	}
}

func TestTaskAt(t *testing.T) {
	task1 := Task{
		Start: test.Time("10:20"),
		Text:  "existing text",
		Tags:  []string{"haora"},
	}
	task2 := Task{
		Start: test.Time("12:30"),
		Text:  "lunch",
		Tags:  nil,
	}
	d := Day{Tasks: []*Task{&task1, &task2}}

	found, err := d.taskAt(test.Time("10:20"))

	if err != nil {
		t.Fatal(err)
	}
	if found.Start != task1.Start {
		t.Errorf("expected found task to be task1, but got %v", found)
	}
}

func TestTags(t *testing.T) {
	d := Day{
		Tasks: []*Task{
			{Tags: []string{"T1"}},
			{Tags: []string{"T2", "T4"}},
			{Tags: []string{"T3", "T4"}},
		},
	}

	tags := d.Tags()

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
			test.Date("21.02.2024 10:00"),
			test.Date("21.02.2024 10:00"),
			true},
		{"dates at same day should return true",
			test.Date("21.02.2024 10:00"),
			test.Date("21.02.2024 15:22"),
			true},
		{"dates at different days should return false",
			test.Date("21.02.2024 10:00"),
			test.Date("12.11.2024 10:00"),
			false},
		{"dates at different month should return false",
			test.Date("21.02.2024 10:00"),
			test.Date("21.03.2024 10:00"),
			false},
		{"dates at different years should return false",
			test.Date("21.02.2023 10:00"),
			test.Date("21.02.2024 10:00"),
			false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsSameDay(tc.date1, tc.date2)
			if result != tc.expected {
				t.Errorf("expected %t, but got %t", tc.expected, result)
			}
		})
	}
}
