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
	"github.com/drademann/haora/cmd/config"
	"testing"
	"time"
)

func TestWeek_TotalWeekDurationWithoutVacation(t *testing.T) {
	config.SetDurationPerWeek(t, 40.0*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.ApplyConfigOptions(t)

	w := Week{
		Days: [7]Day{
			{test.Date("10.02.2025 00:00"), nil, time.Time{}, false},
			{test.Date("11.02.2025 00:00"), nil, time.Time{}, false},
			{test.Date("12.02.2025 00:00"), nil, time.Time{}, false},
			{test.Date("13.02.2025 00:00"), nil, time.Time{}, false},
			{test.Date("14.02.2025 00:00"), nil, time.Time{}, true},
		},
	}

	twd, ok := w.TotalWeekDurationWithoutVacation()

	if !ok {
		t.Fatal("expected a successful calculation, are config settings correct?")
	}
	if twd != 32*time.Hour {
		t.Errorf("expected total week duration to be 32 hours, but got %s", twd)
	}
}
