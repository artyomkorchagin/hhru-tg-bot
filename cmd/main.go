package main

import (
	"log"

	"github.com/artyomkorchagin/hhru-tg-bot/internal/storage/sqlite"
	bot "github.com/artyomkorchagin/hhru-tg-bot/internal/tg-bot"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	sqlite.InitSqliteDB()
	bot.InitBot()
}
