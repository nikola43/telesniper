package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
)

func main() {
	//botToken := os.Getenv("TOKEN")
	// init bot
	botToken := "5931596960:AAEpkYhCtdUj6PQhZeDjMW-QOOYJMeEVShA"
	//bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	bot, err := telego.NewBot(botToken, telego.WithWarnings())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Receiving updates
	err = HandleUpdates(bot)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
