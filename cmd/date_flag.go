package cmd

import (
	"errors"
	"fmt"
	"haora/app"
	"time"
)

var (
	workingDateFlag *string
	workingDate     time.Time
)

func ParseDateFlag() error {
	// no date flag given
	if workingDateFlag == nil || *workingDateFlag == "" {
		workingDate = app.Now()
		return nil
	}

	var p time.Time
	var err error
	p, err = time.Parse("2.1.2006", *workingDateFlag)
	if err == nil {
		workingDate = time.Date(p.Year(), p.Month(), p.Day(), 0, 0, 0, 0, time.Local)
		return nil
	}
	p, err = time.Parse("2.1.", *workingDateFlag)
	if err == nil {
		workingDate = time.Date(app.Now().Year(), p.Month(), p.Day(), 0, 0, 0, 0, time.Local)
		return nil
	}
	p, err = time.Parse("2.1", *workingDateFlag)
	if err == nil {
		workingDate = time.Date(app.Now().Year(), p.Month(), p.Day(), 0, 0, 0, 0, time.Local)
		return nil
	}
	p, err = time.Parse("2.", *workingDateFlag)
	if err == nil {
		workingDate = time.Date(app.Now().Year(), app.Now().Month(), p.Day(), 0, 0, 0, 0, time.Local)
		return nil
	}
	p, err = time.Parse("2", *workingDateFlag)
	if err == nil {
		workingDate = time.Date(app.Now().Year(), app.Now().Month(), p.Day(), 0, 0, 0, 0, time.Local)
		return nil
	}

	return errors.Join(fmt.Errorf("unable to parse date flag %q", *workingDateFlag), err)
}
