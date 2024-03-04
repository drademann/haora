package app

import (
	"encoding/json"
	"github.com/drademann/haora/app/data"
	"io"
	"os"
	"path/filepath"
)

const (
	dataDir  = ".haora"
	dataFile = "workbook"
)

var userHomeDir = os.UserHomeDir

func Load() error {
	if err := ensureDataDirExists(); err != nil {
		return err
	}
	homeDir, err := userHomeDir()
	if err != nil {
		return err
	}
	filePath := filepath.Join(homeDir, dataDir, dataFile)
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			data.State.DayList = data.DayListType{}
			return nil
		}
		return err
	}
	defer file.Close()
	if err = read(file); err != nil {
		return err
	}
	return nil
}

func read(r io.Reader) error {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bytes, &data.State.DayList.Days); err != nil {
		return err
	}
	return nil
}

func Save() error {
	if err := ensureDataDirExists(); err != nil {
		return err
	}
	homeDir, err := userHomeDir()
	if err != nil {
		return err
	}
	filePath := filepath.Join(homeDir, dataDir, dataFile)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	if err = write(file); err != nil {
		return err
	}
	return nil
}

func write(w io.Writer) error {
	bytes, err := json.MarshalIndent(nonEmptyDays(data.State.DayList.Days), "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func nonEmptyDays(days []data.Day) []data.Day {
	var filtered = make([]data.Day, 0)
	for _, day := range days {
		if !day.IsEmpty() {
			filtered = append(filtered, day)
		}
	}
	return filtered
}

func ensureDataDirExists() error {
	homeDir, err := userHomeDir()
	if err != nil {
		return err
	}
	dirPath := filepath.Join(homeDir, dataDir)
	_, err = os.Stat(dirPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirPath, 0700)
		if errDir != nil {
			return err
		}
	}
	return nil
}
