//
// Copyright 2024-2025 The Haora Authors
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

package data

import (
	"encoding/json"
	"github.com/drademann/haora/cmd/config"
	"io"
	"os"
	"path/filepath"
	"testing"
)

var Load = loadFunc

func loadFunc() (*DayList, error) {
	if err := ensureDataDirExists(); err != nil {
		return nil, err
	}
	homeDir, err := config.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(homeDir, config.HaoraDir, config.DataFile)
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &DayList{}, nil
		}
		return nil, err
	}
	defer file.Close()
	dayList, err := read(file)
	if err != nil {
		return nil, err
	}
	return dayList, nil
}

func read(r io.Reader) (*DayList, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	dayList := DayList{}
	if err = json.Unmarshal(bytes, &dayList.Days); err != nil {
		return nil, err
	}
	return &dayList, nil
}

var Save = saveFunc

func saveFunc(dayList *DayList) error {
	if err := ensureDataDirExists(); err != nil {
		return err
	}
	homeDir, err := config.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := filepath.Join(homeDir, config.HaoraDir, config.DataFile)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	if err = write(file, dayList); err != nil {
		return err
	}
	return nil
}

func write(w io.Writer, dayList *DayList) error {
	bytes, err := json.MarshalIndent(dayList.SanitizedDays(), "", "  ")
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
	homeDir, err := config.UserHomeDir()
	if err != nil {
		return err
	}
	dirPath := filepath.Join(homeDir, config.HaoraDir)
	_, err = os.Stat(dirPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirPath, 0700)
		if errDir != nil {
			return err
		}
	}
	return nil
}

// MockLoadSave is a function that mocks the behavior of the data.Load and data.Save functions by replacing
// them with custom implementations.
// The original functions are stored and can be restored later by calling the returned function.
//
// The mocked data.Load function always returns the provided dayList and nil error.
// The mocked data.Save function always returns nil error.
//
// Should not be called from production code!
func MockLoadSave(t *testing.T, dayList *DayList) {
	t.Helper()
	originalLoad := Load
	Load = func() (*DayList, error) {
		return dayList, nil
	}
	originalSave := Save
	Save = func(dayList *DayList) error {
		return nil
	}
	t.Cleanup(func() {
		Load = originalLoad
		Save = originalSave
	})
}
