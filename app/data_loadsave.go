package app

import (
	"encoding/json"
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
			Data = DayList{}
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
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &Data.Days); err != nil {
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
	bytes, err := json.Marshal(Data.Days)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	if err != nil {
		return err
	}
	return nil
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
