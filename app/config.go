package app

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const (
	HaoraDir = ".haora"

	filename = "config.json"
)

var UserHomeDir = os.UserHomeDir

type Type struct {
	Times struct {
		PeriodPerWeek int `json:"PeriodPerWeek"`
		DaysPerWeek   int `json:"DaysPerWeek"`
	} `json:"times"`
}

var Config Type

func init() {
	// set defaults
	Config.Times.PeriodPerWeek = 40
	Config.Times.DaysPerWeek = 5
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
