package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DB_url 			string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	var conf Config 
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}

func (conf *Config) SetUser(username string) error {
	conf.CurrentUserName = username
	err := write(*conf)
	return err
}

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home_dir, configFileName), nil
}

func write(conf Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(conf)
	return err
}