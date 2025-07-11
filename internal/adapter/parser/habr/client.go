package habr

import "vacantr/internal/core"

type HabrParser struct{}

func NewHabrParser() *HabrParser {
	return &HabrParser{}
}

func (p *HabrParser) Fetch() []core.Vacancy {
	v := []core.Vacancy{
		{Title: "Go разработчик Habr", URL: "https://career.habr.com/vacancies/1"},
		{Title: "Backend на Go", URL: "https://career.habr.com/vacancies/2"},
	}

	return v
}
