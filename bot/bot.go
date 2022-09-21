package bot

import (
	"log"
	"fmt"
	"regexp"
	"golang-discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId	string
	goBot	*discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		log.Fatal(err)
	}

	u, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
	}

	BotId = u.ID
	goBot.AddHandler(messageHandler)
	goBot.AddHandler(statusHandler)

	err = goBot.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot is running!")

}

func statusHandler(s *discordgo.Session, usd discordgo.UpdateStatusData) {
	usd.Status = "online"
	s.UpdateStatusComplex(usd)
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	matched, rerr := regexp.MatchString(config.Pattern, m.Content)
	if rerr != nil {
		log.Fatal(rerr)
	}
	
	if matched && (config.ChannelID == m.ChannelID) {
	// if matched {
		fmt.Println("matched")
		// fmt.Println(m.ChannelID)
		// fmt.Println(m.ID)
		
		merr := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if merr != nil {
			fmt.Println(merr)
			
		}

		_, merr = s.ChannelMessageSend(m.ChannelID, "Oye!!! Mera mi mejol exito! https://www.youtube.com/watch?v=kJQP7kiw5Fk")
		if merr != nil {
			fmt.Println(merr)
			
		}
		
	} 
	// else {
	// 	fmt.Println("nada")
	// 	fmt.Println(m.ID)
	// }
}
