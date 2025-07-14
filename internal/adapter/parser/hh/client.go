package hh

import "vacantr/internal/core"

type HHParser struct{}

func NewHHParser() *HHParser {
	return &HHParser{}
}

func (p *HHParser) Fetch() []core.Vacancy {
	v := []core.Vacancy{
		{Title: "Golang Dev", URL: "https://hh.ru/vacancy/123"},
		{Title: "Senior Go", URL: "https://hh.ru/vacancy/456"},
		{Title: "Senior Go dev", URL: "https://hh.ru/vacancy/46"},
		{Title: "Go-developer", URL: "https://hh.ru/vacancy/1"},
		{Title: "Go-developer срочно", URL: "https://hh.ru/vacancy/2"},
		{Title: "Go developer 2000$", URL: "https://hh.ru/vacancy/12"},
	}
	return v
}
