package app

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func BotRunner() error {
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	session, err := discordgo.New(config.BotToken)
	if err != nil {
		return err
	}

	session.AddHandler(botHandler)
	return nil
}
