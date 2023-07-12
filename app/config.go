package app

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

func loadConfig() (*Config, error) {
	filename, err := filepath.Abs("config.json")
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err

	}
	var config Config
	err = json.Unmarshal(file, &config)
	return &config, nil
}

type Config struct {
	BotToken string `json:"botToken"`
	ApiURL   string `json:"apiUrl`
	ApiPort  string `json:"apiPort"`
}
