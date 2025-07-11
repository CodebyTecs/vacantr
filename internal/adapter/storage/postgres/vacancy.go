package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
	"vacantr/internal/core"
)

func SaveVacancy(db *sqlx.DB, u core.Vacancy) {
	_, err := db.Exec("INSERT INTO vacancies (title, url) VALUES ($1, $2)", u.Title, u.URL)
	if err != nil {
		log.Println("insert vacancy error:", err)
	}
}
