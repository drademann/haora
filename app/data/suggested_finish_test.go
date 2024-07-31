//
// Copyright 2024-2024 The Haora Authors
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
	"github.com/drademann/haora/cmd/config"
	"testing"
	"time"
)

func TestSuggestedFinishTime(t *testing.T) {
	config.SetDurationPerWeek(t, 35*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.ApplyConfigOptions(t)

	d := Day{
		Tasks: []*Task{
			{Start: test.Time("10:20")},
		},
	}

	suggestion, ok := d.SuggestedFinish()

	if !ok {
		t.Fatal("expected a successful time calculation")
	}
	if suggestion.IsZero() {
		t.Fatal("expected a valid suggestion time")
	}
	if suggestion != test.Time("17:20") {
		t.Errorf("expected the suggested finish time to be 17:20 but got %v", suggestion)
	}
}

func TestSuggestedFinishTime_NoTasks(t *testing.T) {
	d := Day{}

	_, ok := d.SuggestedFinish()

	if ok {
		t.Errorf("expected a suggestion to be NOT ok because there is no task to begin with")
	}
}

func TestSuggestedFinishTime_WithPause(t *testing.T) {
	config.SetDurationPerWeek(t, 35*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.ApplyConfigOptions(t)

	d := Day{
		Tasks: []*Task{
			{Start: test.Time("10:20")},
			{Start: test.Time("12:00"), IsPause: true},
			{Start: test.Time("12:45")},
		},
	}

	suggestion, _ := d.SuggestedFinish()

	if suggestion != test.Time("18:05") {
		t.Errorf("expected the suggested finish time to be 18:05 but got %v", suggestion)
	}
}

func TestSuggestedFinishTime_WithSetFinish(t *testing.T) {
	d := Day{
		Tasks: []*Task{
			{Start: test.Time("10:20")},
			{Start: test.Time("12:00"), IsPause: true},
			{Start: test.Time("12:45")},
		},
		Finished: test.Time("17:00"),
	}

	_, ok := d.SuggestedFinish()

	if ok {
		t.Errorf("expected a suggestion to be NOT ok because the day is already finished")
	}
}
