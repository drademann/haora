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

package add

import (
	"github.com/drademann/fugo/test"
	"github.com/drademann/fugo/test/assert"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/cmd"
	"github.com/drademann/haora/cmd/root"
	"testing"
)

func TestAddVacationCmd(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("26.02.2024 13:37"))
	dayList := data.DayList{
		Days: []*data.Day{
			{
				Date:       test.Date("26.02.2024 00:00"),
				Tasks:      []*data.Task{},
				IsVacation: false,
			},
		},
	}
	data.MockLoadSave(t, &dayList)

	cmd.TestExecute(t, root.Command, "add vacation")

	d := dayList.Day(test.Date("26.02.2024 00:00"))
	assert.True(t, d.IsVacation)
}
