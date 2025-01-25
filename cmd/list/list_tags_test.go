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

package list

import (
	"fmt"
	"github.com/drademann/fugo/test"
	"github.com/drademann/fugo/test/assert"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/cmd"
	"github.com/drademann/haora/cmd/root"
	"testing"
	"time"
)

func TestListTagsCmd_givenNoTasks(t *testing.T) {
	datetime.AssumeForTestNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))

	t.Run("no days and thus no tasks for today", func(t *testing.T) {
		data.MockLoadSave(t, &data.DayList{})

		out := cmd.TestExecute(t, root.Command, "-d 22.02.2024 list --tags-per-day")

		assert.Output(t, out,
			`
			Tag summary for today, 22.02.2024 (Thu)

			no tags found
			`)
	})
	t.Run("no tasks for other day than today", func(t *testing.T) {
		data.MockLoadSave(t, &data.DayList{})

		out := cmd.TestExecute(t, root.Command, "-d 20.02.2024 list --tags-per-day")

		assert.Output(t, out,
			`
			Tag summary for 20.02.2024 (Tue)
			
			no tags found
			`)
	})
}

func TestListTagsDayCmd(t *testing.T) {
	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "a task", "haora"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 12:00"), "a task", "learning"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 15:00"), "a task", "go", "learning"))
	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d}})

	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))

	flagCases := []string{"--tags-per-day", "-t"}
	for _, fc := range flagCases {
		command := fmt.Sprintf("list %s", fc)
		t.Run(command, func(t *testing.T) {
			out := cmd.TestExecute(t, root.Command, command)

			assert.Output(t, out,
				`
				Tag summary for today, 22.02.2024 (Thu)
		
				 1h 32m   1.53h   1.50h   #go
				 3h  0m   3.00h   3.00h   #haora
				 4h 32m   4.53h   4.50h   #learning
				
				 9h  4m   9.07h   9.00h
				`)
		})
	}
}

func TestListTagsMonthCmd(t *testing.T) {
	d1 := data.Day{Date: test.Date("22.02.2024 00:00")}
	d1.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "a task", "haora"))
	d1.AddTask(data.NewTask(test.Date("22.02.2024 12:00"), "a task", "learning"))
	d1.AddTask(data.NewTask(test.Date("22.02.2024 15:00"), "a task", "go", "learning"))
	d1.Finished = test.Date("22.02.2024 17:00")

	d2 := data.Day{Date: test.Date("24.02.2024 00:00")}
	d2.AddTask(data.NewTask(test.Date("24.02.2024 10:00"), "a task", "haora"))
	d2.AddTask(data.NewTask(test.Date("24.02.2024 14:00"), "a task", "learning"))
	d2.Finished = test.Date("24.02.2024 16:00")

	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d1, &d2}})

	flagCases := []string{"-d 26.02.2024 --tags-per-month"}
	for _, fc := range flagCases {
		command := fmt.Sprintf("list %s", fc)
		t.Run(command, func(t *testing.T) {
			out := cmd.TestExecute(t, root.Command, command)

			assert.Output(t, out,
				`
				Tag summary for February 2024
		
				  2h  0m    2.00h    2.00h   #go
				  7h  0m    7.00h    7.00h   #haora
				  7h  0m    7.00h    7.00h   #learning

				 16h  0m   16.00h   16.00h
				`)
		})
	}
}

func TestListTagsMonthCmd_MoreThan100(t *testing.T) {
	d1 := data.Day{Date: test.Date("22.02.2024 00:00")}
	d1.AddTask(data.NewTask(test.Date("22.02.2024 1:00"), "a task", "haora"))
	d1.Finished = test.Date("22.02.2024 23:00")

	d2 := data.Day{Date: test.Date("23.02.2024 00:00")}
	d2.AddTask(data.NewTask(test.Date("23.02.2024 01:00"), "a task", "haora"))
	d2.Finished = test.Date("23.02.2024 23:00")

	d3 := data.Day{Date: test.Date("24.02.2024 00:00")}
	d3.AddTask(data.NewTask(test.Date("24.02.2024 01:00"), "a task", "haora"))
	d3.Finished = test.Date("24.02.2024 23:00")

	d4 := data.Day{Date: test.Date("25.02.2024 00:00")}
	d4.AddTask(data.NewTask(test.Date("25.02.2024 01:00"), "a task", "haora"))
	d4.Finished = test.Date("25.02.2024 23:00")

	d5 := data.Day{Date: test.Date("26.02.2024 00:00")}
	d5.AddTask(data.NewTask(test.Date("26.02.2024 01:00"), "a task", "haora"))
	d5.Finished = test.Date("26.02.2024 23:00")

	d6 := data.Day{Date: test.Date("27.02.2024 00:00")}
	d6.AddTask(data.NewTask(test.Date("27.02.2024 12:00"), "a sandwich", "lunch"))
	d6.Finished = test.Date("27.02.2024 12:45")

	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d1, &d2, &d3, &d4, &d5, &d6}})

	flagCases := []string{"-d 22.02.2024 --tags-per-month"}
	for _, fc := range flagCases {
		command := fmt.Sprintf("list %s", fc)
		t.Run(command, func(t *testing.T) {
			out := cmd.TestExecute(t, root.Command, command)

			assert.Output(t, out,
				`
				Tag summary for February 2024
		
				110h  0m  110.00h  110.00h   #haora
				     45m    0.75h    0.75h   #lunch

				110h 45m  110.75h  110.75h
				`)
		})
	}
}
