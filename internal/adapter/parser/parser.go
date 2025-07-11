package parser

import "vacantr/internal/core"

type VacancyProvider interface {
	Fetch() []core.Vacancy
}
