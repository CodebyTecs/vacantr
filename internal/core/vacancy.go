package core

type Vacancy struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
	URL   string `db:"url"`
}
