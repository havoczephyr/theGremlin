package app

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func discordBot() {
	dg, err := discordgo.New("Bot TOKEN HERE")
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}
}
