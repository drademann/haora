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
	"errors"
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestTaskSuccPred(t *testing.T) {
	testDayDate := test.Time("0:00")
	task1 := NewTask(test.Time("9:00"), "task 1")
	task2 := NewTask(test.Time("10:00"), "task 2")
	task3 := NewTask(test.Time("12:00"), "task 3")
	d := Day{Date: testDayDate,
		Tasks:    []*Task{task1, task2, task3},
		Finished: time.Time{},
	}

	t.Run("find successor", func(t *testing.T) {
		s, err := d.Succ(*task2)

		if err != nil {
			t.Fatal(err)
		}
		if s.Start != task3.Start {
			t.Errorf("expected successor to be %q, but got %q", task3.Text, s.Text)
		}
	})
	t.Run("find no successor", func(t *testing.T) {
		_, err := d.Succ(*task3)

		if err == nil {
			t.Errorf("expected error, but got nil")
		}
		if !errors.Is(err, NoTaskSucc) {
			t.Errorf("expected error %q, but got %q", NoTaskSucc, err)
		}
	})
	t.Run("find predecessor", func(t *testing.T) {
		p, err := d.Pred(*task2)

		if err != nil {
			t.Fatal(err)
		}
		if p.Start != task1.Start {
			t.Errorf("expected predecessor to be %q, but got %q", task1.Text, p.Text)
		}
	})
	t.Run("find no predecessor", func(t *testing.T) {
		_, err := d.Pred(*task1)

		if err == nil {
			t.Errorf("expected error, but got nil")
		}
		if !errors.Is(err, NoTaskPred) {
			t.Errorf("expected error %q, but got %q", NoTaskPred, err)
		}
	})
}
