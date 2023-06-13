package app

import (
	"github.com/bwmarrin/discordgo"
)

func discordBot() (string, error) {
	dg, err := discordgo.New("")
	if err != nil {
		return "", err
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		return "", err
	}

	return "", nil
}
