package main

import (
	"fmt"
	"theGremlin/app"
)

func main() {
	err := app.BotRunner()
	if err != nil {
		fmt.Println("Error occured: %w", err)
		return
	}
}
