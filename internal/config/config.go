package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const DBFileName = ".gatorconfig.json"

var dbPath = ""

type Config struct {
	DBURL    string `json:"db_url"`
	UserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	dbLocation, err := getConfigPath()

	if err != nil {
		return nil, err
	}

	jsonFile, err := os.Open(dbLocation)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var cfg *Config
	json.Unmarshal(byteValue, &cfg)

	return cfg, nil
}

func (config *Config) SetUser(userName string) error {
	config.UserName = userName
	return write(config)
}

func write(config *Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return nil
	}

	err = os.WriteFile(dbPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getConfigPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	location := fmt.Sprintf("%s/%s", dir, DBFileName)
	dbPath = location
	return location, nil
}
