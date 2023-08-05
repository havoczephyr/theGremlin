package app

import (
	"strings"
	"theGremlin/logic"

	"github.com/bwmarrin/discordgo"
)

func botHandler(s *discordgo.Session, m *discordgo.MessageCreate, config JSONConfig) {
	if m.Author.Bot || !strings.HasPrefix(m.Content, "!theGremlin") {
		return
	}
	parts := strings.Fields(m.Content[len("!theGremlin"):])
	command := parts[0]
	authorName := m.Author.Username

	port := config.ApiPort
	url := config.ApiUrl
	token := config.BotToken

	switch command {
	case "curateAllPending":
		// add me!
	case "allPendingToActive":
		// add me!
	case "showPending":
		// add me!
	case "showActive":
		// add me!
	case "showDenied":
		// add me!
	default:
		if len(parts) < 2 {
			//No arguments provided
			return
		}
	}
	arguments := strings.Join(parts[1:], " ")

	switch command {
	case "denyPost":
		// add me!
	case "suggest":
		logic.Suggest(arguments, authorName, url, token, port)
	case "createPriority":
		// add me!
	}

}
