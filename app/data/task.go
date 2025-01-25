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
	"time"
)

type Task struct {
	Start   time.Time
	Text    string
	IsPause bool
	Tags    []string
}

func NewTask(s time.Time, tx string, tgs ...string) *Task {
	return &Task{
		Start:   s.Truncate(time.Minute),
		Text:    tx,
		IsPause: false,
		Tags:    tgs,
	}
}

func NewPause(s time.Time, tx string) *Task {
	return &Task{
		Start:   s.Truncate(time.Minute),
		Text:    tx,
		IsPause: true,
		Tags:    nil,
	}
}

var tasksByStart = func(t1, t2 *Task) int {
	switch {
	case t1.Start.Before(t2.Start):
		return -1
	case t1.Start.After(t2.Start):
		return 1
	default:
		return 0
	}
}
