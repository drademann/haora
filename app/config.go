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
