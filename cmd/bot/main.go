package main

import (
	"github.com/by-thoma/pocketer/pkg/config"
	"github.com/by-thoma/pocketer/pkg/server"
	"log"

	"github.com/boltdb/bolt"
	"github.com/by-thoma/pocketer/pkg/repository"
	"github.com/by-thoma/pocketer/pkg/repository/boltdb"
	"github.com/by-thoma/pocketer/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient(cfg.PocketConsumerKey)
	if err != nil {
		log.Fatal()
	}

	db, err := initDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	tokenRepositiry := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, cfg.AuthServerURL, tokenRepositiry, cfg.Messages)

	authorizationServer := server.NewAuthorizationServer(pocketClient, tokenRepositiry, cfg.TelegramBotURL)

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB(cfg *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cfg.DBPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Batch(func(tx *bolt.Tx) error {
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
