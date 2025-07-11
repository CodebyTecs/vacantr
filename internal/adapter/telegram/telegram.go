package telegram

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
	"vacantr/internal/adapter/parser/hh"
)

func NewBot() *telebot.Bot {
	pref := telebot.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("failed to start bot: %s", err)
	}

	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Добро пожаловать! Отправь /vacancies, чтобы получить вакансии.")
	})

	bot.Handle("/vacancies", func(c telebot.Context) error {
		vacancies := hh.FetchVacanciesMock()
		if err != nil {
			return c.Send("Error get vacancies")
		}

		if len(vacancies) == 0 {
			return c.Send("No vacancies")
		}

		for _, v := range vacancies {
			c.Send(fmt.Sprintf("%s\n%s", v.Name, v.URL))
		}

		return nil
	})

	return bot
}
