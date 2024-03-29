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

package command

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/test"
	"github.com/drademann/haora/test/assert"
	"testing"
	"time"
)

func TestListTagsCmd_givenNoTasks(t *testing.T) {
	datetime.AssumeForTestNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))

	t.Run("no days and thus no tasks for today", func(t *testing.T) {
		data.State.DayList.Days = nil

		out := test.ExecuteCommand(t, Root, "-d 22.02.2024 list --tags")

		assert.Output(t, out,
			`
			Tag summary for today, 22.02.2024 (Thu)

			no tags found
			`)
	})
	t.Run("no tasks for other day than today", func(t *testing.T) {
		data.State.DayList.Days = nil

		out := test.ExecuteCommand(t, Root, "-d 20.02.2024 list --tags")

		assert.Output(t, out,
			`
			Tag summary for 20.02.2024 (Tue)
			
			no tags found
			`)
	})
}

func TestListTagsCmd(t *testing.T) {
	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "a task", "haora"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 12:00"), "a task", "learning"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 15:00"), "a task", "go", "learning"))
	data.State.DayList = &data.DayListType{Days: []*data.Day{&d}}

	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))

	out := test.ExecuteCommand(t, Root, "list --tags")

	assert.Output(t, out,
		`
		Tag summary for today, 22.02.2024 (Thu)

		 1h 32m   1.53h   1.50h   #go
		 3h  0m   3.00h   3.00h   #haora
		 4h 32m   4.53h   4.50h   #learning
		`)
}
