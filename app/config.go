package app

import (
	"errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	haoraDir   = ".haora"
	dataFile   = "workbook.json"
	configFile = "config"
)

var userHomeDir = os.UserHomeDir

func LoadConfig() error {
	viper.SetConfigName(configFile)
	viper.SetConfigType("json")
	homeDir, err := userHomeDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(homeDir, haoraDir)
	viper.AddConfigPath(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}
	return nil
}
