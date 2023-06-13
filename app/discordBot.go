package app

import (
	"github.com/bwmarrin/discordgo"
)

func discordBot() (string, error) {
	dg, err := discordgo.New("MTExODI2NDMyNjc4MDg5MTIzNw.GJEZGy.gmujkycrlutDVb7l36NkHWkNg4-wuTnk1aUoQ8")
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
