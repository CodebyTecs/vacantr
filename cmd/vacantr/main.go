package main

import (
	"github.com/joho/godotenv"
	"log"
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
	db := postgres.Connect()

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
