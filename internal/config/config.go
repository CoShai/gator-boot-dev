package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	file_dir, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	content, err := os.ReadFile(file_dir)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return err
	}

	file_dir, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(file_dir, jsonBytes, 0666)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	file_dir := dir + "/" + configFileName
	return file_dir, nil
}
