package main

import (
	"log"

	"github.com/by-thoma/pocketer/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5426398167:AAEc9FgYasVVJQ59b-XqmX78FBzZL5yXKjs")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("102881-21ac15f947728b356fd644d")
	if err != nil {
		log.Fatal()
	}

	telegramBot := telegram.NewBot(bot, pocketClient, "https://localhost/")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
