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
		{Title: "Backend на Голанг", URL: "https://career.habr.com/vacancies/56"},
		{Title: "Голанг", URL: "https://career.habr.com/vacancies/3"},
		{Title: "Голанг разработчик", URL: "https://career.habr.com/vacancies/4"},
	}

	return v
}
