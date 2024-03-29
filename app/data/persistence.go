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

package data

import (
	"encoding/json"
	"github.com/drademann/haora/app/config"
	"io"
	"os"
	"path/filepath"
)

func Load() error {
	if err := ensureDataDirExists(); err != nil {
		return err
	}
	homeDir, err := config.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := filepath.Join(homeDir, config.HaoraDir, config.DataFile)
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			State.DayList = &DayListType{}
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
	if err = json.Unmarshal(bytes, &State.DayList.Days); err != nil {
		return err
	}
	return nil
}

func Save() error {
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
	if err = write(file); err != nil {
		return err
	}
	return nil
}

func write(w io.Writer) error {
	bytes, err := json.MarshalIndent(State.SanitizedDays(), "", "  ")
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
