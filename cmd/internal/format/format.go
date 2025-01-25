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

package format

import (
	"fmt"
	"math"
	"time"
)

func Duration(d time.Duration) string {
	h, m := hhmm(d)
	var hStr, mStr string
	if h > 0 {
		hStr = fmt.Sprintf("%dh", h)
	}
	mStr = fmt.Sprintf("%dm", m)
	return fmt.Sprintf("%3v %3v", hStr, mStr)
}

func hhmm(d time.Duration) (h int, m int) {
	s := int(d.Seconds())
	h = abs(s / 3600)
	s %= 3600
	m = abs(s / 60)
	return
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func DurationShort(d time.Duration) string {
	h, m := hhmm(d)
	switch {
	case h > 0 && m > 0:
		return fmt.Sprintf("%2dh %2dm", h, m)
	case h > 0:
		return fmt.Sprintf("%2dh", h)
	case m > 0:
		return fmt.Sprintf("%2dm", m)
	default:
		return ""
	}
}

func DurationDecimal(d time.Duration) string {
	v := math.Round(d.Hours()*100.0) / 100.0
	return fmt.Sprintf("%5.2fh", v)
}

func DurationDecimalRounded(d time.Duration, r time.Duration) string {
	v := d.Round(r)
	return DurationDecimal(v)
}
