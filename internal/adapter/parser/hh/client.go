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
	}
	return v
}
