package app

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type JSONConfig struct {
	BotToken string `json:"botToken"`
	ApiUrl   string `json:"apiUrl"`
	ApiPort  int    `json:"apiPort"`
}

func (j *JSONConfig) LoadConfig() error {
	filename, err := filepath.Abs("config.json")
	if err != nil {
		return err
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return err

	}
	err = json.Unmarshal(file, j)
	if err != nil {
		return err
	}
	return nil
}
