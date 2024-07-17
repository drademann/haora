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
	"time"
)

type Task struct {
	ID      uint      `gorm:"primary_key;"`
	DayID   uint      `gorm:"not null;"`
	Start   time.Time `gorm:"not null;"`
	Text    string    `gorm:"not null;"`
	IsPause bool      `gorm:"not null;default false;"`
	Tags    []string  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewTask(start time.Time, text string, tags ...string) Task {
	return Task{
		Start:   start.Truncate(time.Minute),
		Text:    text,
		IsPause: false,
		Tags:    tags,
	}
}

func NewPause(start time.Time, text string) Task {
	return Task{
		Start:   start.Truncate(time.Minute),
		Text:    text,
		IsPause: true,
		Tags:    nil,
	}
}

var tasksByStart = func(t1, t2 Task) int {
	switch {
	case t1.Start.Before(t2.Start):
		return -1
	case t1.Start.After(t2.Start):
		return 1
	default:
		return 0
	}
}
