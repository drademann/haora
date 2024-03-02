package format

import (
	"fmt"
	"time"
)

func Duration(d time.Duration) string {
	s := int(d.Seconds())
	h := s / 3600
	s %= 3600
	m := s / 60

	var hStr, mStr string
	if h > 0 {
		hStr = fmt.Sprintf("%dh", h)
	}
	mStr = fmt.Sprintf("%dm", m)
	return fmt.Sprintf("%3v %3v", hStr, mStr)
}
