package telegram

import (
	"log"
	"os"

	"gopkg.in/telebot.v3"
	"time"
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
		// пока заглушка
		return c.Send("Здесь будут вакансии.")
	})

	return bot
}
