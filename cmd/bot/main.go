package main

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/by-thoma/pocketer/pkg/repository"
	"github.com/by-thoma/pocketer/pkg/repository/boltdb"
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

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	tokenRepositiry := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, "https://localhost/", tokenRepositiry)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
