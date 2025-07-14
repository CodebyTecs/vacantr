package usecase

import (
	"github.com/jmoiron/sqlx"
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

func (v *VacancyUseCase) GetTopVacancies() []core.Vacancy {
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

	return result
}
