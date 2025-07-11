package usecase

import (
	"vacantr/internal/adapter/parser"
	"vacantr/internal/core"
)

type VacancyUseCase struct {
	providers []parser.VacancyProvider
}

func NewVacancyUseCase(p []parser.VacancyProvider) *VacancyUseCase {
	return &VacancyUseCase{providers: p}
}

func (v *VacancyUseCase) GetTopVacancies() []core.Vacancy {
	var result []core.Vacancy

	for _, provider := range v.providers {
		result = append(result, provider.Fetch()...)
	}

	return result
}
