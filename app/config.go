package app

import (
	"os"
)

const (
	haoraDir   = ".haora"
	dataFile   = "workbook.json"
	configFile = "config"
)

var userHomeDir = os.UserHomeDir

func LoadConfig() error {
	/*
		homeDir, err := userHomeDir()
		if err != nil {
			return err
		}
	*/
	return nil
}
