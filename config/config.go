package config

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
)

var (
	Token		string
	BotPrefix	string
	Pattern		string
	ChannelID	string
	config		*configStruct
)

type configStruct struct {
	Token		string	`json : "Token"`
	BotPrefix	string	`json : "BotPrefix"`
	Pattern		string	`json : "Pattern"`
	ChannelID	string	`json : "ChannelID"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	Pattern = config.Pattern
	ChannelID = config.ChannelID

	return nil
}
