package cmd

import (
	"errors"
	"fmt"
	"haora/app"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	workingDateFlag *string
	workingDate     time.Time

	// RE for parsing dates like 02.01.2006 or 02.01. or 02. ...
	re = regexp.MustCompile(`(\d+)(?:\.(\d+)(?:\.(\d+)?)?)?`)
	// weekdays for selecting the preceding weekday
	weekdays = map[string]time.Weekday{
		"mo": time.Monday,
		"tu": time.Tuesday,
		"we": time.Wednesday,
		"th": time.Thursday,
		"fr": time.Friday,
		"sa": time.Saturday,
		"su": time.Sunday,
	}
)

func ParseDateFlag() error {
	// no date flag given
	if workingDateFlag == nil || *workingDateFlag == "" {
		workingDate = app.Now()
		return nil
	}

	var dateStringErr, weekdayErr error
	dateStringErr = tryDateString()
	if dateStringErr == nil {
		return nil
	}
	weekdayErr = tryWeekdayString()
	if weekdayErr == nil {
		return nil
	}

	return errors.Join(dateStringErr, weekdayErr)
}

func tryDateString() error {
	var err error
	groups := re.FindStringSubmatch(*workingDateFlag)
	if len(groups) == 0 {
		return errors.New("no date string match")
	}

	var now = app.Now()
	var day = now.Day()
	var month = int(now.Month())
	var year = now.Year()
	if err = parse(&day, groups[1]); err != nil {
		return err
	}
	if err = parse(&month, groups[2]); err != nil {
		return err
	}
	if err = parse(&year, groups[3]); err != nil {
		return err
	}

	if day < 1 || day > daysInMonth(year, month) || month < 1 || month > 12 {
		return fmt.Errorf("unable to parse date flag %q", *workingDateFlag)
	}
	workingDate = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return nil
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

func tryWeekdayString() error {
	for s, wd := range weekdays {
		if strings.HasPrefix(strings.ToLower(*workingDateFlag), s) {
			workingDate = previous(wd)
			return nil
		}
	}
	return errors.New("no weekday string match")
}

func previous(weekday time.Weekday) time.Time {
	d := app.Now().Add(-24 * time.Hour)
	for d.Weekday() != weekday {
		d = d.Add(-24 * time.Hour)
	}
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
