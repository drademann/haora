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

package app

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	HaoraDir = ".haora"

	filename = "config.json"
)

var UserHomeDir = os.UserHomeDir

type Type struct {
	Times struct {
		DurationPerWeek string `json:"DurationPerWeek"`
		DaysPerWeek     int    `json:"DaysPerWeek"`
	} `json:"times"`
}

var Config Type

func init() {
	// set defaults
	Config.Times.DurationPerWeek = "40h"
	Config.Times.DaysPerWeek = 5
}

func Duration(durStr string) (time.Duration, error) {
	durStr = strings.ReplaceAll(durStr, " ", "")
	return time.ParseDuration(durStr)
}

func DurationPerWeek() (time.Duration, error) {
	durationPerWeek, err := Duration(Config.Times.DurationPerWeek)
	if err != nil {
		return 0, err
	}
	return durationPerWeek, nil
}

func DurationPerDay() (time.Duration, error) {
	durationPerWeek, err := Duration(Config.Times.DurationPerWeek)
	if err != nil {
		return 0, err
	}
	nanos := durationPerWeek.Nanoseconds() / int64(Config.Times.DaysPerWeek)
	return time.Duration(nanos), nil
}

func Load() error {
	homeDir, err := UserHomeDir()
	if err != nil {
		return err
	}
	filePath := filepath.Join(homeDir, HaoraDir, filename)
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // uses defaults
		}
		return err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &Config)
	if err != nil {
		return err
	}
	return nil
}
