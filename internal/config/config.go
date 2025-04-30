package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	DbURL string `json:"db_url"`
	UserName string `json="current_user_name"`
}

func NewConfig(dbURL string) (*Config, error) {
	dbLocation, err := getConfigPath(dbURL)
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

	fmt.Println(cfg.DbURL)

	return cfg, nil
}

func getConfigPath(dbURL string) (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", dir, dbURL), nil
}