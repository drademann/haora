package parsing

import (
	"errors"
	"github.com/drademann/haora/app/datetime"
	"time"
)

func Time(flag string, args []string) (time.Time, []string, error) {
	if flag != "" {
		if flag == "now" {
			return datetime.Now(), args, nil
		}
		t, err := time.Parse("15:04", flag)
		if err != nil {
			return time.Time{}, args, err
		}
		return t, args, nil
	}
	if len(args) > 0 {
		if args[0] == "now" {
			return datetime.Now(), args[1:], nil
		}
		t, err := time.Parse("15:04", args[0])
		if err != nil {
			return time.Time{}, args, err
		}
		return t, args[1:], nil
	}
	return time.Time{}, args, errors.New("no task starting time found")
}
