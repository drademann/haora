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
	"slices"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	task := NewTask(time.Now(), "a task")

	if task.Start.Second() != 0 {
		t.Errorf("expected task start time seconds to be 0, but got %d", task.Start.Second())
	}
	if task.Start.Nanosecond() != 0 {
		t.Errorf("expected task start time nanoseconds to be 0, but got %d", task.Start.Nanosecond())
	}
}

func TestTasksByStart(t *testing.T) {
	tasks := []*Task{
		NewTask(test.Time("10:00"), "Y"),
		NewTask(test.Time("9:00"), "Z"),
		NewTask(test.Time("12:00"), "X"),
	}

	slices.SortFunc(tasks, tasksByStart)

	if tasks[0].Text != "Z" || tasks[1].Text != "Y" || tasks[2].Text != "X" {
		t.Errorf("expected sorted ordering to be task Z, Y, X, but got %s, %s, %s", tasks[0].Text, tasks[1].Text, tasks[2].Text)
	}
}
