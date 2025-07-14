package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
	"vacantr/internal/core"
)

func SaveUser(db *sqlx.DB, u core.User) {
	_, err := db.Exec("INSERT INTO users (telegram_id, username) VALUES ($1, $2)", u.TelegramID, u.Username)
	if err != nil {
		log.Println("insert user error:", err)
	}
}
