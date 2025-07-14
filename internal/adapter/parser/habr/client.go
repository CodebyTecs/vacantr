package habr

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"vacantr/internal/core"
)

const limitPages = 5

type HabrParser struct{}

func NewHabrParser() *HabrParser {
	return &HabrParser{}
}

func (p *HabrParser) Fetch() []core.Vacancy {
	var vacancies []core.Vacancy

	for page := 1; page <= limitPages; page++ {
		url := fmt.Sprintf("https://career.habr.com/vacancies?q=go&sort=date&type=all&page=%d", page)

		resp, _ := http.Get(url)
		defer resp.Body.Close()

		doc, _ := goquery.NewDocumentFromReader(resp.Body)

		doc.Find(".vacancy-card__title").Each(func(i int, s *goquery.Selection) {
			title := strings.TrimSpace(s.Text())
			link, _ := s.Find("a").Attr("href")
			url := "https://career.habr.com" + link

			vacancies = append(vacancies, core.Vacancy{
				Title: title,
				URL:   url,
			})
		})
	}

	return vacancies
}
