package usecase

import (
	"gopkg.in/telebot.v3"
	"time"
)

func StartBackgroundParser(bot *telebot.Bot, vacancyUC *VacancyUseCase) {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			<-ticker.C
			vacancyUC.GetTopVacancies(bot, vacancyUC)
		}
	}()
}
