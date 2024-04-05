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
	"github.com/drademann/haora/app/config"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/test"
	"github.com/drademann/haora/test/assert"
	"testing"
	"time"
)

func TestListCmd_givenNoTasks(t *testing.T) {
	datetime.AssumeForTestNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))

	t.Run("no days and thus no tasks for today", func(t *testing.T) {
		defer data.MockLoadSave(&data.DayList{})()

		out := test.ExecuteCommand(t, Root, "-d 22.02.2024 list")

		assert.Output(t, out,
			`
			Tasks for today, 22.02.2024 (Thu)

			no tasks recorded
			`)
	})
	t.Run("no tasks for other day than today", func(t *testing.T) {
		defer data.MockLoadSave(&data.DayList{})()

		out := test.ExecuteCommand(t, Root, "-d 20.02.2024 list")

		assert.Output(t, out,
			`
			Tasks for 20.02.2024 (Tue)
			
			no tasks recorded
			`)
	})
}

func TestListCmd_oneOpenTaskForToday(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))
	config.SetDurationPerWeek(t, 35*time.Hour)
	config.SetDaysPerWeek(t, 5)

	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "a task", "Haora"))
	defer data.MockLoadSave(&data.DayList{Days: []*data.Day{&d}})()

	out := test.ExecuteCommand(t, Root, "list")

	assert.Output(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 -  now     7h 32m   a task #Haora
		
		         total   7h 32m
		        paused       0m
		        worked   7h 32m   (+ 32m)
		`)
}

func TestListCmd_multipleTasksLastOpen(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)

	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "some programming", "Haora"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 10:00"), "fixing bugs"))
	defer data.MockLoadSave(&data.DayList{Days: []*data.Day{&d}})()

	out := test.ExecuteCommand(t, Root, "list")

	assert.Output(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - 10:00    1h  0m   some programming #Haora
		10:00 -  now     6h 32m   fixing bugs

		         total   7h 32m
		        paused       0m
		        worked   7h 32m   (- 28m)
		`)
}

func TestListCmd_multipleTasksLastInFuture(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)

	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "some programming", "Haora"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 17:00"), "fixing bugs"))
	defer data.MockLoadSave(&data.DayList{Days: []*data.Day{&d}})()

	out := test.ExecuteCommand(t, Root, "list")

	assert.Output(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - 17:00    8h  0m   some programming #Haora
		17:00 -  now         0m   fixing bugs

		         total   8h  0m
		        paused       0m
		        worked   8h  0m
		`)
}

func TestListCmd_multipleTasksLastInFuture_andFinishedInFuture(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)

	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "some programming", "Haora"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 17:00"), "fixing bugs"))
	d.Finished = test.Date("22.02.2024 18:00")
	defer data.MockLoadSave(&data.DayList{Days: []*data.Day{&d}})()

	out := test.ExecuteCommand(t, Root, "list")

	assert.Output(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - 17:00    8h  0m   some programming #Haora
		17:00 - 18:00    1h  0m   fixing bugs

		         total   9h  0m
		        paused       0m
		        worked   9h  0m   (+  1h)
		`)
}

func TestListCmd_withPause(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)

	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "some programming", "Haora"))
	d.AddTask(data.NewPause(test.Date("22.02.2024 12:15"), ""))
	d.AddTask(data.NewTask(test.Date("22.02.2024 13:00"), "fixing bugs"))
	defer data.MockLoadSave(&data.DayList{Days: []*data.Day{&d}})()

	out := test.ExecuteCommand(t, Root, "list")

	assert.Output(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - 12:15    3h 15m   some programming #Haora
		      |             45m   
		13:00 -  now     3h 32m   fixing bugs

		         total   7h 32m
		        paused      45m
		        worked   6h 47m   (-  1h 13m)
		`)
}

func TestListCmd_withFinished(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))
	config.SetDurationPerWeek(t, 36*time.Hour+15*time.Minute)
	config.SetDaysPerWeek(t, 5)

	d := data.Day{Date: test.Date("22.02.2024 00:00")}

	d.AddTask(data.NewTask(test.Date("22.02.2024 9:00"), "some programming", "Haora"))
	d.AddTask(data.NewPause(test.Date("22.02.2024 12:15"), "lunch"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 13:00"), "fixing bugs"))
	d.Finished = test.Date("22.02.2024 17:00")
	defer data.MockLoadSave(&data.DayList{Days: []*data.Day{&d}})()

	out := test.ExecuteCommand(t, Root, "list")

	assert.Output(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - 12:15    3h 15m   some programming #Haora
		      |             45m   lunch
		13:00 - 17:00    4h  0m   fixing bugs

		         total   8h  0m
		        paused      45m
		        worked   7h 15m
		`)
}
