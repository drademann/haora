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

package test

import (
	"time"
)

func Date(dateStr string) time.Time {
	tm, err := time.Parse("02.01.2006 15:04", dateStr)
	if err != nil {
		panic(err)
	}
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), 0, 0, time.Local)
}

func Time(timeStr string) time.Time {
	return Date("21.06.2024 " + timeStr)
}
