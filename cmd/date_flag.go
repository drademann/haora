package cmd

import (
	"fmt"
	"haora/app"
	"regexp"
	"strconv"
	"time"
)

var (
	workingDateFlag *string
	workingDate     time.Time

	re = regexp.MustCompile(`(\d+)(?:\.(\d+)(?:\.(\d+)?)?)?`)
)

func ParseDateFlag() error {
	// no date flag given
	if workingDateFlag == nil || *workingDateFlag == "" {
		workingDate = app.Now()
		return nil
	}

	var err error
	groups := re.FindStringSubmatch(*workingDateFlag)

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
