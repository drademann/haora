package format

import (
	"fmt"
	"math"
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

func DurationDecimal(d time.Duration) string {
	v := math.Round(d.Hours()*100.0) / 100.0
	return fmt.Sprintf("%5.2fh", v)
}

func DurationDecimalRounded(d time.Duration, r time.Duration) string {
	v := d.Round(r)
	return DurationDecimal(v)
}
