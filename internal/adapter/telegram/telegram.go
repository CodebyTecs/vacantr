package telegram

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
	"vacantr/internal/adapter/storage/postgres"
	"vacantr/internal/core"
	"vacantr/internal/usecase"
)

type Handler struct {
	Vacancy *usecase.VacancyUseCase
}

func NewBot(handler Handler) *telebot.Bot {
	pref := telebot.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("failed to start bot: %s", err)
	}

	bot.Handle("/start", func(c telebot.Context) error {
		handler.Vacancy.SaveUser(core.User{
			TelegramID: c.Sender().ID,
			Username:   c.Sender().Username,
		})
		return c.Send("Добро пожаловать! Вакансии - /vacancies")
	})

	bot.Handle("/vacancies", func(c telebot.Context) error {
		userID := c.Sender().ID
		db := handler.Vacancy.DB()

		vacancies := postgres.GetUnseenVacancies(db, userID)

		if len(vacancies) == 0 {
			return c.Send("No vacancies")
		}

		for _, v := range vacancies {
			c.Send(fmt.Sprintf("%s\n%s", v.Title, v.URL))
			postgres.MarkVacancySeen(db, userID, v.ID)
		}

		return nil
	})

	bot.Handle("/setfilter", func(c telebot.Context) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Send("Пример: /setfilter golang junior")
		}

		handler.Vacancy.SaveFilters(c.Sender().ID, args)
		return c.Send("filters saved")
	})

	bot.Handle("/subscribe", func(c telebot.Context) error {
		db := handler.Vacancy.DB()

		_, err := db.Exec(`
			INSERT INTO subscriptions (user_id)
			VALUES ($1)
			ON CONFLICT DO NOTHING
			`, c.Sender().ID)

		if err != nil {
			return c.Send("Subscribe error")
		}

		return c.Send("Subscribe on")
	})

	bot.Handle("/unsubscribe", func(c telebot.Context) error {
		db := handler.Vacancy.DB()

		_, err := db.Exec(`
			DELETE FROM subscriptions
			WHERE user_id = $1
			`, c.Sender().ID)

		if err != nil {
			return c.Send("Unsubscribe error")
		}

		return c.Send("Subscribe off")
	})

	return bot
}
