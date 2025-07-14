package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"vacantr/internal/core"
)

func SaveVacancy(db *sqlx.DB, u core.Vacancy) {
	_, err := db.Exec("INSERT INTO vacancies (title, url) VALUES ($1, $2)", u.Title, u.URL)
	if err != nil {
		log.Println("insert vacancy error:", err)
	}
}

func VacancyExists(db *sqlx.DB, url string) bool {
	var exists bool
	err := db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM vacancies WHERE url = $1)", url)
	if err != nil {
		log.Println("vacancy exists error:", err)
		return false
	}
	return exists
}

func GetUnseenVacancies(db *sqlx.DB, userID int64) []core.Vacancy {
	var vacancies []core.Vacancy

	query := `
		SELECT id, title, url
		FROM vacancies
		WHERE id NOT IN (
		    SELECT vacancy_id FROM user_vacancies WHERE user_id = $1
		) AND ((
		    SELECT COUNT(*) FROM user_filters WHERE user_id = $1
		) = 0
		OR title ILIKE ANY (
			SELECT unnest(array(
        		SELECT '%' || f || '%' FROM unnest(keywords) AS f
    		)) FROM user_filters WHERE user_id = $1
		))
		LIMIT 10
	`

	err := db.Select(&vacancies, query, userID)
	if err != nil {
		log.Println("getUnseenVacancies error:", err)
	}
	return vacancies
}

func MarkVacancySeen(db *sqlx.DB, userID int64, vacancyID int64) {
	_, err := db.Exec(`
		INSERT INTO user_vacancies (user_id, vacancy_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
		`, userID, vacancyID)

	if err != nil {
		log.Println("mark vacancy seen error", err)
	}
}

func SaveUserFilters(db *sqlx.DB, userID int64, filters []string) {
	_, err := db.Exec(`
		INSERT INTO user_filters (user_id, keywords)
		VALUES ($1, $2)
		ON CONFLICT (user_id) DO UPDATE SET keywords = EXCLUDED.keywords
		`, userID, pq.Array(filters))

	if err != nil {
		log.Println("save user filters error:", err)
	}
}
