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

package config

import (
	"errors"
	"github.com/drademann/haora/cmd/internal/parsing"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"slices"
	"testing"
	"time"
)

const (
	HaoraDir = ".haora"

	DataFile = "workbook.json"

	durationPerWeekKey = "times.durationPerWeek"
	daysPerWeekKey     = "times.daysPerWeek"
	hiddenWeekdaysKey  = "times.hiddenWeekdays"
)

var UserHomeDir = os.UserHomeDir

var (
	durationPerWeek *time.Duration
	daysPerWeek     *int
	hiddenWeekdays  *[]time.Weekday
)

func InitViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	homeDir, err := UserHomeDir()
	if err != nil {
		return
	}
	viper.AddConfigPath(filepath.Join(homeDir, HaoraDir))
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// ignore, use defaults
		} else {
			panic(err)
		}
	}
	if viper.IsSet(durationPerWeekKey) {
		duration := viper.GetDuration(durationPerWeekKey)
		durationPerWeek = &duration
	}
	if viper.IsSet(daysPerWeekKey) {
		days := viper.GetInt(daysPerWeekKey)
		daysPerWeek = &days
	}
	hiddenWeekdays = new([]time.Weekday)
	if viper.IsSet(hiddenWeekdaysKey) {
		weekdays := viper.GetStringSlice(hiddenWeekdaysKey)
		for _, weekday := range weekdays {
			weekdayTime, err := parsing.Weekday(weekday)
			if err == nil {
				*hiddenWeekdays = append(*hiddenWeekdays, weekdayTime)
			}
		}
	}
}

func DurationPerWeek() (time.Duration, bool) {
	if durationPerWeek == nil {
		return 0, false
	}
	return *durationPerWeek, true
}

func DurationPerDay() (time.Duration, bool) {
	if durationPerWeek == nil || daysPerWeek == nil {
		return 0, false
	}
	nanos := durationPerWeek.Nanoseconds() / int64(*daysPerWeek)
	return time.Duration(nanos), true
}

func IsHidden(weekday time.Weekday) bool {
	return slices.Contains(*hiddenWeekdays, weekday)
}

// for testing

func SetDurationPerWeek(t *testing.T, d time.Duration) {
	t.Helper()
	viper.Set(durationPerWeekKey, d)
	durationPerWeek = nil
}

func SetNoDurationPerWeek(t *testing.T) {
	t.Helper()
	viper.Set(durationPerWeekKey, nil)
	durationPerWeek = nil
}

func SetDaysPerWeek(t *testing.T, n int) {
	t.Helper()
	viper.Set(daysPerWeekKey, n)
	daysPerWeek = nil
}

func SetNoDaysPerWeek(t *testing.T) {
	t.Helper()
	viper.Set(daysPerWeekKey, nil)
	daysPerWeek = nil
}

func SetHiddenWeekdays(t *testing.T, hiddenWeekdaysStr string) {
	t.Helper()
	viper.Set(hiddenWeekdaysKey, hiddenWeekdaysStr)
	hiddenWeekdays = nil
}
