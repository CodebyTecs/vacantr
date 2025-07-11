package hh

type Vacancy struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type hhResponse struct {
	Items []Vacancy `json:"items"`
}

func FetchVacanciesMock() []Vacancy {
	return []Vacancy{
		{"Golang Developer", "https://hh.ru/vacancy/123"},
		{"Senior Go Engineer", "https://hh.ru/vacancy/456"},
	}
}
