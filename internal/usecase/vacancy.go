package usecase

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopkg.in/telebot.v3"
	"vacantr/internal/adapter/parser"
	"vacantr/internal/adapter/storage/postgres"
	"vacantr/internal/core"
)

type VacancyUseCase struct {
	providers []parser.VacancyProvider
	db        *sqlx.DB
}

func NewVacancyUseCase(db *sqlx.DB, p []parser.VacancyProvider) *VacancyUseCase {
	return &VacancyUseCase{providers: p, db: db}
}

func (v *VacancyUseCase) SaveUser(user core.User) {
	postgres.SaveUser(v.db, user)
}

func (v *VacancyUseCase) DB() *sqlx.DB {
	return v.db
}

func (v *VacancyUseCase) SaveFilters(userID int64, filters []string) {
	postgres.SaveUserFilters(v.db, userID, filters)
}

func (v *VacancyUseCase) GetTopVacancies(bot *telebot.Bot, vacancyUC *VacancyUseCase) []core.Vacancy {
	cached := GetCachedVacancies()
	if len(cached) > 0 {
		return cached
	}

	var result []core.Vacancy

	for _, provider := range v.providers {
		vacancies := provider.Fetch()
		for _, vacancy := range vacancies {
			if !postgres.VacancyExists(v.db, vacancy.URL) {
				postgres.SaveVacancy(v.db, vacancy)
				result = append(result, vacancy)
			}
		}
	}

	CacheVacancies(result)

	subscribers := postgres.GetSubscribers(vacancyUC.db)

	for _, userID := range subscribers {
		vacancies := postgres.GetUnseenVacancies(vacancyUC.db, userID)

		for _, vac := range vacancies {
			bot.Send(&telebot.User{ID: userID}, fmt.Sprintf("%s\n%s", vac.Title, vac.URL))
			postgres.MarkVacancySeen(vacancyUC.db, userID, vac.ID)
		}
	}

	return result
}
