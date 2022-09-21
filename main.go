package main

import (
	"log"
	"golang-discord-bot/bot"
	"golang-discord-bot/config"
)


func main() {
	err := config.ReadConfig()

	if err != nil {
		log.Fatal(err)
	}

	bot.Start()

	message := make(chan struct{})
	<-message
	return
}
