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

	return bot
}
