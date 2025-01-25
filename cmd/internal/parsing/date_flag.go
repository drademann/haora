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

package parsing

import (
	"errors"
	"fmt"
	"github.com/drademann/haora/app/datetime"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// RE for parsing dates like 02.01.2006 or 02.01. or 02.
	dateFlagRE = regexp.MustCompile(`(\d+)(?:\.(\d+)(?:\.(\d+)?)?)?`)
)

func WorkingDate(workingDateFlag string) (time.Time, error) {
	// no date flag given
	if workingDateFlag == "" {
		return datetime.Now(), nil
	}

	var dateStringErr, weekdayErr, yesterdayErr error
	workingDate, dateStringErr := tryDateString(workingDateFlag)
	if dateStringErr == nil {
		return workingDate, nil
	}
	workingDate, weekdayErr = tryWeekdayString(workingDateFlag)
	if weekdayErr == nil {
		return workingDate, nil
	}
	workingDate, yesterdayErr = tryYesterdayString(workingDateFlag)
	if yesterdayErr == nil {
		return workingDate, nil
	}

	return time.Time{}, errors.Join(dateStringErr, weekdayErr, yesterdayErr)
}

func tryDateString(workingDateFlag string) (time.Time, error) {
	var err error
	groups := dateFlagRE.FindStringSubmatch(workingDateFlag)
	if len(groups) == 0 {
		return time.Time{}, errors.New("no date string match")
	}

	var now = datetime.Now()
	var day = now.Day()
	var month = int(now.Month())
	var year = now.Year()
	if err = parse(&day, groups[1]); err != nil {
		return time.Time{}, err
	}
	if err = parse(&month, groups[2]); err != nil {
		return time.Time{}, err
	}
	switch len(groups[3]) {
	case 2:
		if err = parse(&year, "20"+groups[3]); err != nil {
			return time.Time{}, err
		}
	default:
		if err = parse(&year, groups[3]); err != nil {
			return time.Time{}, err
		}
	}
	if day < 1 || day > daysInMonth(year, month) || month < 1 || month > 12 {
		return time.Time{}, fmt.Errorf("unable to parse date flag %q", workingDateFlag)
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
}

func parse(v *int, s string) error {
	if s != "" {
		var err error
		*v, err = strconv.Atoi(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func daysInMonth(year, month int) int {
	t := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	return t.Day()
}

func tryWeekdayString(workingDateFlag string) (time.Time, error) {
	for s, wd := range weekdays {
		if strings.HasPrefix(strings.ToLower(workingDateFlag), s) {
			return previous(wd), nil
		}
	}
	return time.Time{}, errors.New("no weekday string match")
}

func previous(weekday time.Weekday) time.Time {
	d := datetime.Now().Add(-24 * time.Hour)
	for d.Weekday() != weekday {
		d = d.Add(-24 * time.Hour)
	}
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func tryYesterdayString(workingDateFlag string) (time.Time, error) {
	for _, yd := range yesterdays {
		if strings.ToLower(workingDateFlag) == yd {
			d := datetime.Now().Add(-24 * time.Hour)
			return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location()), nil
		}
	}
	return time.Time{}, errors.New("no yesterday string match")
}
