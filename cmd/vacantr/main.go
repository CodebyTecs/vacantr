package main

import (
	"github.com/joho/godotenv"
	"log"
	"vacantr/internal/adapter/cache"
	"vacantr/internal/adapter/kafka"
	"vacantr/internal/adapter/parser"
	"vacantr/internal/adapter/parser/habr"
	"vacantr/internal/adapter/parser/hh"
	"vacantr/internal/adapter/storage/postgres"
	"vacantr/internal/adapter/telegram"
	"vacantr/internal/usecase"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	_ = godotenv.Load()

	db := postgres.Connect()
	cache.InitRedis()
	kafka.InitKafkaWriter()

	vacancyUC := usecase.NewVacancyUseCase(db, []parser.VacancyProvider{
		hh.NewHHParser(),
		habr.NewHabrParser(),
	})

	handler := telegram.Handler{Vacancy: vacancyUC}
	bot := telegram.NewBot(handler)

	vacancyUC.GetTopVacancies(bot, vacancyUC)
	usecase.StartBackgroundParser(bot, vacancyUC)

	bot.Start()
}
