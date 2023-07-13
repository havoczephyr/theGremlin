package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func BotRunner() error {

	config := JSONConfig{}

	err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error occured: %w", err)
	}

	session, err := discordgo.New(config.BotToken)
	if err != nil {
		return err
	}

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		botHandler(s, m, config)
	})
	return nil
}
