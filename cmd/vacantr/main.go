package main

import (
	"github.com/joho/godotenv"
	"log"
	"vacantr/internal/adapter/telegram"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	bot := telegram.NewBot()
	log.Println("Bot started")
	bot.Start()
}
