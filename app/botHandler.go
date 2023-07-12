package app

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func botHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || !strings.HasPrefix(m.Content, "!theGremlin") {
		return
	}
	parts := strings.Split(m.Content[1:], " ")
	command := parts[0]
	// args := parts[1:]

	switch command {
	case "curateAllPending":
		// add me!
	case "declineQuestion":
		// add me!
	case "suggest":
		// add me!
	case "showPending":
		// add me!
	case "showActive":
		// add me!
	case "showDenied":
		// add me!
	}

}
