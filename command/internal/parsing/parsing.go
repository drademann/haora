package parsing

import (
	"errors"
	"fmt"
	"github.com/drademann/haora/app/datetime"
	"regexp"
	"strconv"
	"time"
)

func Time(flag string, args []string) (time.Time, []string, error) {
	if flag != "" {
		t, err := parseTime(flag)
		if err != nil {
			return t, args, err
		}
		return t, args, nil
	}
	if len(args) > 0 {
		t, err := parseTime(args[0])
		if err != nil {
			return t, args, err
		}
		return t, args[1:], nil
	}
	return time.Time{}, args, errors.New("no time found")
}

var timeRE = regexp.MustCompile(`(\d?\d):?(\d\d)`)

func parseTime(timeStr string) (time.Time, error) {
	if timeStr == "now" {
		return datetime.Now(), nil
	}
	groups := timeRE.FindStringSubmatch(timeStr)
	if len(groups) == 0 {
		return time.Time{}, errors.New("invalid time format")
	}
	hour, err := strconv.Atoi(groups[1])
	if err != nil {
		return time.Time{}, err
	}
	if hour > 23 {
		return time.Time{}, fmt.Errorf("invalid hour: %d", hour)
	}
	minute, err := strconv.Atoi(groups[2])
	if err != nil {
		return time.Time{}, err
	}
	if minute > 59 {
		return time.Time{}, fmt.Errorf("invalid minute: %d", minute)
	}
	t := time.Time{}
	return time.Date(t.Year(), t.Month(), t.Day(), hour, minute, 0, 0, t.Location()), nil
}
